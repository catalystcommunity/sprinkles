package postgres

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
)

func (p PostgresStorage) UpsertOptionDefinitions(ctx context.Context, request *sprinklesv1.UpsertOptionDefinitionsRequest) ([]*sprinklesv1.OptionDefinition, error) {
	protos := sprinklesv1.OptionDefinitionProtos(request.OptionDefinitions)
	_, err := protos.Upsert(ctx, db)
	return protos, err
}

func (p PostgresStorage) DeleteOptionDefinitions(ctx context.Context, ids []string) error {
	return sprinklesv1.DeleteOptionDefinitionGormModels(ctx, db, ids)
}

func (p PostgresStorage) ListOptionDefinitions(ctx context.Context, request *sprinklesv1.ListRequest) ([]*sprinklesv1.OptionDefinition, error) {
	protos := sprinklesv1.OptionDefinitionProtos{}
	var orderBy string
	if orderBy = request.OrderBy; orderBy == "" {
		orderBy = "created_at"
	}
	err := protos.List(ctx, db, int(request.Limit), int(request.Offset), orderBy)
	return protos, err
}

func (p PostgresStorage) GetOptionDefinitions(ctx context.Context, request *sprinklesv1.GetRequest) ([]*sprinklesv1.OptionDefinition, error) {
	protos := sprinklesv1.OptionDefinitionProtos{}
	err := protos.GetByIds(ctx, db, request.Ids)
	return protos, err
}
