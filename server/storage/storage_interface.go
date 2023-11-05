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
}
