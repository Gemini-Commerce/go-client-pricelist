# \PriceListAPI

All URIs are relative to *https://pricelist.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePriceList**](PriceListAPI.md#CreatePriceList) | **Post** /pricelist.PriceList/CreatePriceList | Create Price List
[**DeletePriceListItems**](PriceListAPI.md#DeletePriceListItems) | **Post** /pricelist.PriceList/DeletePriceListItems | Delete Price List Items
[**DeletePriceListItemsAcrossPriceLists**](PriceListAPI.md#DeletePriceListItemsAcrossPriceLists) | **Post** /pricelist.PriceList/DeletePriceListItemsAcrossPriceLists | Delete Price List Items Across Price Lists
[**GetFullPriceItemsByPricelistId**](PriceListAPI.md#GetFullPriceItemsByPricelistId) | **Post** /pricelist.PriceList/GetFullPriceItemsByPricelistId | Get Full Price Items By Price List Id
[**GetPriceItemsByPriceListItemIds**](PriceListAPI.md#GetPriceItemsByPriceListItemIds) | **Post** /pricelist.PriceList/GetPriceItemsByPriceListItemIds | Get Prices Items By Items Ids
[**GetPriceList**](PriceListAPI.md#GetPriceList) | **Post** /pricelist.PriceList/GetPriceList | Get Price List
[**GetPriceListByCode**](PriceListAPI.md#GetPriceListByCode) | **Post** /pricelist.PriceList/GetPriceListByCode | Get Price List By Code
[**GetPriceListItems**](PriceListAPI.md#GetPriceListItems) | **Post** /pricelist.PriceList/GetPriceListItems | Get Price List Items
[**GetPricesItems**](PriceListAPI.md#GetPricesItems) | **Post** /pricelist.PriceList/GetPricesItems | Get Prices Items
[**ListFullPriceItemsByPricelistId**](PriceListAPI.md#ListFullPriceItemsByPricelistId) | **Post** /pricelist.PriceList/ListFullPriceItemsByPricelistId | List Full Price Items By Price List Id
[**ListPriceLists**](PriceListAPI.md#ListPriceLists) | **Post** /pricelist.PriceList/ListPriceLists | List Price Lists
[**SetPriceListItems**](PriceListAPI.md#SetPriceListItems) | **Post** /pricelist.PriceList/SetPriceListItems | Set Price List Items
[**UpdatePriceList**](PriceListAPI.md#UpdatePriceList) | **Post** /pricelist.PriceList/UpdatePriceList | Update Price List



## CreatePriceList

> PricelistCreatePriceListResponse CreatePriceList(ctx).Body(body).Execute()

Create Price List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistCreatePriceListRequest() // PricelistCreatePriceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.CreatePriceList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.CreatePriceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePriceList`: PricelistCreatePriceListResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.CreatePriceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePriceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistCreatePriceListRequest**](PricelistCreatePriceListRequest.md) |  | 

### Return type

[**PricelistCreatePriceListResponse**](PricelistCreatePriceListResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriceListItems

> map[string]interface{} DeletePriceListItems(ctx).Body(body).Execute()

Delete Price List Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistDeletePriceListItemsRequest() // PricelistDeletePriceListItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.DeletePriceListItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.DeletePriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePriceListItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.DeletePriceListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistDeletePriceListItemsRequest**](PricelistDeletePriceListItemsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriceListItemsAcrossPriceLists

> map[string]interface{} DeletePriceListItemsAcrossPriceLists(ctx).Body(body).Execute()

Delete Price List Items Across Price Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistDeletePriceListItemsAcrossPriceListsRequest() // PricelistDeletePriceListItemsAcrossPriceListsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.DeletePriceListItemsAcrossPriceLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.DeletePriceListItemsAcrossPriceLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePriceListItemsAcrossPriceLists`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.DeletePriceListItemsAcrossPriceLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriceListItemsAcrossPriceListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistDeletePriceListItemsAcrossPriceListsRequest**](PricelistDeletePriceListItemsAcrossPriceListsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullPriceItemsByPricelistId

> PricelistGetFullPriceItemsResponse GetFullPriceItemsByPricelistId(ctx).Body(body).Execute()

Get Full Price Items By Price List Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistGetFullPriceItemsRequest() // PricelistGetFullPriceItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.GetFullPriceItemsByPricelistId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.GetFullPriceItemsByPricelistId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFullPriceItemsByPricelistId`: PricelistGetFullPriceItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.GetFullPriceItemsByPricelistId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFullPriceItemsByPricelistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetFullPriceItemsRequest**](PricelistGetFullPriceItemsRequest.md) |  | 

### Return type

[**PricelistGetFullPriceItemsResponse**](PricelistGetFullPriceItemsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceItemsByPriceListItemIds

> PricelistGetPriceItemsByPriceListItemIdsResponse GetPriceItemsByPriceListItemIds(ctx).Body(body).Execute()

Get Prices Items By Items Ids



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceItemsByPriceListItemIdsRequest() // PricelistGetPriceItemsByPriceListItemIdsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.GetPriceItemsByPriceListItemIds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.GetPriceItemsByPriceListItemIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceItemsByPriceListItemIds`: PricelistGetPriceItemsByPriceListItemIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.GetPriceItemsByPriceListItemIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceItemsByPriceListItemIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceItemsByPriceListItemIdsRequest**](PricelistGetPriceItemsByPriceListItemIdsRequest.md) |  | 

### Return type

[**PricelistGetPriceItemsByPriceListItemIdsResponse**](PricelistGetPriceItemsByPriceListItemIdsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceList

> PricelistGetPriceListResponse GetPriceList(ctx).Body(body).Execute()

Get Price List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceListRequest() // PricelistGetPriceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.GetPriceList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.GetPriceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceList`: PricelistGetPriceListResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.GetPriceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceListRequest**](PricelistGetPriceListRequest.md) |  | 

### Return type

[**PricelistGetPriceListResponse**](PricelistGetPriceListResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListByCode

> PricelistGetPriceListByCodeResponse GetPriceListByCode(ctx).Body(body).Execute()

Get Price List By Code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceListByCodeRequest() // PricelistGetPriceListByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.GetPriceListByCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.GetPriceListByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListByCode`: PricelistGetPriceListByCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.GetPriceListByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceListByCodeRequest**](PricelistGetPriceListByCodeRequest.md) |  | 

### Return type

[**PricelistGetPriceListByCodeResponse**](PricelistGetPriceListByCodeResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListItems

> PricelistGetPriceListItemsResponse GetPriceListItems(ctx).Body(body).Execute()

Get Price List Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceListItemsRequest() // PricelistGetPriceListItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.GetPriceListItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.GetPriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListItems`: PricelistGetPriceListItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.GetPriceListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceListItemsRequest**](PricelistGetPriceListItemsRequest.md) |  | 

### Return type

[**PricelistGetPriceListItemsResponse**](PricelistGetPriceListItemsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricesItems

> PricelistGetPricesResponse GetPricesItems(ctx).Body(body).Execute()

Get Prices Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistGetPricesRequest() // PricelistGetPricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.GetPricesItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.GetPricesItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricesItems`: PricelistGetPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.GetPricesItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPricesRequest**](PricelistGetPricesRequest.md) |  | 

### Return type

[**PricelistGetPricesResponse**](PricelistGetPricesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFullPriceItemsByPricelistId

> PricelistListFullPriceItemsResponse ListFullPriceItemsByPricelistId(ctx).Body(body).Execute()

List Full Price Items By Price List Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistListFullPriceItemsRequest() // PricelistListFullPriceItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.ListFullPriceItemsByPricelistId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.ListFullPriceItemsByPricelistId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFullPriceItemsByPricelistId`: PricelistListFullPriceItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.ListFullPriceItemsByPricelistId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFullPriceItemsByPricelistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistListFullPriceItemsRequest**](PricelistListFullPriceItemsRequest.md) |  | 

### Return type

[**PricelistListFullPriceItemsResponse**](PricelistListFullPriceItemsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPriceLists

> PricelistListPriceListsResponse ListPriceLists(ctx).Body(body).Execute()

List Price Lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistListPriceListsRequest() // PricelistListPriceListsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.ListPriceLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.ListPriceLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPriceLists`: PricelistListPriceListsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.ListPriceLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPriceListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistListPriceListsRequest**](PricelistListPriceListsRequest.md) |  | 

### Return type

[**PricelistListPriceListsResponse**](PricelistListPriceListsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPriceListItems

> PricelistSetPriceListItemsResponse SetPriceListItems(ctx).Body(body).Execute()

Set Price List Items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistSetPriceListItemsRequest() // PricelistSetPriceListItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.SetPriceListItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.SetPriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPriceListItems`: PricelistSetPriceListItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.SetPriceListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistSetPriceListItemsRequest**](PricelistSetPriceListItemsRequest.md) |  | 

### Return type

[**PricelistSetPriceListItemsResponse**](PricelistSetPriceListItemsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePriceList

> map[string]interface{} UpdatePriceList(ctx).Body(body).Execute()

Update Price List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist/v2"
)

func main() {
	body := *openapiclient.NewPricelistUpdatePriceListRequest() // PricelistUpdatePriceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.UpdatePriceList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.UpdatePriceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePriceList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.UpdatePriceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistUpdatePriceListRequest**](PricelistUpdatePriceListRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

