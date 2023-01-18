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

// POSTWebhooks201Response struct for POSTWebhooks201Response
type POSTWebhooks201Response struct {
	Data *POSTWebhooks201ResponseData `json:"data,omitempty"`
}

// NewPOSTWebhooks201Response instantiates a new POSTWebhooks201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTWebhooks201Response() *POSTWebhooks201Response {
	this := POSTWebhooks201Response{}
	return &this
}

// NewPOSTWebhooks201ResponseWithDefaults instantiates a new POSTWebhooks201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTWebhooks201ResponseWithDefaults() *POSTWebhooks201Response {
	this := POSTWebhooks201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTWebhooks201Response) GetData() POSTWebhooks201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTWebhooks201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTWebhooks201Response) GetDataOk() (*POSTWebhooks201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTWebhooks201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTWebhooks201ResponseData and assigns it to the Data field.
func (o *POSTWebhooks201Response) SetData(v POSTWebhooks201ResponseData) {
	o.Data = &v
}

func (o POSTWebhooks201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTWebhooks201Response struct {
	value *POSTWebhooks201Response
	isSet bool
}

func (v NullablePOSTWebhooks201Response) Get() *POSTWebhooks201Response {
	return v.value
}

func (v *NullablePOSTWebhooks201Response) Set(val *POSTWebhooks201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTWebhooks201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTWebhooks201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTWebhooks201Response(val *POSTWebhooks201Response) *NullablePOSTWebhooks201Response {
	return &NullablePOSTWebhooks201Response{value: val, isSet: true}
}

func (v NullablePOSTWebhooks201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTWebhooks201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
