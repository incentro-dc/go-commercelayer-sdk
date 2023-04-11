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

// MarketDataRelationshipsMerchant struct for MarketDataRelationshipsMerchant
type MarketDataRelationshipsMerchant struct {
	Data *MarketDataRelationshipsMerchantData `json:"data,omitempty"`
}

// NewMarketDataRelationshipsMerchant instantiates a new MarketDataRelationshipsMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketDataRelationshipsMerchant() *MarketDataRelationshipsMerchant {
	this := MarketDataRelationshipsMerchant{}
	return &this
}

// NewMarketDataRelationshipsMerchantWithDefaults instantiates a new MarketDataRelationshipsMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataRelationshipsMerchantWithDefaults() *MarketDataRelationshipsMerchant {
	this := MarketDataRelationshipsMerchant{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MarketDataRelationshipsMerchant) GetData() MarketDataRelationshipsMerchantData {
	if o == nil || o.Data == nil {
		var ret MarketDataRelationshipsMerchantData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataRelationshipsMerchant) GetDataOk() (*MarketDataRelationshipsMerchantData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MarketDataRelationshipsMerchant) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given MarketDataRelationshipsMerchantData and assigns it to the Data field.
func (o *MarketDataRelationshipsMerchant) SetData(v MarketDataRelationshipsMerchantData) {
	o.Data = &v
}

func (o MarketDataRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketDataRelationshipsMerchant struct {
	value *MarketDataRelationshipsMerchant
	isSet bool
}

func (v NullableMarketDataRelationshipsMerchant) Get() *MarketDataRelationshipsMerchant {
	return v.value
}

func (v *NullableMarketDataRelationshipsMerchant) Set(val *MarketDataRelationshipsMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketDataRelationshipsMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketDataRelationshipsMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketDataRelationshipsMerchant(val *MarketDataRelationshipsMerchant) *NullableMarketDataRelationshipsMerchant {
	return &NullableMarketDataRelationshipsMerchant{value: val, isSet: true}
}

func (v NullableMarketDataRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketDataRelationshipsMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
