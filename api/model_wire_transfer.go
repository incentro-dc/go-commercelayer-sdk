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

// WireTransfer struct for WireTransfer
type WireTransfer struct {
	Data WireTransferData `json:"data"`
}

// NewWireTransfer instantiates a new WireTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransfer(data WireTransferData) *WireTransfer {
	this := WireTransfer{}
	this.Data = data
	return &this
}

// NewWireTransferWithDefaults instantiates a new WireTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransferWithDefaults() *WireTransfer {
	this := WireTransfer{}
	return &this
}

// GetData returns the Data field value
func (o *WireTransfer) GetData() WireTransferData {
	if o == nil {
		var ret WireTransferData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WireTransfer) GetDataOk() (*WireTransferData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WireTransfer) SetData(v WireTransferData) {
	o.Data = v
}

func (o WireTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableWireTransfer struct {
	value *WireTransfer
	isSet bool
}

func (v NullableWireTransfer) Get() *WireTransfer {
	return v.value
}

func (v *NullableWireTransfer) Set(val *WireTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransfer(val *WireTransfer) *NullableWireTransfer {
	return &NullableWireTransfer{value: val, isSet: true}
}

func (v NullableWireTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}