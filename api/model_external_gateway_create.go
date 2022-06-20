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

// ExternalGatewayCreate struct for ExternalGatewayCreate
type ExternalGatewayCreate struct {
	Data ExternalGatewayCreateData `json:"data"`
}

// NewExternalGatewayCreate instantiates a new ExternalGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGatewayCreate(data ExternalGatewayCreateData) *ExternalGatewayCreate {
	this := ExternalGatewayCreate{}
	this.Data = data
	return &this
}

// NewExternalGatewayCreateWithDefaults instantiates a new ExternalGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGatewayCreateWithDefaults() *ExternalGatewayCreate {
	this := ExternalGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalGatewayCreate) GetData() ExternalGatewayCreateData {
	if o == nil {
		var ret ExternalGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalGatewayCreate) GetDataOk() (*ExternalGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalGatewayCreate) SetData(v ExternalGatewayCreateData) {
	o.Data = v
}

func (o ExternalGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGatewayCreate struct {
	value *ExternalGatewayCreate
	isSet bool
}

func (v NullableExternalGatewayCreate) Get() *ExternalGatewayCreate {
	return v.value
}

func (v *NullableExternalGatewayCreate) Set(val *ExternalGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGatewayCreate(val *ExternalGatewayCreate) *NullableExternalGatewayCreate {
	return &NullableExternalGatewayCreate{value: val, isSet: true}
}

func (v NullableExternalGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
