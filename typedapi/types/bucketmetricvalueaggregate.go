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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// BucketMetricValueAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L224-L227
type BucketMetricValueAggregate struct {
	Keys []string  `json:"keys"`
	Meta *Metadata `json:"meta,omitempty"`
	// Value The metric value. A missing value generally means that there was no data to
	// aggregate,
	// unless specified otherwise.
	Value         float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

// BucketMetricValueAggregateBuilder holds BucketMetricValueAggregate struct and provides a builder API.
type BucketMetricValueAggregateBuilder struct {
	v *BucketMetricValueAggregate
}

// NewBucketMetricValueAggregate provides a builder for the BucketMetricValueAggregate struct.
func NewBucketMetricValueAggregateBuilder() *BucketMetricValueAggregateBuilder {
	r := BucketMetricValueAggregateBuilder{
		&BucketMetricValueAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the BucketMetricValueAggregate struct
func (rb *BucketMetricValueAggregateBuilder) Build() BucketMetricValueAggregate {
	return *rb.v
}

func (rb *BucketMetricValueAggregateBuilder) Keys(keys ...string) *BucketMetricValueAggregateBuilder {
	rb.v.Keys = keys
	return rb
}

func (rb *BucketMetricValueAggregateBuilder) Meta(meta *MetadataBuilder) *BucketMetricValueAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

// Value The metric value. A missing value generally means that there was no data to
// aggregate,
// unless specified otherwise.

func (rb *BucketMetricValueAggregateBuilder) Value(value float64) *BucketMetricValueAggregateBuilder {
	rb.v.Value = value
	return rb
}

func (rb *BucketMetricValueAggregateBuilder) ValueAsString(valueasstring string) *BucketMetricValueAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
