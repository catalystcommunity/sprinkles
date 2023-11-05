package handlers

import (
	"context"
	"errors"
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/catalystcommunity/sprinkles/server/storage"
)

type ApiServer struct{}

func (a ApiServer) Healthy(ctx context.Context, empty *sprinklesv1.Empty) (*sprinklesv1.Empty, error) {
	return &sprinklesv1.Empty{}, nil
}

func (a ApiServer) Ready(ctx context.Context, empty *sprinklesv1.Empty) (*sprinklesv1.Empty, error) {
	if storage.Storage.Ready() {
		return &sprinklesv1.Empty{}, nil
	}
	return nil, errors.New("")
}
