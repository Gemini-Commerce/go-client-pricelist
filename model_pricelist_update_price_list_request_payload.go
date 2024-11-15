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

// checks if the PricelistUpdatePriceListRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistUpdatePriceListRequestPayload{}

// PricelistUpdatePriceListRequestPayload struct for PricelistUpdatePriceListRequestPayload
type PricelistUpdatePriceListRequestPayload struct {
	IsActive *bool `json:"isActive,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	Currency *PricelistCurrency `json:"currency,omitempty"`
	VatIncluded *bool `json:"vatIncluded,omitempty"`
	DeliveredDutyPaid *bool `json:"deliveredDutyPaid,omitempty"`
	Segments []string `json:"segments,omitempty"`
	Markets []string `json:"markets,omitempty"`
	Channels []string `json:"channels,omitempty"`
	Type *PricelistPriceListType `json:"type,omitempty"`
	IsSystem *bool `json:"isSystem,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistUpdatePriceListRequestPayload PricelistUpdatePriceListRequestPayload

// NewPricelistUpdatePriceListRequestPayload instantiates a new PricelistUpdatePriceListRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistUpdatePriceListRequestPayload() *PricelistUpdatePriceListRequestPayload {
	this := PricelistUpdatePriceListRequestPayload{}
	var currency PricelistCurrency = PRICELISTCURRENCY_XXX
	this.Currency = &currency
	var type_ PricelistPriceListType = PRICELISTPRICELISTTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewPricelistUpdatePriceListRequestPayloadWithDefaults instantiates a new PricelistUpdatePriceListRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistUpdatePriceListRequestPayloadWithDefaults() *PricelistUpdatePriceListRequestPayload {
	this := PricelistUpdatePriceListRequestPayload{}
	var currency PricelistCurrency = PRICELISTCURRENCY_XXX
	this.Currency = &currency
	var type_ PricelistPriceListType = PRICELISTPRICELISTTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PricelistUpdatePriceListRequestPayload) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *PricelistUpdatePriceListRequestPayload) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetCurrency() PricelistCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret PricelistCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetCurrencyOk() (*PricelistCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given PricelistCurrency and assigns it to the Currency field.
func (o *PricelistUpdatePriceListRequestPayload) SetCurrency(v PricelistCurrency) {
	o.Currency = &v
}

// GetVatIncluded returns the VatIncluded field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetVatIncluded() bool {
	if o == nil || IsNil(o.VatIncluded) {
		var ret bool
		return ret
	}
	return *o.VatIncluded
}

// GetVatIncludedOk returns a tuple with the VatIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetVatIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatIncluded) {
		return nil, false
	}
	return o.VatIncluded, true
}

// HasVatIncluded returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasVatIncluded() bool {
	if o != nil && !IsNil(o.VatIncluded) {
		return true
	}

	return false
}

// SetVatIncluded gets a reference to the given bool and assigns it to the VatIncluded field.
func (o *PricelistUpdatePriceListRequestPayload) SetVatIncluded(v bool) {
	o.VatIncluded = &v
}

// GetDeliveredDutyPaid returns the DeliveredDutyPaid field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetDeliveredDutyPaid() bool {
	if o == nil || IsNil(o.DeliveredDutyPaid) {
		var ret bool
		return ret
	}
	return *o.DeliveredDutyPaid
}

// GetDeliveredDutyPaidOk returns a tuple with the DeliveredDutyPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetDeliveredDutyPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveredDutyPaid) {
		return nil, false
	}
	return o.DeliveredDutyPaid, true
}

// HasDeliveredDutyPaid returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasDeliveredDutyPaid() bool {
	if o != nil && !IsNil(o.DeliveredDutyPaid) {
		return true
	}

	return false
}

// SetDeliveredDutyPaid gets a reference to the given bool and assigns it to the DeliveredDutyPaid field.
func (o *PricelistUpdatePriceListRequestPayload) SetDeliveredDutyPaid(v bool) {
	o.DeliveredDutyPaid = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetSegments() []string {
	if o == nil || IsNil(o.Segments) {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *PricelistUpdatePriceListRequestPayload) SetSegments(v []string) {
	o.Segments = v
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetMarkets() []string {
	if o == nil || IsNil(o.Markets) {
		var ret []string
		return ret
	}
	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetMarketsOk() ([]string, bool) {
	if o == nil || IsNil(o.Markets) {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasMarkets() bool {
	if o != nil && !IsNil(o.Markets) {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given []string and assigns it to the Markets field.
func (o *PricelistUpdatePriceListRequestPayload) SetMarkets(v []string) {
	o.Markets = v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetChannels() []string {
	if o == nil || IsNil(o.Channels) {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *PricelistUpdatePriceListRequestPayload) SetChannels(v []string) {
	o.Channels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetType() PricelistPriceListType {
	if o == nil || IsNil(o.Type) {
		var ret PricelistPriceListType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetTypeOk() (*PricelistPriceListType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PricelistPriceListType and assigns it to the Type field.
func (o *PricelistUpdatePriceListRequestPayload) SetType(v PricelistPriceListType) {
	o.Type = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasIsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *PricelistUpdatePriceListRequestPayload) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PricelistUpdatePriceListRequestPayload) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PricelistUpdatePriceListRequestPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistUpdatePriceListRequestPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PricelistUpdatePriceListRequestPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PricelistUpdatePriceListRequestPayload) SetDescription(v string) {
	o.Description = &v
}

func (o PricelistUpdatePriceListRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistUpdatePriceListRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.VatIncluded) {
		toSerialize["vatIncluded"] = o.VatIncluded
	}
	if !IsNil(o.DeliveredDutyPaid) {
		toSerialize["deliveredDutyPaid"] = o.DeliveredDutyPaid
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Markets) {
		toSerialize["markets"] = o.Markets
	}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsSystem) {
		toSerialize["isSystem"] = o.IsSystem
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistUpdatePriceListRequestPayload) UnmarshalJSON(data []byte) (err error) {
	varPricelistUpdatePriceListRequestPayload := _PricelistUpdatePriceListRequestPayload{}

	err = json.Unmarshal(data, &varPricelistUpdatePriceListRequestPayload)

	if err != nil {
		return err
	}

	*o = PricelistUpdatePriceListRequestPayload(varPricelistUpdatePriceListRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "vatIncluded")
		delete(additionalProperties, "deliveredDutyPaid")
		delete(additionalProperties, "segments")
		delete(additionalProperties, "markets")
		delete(additionalProperties, "channels")
		delete(additionalProperties, "type")
		delete(additionalProperties, "isSystem")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistUpdatePriceListRequestPayload) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PricelistUpdatePriceListRequestPayload) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePricelistUpdatePriceListRequestPayload struct {
	value *PricelistUpdatePriceListRequestPayload
	isSet bool
}

func (v NullablePricelistUpdatePriceListRequestPayload) Get() *PricelistUpdatePriceListRequestPayload {
	return v.value
}

func (v *NullablePricelistUpdatePriceListRequestPayload) Set(val *PricelistUpdatePriceListRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistUpdatePriceListRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistUpdatePriceListRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistUpdatePriceListRequestPayload(val *PricelistUpdatePriceListRequestPayload) *NullablePricelistUpdatePriceListRequestPayload {
	return &NullablePricelistUpdatePriceListRequestPayload{value: val, isSet: true}
}

func (v NullablePricelistUpdatePriceListRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistUpdatePriceListRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


