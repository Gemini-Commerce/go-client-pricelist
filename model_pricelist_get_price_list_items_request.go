/*
PriceList Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pricelist

import (
	"encoding/json"
)

// checks if the PricelistGetPriceListItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetPriceListItemsRequest{}

// PricelistGetPriceListItemsRequest struct for PricelistGetPriceListItemsRequest
type PricelistGetPriceListItemsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewPricelistGetPriceListItemsRequest instantiates a new PricelistGetPriceListItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetPriceListItemsRequest() *PricelistGetPriceListItemsRequest {
	this := PricelistGetPriceListItemsRequest{}
	return &this
}

// NewPricelistGetPriceListItemsRequestWithDefaults instantiates a new PricelistGetPriceListItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetPriceListItemsRequestWithDefaults() *PricelistGetPriceListItemsRequest {
	this := PricelistGetPriceListItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PricelistGetPriceListItemsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PricelistGetPriceListItemsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PricelistGetPriceListItemsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PricelistGetPriceListItemsRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItemsRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PricelistGetPriceListItemsRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PricelistGetPriceListItemsRequest) SetId(v string) {
	o.Id = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PricelistGetPriceListItemsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItemsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PricelistGetPriceListItemsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *PricelistGetPriceListItemsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *PricelistGetPriceListItemsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItemsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *PricelistGetPriceListItemsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *PricelistGetPriceListItemsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o PricelistGetPriceListItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetPriceListItemsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return toSerialize, nil
}

type NullablePricelistGetPriceListItemsRequest struct {
	value *PricelistGetPriceListItemsRequest
	isSet bool
}

func (v NullablePricelistGetPriceListItemsRequest) Get() *PricelistGetPriceListItemsRequest {
	return v.value
}

func (v *NullablePricelistGetPriceListItemsRequest) Set(val *PricelistGetPriceListItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetPriceListItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetPriceListItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetPriceListItemsRequest(val *PricelistGetPriceListItemsRequest) *NullablePricelistGetPriceListItemsRequest {
	return &NullablePricelistGetPriceListItemsRequest{value: val, isSet: true}
}

func (v NullablePricelistGetPriceListItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetPriceListItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


