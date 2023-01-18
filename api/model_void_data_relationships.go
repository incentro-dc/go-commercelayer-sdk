/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// VoidDataRelationships struct for VoidDataRelationships
type VoidDataRelationships struct {
	Order                  *AdyenPaymentDataRelationshipsOrder             `json:"order,omitempty"`
	ReferenceAuthorization *CaptureDataRelationshipsReferenceAuthorization `json:"reference_authorization,omitempty"`
}

// NewVoidDataRelationships instantiates a new VoidDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidDataRelationships() *VoidDataRelationships {
	this := VoidDataRelationships{}
	return &this
}

// NewVoidDataRelationshipsWithDefaults instantiates a new VoidDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidDataRelationshipsWithDefaults() *VoidDataRelationships {
	this := VoidDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *VoidDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetReferenceAuthorization returns the ReferenceAuthorization field value if set, zero value otherwise.
func (o *VoidDataRelationships) GetReferenceAuthorization() CaptureDataRelationshipsReferenceAuthorization {
	if o == nil || o.ReferenceAuthorization == nil {
		var ret CaptureDataRelationshipsReferenceAuthorization
		return ret
	}
	return *o.ReferenceAuthorization
}

// GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidDataRelationships) GetReferenceAuthorizationOk() (*CaptureDataRelationshipsReferenceAuthorization, bool) {
	if o == nil || o.ReferenceAuthorization == nil {
		return nil, false
	}
	return o.ReferenceAuthorization, true
}

// HasReferenceAuthorization returns a boolean if a field has been set.
func (o *VoidDataRelationships) HasReferenceAuthorization() bool {
	if o != nil && o.ReferenceAuthorization != nil {
		return true
	}

	return false
}

// SetReferenceAuthorization gets a reference to the given CaptureDataRelationshipsReferenceAuthorization and assigns it to the ReferenceAuthorization field.
func (o *VoidDataRelationships) SetReferenceAuthorization(v CaptureDataRelationshipsReferenceAuthorization) {
	o.ReferenceAuthorization = &v
}

func (o VoidDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ReferenceAuthorization != nil {
		toSerialize["reference_authorization"] = o.ReferenceAuthorization
	}
	return json.Marshal(toSerialize)
}

type NullableVoidDataRelationships struct {
	value *VoidDataRelationships
	isSet bool
}

func (v NullableVoidDataRelationships) Get() *VoidDataRelationships {
	return v.value
}

func (v *NullableVoidDataRelationships) Set(val *VoidDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidDataRelationships(val *VoidDataRelationships) *NullableVoidDataRelationships {
	return &NullableVoidDataRelationships{value: val, isSet: true}
}

func (v NullableVoidDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
