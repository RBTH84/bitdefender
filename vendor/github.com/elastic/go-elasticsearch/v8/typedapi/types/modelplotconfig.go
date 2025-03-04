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

// ModelPlotConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/ModelPlot.ts#L23-L40
type ModelPlotConfig struct {
	// AnnotationsEnabled If true, enables calculation and storage of the model change annotations for
	// each entity that is being analyzed.
	AnnotationsEnabled *bool `json:"annotations_enabled,omitempty"`
	// Enabled If true, enables calculation and storage of the model bounds for each entity
	// that is being analyzed.
	Enabled *bool `json:"enabled,omitempty"`
	// Terms Limits data collection to this comma separated list of partition or by field
	// values. If terms are not specified or it is an empty string, no filtering is
	// applied. Wildcards are not supported. Only the specified terms can be viewed
	// when using the Single Metric Viewer.
	Terms *string `json:"terms,omitempty"`
}

// NewModelPlotConfig returns a ModelPlotConfig.
func NewModelPlotConfig() *ModelPlotConfig {
	r := &ModelPlotConfig{}

	return r
}
