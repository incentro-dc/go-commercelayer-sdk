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

// GETCouponRecipients200Response struct for GETCouponRecipients200Response
type GETCouponRecipients200Response struct {
	Data []GETCouponRecipients200ResponseDataInner `json:"data,omitempty"`
}

// NewGETCouponRecipients200Response instantiates a new GETCouponRecipients200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponRecipients200Response() *GETCouponRecipients200Response {
	this := GETCouponRecipients200Response{}
	return &this
}

// NewGETCouponRecipients200ResponseWithDefaults instantiates a new GETCouponRecipients200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponRecipients200ResponseWithDefaults() *GETCouponRecipients200Response {
	this := GETCouponRecipients200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCouponRecipients200Response) GetData() []GETCouponRecipients200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETCouponRecipients200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponRecipients200Response) GetDataOk() ([]GETCouponRecipients200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCouponRecipients200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETCouponRecipients200ResponseDataInner and assigns it to the Data field.
func (o *GETCouponRecipients200Response) SetData(v []GETCouponRecipients200ResponseDataInner) {
	o.Data = v
}

func (o GETCouponRecipients200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCouponRecipients200Response struct {
	value *GETCouponRecipients200Response
	isSet bool
}

func (v NullableGETCouponRecipients200Response) Get() *GETCouponRecipients200Response {
	return v.value
}

func (v *NullableGETCouponRecipients200Response) Set(val *GETCouponRecipients200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponRecipients200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponRecipients200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponRecipients200Response(val *GETCouponRecipients200Response) *NullableGETCouponRecipients200Response {
	return &NullableGETCouponRecipients200Response{value: val, isSet: true}
}

func (v NullableGETCouponRecipients200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponRecipients200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
