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

// GETExternalTaxCalculators200Response struct for GETExternalTaxCalculators200Response
type GETExternalTaxCalculators200Response struct {
	Data []GETExternalTaxCalculators200ResponseDataInner `json:"data,omitempty"`
}

// NewGETExternalTaxCalculators200Response instantiates a new GETExternalTaxCalculators200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalTaxCalculators200Response() *GETExternalTaxCalculators200Response {
	this := GETExternalTaxCalculators200Response{}
	return &this
}

// NewGETExternalTaxCalculators200ResponseWithDefaults instantiates a new GETExternalTaxCalculators200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalTaxCalculators200ResponseWithDefaults() *GETExternalTaxCalculators200Response {
	this := GETExternalTaxCalculators200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalTaxCalculators200Response) GetData() []GETExternalTaxCalculators200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETExternalTaxCalculators200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalTaxCalculators200Response) GetDataOk() ([]GETExternalTaxCalculators200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalTaxCalculators200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETExternalTaxCalculators200ResponseDataInner and assigns it to the Data field.
func (o *GETExternalTaxCalculators200Response) SetData(v []GETExternalTaxCalculators200ResponseDataInner) {
	o.Data = v
}

func (o GETExternalTaxCalculators200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalTaxCalculators200Response struct {
	value *GETExternalTaxCalculators200Response
	isSet bool
}

func (v NullableGETExternalTaxCalculators200Response) Get() *GETExternalTaxCalculators200Response {
	return v.value
}

func (v *NullableGETExternalTaxCalculators200Response) Set(val *GETExternalTaxCalculators200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalTaxCalculators200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalTaxCalculators200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalTaxCalculators200Response(val *GETExternalTaxCalculators200Response) *NullableGETExternalTaxCalculators200Response {
	return &NullableGETExternalTaxCalculators200Response{value: val, isSet: true}
}

func (v NullableGETExternalTaxCalculators200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalTaxCalculators200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
