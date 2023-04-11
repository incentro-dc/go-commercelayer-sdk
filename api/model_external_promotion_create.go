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

// ExternalPromotionCreate struct for ExternalPromotionCreate
type ExternalPromotionCreate struct {
	Data ExternalPromotionCreateData `json:"data"`
}

// NewExternalPromotionCreate instantiates a new ExternalPromotionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionCreate(data ExternalPromotionCreateData) *ExternalPromotionCreate {
	this := ExternalPromotionCreate{}
	this.Data = data
	return &this
}

// NewExternalPromotionCreateWithDefaults instantiates a new ExternalPromotionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionCreateWithDefaults() *ExternalPromotionCreate {
	this := ExternalPromotionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalPromotionCreate) GetData() ExternalPromotionCreateData {
	if o == nil {
		var ret ExternalPromotionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreate) GetDataOk() (*ExternalPromotionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalPromotionCreate) SetData(v ExternalPromotionCreateData) {
	o.Data = v
}

func (o ExternalPromotionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionCreate struct {
	value *ExternalPromotionCreate
	isSet bool
}

func (v NullableExternalPromotionCreate) Get() *ExternalPromotionCreate {
	return v.value
}

func (v *NullableExternalPromotionCreate) Set(val *ExternalPromotionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionCreate(val *ExternalPromotionCreate) *NullableExternalPromotionCreate {
	return &NullableExternalPromotionCreate{value: val, isSet: true}
}

func (v NullableExternalPromotionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
