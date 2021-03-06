/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleGeocoder struct for GoogleGeocoder
type GoogleGeocoder struct {
	Data GoogleGeocoderData `json:"data"`
}

// NewGoogleGeocoder instantiates a new GoogleGeocoder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoder(data GoogleGeocoderData) *GoogleGeocoder {
	this := GoogleGeocoder{}
	this.Data = data
	return &this
}

// NewGoogleGeocoderWithDefaults instantiates a new GoogleGeocoder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderWithDefaults() *GoogleGeocoder {
	this := GoogleGeocoder{}
	return &this
}

// GetData returns the Data field value
func (o *GoogleGeocoder) GetData() GoogleGeocoderData {
	if o == nil {
		var ret GoogleGeocoderData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoder) GetDataOk() (*GoogleGeocoderData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GoogleGeocoder) SetData(v GoogleGeocoderData) {
	o.Data = v
}

func (o GoogleGeocoder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleGeocoder struct {
	value *GoogleGeocoder
	isSet bool
}

func (v NullableGoogleGeocoder) Get() *GoogleGeocoder {
	return v.value
}

func (v *NullableGoogleGeocoder) Set(val *GoogleGeocoder) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoder) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoder(val *GoogleGeocoder) *NullableGoogleGeocoder {
	return &NullableGoogleGeocoder{value: val, isSet: true}
}

func (v NullableGoogleGeocoder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
