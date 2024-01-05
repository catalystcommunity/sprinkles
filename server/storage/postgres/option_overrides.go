package postgres

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
)

func (p PostgresStorage) GetOptionOverrides(ctx context.Context, request *sprinklesv1.GetRequest) ([]*sprinklesv1.OptionOverride, error) {
	protos := sprinklesv1.OptionOverrideProtos{}
	err := protos.GetByIds(ctx, db, request.Ids)
	return protos, err
}

func (p PostgresStorage) ListOptionOverrides(ctx context.Context, request *sprinklesv1.ListRequest) ([]*sprinklesv1.OptionOverride, error) {
	protos := sprinklesv1.OptionOverrideProtos{}
	var orderBy string
	if orderBy = request.OrderBy; orderBy == "" {
		orderBy = "created_at"
	}
	err := protos.List(ctx, db, int(request.Limit), int(request.Offset), orderBy)
	return protos, err
}

func (p PostgresStorage) DeleteOptionOverrides(ctx context.Context, ids []string) error {
	return sprinklesv1.DeleteOptionOverrideGormModels(ctx, db, ids)
}

func (p PostgresStorage) UpsertOptionOverrides(ctx context.Context, request *sprinklesv1.UpsertOptionOverridesRequest) ([]*sprinklesv1.OptionOverride, error) {
	helloProtos := sprinklesv1.OptionOverrideProtos(request.OptionOverrides)
	_, err := helloProtos.Upsert(ctx, db)
	return helloProtos, err
}

func (p PostgresStorage) GetOptionsByGroup(ctx context.Context, request *sprinklesv1.OptionGroupRequest) ([]*sprinklesv1.OptionOverride, error) {
	protos := sprinklesv1.OptionOverrideProtos{}
	err := protos.GetOptionsByGroup(ctx, db, request.Groups)
	return protos, err
}

func (p PostgresStorage) GetOptionValue(ctx context.Context, request *sprinklesv1.OptionValueRequest) (*sprinklesv1.OptionValueResponse, error) {
	protos := sprinklesv1.OptionOverrideProtos{}
	option_value, err := protos.GetOptionValue(ctx, db, request.IgnoreGroups)
	return &sprinklesv1.OptionValueResponse{OptionValue: option_value}, err
}
