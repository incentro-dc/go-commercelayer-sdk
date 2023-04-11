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

// GETOrderAmountPromotionRules200Response struct for GETOrderAmountPromotionRules200Response
type GETOrderAmountPromotionRules200Response struct {
	Data []GETOrderAmountPromotionRules200ResponseDataInner `json:"data,omitempty"`
}

// NewGETOrderAmountPromotionRules200Response instantiates a new GETOrderAmountPromotionRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderAmountPromotionRules200Response() *GETOrderAmountPromotionRules200Response {
	this := GETOrderAmountPromotionRules200Response{}
	return &this
}

// NewGETOrderAmountPromotionRules200ResponseWithDefaults instantiates a new GETOrderAmountPromotionRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderAmountPromotionRules200ResponseWithDefaults() *GETOrderAmountPromotionRules200Response {
	this := GETOrderAmountPromotionRules200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrderAmountPromotionRules200Response) GetData() []GETOrderAmountPromotionRules200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETOrderAmountPromotionRules200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderAmountPromotionRules200Response) GetDataOk() ([]GETOrderAmountPromotionRules200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrderAmountPromotionRules200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETOrderAmountPromotionRules200ResponseDataInner and assigns it to the Data field.
func (o *GETOrderAmountPromotionRules200Response) SetData(v []GETOrderAmountPromotionRules200ResponseDataInner) {
	o.Data = v
}

func (o GETOrderAmountPromotionRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderAmountPromotionRules200Response struct {
	value *GETOrderAmountPromotionRules200Response
	isSet bool
}

func (v NullableGETOrderAmountPromotionRules200Response) Get() *GETOrderAmountPromotionRules200Response {
	return v.value
}

func (v *NullableGETOrderAmountPromotionRules200Response) Set(val *GETOrderAmountPromotionRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderAmountPromotionRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderAmountPromotionRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderAmountPromotionRules200Response(val *GETOrderAmountPromotionRules200Response) *NullableGETOrderAmountPromotionRules200Response {
	return &NullableGETOrderAmountPromotionRules200Response{value: val, isSet: true}
}

func (v NullableGETOrderAmountPromotionRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderAmountPromotionRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
