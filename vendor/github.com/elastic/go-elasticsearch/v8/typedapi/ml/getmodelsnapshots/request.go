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

package getmodelsnapshots

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package getmodelsnapshots
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/get_model_snapshots/MlGetModelSnapshotsRequest.ts#L26-L96
type Request struct {

	// Desc Refer to the description for the `desc` query parameter.
	Desc *bool `json:"desc,omitempty"`
	// End Refer to the description for the `end` query parameter.
	End  types.DateTime `json:"end,omitempty"`
	Page *types.Page    `json:"page,omitempty"`
	// Sort Refer to the description for the `sort` query parameter.
	Sort *string `json:"sort,omitempty"`
	// Start Refer to the description for the `start` query parameter.
	Start types.DateTime `json:"start,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Getmodelsnapshots request: %w", err)
	}

	return &req, nil
}
