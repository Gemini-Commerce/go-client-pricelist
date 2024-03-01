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

// checks if the PricelistListPriceListsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistListPriceListsResponse{}

// PricelistListPriceListsResponse struct for PricelistListPriceListsResponse
type PricelistListPriceListsResponse struct {
	Pricelists []ListPriceListsResponsePriceList `json:"pricelists,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewPricelistListPriceListsResponse instantiates a new PricelistListPriceListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistListPriceListsResponse() *PricelistListPriceListsResponse {
	this := PricelistListPriceListsResponse{}
	return &this
}

// NewPricelistListPriceListsResponseWithDefaults instantiates a new PricelistListPriceListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistListPriceListsResponseWithDefaults() *PricelistListPriceListsResponse {
	this := PricelistListPriceListsResponse{}
	return &this
}

// GetPricelists returns the Pricelists field value if set, zero value otherwise.
func (o *PricelistListPriceListsResponse) GetPricelists() []ListPriceListsResponsePriceList {
	if o == nil || IsNil(o.Pricelists) {
		var ret []ListPriceListsResponsePriceList
		return ret
	}
	return o.Pricelists
}

// GetPricelistsOk returns a tuple with the Pricelists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistListPriceListsResponse) GetPricelistsOk() ([]ListPriceListsResponsePriceList, bool) {
	if o == nil || IsNil(o.Pricelists) {
		return nil, false
	}
	return o.Pricelists, true
}

// HasPricelists returns a boolean if a field has been set.
func (o *PricelistListPriceListsResponse) HasPricelists() bool {
	if o != nil && !IsNil(o.Pricelists) {
		return true
	}

	return false
}

// SetPricelists gets a reference to the given []ListPriceListsResponsePriceList and assigns it to the Pricelists field.
func (o *PricelistListPriceListsResponse) SetPricelists(v []ListPriceListsResponsePriceList) {
	o.Pricelists = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *PricelistListPriceListsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistListPriceListsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *PricelistListPriceListsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *PricelistListPriceListsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o PricelistListPriceListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistListPriceListsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pricelists) {
		toSerialize["pricelists"] = o.Pricelists
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullablePricelistListPriceListsResponse struct {
	value *PricelistListPriceListsResponse
	isSet bool
}

func (v NullablePricelistListPriceListsResponse) Get() *PricelistListPriceListsResponse {
	return v.value
}

func (v *NullablePricelistListPriceListsResponse) Set(val *PricelistListPriceListsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistListPriceListsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistListPriceListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistListPriceListsResponse(val *PricelistListPriceListsResponse) *NullablePricelistListPriceListsResponse {
	return &NullablePricelistListPriceListsResponse{value: val, isSet: true}
}

func (v NullablePricelistListPriceListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistListPriceListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


