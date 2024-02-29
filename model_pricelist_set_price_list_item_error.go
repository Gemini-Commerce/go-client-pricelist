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

// checks if the PricelistSetPriceListItemError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistSetPriceListItemError{}

// PricelistSetPriceListItemError struct for PricelistSetPriceListItemError
type PricelistSetPriceListItemError struct {
	ItemGrn *string `json:"itemGrn,omitempty"`
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewPricelistSetPriceListItemError instantiates a new PricelistSetPriceListItemError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistSetPriceListItemError() *PricelistSetPriceListItemError {
	this := PricelistSetPriceListItemError{}
	return &this
}

// NewPricelistSetPriceListItemErrorWithDefaults instantiates a new PricelistSetPriceListItemError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistSetPriceListItemErrorWithDefaults() *PricelistSetPriceListItemError {
	this := PricelistSetPriceListItemError{}
	return &this
}

// GetItemGrn returns the ItemGrn field value if set, zero value otherwise.
func (o *PricelistSetPriceListItemError) GetItemGrn() string {
	if o == nil || IsNil(o.ItemGrn) {
		var ret string
		return ret
	}
	return *o.ItemGrn
}

// GetItemGrnOk returns a tuple with the ItemGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItemError) GetItemGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGrn) {
		return nil, false
	}
	return o.ItemGrn, true
}

// HasItemGrn returns a boolean if a field has been set.
func (o *PricelistSetPriceListItemError) HasItemGrn() bool {
	if o != nil && !IsNil(o.ItemGrn) {
		return true
	}

	return false
}

// SetItemGrn gets a reference to the given string and assigns it to the ItemGrn field.
func (o *PricelistSetPriceListItemError) SetItemGrn(v string) {
	o.ItemGrn = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PricelistSetPriceListItemError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItemError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PricelistSetPriceListItemError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PricelistSetPriceListItemError) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PricelistSetPriceListItemError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItemError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PricelistSetPriceListItemError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PricelistSetPriceListItemError) SetDescription(v string) {
	o.Description = &v
}

func (o PricelistSetPriceListItemError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistSetPriceListItemError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemGrn) {
		toSerialize["itemGrn"] = o.ItemGrn
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullablePricelistSetPriceListItemError struct {
	value *PricelistSetPriceListItemError
	isSet bool
}

func (v NullablePricelistSetPriceListItemError) Get() *PricelistSetPriceListItemError {
	return v.value
}

func (v *NullablePricelistSetPriceListItemError) Set(val *PricelistSetPriceListItemError) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistSetPriceListItemError) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistSetPriceListItemError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistSetPriceListItemError(val *PricelistSetPriceListItemError) *NullablePricelistSetPriceListItemError {
	return &NullablePricelistSetPriceListItemError{value: val, isSet: true}
}

func (v NullablePricelistSetPriceListItemError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistSetPriceListItemError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

