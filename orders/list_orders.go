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

package orders

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type ListOrdersRequest struct {
	ProductId            string                  `json:"product_id,omitempty"`
	OrderStatus          []string                `json:"order_status,omitempty"`
	StartDate            string                  `json:"start_date,omitempty"`
	EndDate              string                  `json:"end_date,omitempty"`
	OrderType            string                  `json:"order_type,omitempty"`
	OrderSide            string                  `json:"order_side,omitempty"`
	ProductType          string                  `json:"product_type,omitempty"`
	OrderPlacementSource string                  `json:"order_placement_source,omitempty"`
	ContractExpiryType   string                  `json:"contract_expiry_type,omitempty"`
	AssetFilters         []string                `json:"asset_filters,omitempty"`
	RetailPortfolioId    string                  `json:"retail_portfolio_id,omitempty"`
	Pagination           *model.PaginationParams `json:"pagination,omitempty"`
}

type ListOrdersResponse struct {
	Orders     []*model.Order `json:"orders"`
	Sequence   string         `json:"sequence"`
	Pagination *model.Pagination
	Request    *ListOrdersRequest `json:"request"`
}

func (s ordersServiceImpl) ListOrders(
	ctx context.Context,
	request *ListOrdersRequest,
) (*ListOrdersResponse, error) {

	path := "/brokerage/orders/historical/batch"

	var queryParams string

	if len(request.ProductId) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_id", request.ProductId)
	}
	for _, status := range request.OrderStatus {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_status", status)
	}
	if len(request.StartDate) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "start_date", request.StartDate)
	}
	if len(request.EndDate) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_date", request.EndDate)
	}
	if len(request.OrderType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_type", request.OrderType)
	}
	if len(request.OrderSide) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_side", request.OrderSide)
	}
	if len(request.ProductType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_type", request.ProductType)
	}
	if len(request.OrderPlacementSource) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_placement_source", request.OrderPlacementSource)
	}
	if len(request.ContractExpiryType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "contract_expiry_type", request.ContractExpiryType)
	}
	for _, filter := range request.AssetFilters {
		queryParams = core.AppendHttpQueryParam(queryParams, "asset_filters", filter)
	}
	if len(request.RetailPortfolioId) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "retail_portfolio_id", request.RetailPortfolioId)
	}

	response := &ListOrdersResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
