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
	"time"
)

// checks if the PricelistGetPriceListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetPriceListItem{}

// PricelistGetPriceListItem struct for PricelistGetPriceListItem
type PricelistGetPriceListItem struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	ItemGrn *string `json:"itemGrn,omitempty"`
	Price *PricelistMoney `json:"price,omitempty"`
	DoubleFormatPrice *float64 `json:"doubleFormatPrice,omitempty"`
	EndDatePrice *time.Time `json:"endDatePrice,omitempty"`
	BasePrice *PricelistMoney `json:"basePrice,omitempty"`
	DoubleFormatBasePrice *float64 `json:"doubleFormatBasePrice,omitempty"`
	Currency *PricelistCurrency `json:"currency,omitempty"`
	HasTierPrices *bool `json:"hasTierPrices,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewPricelistGetPriceListItem instantiates a new PricelistGetPriceListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetPriceListItem() *PricelistGetPriceListItem {
	this := PricelistGetPriceListItem{}
	var currency PricelistCurrency = PRICELISTCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// NewPricelistGetPriceListItemWithDefaults instantiates a new PricelistGetPriceListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetPriceListItemWithDefaults() *PricelistGetPriceListItem {
	this := PricelistGetPriceListItem{}
	var currency PricelistCurrency = PRICELISTCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PricelistGetPriceListItem) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *PricelistGetPriceListItem) SetGrn(v string) {
	o.Grn = &v
}

// GetItemGrn returns the ItemGrn field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetItemGrn() string {
	if o == nil || IsNil(o.ItemGrn) {
		var ret string
		return ret
	}
	return *o.ItemGrn
}

// GetItemGrnOk returns a tuple with the ItemGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetItemGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ItemGrn) {
		return nil, false
	}
	return o.ItemGrn, true
}

// HasItemGrn returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasItemGrn() bool {
	if o != nil && !IsNil(o.ItemGrn) {
		return true
	}

	return false
}

// SetItemGrn gets a reference to the given string and assigns it to the ItemGrn field.
func (o *PricelistGetPriceListItem) SetItemGrn(v string) {
	o.ItemGrn = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetPrice() PricelistMoney {
	if o == nil || IsNil(o.Price) {
		var ret PricelistMoney
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetPriceOk() (*PricelistMoney, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PricelistMoney and assigns it to the Price field.
func (o *PricelistGetPriceListItem) SetPrice(v PricelistMoney) {
	o.Price = &v
}

// GetDoubleFormatPrice returns the DoubleFormatPrice field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetDoubleFormatPrice() float64 {
	if o == nil || IsNil(o.DoubleFormatPrice) {
		var ret float64
		return ret
	}
	return *o.DoubleFormatPrice
}

// GetDoubleFormatPriceOk returns a tuple with the DoubleFormatPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetDoubleFormatPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.DoubleFormatPrice) {
		return nil, false
	}
	return o.DoubleFormatPrice, true
}

// HasDoubleFormatPrice returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasDoubleFormatPrice() bool {
	if o != nil && !IsNil(o.DoubleFormatPrice) {
		return true
	}

	return false
}

// SetDoubleFormatPrice gets a reference to the given float64 and assigns it to the DoubleFormatPrice field.
func (o *PricelistGetPriceListItem) SetDoubleFormatPrice(v float64) {
	o.DoubleFormatPrice = &v
}

// GetEndDatePrice returns the EndDatePrice field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetEndDatePrice() time.Time {
	if o == nil || IsNil(o.EndDatePrice) {
		var ret time.Time
		return ret
	}
	return *o.EndDatePrice
}

// GetEndDatePriceOk returns a tuple with the EndDatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetEndDatePriceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDatePrice) {
		return nil, false
	}
	return o.EndDatePrice, true
}

// HasEndDatePrice returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasEndDatePrice() bool {
	if o != nil && !IsNil(o.EndDatePrice) {
		return true
	}

	return false
}

// SetEndDatePrice gets a reference to the given time.Time and assigns it to the EndDatePrice field.
func (o *PricelistGetPriceListItem) SetEndDatePrice(v time.Time) {
	o.EndDatePrice = &v
}

// GetBasePrice returns the BasePrice field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetBasePrice() PricelistMoney {
	if o == nil || IsNil(o.BasePrice) {
		var ret PricelistMoney
		return ret
	}
	return *o.BasePrice
}

// GetBasePriceOk returns a tuple with the BasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetBasePriceOk() (*PricelistMoney, bool) {
	if o == nil || IsNil(o.BasePrice) {
		return nil, false
	}
	return o.BasePrice, true
}

// HasBasePrice returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasBasePrice() bool {
	if o != nil && !IsNil(o.BasePrice) {
		return true
	}

	return false
}

// SetBasePrice gets a reference to the given PricelistMoney and assigns it to the BasePrice field.
func (o *PricelistGetPriceListItem) SetBasePrice(v PricelistMoney) {
	o.BasePrice = &v
}

// GetDoubleFormatBasePrice returns the DoubleFormatBasePrice field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetDoubleFormatBasePrice() float64 {
	if o == nil || IsNil(o.DoubleFormatBasePrice) {
		var ret float64
		return ret
	}
	return *o.DoubleFormatBasePrice
}

// GetDoubleFormatBasePriceOk returns a tuple with the DoubleFormatBasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetDoubleFormatBasePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.DoubleFormatBasePrice) {
		return nil, false
	}
	return o.DoubleFormatBasePrice, true
}

// HasDoubleFormatBasePrice returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasDoubleFormatBasePrice() bool {
	if o != nil && !IsNil(o.DoubleFormatBasePrice) {
		return true
	}

	return false
}

// SetDoubleFormatBasePrice gets a reference to the given float64 and assigns it to the DoubleFormatBasePrice field.
func (o *PricelistGetPriceListItem) SetDoubleFormatBasePrice(v float64) {
	o.DoubleFormatBasePrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetCurrency() PricelistCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret PricelistCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetCurrencyOk() (*PricelistCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given PricelistCurrency and assigns it to the Currency field.
func (o *PricelistGetPriceListItem) SetCurrency(v PricelistCurrency) {
	o.Currency = &v
}

// GetHasTierPrices returns the HasTierPrices field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetHasTierPrices() bool {
	if o == nil || IsNil(o.HasTierPrices) {
		var ret bool
		return ret
	}
	return *o.HasTierPrices
}

// GetHasTierPricesOk returns a tuple with the HasTierPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetHasTierPricesOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTierPrices) {
		return nil, false
	}
	return o.HasTierPrices, true
}

// HasHasTierPrices returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasHasTierPrices() bool {
	if o != nil && !IsNil(o.HasTierPrices) {
		return true
	}

	return false
}

// SetHasTierPrices gets a reference to the given bool and assigns it to the HasTierPrices field.
func (o *PricelistGetPriceListItem) SetHasTierPrices(v bool) {
	o.HasTierPrices = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PricelistGetPriceListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PricelistGetPriceListItem) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PricelistGetPriceListItem) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PricelistGetPriceListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PricelistGetPriceListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetPriceListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.ItemGrn) {
		toSerialize["itemGrn"] = o.ItemGrn
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.DoubleFormatPrice) {
		toSerialize["doubleFormatPrice"] = o.DoubleFormatPrice
	}
	if !IsNil(o.EndDatePrice) {
		toSerialize["endDatePrice"] = o.EndDatePrice
	}
	if !IsNil(o.BasePrice) {
		toSerialize["basePrice"] = o.BasePrice
	}
	if !IsNil(o.DoubleFormatBasePrice) {
		toSerialize["doubleFormatBasePrice"] = o.DoubleFormatBasePrice
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.HasTierPrices) {
		toSerialize["hasTierPrices"] = o.HasTierPrices
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePricelistGetPriceListItem struct {
	value *PricelistGetPriceListItem
	isSet bool
}

func (v NullablePricelistGetPriceListItem) Get() *PricelistGetPriceListItem {
	return v.value
}

func (v *NullablePricelistGetPriceListItem) Set(val *PricelistGetPriceListItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetPriceListItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetPriceListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetPriceListItem(val *PricelistGetPriceListItem) *NullablePricelistGetPriceListItem {
	return &NullablePricelistGetPriceListItem{value: val, isSet: true}
}

func (v NullablePricelistGetPriceListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetPriceListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


