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

// Shipment struct for Shipment
type Shipment struct {
	Data *ShipmentData `json:"data,omitempty"`
}

// NewShipment instantiates a new Shipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipment() *Shipment {
	this := Shipment{}
	return &this
}

// NewShipmentWithDefaults instantiates a new Shipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentWithDefaults() *Shipment {
	this := Shipment{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Shipment) GetData() ShipmentData {
	if o == nil || o.Data == nil {
		var ret ShipmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shipment) GetDataOk() (*ShipmentData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Shipment) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShipmentData and assigns it to the Data field.
func (o *Shipment) SetData(v ShipmentData) {
	o.Data = &v
}

func (o Shipment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShipment struct {
	value *Shipment
	isSet bool
}

func (v NullableShipment) Get() *Shipment {
	return v.value
}

func (v *NullableShipment) Set(val *Shipment) {
	v.value = val
	v.isSet = true
}

func (v NullableShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipment(val *Shipment) *NullableShipment {
	return &NullableShipment{value: val, isSet: true}
}

func (v NullableShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
