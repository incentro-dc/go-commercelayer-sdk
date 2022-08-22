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

// POSTExternalPayments201Response struct for POSTExternalPayments201Response
type POSTExternalPayments201Response struct {
	Data *POSTExternalPayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTExternalPayments201Response instantiates a new POSTExternalPayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPayments201Response() *POSTExternalPayments201Response {
	this := POSTExternalPayments201Response{}
	return &this
}

// NewPOSTExternalPayments201ResponseWithDefaults instantiates a new POSTExternalPayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPayments201ResponseWithDefaults() *POSTExternalPayments201Response {
	this := POSTExternalPayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalPayments201Response) GetData() POSTExternalPayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTExternalPayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPayments201Response) GetDataOk() (*POSTExternalPayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalPayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalPayments201ResponseData and assigns it to the Data field.
func (o *POSTExternalPayments201Response) SetData(v POSTExternalPayments201ResponseData) {
	o.Data = &v
}

func (o POSTExternalPayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTExternalPayments201Response struct {
	value *POSTExternalPayments201Response
	isSet bool
}

func (v NullablePOSTExternalPayments201Response) Get() *POSTExternalPayments201Response {
	return v.value
}

func (v *NullablePOSTExternalPayments201Response) Set(val *POSTExternalPayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPayments201Response(val *POSTExternalPayments201Response) *NullablePOSTExternalPayments201Response {
	return &NullablePOSTExternalPayments201Response{value: val, isSet: true}
}

func (v NullablePOSTExternalPayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
