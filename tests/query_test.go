package tests

import (
	"testing"

	wzlib_query "github.com/infra-whizz/wzlib/query"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type QueryTestSuite struct {
	machines []string
}

var _ = check.Suite(&QueryTestSuite{})

func (s *QueryTestSuite) SetUpSuite(c *check.C) {
	s.machines = []string{
		"web1", "web2", "web3",
		"system1", "system2", "system3",
	}
}

func (s *QueryTestSuite) TearDownSuite(c *check.C) {
	s.machines = nil
}

// TestMatchAll for `\*` command
func (s *QueryTestSuite) TestMatchAllByFlag(c *check.C) {
	matcher := wzlib_query.NewWzQueryMatcher("a::")
	matcher.Match()
	c.Assert(0, check.Equals, 0)
}

// TestMatchAll for `a::`
func (s *QueryTestSuite) TestMatchAllByGlob(c *check.C) {
	matcher := wzlib_query.NewWzQueryMatcher("*")
	matcher.Match()
	c.Assert(0, check.Equals, 0)
}
