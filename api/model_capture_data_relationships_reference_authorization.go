/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CaptureDataRelationshipsReferenceAuthorization struct for CaptureDataRelationshipsReferenceAuthorization
type CaptureDataRelationshipsReferenceAuthorization struct {
	Data CaptureDataRelationshipsReferenceAuthorizationData `json:"data"`
}

// NewCaptureDataRelationshipsReferenceAuthorization instantiates a new CaptureDataRelationshipsReferenceAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureDataRelationshipsReferenceAuthorization(data CaptureDataRelationshipsReferenceAuthorizationData) *CaptureDataRelationshipsReferenceAuthorization {
	this := CaptureDataRelationshipsReferenceAuthorization{}
	this.Data = data
	return &this
}

// NewCaptureDataRelationshipsReferenceAuthorizationWithDefaults instantiates a new CaptureDataRelationshipsReferenceAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureDataRelationshipsReferenceAuthorizationWithDefaults() *CaptureDataRelationshipsReferenceAuthorization {
	this := CaptureDataRelationshipsReferenceAuthorization{}
	return &this
}

// GetData returns the Data field value
func (o *CaptureDataRelationshipsReferenceAuthorization) GetData() CaptureDataRelationshipsReferenceAuthorizationData {
	if o == nil {
		var ret CaptureDataRelationshipsReferenceAuthorizationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CaptureDataRelationshipsReferenceAuthorization) GetDataOk() (*CaptureDataRelationshipsReferenceAuthorizationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CaptureDataRelationshipsReferenceAuthorization) SetData(v CaptureDataRelationshipsReferenceAuthorizationData) {
	o.Data = v
}

func (o CaptureDataRelationshipsReferenceAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCaptureDataRelationshipsReferenceAuthorization struct {
	value *CaptureDataRelationshipsReferenceAuthorization
	isSet bool
}

func (v NullableCaptureDataRelationshipsReferenceAuthorization) Get() *CaptureDataRelationshipsReferenceAuthorization {
	return v.value
}

func (v *NullableCaptureDataRelationshipsReferenceAuthorization) Set(val *CaptureDataRelationshipsReferenceAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureDataRelationshipsReferenceAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureDataRelationshipsReferenceAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureDataRelationshipsReferenceAuthorization(val *CaptureDataRelationshipsReferenceAuthorization) *NullableCaptureDataRelationshipsReferenceAuthorization {
	return &NullableCaptureDataRelationshipsReferenceAuthorization{value: val, isSet: true}
}

func (v NullableCaptureDataRelationshipsReferenceAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureDataRelationshipsReferenceAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}