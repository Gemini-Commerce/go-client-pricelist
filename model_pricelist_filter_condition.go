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
	"fmt"
)

// PricelistFilterCondition the model 'PricelistFilterCondition'
type PricelistFilterCondition string

// List of pricelistFilterCondition
const (
	PRICELISTFILTERCONDITION_IN PricelistFilterCondition = "IN"
	PRICELISTFILTERCONDITION_NOT_IN PricelistFilterCondition = "NOT_IN"
)

// All allowed values of PricelistFilterCondition enum
var AllowedPricelistFilterConditionEnumValues = []PricelistFilterCondition{
	"IN",
	"NOT_IN",
}

func (v *PricelistFilterCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PricelistFilterCondition(value)
	for _, existing := range AllowedPricelistFilterConditionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PricelistFilterCondition", value)
}

// NewPricelistFilterConditionFromValue returns a pointer to a valid PricelistFilterCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPricelistFilterConditionFromValue(v string) (*PricelistFilterCondition, error) {
	ev := PricelistFilterCondition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PricelistFilterCondition: valid values are %v", v, AllowedPricelistFilterConditionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PricelistFilterCondition) IsValid() bool {
	for _, existing := range AllowedPricelistFilterConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to pricelistFilterCondition value
func (v PricelistFilterCondition) Ptr() *PricelistFilterCondition {
	return &v
}

type NullablePricelistFilterCondition struct {
	value *PricelistFilterCondition
	isSet bool
}

func (v NullablePricelistFilterCondition) Get() *PricelistFilterCondition {
	return v.value
}

func (v *NullablePricelistFilterCondition) Set(val *PricelistFilterCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistFilterCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistFilterCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistFilterCondition(val *PricelistFilterCondition) *NullablePricelistFilterCondition {
	return &NullablePricelistFilterCondition{value: val, isSet: true}
}

func (v NullablePricelistFilterCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistFilterCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

