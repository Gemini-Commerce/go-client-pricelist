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

// checks if the PricelistGetFullPriceItemsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetFullPriceItemsRequest{}

// PricelistGetFullPriceItemsRequest struct for PricelistGetFullPriceItemsRequest
type PricelistGetFullPriceItemsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	PricelistId *string `json:"pricelistId,omitempty"`
	ItemsGrn []string `json:"itemsGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistGetFullPriceItemsRequest PricelistGetFullPriceItemsRequest

// NewPricelistGetFullPriceItemsRequest instantiates a new PricelistGetFullPriceItemsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetFullPriceItemsRequest() *PricelistGetFullPriceItemsRequest {
	this := PricelistGetFullPriceItemsRequest{}
	return &this
}

// NewPricelistGetFullPriceItemsRequestWithDefaults instantiates a new PricelistGetFullPriceItemsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetFullPriceItemsRequestWithDefaults() *PricelistGetFullPriceItemsRequest {
	this := PricelistGetFullPriceItemsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PricelistGetFullPriceItemsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetFullPriceItemsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// &#39;Has&#39;TenantId returns a boolean if a field has been set.
func (o *PricelistGetFullPriceItemsRequest) &#39;Has&#39;TenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PricelistGetFullPriceItemsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetPricelistId returns the PricelistId field value if set, zero value otherwise.
func (o *PricelistGetFullPriceItemsRequest) GetPricelistId() string {
	if o == nil || IsNil(o.PricelistId) {
		var ret string
		return ret
	}
	return *o.PricelistId
}

// GetPricelistIdOk returns a tuple with the PricelistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetFullPriceItemsRequest) GetPricelistIdOk() (*string, bool) {
	if o == nil || IsNil(o.PricelistId) {
		return nil, false
	}
	return o.PricelistId, true
}

// &#39;Has&#39;PricelistId returns a boolean if a field has been set.
func (o *PricelistGetFullPriceItemsRequest) &#39;Has&#39;PricelistId() bool {
	if o != nil && !IsNil(o.PricelistId) {
		return true
	}

	return false
}

// SetPricelistId gets a reference to the given string and assigns it to the PricelistId field.
func (o *PricelistGetFullPriceItemsRequest) SetPricelistId(v string) {
	o.PricelistId = &v
}

// GetItemsGrn returns the ItemsGrn field value if set, zero value otherwise.
func (o *PricelistGetFullPriceItemsRequest) GetItemsGrn() []string {
	if o == nil || IsNil(o.ItemsGrn) {
		var ret []string
		return ret
	}
	return o.ItemsGrn
}

// GetItemsGrnOk returns a tuple with the ItemsGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetFullPriceItemsRequest) GetItemsGrnOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsGrn) {
		return nil, false
	}
	return o.ItemsGrn, true
}

// &#39;Has&#39;ItemsGrn returns a boolean if a field has been set.
func (o *PricelistGetFullPriceItemsRequest) &#39;Has&#39;ItemsGrn() bool {
	if o != nil && !IsNil(o.ItemsGrn) {
		return true
	}

	return false
}

// SetItemsGrn gets a reference to the given []string and assigns it to the ItemsGrn field.
func (o *PricelistGetFullPriceItemsRequest) SetItemsGrn(v []string) {
	o.ItemsGrn = v
}

func (o PricelistGetFullPriceItemsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetFullPriceItemsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.PricelistId) {
		toSerialize["pricelistId"] = o.PricelistId
	}
	if !IsNil(o.ItemsGrn) {
		toSerialize["itemsGrn"] = o.ItemsGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistGetFullPriceItemsRequest) UnmarshalJSON(data []byte) (err error) {
	varPricelistGetFullPriceItemsRequest := _PricelistGetFullPriceItemsRequest{}

	err = json.Unmarshal(data, &varPricelistGetFullPriceItemsRequest)

	if err != nil {
		return err
	}

	*o = PricelistGetFullPriceItemsRequest(varPricelistGetFullPriceItemsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "pricelistId")
		delete(additionalProperties, "itemsGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistGetFullPriceItemsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PricelistGetFullPriceItemsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePricelistGetFullPriceItemsRequest struct {
	value *PricelistGetFullPriceItemsRequest
	isSet bool
}

func (v NullablePricelistGetFullPriceItemsRequest) Get() *PricelistGetFullPriceItemsRequest {
	return v.value
}

func (v *NullablePricelistGetFullPriceItemsRequest) Set(val *PricelistGetFullPriceItemsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetFullPriceItemsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetFullPriceItemsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetFullPriceItemsRequest(val *PricelistGetFullPriceItemsRequest) *NullablePricelistGetFullPriceItemsRequest {
	return &NullablePricelistGetFullPriceItemsRequest{value: val, isSet: true}
}

func (v NullablePricelistGetFullPriceItemsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetFullPriceItemsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


