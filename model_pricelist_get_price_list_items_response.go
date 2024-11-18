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

// checks if the PricelistGetPriceListItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetPriceListItemsResponse{}

// PricelistGetPriceListItemsResponse struct for PricelistGetPriceListItemsResponse
type PricelistGetPriceListItemsResponse struct {
	Items                []PricelistGetPriceListItem `json:"items,omitempty"`
	NextPageToken        *string                     `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistGetPriceListItemsResponse PricelistGetPriceListItemsResponse

// NewPricelistGetPriceListItemsResponse instantiates a new PricelistGetPriceListItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetPriceListItemsResponse() *PricelistGetPriceListItemsResponse {
	this := PricelistGetPriceListItemsResponse{}
	return &this
}

// NewPricelistGetPriceListItemsResponseWithDefaults instantiates a new PricelistGetPriceListItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetPriceListItemsResponseWithDefaults() *PricelistGetPriceListItemsResponse {
	this := PricelistGetPriceListItemsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PricelistGetPriceListItemsResponse) GetItems() []PricelistGetPriceListItem {
	if o == nil || IsNil(o.Items) {
		var ret []PricelistGetPriceListItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItemsResponse) GetItemsOk() ([]PricelistGetPriceListItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PricelistGetPriceListItemsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PricelistGetPriceListItem and assigns it to the Items field.
func (o *PricelistGetPriceListItemsResponse) SetItems(v []PricelistGetPriceListItem) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *PricelistGetPriceListItemsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPriceListItemsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *PricelistGetPriceListItemsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *PricelistGetPriceListItemsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o PricelistGetPriceListItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetPriceListItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistGetPriceListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	varPricelistGetPriceListItemsResponse := _PricelistGetPriceListItemsResponse{}

	err = json.Unmarshal(data, &varPricelistGetPriceListItemsResponse)

	if err != nil {
		return err
	}

	*o = PricelistGetPriceListItemsResponse(varPricelistGetPriceListItemsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistGetPriceListItemsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PricelistGetPriceListItemsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePricelistGetPriceListItemsResponse struct {
	value *PricelistGetPriceListItemsResponse
	isSet bool
}

func (v NullablePricelistGetPriceListItemsResponse) Get() *PricelistGetPriceListItemsResponse {
	return v.value
}

func (v *NullablePricelistGetPriceListItemsResponse) Set(val *PricelistGetPriceListItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetPriceListItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetPriceListItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetPriceListItemsResponse(val *PricelistGetPriceListItemsResponse) *NullablePricelistGetPriceListItemsResponse {
	return &NullablePricelistGetPriceListItemsResponse{value: val, isSet: true}
}

func (v NullablePricelistGetPriceListItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetPriceListItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
