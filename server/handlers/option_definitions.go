package handlers

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/catalystcommunity/sprinkles/server/storage"
)

func (a ApiServer) GetOptionDefinitions(ctx context.Context, request *sprinklesv1.GetRequest) (*sprinklesv1.OptionDefinitions, error) {
	optionDefinitions, err := storage.Storage.GetOptionDefinitions(ctx, request)
	return &sprinklesv1.OptionDefinitions{OptionDefinitions: optionDefinitions}, err
}

func (a ApiServer) ListOptionDefinitions(ctx context.Context, request *sprinklesv1.ListRequest) (*sprinklesv1.OptionDefinitions, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	optionDefinitions, err := storage.Storage.ListOptionDefinitions(ctx, request)
	return &sprinklesv1.OptionDefinitions{OptionDefinitions: optionDefinitions}, err
}

func (a ApiServer) DeleteOptionDefinitions(ctx context.Context, request *sprinklesv1.DeleteRequest) (*sprinklesv1.DeleteResponse, error) {
	return &sprinklesv1.DeleteResponse{}, storage.Storage.DeleteOptionDefinitions(ctx, request.Ids)
}

func (a ApiServer) UpsertOptionDefinitions(ctx context.Context, request *sprinklesv1.UpsertOptionDefinitionsRequest) (*sprinklesv1.OptionDefinitions, error) {
	upsertedOptionDefinitions, err := storage.Storage.UpsertOptionDefinitions(ctx, request)
	return &sprinklesv1.OptionDefinitions{OptionDefinitions: upsertedOptionDefinitions}, err
}
