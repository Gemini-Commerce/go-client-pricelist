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

// checks if the PricelistSetPriceListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistSetPriceListItem{}

// PricelistSetPriceListItem struct for PricelistSetPriceListItem
type PricelistSetPriceListItem struct {
	ItemGrn              *string                          `json:"itemGrn,omitempty"`
	BasePrice            *PricelistMoney                  `json:"basePrice,omitempty"`
	PriceItems           []PricelistSetPriceListItemPrice `json:"priceItems,omitempty"`
	HasTierPrices        *bool                            `json:"hasTierPrices,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistSetPriceListItem PricelistSetPriceListItem

// NewPricelistSetPriceListItem instantiates a new PricelistSetPriceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistSetPriceListItem() *PricelistSetPriceListItem {
	this := PricelistSetPriceListItem{}
	return &this
}

// NewPricelistSetPriceListItemWithDefaults instantiates a new PricelistSetPriceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistSetPriceListItemWithDefaults() *PricelistSetPriceListItem {
	this := PricelistSetPriceListItem{}
	return &this
}

// GetItemGrn returns the ItemGrn field value if set, zero value otherwise.
func (o *PricelistSetPriceListItem) GetItemGrn() string {
	if o == nil || IsNil(o.ItemGrn) {
		var ret string
		return ret
	}
	return *o.ItemGrn
}

// GetItemGrnOk returns a tuple with the ItemGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItem) GetItemGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGrn) {
		return nil, false
	}
	return o.ItemGrn, true
}

// HasItemGrn returns a boolean if a field has been set.
func (o *PricelistSetPriceListItem) HasItemGrn() bool {
	if o != nil && !IsNil(o.ItemGrn) {
		return true
	}

	return false
}

// SetItemGrn gets a reference to the given string and assigns it to the ItemGrn field.
func (o *PricelistSetPriceListItem) SetItemGrn(v string) {
	o.ItemGrn = &v
}

// GetBasePrice returns the BasePrice field value if set, zero value otherwise.
func (o *PricelistSetPriceListItem) GetBasePrice() PricelistMoney {
	if o == nil || IsNil(o.BasePrice) {
		var ret PricelistMoney
		return ret
	}
	return *o.BasePrice
}

// GetBasePriceOk returns a tuple with the BasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItem) GetBasePriceOk() (*PricelistMoney, bool) {
	if o == nil || IsNil(o.BasePrice) {
		return nil, false
	}
	return o.BasePrice, true
}

// HasBasePrice returns a boolean if a field has been set.
func (o *PricelistSetPriceListItem) HasBasePrice() bool {
	if o != nil && !IsNil(o.BasePrice) {
		return true
	}

	return false
}

// SetBasePrice gets a reference to the given PricelistMoney and assigns it to the BasePrice field.
func (o *PricelistSetPriceListItem) SetBasePrice(v PricelistMoney) {
	o.BasePrice = &v
}

// GetPriceItems returns the PriceItems field value if set, zero value otherwise.
func (o *PricelistSetPriceListItem) GetPriceItems() []PricelistSetPriceListItemPrice {
	if o == nil || IsNil(o.PriceItems) {
		var ret []PricelistSetPriceListItemPrice
		return ret
	}
	return o.PriceItems
}

// GetPriceItemsOk returns a tuple with the PriceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItem) GetPriceItemsOk() ([]PricelistSetPriceListItemPrice, bool) {
	if o == nil || IsNil(o.PriceItems) {
		return nil, false
	}
	return o.PriceItems, true
}

// HasPriceItems returns a boolean if a field has been set.
func (o *PricelistSetPriceListItem) HasPriceItems() bool {
	if o != nil && !IsNil(o.PriceItems) {
		return true
	}

	return false
}

// SetPriceItems gets a reference to the given []PricelistSetPriceListItemPrice and assigns it to the PriceItems field.
func (o *PricelistSetPriceListItem) SetPriceItems(v []PricelistSetPriceListItemPrice) {
	o.PriceItems = v
}

// GetHasTierPrices returns the HasTierPrices field value if set, zero value otherwise.
func (o *PricelistSetPriceListItem) GetHasTierPrices() bool {
	if o == nil || IsNil(o.HasTierPrices) {
		var ret bool
		return ret
	}
	return *o.HasTierPrices
}

// GetHasTierPricesOk returns a tuple with the HasTierPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistSetPriceListItem) GetHasTierPricesOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTierPrices) {
		return nil, false
	}
	return o.HasTierPrices, true
}

// HasHasTierPrices returns a boolean if a field has been set.
func (o *PricelistSetPriceListItem) HasHasTierPrices() bool {
	if o != nil && !IsNil(o.HasTierPrices) {
		return true
	}

	return false
}

// SetHasTierPrices gets a reference to the given bool and assigns it to the HasTierPrices field.
func (o *PricelistSetPriceListItem) SetHasTierPrices(v bool) {
	o.HasTierPrices = &v
}

func (o PricelistSetPriceListItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistSetPriceListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemGrn) {
		toSerialize["itemGrn"] = o.ItemGrn
	}
	if !IsNil(o.BasePrice) {
		toSerialize["basePrice"] = o.BasePrice
	}
	if !IsNil(o.PriceItems) {
		toSerialize["priceItems"] = o.PriceItems
	}
	if !IsNil(o.HasTierPrices) {
		toSerialize["hasTierPrices"] = o.HasTierPrices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistSetPriceListItem) UnmarshalJSON(data []byte) (err error) {
	varPricelistSetPriceListItem := _PricelistSetPriceListItem{}

	err = json.Unmarshal(data, &varPricelistSetPriceListItem)

	if err != nil {
		return err
	}

	*o = PricelistSetPriceListItem(varPricelistSetPriceListItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "itemGrn")
		delete(additionalProperties, "basePrice")
		delete(additionalProperties, "priceItems")
		delete(additionalProperties, "hasTierPrices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistSetPriceListItem) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PricelistSetPriceListItem) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePricelistSetPriceListItem struct {
	value *PricelistSetPriceListItem
	isSet bool
}

func (v NullablePricelistSetPriceListItem) Get() *PricelistSetPriceListItem {
	return v.value
}

func (v *NullablePricelistSetPriceListItem) Set(val *PricelistSetPriceListItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistSetPriceListItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistSetPriceListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistSetPriceListItem(val *PricelistSetPriceListItem) *NullablePricelistSetPriceListItem {
	return &NullablePricelistSetPriceListItem{value: val, isSet: true}
}

func (v NullablePricelistSetPriceListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistSetPriceListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
