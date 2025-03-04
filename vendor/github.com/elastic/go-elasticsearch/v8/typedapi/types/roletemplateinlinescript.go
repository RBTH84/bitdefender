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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// RoleTemplateInlineScript type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/security/_types/Privileges.ts#L152-L157
type RoleTemplateInlineScript struct {
	Lang    *scriptlanguage.ScriptLanguage `json:"lang,omitempty"`
	Options map[string]string              `json:"options,omitempty"`
	Params  map[string]json.RawMessage     `json:"params,omitempty"`
	Source  RoleTemplateInlineQuery        `json:"source"`
}

func (s *RoleTemplateInlineScript) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "lang":
			if err := dec.Decode(&s.Lang); err != nil {
				return err
			}

		case "options":
			if err := dec.Decode(&s.Options); err != nil {
				return err
			}

		case "params":
			if err := dec.Decode(&s.Params); err != nil {
				return err
			}

		case "source":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := NewQuery()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.Source = *o

			default:
				if err := localDec.Decode(&s.Source); err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// NewRoleTemplateInlineScript returns a RoleTemplateInlineScript.
func NewRoleTemplateInlineScript() *RoleTemplateInlineScript {
	r := &RoleTemplateInlineScript{
		Options: make(map[string]string, 0),
		Params:  make(map[string]json.RawMessage, 0),
	}

	return r
}
