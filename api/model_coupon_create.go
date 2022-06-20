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

// CouponCreate struct for CouponCreate
type CouponCreate struct {
	Data CouponCreateData `json:"data"`
}

// NewCouponCreate instantiates a new CouponCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCreate(data CouponCreateData) *CouponCreate {
	this := CouponCreate{}
	this.Data = data
	return &this
}

// NewCouponCreateWithDefaults instantiates a new CouponCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCreateWithDefaults() *CouponCreate {
	this := CouponCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCreate) GetData() CouponCreateData {
	if o == nil {
		var ret CouponCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCreate) GetDataOk() (*CouponCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCreate) SetData(v CouponCreateData) {
	o.Data = v
}

func (o CouponCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCreate struct {
	value *CouponCreate
	isSet bool
}

func (v NullableCouponCreate) Get() *CouponCreate {
	return v.value
}

func (v *NullableCouponCreate) Set(val *CouponCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCreate(val *CouponCreate) *NullableCouponCreate {
	return &NullableCouponCreate{value: val, isSet: true}
}

func (v NullableCouponCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
