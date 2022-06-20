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

// ModelPackage struct for ModelPackage
type ModelPackage struct {
	Data PackageData `json:"data"`
}

// NewModelPackage instantiates a new ModelPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPackage(data PackageData) *ModelPackage {
	this := ModelPackage{}
	this.Data = data
	return &this
}

// NewModelPackageWithDefaults instantiates a new ModelPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPackageWithDefaults() *ModelPackage {
	this := ModelPackage{}
	return &this
}

// GetData returns the Data field value
func (o *ModelPackage) GetData() PackageData {
	if o == nil {
		var ret PackageData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ModelPackage) GetDataOk() (*PackageData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ModelPackage) SetData(v PackageData) {
	o.Data = v
}

func (o ModelPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableModelPackage struct {
	value *ModelPackage
	isSet bool
}

func (v NullableModelPackage) Get() *ModelPackage {
	return v.value
}

func (v *NullableModelPackage) Set(val *ModelPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPackage(val *ModelPackage) *NullableModelPackage {
	return &NullableModelPackage{value: val, isSet: true}
}

func (v NullableModelPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
