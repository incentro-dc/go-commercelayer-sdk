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

// GETCleanupsCleanupId200Response struct for GETCleanupsCleanupId200Response
type GETCleanupsCleanupId200Response struct {
	Data *GETCleanups200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCleanupsCleanupId200Response instantiates a new GETCleanupsCleanupId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCleanupsCleanupId200Response() *GETCleanupsCleanupId200Response {
	this := GETCleanupsCleanupId200Response{}
	return &this
}

// NewGETCleanupsCleanupId200ResponseWithDefaults instantiates a new GETCleanupsCleanupId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCleanupsCleanupId200ResponseWithDefaults() *GETCleanupsCleanupId200Response {
	this := GETCleanupsCleanupId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCleanupsCleanupId200Response) GetData() GETCleanups200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETCleanups200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCleanupsCleanupId200Response) GetDataOk() (*GETCleanups200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCleanupsCleanupId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCleanups200ResponseDataInner and assigns it to the Data field.
func (o *GETCleanupsCleanupId200Response) SetData(v GETCleanups200ResponseDataInner) {
	o.Data = &v
}

func (o GETCleanupsCleanupId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCleanupsCleanupId200Response struct {
	value *GETCleanupsCleanupId200Response
	isSet bool
}

func (v NullableGETCleanupsCleanupId200Response) Get() *GETCleanupsCleanupId200Response {
	return v.value
}

func (v *NullableGETCleanupsCleanupId200Response) Set(val *GETCleanupsCleanupId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCleanupsCleanupId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCleanupsCleanupId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCleanupsCleanupId200Response(val *GETCleanupsCleanupId200Response) *NullableGETCleanupsCleanupId200Response {
	return &NullableGETCleanupsCleanupId200Response{value: val, isSet: true}
}

func (v NullableGETCleanupsCleanupId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCleanupsCleanupId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
