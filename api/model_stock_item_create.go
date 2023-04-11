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

// StockItemCreate struct for StockItemCreate
type StockItemCreate struct {
	Data StockItemCreateData `json:"data"`
}

// NewStockItemCreate instantiates a new StockItemCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemCreate(data StockItemCreateData) *StockItemCreate {
	this := StockItemCreate{}
	this.Data = data
	return &this
}

// NewStockItemCreateWithDefaults instantiates a new StockItemCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemCreateWithDefaults() *StockItemCreate {
	this := StockItemCreate{}
	return &this
}

// GetData returns the Data field value
func (o *StockItemCreate) GetData() StockItemCreateData {
	if o == nil {
		var ret StockItemCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockItemCreate) GetDataOk() (*StockItemCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockItemCreate) SetData(v StockItemCreateData) {
	o.Data = v
}

func (o StockItemCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockItemCreate struct {
	value *StockItemCreate
	isSet bool
}

func (v NullableStockItemCreate) Get() *StockItemCreate {
	return v.value
}

func (v *NullableStockItemCreate) Set(val *StockItemCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemCreate(val *StockItemCreate) *NullableStockItemCreate {
	return &NullableStockItemCreate{value: val, isSet: true}
}

func (v NullableStockItemCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
