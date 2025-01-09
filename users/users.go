package users

import (
	"context"
	"errors"
	"github.com/ObjectNimesis/sdk/models"

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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

// convertUser converts a User to the SDK's User model.
func convertUser(user *User) *models.User {
	var emails []models.Email
	for _, email := range user.Emails {
		emails = append(emails, models.Email{
			Address:  email.Address,
			Type:     models.EmailAddressType(EmailAddressType_name[int32(email.Type)]),
			Verified: email.Verified,
			UserID:   email.UserId,
		})
	}

	return &models.User{
		ID:          user.Id,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Password:    user.Password,
		Gender:      user.Gender,
		Pronouns:    user.Pronouns,
		Avatar:      user.Avatar,
		Emails:      emails,
	}
}

// Create creates a new user account.
func (o *ObjectNimesisUsers) Create(ctx context.Context, username, password, email string) (*models.User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	req := &CreateUserRequest{
		Username: username,
		Password: password,
		Email:    email,
	}

	resp, err := o.client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return convertUser(resp.User), nil
}

// EditByID edits a user by their ID.
func (o *ObjectNimesisUsers) EditByID(ctx context.Context, userID uint32, userData *User) (bool, error) {
	if err := o.ensureInitialized(); err != nil {
		return false, err
	}

	req := &EditUserByIdRequest{
		UserId:   userID,
		UserData: userData,
	}

	_, err := o.client.EditUserById(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetByID retrieves a user by their ID.
func (o *ObjectNimesisUsers) GetByID(ctx context.Context, userID uint32) (*models.User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	req := &GetUserByIdRequest{
		UserId: userID,
	}

	resp, err := o.client.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}

	return convertUser(resp.User), nil
}

// GetByEmail retrieves a user by their email.
func (o *ObjectNimesisUsers) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	req := &GetUserByEmailRequest{
		Email: email,
	}

	resp, err := o.client.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return convertUser(resp.User), nil
}

// GetByUsername retrieves a user by their username.
func (o *ObjectNimesisUsers) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	if err := o.ensureInitialized(); err != nil {
		return nil, err
	}

	req := &GetUserByUsernameRequest{
		Username: username,
	}

	resp, err := o.client.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, err
	}

	return convertUser(resp.User), nil
}

// DeleteByID deletes a user by their ID.
func (o *ObjectNimesisUsers) DeleteByID(ctx context.Context, userID uint32) (bool, error) {
	if err := o.ensureInitialized(); err != nil {
		return false, err
	}

	req := &DeleteUserByIdRequest{
		UserId: userID,
	}

	resp, err := o.client.DeleteUserById(ctx, req)
	if err != nil {
		return false, err
	}

	return resp.Success, nil
}
