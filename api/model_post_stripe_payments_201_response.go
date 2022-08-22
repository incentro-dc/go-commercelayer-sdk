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

// POSTStripePayments201Response struct for POSTStripePayments201Response
type POSTStripePayments201Response struct {
	Data *POSTStripePayments201ResponseData `json:"data,omitempty"`
}

// NewPOSTStripePayments201Response instantiates a new POSTStripePayments201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStripePayments201Response() *POSTStripePayments201Response {
	this := POSTStripePayments201Response{}
	return &this
}

// NewPOSTStripePayments201ResponseWithDefaults instantiates a new POSTStripePayments201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStripePayments201ResponseWithDefaults() *POSTStripePayments201Response {
	this := POSTStripePayments201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTStripePayments201Response) GetData() POSTStripePayments201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTStripePayments201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripePayments201Response) GetDataOk() (*POSTStripePayments201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTStripePayments201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTStripePayments201ResponseData and assigns it to the Data field.
func (o *POSTStripePayments201Response) SetData(v POSTStripePayments201ResponseData) {
	o.Data = &v
}

func (o POSTStripePayments201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTStripePayments201Response struct {
	value *POSTStripePayments201Response
	isSet bool
}

func (v NullablePOSTStripePayments201Response) Get() *POSTStripePayments201Response {
	return v.value
}

func (v *NullablePOSTStripePayments201Response) Set(val *POSTStripePayments201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStripePayments201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStripePayments201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStripePayments201Response(val *POSTStripePayments201Response) *NullablePOSTStripePayments201Response {
	return &NullablePOSTStripePayments201Response{value: val, isSet: true}
}

func (v NullablePOSTStripePayments201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStripePayments201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
