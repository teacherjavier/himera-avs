package operator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// registerOperatorOnStartup registers the operator with the chain and the AVS on startup.
func (o *Operator) registerOperatorOnStartup() {
	err := o.RegisterOperatorWithChain()
	if err != nil {
		// This error might only be that the operator was already registered with chain, so we don't want to fatal
		o.logger.Error("Error registering operator with chain", "err", err)
	} else {
		o.logger.Infof("Registered operator with chain")
	}

	err = o.RegisterOperatorWithAvs()
	if err != nil {
		o.logger.Fatal("Error registering operator with avs", "err", err)
	}
}

// RegisterOperatorWithChain checks if the operator is registered with the chain.
// If the operator is not registered, it will panic.
//
// Returns:
//   - An error if the operator is not registered or if there is an error checking the registration status.
func (o *Operator) RegisterOperatorWithChain() error {
	flag, err := o.avsReader.IsOperator(&bind.CallOpts{}, o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}
	if !flag {
		o.logger.Info("Operator is not registered.")
		panic(fmt.Sprintf("Operator is not registered: %s", o.operatorAddr.String()))

	}
	return nil
}

// RegisterOperatorWithAvs registers the operator with the AVS.
// If the operator is not already opted-in, it will send a transaction to do so.
//
// Returns:
//   - An error if the registration fails.
func (o *Operator) RegisterOperatorWithAvs() error {
	operators, err := o.avsReader.GetOptInOperators(&bind.CallOpts{}, o.avsAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}
	found := false
	for _, addr := range operators {
		if addr == o.operatorAddr {
			found = true
			break
		}
	}

	if !found {
		o.logger.Info("Operator is not opt-in this avs.")
		_, err = o.avsWriter.RegisterOperatorToAVS(context.Background())
		if err != nil {
			o.logger.Error("Avs failed to RegisterOperatorToAVS", "err", err)
			return err
		}
	}

	o.logger.Info("Operator has opt-in this avs:", "avsAddr", o.avsAddr.String())

	return nil
}

// OperatorStatus represents the status of the operator.
type OperatorStatus struct {
	EcdsaAddress      string
	PubkeysRegistered bool
	Pubkey            string
	RegisteredWithAvs bool
}

// PrintOperatorStatus prints the current status of the operator.
//
// Returns:
//   - An error if the status cannot be marshaled to JSON.
func (o *Operator) PrintOperatorStatus() error {
	o.logger.Info("Printing operator status")
	operatorStatus := OperatorStatus{
		EcdsaAddress:      o.operatorAddr.String(),
		PubkeysRegistered: true,
		Pubkey:            string(o.blsKeypair.PublicKey().Marshal()),
		RegisteredWithAvs: false,
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	o.logger.Info(string(operatorStatusJson))
	return nil
}
