// core/chainio/bindings.go
package chainio

import (
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/imua-xyz/imua-avs-sdk/logging"

	// 1. Importamos nuestros bindings generados, dándoles el alias que usaremos en todo el proyecto: 'himera_avs'
	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"

	"github.com/imua-xyz/imua-avs/core/chainio/eth"
)

// 2. La struct ahora apunta a nuestro tipo de contrato HimeraAvs, usando el alias correcto.
type ContractBindings struct {
	AvsAddr    gethcommon.Address
	AVSManager *himera_avs.ContractHimeraAvs 
}

func NewContractBindings(
	avsAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*ContractBindings, error) {
	// 3. Llamamos a la función "New" de nuestro binding, no a la del ejemplo.
	contractAvsManager, err := himera_avs.NewContractHimeraAvs(avsAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch/bind to HimeraAvs contract", "err", err)
		return nil, err
	}

	return &ContractBindings{
		AvsAddr:    avsAddr,
		AVSManager: contractAvsManager,
	}, nil
}
