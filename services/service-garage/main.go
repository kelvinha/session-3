package services

import (
	"context"
	"session-3/common/models"

	"github.com/golang/protobuf/ptypes/empty"
)

var localStorage *models.GarageListBuyer

func init() {
	localStorage = new(models.GarageListBuyer)
	localStorage.List = make(map[string]*models.GarageList)
}

type GaragesServer struct{}

func (GaragesServer) Add(ctx context.Context, param *models.GarageAndUserID) (*empty.Empty, error) {
	return new(empty.Empty), nil
}
