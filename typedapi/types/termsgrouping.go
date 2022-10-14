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

// TermsGrouping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/rollup/_types/Groupings.ts#L40-L42
type TermsGrouping struct {
	Fields Fields `json:"fields"`
}

// TermsGroupingBuilder holds TermsGrouping struct and provides a builder API.
type TermsGroupingBuilder struct {
	v *TermsGrouping
}

// NewTermsGrouping provides a builder for the TermsGrouping struct.
func NewTermsGroupingBuilder() *TermsGroupingBuilder {
	r := TermsGroupingBuilder{
		&TermsGrouping{},
	}

	return &r
}

// Build finalize the chain and returns the TermsGrouping struct
func (rb *TermsGroupingBuilder) Build() TermsGrouping {
	return *rb.v
}

func (rb *TermsGroupingBuilder) Fields(fields *FieldsBuilder) *TermsGroupingBuilder {
	v := fields.Build()
	rb.v.Fields = v
	return rb
}
