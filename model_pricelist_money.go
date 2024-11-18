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

// checks if the PricelistMoney type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistMoney{}

// PricelistMoney Represents an amount of money.
type PricelistMoney struct {
	// The whole units of the amount. For example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.
	Units *string `json:"units,omitempty"`
	// Number of micro (10^-6) units of the amount. The value must be between -999,999 and +999,999 inclusive. If `units` is positive, `micros` must be positive or zero. If `units` is zero, `micros` can be positive, zero, or negative. If `units` is negative, `micros` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `micros`=-750,000.
	Micros               *int32 `json:"micros,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistMoney PricelistMoney

// NewPricelistMoney instantiates a new PricelistMoney object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistMoney() *PricelistMoney {
	this := PricelistMoney{}
	return &this
}

// NewPricelistMoneyWithDefaults instantiates a new PricelistMoney object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistMoneyWithDefaults() *PricelistMoney {
	this := PricelistMoney{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *PricelistMoney) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistMoney) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *PricelistMoney) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *PricelistMoney) SetUnits(v string) {
	o.Units = &v
}

// GetMicros returns the Micros field value if set, zero value otherwise.
func (o *PricelistMoney) GetMicros() int32 {
	if o == nil || IsNil(o.Micros) {
		var ret int32
		return ret
	}
	return *o.Micros
}

// GetMicrosOk returns a tuple with the Micros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistMoney) GetMicrosOk() (*int32, bool) {
	if o == nil || IsNil(o.Micros) {
		return nil, false
	}
	return o.Micros, true
}

// HasMicros returns a boolean if a field has been set.
func (o *PricelistMoney) HasMicros() bool {
	if o != nil && !IsNil(o.Micros) {
		return true
	}

	return false
}

// SetMicros gets a reference to the given int32 and assigns it to the Micros field.
func (o *PricelistMoney) SetMicros(v int32) {
	o.Micros = &v
}

func (o PricelistMoney) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistMoney) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.Micros) {
		toSerialize["micros"] = o.Micros
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistMoney) UnmarshalJSON(data []byte) (err error) {
	varPricelistMoney := _PricelistMoney{}

	err = json.Unmarshal(data, &varPricelistMoney)

	if err != nil {
		return err
	}

	*o = PricelistMoney(varPricelistMoney)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "units")
		delete(additionalProperties, "micros")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistMoney) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PricelistMoney) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePricelistMoney struct {
	value *PricelistMoney
	isSet bool
}

func (v NullablePricelistMoney) Get() *PricelistMoney {
	return v.value
}

func (v *NullablePricelistMoney) Set(val *PricelistMoney) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistMoney) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistMoney(val *PricelistMoney) *NullablePricelistMoney {
	return &NullablePricelistMoney{value: val, isSet: true}
}

func (v NullablePricelistMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
