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

// GETKlarnaGateways200ResponseDataInnerRelationships struct for GETKlarnaGateways200ResponseDataInnerRelationships
type GETKlarnaGateways200ResponseDataInnerRelationships struct {
	PaymentMethods *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	KlarnaPayments *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments `json:"klarna_payments,omitempty"`
}

// NewGETKlarnaGateways200ResponseDataInnerRelationships instantiates a new GETKlarnaGateways200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETKlarnaGateways200ResponseDataInnerRelationships() *GETKlarnaGateways200ResponseDataInnerRelationships {
	this := GETKlarnaGateways200ResponseDataInnerRelationships{}
	return &this
}

// NewGETKlarnaGateways200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETKlarnaGateways200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETKlarnaGateways200ResponseDataInnerRelationshipsWithDefaults() *GETKlarnaGateways200ResponseDataInnerRelationships {
	this := GETKlarnaGateways200ResponseDataInnerRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) GetPaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) GetPaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) SetPaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetKlarnaPayments returns the KlarnaPayments field value if set, zero value otherwise.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) GetKlarnaPayments() GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments {
	if o == nil || o.KlarnaPayments == nil {
		var ret GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments
		return ret
	}
	return *o.KlarnaPayments
}

// GetKlarnaPaymentsOk returns a tuple with the KlarnaPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) GetKlarnaPaymentsOk() (*GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments, bool) {
	if o == nil || o.KlarnaPayments == nil {
		return nil, false
	}
	return o.KlarnaPayments, true
}

// HasKlarnaPayments returns a boolean if a field has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) HasKlarnaPayments() bool {
	if o != nil && o.KlarnaPayments != nil {
		return true
	}

	return false
}

// SetKlarnaPayments gets a reference to the given GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments and assigns it to the KlarnaPayments field.
func (o *GETKlarnaGateways200ResponseDataInnerRelationships) SetKlarnaPayments(v GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) {
	o.KlarnaPayments = &v
}

func (o GETKlarnaGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.KlarnaPayments != nil {
		toSerialize["klarna_payments"] = o.KlarnaPayments
	}
	return json.Marshal(toSerialize)
}

type NullableGETKlarnaGateways200ResponseDataInnerRelationships struct {
	value *GETKlarnaGateways200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationships) Get() *GETKlarnaGateways200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationships) Set(val *GETKlarnaGateways200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETKlarnaGateways200ResponseDataInnerRelationships(val *GETKlarnaGateways200ResponseDataInnerRelationships) *NullableGETKlarnaGateways200ResponseDataInnerRelationships {
	return &NullableGETKlarnaGateways200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}