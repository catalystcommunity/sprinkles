package test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

type TestData struct {
	Hellos              []*sprinklesv1.Hello
	HelloIds            *hashset.Set
	OptionDefinitions   []*sprinklesv1.OptionDefinition
	OptionDefinitionIds *hashset.Set
	OptionOverrides     []*sprinklesv1.OptionOverride
	OptionOverrideIds   *hashset.Set
	Groups              []*sprinklesv1.Group
	GroupNames          []string
}

var LoadedTestData = TestData{}

func loadTestData(t *testing.T) {
	deleteAllData(t)
	loadAllData(t)
}

func loadAllData(t *testing.T) {
	loadHellos(t)
	loadOptionDefinitions(t)
	loadGroups(t)
	loadOptionOverrides(t)
}

func deleteAllData(t *testing.T) {
	deleteHellos(t)
	deleteOptionOverrides(t)
	deleteGroups(t)
	deleteOptionDefinitions(t)
}

func deleteHellos(t *testing.T) {
	hellos, err := ListHellos(1000, 0, "")
	require.NoError(t, err)
	ids := lo.Map(hellos, func(item *sprinklesv1.Hello, index int) string {
		return lo.FromPtr(item.Id)
	})
	require.NoError(t, DeleteHellos(ids))
}

func loadHellos(t *testing.T) {
	hellos := CreateRandomNumHellos(t)
	LoadedTestData.Hellos = hellos
	LoadedTestData.HelloIds = hashset.New()
	for _, hello := range hellos {
		LoadedTestData.HelloIds.Add(*hello.Id)
	}
}

func CreateRandomNumHellos(t *testing.T) []*sprinklesv1.Hello {
	return CreateHellos(t, gofakeit.Number(5, 10))
}

func CreateHellos(t *testing.T, num int) []*sprinklesv1.Hello {
	var err error
	hellos := []*sprinklesv1.Hello{}
	for i := 0; i < num; i++ {
		hello := createRandomHelloProto(t)
		hellos = append(hellos, hello)
	}
	hellos, err = UpsertHellos(hellos)
	require.NoError(t, err)
	return hellos
}

func createRandomHelloProto(t *testing.T) *sprinklesv1.Hello {
	hello := &sprinklesv1.Hello{}
	err := gofakeit.Struct(hello)
	require.NoError(t, err)
	return hello
}

func randomizeHellos(t *testing.T, hellos []*sprinklesv1.Hello) []*sprinklesv1.Hello {
	randomized := []*sprinklesv1.Hello{}
	for _, hello := range hellos {
		random := createRandomHelloProto(t)
		random.Id = hello.Id
		random.UpdatedAt = hello.UpdatedAt
		random.CreatedAt = hello.CreatedAt
		randomized = append(randomized, random)
	}
	return randomized
}

func deleteOptionDefinitions(t *testing.T) {
	option_definitions, err := ListOptionDefinitions(1000, 0, "")
	require.NoError(t, err)
	ids := lo.Map(option_definitions, func(item *sprinklesv1.OptionDefinition, index int) string {
		return lo.FromPtr(item.Id)
	})
	require.NoError(t, DeleteOptionDefinitions(ids))
}

func loadOptionDefinitions(t *testing.T) {
	option_definitions := CreateRandomNumOptionDefinitions(t)
	LoadedTestData.OptionDefinitions = option_definitions
	LoadedTestData.OptionDefinitionIds = hashset.New()
	for _, option_definition := range option_definitions {
		LoadedTestData.OptionDefinitionIds.Add(*option_definition.Id)
	}
}

func CreateRandomNumOptionDefinitions(t *testing.T) []*sprinklesv1.OptionDefinition {
	return CreateOptionDefinitions(t, gofakeit.Number(5, 10))
}

func CreateOptionDefinitions(t *testing.T, num int) []*sprinklesv1.OptionDefinition {
	var err error
	option_definitions := []*sprinklesv1.OptionDefinition{}
	for i := 0; i < num; i++ {
		option_definition := createRandomOptionDefinitionProto(t)
		option_definitions = append(option_definitions, option_definition)
	}
	option_definitions, err = UpsertOptionDefinitions(option_definitions)
	require.NoError(t, err)
	return option_definitions
}

func createRandomOptionDefinitionProto(t *testing.T) *sprinklesv1.OptionDefinition {
	option_definition := &sprinklesv1.OptionDefinition{}
	err := gofakeit.Struct(option_definition)
	require.NoError(t, err)
	return option_definition
}

