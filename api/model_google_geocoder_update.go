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

// GoogleGeocoderUpdate struct for GoogleGeocoderUpdate
type GoogleGeocoderUpdate struct {
	Data GoogleGeocoderUpdateData `json:"data"`
}

// NewGoogleGeocoderUpdate instantiates a new GoogleGeocoderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoderUpdate(data GoogleGeocoderUpdateData) *GoogleGeocoderUpdate {
	this := GoogleGeocoderUpdate{}
	this.Data = data
	return &this
}

// NewGoogleGeocoderUpdateWithDefaults instantiates a new GoogleGeocoderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderUpdateWithDefaults() *GoogleGeocoderUpdate {
	this := GoogleGeocoderUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *GoogleGeocoderUpdate) GetData() GoogleGeocoderUpdateData {
	if o == nil {
		var ret GoogleGeocoderUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderUpdate) GetDataOk() (*GoogleGeocoderUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GoogleGeocoderUpdate) SetData(v GoogleGeocoderUpdateData) {
	o.Data = v
}

func (o GoogleGeocoderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleGeocoderUpdate struct {
	value *GoogleGeocoderUpdate
	isSet bool
}

func (v NullableGoogleGeocoderUpdate) Get() *GoogleGeocoderUpdate {
	return v.value
}

func (v *NullableGoogleGeocoderUpdate) Set(val *GoogleGeocoderUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoderUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoderUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoderUpdate(val *GoogleGeocoderUpdate) *NullableGoogleGeocoderUpdate {
	return &NullableGoogleGeocoderUpdate{value: val, isSet: true}
}

func (v NullableGoogleGeocoderUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoderUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
