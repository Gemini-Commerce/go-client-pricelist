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

// checks if the ListPriceListsRequestFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPriceListsRequestFilter{}

// ListPriceListsRequestFilter struct for ListPriceListsRequestFilter
type ListPriceListsRequestFilter struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	CurrencyFilter *PricelistCurrencyFilter `json:"currencyFilter,omitempty"`
	VatIncluded *bool `json:"vatIncluded,omitempty"`
	DeliveredDutyPaid *bool `json:"deliveredDutyPaid,omitempty"`
	SegmentsFilter *PricelistSegmentFilter `json:"segmentsFilter,omitempty"`
	MarketsFilter *PricelistMarketFilter `json:"marketsFilter,omitempty"`
	ChannelsFilter *PricelistChannelFilter `json:"channelsFilter,omitempty"`
	TypeFilter *PricelistPriceListTypeFilter `json:"typeFilter,omitempty"`
	IsSystem *bool `json:"isSystem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPriceListsRequestFilter ListPriceListsRequestFilter

// NewListPriceListsRequestFilter instantiates a new ListPriceListsRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPriceListsRequestFilter() *ListPriceListsRequestFilter {
	this := ListPriceListsRequestFilter{}
	return &this
}

// NewListPriceListsRequestFilterWithDefaults instantiates a new ListPriceListsRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPriceListsRequestFilterWithDefaults() *ListPriceListsRequestFilter {
	this := ListPriceListsRequestFilter{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// &#39;Has&#39;Code returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;Code() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ListPriceListsRequestFilter) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// &#39;Has&#39;Name returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;Name() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListPriceListsRequestFilter) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// &#39;Has&#39;Description returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;Description() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListPriceListsRequestFilter) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// &#39;Has&#39;IsActive returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;IsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ListPriceListsRequestFilter) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// &#39;Has&#39;IsDefault returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;IsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ListPriceListsRequestFilter) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetCurrencyFilter returns the CurrencyFilter field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetCurrencyFilter() PricelistCurrencyFilter {
	if o == nil || IsNil(o.CurrencyFilter) {
		var ret PricelistCurrencyFilter
		return ret
	}
	return *o.CurrencyFilter
}

// GetCurrencyFilterOk returns a tuple with the CurrencyFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetCurrencyFilterOk() (*PricelistCurrencyFilter, bool) {
	if o == nil || IsNil(o.CurrencyFilter) {
		return nil, false
	}
	return o.CurrencyFilter, true
}

// &#39;Has&#39;CurrencyFilter returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;CurrencyFilter() bool {
	if o != nil && !IsNil(o.CurrencyFilter) {
		return true
	}

	return false
}

// SetCurrencyFilter gets a reference to the given PricelistCurrencyFilter and assigns it to the CurrencyFilter field.
func (o *ListPriceListsRequestFilter) SetCurrencyFilter(v PricelistCurrencyFilter) {
	o.CurrencyFilter = &v
}

// GetVatIncluded returns the VatIncluded field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetVatIncluded() bool {
	if o == nil || IsNil(o.VatIncluded) {
		var ret bool
		return ret
	}
	return *o.VatIncluded
}

// GetVatIncludedOk returns a tuple with the VatIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetVatIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatIncluded) {
		return nil, false
	}
	return o.VatIncluded, true
}

// &#39;Has&#39;VatIncluded returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;VatIncluded() bool {
	if o != nil && !IsNil(o.VatIncluded) {
		return true
	}

	return false
}

// SetVatIncluded gets a reference to the given bool and assigns it to the VatIncluded field.
func (o *ListPriceListsRequestFilter) SetVatIncluded(v bool) {
	o.VatIncluded = &v
}

// GetDeliveredDutyPaid returns the DeliveredDutyPaid field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetDeliveredDutyPaid() bool {
	if o == nil || IsNil(o.DeliveredDutyPaid) {
		var ret bool
		return ret
	}
	return *o.DeliveredDutyPaid
}

// GetDeliveredDutyPaidOk returns a tuple with the DeliveredDutyPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetDeliveredDutyPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.DeliveredDutyPaid) {
		return nil, false
	}
	return o.DeliveredDutyPaid, true
}

// &#39;Has&#39;DeliveredDutyPaid returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;DeliveredDutyPaid() bool {
	if o != nil && !IsNil(o.DeliveredDutyPaid) {
		return true
	}

	return false
}

// SetDeliveredDutyPaid gets a reference to the given bool and assigns it to the DeliveredDutyPaid field.
func (o *ListPriceListsRequestFilter) SetDeliveredDutyPaid(v bool) {
	o.DeliveredDutyPaid = &v
}

// GetSegmentsFilter returns the SegmentsFilter field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetSegmentsFilter() PricelistSegmentFilter {
	if o == nil || IsNil(o.SegmentsFilter) {
		var ret PricelistSegmentFilter
		return ret
	}
	return *o.SegmentsFilter
}

// GetSegmentsFilterOk returns a tuple with the SegmentsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetSegmentsFilterOk() (*PricelistSegmentFilter, bool) {
	if o == nil || IsNil(o.SegmentsFilter) {
		return nil, false
	}
	return o.SegmentsFilter, true
}

// &#39;Has&#39;SegmentsFilter returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;SegmentsFilter() bool {
	if o != nil && !IsNil(o.SegmentsFilter) {
		return true
	}

	return false
}

// SetSegmentsFilter gets a reference to the given PricelistSegmentFilter and assigns it to the SegmentsFilter field.
func (o *ListPriceListsRequestFilter) SetSegmentsFilter(v PricelistSegmentFilter) {
	o.SegmentsFilter = &v
}

// GetMarketsFilter returns the MarketsFilter field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetMarketsFilter() PricelistMarketFilter {
	if o == nil || IsNil(o.MarketsFilter) {
		var ret PricelistMarketFilter
		return ret
	}
	return *o.MarketsFilter
}

// GetMarketsFilterOk returns a tuple with the MarketsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetMarketsFilterOk() (*PricelistMarketFilter, bool) {
	if o == nil || IsNil(o.MarketsFilter) {
		return nil, false
	}
	return o.MarketsFilter, true
}

// &#39;Has&#39;MarketsFilter returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;MarketsFilter() bool {
	if o != nil && !IsNil(o.MarketsFilter) {
		return true
	}

	return false
}

// SetMarketsFilter gets a reference to the given PricelistMarketFilter and assigns it to the MarketsFilter field.
func (o *ListPriceListsRequestFilter) SetMarketsFilter(v PricelistMarketFilter) {
	o.MarketsFilter = &v
}

// GetChannelsFilter returns the ChannelsFilter field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetChannelsFilter() PricelistChannelFilter {
	if o == nil || IsNil(o.ChannelsFilter) {
		var ret PricelistChannelFilter
		return ret
	}
	return *o.ChannelsFilter
}

// GetChannelsFilterOk returns a tuple with the ChannelsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetChannelsFilterOk() (*PricelistChannelFilter, bool) {
	if o == nil || IsNil(o.ChannelsFilter) {
		return nil, false
	}
	return o.ChannelsFilter, true
}

// &#39;Has&#39;ChannelsFilter returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;ChannelsFilter() bool {
	if o != nil && !IsNil(o.ChannelsFilter) {
		return true
	}

	return false
}

// SetChannelsFilter gets a reference to the given PricelistChannelFilter and assigns it to the ChannelsFilter field.
func (o *ListPriceListsRequestFilter) SetChannelsFilter(v PricelistChannelFilter) {
	o.ChannelsFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetTypeFilter() PricelistPriceListTypeFilter {
	if o == nil || IsNil(o.TypeFilter) {
		var ret PricelistPriceListTypeFilter
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetTypeFilterOk() (*PricelistPriceListTypeFilter, bool) {
	if o == nil || IsNil(o.TypeFilter) {
		return nil, false
	}
	return o.TypeFilter, true
}

// &#39;Has&#39;TypeFilter returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;TypeFilter() bool {
	if o != nil && !IsNil(o.TypeFilter) {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given PricelistPriceListTypeFilter and assigns it to the TypeFilter field.
func (o *ListPriceListsRequestFilter) SetTypeFilter(v PricelistPriceListTypeFilter) {
	o.TypeFilter = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *ListPriceListsRequestFilter) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPriceListsRequestFilter) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// &#39;Has&#39;IsSystem returns a boolean if a field has been set.
func (o *ListPriceListsRequestFilter) &#39;Has&#39;IsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *ListPriceListsRequestFilter) SetIsSystem(v bool) {
	o.IsSystem = &v
}

func (o ListPriceListsRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPriceListsRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.CurrencyFilter) {
		toSerialize["currencyFilter"] = o.CurrencyFilter
	}
	if !IsNil(o.VatIncluded) {
		toSerialize["vatIncluded"] = o.VatIncluded
	}
	if !IsNil(o.DeliveredDutyPaid) {
		toSerialize["deliveredDutyPaid"] = o.DeliveredDutyPaid
	}
	if !IsNil(o.SegmentsFilter) {
		toSerialize["segmentsFilter"] = o.SegmentsFilter
	}
	if !IsNil(o.MarketsFilter) {
		toSerialize["marketsFilter"] = o.MarketsFilter
	}
	if !IsNil(o.ChannelsFilter) {
		toSerialize["channelsFilter"] = o.ChannelsFilter
	}
	if !IsNil(o.TypeFilter) {
		toSerialize["typeFilter"] = o.TypeFilter
	}
	if !IsNil(o.IsSystem) {
		toSerialize["isSystem"] = o.IsSystem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPriceListsRequestFilter) UnmarshalJSON(data []byte) (err error) {
	varListPriceListsRequestFilter := _ListPriceListsRequestFilter{}

	err = json.Unmarshal(data, &varListPriceListsRequestFilter)

	if err != nil {
		return err
	}

	*o = ListPriceListsRequestFilter(varListPriceListsRequestFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "currencyFilter")
		delete(additionalProperties, "vatIncluded")
		delete(additionalProperties, "deliveredDutyPaid")
		delete(additionalProperties, "segmentsFilter")
		delete(additionalProperties, "marketsFilter")
		delete(additionalProperties, "channelsFilter")
		delete(additionalProperties, "typeFilter")
		delete(additionalProperties, "isSystem")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ListPriceListsRequestFilter) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ListPriceListsRequestFilter) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableListPriceListsRequestFilter struct {
	value *ListPriceListsRequestFilter
	isSet bool
}

func (v NullableListPriceListsRequestFilter) Get() *ListPriceListsRequestFilter {
	return v.value
}

func (v *NullableListPriceListsRequestFilter) Set(val *ListPriceListsRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableListPriceListsRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableListPriceListsRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPriceListsRequestFilter(val *ListPriceListsRequestFilter) *NullableListPriceListsRequestFilter {
	return &NullableListPriceListsRequestFilter{value: val, isSet: true}
}

func (v NullableListPriceListsRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPriceListsRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


