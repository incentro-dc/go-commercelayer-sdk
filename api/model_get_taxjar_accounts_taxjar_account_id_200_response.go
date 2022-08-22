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

// GETTaxjarAccountsTaxjarAccountId200Response struct for GETTaxjarAccountsTaxjarAccountId200Response
type GETTaxjarAccountsTaxjarAccountId200Response struct {
	Data *GETTaxjarAccounts200ResponseDataInner `json:"data,omitempty"`
}

// NewGETTaxjarAccountsTaxjarAccountId200Response instantiates a new GETTaxjarAccountsTaxjarAccountId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETTaxjarAccountsTaxjarAccountId200Response() *GETTaxjarAccountsTaxjarAccountId200Response {
	this := GETTaxjarAccountsTaxjarAccountId200Response{}
	return &this
}

// NewGETTaxjarAccountsTaxjarAccountId200ResponseWithDefaults instantiates a new GETTaxjarAccountsTaxjarAccountId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETTaxjarAccountsTaxjarAccountId200ResponseWithDefaults() *GETTaxjarAccountsTaxjarAccountId200Response {
	this := GETTaxjarAccountsTaxjarAccountId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETTaxjarAccountsTaxjarAccountId200Response) GetData() GETTaxjarAccounts200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETTaxjarAccounts200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETTaxjarAccountsTaxjarAccountId200Response) GetDataOk() (*GETTaxjarAccounts200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETTaxjarAccountsTaxjarAccountId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETTaxjarAccounts200ResponseDataInner and assigns it to the Data field.
func (o *GETTaxjarAccountsTaxjarAccountId200Response) SetData(v GETTaxjarAccounts200ResponseDataInner) {
	o.Data = &v
}

func (o GETTaxjarAccountsTaxjarAccountId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETTaxjarAccountsTaxjarAccountId200Response struct {
	value *GETTaxjarAccountsTaxjarAccountId200Response
	isSet bool
}

func (v NullableGETTaxjarAccountsTaxjarAccountId200Response) Get() *GETTaxjarAccountsTaxjarAccountId200Response {
	return v.value
}

func (v *NullableGETTaxjarAccountsTaxjarAccountId200Response) Set(val *GETTaxjarAccountsTaxjarAccountId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETTaxjarAccountsTaxjarAccountId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETTaxjarAccountsTaxjarAccountId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETTaxjarAccountsTaxjarAccountId200Response(val *GETTaxjarAccountsTaxjarAccountId200Response) *NullableGETTaxjarAccountsTaxjarAccountId200Response {
	return &NullableGETTaxjarAccountsTaxjarAccountId200Response{value: val, isSet: true}
}

func (v NullableGETTaxjarAccountsTaxjarAccountId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETTaxjarAccountsTaxjarAccountId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
