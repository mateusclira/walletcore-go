package gateway

import "github.com/mateusclira/walletcore-go/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
