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
	"encoding/json"
)

// PercolateQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/specialized.ts#L110-L120
type PercolateQuery struct {
	Boost      *float32          `json:"boost,omitempty"`
	Document   json.RawMessage   `json:"document,omitempty"`
	Documents  []json.RawMessage `json:"documents,omitempty"`
	Field      string            `json:"field"`
	Id         *string           `json:"id,omitempty"`
	Index      *string           `json:"index,omitempty"`
	Name       *string           `json:"name,omitempty"`
	Preference *string           `json:"preference,omitempty"`
	QueryName_ *string           `json:"_name,omitempty"`
	Routing    *string           `json:"routing,omitempty"`
	Version    *int64            `json:"version,omitempty"`
}

// NewPercolateQuery returns a PercolateQuery.
func NewPercolateQuery() *PercolateQuery {
	r := &PercolateQuery{}

	return r
}
