# # PricelistGetPriceListResponse


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **string** |   | [optional]
**Grn**| **string** |   | [optional]
**Code**| **string** |   | [optional]
**Name**| **string** |   | [optional]
**Description**| **string** |   | [optional]
**IsActive**| **bool** |   | [optional]
**IsDefault**| **bool** |   | [optional]
**Currency**| [**PricelistCurrency**](PricelistCurrency.md) |  for more information please, see Model/PricelistCurrency.php  | [optional] [default to PRICELISTCURRENCY_XXX]
**VatIncluded**| **bool** |   | [optional]
**DeliveredDutyPaid**| **bool** |   | [optional]
**Segments**| **[]string** |   | [optional]
**Markets**| **[]string** |   | [optional]
**Channels**| **[]string** |   | [optional]
**CreatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**UpdatedAt**| [**time.Time**](time.Time.md) |   | [optional]
**Type**| [**PricelistPriceListType**](PricelistPriceListType.md) |  for more information please, see Model/PricelistPriceListType.php  | [optional] [default to PRICELISTPRICELISTTYPE_UNKNOWN]
**IsSystem**| **bool** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

