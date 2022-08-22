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

// GETOrders200Response struct for GETOrders200Response
type GETOrders200Response struct {
	Data []GETOrders200ResponseDataInner `json:"data,omitempty"`
}

// NewGETOrders200Response instantiates a new GETOrders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200Response() *GETOrders200Response {
	this := GETOrders200Response{}
	return &this
}

// NewGETOrders200ResponseWithDefaults instantiates a new GETOrders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseWithDefaults() *GETOrders200Response {
	this := GETOrders200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200Response) GetData() []GETOrders200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrders200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200Response) GetDataOk() ([]GETOrders200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrders200ResponseDataInner and assigns it to the Data field.
func (o *GETOrders200Response) SetData(v []GETOrders200ResponseDataInner) {
	o.Data = v
}

func (o GETOrders200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200Response struct {
	value *GETOrders200Response
	isSet bool
}

func (v NullableGETOrders200Response) Get() *GETOrders200Response {
	return v.value
}

func (v *NullableGETOrders200Response) Set(val *GETOrders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200Response(val *GETOrders200Response) *NullableGETOrders200Response {
	return &NullableGETOrders200Response{value: val, isSet: true}
}

func (v NullableGETOrders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
