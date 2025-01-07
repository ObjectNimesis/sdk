package users

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ObjectNimesisUsers provides methods to interact with the Object Nimesis Users service.
type ObjectNimesisUsers struct {
	client      UsersServiceClient
	initialized bool
}

// NewObjectNimesisUsers creates a new ObjectNimesisUsers instance.
func NewObjectNimesisUsers(address string) (*ObjectNimesisUsers, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := NewUsersServiceClient(conn)
	return &ObjectNimesisUsers{
		client:      client,
		initialized: true,
	}, nil
}

// ensureInitialized ensures the service is initialized before making a request.
func (o *ObjectNimesisUsers) ensureInitialized() error {
	if !o.initialized {
		return errors.New("service is not initialized")
	}
	return nil
}

// Create creates a new user account.
func (o *ObjectNimesisUsers) Create(ctx context.Context, username, password, email string) (*User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	resp, err := o.client.Create(ctx, &CreateRequest{
		Username: username,
		Password: password,
		Email:    email,
	})
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}

// EditById edits a user by ID.
func (o *ObjectNimesisUsers) EditById(ctx context.Context, userID uint32, data *User) (bool, error) {
	if err := o.ensureInitialized(); err != nil {
		return false, err
	}

	_, err := o.client.EditById(ctx, &EditByIdRequest{
		UserId: userID,
		Data:   data,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetById retrieves a user by their ID.
func (o *ObjectNimesisUsers) GetById(ctx context.Context, userID uint32) (*User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	resp, err := o.client.GetById(ctx, &GetByIdRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}

// GetByEmail retrieves a user by their email.
func (o *ObjectNimesisUsers) GetByEmail(ctx context.Context, email string) (*User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	resp, err := o.client.GetByEmail(ctx, &GetByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}

// GetByUsername retrieves a user by their username.
func (o *ObjectNimesisUsers) GetByUsername(ctx context.Context, username string) (*User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	resp, err := o.client.GetByUsername(ctx, &GetByUsernameRequest{
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}

// DeleteByID deletes a user by their ID.
func (o *ObjectNimesisUsers) DeleteByID(ctx context.Context, userID uint32) error {
	if err := o.ensureInitialized(); err != nil {
		return err
	}

	_, err := o.client.DeleteById(ctx, &DeleteByIdRequest{
		UserId: userID,
	})
	return err
}
