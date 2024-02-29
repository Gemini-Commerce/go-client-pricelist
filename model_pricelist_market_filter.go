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

// checks if the PricelistMarketFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistMarketFilter{}

// PricelistMarketFilter struct for PricelistMarketFilter
type PricelistMarketFilter struct {
	Markets []string `json:"markets,omitempty"`
	Condition *PricelistFilterCondition `json:"condition,omitempty"`
}

// NewPricelistMarketFilter instantiates a new PricelistMarketFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistMarketFilter() *PricelistMarketFilter {
	this := PricelistMarketFilter{}
	var condition PricelistFilterCondition = IN
	this.Condition = &condition
	return &this
}

// NewPricelistMarketFilterWithDefaults instantiates a new PricelistMarketFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistMarketFilterWithDefaults() *PricelistMarketFilter {
	this := PricelistMarketFilter{}
	var condition PricelistFilterCondition = IN
	this.Condition = &condition
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *PricelistMarketFilter) GetMarkets() []string {
	if o == nil || IsNil(o.Markets) {
		var ret []string
		return ret
	}
	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistMarketFilter) GetMarketsOk() ([]string, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *PricelistMarketFilter) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given []string and assigns it to the Markets field.
func (o *PricelistMarketFilter) SetMarkets(v []string) {
	o.Markets = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *PricelistMarketFilter) GetCondition() PricelistFilterCondition {
	if o == nil || IsNil(o.Condition) {
		var ret PricelistFilterCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistMarketFilter) GetConditionOk() (*PricelistFilterCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *PricelistMarketFilter) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given PricelistFilterCondition and assigns it to the Condition field.
func (o *PricelistMarketFilter) SetCondition(v PricelistFilterCondition) {
	o.Condition = &v
}

func (o PricelistMarketFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistMarketFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullablePricelistMarketFilter struct {
	value *PricelistMarketFilter
	isSet bool
}

func (v NullablePricelistMarketFilter) Get() *PricelistMarketFilter {
	return v.value
}

func (v *NullablePricelistMarketFilter) Set(val *PricelistMarketFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistMarketFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistMarketFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistMarketFilter(val *PricelistMarketFilter) *NullablePricelistMarketFilter {
	return &NullablePricelistMarketFilter{value: val, isSet: true}
}

func (v NullablePricelistMarketFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistMarketFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

