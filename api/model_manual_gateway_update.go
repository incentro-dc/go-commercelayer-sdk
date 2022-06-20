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

// ManualGatewayUpdate struct for ManualGatewayUpdate
type ManualGatewayUpdate struct {
	Data ManualGatewayUpdateData `json:"data"`
}

// NewManualGatewayUpdate instantiates a new ManualGatewayUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayUpdate(data ManualGatewayUpdateData) *ManualGatewayUpdate {
	this := ManualGatewayUpdate{}
	this.Data = data
	return &this
}

// NewManualGatewayUpdateWithDefaults instantiates a new ManualGatewayUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayUpdateWithDefaults() *ManualGatewayUpdate {
	this := ManualGatewayUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ManualGatewayUpdate) GetData() ManualGatewayUpdateData {
	if o == nil {
		var ret ManualGatewayUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayUpdate) GetDataOk() (*ManualGatewayUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManualGatewayUpdate) SetData(v ManualGatewayUpdateData) {
	o.Data = v
}

func (o ManualGatewayUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualGatewayUpdate struct {
	value *ManualGatewayUpdate
	isSet bool
}

func (v NullableManualGatewayUpdate) Get() *ManualGatewayUpdate {
	return v.value
}

func (v *NullableManualGatewayUpdate) Set(val *ManualGatewayUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayUpdate(val *ManualGatewayUpdate) *NullableManualGatewayUpdate {
	return &NullableManualGatewayUpdate{value: val, isSet: true}
}

func (v NullableManualGatewayUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
