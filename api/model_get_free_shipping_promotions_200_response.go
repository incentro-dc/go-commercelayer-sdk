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

// GETFreeShippingPromotions200Response struct for GETFreeShippingPromotions200Response
type GETFreeShippingPromotions200Response struct {
	Data []GETFreeShippingPromotions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETFreeShippingPromotions200Response instantiates a new GETFreeShippingPromotions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFreeShippingPromotions200Response() *GETFreeShippingPromotions200Response {
	this := GETFreeShippingPromotions200Response{}
	return &this
}

// NewGETFreeShippingPromotions200ResponseWithDefaults instantiates a new GETFreeShippingPromotions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFreeShippingPromotions200ResponseWithDefaults() *GETFreeShippingPromotions200Response {
	this := GETFreeShippingPromotions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETFreeShippingPromotions200Response) GetData() []GETFreeShippingPromotions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETFreeShippingPromotions200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeShippingPromotions200Response) GetDataOk() ([]GETFreeShippingPromotions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETFreeShippingPromotions200ResponseDataInner and assigns it to the Data field.
func (o *GETFreeShippingPromotions200Response) SetData(v []GETFreeShippingPromotions200ResponseDataInner) {
	o.Data = v
}

func (o GETFreeShippingPromotions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETFreeShippingPromotions200Response struct {
	value *GETFreeShippingPromotions200Response
	isSet bool
}

func (v NullableGETFreeShippingPromotions200Response) Get() *GETFreeShippingPromotions200Response {
	return v.value
}

func (v *NullableGETFreeShippingPromotions200Response) Set(val *GETFreeShippingPromotions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFreeShippingPromotions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFreeShippingPromotions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFreeShippingPromotions200Response(val *GETFreeShippingPromotions200Response) *NullableGETFreeShippingPromotions200Response {
	return &NullableGETFreeShippingPromotions200Response{value: val, isSet: true}
}

func (v NullableGETFreeShippingPromotions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFreeShippingPromotions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
