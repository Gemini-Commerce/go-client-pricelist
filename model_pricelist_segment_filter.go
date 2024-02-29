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

// checks if the PricelistSegmentFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistSegmentFilter{}

// PricelistSegmentFilter struct for PricelistSegmentFilter
type PricelistSegmentFilter struct {
	Segments []string `json:"segments,omitempty"`
	Condition *PricelistFilterCondition `json:"condition,omitempty"`
}

// NewPricelistSegmentFilter instantiates a new PricelistSegmentFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistSegmentFilter() *PricelistSegmentFilter {
	this := PricelistSegmentFilter{}
	var condition PricelistFilterCondition = IN
	this.Condition = &condition
	return &this
}

// NewPricelistSegmentFilterWithDefaults instantiates a new PricelistSegmentFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistSegmentFilterWithDefaults() *PricelistSegmentFilter {
	this := PricelistSegmentFilter{}
	var condition PricelistFilterCondition = IN
	this.Condition = &condition
	return &this
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *PricelistSegmentFilter) GetSegments() []string {
	if o == nil || IsNil(o.Segments) {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSegmentFilter) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *PricelistSegmentFilter) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *PricelistSegmentFilter) SetSegments(v []string) {
	o.Segments = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *PricelistSegmentFilter) GetCondition() PricelistFilterCondition {
	if o == nil || IsNil(o.Condition) {
		var ret PricelistFilterCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSegmentFilter) GetConditionOk() (*PricelistFilterCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *PricelistSegmentFilter) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given PricelistFilterCondition and assigns it to the Condition field.
func (o *PricelistSegmentFilter) SetCondition(v PricelistFilterCondition) {
	o.Condition = &v
}

func (o PricelistSegmentFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistSegmentFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullablePricelistSegmentFilter struct {
	value *PricelistSegmentFilter
	isSet bool
}

func (v NullablePricelistSegmentFilter) Get() *PricelistSegmentFilter {
	return v.value
}

func (v *NullablePricelistSegmentFilter) Set(val *PricelistSegmentFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistSegmentFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistSegmentFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistSegmentFilter(val *PricelistSegmentFilter) *NullablePricelistSegmentFilter {
	return &NullablePricelistSegmentFilter{value: val, isSet: true}
}

func (v NullablePricelistSegmentFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistSegmentFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

