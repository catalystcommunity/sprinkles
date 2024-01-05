package storage

import (
	"context"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
)

var Storage StorageInterface

type StorageInterface interface {
	Initialize() (shutdown func(), err error)
	Ready() bool
	UpsertHellos(ctx context.Context, request *sprinklesv1.UpsertHellosRequest) ([]*sprinklesv1.Hello, error)
	DeleteHellos(ctx context.Context, ids []string) error
	ListHellos(ctx context.Context, request *sprinklesv1.ListRequest) ([]*sprinklesv1.Hello, error)
	GetHellos(ctx context.Context, request *sprinklesv1.GetRequest) ([]*sprinklesv1.Hello, error)
	GetHello(ctx context.Context, id string) (*sprinklesv1.Hello, error)
	GetOptionDefinitions(ctx context.Context, request *sprinklesv1.GetRequest) ([]*sprinklesv1.OptionDefinition, error)
	ListOptionDefinitions(ctx context.Context, request *sprinklesv1.ListRequest) ([]*sprinklesv1.OptionDefinition, error)
	DeleteOptionDefinitions(ctx context.Context, ids []string) error
	UpsertOptionDefinitions(ctx context.Context, request *sprinklesv1.UpsertOptionDefinitionsRequest) ([]*sprinklesv1.OptionDefinition, error)
	GetOptionOverrides(ctx context.Context, request *sprinklesv1.GetRequest) ([]*sprinklesv1.OptionOverride, error)
	ListOptionOverrides(ctx context.Context, request *sprinklesv1.ListRequest) ([]*sprinklesv1.OptionOverride, error)
	DeleteOptionOverrides(ctx context.Context, ids []string) error
	UpsertOptionOverrides(ctx context.Context, request *sprinklesv1.UpsertOptionOverridesRequest) ([]*sprinklesv1.OptionOverride, error)
	GetOptionsByGroup(ctx context.Context, request *sprinklesv1.OptionGroupRequest) ([]*sprinklesv1.OptionOverride, error)
	GetOptionValue(ctx context.Context, request *sprinklesv1.OptionValueRequest) (*sprinklesv1.OptionValueResponse, error)
	ListGroups(ctx context.Context, request *sprinklesv1.ListRequest) (*sprinklesv1.Groups, error)
	DeleteGroups(ctx context.Context, request *sprinklesv1.DeleteGroupRequest) error
	UpsertGroups(ctx context.Context, request *sprinklesv1.UpsertGroupsRequest) (*sprinklesv1.Groups, error)
}
