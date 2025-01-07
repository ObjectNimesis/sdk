package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ObjectNimesisAuth provides methods to interact with the Object Nimesis Users service.
type ObjectNimesisAuth struct {
	client      AuthServiceClient
	initialized bool
}

// NewObjectNimesisAuth creates a new ObjectNimesisUsers instance.
func NewObjectNimesisAuth(address string) (*ObjectNimesisAuth, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := NewAuthServiceClient(conn)
	return &ObjectNimesisAuth{
		client:      client,
		initialized: true,
	}, nil
}

// ensureInitialized ensures the service is initialized before making a request.
func (o *ObjectNimesisAuth) ensureInitialized() error {
	if !o.initialized {
		return errors.New("service is not initialized")
	}
	return nil
}

// ValidateToken creates a new user account.
func (o *ObjectNimesisAuth) ValidateToken(ctx context.Context, token string) (*ValidateTokenResponse, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	resp, err := o.client.ValidateToken(ctx, &ValidateTokenRequest{
		AccessToken: token,
	})

	if err != nil {
		return nil, err
	}

	return &ValidateTokenResponse{Valid: resp.Valid, UserId: resp.UserId}, nil
}
