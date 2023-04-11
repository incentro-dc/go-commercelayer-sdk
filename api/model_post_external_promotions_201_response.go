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

// POSTExternalPromotions201Response struct for POSTExternalPromotions201Response
type POSTExternalPromotions201Response struct {
	Data *POSTExternalPromotions201ResponseData `json:"data,omitempty"`
}

// NewPOSTExternalPromotions201Response instantiates a new POSTExternalPromotions201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPromotions201Response() *POSTExternalPromotions201Response {
	this := POSTExternalPromotions201Response{}
	return &this
}

// NewPOSTExternalPromotions201ResponseWithDefaults instantiates a new POSTExternalPromotions201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPromotions201ResponseWithDefaults() *POSTExternalPromotions201Response {
	this := POSTExternalPromotions201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalPromotions201Response) GetData() POSTExternalPromotions201ResponseData {
	if o == nil || o.Data == nil {
		var ret POSTExternalPromotions201ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPromotions201Response) GetDataOk() (*POSTExternalPromotions201ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalPromotions201Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalPromotions201ResponseData and assigns it to the Data field.
func (o *POSTExternalPromotions201Response) SetData(v POSTExternalPromotions201ResponseData) {
	o.Data = &v
}

func (o POSTExternalPromotions201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTExternalPromotions201Response struct {
	value *POSTExternalPromotions201Response
	isSet bool
}

func (v NullablePOSTExternalPromotions201Response) Get() *POSTExternalPromotions201Response {
	return v.value
}

func (v *NullablePOSTExternalPromotions201Response) Set(val *POSTExternalPromotions201Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPromotions201Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPromotions201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPromotions201Response(val *POSTExternalPromotions201Response) *NullablePOSTExternalPromotions201Response {
	return &NullablePOSTExternalPromotions201Response{value: val, isSet: true}
}

func (v NullablePOSTExternalPromotions201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPromotions201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
