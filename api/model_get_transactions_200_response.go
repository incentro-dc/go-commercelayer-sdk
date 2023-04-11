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

// GETTransactions200Response struct for GETTransactions200Response
type GETTransactions200Response struct {
	Data []GETTransactions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETTransactions200Response instantiates a new GETTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTransactions200Response() *GETTransactions200Response {
	this := GETTransactions200Response{}
	return &this
}

// NewGETTransactions200ResponseWithDefaults instantiates a new GETTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTransactions200ResponseWithDefaults() *GETTransactions200Response {
	this := GETTransactions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETTransactions200Response) GetData() []GETTransactions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETTransactions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTransactions200Response) GetDataOk() ([]GETTransactions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETTransactions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETTransactions200ResponseDataInner and assigns it to the Data field.
func (o *GETTransactions200Response) SetData(v []GETTransactions200ResponseDataInner) {
	o.Data = v
}

func (o GETTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETTransactions200Response struct {
	value *GETTransactions200Response
	isSet bool
}

func (v NullableGETTransactions200Response) Get() *GETTransactions200Response {
	return v.value
}

func (v *NullableGETTransactions200Response) Set(val *GETTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTransactions200Response(val *GETTransactions200Response) *NullableGETTransactions200Response {
	return &NullableGETTransactions200Response{value: val, isSet: true}
}

func (v NullableGETTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
