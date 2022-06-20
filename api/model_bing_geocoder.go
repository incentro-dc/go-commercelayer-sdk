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

// BingGeocoder struct for BingGeocoder
type BingGeocoder struct {
	Data BingGeocoderData `json:"data"`
}

// NewBingGeocoder instantiates a new BingGeocoder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoder(data BingGeocoderData) *BingGeocoder {
	this := BingGeocoder{}
	this.Data = data
	return &this
}

// NewBingGeocoderWithDefaults instantiates a new BingGeocoder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderWithDefaults() *BingGeocoder {
	this := BingGeocoder{}
	return &this
}

// GetData returns the Data field value
func (o *BingGeocoder) GetData() BingGeocoderData {
	if o == nil {
		var ret BingGeocoderData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BingGeocoder) GetDataOk() (*BingGeocoderData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BingGeocoder) SetData(v BingGeocoderData) {
	o.Data = v
}

func (o BingGeocoder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoder struct {
	value *BingGeocoder
	isSet bool
}

func (v NullableBingGeocoder) Get() *BingGeocoder {
	return v.value
}

func (v *NullableBingGeocoder) Set(val *BingGeocoder) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoder) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoder(val *BingGeocoder) *NullableBingGeocoder {
	return &NullableBingGeocoder{value: val, isSet: true}
}

func (v NullableBingGeocoder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
