package test

import (
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
)

func (s *SprinklesSuite) TestListOptionOverrides() {
	OptionOverrides, err := ListOptionOverrides(1000, 0, "")
	require.NoError(s.T(), err)
	OptionOverrides = lo.Filter(OptionOverrides, func(item *sprinklesv1.OptionOverride, index int) bool {
		return LoadedTestData.OptionOverrideIds.Contains(lo.FromPtr(item.Id))
	})
	assertProtoEqualitySortById(s.T(), LoadedTestData.OptionOverrides, OptionOverrides)
}

func (s *SprinklesSuite) TestGetOptionOverrides() {
	ids := []string{}
	for _, id := range LoadedTestData.OptionOverrideIds.Values() {
		ids = append(ids, id.(string))
	}
	OptionOverrides, err := GetOptionOverrides(ids)
	require.NoError(s.T(), err)
	assertProtoEqualitySortById(s.T(), LoadedTestData.OptionOverrides, OptionOverrides)
}

func (s *SprinklesSuite) TestUpdateOptionOverrides() {
	optionOverrides := CreateRandomNumOptionOverrides(s.T())
	randomizedOptionOverrides := randomizeOptionOverrides(s.T(), optionOverrides)
	updatedOptionOverrides, err := UpsertOptionOverrides(randomizedOptionOverrides)
	require.NoError(s.T(), err)

	assertProtoEqualitySortById(s.T(), randomizedOptionOverrides, updatedOptionOverrides,
		protocmp.IgnoreFields(&sprinklesv1.OptionOverride{}, "id", "updated_at", "option_value", "group"),
	)
}

func (s *SprinklesSuite) TestDeleteOptionOverrides() {
	OptionOverrides := CreateRandomNumOptionOverrides(s.T())
	ids := lo.Map(OptionOverrides, func(item *sprinklesv1.OptionOverride, index int) string {
		return lo.FromPtr(item.Id)
	})
	err := DeleteOptionOverrides(ids)
	require.NoError(s.T(), err)
	OptionOverrides, err = GetOptionOverrides(ids)
	require.NoError(s.T(), err)
	require.Len(s.T(), OptionOverrides, 0)
}
