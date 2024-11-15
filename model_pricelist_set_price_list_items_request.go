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

// checks if the PricelistSetPriceListItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistSetPriceListItemsRequest{}

// PricelistSetPriceListItemsRequest struct for PricelistSetPriceListItemsRequest
type PricelistSetPriceListItemsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
	Items []PricelistSetPriceListItem `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistSetPriceListItemsRequest PricelistSetPriceListItemsRequest

// NewPricelistSetPriceListItemsRequest instantiates a new PricelistSetPriceListItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistSetPriceListItemsRequest() *PricelistSetPriceListItemsRequest {
	this := PricelistSetPriceListItemsRequest{}
	return &this
}

// NewPricelistSetPriceListItemsRequestWithDefaults instantiates a new PricelistSetPriceListItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistSetPriceListItemsRequestWithDefaults() *PricelistSetPriceListItemsRequest {
	this := PricelistSetPriceListItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PricelistSetPriceListItemsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PricelistSetPriceListItemsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PricelistSetPriceListItemsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PricelistSetPriceListItemsRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItemsRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PricelistSetPriceListItemsRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PricelistSetPriceListItemsRequest) SetId(v string) {
	o.Id = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PricelistSetPriceListItemsRequest) GetItems() []PricelistSetPriceListItem {
	if o == nil || IsNil(o.Items) {
		var ret []PricelistSetPriceListItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItemsRequest) GetItemsOk() ([]PricelistSetPriceListItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PricelistSetPriceListItemsRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PricelistSetPriceListItem and assigns it to the Items field.
func (o *PricelistSetPriceListItemsRequest) SetItems(v []PricelistSetPriceListItem) {
	o.Items = v
}

func (o PricelistSetPriceListItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistSetPriceListItemsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistSetPriceListItemsRequest) UnmarshalJSON(data []byte) (err error) {
	varPricelistSetPriceListItemsRequest := _PricelistSetPriceListItemsRequest{}

	err = json.Unmarshal(data, &varPricelistSetPriceListItemsRequest)

	if err != nil {
		return err
	}

	*o = PricelistSetPriceListItemsRequest(varPricelistSetPriceListItemsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistSetPriceListItemsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PricelistSetPriceListItemsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePricelistSetPriceListItemsRequest struct {
	value *PricelistSetPriceListItemsRequest
	isSet bool
}

func (v NullablePricelistSetPriceListItemsRequest) Get() *PricelistSetPriceListItemsRequest {
	return v.value
}

func (v *NullablePricelistSetPriceListItemsRequest) Set(val *PricelistSetPriceListItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistSetPriceListItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistSetPriceListItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistSetPriceListItemsRequest(val *PricelistSetPriceListItemsRequest) *NullablePricelistSetPriceListItemsRequest {
	return &NullablePricelistSetPriceListItemsRequest{value: val, isSet: true}
}

func (v NullablePricelistSetPriceListItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistSetPriceListItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


