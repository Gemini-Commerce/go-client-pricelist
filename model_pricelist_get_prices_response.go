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

// checks if the PricelistGetPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetPricesResponse{}

// PricelistGetPricesResponse struct for PricelistGetPricesResponse
type PricelistGetPricesResponse struct {
	Prices []PricelistGetPriceItem `json:"prices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistGetPricesResponse PricelistGetPricesResponse

// NewPricelistGetPricesResponse instantiates a new PricelistGetPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetPricesResponse() *PricelistGetPricesResponse {
	this := PricelistGetPricesResponse{}
	return &this
}

// NewPricelistGetPricesResponseWithDefaults instantiates a new PricelistGetPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetPricesResponseWithDefaults() *PricelistGetPricesResponse {
	this := PricelistGetPricesResponse{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *PricelistGetPricesResponse) GetPrices() []PricelistGetPriceItem {
	if o == nil || IsNil(o.Prices) {
		var ret []PricelistGetPriceItem
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPricesResponse) GetPricesOk() ([]PricelistGetPriceItem, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *PricelistGetPricesResponse) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []PricelistGetPriceItem and assigns it to the Prices field.
func (o *PricelistGetPricesResponse) SetPrices(v []PricelistGetPriceItem) {
	o.Prices = v
}

func (o PricelistGetPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetPricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistGetPricesResponse) UnmarshalJSON(data []byte) (err error) {
	varPricelistGetPricesResponse := _PricelistGetPricesResponse{}

	err = json.Unmarshal(data, &varPricelistGetPricesResponse)

	if err != nil {
		return err
	}

	*o = PricelistGetPricesResponse(varPricelistGetPricesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "prices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistGetPricesResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PricelistGetPricesResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePricelistGetPricesResponse struct {
	value *PricelistGetPricesResponse
	isSet bool
}

func (v NullablePricelistGetPricesResponse) Get() *PricelistGetPricesResponse {
	return v.value
}

func (v *NullablePricelistGetPricesResponse) Set(val *PricelistGetPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetPricesResponse(val *PricelistGetPricesResponse) *NullablePricelistGetPricesResponse {
	return &NullablePricelistGetPricesResponse{value: val, isSet: true}
}

func (v NullablePricelistGetPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


