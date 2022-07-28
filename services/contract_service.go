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

func (c ContractService) GetGreeting() string {
	return c.client.GetGreetingFromContract()
}

func (c ContractService) PostGreeting(greeting string) string {
	return c.client.PostGreeting(greeting)
}
