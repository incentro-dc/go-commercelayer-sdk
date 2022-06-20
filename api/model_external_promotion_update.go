/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPromotionUpdate struct for ExternalPromotionUpdate
type ExternalPromotionUpdate struct {
	Data ExternalPromotionUpdateData `json:"data"`
}

// NewExternalPromotionUpdate instantiates a new ExternalPromotionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionUpdate(data ExternalPromotionUpdateData) *ExternalPromotionUpdate {
	this := ExternalPromotionUpdate{}
	this.Data = data
	return &this
}

// NewExternalPromotionUpdateWithDefaults instantiates a new ExternalPromotionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionUpdateWithDefaults() *ExternalPromotionUpdate {
	this := ExternalPromotionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalPromotionUpdate) GetData() ExternalPromotionUpdateData {
	if o == nil {
		var ret ExternalPromotionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalPromotionUpdate) GetDataOk() (*ExternalPromotionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalPromotionUpdate) SetData(v ExternalPromotionUpdateData) {
	o.Data = v
}

func (o ExternalPromotionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionUpdate struct {
	value *ExternalPromotionUpdate
	isSet bool
}

func (v NullableExternalPromotionUpdate) Get() *ExternalPromotionUpdate {
	return v.value
}

func (v *NullableExternalPromotionUpdate) Set(val *ExternalPromotionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionUpdate(val *ExternalPromotionUpdate) *NullableExternalPromotionUpdate {
	return &NullableExternalPromotionUpdate{value: val, isSet: true}
}

func (v NullableExternalPromotionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}