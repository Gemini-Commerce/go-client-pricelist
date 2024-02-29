# GeminiCommerce\Pricelist\PriceListAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PriceListCreatePriceList**](PriceListAPI.md#PriceListCreatePriceList) | **Post** /pricelist.PriceList/CreatePriceList | 
[**PriceListDeletePriceListItems**](PriceListAPI.md#PriceListDeletePriceListItems) | **Post** /pricelist.PriceList/DeletePriceListItems | 
[**PriceListGetFullPriceItemsByPricelistId**](PriceListAPI.md#PriceListGetFullPriceItemsByPricelistId) | **Post** /pricelist.PriceList/GetFullPriceItemsByPricelistId | 
[**PriceListGetPriceItemsByPriceListItemIds**](PriceListAPI.md#PriceListGetPriceItemsByPriceListItemIds) | **Post** /pricelist.PriceList/GetPriceItemsByPriceListItemIds | 
[**PriceListGetPriceList**](PriceListAPI.md#PriceListGetPriceList) | **Post** /pricelist.PriceList/GetPriceList | 
[**PriceListGetPriceListByCode**](PriceListAPI.md#PriceListGetPriceListByCode) | **Post** /pricelist.PriceList/GetPriceListByCode | 
[**PriceListGetPriceListItems**](PriceListAPI.md#PriceListGetPriceListItems) | **Post** /pricelist.PriceList/GetPriceListItems | 
[**PriceListGetPricesItems**](PriceListAPI.md#PriceListGetPricesItems) | **Post** /pricelist.PriceList/GetPricesItems | 
[**PriceListListFullPriceItemsByPricelistId**](PriceListAPI.md#PriceListListFullPriceItemsByPricelistId) | **Post** /pricelist.PriceList/ListFullPriceItemsByPricelistId | 
[**PriceListListPriceLists**](PriceListAPI.md#PriceListListPriceLists) | **Post** /pricelist.PriceList/ListPriceLists | 
[**PriceListSetPriceListItems**](PriceListAPI.md#PriceListSetPriceListItems) | **Post** /pricelist.PriceList/SetPriceListItems | 
[**PriceListUpdatePriceList**](PriceListAPI.md#PriceListUpdatePriceList) | **Post** /pricelist.PriceList/UpdatePriceList | 



## PriceListCreatePriceList

> PricelistCreatePriceListResponse PriceListCreatePriceList(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistCreatePriceListRequest() // PricelistCreatePriceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListCreatePriceList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListCreatePriceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListCreatePriceList`: PricelistCreatePriceListResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListCreatePriceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListCreatePriceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistCreatePriceListRequest**](PricelistCreatePriceListRequest.md) |  | 

### Return type

[**PricelistCreatePriceListResponse**](PricelistCreatePriceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListDeletePriceListItems

> map[string]interface{} PriceListDeletePriceListItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistDeletePriceListItemsRequest() // PricelistDeletePriceListItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListDeletePriceListItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListDeletePriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListDeletePriceListItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListDeletePriceListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListDeletePriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistDeletePriceListItemsRequest**](PricelistDeletePriceListItemsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListGetFullPriceItemsByPricelistId

> PricelistGetFullPriceItemsResponse PriceListGetFullPriceItemsByPricelistId(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistGetFullPriceItemsRequest() // PricelistGetFullPriceItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListGetFullPriceItemsByPricelistId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListGetFullPriceItemsByPricelistId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListGetFullPriceItemsByPricelistId`: PricelistGetFullPriceItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListGetFullPriceItemsByPricelistId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListGetFullPriceItemsByPricelistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetFullPriceItemsRequest**](PricelistGetFullPriceItemsRequest.md) |  | 

### Return type

[**PricelistGetFullPriceItemsResponse**](PricelistGetFullPriceItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListGetPriceItemsByPriceListItemIds

> PricelistGetPriceItemsByPriceListItemIdsResponse PriceListGetPriceItemsByPriceListItemIds(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceItemsByPriceListItemIdsRequest() // PricelistGetPriceItemsByPriceListItemIdsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListGetPriceItemsByPriceListItemIds(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListGetPriceItemsByPriceListItemIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListGetPriceItemsByPriceListItemIds`: PricelistGetPriceItemsByPriceListItemIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListGetPriceItemsByPriceListItemIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListGetPriceItemsByPriceListItemIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceItemsByPriceListItemIdsRequest**](PricelistGetPriceItemsByPriceListItemIdsRequest.md) |  | 

### Return type

[**PricelistGetPriceItemsByPriceListItemIdsResponse**](PricelistGetPriceItemsByPriceListItemIdsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListGetPriceList

> PricelistGetPriceListResponse PriceListGetPriceList(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceListRequest() // PricelistGetPriceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListGetPriceList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListGetPriceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListGetPriceList`: PricelistGetPriceListResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListGetPriceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListGetPriceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceListRequest**](PricelistGetPriceListRequest.md) |  | 

### Return type

[**PricelistGetPriceListResponse**](PricelistGetPriceListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListGetPriceListByCode

> PricelistGetPriceListByCodeResponse PriceListGetPriceListByCode(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceListByCodeRequest() // PricelistGetPriceListByCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListGetPriceListByCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListGetPriceListByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListGetPriceListByCode`: PricelistGetPriceListByCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListGetPriceListByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListGetPriceListByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceListByCodeRequest**](PricelistGetPriceListByCodeRequest.md) |  | 

### Return type

[**PricelistGetPriceListByCodeResponse**](PricelistGetPriceListByCodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListGetPriceListItems

> PricelistGetPriceListItemsResponse PriceListGetPriceListItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistGetPriceListItemsRequest() // PricelistGetPriceListItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListGetPriceListItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListGetPriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListGetPriceListItems`: PricelistGetPriceListItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListGetPriceListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListGetPriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPriceListItemsRequest**](PricelistGetPriceListItemsRequest.md) |  | 

### Return type

[**PricelistGetPriceListItemsResponse**](PricelistGetPriceListItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListGetPricesItems

> PricelistGetPricesResponse PriceListGetPricesItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistGetPricesRequest() // PricelistGetPricesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListGetPricesItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListGetPricesItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListGetPricesItems`: PricelistGetPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListGetPricesItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListGetPricesItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistGetPricesRequest**](PricelistGetPricesRequest.md) |  | 

### Return type

[**PricelistGetPricesResponse**](PricelistGetPricesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListListFullPriceItemsByPricelistId

> PricelistListFullPriceItemsResponse PriceListListFullPriceItemsByPricelistId(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistListFullPriceItemsRequest() // PricelistListFullPriceItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListListFullPriceItemsByPricelistId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListListFullPriceItemsByPricelistId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListListFullPriceItemsByPricelistId`: PricelistListFullPriceItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListListFullPriceItemsByPricelistId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListListFullPriceItemsByPricelistIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistListFullPriceItemsRequest**](PricelistListFullPriceItemsRequest.md) |  | 

### Return type

[**PricelistListFullPriceItemsResponse**](PricelistListFullPriceItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListListPriceLists

> PricelistListPriceListsResponse PriceListListPriceLists(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistListPriceListsRequest() // PricelistListPriceListsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListListPriceLists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListListPriceLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListListPriceLists`: PricelistListPriceListsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListListPriceLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListListPriceListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistListPriceListsRequest**](PricelistListPriceListsRequest.md) |  | 

### Return type

[**PricelistListPriceListsResponse**](PricelistListPriceListsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListSetPriceListItems

> PricelistSetPriceListItemsResponse PriceListSetPriceListItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistSetPriceListItemsRequest() // PricelistSetPriceListItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListSetPriceListItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListSetPriceListItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListSetPriceListItems`: PricelistSetPriceListItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListSetPriceListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListSetPriceListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistSetPriceListItemsRequest**](PricelistSetPriceListItemsRequest.md) |  | 

### Return type

[**PricelistSetPriceListItemsResponse**](PricelistSetPriceListItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PriceListUpdatePriceList

> map[string]interface{} PriceListUpdatePriceList(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-pricelist"
)

func main() {
	body := *openapiclient.NewPricelistUpdatePriceListRequest() // PricelistUpdatePriceListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListAPI.PriceListUpdatePriceList(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListAPI.PriceListUpdatePriceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PriceListUpdatePriceList`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PriceListAPI.PriceListUpdatePriceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPriceListUpdatePriceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PricelistUpdatePriceListRequest**](PricelistUpdatePriceListRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

