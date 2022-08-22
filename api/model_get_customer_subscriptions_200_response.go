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

// GETCustomerSubscriptions200Response struct for GETCustomerSubscriptions200Response
type GETCustomerSubscriptions200Response struct {
	Data []GETCustomerSubscriptions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCustomerSubscriptions200Response instantiates a new GETCustomerSubscriptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerSubscriptions200Response() *GETCustomerSubscriptions200Response {
	this := GETCustomerSubscriptions200Response{}
	return &this
}

// NewGETCustomerSubscriptions200ResponseWithDefaults instantiates a new GETCustomerSubscriptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerSubscriptions200ResponseWithDefaults() *GETCustomerSubscriptions200Response {
	this := GETCustomerSubscriptions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerSubscriptions200Response) GetData() []GETCustomerSubscriptions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCustomerSubscriptions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerSubscriptions200Response) GetDataOk() ([]GETCustomerSubscriptions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerSubscriptions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCustomerSubscriptions200ResponseDataInner and assigns it to the Data field.
func (o *GETCustomerSubscriptions200Response) SetData(v []GETCustomerSubscriptions200ResponseDataInner) {
	o.Data = v
}

func (o GETCustomerSubscriptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerSubscriptions200Response struct {
	value *GETCustomerSubscriptions200Response
	isSet bool
}

func (v NullableGETCustomerSubscriptions200Response) Get() *GETCustomerSubscriptions200Response {
	return v.value
}

func (v *NullableGETCustomerSubscriptions200Response) Set(val *GETCustomerSubscriptions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerSubscriptions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerSubscriptions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerSubscriptions200Response(val *GETCustomerSubscriptions200Response) *NullableGETCustomerSubscriptions200Response {
	return &NullableGETCustomerSubscriptions200Response{value: val, isSet: true}
}

func (v NullableGETCustomerSubscriptions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerSubscriptions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
