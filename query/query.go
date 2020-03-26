/*
	Full query parser.
*/

package wzlib_query

type WzQuery struct {
	mAny []WzQueryMatcher
	mAll []WzQueryMatcher
}

func NewWzQuery(query string) *WzQuery {
	wq := new(WzQuery)
	wq.mAny = make([]WzQueryMatcher, 0)
	wq.mAll = make([]WzQueryMatcher, 0)

	/*
		Steps to match hosts:

		1. [ ] Sort out matchers right
		2. [ ] Filter-out non-matching FQDNs
		3. [ ] Sum all up

		Dropped from Salt's query:

		1. and/or grouping. This is confusing and is not about selecting.
	*/

	wq.getMatchers()

	return wq
}

// getMatchers interprets the logical operators
// and splits matchers where they belong to.
func (wq *WzQuery) getMatchers() {

}

// Select hosts by this query
func (wq *WzQuery) Select() []string {
	return nil
}
