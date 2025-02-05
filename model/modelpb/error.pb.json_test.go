// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package modelpb

import (
	"testing"

	"github.com/elastic/apm-data/model/internal/modeljson"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestErrorToModelJSON(t *testing.T) {
	handled := true

	attrs, attrsMap := randomStruct(t)

	testCases := map[string]struct {
		proto    *Error
		expected *modeljson.Error
	}{
		"empty": {
			proto:    &Error{},
			expected: &modeljson.Error{},
		},
		"full": {
			proto: &Error{
				Exception: &Exception{
					Message:    "ex_message",
					Module:     "ex_module",
					Code:       "ex_code",
					Attributes: attrs,
					Type:       "ex_type",
					Handled:    &handled,
					Cause: []*Exception{
						{
							Message: "ex1_message",
							Module:  "ex1_module",
							Code:    "ex1_code",
							Type:    "ex_type",
						},
					},
				},
				Log: &ErrorLog{
					Message:      "log_message",
					Level:        "log_level",
					ParamMessage: "log_parammessage",
					LoggerName:   "log_loggername",
				},
				Id:          "id",
				GroupingKey: "groupingkey",
				Culprit:     "culprit",
				StackTrace:  "stacktrace",
				Message:     "message",
				Type:        "type",
			},
			expected: &modeljson.Error{
				Exception: &modeljson.Exception{
					Message:    "ex_message",
					Module:     "ex_module",
					Code:       "ex_code",
					Attributes: attrsMap,
					Type:       "ex_type",
					Handled:    &handled,
					Cause: []modeljson.Exception{
						{
							Message: "ex1_message",
							Module:  "ex1_module",
							Code:    "ex1_code",
							Type:    "ex_type",
						},
					},
				},
				Log: &modeljson.ErrorLog{
					Message:      "log_message",
					Level:        "log_level",
					ParamMessage: "log_parammessage",
					LoggerName:   "log_loggername",
				},
				ID:          "id",
				GroupingKey: "groupingkey",
				Culprit:     "culprit",
				StackTrace:  "stacktrace",
				Message:     "message",
				Type:        "type",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := modeljson.Error{
				Exception: &modeljson.Exception{},
				Log:       &modeljson.ErrorLog{},
			}
			tc.proto.toModelJSON(&out)
			diff := cmp.Diff(*tc.expected, out)
			require.Empty(t, diff)
		})
	}
}
