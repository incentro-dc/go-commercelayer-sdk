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

// OrderCopyCreateDataRelationships struct for OrderCopyCreateDataRelationships
type OrderCopyCreateDataRelationships struct {
	SourceOrder AdyenPaymentDataRelationshipsOrder `json:"source_order"`
}

// NewOrderCopyCreateDataRelationships instantiates a new OrderCopyCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyCreateDataRelationships(sourceOrder AdyenPaymentDataRelationshipsOrder) *OrderCopyCreateDataRelationships {
	this := OrderCopyCreateDataRelationships{}
	this.SourceOrder = sourceOrder
	return &this
}

// NewOrderCopyCreateDataRelationshipsWithDefaults instantiates a new OrderCopyCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyCreateDataRelationshipsWithDefaults() *OrderCopyCreateDataRelationships {
	this := OrderCopyCreateDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value
func (o *OrderCopyCreateDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}

	return o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value
// and a boolean to check if the value has been set.
func (o *OrderCopyCreateDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceOrder, true
}

// SetSourceOrder sets field value
func (o *OrderCopyCreateDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.SourceOrder = v
}

func (o OrderCopyCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source_order"] = o.SourceOrder
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCopyCreateDataRelationships struct {
	value *OrderCopyCreateDataRelationships
	isSet bool
}

func (v NullableOrderCopyCreateDataRelationships) Get() *OrderCopyCreateDataRelationships {
	return v.value
}

func (v *NullableOrderCopyCreateDataRelationships) Set(val *OrderCopyCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyCreateDataRelationships(val *OrderCopyCreateDataRelationships) *NullableOrderCopyCreateDataRelationships {
	return &NullableOrderCopyCreateDataRelationships{value: val, isSet: true}
}

func (v NullableOrderCopyCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
