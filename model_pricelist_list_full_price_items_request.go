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

// checks if the PricelistListFullPriceItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistListFullPriceItemsRequest{}

// PricelistListFullPriceItemsRequest struct for PricelistListFullPriceItemsRequest
type PricelistListFullPriceItemsRequest struct {
	TenantId             *string `json:"tenantId,omitempty"`
	Id                   *string `json:"id,omitempty"`
	PageSize             *int64  `json:"pageSize,omitempty"`
	PageToken            *string `json:"pageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistListFullPriceItemsRequest PricelistListFullPriceItemsRequest

// NewPricelistListFullPriceItemsRequest instantiates a new PricelistListFullPriceItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistListFullPriceItemsRequest() *PricelistListFullPriceItemsRequest {
	this := PricelistListFullPriceItemsRequest{}
	return &this
}

// NewPricelistListFullPriceItemsRequestWithDefaults instantiates a new PricelistListFullPriceItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistListFullPriceItemsRequestWithDefaults() *PricelistListFullPriceItemsRequest {
	this := PricelistListFullPriceItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PricelistListFullPriceItemsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistListFullPriceItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PricelistListFullPriceItemsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PricelistListFullPriceItemsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PricelistListFullPriceItemsRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistListFullPriceItemsRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PricelistListFullPriceItemsRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PricelistListFullPriceItemsRequest) SetId(v string) {
	o.Id = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PricelistListFullPriceItemsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistListFullPriceItemsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PricelistListFullPriceItemsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *PricelistListFullPriceItemsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *PricelistListFullPriceItemsRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistListFullPriceItemsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *PricelistListFullPriceItemsRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *PricelistListFullPriceItemsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o PricelistListFullPriceItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistListFullPriceItemsRequest) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistListFullPriceItemsRequest) UnmarshalJSON(data []byte) (err error) {
	varPricelistListFullPriceItemsRequest := _PricelistListFullPriceItemsRequest{}

	err = json.Unmarshal(data, &varPricelistListFullPriceItemsRequest)

	if err != nil {
		return err
	}

	*o = PricelistListFullPriceItemsRequest(varPricelistListFullPriceItemsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "pageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistListFullPriceItemsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PricelistListFullPriceItemsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePricelistListFullPriceItemsRequest struct {
	value *PricelistListFullPriceItemsRequest
	isSet bool
}

func (v NullablePricelistListFullPriceItemsRequest) Get() *PricelistListFullPriceItemsRequest {
	return v.value
}

func (v *NullablePricelistListFullPriceItemsRequest) Set(val *PricelistListFullPriceItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistListFullPriceItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistListFullPriceItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistListFullPriceItemsRequest(val *PricelistListFullPriceItemsRequest) *NullablePricelistListFullPriceItemsRequest {
	return &NullablePricelistListFullPriceItemsRequest{value: val, isSet: true}
}

func (v NullablePricelistListFullPriceItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistListFullPriceItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
