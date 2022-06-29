package services

import (
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"
	"math/big"
)

// ContractService service layer
type ContractService struct {
	logger loggers.Logger
	client lib.EtherClient
}

// NewContractService creates a new userservice
func NewContractService(logger loggers.Logger, client lib.EtherClient) ContractService {
	return ContractService{
		logger: logger,
		client: client,
	}
}

func (c ContractService) GetBalance(address string) *big.Float {
	return c.client.GetBalance(address)
}
