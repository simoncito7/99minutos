package repository

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateClient(t *testing.T) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=root password=root dbname=logistic_app sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	// create a test user
	user := Client{
		Name:      "Darth",
		LastName:  "Vader",
		Email:     "vader@imp.com",
		Password:  "12345",
		CreatedAt: time.Now(),
		Token:     "token",
	}

	// insert the user into the test database using the CreateUser method
	err = repo.CreateClient(context.Background(), db, user)
	require.NoError(t, err)

	// fetch the user from the database
	var fetchedUser Client
	err = db.GetContext(context.Background(), &fetchedUser, "SELECT * FROM client WHERE email = $1", user.Email)
	require.NoError(t, err)

	require.Equal(t, user.Name, fetchedUser.Name)
	require.Equal(t, user.LastName, fetchedUser.LastName)
	require.Equal(t, user.Email, fetchedUser.Email)
	require.Equal(t, user.Password, fetchedUser.Password)
	require.Equal(t, user.Token, fetchedUser.Token)
}

func TestCreateUser_ErrorUnknownDriver(t *testing.T) {
	_, err := sqlx.Open("postgresa", "host=localhost port=5432 user=root password=root dbname=logistic_app sslmode=disable")
	require.Error(t, err)
}

func createTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=root password=root dbname=logistic_app sslmode=disable")
	require.NoError(t, err)

	return db
}
