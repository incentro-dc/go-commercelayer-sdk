/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTPaypalPayments201Response struct for POSTPaypalPayments201Response
type POSTPaypalPayments201Response struct {
	Data *POSTPaypalPayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTPaypalPayments201Response instantiates a new POSTPaypalPayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPaypalPayments201Response() *POSTPaypalPayments201Response {
	this := POSTPaypalPayments201Response{}
	return &this
}

// NewPOSTPaypalPayments201ResponseWithDefaults instantiates a new POSTPaypalPayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPaypalPayments201ResponseWithDefaults() *POSTPaypalPayments201Response {
	this := POSTPaypalPayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPaypalPayments201Response) GetData() POSTPaypalPayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTPaypalPayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPaypalPayments201Response) GetDataOk() (*POSTPaypalPayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPaypalPayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPaypalPayments201ResponseData and assigns it to the Data field.
func (o *POSTPaypalPayments201Response) SetData(v POSTPaypalPayments201ResponseData) {
	o.Data = &v
}

func (o POSTPaypalPayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTPaypalPayments201Response struct {
	value *POSTPaypalPayments201Response
	isSet bool
}

func (v NullablePOSTPaypalPayments201Response) Get() *POSTPaypalPayments201Response {
	return v.value
}

func (v *NullablePOSTPaypalPayments201Response) Set(val *POSTPaypalPayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPaypalPayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPaypalPayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPaypalPayments201Response(val *POSTPaypalPayments201Response) *NullablePOSTPaypalPayments201Response {
	return &NullablePOSTPaypalPayments201Response{value: val, isSet: true}
}

func (v NullablePOSTPaypalPayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPaypalPayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
