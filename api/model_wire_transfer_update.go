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

// WireTransferUpdate struct for WireTransferUpdate
type WireTransferUpdate struct {
	Data WireTransferUpdateData `json:"data"`
}

// NewWireTransferUpdate instantiates a new WireTransferUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransferUpdate(data WireTransferUpdateData) *WireTransferUpdate {
	this := WireTransferUpdate{}
	this.Data = data
	return &this
}

// NewWireTransferUpdateWithDefaults instantiates a new WireTransferUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransferUpdateWithDefaults() *WireTransferUpdate {
	this := WireTransferUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *WireTransferUpdate) GetData() WireTransferUpdateData {
	if o == nil {
		var ret WireTransferUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WireTransferUpdate) GetDataOk() (*WireTransferUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WireTransferUpdate) SetData(v WireTransferUpdateData) {
	o.Data = v
}

func (o WireTransferUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableWireTransferUpdate struct {
	value *WireTransferUpdate
	isSet bool
}

func (v NullableWireTransferUpdate) Get() *WireTransferUpdate {
	return v.value
}

func (v *NullableWireTransferUpdate) Set(val *WireTransferUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransferUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransferUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransferUpdate(val *WireTransferUpdate) *NullableWireTransferUpdate {
	return &NullableWireTransferUpdate{value: val, isSet: true}
}

func (v NullableWireTransferUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransferUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
