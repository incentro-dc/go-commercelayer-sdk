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

// GETKlarnaPaymentsKlarnaPaymentId200Response struct for GETKlarnaPaymentsKlarnaPaymentId200Response
type GETKlarnaPaymentsKlarnaPaymentId200Response struct {
	Data *GETKlarnaPayments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETKlarnaPaymentsKlarnaPaymentId200Response instantiates a new GETKlarnaPaymentsKlarnaPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETKlarnaPaymentsKlarnaPaymentId200Response() *GETKlarnaPaymentsKlarnaPaymentId200Response {
	this := GETKlarnaPaymentsKlarnaPaymentId200Response{}
	return &this
}

// NewGETKlarnaPaymentsKlarnaPaymentId200ResponseWithDefaults instantiates a new GETKlarnaPaymentsKlarnaPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETKlarnaPaymentsKlarnaPaymentId200ResponseWithDefaults() *GETKlarnaPaymentsKlarnaPaymentId200Response {
	this := GETKlarnaPaymentsKlarnaPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETKlarnaPaymentsKlarnaPaymentId200Response) GetData() GETKlarnaPayments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETKlarnaPayments200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaPaymentsKlarnaPaymentId200Response) GetDataOk() (*GETKlarnaPayments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETKlarnaPaymentsKlarnaPaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETKlarnaPayments200ResponseDataInner and assigns it to the Data field.
func (o *GETKlarnaPaymentsKlarnaPaymentId200Response) SetData(v GETKlarnaPayments200ResponseDataInner) {
	o.Data = &v
}

func (o GETKlarnaPaymentsKlarnaPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETKlarnaPaymentsKlarnaPaymentId200Response struct {
	value *GETKlarnaPaymentsKlarnaPaymentId200Response
	isSet bool
}

func (v NullableGETKlarnaPaymentsKlarnaPaymentId200Response) Get() *GETKlarnaPaymentsKlarnaPaymentId200Response {
	return v.value
}

func (v *NullableGETKlarnaPaymentsKlarnaPaymentId200Response) Set(val *GETKlarnaPaymentsKlarnaPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETKlarnaPaymentsKlarnaPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETKlarnaPaymentsKlarnaPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETKlarnaPaymentsKlarnaPaymentId200Response(val *GETKlarnaPaymentsKlarnaPaymentId200Response) *NullableGETKlarnaPaymentsKlarnaPaymentId200Response {
	return &NullableGETKlarnaPaymentsKlarnaPaymentId200Response{value: val, isSet: true}
}

func (v NullableGETKlarnaPaymentsKlarnaPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETKlarnaPaymentsKlarnaPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
