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

// POSTAdjustments201Response struct for POSTAdjustments201Response
type POSTAdjustments201Response struct {
	Data *POSTAdjustments201ResponseData `json:"data,omitempty"`
}

// NewPOSTAdjustments201Response instantiates a new POSTAdjustments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdjustments201Response() *POSTAdjustments201Response {
	this := POSTAdjustments201Response{}
	return &this
}

// NewPOSTAdjustments201ResponseWithDefaults instantiates a new POSTAdjustments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdjustments201ResponseWithDefaults() *POSTAdjustments201Response {
	this := POSTAdjustments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAdjustments201Response) GetData() POSTAdjustments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTAdjustments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdjustments201Response) GetDataOk() (*POSTAdjustments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAdjustments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAdjustments201ResponseData and assigns it to the Data field.
func (o *POSTAdjustments201Response) SetData(v POSTAdjustments201ResponseData) {
	o.Data = &v
}

func (o POSTAdjustments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAdjustments201Response struct {
	value *POSTAdjustments201Response
	isSet bool
}

func (v NullablePOSTAdjustments201Response) Get() *POSTAdjustments201Response {
	return v.value
}

func (v *NullablePOSTAdjustments201Response) Set(val *POSTAdjustments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdjustments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdjustments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdjustments201Response(val *POSTAdjustments201Response) *NullablePOSTAdjustments201Response {
	return &NullablePOSTAdjustments201Response{value: val, isSet: true}
}

func (v NullablePOSTAdjustments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdjustments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
