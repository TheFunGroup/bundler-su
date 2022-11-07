package bundler

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stackup-wallet/stackup-bundler/pkg/mempool"
	"github.com/stackup-wallet/stackup-bundler/pkg/modules/noop"
)

// New initializes a new ERC-4337 bundler for processing UserOperations from the mempool.
func New(ethClient *ethclient.Client, mempool *mempool.Interface, chainID *big.Int, supportedEntryPoints []common.Address) *Instance {
	return &Instance{
		ethClient:            ethClient,
		mempool:              mempool,
		chainID:              chainID,
		supportedEntryPoints: supportedEntryPoints,
		batchHandler:         noop.BatchHandler,
		errorHandler:         noop.ErrorHandler,
	}
}
