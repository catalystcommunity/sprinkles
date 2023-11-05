package test

import (
	"context"
	"time"

	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"google.golang.org/grpc"
)

var ApiClient = initializeApiClient(5 * time.Second)

func initializeApiClient(timeout time.Duration) sprinklesv1.ApiClient {
	connectTo := "127.0.0.1:6000"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return sprinklesv1.NewApiClient(conn)
}

func UpsertHellos(hellos []*sprinklesv1.Hello) ([]*sprinklesv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertHellos(ctx, &sprinklesv1.UpsertHellosRequest{Hellos: hellos})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.Hellos, err
	}
}

func DeleteHellos(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteHellos(ctx, &sprinklesv1.DeleteRequest{Ids: ids})
	return err
}

func ListHellos(limit, offset int, orderBy string) ([]*sprinklesv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListHellos(ctx, &sprinklesv1.ListRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.Hellos, err
}

func GetHellosById(ids []string) ([]*sprinklesv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.GetHellos(ctx, &sprinklesv1.GetRequest{Ids: ids})
	return response.Hellos, err
}
