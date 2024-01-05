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

func GetOptionDefinitions(ids []string) ([]*sprinklesv1.OptionDefinition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.GetOptionDefinitions(ctx, &sprinklesv1.GetRequest{Ids: ids})
	return response.OptionDefinitions, err
}

func ListOptionDefinitions(limit, offset int, orderBy string) ([]*sprinklesv1.OptionDefinition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListOptionDefinitions(ctx, &sprinklesv1.ListRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.OptionDefinitions, err
}

func DeleteOptionDefinitions(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteOptionDefinitions(ctx, &sprinklesv1.DeleteRequest{Ids: ids})
	return err
}

func UpsertOptionDefinitions(option_definitions []*sprinklesv1.OptionDefinition) ([]*sprinklesv1.OptionDefinition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertOptionDefinitions(ctx, &sprinklesv1.UpsertOptionDefinitionsRequest{OptionDefinitions: option_definitions})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.OptionDefinitions, err
	}
}

func GetOptionOverrides(ids []string) ([]*sprinklesv1.OptionOverride, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.GetOptionOverrides(ctx, &sprinklesv1.GetRequest{Ids: ids})
	return response.OptionOverrides, err
}

func ListOptionOverrides(limit, offset int, orderBy string) ([]*sprinklesv1.OptionOverride, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListOptionOverrides(ctx, &sprinklesv1.ListRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.OptionOverrides, err
}

func DeleteOptionOverrides(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteOptionOverrides(ctx, &sprinklesv1.DeleteRequest{Ids: ids})
	return err
}

func UpsertOptionOverrides(option_overrides []*sprinklesv1.OptionOverride) ([]*sprinklesv1.OptionOverride, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertOptionOverrides(ctx, &sprinklesv1.UpsertOptionOverridesRequest{OptionOverrides: option_overrides})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.OptionOverrides, err
	}
}

func ListGroups(limit, offset int, orderBy string) ([]*sprinklesv1.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListGroups(ctx, &sprinklesv1.ListRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.Groups, err
}

func DeleteGroups(names []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteGroups(ctx, &sprinklesv1.DeleteGroupRequest{Names: names, ForceCascade: true})
	return err
}

func UpsertGroups(option_overrides []*sprinklesv1.Group) ([]*sprinklesv1.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertGroups(ctx, &sprinklesv1.UpsertGroupsRequest{Groups: option_overrides})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.Groups, err
	}
}
