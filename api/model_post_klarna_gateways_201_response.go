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

// POSTKlarnaGateways201Response struct for POSTKlarnaGateways201Response
type POSTKlarnaGateways201Response struct {
	Data *POSTKlarnaGateways201ResponseData `json:"data,omitempty"`
}

// NewPOSTKlarnaGateways201Response instantiates a new POSTKlarnaGateways201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTKlarnaGateways201Response() *POSTKlarnaGateways201Response {
	this := POSTKlarnaGateways201Response{}
	return &this
}

// NewPOSTKlarnaGateways201ResponseWithDefaults instantiates a new POSTKlarnaGateways201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTKlarnaGateways201ResponseWithDefaults() *POSTKlarnaGateways201Response {
	this := POSTKlarnaGateways201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTKlarnaGateways201Response) GetData() POSTKlarnaGateways201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTKlarnaGateways201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGateways201Response) GetDataOk() (*POSTKlarnaGateways201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTKlarnaGateways201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTKlarnaGateways201ResponseData and assigns it to the Data field.
func (o *POSTKlarnaGateways201Response) SetData(v POSTKlarnaGateways201ResponseData) {
	o.Data = &v
}

func (o POSTKlarnaGateways201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTKlarnaGateways201Response struct {
	value *POSTKlarnaGateways201Response
	isSet bool
}

func (v NullablePOSTKlarnaGateways201Response) Get() *POSTKlarnaGateways201Response {
	return v.value
}

func (v *NullablePOSTKlarnaGateways201Response) Set(val *POSTKlarnaGateways201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTKlarnaGateways201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTKlarnaGateways201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTKlarnaGateways201Response(val *POSTKlarnaGateways201Response) *NullablePOSTKlarnaGateways201Response {
	return &NullablePOSTKlarnaGateways201Response{value: val, isSet: true}
}

func (v NullablePOSTKlarnaGateways201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTKlarnaGateways201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
