package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreateAt  time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		fmt.Println("amount must be greater than zero")
		return nil
	}
	a.Balance += amount
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		fmt.Println("amount must be greater than zero")
		return nil
	}
	if amount > a.Balance {
		fmt.Println("insufficient funds")
		return nil
	}
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return nil
}
