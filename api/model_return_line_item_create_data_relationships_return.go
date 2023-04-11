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

// ReturnLineItemCreateDataRelationshipsReturn struct for ReturnLineItemCreateDataRelationshipsReturn
type ReturnLineItemCreateDataRelationshipsReturn struct {
	Data CustomerDataRelationshipsReturnsData `json:"data"`
}

// NewReturnLineItemCreateDataRelationshipsReturn instantiates a new ReturnLineItemCreateDataRelationshipsReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemCreateDataRelationshipsReturn(data CustomerDataRelationshipsReturnsData) *ReturnLineItemCreateDataRelationshipsReturn {
	this := ReturnLineItemCreateDataRelationshipsReturn{}
	this.Data = data
	return &this
}

// NewReturnLineItemCreateDataRelationshipsReturnWithDefaults instantiates a new ReturnLineItemCreateDataRelationshipsReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemCreateDataRelationshipsReturnWithDefaults() *ReturnLineItemCreateDataRelationshipsReturn {
	this := ReturnLineItemCreateDataRelationshipsReturn{}
	return &this
}

// GetData returns the Data field value
func (o *ReturnLineItemCreateDataRelationshipsReturn) GetData() CustomerDataRelationshipsReturnsData {
	if o == nil {
		var ret CustomerDataRelationshipsReturnsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreateDataRelationshipsReturn) GetDataOk() (*CustomerDataRelationshipsReturnsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReturnLineItemCreateDataRelationshipsReturn) SetData(v CustomerDataRelationshipsReturnsData) {
	o.Data = v
}

func (o ReturnLineItemCreateDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnLineItemCreateDataRelationshipsReturn struct {
	value *ReturnLineItemCreateDataRelationshipsReturn
	isSet bool
}

func (v NullableReturnLineItemCreateDataRelationshipsReturn) Get() *ReturnLineItemCreateDataRelationshipsReturn {
	return v.value
}

func (v *NullableReturnLineItemCreateDataRelationshipsReturn) Set(val *ReturnLineItemCreateDataRelationshipsReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemCreateDataRelationshipsReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemCreateDataRelationshipsReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemCreateDataRelationshipsReturn(val *ReturnLineItemCreateDataRelationshipsReturn) *NullableReturnLineItemCreateDataRelationshipsReturn {
	return &NullableReturnLineItemCreateDataRelationshipsReturn{value: val, isSet: true}
}

func (v NullableReturnLineItemCreateDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemCreateDataRelationshipsReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
