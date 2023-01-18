/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemDataRelationshipsStockLineItemsData struct for LineItemDataRelationshipsStockLineItemsData
type LineItemDataRelationshipsStockLineItemsData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewLineItemDataRelationshipsStockLineItemsData instantiates a new LineItemDataRelationshipsStockLineItemsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsStockLineItemsData() *LineItemDataRelationshipsStockLineItemsData {
	this := LineItemDataRelationshipsStockLineItemsData{}
	return &this
}

// NewLineItemDataRelationshipsStockLineItemsDataWithDefaults instantiates a new LineItemDataRelationshipsStockLineItemsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsStockLineItemsDataWithDefaults() *LineItemDataRelationshipsStockLineItemsData {
	this := LineItemDataRelationshipsStockLineItemsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsStockLineItemsData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsStockLineItemsData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsStockLineItemsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LineItemDataRelationshipsStockLineItemsData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsStockLineItemsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsStockLineItemsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsStockLineItemsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LineItemDataRelationshipsStockLineItemsData) SetId(v string) {
	o.Id = &v
}

func (o LineItemDataRelationshipsStockLineItemsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemDataRelationshipsStockLineItemsData struct {
	value *LineItemDataRelationshipsStockLineItemsData
	isSet bool
}

func (v NullableLineItemDataRelationshipsStockLineItemsData) Get() *LineItemDataRelationshipsStockLineItemsData {
	return v.value
}

func (v *NullableLineItemDataRelationshipsStockLineItemsData) Set(val *LineItemDataRelationshipsStockLineItemsData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsStockLineItemsData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsStockLineItemsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsStockLineItemsData(val *LineItemDataRelationshipsStockLineItemsData) *NullableLineItemDataRelationshipsStockLineItemsData {
	return &NullableLineItemDataRelationshipsStockLineItemsData{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsStockLineItemsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsStockLineItemsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
