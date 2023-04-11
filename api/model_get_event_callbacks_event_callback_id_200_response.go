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

// GETEventCallbacksEventCallbackId200Response struct for GETEventCallbacksEventCallbackId200Response
type GETEventCallbacksEventCallbackId200Response struct {
	Data *GETEventCallbacks200ResponseDataInner `json:"data,omitempty"`
}

// NewGETEventCallbacksEventCallbackId200Response instantiates a new GETEventCallbacksEventCallbackId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacksEventCallbackId200Response() *GETEventCallbacksEventCallbackId200Response {
	this := GETEventCallbacksEventCallbackId200Response{}
	return &this
}

// NewGETEventCallbacksEventCallbackId200ResponseWithDefaults instantiates a new GETEventCallbacksEventCallbackId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacksEventCallbackId200ResponseWithDefaults() *GETEventCallbacksEventCallbackId200Response {
	this := GETEventCallbacksEventCallbackId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETEventCallbacksEventCallbackId200Response) GetData() GETEventCallbacks200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETEventCallbacks200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacksEventCallbackId200Response) GetDataOk() (*GETEventCallbacks200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETEventCallbacksEventCallbackId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETEventCallbacks200ResponseDataInner and assigns it to the Data field.
func (o *GETEventCallbacksEventCallbackId200Response) SetData(v GETEventCallbacks200ResponseDataInner) {
	o.Data = &v
}

func (o GETEventCallbacksEventCallbackId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETEventCallbacksEventCallbackId200Response struct {
	value *GETEventCallbacksEventCallbackId200Response
	isSet bool
}

func (v NullableGETEventCallbacksEventCallbackId200Response) Get() *GETEventCallbacksEventCallbackId200Response {
	return v.value
}

func (v *NullableGETEventCallbacksEventCallbackId200Response) Set(val *GETEventCallbacksEventCallbackId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacksEventCallbackId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacksEventCallbackId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacksEventCallbackId200Response(val *GETEventCallbacksEventCallbackId200Response) *NullableGETEventCallbacksEventCallbackId200Response {
	return &NullableGETEventCallbacksEventCallbackId200Response{value: val, isSet: true}
}

func (v NullableGETEventCallbacksEventCallbackId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacksEventCallbackId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
