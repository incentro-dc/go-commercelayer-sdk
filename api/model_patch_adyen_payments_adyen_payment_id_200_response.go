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

// PATCHAdyenPaymentsAdyenPaymentId200Response struct for PATCHAdyenPaymentsAdyenPaymentId200Response
type PATCHAdyenPaymentsAdyenPaymentId200Response struct {
	Data *PATCHAdyenPaymentsAdyenPaymentId200ResponseData `json:"data,omitempty"`
}

// NewPATCHAdyenPaymentsAdyenPaymentId200Response instantiates a new PATCHAdyenPaymentsAdyenPaymentId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAdyenPaymentsAdyenPaymentId200Response() *PATCHAdyenPaymentsAdyenPaymentId200Response {
	this := PATCHAdyenPaymentsAdyenPaymentId200Response{}
	return &this
}

// NewPATCHAdyenPaymentsAdyenPaymentId200ResponseWithDefaults instantiates a new PATCHAdyenPaymentsAdyenPaymentId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAdyenPaymentsAdyenPaymentId200ResponseWithDefaults() *PATCHAdyenPaymentsAdyenPaymentId200Response {
	this := PATCHAdyenPaymentsAdyenPaymentId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHAdyenPaymentsAdyenPaymentId200Response) GetData() PATCHAdyenPaymentsAdyenPaymentId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHAdyenPaymentsAdyenPaymentId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenPaymentsAdyenPaymentId200Response) GetDataOk() (*PATCHAdyenPaymentsAdyenPaymentId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHAdyenPaymentsAdyenPaymentId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHAdyenPaymentsAdyenPaymentId200ResponseData and assigns it to the Data field.
func (o *PATCHAdyenPaymentsAdyenPaymentId200Response) SetData(v PATCHAdyenPaymentsAdyenPaymentId200ResponseData) {
	o.Data = &v
}

func (o PATCHAdyenPaymentsAdyenPaymentId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHAdyenPaymentsAdyenPaymentId200Response struct {
	value *PATCHAdyenPaymentsAdyenPaymentId200Response
	isSet bool
}

func (v NullablePATCHAdyenPaymentsAdyenPaymentId200Response) Get() *PATCHAdyenPaymentsAdyenPaymentId200Response {
	return v.value
}

func (v *NullablePATCHAdyenPaymentsAdyenPaymentId200Response) Set(val *PATCHAdyenPaymentsAdyenPaymentId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAdyenPaymentsAdyenPaymentId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAdyenPaymentsAdyenPaymentId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAdyenPaymentsAdyenPaymentId200Response(val *PATCHAdyenPaymentsAdyenPaymentId200Response) *NullablePATCHAdyenPaymentsAdyenPaymentId200Response {
	return &NullablePATCHAdyenPaymentsAdyenPaymentId200Response{value: val, isSet: true}
}

func (v NullablePATCHAdyenPaymentsAdyenPaymentId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAdyenPaymentsAdyenPaymentId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}