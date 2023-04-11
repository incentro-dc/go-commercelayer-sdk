/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MarketDataRelationshipsTaxCalculator struct for MarketDataRelationshipsTaxCalculator
type MarketDataRelationshipsTaxCalculator struct {
	Data *MarketDataRelationshipsTaxCalculatorData `json:"data,omitempty"`
}

// NewMarketDataRelationshipsTaxCalculator instantiates a new MarketDataRelationshipsTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketDataRelationshipsTaxCalculator() *MarketDataRelationshipsTaxCalculator {
	this := MarketDataRelationshipsTaxCalculator{}
	return &this
}

// NewMarketDataRelationshipsTaxCalculatorWithDefaults instantiates a new MarketDataRelationshipsTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataRelationshipsTaxCalculatorWithDefaults() *MarketDataRelationshipsTaxCalculator {
	this := MarketDataRelationshipsTaxCalculator{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MarketDataRelationshipsTaxCalculator) GetData() MarketDataRelationshipsTaxCalculatorData {
	if o == nil || o.Data == nil {
		var ret MarketDataRelationshipsTaxCalculatorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataRelationshipsTaxCalculator) GetDataOk() (*MarketDataRelationshipsTaxCalculatorData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MarketDataRelationshipsTaxCalculator) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given MarketDataRelationshipsTaxCalculatorData and assigns it to the Data field.
func (o *MarketDataRelationshipsTaxCalculator) SetData(v MarketDataRelationshipsTaxCalculatorData) {
	o.Data = &v
}

func (o MarketDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketDataRelationshipsTaxCalculator struct {
	value *MarketDataRelationshipsTaxCalculator
	isSet bool
}

func (v NullableMarketDataRelationshipsTaxCalculator) Get() *MarketDataRelationshipsTaxCalculator {
	return v.value
}

func (v *NullableMarketDataRelationshipsTaxCalculator) Set(val *MarketDataRelationshipsTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketDataRelationshipsTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketDataRelationshipsTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketDataRelationshipsTaxCalculator(val *MarketDataRelationshipsTaxCalculator) *NullableMarketDataRelationshipsTaxCalculator {
	return &NullableMarketDataRelationshipsTaxCalculator{value: val, isSet: true}
}

func (v NullableMarketDataRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketDataRelationshipsTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
