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

// checks if the PricelistGetPriceItemsByPriceListItemIdsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetPriceItemsByPriceListItemIdsRequest{}

// PricelistGetPriceItemsByPriceListItemIdsRequest struct for PricelistGetPriceItemsByPriceListItemIdsRequest
type PricelistGetPriceItemsByPriceListItemIdsRequest struct {
	TenantId             *string  `json:"tenantId,omitempty"`
	PriceListItemId      []string `json:"priceListItemId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistGetPriceItemsByPriceListItemIdsRequest PricelistGetPriceItemsByPriceListItemIdsRequest

// NewPricelistGetPriceItemsByPriceListItemIdsRequest instantiates a new PricelistGetPriceItemsByPriceListItemIdsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetPriceItemsByPriceListItemIdsRequest() *PricelistGetPriceItemsByPriceListItemIdsRequest {
	this := PricelistGetPriceItemsByPriceListItemIdsRequest{}
	return &this
}

// NewPricelistGetPriceItemsByPriceListItemIdsRequestWithDefaults instantiates a new PricelistGetPriceItemsByPriceListItemIdsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetPriceItemsByPriceListItemIdsRequestWithDefaults() *PricelistGetPriceItemsByPriceListItemIdsRequest {
	this := PricelistGetPriceItemsByPriceListItemIdsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPriceListItemId returns the PriceListItemId field value if set, zero value otherwise.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) GetPriceListItemId() []string {
	if o == nil || IsNil(o.PriceListItemId) {
		var ret []string
		return ret
	}
	return o.PriceListItemId
}

// GetPriceListItemIdOk returns a tuple with the PriceListItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) GetPriceListItemIdOk() ([]string, bool) {
	if o == nil || IsNil(o.PriceListItemId) {
		return nil, false
	}
	return o.PriceListItemId, true
}

// HasPriceListItemId returns a boolean if a field has been set.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) HasPriceListItemId() bool {
	if o != nil && !IsNil(o.PriceListItemId) {
		return true
	}

	return false
}

// SetPriceListItemId gets a reference to the given []string and assigns it to the PriceListItemId field.
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) SetPriceListItemId(v []string) {
	o.PriceListItemId = v
}

func (o PricelistGetPriceItemsByPriceListItemIdsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetPriceItemsByPriceListItemIdsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.PriceListItemId) {
		toSerialize["priceListItemId"] = o.PriceListItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) UnmarshalJSON(data []byte) (err error) {
	varPricelistGetPriceItemsByPriceListItemIdsRequest := _PricelistGetPriceItemsByPriceListItemIdsRequest{}

	err = json.Unmarshal(data, &varPricelistGetPriceItemsByPriceListItemIdsRequest)

	if err != nil {
		return err
	}

	*o = PricelistGetPriceItemsByPriceListItemIdsRequest(varPricelistGetPriceItemsByPriceListItemIdsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "priceListItemId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PricelistGetPriceItemsByPriceListItemIdsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePricelistGetPriceItemsByPriceListItemIdsRequest struct {
	value *PricelistGetPriceItemsByPriceListItemIdsRequest
	isSet bool
}

func (v NullablePricelistGetPriceItemsByPriceListItemIdsRequest) Get() *PricelistGetPriceItemsByPriceListItemIdsRequest {
	return v.value
}

func (v *NullablePricelistGetPriceItemsByPriceListItemIdsRequest) Set(val *PricelistGetPriceItemsByPriceListItemIdsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetPriceItemsByPriceListItemIdsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetPriceItemsByPriceListItemIdsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetPriceItemsByPriceListItemIdsRequest(val *PricelistGetPriceItemsByPriceListItemIdsRequest) *NullablePricelistGetPriceItemsByPriceListItemIdsRequest {
	return &NullablePricelistGetPriceItemsByPriceListItemIdsRequest{value: val, isSet: true}
}

func (v NullablePricelistGetPriceItemsByPriceListItemIdsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetPriceItemsByPriceListItemIdsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
