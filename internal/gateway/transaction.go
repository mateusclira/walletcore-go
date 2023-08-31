package gateway

import "github.com/mateusclira/walletcore-go/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
