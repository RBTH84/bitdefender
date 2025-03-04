// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestsort"
)

// TermSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/search/_types/suggester.ts#L252-L265
type TermSuggester struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Field          string                         `json:"field"`
	LowercaseTerms *bool                          `json:"lowercase_terms,omitempty"`
	MaxEdits       *int                           `json:"max_edits,omitempty"`
	MaxInspections *int                           `json:"max_inspections,omitempty"`
	MaxTermFreq    *float32                       `json:"max_term_freq,omitempty"`
	MinDocFreq     *float32                       `json:"min_doc_freq,omitempty"`
	MinWordLength  *int                           `json:"min_word_length,omitempty"`
	PrefixLength   *int                           `json:"prefix_length,omitempty"`
	ShardSize      *int                           `json:"shard_size,omitempty"`
	Size           *int                           `json:"size,omitempty"`
	Sort           *suggestsort.SuggestSort       `json:"sort,omitempty"`
	StringDistance *stringdistance.StringDistance `json:"string_distance,omitempty"`
	SuggestMode    *suggestmode.SuggestMode       `json:"suggest_mode,omitempty"`
	Text           *string                        `json:"text,omitempty"`
}

// NewTermSuggester returns a TermSuggester.
func NewTermSuggester() *TermSuggester {
	r := &TermSuggester{}

	return r
}
