/*
	Matcher is a sub-query or a part of the entire query.
	Query on its own can consist of several queries
	and have logical expressions, like "and", "or".
*/

package wzlib_query

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	Q_FH = iota
	Q_FTV
	Q_TV
	Q_H
)

type WzQueryMatcher struct {
	query    string
	fqdns    []string // hostnames
	flags    []string
	traitKey string
	qType    int
}

// Constructor
func NewWzQueryMatcher(q string) *WzQueryMatcher {
	wqm := new(WzQueryMatcher)
	wqm.query = q
	wqm.parseQuery()
	return wqm
}

func (wqm *WzQueryMatcher) parseQuery() {
	// Remove alias
	if wqm.query == "a::" || wqm.query == "a:" || wqm.query == "all" {
		wqm.query = "*"
	}

	if wqm.query == "*" {
		return
	}

	parts := strings.Split(wqm.query, ":")
	if len(parts) == 3 {
		// Get all flags
		for _, flag := range parts[0] {
			wqm.flags = append(wqm.flags, string(flag))
		}

		// Set query type
		if parts[1] == "" {
			// flag::hostname
			wqm.qType = Q_FH
			wqm.fqdns = wqm.getHostnames(parts[2])
		} else {
			// flag:trait:value
			wqm.qType = Q_FTV
		}

	} else if len(parts) == 2 {
		// trait:value
		wqm.qType = Q_TV
	} else {
		// hostname
		wqm.qType = Q_H
		wqm.fqdns = wqm.getHostnames(parts[0])
	}
}

// Get hostnames from the expression
func (wqm *WzQueryMatcher) getHostnames(expr string) []string {
	if expr == "" {
		return nil
	}

	return strings.Split(strings.ReplaceAll(expr, " ", ""), ",")
}

/*
	Query has a flag against a hostname expression.
	This expects the following syntax:

		flag::hostname
*/
func (wqm *WzQueryMatcher) parseFlagToHostname(parts []string) {
	fmt.Println("Flag to hostname")
}

/*
	Query should have a flag against a trait and its value.
	This expects the following syntax:

		flag:trait:value
*/
func (wqm *WzQueryMatcher) parseFlagToTraitValue(parts []string) {
	fmt.Println("Flag to trait value")
}

/*
	Query describes a trait key and its value.
	This expects the following syntax:

		trait:value
*/
func (wqm *WzQueryMatcher) parseTraitValue(parts []string) {
	fmt.Println("Plain trait value")
}

/*
	Query describes an expression of the hostname.
	This expects the following syntax:

		hostname
*/
func (wqm *WzQueryMatcher) parseHostnameExpression() {
	fmt.Println("Hostname expression")
}

// matchByHostname matches the hostname if the query type is for matching FQDNs only
func (wqm *WzQueryMatcher) filterFqdns(fqdns []string) []string {
	out := make([]string, 0)

	for _, fqdn := range fqdns {
		for _, queryFqdn := range wqm.fqdns {
			result, _ := filepath.Match(queryFqdn, fqdn)
			if result {
				out = append(out, fqdn)
			}
		}
	}

	return out
}

// Select from an array of fqdns
func (wqm *WzQueryMatcher) Select(fqdns []string) []string {
	switch wqm.qType {
	case Q_FH, Q_H:
		return wqm.filterFqdns(fqdns)
	default:
		return make([]string, 0)
	}
}
