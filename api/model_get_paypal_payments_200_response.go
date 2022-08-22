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

// GETPaypalPayments200Response struct for GETPaypalPayments200Response
type GETPaypalPayments200Response struct {
	Data []GETPaypalPayments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPaypalPayments200Response instantiates a new GETPaypalPayments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaypalPayments200Response() *GETPaypalPayments200Response {
	this := GETPaypalPayments200Response{}
	return &this
}

// NewGETPaypalPayments200ResponseWithDefaults instantiates a new GETPaypalPayments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaypalPayments200ResponseWithDefaults() *GETPaypalPayments200Response {
	this := GETPaypalPayments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPaypalPayments200Response) GetData() []GETPaypalPayments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETPaypalPayments200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaypalPayments200Response) GetDataOk() ([]GETPaypalPayments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPaypalPayments200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETPaypalPayments200ResponseDataInner and assigns it to the Data field.
func (o *GETPaypalPayments200Response) SetData(v []GETPaypalPayments200ResponseDataInner) {
	o.Data = v
}

func (o GETPaypalPayments200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPaypalPayments200Response struct {
	value *GETPaypalPayments200Response
	isSet bool
}

func (v NullableGETPaypalPayments200Response) Get() *GETPaypalPayments200Response {
	return v.value
}

func (v *NullableGETPaypalPayments200Response) Set(val *GETPaypalPayments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaypalPayments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaypalPayments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaypalPayments200Response(val *GETPaypalPayments200Response) *NullableGETPaypalPayments200Response {
	return &NullableGETPaypalPayments200Response{value: val, isSet: true}
}

func (v NullableGETPaypalPayments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaypalPayments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
