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

// Alias type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/indices/_types/Alias.ts#L23-L30
type Alias struct {
	Filter        *Query  `json:"filter,omitempty"`
	IndexRouting  *string `json:"index_routing,omitempty"`
	IsHidden      *bool   `json:"is_hidden,omitempty"`
	IsWriteIndex  *bool   `json:"is_write_index,omitempty"`
	Routing       *string `json:"routing,omitempty"`
	SearchRouting *string `json:"search_routing,omitempty"`
}

// NewAlias returns a Alias.
func NewAlias() *Alias {
	r := &Alias{}

	return r
}
