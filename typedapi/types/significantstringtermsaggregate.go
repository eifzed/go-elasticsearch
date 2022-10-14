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

// SignificantStringTermsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L586-L588
type SignificantStringTermsAggregate struct {
	BgCount  *int64                              `json:"bg_count,omitempty"`
	Buckets  BucketsSignificantStringTermsBucket `json:"buckets"`
	DocCount *int64                              `json:"doc_count,omitempty"`
	Meta     *Metadata                           `json:"meta,omitempty"`
}

// SignificantStringTermsAggregateBuilder holds SignificantStringTermsAggregate struct and provides a builder API.
type SignificantStringTermsAggregateBuilder struct {
	v *SignificantStringTermsAggregate
}

// NewSignificantStringTermsAggregate provides a builder for the SignificantStringTermsAggregate struct.
func NewSignificantStringTermsAggregateBuilder() *SignificantStringTermsAggregateBuilder {
	r := SignificantStringTermsAggregateBuilder{
		&SignificantStringTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the SignificantStringTermsAggregate struct
func (rb *SignificantStringTermsAggregateBuilder) Build() SignificantStringTermsAggregate {
	return *rb.v
}

func (rb *SignificantStringTermsAggregateBuilder) BgCount(bgcount int64) *SignificantStringTermsAggregateBuilder {
	rb.v.BgCount = &bgcount
	return rb
}

func (rb *SignificantStringTermsAggregateBuilder) Buckets(buckets *BucketsSignificantStringTermsBucketBuilder) *SignificantStringTermsAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *SignificantStringTermsAggregateBuilder) DocCount(doccount int64) *SignificantStringTermsAggregateBuilder {
	rb.v.DocCount = &doccount
	return rb
}

func (rb *SignificantStringTermsAggregateBuilder) Meta(meta *MetadataBuilder) *SignificantStringTermsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}