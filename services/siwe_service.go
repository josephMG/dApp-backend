package services

import (
	"hardhat-backend/lib/loggers"

	siwe "github.com/spruceid/siwe-go"
)

// SiweService service layer
type SiweService struct {
	logger loggers.Logger
}

// NewSiweService creates a new userservice
func NewSiweService(
	logger loggers.Logger,
) *SiweService {
	return &SiweService{
		logger: logger,
	}
}

// Verify the messageStr and signature
func (s SiweService) Verify(messageStr string, signature string) (message *siwe.Message, err error) {
	message, err = siwe.ParseMessage(messageStr)
	if err != nil {
		return nil, err
	}

	_, err = message.Verify(signature, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return message, err
}