func randomizeOptionDefinitions(t *testing.T, option_definitions []*sprinklesv1.OptionDefinition) []*sprinklesv1.OptionDefinition {
	randomized := []*sprinklesv1.OptionDefinition{}
	for _, option_definition := range option_definitions {
		random := createRandomOptionDefinitionProto(t)
		random.Id = option_definition.Id
		random.UpdatedAt = option_definition.UpdatedAt
		random.CreatedAt = option_definition.CreatedAt
		randomized = append(randomized, random)
	}
	return randomized
}

func deleteOptionOverrides(t *testing.T) {
	option_overrides, err := ListOptionOverrides(1000, 0, "")
	require.NoError(t, err)
	ids := lo.Map(option_overrides, func(item *sprinklesv1.OptionOverride, index int) string {
		return lo.FromPtr(item.Id)
	})
	require.NoError(t, DeleteOptionOverrides(ids))
}

func loadOptionOverrides(t *testing.T) {
	option_overrides := CreateRandomNumOptionOverrides(t)
	LoadedTestData.OptionOverrides = option_overrides
	LoadedTestData.OptionOverrideIds = hashset.New()
	for _, option_override := range option_overrides {
		LoadedTestData.OptionOverrideIds.Add(*option_override.Id)
	}
}

func CreateRandomNumOptionOverrides(t *testing.T) []*sprinklesv1.OptionOverride {
	return CreateOptionOverrides(t, gofakeit.Number(5, 10))
}

func CreateOptionOverrides(t *testing.T, num int) []*sprinklesv1.OptionOverride {
	var err error
	option_overrides := []*sprinklesv1.OptionOverride{}
	for i := 0; i < num; i++ {
		option_override := createRandomOptionOverrideProto(t)
		option_overrides = append(option_overrides, option_override)
	}
	option_overrides, err = UpsertOptionOverrides(option_overrides)
	require.NoError(t, err)
	return option_overrides
}

func createRandomOptionOverrideProto(t *testing.T) *sprinklesv1.OptionOverride {
	option_override := &sprinklesv1.OptionOverride{}
	err := gofakeit.Struct(option_override)
	require.NoError(t, err)
	return option_override
}

func randomizeOptionOverrides(t *testing.T, option_overrides []*sprinklesv1.OptionOverride) []*sprinklesv1.OptionOverride {
	randomized := []*sprinklesv1.OptionOverride{}
	for _, option_override := range option_overrides {
		random := createRandomOptionOverrideProto(t)
		random.Id = option_override.Id
		random.UpdatedAt = option_override.UpdatedAt
		random.CreatedAt = option_override.CreatedAt
		randomized = append(randomized, random)
	}
	return randomized
}

func deleteGroups(t *testing.T) {
	groups, err := ListGroups(1000, 0, "")
	require.NoError(t, err)
	ids := lo.Map(groups, func(item *sprinklesv1.Group, index int) string {
		return lo.FromPtr(item.Id)
	})
	require.NoError(t, DeleteGroups(ids))
}

func loadGroups(t *testing.T) {
	groups := CreateRandomNumGroups(t)
	LoadedTestData.Groups = groups
	LoadedTestData.GroupNames = []string{}
	for _, group := range groups {
		LoadedTestData.GroupNames = append(LoadedTestData.GroupNames, group.GroupName)
	}
}

func CreateRandomNumGroups(t *testing.T) []*sprinklesv1.Group {
	return CreateGroups(t, gofakeit.Number(5, 10))
}

func CreateGroups(t *testing.T, num int) []*sprinklesv1.Group {
	var err error
	groups := []*sprinklesv1.Group{}
	for i := 0; i < num; i++ {
		group := createRandomGroupProto(t)
		groups = append(groups, group)
	}
	groups, err = UpsertGroups(groups)
	require.NoError(t, err)
	return groups
}

func createRandomGroupProto(t *testing.T) *sprinklesv1.Group {
	group := &sprinklesv1.Group{}
	err := gofakeit.Struct(group)
	require.NoError(t, err)
	return group
}

func randomizeGroups(t *testing.T, groups []*sprinklesv1.Group) []*sprinklesv1.Group {
	randomized := []*sprinklesv1.Group{}
	for _, group := range groups {
		random := createRandomGroupProto(t)
		random.GroupName = group.GroupName
		random.UpdatedAt = group.UpdatedAt
		random.CreatedAt = group.CreatedAt
		randomized = append(randomized, random)
	}
	return randomized
}
