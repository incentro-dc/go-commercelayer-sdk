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

// GETShippingWeightTiers200Response struct for GETShippingWeightTiers200Response
type GETShippingWeightTiers200Response struct {
	Data []GETShippingWeightTiers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShippingWeightTiers200Response instantiates a new GETShippingWeightTiers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingWeightTiers200Response() *GETShippingWeightTiers200Response {
	this := GETShippingWeightTiers200Response{}
	return &this
}

// NewGETShippingWeightTiers200ResponseWithDefaults instantiates a new GETShippingWeightTiers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingWeightTiers200ResponseWithDefaults() *GETShippingWeightTiers200Response {
	this := GETShippingWeightTiers200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShippingWeightTiers200Response) GetData() []GETShippingWeightTiers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETShippingWeightTiers200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingWeightTiers200Response) GetDataOk() ([]GETShippingWeightTiers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShippingWeightTiers200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETShippingWeightTiers200ResponseDataInner and assigns it to the Data field.
func (o *GETShippingWeightTiers200Response) SetData(v []GETShippingWeightTiers200ResponseDataInner) {
	o.Data = v
}

func (o GETShippingWeightTiers200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingWeightTiers200Response struct {
	value *GETShippingWeightTiers200Response
	isSet bool
}

func (v NullableGETShippingWeightTiers200Response) Get() *GETShippingWeightTiers200Response {
	return v.value
}

func (v *NullableGETShippingWeightTiers200Response) Set(val *GETShippingWeightTiers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingWeightTiers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingWeightTiers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingWeightTiers200Response(val *GETShippingWeightTiers200Response) *NullableGETShippingWeightTiers200Response {
	return &NullableGETShippingWeightTiers200Response{value: val, isSet: true}
}

func (v NullableGETShippingWeightTiers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingWeightTiers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}