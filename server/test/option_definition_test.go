package test

import (
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
)

func (s *SprinklesSuite) TestListOptionDefinitions() {
	OptionDefinitions, err := ListOptionDefinitions(1000, 0, "")
	require.NoError(s.T(), err)
	OptionDefinitions = lo.Filter(OptionDefinitions, func(item *sprinklesv1.OptionDefinition, index int) bool {
		return LoadedTestData.OptionDefinitionIds.Contains(lo.FromPtr(item.Id))
	})
	assertProtoEqualitySortById(s.T(), LoadedTestData.OptionDefinitions, OptionDefinitions)
}

func (s *SprinklesSuite) TestGetOptionDefinitions() {
	ids := []string{}
	for _, id := range LoadedTestData.OptionDefinitionIds.Values() {
		ids = append(ids, id.(string))
	}
	OptionDefinitions, err := GetOptionDefinitions(ids)
	require.NoError(s.T(), err)
	assertProtoEqualitySortById(s.T(), LoadedTestData.OptionDefinitions, OptionDefinitions)
}

func (s *SprinklesSuite) TestUpdateOptionDefinitions() {
	optionDefinitions := CreateRandomNumOptionDefinitions(s.T())
	randomizedOptionDefinitions := randomizeOptionDefinitions(s.T(), optionDefinitions)
	updatedOptionDefinitions, err := UpsertOptionDefinitions(randomizedOptionDefinitions)
	require.NoError(s.T(), err)

	assertProtoEqualitySortById(s.T(), randomizedOptionDefinitions, updatedOptionDefinitions,
		protocmp.IgnoreFields(&sprinklesv1.OptionDefinition{}, "id", "updated_at", "name", "description", "default_value", "option_type", "schema"),
	)
}

func (s *SprinklesSuite) TestDeleteOptionDefinitions() {
	OptionDefinitions := CreateRandomNumOptionDefinitions(s.T())
	ids := lo.Map(OptionDefinitions, func(item *sprinklesv1.OptionDefinition, index int) string {
		return lo.FromPtr(item.Id)
	})
	err := DeleteOptionDefinitions(ids)
	require.NoError(s.T(), err)
	OptionDefinitions, err = GetOptionDefinitions(ids)
	require.NoError(s.T(), err)
	require.Len(s.T(), OptionDefinitions, 0)
}
