package sdk

import (
	"fmt"
	"github.com/ObjectNimesis/sdk/auth"
	"github.com/ObjectNimesis/sdk/users"
)

// SDK is the aggregator for all services.
type SDK struct {
	Users *users.ObjectNimesisUsers
	Auth  *auth.ObjectNimesisAuth
}

type Service struct {
	Address string
}

type Options struct {
	Auth  Service
	Users Service
}

// NewSDK initializes and returns a new SDK instance with all services.
func NewSDK(options Options) (*SDK, error) {
	userService, err := users.NewObjectNimesisUsers(options.Users.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize user service: %w", err)
	}

	authService, err := auth.NewObjectNimesisAuth(options.Auth.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize auth service: %w", err)
	}

	return &SDK{
		Users: userService,
		Auth:  authService,
	}, nil
}
