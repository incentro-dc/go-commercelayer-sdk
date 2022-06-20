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

// StockTransferData struct for StockTransferData
type StockTransferData struct {
	// The resource's type
	Type          string                          `json:"type"`
	Attributes    StockTransferDataAttributes     `json:"attributes"`
	Relationships *StockTransferDataRelationships `json:"relationships,omitempty"`
}

// NewStockTransferData instantiates a new StockTransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferData(type_ string, attributes StockTransferDataAttributes) *StockTransferData {
	this := StockTransferData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockTransferDataWithDefaults instantiates a new StockTransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferDataWithDefaults() *StockTransferData {
	this := StockTransferData{}
	var type_ string = "stock_transfers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StockTransferData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StockTransferData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockTransferData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockTransferData) GetAttributes() StockTransferDataAttributes {
	if o == nil {
		var ret StockTransferDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockTransferData) GetAttributesOk() (*StockTransferDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockTransferData) SetAttributes(v StockTransferDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockTransferData) GetRelationships() StockTransferDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StockTransferDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferData) GetRelationshipsOk() (*StockTransferDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockTransferData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockTransferDataRelationships and assigns it to the Relationships field.
func (o *StockTransferData) SetRelationships(v StockTransferDataRelationships) {
	o.Relationships = &v
}

func (o StockTransferData) MarshalJSON() ([]byte, error) {
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

type NullableStockTransferData struct {
	value *StockTransferData
	isSet bool
}

func (v NullableStockTransferData) Get() *StockTransferData {
	return v.value
}

func (v *NullableStockTransferData) Set(val *StockTransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferData(val *StockTransferData) *NullableStockTransferData {
	return &NullableStockTransferData{value: val, isSet: true}
}

func (v NullableStockTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
