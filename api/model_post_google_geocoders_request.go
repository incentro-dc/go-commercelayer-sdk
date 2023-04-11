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

// checks if the POSTGoogleGeocodersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTGoogleGeocodersRequest{}

// POSTGoogleGeocodersRequest struct for POSTGoogleGeocodersRequest
type POSTGoogleGeocodersRequest struct {
	Data POSTGoogleGeocodersRequestData `json:"data"`
}

// NewPOSTGoogleGeocodersRequest instantiates a new POSTGoogleGeocodersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGoogleGeocodersRequest(data POSTGoogleGeocodersRequestData) *POSTGoogleGeocodersRequest {
	this := POSTGoogleGeocodersRequest{}
	this.Data = data
	return &this
}

// NewPOSTGoogleGeocodersRequestWithDefaults instantiates a new POSTGoogleGeocodersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGoogleGeocodersRequestWithDefaults() *POSTGoogleGeocodersRequest {
	this := POSTGoogleGeocodersRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTGoogleGeocodersRequest) GetData() POSTGoogleGeocodersRequestData {
	if o == nil {
		var ret POSTGoogleGeocodersRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTGoogleGeocodersRequest) GetDataOk() (*POSTGoogleGeocodersRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTGoogleGeocodersRequest) SetData(v POSTGoogleGeocodersRequestData) {
	o.Data = v
}

func (o POSTGoogleGeocodersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTGoogleGeocodersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTGoogleGeocodersRequest struct {
	value *POSTGoogleGeocodersRequest
	isSet bool
}

func (v NullablePOSTGoogleGeocodersRequest) Get() *POSTGoogleGeocodersRequest {
	return v.value
}

func (v *NullablePOSTGoogleGeocodersRequest) Set(val *POSTGoogleGeocodersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGoogleGeocodersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGoogleGeocodersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGoogleGeocodersRequest(val *POSTGoogleGeocodersRequest) *NullablePOSTGoogleGeocodersRequest {
	return &NullablePOSTGoogleGeocodersRequest{value: val, isSet: true}
}

func (v NullablePOSTGoogleGeocodersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGoogleGeocodersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}