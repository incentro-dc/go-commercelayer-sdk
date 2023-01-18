/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETRefundsRefundId200Response struct for GETRefundsRefundId200Response
type GETRefundsRefundId200Response struct {
	Data *GETRefunds200ResponseDataInner `json:"data,omitempty"`
}

// NewGETRefundsRefundId200Response instantiates a new GETRefundsRefundId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETRefundsRefundId200Response() *GETRefundsRefundId200Response {
	this := GETRefundsRefundId200Response{}
	return &this
}

// NewGETRefundsRefundId200ResponseWithDefaults instantiates a new GETRefundsRefundId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETRefundsRefundId200ResponseWithDefaults() *GETRefundsRefundId200Response {
	this := GETRefundsRefundId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETRefundsRefundId200Response) GetData() GETRefunds200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETRefunds200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETRefundsRefundId200Response) GetDataOk() (*GETRefunds200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETRefundsRefundId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETRefunds200ResponseDataInner and assigns it to the Data field.
func (o *GETRefundsRefundId200Response) SetData(v GETRefunds200ResponseDataInner) {
	o.Data = &v
}

func (o GETRefundsRefundId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETRefundsRefundId200Response struct {
	value *GETRefundsRefundId200Response
	isSet bool
}

func (v NullableGETRefundsRefundId200Response) Get() *GETRefundsRefundId200Response {
	return v.value
}

func (v *NullableGETRefundsRefundId200Response) Set(val *GETRefundsRefundId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETRefundsRefundId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETRefundsRefundId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETRefundsRefundId200Response(val *GETRefundsRefundId200Response) *NullableGETRefundsRefundId200Response {
	return &NullableGETRefundsRefundId200Response{value: val, isSet: true}
}

func (v NullableGETRefundsRefundId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETRefundsRefundId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
