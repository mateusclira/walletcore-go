package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@j")
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account1, account2, 500)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1500.0, account2.Balance)
	assert.Equal(t, 500.0, account1.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Jane Doe", "jane@j")
	account2 := NewAccount(client2)

	account1.Deposit(1000)
	account2.Deposit(1000)

	transaction, err := NewTransaction(account1, account2, 1500)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}
