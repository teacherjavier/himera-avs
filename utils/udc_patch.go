package utils

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
)

// Builds the INVOKE txn function call to deploy a contract using the UDC.
//
// Parameters:
//   - classHash: the class hash of the contract to deploy
//   - constructorCalldata: the calldata to pass to the constructor of the contract
//   - opts: the options for the UDC deployment. If nil, the default options will be used.
//
// Returns:
//   - the INVOKE txn function call to deploy the contract, including the UDC
//     address and the calldata
//   - the salt used for the UDC deployment (either the provided one or the random one)
//   - an error if any
func BuildUDCCalldata(
	classHash *felt.Felt,
	constructorCalldata []*felt.Felt,
	opts *UDCOptions,
) (rpc.InvokeFunctionCall, *felt.Felt, error) {
	var result rpc.InvokeFunctionCall
	if classHash == nil {
		return result, nil, errClassHashNotProvided
	}

	if opts == nil {
		opts = new(UDCOptions)
	}

	// salt
	if opts.Salt == nil {
		randFelt := new(felt.Felt).SetRandom()
		opts.Salt = randFelt
	}

	// response
	var udcCallData []*felt.Felt
	var udcAddress *felt.Felt
	var methodName string
	var originIndFelt *felt.Felt

	switch opts.UDCVersion {
	case UDCCairoV0:
		originIndFelt = new(felt.Felt).SetUint64(1)
		if opts.OriginIndependent {
			originIndFelt.SetUint64(0)
		}

		calldataLen := new(felt.Felt).SetUint64(uint64(len(constructorCalldata)))
		udcCallData = append(
			[]*felt.Felt{classHash, opts.Salt, originIndFelt, calldataLen},
			constructorCalldata...)
		udcAddress = udcAddressCairoV0
		methodName = "deployContract"
	case UDCCairoV2:
		originIndFelt = new(felt.Felt).SetUint64(0)
		if opts.OriginIndependent {
			originIndFelt.SetUint64(1)
		}

		// The UDCCairoV2 `calldata` constructor parameter is of type `Span<felt>`, so
		// we need to pass the length of the array as the first element.
		// ref: https://book.cairo-lang.org/ch102-04-serialization-of-cairo-types.html#serialization-of-arrays-and-spans
		//nolint:lll // The link would be unclickable if we break the line.
		constructorCalldata = append(
			[]*felt.Felt{new(felt.Felt).SetUint64(uint64(len(constructorCalldata)))},
			constructorCalldata...)

		udcCallData = append(
			[]*felt.Felt{classHash, opts.Salt, originIndFelt},
			constructorCalldata...)
		udcAddress = udcAddressCairoV2
		methodName = "deploy_contract"
	default:
		return result, nil, errInvalidUDCVersion
	}

	result = rpc.InvokeFunctionCall{
		ContractAddress: udcAddress,
		FunctionName:    methodName,
		CallData:        udcCallData,
	}

	return result, opts.Salt, nil
}
