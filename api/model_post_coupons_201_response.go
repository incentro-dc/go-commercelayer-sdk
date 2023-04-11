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

// POSTCoupons201Response struct for POSTCoupons201Response
type POSTCoupons201Response struct {
	Data *POSTCoupons201ResponseData `json:"data,omitempty"`
}

// NewPOSTCoupons201Response instantiates a new POSTCoupons201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCoupons201Response() *POSTCoupons201Response {
	this := POSTCoupons201Response{}
	return &this
}

// NewPOSTCoupons201ResponseWithDefaults instantiates a new POSTCoupons201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCoupons201ResponseWithDefaults() *POSTCoupons201Response {
	this := POSTCoupons201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCoupons201Response) GetData() POSTCoupons201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTCoupons201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201Response) GetDataOk() (*POSTCoupons201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCoupons201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCoupons201ResponseData and assigns it to the Data field.
func (o *POSTCoupons201Response) SetData(v POSTCoupons201ResponseData) {
	o.Data = &v
}

func (o POSTCoupons201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTCoupons201Response struct {
	value *POSTCoupons201Response
	isSet bool
}

func (v NullablePOSTCoupons201Response) Get() *POSTCoupons201Response {
	return v.value
}

func (v *NullablePOSTCoupons201Response) Set(val *POSTCoupons201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCoupons201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCoupons201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCoupons201Response(val *POSTCoupons201Response) *NullablePOSTCoupons201Response {
	return &NullablePOSTCoupons201Response{value: val, isSet: true}
}

func (v NullablePOSTCoupons201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCoupons201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
