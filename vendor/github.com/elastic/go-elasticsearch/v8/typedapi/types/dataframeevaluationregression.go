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

// DataframeEvaluationRegression type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/DataframeEvaluation.ts#L55-L62
type DataframeEvaluationRegression struct {
	// ActualField The field of the index which contains the ground truth. The data type of this
	// field must be numerical.
	ActualField string `json:"actual_field"`
	// Metrics Specifies the metrics that are used for the evaluation. For more information
	// on mse, msle, and huber, consult the Jupyter notebook on regression loss
	// functions.
	Metrics *DataframeEvaluationRegressionMetrics `json:"metrics,omitempty"`
	// PredictedField The field in the index that contains the predicted value, in other words the
	// results of the regression analysis.
	PredictedField string `json:"predicted_field"`
}

// NewDataframeEvaluationRegression returns a DataframeEvaluationRegression.
func NewDataframeEvaluationRegression() *DataframeEvaluationRegression {
	r := &DataframeEvaluationRegression{}

	return r
}
