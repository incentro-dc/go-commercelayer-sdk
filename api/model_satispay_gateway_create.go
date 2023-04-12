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

// checks if the SatispayGatewayCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayGatewayCreate{}

// SatispayGatewayCreate struct for SatispayGatewayCreate
type SatispayGatewayCreate struct {
	Data SatispayGatewayCreateData `json:"data"`
}

// NewSatispayGatewayCreate instantiates a new SatispayGatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGatewayCreate(data SatispayGatewayCreateData) *SatispayGatewayCreate {
	this := SatispayGatewayCreate{}
	this.Data = data
	return &this
}

// NewSatispayGatewayCreateWithDefaults instantiates a new SatispayGatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayCreateWithDefaults() *SatispayGatewayCreate {
	this := SatispayGatewayCreate{}
	return &this
}

// GetData returns the Data field value
func (o *SatispayGatewayCreate) GetData() SatispayGatewayCreateData {
	if o == nil {
		var ret SatispayGatewayCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SatispayGatewayCreate) GetDataOk() (*SatispayGatewayCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SatispayGatewayCreate) SetData(v SatispayGatewayCreateData) {
	o.Data = v
}

func (o SatispayGatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayGatewayCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSatispayGatewayCreate struct {
	value *SatispayGatewayCreate
	isSet bool
}

func (v NullableSatispayGatewayCreate) Get() *SatispayGatewayCreate {
	return v.value
}

func (v *NullableSatispayGatewayCreate) Set(val *SatispayGatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGatewayCreate(val *SatispayGatewayCreate) *NullableSatispayGatewayCreate {
	return &NullableSatispayGatewayCreate{value: val, isSet: true}
}

func (v NullableSatispayGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
