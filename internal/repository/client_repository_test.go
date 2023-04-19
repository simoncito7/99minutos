package repository

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateClient(t *testing.T) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=root password=root dbname=logistic_app sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)
	pass, err := hashPassword("12345")
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte("12345"))
	require.NoError(t, err)
	// create a test user
	user := Client{
		Username:  generateRandomClientID(),
		Fullname:  "Vader",
		Email:     "vader@imp.com",
		Password:  string(pass),
		CreatedAt: time.Now(),
		Token:     "token",
	}

	// insert the user into the test database using the CreateUser method
	err = repo.CreateClient(context.Background(), user)
	require.NoError(t, err)

	// fetch the user from the database
	var fetchedUser Client
	err = db.GetContext(context.Background(), &fetchedUser, "SELECT * FROM clients WHERE username = $1", user.Username)
	require.NoError(t, err)

	require.Equal(t, user.Username, fetchedUser.Username)
	require.Equal(t, user.Fullname, fetchedUser.Fullname)
	require.Equal(t, user.Email, fetchedUser.Email)
	require.Equal(t, user.Token, fetchedUser.Token)
}

func TestCreateUser_ErrorUnknownDriver(t *testing.T) {
	_, err := sqlx.Open("postgresa", "host=localhost port=5432 user=root password=root dbname=logistic_app sslmode=disable")
	require.Error(t, err)
}

func TestRepository_GetClient(t *testing.T) {
	db := createTestDB(t)
	defer db.Close()
	repo := New(db)

	expectedClient := Client{
		Username:  generateRandomClientID(),
		Fullname:  "Vader",
		Email:     "vader@imp.com",
		Password:  "12345",
		CreatedAt: time.Now(),
		Token:     "token",
	}

	// Insert the expected client into the database
	_, err := db.Exec(`INSERT INTO clients (username, fullname, email, password, created_at, token) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		expectedClient.Username, expectedClient.Fullname, expectedClient.Email,
		expectedClient.Password, expectedClient.CreatedAt, expectedClient.Token)
	require.NoError(t, err)

	// Call the GetClient function to retrieve the client from the database
	client, err := repo.GetClient(context.Background(), expectedClient.Username)
	require.NoError(t, err)

	// Compare the expected client to the actual client returned from the database
	require.Equal(t, expectedClient.Username, client.Username)
	require.Equal(t, expectedClient.Fullname, client.Fullname)
	require.Equal(t, expectedClient.Email, client.Email)
	require.Equal(t, expectedClient.Token, client.Token)
}

func TestRepository_GetClient_ClientDoesntExist(t *testing.T) {
	// set up test
	db := createTestDB(t)
	repo := New(db)

	// call function with non-existent username
	username := "nonexistent"
	client, err := repo.GetClient(context.Background(), username)
	require.EqualError(t, err, "sql: no rows in result set")
	// verify result
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	if client != (Client{}) {
		t.Errorf("Expected empty client but got %+v", client)
	}
}

func createTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=root password=root dbname=logistic_app sslmode=disable")
	require.NoError(t, err)

	return db
}
