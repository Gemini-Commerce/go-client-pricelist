/*
pricelist/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pricelist

import (
	"encoding/json"
)

// checks if the PricelistCurrencyFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistCurrencyFilter{}

// PricelistCurrencyFilter struct for PricelistCurrencyFilter
type PricelistCurrencyFilter struct {
	Currencies []PricelistCurrency `json:"currencies,omitempty"`
	Condition *PricelistFilterCondition `json:"condition,omitempty"`
}

// NewPricelistCurrencyFilter instantiates a new PricelistCurrencyFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistCurrencyFilter() *PricelistCurrencyFilter {
	this := PricelistCurrencyFilter{}
	var condition PricelistFilterCondition = IN
	this.Condition = &condition
	return &this
}

// NewPricelistCurrencyFilterWithDefaults instantiates a new PricelistCurrencyFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistCurrencyFilterWithDefaults() *PricelistCurrencyFilter {
	this := PricelistCurrencyFilter{}
	var condition PricelistFilterCondition = IN
	this.Condition = &condition
	return &this
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *PricelistCurrencyFilter) GetCurrencies() []PricelistCurrency {
	if o == nil || IsNil(o.Currencies) {
		var ret []PricelistCurrency
		return ret
	}
	return o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistCurrencyFilter) GetCurrenciesOk() ([]PricelistCurrency, bool) {
	if o == nil || IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *PricelistCurrencyFilter) HasCurrencies() bool {
	if o != nil && !IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given []PricelistCurrency and assigns it to the Currencies field.
func (o *PricelistCurrencyFilter) SetCurrencies(v []PricelistCurrency) {
	o.Currencies = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *PricelistCurrencyFilter) GetCondition() PricelistFilterCondition {
	if o == nil || IsNil(o.Condition) {
		var ret PricelistFilterCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistCurrencyFilter) GetConditionOk() (*PricelistFilterCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *PricelistCurrencyFilter) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given PricelistFilterCondition and assigns it to the Condition field.
func (o *PricelistCurrencyFilter) SetCondition(v PricelistFilterCondition) {
	o.Condition = &v
}

func (o PricelistCurrencyFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistCurrencyFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currencies) {
		toSerialize["currencies"] = o.Currencies
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullablePricelistCurrencyFilter struct {
	value *PricelistCurrencyFilter
	isSet bool
}

func (v NullablePricelistCurrencyFilter) Get() *PricelistCurrencyFilter {
	return v.value
}

func (v *NullablePricelistCurrencyFilter) Set(val *PricelistCurrencyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistCurrencyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistCurrencyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistCurrencyFilter(val *PricelistCurrencyFilter) *NullablePricelistCurrencyFilter {
	return &NullablePricelistCurrencyFilter{value: val, isSet: true}
}

func (v NullablePricelistCurrencyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistCurrencyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

