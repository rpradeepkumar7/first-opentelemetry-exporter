// Code generated by go generate gonum.org/v1/gonum/graph/formats/rdf; DO NOT EDIT.

// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rdf

import (
	"fmt"
	"net/url"
	"unicode"
)

%%{
	machine nquads;

	include "parse_actions.rl";

	include "nquads.rl";

	statement := (
	                whitespace*  subject    >StartSubject   %SetSubject
	                whitespace*  predicate  >StartPredicate %SetPredicate
	                whitespace*  object     >StartObject    %SetObject
	                (whitespace* graphLabel >StartLabel     %SetLabel)?
	                whitespace*  '.' whitespace* ('#' any*)? >Comment
	             ) %Return @!Error ;

	write data;
}%%

func parse(data []rune) (Statement, error) {
	var (
		cs, p int
		pe    = len(data)
		eof   = pe

		subject   = -1
		predicate = -1
		object    = -1
		label     = -1
		iri       = -1

		s Statement
	)

	%%write init;

	%%write exec;

	return Statement{}, ErrInvalid
}
