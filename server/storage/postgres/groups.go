package postgres

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
)

func (p PostgresStorage) ListGroups(ctx context.Context, request *sprinklesv1.ListRequest) (*sprinklesv1.Groups, error) {
	protos := sprinklesv1.GroupProtos{}
	var orderBy string
	if orderBy = request.OrderBy; orderBy == "" {
		orderBy = "created_at"
	}
	err := protos.List(ctx, db, int(request.Limit), int(request.Offset), orderBy)
	return &sprinklesv1.Groups{Groups: protos}, err
}

func (p PostgresStorage) DeleteGroups(ctx context.Context, request *sprinklesv1.DeleteGroupRequest) error {
	return sprinklesv1.DeleteGroupsByNames(ctx, db, request.Names, request.ForceCascade)
}

func (p PostgresStorage) UpsertGroups(ctx context.Context, request *sprinklesv1.UpsertGroupsRequest) (*sprinklesv1.Groups, error) {
	groupProtos := sprinklesv1.GroupProtos(request.Groups)
	_, err := groupProtos.Upsert(ctx, db)
	return &sprinklesv1.Groups{Groups: groupProtos}, err
}
