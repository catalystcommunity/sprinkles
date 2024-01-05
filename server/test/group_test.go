package test

import (
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
)

func (s *SprinklesSuite) TestListGroups() {
	ServerGroups, err := ListGroups(1000, 0, "")
	require.Len(s.T(), ServerGroups, 0)
	groups := CreateGroups(s.T(), 3)
	require.Len(s.T(), groups, 3)
	ServerGroups, err = ListGroups(1000, 0, "")
	require.NoError(s.T(), err)
	require.Len(s.T(), ServerGroups, 0)
}

func (s *SprinklesSuite) TestUpdateGroups() {
	groups := CreateRandomNumGroups(s.T())
	randomizedGroups := randomizeGroups(s.T(), groups)
	updatedGroups, err := UpsertGroups(randomizedGroups)
	require.NoError(s.T(), err)
	assertProtoEqualitySortById(s.T(), randomizedGroups, updatedGroups,
		protocmp.IgnoreFields(&sprinklesv1.Group{}, "id", "updated_at", "name", "description", "default_value", "option_type", "schema"),
	)
}

func (s *SprinklesSuite) TestDeleteGroups() {
	Groups := CreateRandomNumGroups(s.T())
	names := lo.Map(Groups, func(item *sprinklesv1.Group, index int) string {
		return item.GroupName
	})
	err := DeleteGroups(names)
	require.NoError(s.T(), err)
	Groups, err = ListGroups(1000, 0, "")
	require.NoError(s.T(), err)
	require.Len(s.T(), Groups, 0)
}
