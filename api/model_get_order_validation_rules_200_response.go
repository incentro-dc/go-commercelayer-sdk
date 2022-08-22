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

// GETOrderValidationRules200Response struct for GETOrderValidationRules200Response
type GETOrderValidationRules200Response struct {
	Data []GETOrderValidationRules200ResponseDataInner `json:"data,omitempty"`
}

// NewGETOrderValidationRules200Response instantiates a new GETOrderValidationRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderValidationRules200Response() *GETOrderValidationRules200Response {
	this := GETOrderValidationRules200Response{}
	return &this
}

// NewGETOrderValidationRules200ResponseWithDefaults instantiates a new GETOrderValidationRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderValidationRules200ResponseWithDefaults() *GETOrderValidationRules200Response {
	this := GETOrderValidationRules200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrderValidationRules200Response) GetData() []GETOrderValidationRules200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrderValidationRules200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderValidationRules200Response) GetDataOk() ([]GETOrderValidationRules200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrderValidationRules200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrderValidationRules200ResponseDataInner and assigns it to the Data field.
func (o *GETOrderValidationRules200Response) SetData(v []GETOrderValidationRules200ResponseDataInner) {
	o.Data = v
}

func (o GETOrderValidationRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderValidationRules200Response struct {
	value *GETOrderValidationRules200Response
	isSet bool
}

func (v NullableGETOrderValidationRules200Response) Get() *GETOrderValidationRules200Response {
	return v.value
}

func (v *NullableGETOrderValidationRules200Response) Set(val *GETOrderValidationRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderValidationRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderValidationRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderValidationRules200Response(val *GETOrderValidationRules200Response) *NullableGETOrderValidationRules200Response {
	return &NullableGETOrderValidationRules200Response{value: val, isSet: true}
}

func (v NullableGETOrderValidationRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderValidationRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
