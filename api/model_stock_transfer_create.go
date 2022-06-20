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

// StockTransferCreate struct for StockTransferCreate
type StockTransferCreate struct {
	Data StockTransferCreateData `json:"data"`
}

// NewStockTransferCreate instantiates a new StockTransferCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferCreate(data StockTransferCreateData) *StockTransferCreate {
	this := StockTransferCreate{}
	this.Data = data
	return &this
}

// NewStockTransferCreateWithDefaults instantiates a new StockTransferCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferCreateWithDefaults() *StockTransferCreate {
	this := StockTransferCreate{}
	return &this
}

// GetData returns the Data field value
func (o *StockTransferCreate) GetData() StockTransferCreateData {
	if o == nil {
		var ret StockTransferCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockTransferCreate) GetDataOk() (*StockTransferCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockTransferCreate) SetData(v StockTransferCreateData) {
	o.Data = v
}

func (o StockTransferCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferCreate struct {
	value *StockTransferCreate
	isSet bool
}

func (v NullableStockTransferCreate) Get() *StockTransferCreate {
	return v.value
}

func (v *NullableStockTransferCreate) Set(val *StockTransferCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferCreate(val *StockTransferCreate) *NullableStockTransferCreate {
	return &NullableStockTransferCreate{value: val, isSet: true}
}

func (v NullableStockTransferCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
