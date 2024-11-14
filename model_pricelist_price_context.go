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

// checks if the PricelistPriceContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricelistPriceContext{}

// PricelistPriceContext struct for PricelistPriceContext
type PricelistPriceContext struct {
	Currency PricelistCurrency `json:"currency"`
	Segment *string `json:"segment,omitempty"`
	Segments []string `json:"segments,omitempty"`
	Market *string `json:"market,omitempty"`
	Channel *string `json:"channel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PricelistPriceContext PricelistPriceContext

// NewPricelistPriceContext instantiates a new PricelistPriceContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricelistPriceContext(currency PricelistCurrency) *PricelistPriceContext {
	this := PricelistPriceContext{}
	this.Currency = currency
	return &this
}

// NewPricelistPriceContextWithDefaults instantiates a new PricelistPriceContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricelistPriceContextWithDefaults() *PricelistPriceContext {
	this := PricelistPriceContext{}
	var currency PricelistCurrency = PRICELISTCURRENCY_XXX
	this.Currency = currency
	return &this
}

// GetCurrency returns the Currency field value
func (o *PricelistPriceContext) GetCurrency() PricelistCurrency {
	if o == nil {
		var ret PricelistCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PricelistPriceContext) GetCurrencyOk() (*PricelistCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PricelistPriceContext) SetCurrency(v PricelistCurrency) {
	o.Currency = v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *PricelistPriceContext) GetSegment() string {
	if o == nil || IsNil(o.Segment) {
		var ret string
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistPriceContext) GetSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// &#39;Has&#39;Segment returns a boolean if a field has been set.
func (o *PricelistPriceContext) &#39;Has&#39;Segment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given string and assigns it to the Segment field.
func (o *PricelistPriceContext) SetSegment(v string) {
	o.Segment = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *PricelistPriceContext) GetSegments() []string {
	if o == nil || IsNil(o.Segments) {
		var ret []string
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistPriceContext) GetSegmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// &#39;Has&#39;Segments returns a boolean if a field has been set.
func (o *PricelistPriceContext) &#39;Has&#39;Segments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []string and assigns it to the Segments field.
func (o *PricelistPriceContext) SetSegments(v []string) {
	o.Segments = v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PricelistPriceContext) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistPriceContext) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// &#39;Has&#39;Market returns a boolean if a field has been set.
func (o *PricelistPriceContext) &#39;Has&#39;Market() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *PricelistPriceContext) SetMarket(v string) {
	o.Market = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *PricelistPriceContext) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricelistPriceContext) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// &#39;Has&#39;Channel returns a boolean if a field has been set.
func (o *PricelistPriceContext) &#39;Has&#39;Channel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *PricelistPriceContext) SetChannel(v string) {
	o.Channel = &v
}

func (o PricelistPriceContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricelistPriceContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PricelistPriceContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPricelistPriceContext := _PricelistPriceContext{}

	err = json.Unmarshal(data, &varPricelistPriceContext)

	if err != nil {
		return err
	}

	*o = PricelistPriceContext(varPricelistPriceContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "segment")
		delete(additionalProperties, "segments")
		delete(additionalProperties, "market")
		delete(additionalProperties, "channel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PricelistPriceContext) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *PricelistPriceContext) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullablePricelistPriceContext struct {
	value *PricelistPriceContext
	isSet bool
}

func (v NullablePricelistPriceContext) Get() *PricelistPriceContext {
	return v.value
}

func (v *NullablePricelistPriceContext) Set(val *PricelistPriceContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePricelistPriceContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePricelistPriceContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricelistPriceContext(val *PricelistPriceContext) *NullablePricelistPriceContext {
	return &NullablePricelistPriceContext{value: val, isSet: true}
}

func (v NullablePricelistPriceContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricelistPriceContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


