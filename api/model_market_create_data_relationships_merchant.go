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

// MarketCreateDataRelationshipsMerchant struct for MarketCreateDataRelationshipsMerchant
type MarketCreateDataRelationshipsMerchant struct {
	Data MarketDataRelationshipsMerchantData `json:"data"`
}

// NewMarketCreateDataRelationshipsMerchant instantiates a new MarketCreateDataRelationshipsMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketCreateDataRelationshipsMerchant(data MarketDataRelationshipsMerchantData) *MarketCreateDataRelationshipsMerchant {
	this := MarketCreateDataRelationshipsMerchant{}
	this.Data = data
	return &this
}

// NewMarketCreateDataRelationshipsMerchantWithDefaults instantiates a new MarketCreateDataRelationshipsMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketCreateDataRelationshipsMerchantWithDefaults() *MarketCreateDataRelationshipsMerchant {
	this := MarketCreateDataRelationshipsMerchant{}
	return &this
}

// GetData returns the Data field value
func (o *MarketCreateDataRelationshipsMerchant) GetData() MarketDataRelationshipsMerchantData {
	if o == nil {
		var ret MarketDataRelationshipsMerchantData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MarketCreateDataRelationshipsMerchant) GetDataOk() (*MarketDataRelationshipsMerchantData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MarketCreateDataRelationshipsMerchant) SetData(v MarketDataRelationshipsMerchantData) {
	o.Data = v
}

func (o MarketCreateDataRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketCreateDataRelationshipsMerchant struct {
	value *MarketCreateDataRelationshipsMerchant
	isSet bool
}

func (v NullableMarketCreateDataRelationshipsMerchant) Get() *MarketCreateDataRelationshipsMerchant {
	return v.value
}

func (v *NullableMarketCreateDataRelationshipsMerchant) Set(val *MarketCreateDataRelationshipsMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketCreateDataRelationshipsMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketCreateDataRelationshipsMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketCreateDataRelationshipsMerchant(val *MarketCreateDataRelationshipsMerchant) *NullableMarketCreateDataRelationshipsMerchant {
	return &NullableMarketCreateDataRelationshipsMerchant{value: val, isSet: true}
}

func (v NullableMarketCreateDataRelationshipsMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketCreateDataRelationshipsMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
