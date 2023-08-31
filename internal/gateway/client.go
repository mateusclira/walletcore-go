package gateway

import "github.com/mateusclira/walletcore-go/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
