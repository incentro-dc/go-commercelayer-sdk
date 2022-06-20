/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MarketDataRelationshipsPriceList struct for MarketDataRelationshipsPriceList
type MarketDataRelationshipsPriceList struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewMarketDataRelationshipsPriceList instantiates a new MarketDataRelationshipsPriceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketDataRelationshipsPriceList(type_ string, id string) *MarketDataRelationshipsPriceList {
	this := MarketDataRelationshipsPriceList{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewMarketDataRelationshipsPriceListWithDefaults instantiates a new MarketDataRelationshipsPriceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataRelationshipsPriceListWithDefaults() *MarketDataRelationshipsPriceList {
	this := MarketDataRelationshipsPriceList{}
	var type_ string = "price_lists"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *MarketDataRelationshipsPriceList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarketDataRelationshipsPriceList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketDataRelationshipsPriceList) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *MarketDataRelationshipsPriceList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MarketDataRelationshipsPriceList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MarketDataRelationshipsPriceList) SetId(v string) {
	o.Id = v
}

func (o MarketDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableMarketDataRelationshipsPriceList struct {
	value *MarketDataRelationshipsPriceList
	isSet bool
}

func (v NullableMarketDataRelationshipsPriceList) Get() *MarketDataRelationshipsPriceList {
	return v.value
}

func (v *NullableMarketDataRelationshipsPriceList) Set(val *MarketDataRelationshipsPriceList) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketDataRelationshipsPriceList) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketDataRelationshipsPriceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketDataRelationshipsPriceList(val *MarketDataRelationshipsPriceList) *NullableMarketDataRelationshipsPriceList {
	return &NullableMarketDataRelationshipsPriceList{value: val, isSet: true}
}

func (v NullableMarketDataRelationshipsPriceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketDataRelationshipsPriceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}