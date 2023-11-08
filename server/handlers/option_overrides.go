package handlers

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/catalystcommunity/sprinkles/server/storage"
)

func (a ApiServer) GetOptionOverrides(ctx context.Context, request *sprinklesv1.GetRequest) (*sprinklesv1.OptionOverrides, error) {
	optionOverrides, err := storage.Storage.GetOptionOverrides(ctx, request)
	return &sprinklesv1.OptionOverrides{OptionOverrides: optionOverrides}, err
}

func (a ApiServer) ListOptionOverrides(ctx context.Context, request *sprinklesv1.ListRequest) (*sprinklesv1.OptionOverrides, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	optionOverrides, err := storage.Storage.ListOptionOverrides(ctx, request)
	return &sprinklesv1.OptionOverrides{OptionOverrides: optionOverrides}, err
}

func (a ApiServer) DeleteOptionOverrides(ctx context.Context, request *sprinklesv1.DeleteRequest) (*sprinklesv1.DeleteResponse, error) {
	return &sprinklesv1.DeleteResponse{}, storage.Storage.DeleteOptionOverrides(ctx, request.Ids)
}

func (a ApiServer) UpsertOptionOverrides(ctx context.Context, request *sprinklesv1.UpsertOptionOverridesRequest) (*sprinklesv1.OptionOverrides, error) {
	upsertedOptionOverrides, err := storage.Storage.UpsertOptionOverrides(ctx, request)
	return &sprinklesv1.OptionOverrides{OptionOverrides: upsertedOptionOverrides}, err
}

func (a ApiServer) GetOptionsByGroup(ctx context.Context, request *sprinklesv1.OptionGroupRequest) (*sprinklesv1.OptionOverrides, error) {
	optionOverrides, err := storage.Storage.GetOptionsByGroup(ctx, request)
	return &sprinklesv1.OptionOverrides{OptionOverrides: optionOverrides}, err
}

func (a ApiServer) GetOptionValue(ctx context.Context, request *sprinklesv1.OptionValueRequest) (*sprinklesv1.OptionValueResponse, error) {
	option_value, err := storage.Storage.GetOptionValue(ctx, request)
	return option_value, err
}
