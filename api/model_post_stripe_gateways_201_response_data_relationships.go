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

// checks if the POSTStripeGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStripeGateways201ResponseDataRelationships{}

// POSTStripeGateways201ResponseDataRelationships struct for POSTStripeGateways201ResponseDataRelationships
type POSTStripeGateways201ResponseDataRelationships struct {
	PaymentMethods *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	StripePayments *POSTStripeGateways201ResponseDataRelationshipsStripePayments `json:"stripe_payments,omitempty"`
}

// NewPOSTStripeGateways201ResponseDataRelationships instantiates a new POSTStripeGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStripeGateways201ResponseDataRelationships() *POSTStripeGateways201ResponseDataRelationships {
	this := POSTStripeGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTStripeGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTStripeGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStripeGateways201ResponseDataRelationshipsWithDefaults() *POSTStripeGateways201ResponseDataRelationships {
	this := POSTStripeGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTStripeGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTStripeGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTStripeGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetStripePayments returns the StripePayments field value if set, zero value otherwise.
func (o *POSTStripeGateways201ResponseDataRelationships) GetStripePayments() POSTStripeGateways201ResponseDataRelationshipsStripePayments {
	if o == nil || IsNil(o.StripePayments) {
		var ret POSTStripeGateways201ResponseDataRelationshipsStripePayments
		return ret
	}
	return *o.StripePayments
}

// GetStripePaymentsOk returns a tuple with the StripePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStripeGateways201ResponseDataRelationships) GetStripePaymentsOk() (*POSTStripeGateways201ResponseDataRelationshipsStripePayments, bool) {
	if o == nil || IsNil(o.StripePayments) {
		return nil, false
	}
	return o.StripePayments, true
}

// HasStripePayments returns a boolean if a field has been set.
func (o *POSTStripeGateways201ResponseDataRelationships) HasStripePayments() bool {
	if o != nil && !IsNil(o.StripePayments) {
		return true
	}

	return false
}

// SetStripePayments gets a reference to the given POSTStripeGateways201ResponseDataRelationshipsStripePayments and assigns it to the StripePayments field.
func (o *POSTStripeGateways201ResponseDataRelationships) SetStripePayments(v POSTStripeGateways201ResponseDataRelationshipsStripePayments) {
	o.StripePayments = &v
}

func (o POSTStripeGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStripeGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.StripePayments) {
		toSerialize["stripe_payments"] = o.StripePayments
	}
	return toSerialize, nil
}

type NullablePOSTStripeGateways201ResponseDataRelationships struct {
	value *POSTStripeGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTStripeGateways201ResponseDataRelationships) Get() *POSTStripeGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTStripeGateways201ResponseDataRelationships) Set(val *POSTStripeGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStripeGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStripeGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStripeGateways201ResponseDataRelationships(val *POSTStripeGateways201ResponseDataRelationships) *NullablePOSTStripeGateways201ResponseDataRelationships {
	return &NullablePOSTStripeGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTStripeGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStripeGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}