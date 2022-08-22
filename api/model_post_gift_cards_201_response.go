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

// POSTGiftCards201Response struct for POSTGiftCards201Response
type POSTGiftCards201Response struct {
	Data *POSTGiftCards201ResponseData `json:"data,omitempty"`
}

// NewPOSTGiftCards201Response instantiates a new POSTGiftCards201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGiftCards201Response() *POSTGiftCards201Response {
	this := POSTGiftCards201Response{}
	return &this
}

// NewPOSTGiftCards201ResponseWithDefaults instantiates a new POSTGiftCards201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGiftCards201ResponseWithDefaults() *POSTGiftCards201Response {
	this := POSTGiftCards201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTGiftCards201Response) GetData() POSTGiftCards201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTGiftCards201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201Response) GetDataOk() (*POSTGiftCards201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTGiftCards201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTGiftCards201ResponseData and assigns it to the Data field.
func (o *POSTGiftCards201Response) SetData(v POSTGiftCards201ResponseData) {
	o.Data = &v
}

func (o POSTGiftCards201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTGiftCards201Response struct {
	value *POSTGiftCards201Response
	isSet bool
}

func (v NullablePOSTGiftCards201Response) Get() *POSTGiftCards201Response {
	return v.value
}

func (v *NullablePOSTGiftCards201Response) Set(val *POSTGiftCards201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGiftCards201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGiftCards201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGiftCards201Response(val *POSTGiftCards201Response) *NullablePOSTGiftCards201Response {
	return &NullablePOSTGiftCards201Response{value: val, isSet: true}
}

func (v NullablePOSTGiftCards201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGiftCards201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
