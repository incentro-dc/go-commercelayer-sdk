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

// BundleUpdate struct for BundleUpdate
type BundleUpdate struct {
	Data BundleUpdateData `json:"data"`
}

// NewBundleUpdate instantiates a new BundleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleUpdate(data BundleUpdateData) *BundleUpdate {
	this := BundleUpdate{}
	this.Data = data
	return &this
}

// NewBundleUpdateWithDefaults instantiates a new BundleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleUpdateWithDefaults() *BundleUpdate {
	this := BundleUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *BundleUpdate) GetData() BundleUpdateData {
	if o == nil {
		var ret BundleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleUpdate) GetDataOk() (*BundleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleUpdate) SetData(v BundleUpdateData) {
	o.Data = v
}

func (o BundleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBundleUpdate struct {
	value *BundleUpdate
	isSet bool
}

func (v NullableBundleUpdate) Get() *BundleUpdate {
	return v.value
}

func (v *NullableBundleUpdate) Set(val *BundleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleUpdate(val *BundleUpdate) *NullableBundleUpdate {
	return &NullableBundleUpdate{value: val, isSet: true}
}

func (v NullableBundleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
