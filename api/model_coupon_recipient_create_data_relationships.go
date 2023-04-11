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

// CouponRecipientCreateDataRelationships struct for CouponRecipientCreateDataRelationships
type CouponRecipientCreateDataRelationships struct {
	Customer *CouponRecipientCreateDataRelationshipsCustomer `json:"customer,omitempty"`
}

// NewCouponRecipientCreateDataRelationships instantiates a new CouponRecipientCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponRecipientCreateDataRelationships() *CouponRecipientCreateDataRelationships {
	this := CouponRecipientCreateDataRelationships{}
	return &this
}

// NewCouponRecipientCreateDataRelationshipsWithDefaults instantiates a new CouponRecipientCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponRecipientCreateDataRelationshipsWithDefaults() *CouponRecipientCreateDataRelationships {
	this := CouponRecipientCreateDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CouponRecipientCreateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientCreateDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponRecipientCreateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CouponRecipientCreateDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientCreateDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CouponRecipientCreateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer) {
	o.Customer = &v
}

func (o CouponRecipientCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	return json.Marshal(toSerialize)
}

type NullableCouponRecipientCreateDataRelationships struct {
	value *CouponRecipientCreateDataRelationships
	isSet bool
}

func (v NullableCouponRecipientCreateDataRelationships) Get() *CouponRecipientCreateDataRelationships {
	return v.value
}

func (v *NullableCouponRecipientCreateDataRelationships) Set(val *CouponRecipientCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponRecipientCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponRecipientCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponRecipientCreateDataRelationships(val *CouponRecipientCreateDataRelationships) *NullableCouponRecipientCreateDataRelationships {
	return &NullableCouponRecipientCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCouponRecipientCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponRecipientCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
