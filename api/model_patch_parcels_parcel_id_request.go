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

// checks if the PATCHParcelsParcelIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHParcelsParcelIdRequest{}

// PATCHParcelsParcelIdRequest struct for PATCHParcelsParcelIdRequest
type PATCHParcelsParcelIdRequest struct {
	Data PATCHParcelsParcelIdRequestData `json:"data"`
}

// NewPATCHParcelsParcelIdRequest instantiates a new PATCHParcelsParcelIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHParcelsParcelIdRequest(data PATCHParcelsParcelIdRequestData) *PATCHParcelsParcelIdRequest {
	this := PATCHParcelsParcelIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHParcelsParcelIdRequestWithDefaults instantiates a new PATCHParcelsParcelIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHParcelsParcelIdRequestWithDefaults() *PATCHParcelsParcelIdRequest {
	this := PATCHParcelsParcelIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHParcelsParcelIdRequest) GetData() PATCHParcelsParcelIdRequestData {
	if o == nil {
		var ret PATCHParcelsParcelIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelIdRequest) GetDataOk() (*PATCHParcelsParcelIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHParcelsParcelIdRequest) SetData(v PATCHParcelsParcelIdRequestData) {
	o.Data = v
}

func (o PATCHParcelsParcelIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHParcelsParcelIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHParcelsParcelIdRequest struct {
	value *PATCHParcelsParcelIdRequest
	isSet bool
}

func (v NullablePATCHParcelsParcelIdRequest) Get() *PATCHParcelsParcelIdRequest {
	return v.value
}

func (v *NullablePATCHParcelsParcelIdRequest) Set(val *PATCHParcelsParcelIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHParcelsParcelIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHParcelsParcelIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHParcelsParcelIdRequest(val *PATCHParcelsParcelIdRequest) *NullablePATCHParcelsParcelIdRequest {
	return &NullablePATCHParcelsParcelIdRequest{value: val, isSet: true}
}

func (v NullablePATCHParcelsParcelIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHParcelsParcelIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}