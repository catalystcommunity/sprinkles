package handlers

import (
	"context"
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/catalystcommunity/sprinkles/server/storage"
)

func (a ApiServer) UpsertHellos(ctx context.Context, request *sprinklesv1.UpsertHellosRequest) (*sprinklesv1.Hellos, error) {
	upsertedHellos, err := storage.Storage.UpsertHellos(ctx, request)
	return &sprinklesv1.Hellos{Hellos: upsertedHellos}, err
}

func (a ApiServer) DeleteHellos(ctx context.Context, request *sprinklesv1.DeleteRequest) (*sprinklesv1.DeleteResponse, error) {
	return &sprinklesv1.DeleteResponse{}, storage.Storage.DeleteHellos(ctx, request.Ids)
}

func (a ApiServer) ListHellos(ctx context.Context, request *sprinklesv1.ListRequest) (*sprinklesv1.Hellos, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	hellos, err := storage.Storage.ListHellos(ctx, request)
	return &sprinklesv1.Hellos{Hellos: hellos}, err
}

func (a ApiServer) GetHellos(ctx context.Context, request *sprinklesv1.GetRequest) (*sprinklesv1.Hellos, error) {
	hellos, err := storage.Storage.GetHellos(ctx, request)
	return &sprinklesv1.Hellos{Hellos: hellos}, err
}
