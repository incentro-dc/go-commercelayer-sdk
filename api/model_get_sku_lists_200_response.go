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

// GETSkuLists200Response struct for GETSkuLists200Response
type GETSkuLists200Response struct {
	Data []GETSkuLists200ResponseDataInner `json:"data,omitempty"`
}

// NewGETSkuLists200Response instantiates a new GETSkuLists200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuLists200Response() *GETSkuLists200Response {
	this := GETSkuLists200Response{}
	return &this
}

// NewGETSkuLists200ResponseWithDefaults instantiates a new GETSkuLists200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuLists200ResponseWithDefaults() *GETSkuLists200Response {
	this := GETSkuLists200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETSkuLists200Response) GetData() []GETSkuLists200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETSkuLists200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200Response) GetDataOk() ([]GETSkuLists200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETSkuLists200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETSkuLists200ResponseDataInner and assigns it to the Data field.
func (o *GETSkuLists200Response) SetData(v []GETSkuLists200ResponseDataInner) {
	o.Data = v
}

func (o GETSkuLists200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuLists200Response struct {
	value *GETSkuLists200Response
	isSet bool
}

func (v NullableGETSkuLists200Response) Get() *GETSkuLists200Response {
	return v.value
}

func (v *NullableGETSkuLists200Response) Set(val *GETSkuLists200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuLists200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuLists200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuLists200Response(val *GETSkuLists200Response) *NullableGETSkuLists200Response {
	return &NullableGETSkuLists200Response{value: val, isSet: true}
}

func (v NullableGETSkuLists200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuLists200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
