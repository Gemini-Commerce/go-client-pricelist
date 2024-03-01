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

// checks if the PricelistPriceListTypeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistPriceListTypeFilter{}

// PricelistPriceListTypeFilter struct for PricelistPriceListTypeFilter
type PricelistPriceListTypeFilter struct {
	PricelistTypes []PricelistPriceListType `json:"pricelistTypes,omitempty"`
	Condition *PricelistFilterCondition `json:"condition,omitempty"`
}

// NewPricelistPriceListTypeFilter instantiates a new PricelistPriceListTypeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistPriceListTypeFilter() *PricelistPriceListTypeFilter {
	this := PricelistPriceListTypeFilter{}
	var condition PricelistFilterCondition = PRICELISTFILTERCONDITION_IN
	this.Condition = &condition
	return &this
}

// NewPricelistPriceListTypeFilterWithDefaults instantiates a new PricelistPriceListTypeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistPriceListTypeFilterWithDefaults() *PricelistPriceListTypeFilter {
	this := PricelistPriceListTypeFilter{}
	var condition PricelistFilterCondition = PRICELISTFILTERCONDITION_IN
	this.Condition = &condition
	return &this
}

// GetPricelistTypes returns the PricelistTypes field value if set, zero value otherwise.
func (o *PricelistPriceListTypeFilter) GetPricelistTypes() []PricelistPriceListType {
	if o == nil || IsNil(o.PricelistTypes) {
		var ret []PricelistPriceListType
		return ret
	}
	return o.PricelistTypes
}

// GetPricelistTypesOk returns a tuple with the PricelistTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistPriceListTypeFilter) GetPricelistTypesOk() ([]PricelistPriceListType, bool) {
	if o == nil || IsNil(o.PricelistTypes) {
		return nil, false
	}
	return o.PricelistTypes, true
}

// HasPricelistTypes returns a boolean if a field has been set.
func (o *PricelistPriceListTypeFilter) HasPricelistTypes() bool {
	if o != nil && !IsNil(o.PricelistTypes) {
		return true
	}

	return false
}

// SetPricelistTypes gets a reference to the given []PricelistPriceListType and assigns it to the PricelistTypes field.
func (o *PricelistPriceListTypeFilter) SetPricelistTypes(v []PricelistPriceListType) {
	o.PricelistTypes = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *PricelistPriceListTypeFilter) GetCondition() PricelistFilterCondition {
	if o == nil || IsNil(o.Condition) {
		var ret PricelistFilterCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistPriceListTypeFilter) GetConditionOk() (*PricelistFilterCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *PricelistPriceListTypeFilter) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given PricelistFilterCondition and assigns it to the Condition field.
func (o *PricelistPriceListTypeFilter) SetCondition(v PricelistFilterCondition) {
	o.Condition = &v
}

func (o PricelistPriceListTypeFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistPriceListTypeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PricelistTypes) {
		toSerialize["pricelistTypes"] = o.PricelistTypes
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullablePricelistPriceListTypeFilter struct {
	value *PricelistPriceListTypeFilter
	isSet bool
}

func (v NullablePricelistPriceListTypeFilter) Get() *PricelistPriceListTypeFilter {
	return v.value
}

func (v *NullablePricelistPriceListTypeFilter) Set(val *PricelistPriceListTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistPriceListTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistPriceListTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistPriceListTypeFilter(val *PricelistPriceListTypeFilter) *NullablePricelistPriceListTypeFilter {
	return &NullablePricelistPriceListTypeFilter{value: val, isSet: true}
}

func (v NullablePricelistPriceListTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistPriceListTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


