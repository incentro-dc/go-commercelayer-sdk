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

// GETAxerveGateways200ResponseDataInnerRelationships struct for GETAxerveGateways200ResponseDataInnerRelationships
type GETAxerveGateways200ResponseDataInnerRelationships struct {
	PaymentMethods *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	AxervePayments *GETAxerveGateways200ResponseDataInnerRelationshipsAxervePayments `json:"axerve_payments,omitempty"`
}

// NewGETAxerveGateways200ResponseDataInnerRelationships instantiates a new GETAxerveGateways200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAxerveGateways200ResponseDataInnerRelationships() *GETAxerveGateways200ResponseDataInnerRelationships {
	this := GETAxerveGateways200ResponseDataInnerRelationships{}
	return &this
}

// NewGETAxerveGateways200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETAxerveGateways200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAxerveGateways200ResponseDataInnerRelationshipsWithDefaults() *GETAxerveGateways200ResponseDataInnerRelationships {
	this := GETAxerveGateways200ResponseDataInnerRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) GetPaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) GetPaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) SetPaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetAxervePayments returns the AxervePayments field value if set, zero value otherwise.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) GetAxervePayments() GETAxerveGateways200ResponseDataInnerRelationshipsAxervePayments {
	if o == nil || o.AxervePayments == nil {
		var ret GETAxerveGateways200ResponseDataInnerRelationshipsAxervePayments
		return ret
	}
	return *o.AxervePayments
}

// GetAxervePaymentsOk returns a tuple with the AxervePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) GetAxervePaymentsOk() (*GETAxerveGateways200ResponseDataInnerRelationshipsAxervePayments, bool) {
	if o == nil || o.AxervePayments == nil {
		return nil, false
	}
	return o.AxervePayments, true
}

// HasAxervePayments returns a boolean if a field has been set.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) HasAxervePayments() bool {
	if o != nil && o.AxervePayments != nil {
		return true
	}

	return false
}

// SetAxervePayments gets a reference to the given GETAxerveGateways200ResponseDataInnerRelationshipsAxervePayments and assigns it to the AxervePayments field.
func (o *GETAxerveGateways200ResponseDataInnerRelationships) SetAxervePayments(v GETAxerveGateways200ResponseDataInnerRelationshipsAxervePayments) {
	o.AxervePayments = &v
}

func (o GETAxerveGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.AxervePayments != nil {
		toSerialize["axerve_payments"] = o.AxervePayments
	}
	return json.Marshal(toSerialize)
}

type NullableGETAxerveGateways200ResponseDataInnerRelationships struct {
	value *GETAxerveGateways200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETAxerveGateways200ResponseDataInnerRelationships) Get() *GETAxerveGateways200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETAxerveGateways200ResponseDataInnerRelationships) Set(val *GETAxerveGateways200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAxerveGateways200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAxerveGateways200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAxerveGateways200ResponseDataInnerRelationships(val *GETAxerveGateways200ResponseDataInnerRelationships) *NullableGETAxerveGateways200ResponseDataInnerRelationships {
	return &NullableGETAxerveGateways200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETAxerveGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAxerveGateways200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
