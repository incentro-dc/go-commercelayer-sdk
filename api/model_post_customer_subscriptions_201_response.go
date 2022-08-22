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

// POSTCustomerSubscriptions201Response struct for POSTCustomerSubscriptions201Response
type POSTCustomerSubscriptions201Response struct {
	Data *POSTCustomerSubscriptions201ResponseData `json:"data,omitempty"`
}

// NewPOSTCustomerSubscriptions201Response instantiates a new POSTCustomerSubscriptions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerSubscriptions201Response() *POSTCustomerSubscriptions201Response {
	this := POSTCustomerSubscriptions201Response{}
	return &this
}

// NewPOSTCustomerSubscriptions201ResponseWithDefaults instantiates a new POSTCustomerSubscriptions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerSubscriptions201ResponseWithDefaults() *POSTCustomerSubscriptions201Response {
	this := POSTCustomerSubscriptions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomerSubscriptions201Response) GetData() POSTCustomerSubscriptions201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTCustomerSubscriptions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomerSubscriptions201Response) GetDataOk() (*POSTCustomerSubscriptions201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomerSubscriptions201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomerSubscriptions201ResponseData and assigns it to the Data field.
func (o *POSTCustomerSubscriptions201Response) SetData(v POSTCustomerSubscriptions201ResponseData) {
	o.Data = &v
}

func (o POSTCustomerSubscriptions201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCustomerSubscriptions201Response struct {
	value *POSTCustomerSubscriptions201Response
	isSet bool
}

func (v NullablePOSTCustomerSubscriptions201Response) Get() *POSTCustomerSubscriptions201Response {
	return v.value
}

func (v *NullablePOSTCustomerSubscriptions201Response) Set(val *POSTCustomerSubscriptions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerSubscriptions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerSubscriptions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerSubscriptions201Response(val *POSTCustomerSubscriptions201Response) *NullablePOSTCustomerSubscriptions201Response {
	return &NullablePOSTCustomerSubscriptions201Response{value: val, isSet: true}
}

func (v NullablePOSTCustomerSubscriptions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerSubscriptions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
