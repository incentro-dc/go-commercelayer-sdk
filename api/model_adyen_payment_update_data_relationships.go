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

// AdyenPaymentUpdateDataRelationships struct for AdyenPaymentUpdateDataRelationships
type AdyenPaymentUpdateDataRelationships struct {
	Order *AdyenPaymentDataRelationshipsOrder `json:"order,omitempty"`
}

// NewAdyenPaymentUpdateDataRelationships instantiates a new AdyenPaymentUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentUpdateDataRelationships() *AdyenPaymentUpdateDataRelationships {
	this := AdyenPaymentUpdateDataRelationships{}
	return &this
}

// NewAdyenPaymentUpdateDataRelationshipsWithDefaults instantiates a new AdyenPaymentUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentUpdateDataRelationshipsWithDefaults() *AdyenPaymentUpdateDataRelationships {
	this := AdyenPaymentUpdateDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *AdyenPaymentUpdateDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

func (o AdyenPaymentUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenPaymentUpdateDataRelationships struct {
	value *AdyenPaymentUpdateDataRelationships
	isSet bool
}

func (v NullableAdyenPaymentUpdateDataRelationships) Get() *AdyenPaymentUpdateDataRelationships {
	return v.value
}

func (v *NullableAdyenPaymentUpdateDataRelationships) Set(val *AdyenPaymentUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentUpdateDataRelationships(val *AdyenPaymentUpdateDataRelationships) *NullableAdyenPaymentUpdateDataRelationships {
	return &NullableAdyenPaymentUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableAdyenPaymentUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
