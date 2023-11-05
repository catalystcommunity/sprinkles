package test

import (
	sprinklesv1 "github.com/catalystcommunity/sprinkles/protos/gen/go/sprinkles/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
	"time"
)

type SprinklesSuite struct {
	suite.Suite
}

func TestSprinklesSuite(t *testing.T) {
	suite.Run(t, new(SprinklesSuite))
}

func (s *SprinklesSuite) SetupSuite() {
	initializeApiClient(5 * time.Second)
	loadTestData(s.T())
}

func assertProtoEqualitySortById(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {
	theOpts := []cmp.Option{
		cmpopts.SortSlices(func(x, y *sprinklesv1.Hello) bool {
			return *x.Id < *y.Id
		}),
		protocmp.Transform(),
		protocmp.SortRepeated(func(x, y *sprinklesv1.Hello) bool {
			return *x.Id < *y.Id
		}),
	}
	theOpts = append(theOpts, opts...)
	diff := cmp.Diff(expected, actual, theOpts...)
	require.Empty(t, diff)
}
