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

// StockItem struct for StockItem
type StockItem struct {
	Data *StockItemData `json:"data,omitempty"`
}

// NewStockItem instantiates a new StockItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItem() *StockItem {
	this := StockItem{}
	return &this
}

// NewStockItemWithDefaults instantiates a new StockItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemWithDefaults() *StockItem {
	this := StockItem{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StockItem) GetData() StockItemData {
	if o == nil || o.Data == nil {
		var ret StockItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItem) GetDataOk() (*StockItemData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StockItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given StockItemData and assigns it to the Data field.
func (o *StockItem) SetData(v StockItemData) {
	o.Data = &v
}

func (o StockItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockItem struct {
	value *StockItem
	isSet bool
}

func (v NullableStockItem) Get() *StockItem {
	return v.value
}

func (v *NullableStockItem) Set(val *StockItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItem(val *StockItem) *NullableStockItem {
	return &NullableStockItem{value: val, isSet: true}
}

func (v NullableStockItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
