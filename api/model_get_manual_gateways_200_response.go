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

// GETManualGateways200Response struct for GETManualGateways200Response
type GETManualGateways200Response struct {
	Data []GETManualGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETManualGateways200Response instantiates a new GETManualGateways200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETManualGateways200Response() *GETManualGateways200Response {
	this := GETManualGateways200Response{}
	return &this
}

// NewGETManualGateways200ResponseWithDefaults instantiates a new GETManualGateways200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETManualGateways200ResponseWithDefaults() *GETManualGateways200Response {
	this := GETManualGateways200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETManualGateways200Response) GetData() []GETManualGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETManualGateways200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETManualGateways200Response) GetDataOk() ([]GETManualGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETManualGateways200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETManualGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETManualGateways200Response) SetData(v []GETManualGateways200ResponseDataInner) {
	o.Data = v
}

func (o GETManualGateways200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETManualGateways200Response struct {
	value *GETManualGateways200Response
	isSet bool
}

func (v NullableGETManualGateways200Response) Get() *GETManualGateways200Response {
	return v.value
}

func (v *NullableGETManualGateways200Response) Set(val *GETManualGateways200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETManualGateways200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETManualGateways200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETManualGateways200Response(val *GETManualGateways200Response) *NullableGETManualGateways200Response {
	return &NullableGETManualGateways200Response{value: val, isSet: true}
}

func (v NullableGETManualGateways200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETManualGateways200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
