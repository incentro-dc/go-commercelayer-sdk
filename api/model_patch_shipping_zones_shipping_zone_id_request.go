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

// checks if the PATCHShippingZonesShippingZoneIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHShippingZonesShippingZoneIdRequest{}

// PATCHShippingZonesShippingZoneIdRequest struct for PATCHShippingZonesShippingZoneIdRequest
type PATCHShippingZonesShippingZoneIdRequest struct {
	Data PATCHShippingZonesShippingZoneIdRequestData `json:"data"`
}

// NewPATCHShippingZonesShippingZoneIdRequest instantiates a new PATCHShippingZonesShippingZoneIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingZonesShippingZoneIdRequest(data PATCHShippingZonesShippingZoneIdRequestData) *PATCHShippingZonesShippingZoneIdRequest {
	this := PATCHShippingZonesShippingZoneIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHShippingZonesShippingZoneIdRequestWithDefaults instantiates a new PATCHShippingZonesShippingZoneIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingZonesShippingZoneIdRequestWithDefaults() *PATCHShippingZonesShippingZoneIdRequest {
	this := PATCHShippingZonesShippingZoneIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHShippingZonesShippingZoneIdRequest) GetData() PATCHShippingZonesShippingZoneIdRequestData {
	if o == nil {
		var ret PATCHShippingZonesShippingZoneIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneIdRequest) GetDataOk() (*PATCHShippingZonesShippingZoneIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHShippingZonesShippingZoneIdRequest) SetData(v PATCHShippingZonesShippingZoneIdRequestData) {
	o.Data = v
}

func (o PATCHShippingZonesShippingZoneIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHShippingZonesShippingZoneIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHShippingZonesShippingZoneIdRequest struct {
	value *PATCHShippingZonesShippingZoneIdRequest
	isSet bool
}

func (v NullablePATCHShippingZonesShippingZoneIdRequest) Get() *PATCHShippingZonesShippingZoneIdRequest {
	return v.value
}

func (v *NullablePATCHShippingZonesShippingZoneIdRequest) Set(val *PATCHShippingZonesShippingZoneIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingZonesShippingZoneIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingZonesShippingZoneIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingZonesShippingZoneIdRequest(val *PATCHShippingZonesShippingZoneIdRequest) *NullablePATCHShippingZonesShippingZoneIdRequest {
	return &NullablePATCHShippingZonesShippingZoneIdRequest{value: val, isSet: true}
}

func (v NullablePATCHShippingZonesShippingZoneIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingZonesShippingZoneIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}