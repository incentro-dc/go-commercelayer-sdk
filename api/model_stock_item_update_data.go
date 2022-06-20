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

// StockItemUpdateData struct for StockItemUpdateData
type StockItemUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                            `json:"id"`
	Attributes    StockItemUpdateDataAttributes     `json:"attributes"`
	Relationships *StockItemUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewStockItemUpdateData instantiates a new StockItemUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemUpdateData(type_ string, id string, attributes StockItemUpdateDataAttributes) *StockItemUpdateData {
	this := StockItemUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewStockItemUpdateDataWithDefaults instantiates a new StockItemUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemUpdateDataWithDefaults() *StockItemUpdateData {
	this := StockItemUpdateData{}
	var type_ string = "stock_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StockItemUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StockItemUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockItemUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *StockItemUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StockItemUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StockItemUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *StockItemUpdateData) GetAttributes() StockItemUpdateDataAttributes {
	if o == nil {
		var ret StockItemUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockItemUpdateData) GetAttributesOk() (*StockItemUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockItemUpdateData) SetAttributes(v StockItemUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockItemUpdateData) GetRelationships() StockItemUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StockItemUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemUpdateData) GetRelationshipsOk() (*StockItemUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockItemUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StockItemUpdateDataRelationships and assigns it to the Relationships field.
func (o *StockItemUpdateData) SetRelationships(v StockItemUpdateDataRelationships) {
	o.Relationships = &v
}

func (o StockItemUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableStockItemUpdateData struct {
	value *StockItemUpdateData
	isSet bool
}

func (v NullableStockItemUpdateData) Get() *StockItemUpdateData {
	return v.value
}

func (v *NullableStockItemUpdateData) Set(val *StockItemUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemUpdateData(val *StockItemUpdateData) *NullableStockItemUpdateData {
	return &NullableStockItemUpdateData{value: val, isSet: true}
}

func (v NullableStockItemUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
