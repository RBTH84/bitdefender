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

package restore

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package restore
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/snapshot/restore/SnapshotRestoreRequest.ts#L25-L50
type Request struct {
	IgnoreIndexSettings []string             `json:"ignore_index_settings,omitempty"`
	IgnoreUnavailable   *bool                `json:"ignore_unavailable,omitempty"`
	IncludeAliases      *bool                `json:"include_aliases,omitempty"`
	IncludeGlobalState  *bool                `json:"include_global_state,omitempty"`
	IndexSettings       *types.IndexSettings `json:"index_settings,omitempty"`
	Indices             []string             `json:"indices,omitempty"`
	Partial             *bool                `json:"partial,omitempty"`
	RenamePattern       *string              `json:"rename_pattern,omitempty"`
	RenameReplacement   *string              `json:"rename_replacement,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Restore request: %w", err)
	}

	return &req, nil
}
