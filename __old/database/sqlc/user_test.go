package database

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lucianocorreia/go-starter/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotNil(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.IsActive, user2.IsActive)
	require.Equal(t, user1.TenantID, user2.TenantID)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func createRandomUser(t *testing.T) *User {
	hashaedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	params := &CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashaedPassword,
		Name:           utils.RandomOwner(),
		Email:          utils.RandomEmail(),
		IsActive:       false,
		TenantID:       pgtype.Text{String: uuid.NewString(), Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, user)

	require.Equal(t, params.Username, user.Username)
	require.Equal(t, params.HashedPassword, user.HashedPassword)
	require.Equal(t, params.Name, user.Name)
	require.Equal(t, params.Email, user.Email)
	require.Equal(t, params.IsActive, user.IsActive)
	require.Equal(t, params.TenantID, user.TenantID)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}
