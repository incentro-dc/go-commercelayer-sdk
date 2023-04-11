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

// StockLocationData struct for StockLocationData
type StockLocationData struct {
	// The resource's type
	Type          interface{}                                     `json:"type"`
	Attributes    GETStockLocations200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *StockLocationDataRelationships                 `json:"relationships,omitempty"`
}

// NewStockLocationData instantiates a new StockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLocationData(type_ interface{}, attributes GETStockLocations200ResponseDataInnerAttributes) *StockLocationData {
	this := StockLocationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockLocationDataWithDefaults instantiates a new StockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLocationDataWithDefaults() *StockLocationData {
	this := StockLocationData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StockLocationData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StockLocationData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockLocationData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockLocationData) GetAttributes() GETStockLocations200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETStockLocations200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockLocationData) GetAttributesOk() (*GETStockLocations200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockLocationData) SetAttributes(v GETStockLocations200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockLocationData) GetRelationships() StockLocationDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StockLocationDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationData) GetRelationshipsOk() (*StockLocationDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockLocationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockLocationDataRelationships and assigns it to the Relationships field.
func (o *StockLocationData) SetRelationships(v StockLocationDataRelationships) {
	o.Relationships = &v
}

func (o StockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
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

type NullableStockLocationData struct {
	value *StockLocationData
	isSet bool
}

func (v NullableStockLocationData) Get() *StockLocationData {
	return v.value
}

func (v *NullableStockLocationData) Set(val *StockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLocationData(val *StockLocationData) *NullableStockLocationData {
	return &NullableStockLocationData{value: val, isSet: true}
}

func (v NullableStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
