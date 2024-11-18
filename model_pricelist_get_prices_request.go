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

// checks if the PricelistGetPricesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistGetPricesRequest{}

// PricelistGetPricesRequest struct for PricelistGetPricesRequest
type PricelistGetPricesRequest struct {
	TenantId             *string                `json:"tenantId,omitempty"`
	ItemsGrn             []string               `json:"itemsGrn,omitempty"`
	Context              *PricelistPriceContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistGetPricesRequest PricelistGetPricesRequest

// NewPricelistGetPricesRequest instantiates a new PricelistGetPricesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistGetPricesRequest() *PricelistGetPricesRequest {
	this := PricelistGetPricesRequest{}
	return &this
}

// NewPricelistGetPricesRequestWithDefaults instantiates a new PricelistGetPricesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistGetPricesRequestWithDefaults() *PricelistGetPricesRequest {
	this := PricelistGetPricesRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PricelistGetPricesRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPricesRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PricelistGetPricesRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PricelistGetPricesRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetItemsGrn returns the ItemsGrn field value if set, zero value otherwise.
func (o *PricelistGetPricesRequest) GetItemsGrn() []string {
	if o == nil || IsNil(o.ItemsGrn) {
		var ret []string
		return ret
	}
	return o.ItemsGrn
}

// GetItemsGrnOk returns a tuple with the ItemsGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPricesRequest) GetItemsGrnOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsGrn) {
		return nil, false
	}
	return o.ItemsGrn, true
}

// HasItemsGrn returns a boolean if a field has been set.
func (o *PricelistGetPricesRequest) HasItemsGrn() bool {
	if o != nil && !IsNil(o.ItemsGrn) {
		return true
	}

	return false
}

// SetItemsGrn gets a reference to the given []string and assigns it to the ItemsGrn field.
func (o *PricelistGetPricesRequest) SetItemsGrn(v []string) {
	o.ItemsGrn = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PricelistGetPricesRequest) GetContext() PricelistPriceContext {
	if o == nil || IsNil(o.Context) {
		var ret PricelistPriceContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistGetPricesRequest) GetContextOk() (*PricelistPriceContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PricelistGetPricesRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given PricelistPriceContext and assigns it to the Context field.
func (o *PricelistGetPricesRequest) SetContext(v PricelistPriceContext) {
	o.Context = &v
}

func (o PricelistGetPricesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistGetPricesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.ItemsGrn) {
		toSerialize["itemsGrn"] = o.ItemsGrn
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistGetPricesRequest) UnmarshalJSON(data []byte) (err error) {
	varPricelistGetPricesRequest := _PricelistGetPricesRequest{}

	err = json.Unmarshal(data, &varPricelistGetPricesRequest)

	if err != nil {
		return err
	}

	*o = PricelistGetPricesRequest(varPricelistGetPricesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "itemsGrn")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistGetPricesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PricelistGetPricesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePricelistGetPricesRequest struct {
	value *PricelistGetPricesRequest
	isSet bool
}

func (v NullablePricelistGetPricesRequest) Get() *PricelistGetPricesRequest {
	return v.value
}

func (v *NullablePricelistGetPricesRequest) Set(val *PricelistGetPricesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistGetPricesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistGetPricesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistGetPricesRequest(val *PricelistGetPricesRequest) *NullablePricelistGetPricesRequest {
	return &NullablePricelistGetPricesRequest{value: val, isSet: true}
}

func (v NullablePricelistGetPricesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistGetPricesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
