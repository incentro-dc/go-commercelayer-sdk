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

// POSTPriceLists201Response struct for POSTPriceLists201Response
type POSTPriceLists201Response struct {
	Data *POSTPriceLists201ResponseData `json:"data,omitempty"`
}

// NewPOSTPriceLists201Response instantiates a new POSTPriceLists201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceLists201Response() *POSTPriceLists201Response {
	this := POSTPriceLists201Response{}
	return &this
}

// NewPOSTPriceLists201ResponseWithDefaults instantiates a new POSTPriceLists201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceLists201ResponseWithDefaults() *POSTPriceLists201Response {
	this := POSTPriceLists201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTPriceLists201Response) GetData() POSTPriceLists201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTPriceLists201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201Response) GetDataOk() (*POSTPriceLists201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTPriceLists201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTPriceLists201ResponseData and assigns it to the Data field.
func (o *POSTPriceLists201Response) SetData(v POSTPriceLists201ResponseData) {
	o.Data = &v
}

func (o POSTPriceLists201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTPriceLists201Response struct {
	value *POSTPriceLists201Response
	isSet bool
}

func (v NullablePOSTPriceLists201Response) Get() *POSTPriceLists201Response {
	return v.value
}

func (v *NullablePOSTPriceLists201Response) Set(val *POSTPriceLists201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceLists201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceLists201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceLists201Response(val *POSTPriceLists201Response) *NullablePOSTPriceLists201Response {
	return &NullablePOSTPriceLists201Response{value: val, isSet: true}
}

func (v NullablePOSTPriceLists201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceLists201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
