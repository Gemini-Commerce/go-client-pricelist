# Go API client for pricelist

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import pricelist "github.com/Gemini-Commerce/go-client-pricelist/v2"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `pricelist.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), pricelist.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `pricelist.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), pricelist.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `pricelist.ContextOperationServerIndices` and `pricelist.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), pricelist.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), pricelist.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://pricelist.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PriceListAPI* | [**CreatePriceList**](docs/PriceListAPI.md#createpricelist) | **Post** /pricelist.PriceList/CreatePriceList | Create Price List
*PriceListAPI* | [**DeletePriceListItems**](docs/PriceListAPI.md#deletepricelistitems) | **Post** /pricelist.PriceList/DeletePriceListItems | Delete Price List Items
*PriceListAPI* | [**DeletePriceListItemsAcrossPriceLists**](docs/PriceListAPI.md#deletepricelistitemsacrosspricelists) | **Post** /pricelist.PriceList/DeletePriceListItemsAcrossPriceLists | Delete Price List Items Across Price Lists
*PriceListAPI* | [**GetFullPriceItemsByPricelistId**](docs/PriceListAPI.md#getfullpriceitemsbypricelistid) | **Post** /pricelist.PriceList/GetFullPriceItemsByPricelistId | Get Full Price Items By Price List Id
*PriceListAPI* | [**GetPriceItemsByPriceListItemIds**](docs/PriceListAPI.md#getpriceitemsbypricelistitemids) | **Post** /pricelist.PriceList/GetPriceItemsByPriceListItemIds | Get Prices Items By Items Ids
*PriceListAPI* | [**GetPriceList**](docs/PriceListAPI.md#getpricelist) | **Post** /pricelist.PriceList/GetPriceList | Get Price List
*PriceListAPI* | [**GetPriceListByCode**](docs/PriceListAPI.md#getpricelistbycode) | **Post** /pricelist.PriceList/GetPriceListByCode | Get Price List By Code
*PriceListAPI* | [**GetPriceListItems**](docs/PriceListAPI.md#getpricelistitems) | **Post** /pricelist.PriceList/GetPriceListItems | Get Price List Items
*PriceListAPI* | [**GetPricesItems**](docs/PriceListAPI.md#getpricesitems) | **Post** /pricelist.PriceList/GetPricesItems | Get Prices Items
*PriceListAPI* | [**ListFullPriceItemsByPricelistId**](docs/PriceListAPI.md#listfullpriceitemsbypricelistid) | **Post** /pricelist.PriceList/ListFullPriceItemsByPricelistId | List Full Price Items By Price List Id
*PriceListAPI* | [**ListPriceLists**](docs/PriceListAPI.md#listpricelists) | **Post** /pricelist.PriceList/ListPriceLists | List Price Lists
*PriceListAPI* | [**SetPriceListItems**](docs/PriceListAPI.md#setpricelistitems) | **Post** /pricelist.PriceList/SetPriceListItems | Set Price List Items
*PriceListAPI* | [**UpdatePriceList**](docs/PriceListAPI.md#updatepricelist) | **Post** /pricelist.PriceList/UpdatePriceList | Update Price List


## Documentation For Models

 - [ListPriceListsRequestFilter](docs/ListPriceListsRequestFilter.md)
 - [ListPriceListsResponsePriceList](docs/ListPriceListsResponsePriceList.md)
 - [PricelistChannelFilter](docs/PricelistChannelFilter.md)
 - [PricelistCreatePriceListRequest](docs/PricelistCreatePriceListRequest.md)
 - [PricelistCreatePriceListResponse](docs/PricelistCreatePriceListResponse.md)
 - [PricelistCurrency](docs/PricelistCurrency.md)
 - [PricelistCurrencyFilter](docs/PricelistCurrencyFilter.md)
 - [PricelistDeletePriceListItemsAcrossPriceListsRequest](docs/PricelistDeletePriceListItemsAcrossPriceListsRequest.md)
 - [PricelistDeletePriceListItemsRequest](docs/PricelistDeletePriceListItemsRequest.md)
 - [PricelistFilterCondition](docs/PricelistFilterCondition.md)
 - [PricelistGetFullPriceItem](docs/PricelistGetFullPriceItem.md)
 - [PricelistGetFullPriceItemPrice](docs/PricelistGetFullPriceItemPrice.md)
 - [PricelistGetFullPriceItemsRequest](docs/PricelistGetFullPriceItemsRequest.md)
 - [PricelistGetFullPriceItemsResponse](docs/PricelistGetFullPriceItemsResponse.md)
 - [PricelistGetPriceItem](docs/PricelistGetPriceItem.md)
 - [PricelistGetPriceItemsByPriceListItemIdsRequest](docs/PricelistGetPriceItemsByPriceListItemIdsRequest.md)
 - [PricelistGetPriceItemsByPriceListItemIdsResponse](docs/PricelistGetPriceItemsByPriceListItemIdsResponse.md)
 - [PricelistGetPriceListByCodeRequest](docs/PricelistGetPriceListByCodeRequest.md)
 - [PricelistGetPriceListByCodeResponse](docs/PricelistGetPriceListByCodeResponse.md)
 - [PricelistGetPriceListItem](docs/PricelistGetPriceListItem.md)
 - [PricelistGetPriceListItemsRequest](docs/PricelistGetPriceListItemsRequest.md)
 - [PricelistGetPriceListItemsResponse](docs/PricelistGetPriceListItemsResponse.md)
 - [PricelistGetPriceListRequest](docs/PricelistGetPriceListRequest.md)
 - [PricelistGetPriceListResponse](docs/PricelistGetPriceListResponse.md)
 - [PricelistGetPricesRequest](docs/PricelistGetPricesRequest.md)
 - [PricelistGetPricesResponse](docs/PricelistGetPricesResponse.md)
 - [PricelistListFullPriceItemsRequest](docs/PricelistListFullPriceItemsRequest.md)
 - [PricelistListFullPriceItemsResponse](docs/PricelistListFullPriceItemsResponse.md)
 - [PricelistListPriceListsRequest](docs/PricelistListPriceListsRequest.md)
 - [PricelistListPriceListsResponse](docs/PricelistListPriceListsResponse.md)
 - [PricelistMarketFilter](docs/PricelistMarketFilter.md)
 - [PricelistMoney](docs/PricelistMoney.md)
 - [PricelistPriceContext](docs/PricelistPriceContext.md)
 - [PricelistPriceListType](docs/PricelistPriceListType.md)
 - [PricelistPriceListTypeFilter](docs/PricelistPriceListTypeFilter.md)
 - [PricelistSegmentFilter](docs/PricelistSegmentFilter.md)
 - [PricelistSetPriceListItem](docs/PricelistSetPriceListItem.md)
 - [PricelistSetPriceListItemError](docs/PricelistSetPriceListItemError.md)
 - [PricelistSetPriceListItemPrice](docs/PricelistSetPriceListItemPrice.md)
 - [PricelistSetPriceListItemsRequest](docs/PricelistSetPriceListItemsRequest.md)
 - [PricelistSetPriceListItemsResponse](docs/PricelistSetPriceListItemsResponse.md)
 - [PricelistUpdatePriceListRequest](docs/PricelistUpdatePriceListRequest.md)
 - [PricelistUpdatePriceListRequestPayload](docs/PricelistUpdatePriceListRequestPayload.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		pricelist.ContextAPIKeys,
		map[string]pricelist.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### standardAuthorization


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), pricelist.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, pricelist.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

