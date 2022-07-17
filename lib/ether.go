package lib

import (
	"context"
	"hardhat-backend/config"
	"hardhat-backend/lib/loggers"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EtherClient modal
type EtherClient struct {
	*ethclient.Client
	loggers.Logger
}

// NewEther creates a new ehter client instance
func NewEther(env *config.Env, logger loggers.Logger) EtherClient {
	client, err := ethclient.Dial(env.EthereumURL)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("we have a connection")
	_ = client // we'll use this in the upcoming sections
	return EtherClient{
		Client: client,
		Logger: logger,
	}
}

// GetBalance get balance from address
func (c EtherClient) GetBalance(address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := c.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		c.Logger.Error(err.Error())
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}
