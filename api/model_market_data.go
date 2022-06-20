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

// MarketData struct for MarketData
type MarketData struct {
	// The resource's type
	Type          string                   `json:"type"`
	Attributes    MarketDataAttributes     `json:"attributes"`
	Relationships *MarketDataRelationships `json:"relationships,omitempty"`
}

// NewMarketData instantiates a new MarketData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketData(type_ string, attributes MarketDataAttributes) *MarketData {
	this := MarketData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewMarketDataWithDefaults instantiates a new MarketData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataWithDefaults() *MarketData {
	this := MarketData{}
	var type_ string = "markets"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *MarketData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarketData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *MarketData) GetAttributes() MarketDataAttributes {
	if o == nil {
		var ret MarketDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MarketData) GetAttributesOk() (*MarketDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MarketData) SetAttributes(v MarketDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MarketData) GetRelationships() MarketDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret MarketDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetRelationshipsOk() (*MarketDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MarketData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given MarketDataRelationships and assigns it to the Relationships field.
func (o *MarketData) SetRelationships(v MarketDataRelationships) {
	o.Relationships = &v
}

func (o MarketData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableMarketData struct {
	value *MarketData
	isSet bool
}

func (v NullableMarketData) Get() *MarketData {
	return v.value
}

func (v *NullableMarketData) Set(val *MarketData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketData(val *MarketData) *NullableMarketData {
	return &NullableMarketData{value: val, isSet: true}
}

func (v NullableMarketData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
