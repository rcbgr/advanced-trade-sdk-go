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

package fees

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
)

type FeesService interface {
	GetTransactionsSummary(ctx context.Context, request *GetTransactionsSummaryRequest) (*GetTransactionsSummaryResponse, error)
}

func NewFeesService(c client.RestClient) FeesService {
	return &feesServiceImpl{client: c}
}

type feesServiceImpl struct {
	client client.RestClient
}
