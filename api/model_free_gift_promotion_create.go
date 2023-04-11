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

// FreeGiftPromotionCreate struct for FreeGiftPromotionCreate
type FreeGiftPromotionCreate struct {
	Data FreeGiftPromotionCreateData `json:"data"`
}

// NewFreeGiftPromotionCreate instantiates a new FreeGiftPromotionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeGiftPromotionCreate(data FreeGiftPromotionCreateData) *FreeGiftPromotionCreate {
	this := FreeGiftPromotionCreate{}
	this.Data = data
	return &this
}

// NewFreeGiftPromotionCreateWithDefaults instantiates a new FreeGiftPromotionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeGiftPromotionCreateWithDefaults() *FreeGiftPromotionCreate {
	this := FreeGiftPromotionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *FreeGiftPromotionCreate) GetData() FreeGiftPromotionCreateData {
	if o == nil {
		var ret FreeGiftPromotionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionCreate) GetDataOk() (*FreeGiftPromotionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FreeGiftPromotionCreate) SetData(v FreeGiftPromotionCreateData) {
	o.Data = v
}

func (o FreeGiftPromotionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFreeGiftPromotionCreate struct {
	value *FreeGiftPromotionCreate
	isSet bool
}

func (v NullableFreeGiftPromotionCreate) Get() *FreeGiftPromotionCreate {
	return v.value
}

func (v *NullableFreeGiftPromotionCreate) Set(val *FreeGiftPromotionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeGiftPromotionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeGiftPromotionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeGiftPromotionCreate(val *FreeGiftPromotionCreate) *NullableFreeGiftPromotionCreate {
	return &NullableFreeGiftPromotionCreate{value: val, isSet: true}
}

func (v NullableFreeGiftPromotionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeGiftPromotionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
