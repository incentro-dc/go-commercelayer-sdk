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

// Sku struct for Sku
type Sku struct {
	Data *SkuData `json:"data,omitempty"`
}

// NewSku instantiates a new Sku object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSku() *Sku {
	this := Sku{}
	return &this
}

// NewSkuWithDefaults instantiates a new Sku object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuWithDefaults() *Sku {
	this := Sku{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Sku) GetData() SkuData {
	if o == nil || o.Data == nil {
		var ret SkuData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sku) GetDataOk() (*SkuData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Sku) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SkuData and assigns it to the Data field.
func (o *Sku) SetData(v SkuData) {
	o.Data = &v
}

func (o Sku) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSku struct {
	value *Sku
	isSet bool
}

func (v NullableSku) Get() *Sku {
	return v.value
}

func (v *NullableSku) Set(val *Sku) {
	v.value = val
	v.isSet = true
}

func (v NullableSku) IsSet() bool {
	return v.isSet
}

func (v *NullableSku) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSku(val *Sku) *NullableSku {
	return &NullableSku{value: val, isSet: true}
}

func (v NullableSku) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSku) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
