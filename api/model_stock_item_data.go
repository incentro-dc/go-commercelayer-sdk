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

// StockItemData struct for StockItemData
type StockItemData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    GETStockItems200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *StockItemDataRelationships                 `json:"relationships,omitempty"`
}

// NewStockItemData instantiates a new StockItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemData(type_ interface{}, attributes GETStockItems200ResponseDataInnerAttributes) *StockItemData {
	this := StockItemData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockItemDataWithDefaults instantiates a new StockItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemDataWithDefaults() *StockItemData {
	this := StockItemData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StockItemData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StockItemData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockItemData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockItemData) GetAttributes() GETStockItems200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETStockItems200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockItemData) GetAttributesOk() (*GETStockItems200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockItemData) SetAttributes(v GETStockItems200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockItemData) GetRelationships() StockItemDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StockItemDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemData) GetRelationshipsOk() (*StockItemDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockItemData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockItemDataRelationships and assigns it to the Relationships field.
func (o *StockItemData) SetRelationships(v StockItemDataRelationships) {
	o.Relationships = &v
}

func (o StockItemData) MarshalJSON() ([]byte, error) {
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

type NullableStockItemData struct {
	value *StockItemData
	isSet bool
}

func (v NullableStockItemData) Get() *StockItemData {
	return v.value
}

func (v *NullableStockItemData) Set(val *StockItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemData(val *StockItemData) *NullableStockItemData {
	return &NullableStockItemData{value: val, isSet: true}
}

func (v NullableStockItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
