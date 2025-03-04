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

// Retrieves security privileges for the logged in user.
package getuserprivileges

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetUserPrivileges struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int
}

// NewGetUserPrivileges type alias for index.
type NewGetUserPrivileges func() *GetUserPrivileges

// NewGetUserPrivilegesFunc returns a new instance of GetUserPrivileges with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetUserPrivilegesFunc(tp elastictransport.Interface) NewGetUserPrivileges {
	return func() *GetUserPrivileges {
		n := New(tp)

		return n
	}
}

// Retrieves security privileges for the logged in user.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user-privileges.html
func New(tp elastictransport.Interface) *GetUserPrivileges {
	r := &GetUserPrivileges{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetUserPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("user")
		path.WriteString("/")
		path.WriteString("_privileges")

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetUserPrivileges) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetUserPrivileges query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getuserprivileges.Response
func (r GetUserPrivileges) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil

	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetUserPrivileges) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the GetUserPrivileges headers map.
func (r *GetUserPrivileges) Header(key, value string) *GetUserPrivileges {
	r.headers.Set(key, value)

	return r
}

// Application The name of the application. Application privileges are always associated
// with exactly one application. If you do not specify this parameter, the API
// returns information about all privileges for all applications.
// API name: application
func (r *GetUserPrivileges) Application(v string) *GetUserPrivileges {
	r.values.Set("application", v)

	return r
}

// Priviledge The name of the privilege. If you do not specify this parameter, the API
// returns information about all privileges for the requested application.
// API name: priviledge
func (r *GetUserPrivileges) Priviledge(v string) *GetUserPrivileges {
	r.values.Set("priviledge", v)

	return r
}

// API name: username
func (r *GetUserPrivileges) Username(v string) *GetUserPrivileges {
	r.values.Set("username", v)

	return r
}
