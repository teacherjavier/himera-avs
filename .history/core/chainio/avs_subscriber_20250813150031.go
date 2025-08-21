// core/chainio/avs_subscriber.go
package chainio

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"github.com/imua-xyz/imua-avs-sdk/logging"

	// --- INICIO DE LA CORRECCIÓN ---
	// Usamos la importación y el alias correctos para HIMERA
	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
	// --- FIN DE LA CORRECCIÓN ---
)

// La interfaz se actualiza para esperar un canal del tipo de evento de HIMERA.
type AvsRegistrySubscriber interface {
	SubscribeToNewTasks(newTaskCreatedChan chan *himera_avs.ContractHimeraAvsHimeraTaskCreated) (event.Subscription, error)
}

// La struct ahora almacena una instancia de nuestro contrato HimeraAvs.
type AvsRegistryChainSubscriber struct {
	logger logging.Logger
	avssub *himera_avs.ContracthelloWorld
}

var _ AvsRegistrySubscriber = (*AvsRegistryChainSubscriber)(nil)

func NewAvsRegistryChainSubscriber(
	avssub *himera_avs.ContractHimeraAvs, // Acepta el tipo de HIMERA
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	return &AvsRegistryChainSubscriber{
		logger: logger,
		avssub: avssub,
	}, nil
}

func BuildAvsRegistryChainSubscriber(
	avssubAddr common.Address,
	ethWsClient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryChainSubscriber, error) {
	// Llamamos a la función "New" de nuestro binding
	avssub, err := himera_avs.NewHimeraAvs(avssubAddr, ethWsClient)
	if err != nil {
		logger.Error("Failed to create HimeraAvs contract binding for subscriber", "err", err)
		return nil, err
	}
	return NewAvsRegistryChainSubscriber(avssub, logger)
}

// La función ahora acepta el tipo de canal correcto y llama a la función de 'Watch' correcta.
func (s *AvsRegistryChainSubscriber) SubscribeToNewTasks(newTaskCreatedChan chan *himera_avs.ContractHimeraAvsHimeraTaskCreated) (event.Subscription, error) {
	sub, err := s.avssub.WatchHimeraTaskCreated(
		&bind.WatchOpts{}, newTaskCreatedChan, nil, nil, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new HimeraTaskCreated events", "err", err)
		return nil, err
	}
	s.logger.Infof("Subscribed to new HimeraTaskCreated events")
	return sub, nil
}
