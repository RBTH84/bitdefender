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

// ZeroShotClassificationInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/inference.ts#L339-L348
type ZeroShotClassificationInferenceUpdateOptions struct {
	// Labels The labels to predict.
	Labels []string `json:"labels"`
	// MultiLabel Update the configured multi label option. Indicates if more than one true
	// label exists. Defaults to the configured value.
	MultiLabel *bool `json:"multi_label,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// NewZeroShotClassificationInferenceUpdateOptions returns a ZeroShotClassificationInferenceUpdateOptions.
func NewZeroShotClassificationInferenceUpdateOptions() *ZeroShotClassificationInferenceUpdateOptions {
	r := &ZeroShotClassificationInferenceUpdateOptions{}

	return r
}
