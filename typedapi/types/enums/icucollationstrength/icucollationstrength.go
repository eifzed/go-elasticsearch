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


// Package icucollationstrength
package icucollationstrength

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/analysis/icu-plugin.ts#L104-L110
type IcuCollationStrength struct {
	name string
}

var (
	Primary = IcuCollationStrength{"primary"}

	Secondary = IcuCollationStrength{"secondary"}

	Tertiary = IcuCollationStrength{"tertiary"}

	Quaternary = IcuCollationStrength{"quaternary"}

	Identical = IcuCollationStrength{"identical"}
)

func (i IcuCollationStrength) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IcuCollationStrength) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "primary":
		*i = Primary
	case "secondary":
		*i = Secondary
	case "tertiary":
		*i = Tertiary
	case "quaternary":
		*i = Quaternary
	case "identical":
		*i = Identical
	default:
		*i = IcuCollationStrength{string(text)}
	}

	return nil
}

func (i IcuCollationStrength) String() string {
	return i.name
}
