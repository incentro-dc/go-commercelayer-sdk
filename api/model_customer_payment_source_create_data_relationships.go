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

// CustomerPaymentSourceCreateDataRelationships struct for CustomerPaymentSourceCreateDataRelationships
type CustomerPaymentSourceCreateDataRelationships struct {
	Customer      CouponRecipientCreateDataRelationshipsCustomer            `json:"customer"`
	PaymentSource CustomerPaymentSourceCreateDataRelationshipsPaymentSource `json:"payment_source"`
}

// NewCustomerPaymentSourceCreateDataRelationships instantiates a new CustomerPaymentSourceCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceCreateDataRelationships(customer CouponRecipientCreateDataRelationshipsCustomer, paymentSource CustomerPaymentSourceCreateDataRelationshipsPaymentSource) *CustomerPaymentSourceCreateDataRelationships {
	this := CustomerPaymentSourceCreateDataRelationships{}
	this.Customer = customer
	this.PaymentSource = paymentSource
	return &this
}

// NewCustomerPaymentSourceCreateDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceCreateDataRelationshipsWithDefaults() *CustomerPaymentSourceCreateDataRelationships {
	this := CustomerPaymentSourceCreateDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value
func (o *CustomerPaymentSourceCreateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer {
	if o == nil {
		var ret CouponRecipientCreateDataRelationshipsCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *CustomerPaymentSourceCreateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer) {
	o.Customer = v
}

// GetPaymentSource returns the PaymentSource field value
func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentSource() CustomerPaymentSourceCreateDataRelationshipsPaymentSource {
	if o == nil {
		var ret CustomerPaymentSourceCreateDataRelationshipsPaymentSource
		return ret
	}

	return o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceCreateDataRelationshipsPaymentSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentSource, true
}

// SetPaymentSource sets field value
func (o *CustomerPaymentSourceCreateDataRelationships) SetPaymentSource(v CustomerPaymentSourceCreateDataRelationshipsPaymentSource) {
	o.PaymentSource = v
}

func (o CustomerPaymentSourceCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer"] = o.Customer
	}
	if true {
		toSerialize["payment_source"] = o.PaymentSource
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSourceCreateDataRelationships struct {
	value *CustomerPaymentSourceCreateDataRelationships
	isSet bool
}

func (v NullableCustomerPaymentSourceCreateDataRelationships) Get() *CustomerPaymentSourceCreateDataRelationships {
	return v.value
}

func (v *NullableCustomerPaymentSourceCreateDataRelationships) Set(val *CustomerPaymentSourceCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceCreateDataRelationships(val *CustomerPaymentSourceCreateDataRelationships) *NullableCustomerPaymentSourceCreateDataRelationships {
	return &NullableCustomerPaymentSourceCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
