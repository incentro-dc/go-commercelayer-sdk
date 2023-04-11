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

// checks if the PATCHPackagesPackageIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPackagesPackageIdRequest{}

// PATCHPackagesPackageIdRequest struct for PATCHPackagesPackageIdRequest
type PATCHPackagesPackageIdRequest struct {
	Data PATCHPackagesPackageIdRequestData `json:"data"`
}

// NewPATCHPackagesPackageIdRequest instantiates a new PATCHPackagesPackageIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPackagesPackageIdRequest(data PATCHPackagesPackageIdRequestData) *PATCHPackagesPackageIdRequest {
	this := PATCHPackagesPackageIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHPackagesPackageIdRequestWithDefaults instantiates a new PATCHPackagesPackageIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPackagesPackageIdRequestWithDefaults() *PATCHPackagesPackageIdRequest {
	this := PATCHPackagesPackageIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHPackagesPackageIdRequest) GetData() PATCHPackagesPackageIdRequestData {
	if o == nil {
		var ret PATCHPackagesPackageIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHPackagesPackageIdRequest) GetDataOk() (*PATCHPackagesPackageIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHPackagesPackageIdRequest) SetData(v PATCHPackagesPackageIdRequestData) {
	o.Data = v
}

func (o PATCHPackagesPackageIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPackagesPackageIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHPackagesPackageIdRequest struct {
	value *PATCHPackagesPackageIdRequest
	isSet bool
}

func (v NullablePATCHPackagesPackageIdRequest) Get() *PATCHPackagesPackageIdRequest {
	return v.value
}

func (v *NullablePATCHPackagesPackageIdRequest) Set(val *PATCHPackagesPackageIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPackagesPackageIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPackagesPackageIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPackagesPackageIdRequest(val *PATCHPackagesPackageIdRequest) *NullablePATCHPackagesPackageIdRequest {
	return &NullablePATCHPackagesPackageIdRequest{value: val, isSet: true}
}

func (v NullablePATCHPackagesPackageIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPackagesPackageIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}