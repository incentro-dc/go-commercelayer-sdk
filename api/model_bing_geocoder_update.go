/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BingGeocoderUpdate struct for BingGeocoderUpdate
type BingGeocoderUpdate struct {
	Data BingGeocoderUpdateData `json:"data"`
}

// NewBingGeocoderUpdate instantiates a new BingGeocoderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderUpdate(data BingGeocoderUpdateData) *BingGeocoderUpdate {
	this := BingGeocoderUpdate{}
	this.Data = data
	return &this
}

// NewBingGeocoderUpdateWithDefaults instantiates a new BingGeocoderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderUpdateWithDefaults() *BingGeocoderUpdate {
	this := BingGeocoderUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *BingGeocoderUpdate) GetData() BingGeocoderUpdateData {
	if o == nil {
		var ret BingGeocoderUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BingGeocoderUpdate) GetDataOk() (*BingGeocoderUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BingGeocoderUpdate) SetData(v BingGeocoderUpdateData) {
	o.Data = v
}

func (o BingGeocoderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoderUpdate struct {
	value *BingGeocoderUpdate
	isSet bool
}

func (v NullableBingGeocoderUpdate) Get() *BingGeocoderUpdate {
	return v.value
}

func (v *NullableBingGeocoderUpdate) Set(val *BingGeocoderUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderUpdate(val *BingGeocoderUpdate) *NullableBingGeocoderUpdate {
	return &NullableBingGeocoderUpdate{value: val, isSet: true}
}

func (v NullableBingGeocoderUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
