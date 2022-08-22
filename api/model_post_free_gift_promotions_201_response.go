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

// POSTFreeGiftPromotions201Response struct for POSTFreeGiftPromotions201Response
type POSTFreeGiftPromotions201Response struct {
	Data *POSTFreeGiftPromotions201ResponseData `json:"data,omitempty"`
}

// NewPOSTFreeGiftPromotions201Response instantiates a new POSTFreeGiftPromotions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFreeGiftPromotions201Response() *POSTFreeGiftPromotions201Response {
	this := POSTFreeGiftPromotions201Response{}
	return &this
}

// NewPOSTFreeGiftPromotions201ResponseWithDefaults instantiates a new POSTFreeGiftPromotions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFreeGiftPromotions201ResponseWithDefaults() *POSTFreeGiftPromotions201Response {
	this := POSTFreeGiftPromotions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTFreeGiftPromotions201Response) GetData() POSTFreeGiftPromotions201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTFreeGiftPromotions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFreeGiftPromotions201Response) GetDataOk() (*POSTFreeGiftPromotions201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTFreeGiftPromotions201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTFreeGiftPromotions201ResponseData and assigns it to the Data field.
func (o *POSTFreeGiftPromotions201Response) SetData(v POSTFreeGiftPromotions201ResponseData) {
	o.Data = &v
}

func (o POSTFreeGiftPromotions201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTFreeGiftPromotions201Response struct {
	value *POSTFreeGiftPromotions201Response
	isSet bool
}

func (v NullablePOSTFreeGiftPromotions201Response) Get() *POSTFreeGiftPromotions201Response {
	return v.value
}

func (v *NullablePOSTFreeGiftPromotions201Response) Set(val *POSTFreeGiftPromotions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFreeGiftPromotions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFreeGiftPromotions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFreeGiftPromotions201Response(val *POSTFreeGiftPromotions201Response) *NullablePOSTFreeGiftPromotions201Response {
	return &NullablePOSTFreeGiftPromotions201Response{value: val, isSet: true}
}

func (v NullablePOSTFreeGiftPromotions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFreeGiftPromotions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
