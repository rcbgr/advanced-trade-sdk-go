/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"net/http"

	"github.com/coinbase-samples/advanced-trade-sdk-go/credentials"
)

var defaultV3ApiBaseUrl = "https://api.coinbase.com/api/v3"

type RestClient interface {
	SetBaseUrl(u string) RestClient
	BaseUrl() string
	HttpClient() *http.Client
	Credentials() *credentials.Credentials
}

type restClientImpl struct {
	httpClient  http.Client
	credentials *credentials.Credentials
	baseUrl     string
}

func (c *restClientImpl) BaseUrl() string {
	return c.baseUrl
}

func (c *restClientImpl) SetBaseUrl(u string) RestClient {
	c.baseUrl = u
	return c
}

func (c *restClientImpl) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *restClientImpl) Credentials() *credentials.Credentials {
	return c.credentials
}

func NewRestClient(credentials *credentials.Credentials, httpClient http.Client) RestClient {
	return &restClientImpl{
		credentials: credentials,
		httpClient:  httpClient,
		baseUrl:     defaultV3ApiBaseUrl,
	}
}
