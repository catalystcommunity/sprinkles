package handlers

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/catalystcommunity/sprinkles/server/storage"
)

func (a ApiServer) ListGroups(ctx context.Context, request *sprinklesv1.ListRequest) (*sprinklesv1.Groups, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	groups, err := storage.Storage.ListGroups(ctx, request)
	return groups, err
}

func (a ApiServer) DeleteGroups(ctx context.Context, request *sprinklesv1.DeleteGroupRequest) (*sprinklesv1.DeleteResponse, error) {
	return &sprinklesv1.DeleteResponse{}, storage.Storage.DeleteGroups(ctx, request)
}

func (a ApiServer) UpsertGroups(ctx context.Context, request *sprinklesv1.UpsertGroupsRequest) (*sprinklesv1.Groups, error) {
	groups, err := storage.Storage.UpsertGroups(ctx, request)
	return groups, err
}
