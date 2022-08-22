/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETPaypalGateways200ResponseDataInnerRelationships struct for GETPaypalGateways200ResponseDataInnerRelationships
type GETPaypalGateways200ResponseDataInnerRelationships struct {
	PaymentMethods *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	PaypalPayments *GETPaypalGateways200ResponseDataInnerRelationshipsPaypalPayments `json:"paypal_payments,omitempty"`
}

// NewGETPaypalGateways200ResponseDataInnerRelationships instantiates a new GETPaypalGateways200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaypalGateways200ResponseDataInnerRelationships() *GETPaypalGateways200ResponseDataInnerRelationships {
	this := GETPaypalGateways200ResponseDataInnerRelationships{}
	return &this
}

// NewGETPaypalGateways200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETPaypalGateways200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaypalGateways200ResponseDataInnerRelationshipsWithDefaults() *GETPaypalGateways200ResponseDataInnerRelationships {
	this := GETPaypalGateways200ResponseDataInnerRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) GetPaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) GetPaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) SetPaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetPaypalPayments returns the PaypalPayments field value if set, zero value otherwise.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) GetPaypalPayments() GETPaypalGateways200ResponseDataInnerRelationshipsPaypalPayments {
	if o == nil || o.PaypalPayments == nil {
		var ret GETPaypalGateways200ResponseDataInnerRelationshipsPaypalPayments
		return ret
	}
	return *o.PaypalPayments
}

// GetPaypalPaymentsOk returns a tuple with the PaypalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) GetPaypalPaymentsOk() (*GETPaypalGateways200ResponseDataInnerRelationshipsPaypalPayments, bool) {
	if o == nil || o.PaypalPayments == nil {
		return nil, false
	}
	return o.PaypalPayments, true
}

// HasPaypalPayments returns a boolean if a field has been set.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) HasPaypalPayments() bool {
	if o != nil && o.PaypalPayments != nil {
		return true
	}

	return false
}

// SetPaypalPayments gets a reference to the given GETPaypalGateways200ResponseDataInnerRelationshipsPaypalPayments and assigns it to the PaypalPayments field.
func (o *GETPaypalGateways200ResponseDataInnerRelationships) SetPaypalPayments(v GETPaypalGateways200ResponseDataInnerRelationshipsPaypalPayments) {
	o.PaypalPayments = &v
}

func (o GETPaypalGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.PaypalPayments != nil {
		toSerialize["paypal_payments"] = o.PaypalPayments
	}
	return json.Marshal(toSerialize)
}

type NullableGETPaypalGateways200ResponseDataInnerRelationships struct {
	value *GETPaypalGateways200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETPaypalGateways200ResponseDataInnerRelationships) Get() *GETPaypalGateways200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETPaypalGateways200ResponseDataInnerRelationships) Set(val *GETPaypalGateways200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaypalGateways200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaypalGateways200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaypalGateways200ResponseDataInnerRelationships(val *GETPaypalGateways200ResponseDataInnerRelationships) *NullableGETPaypalGateways200ResponseDataInnerRelationships {
	return &NullableGETPaypalGateways200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETPaypalGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaypalGateways200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
