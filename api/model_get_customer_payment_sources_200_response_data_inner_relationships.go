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

// GETCustomerPaymentSources200ResponseDataInnerRelationships struct for GETCustomerPaymentSources200ResponseDataInnerRelationships
type GETCustomerPaymentSources200ResponseDataInnerRelationships struct {
	Customer      *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer            `json:"customer,omitempty"`
	PaymentSource *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource `json:"payment_source,omitempty"`
}

// NewGETCustomerPaymentSources200ResponseDataInnerRelationships instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerPaymentSources200ResponseDataInnerRelationships() *GETCustomerPaymentSources200ResponseDataInnerRelationships {
	this := GETCustomerPaymentSources200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsWithDefaults() *GETCustomerPaymentSources200ResponseDataInnerRelationships {
	this := GETCustomerPaymentSources200ResponseDataInnerRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) GetPaymentSource() GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	if o == nil || o.PaymentSource == nil {
		var ret GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) GetPaymentSourceOk() (*GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource and assigns it to the PaymentSource field.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationships) SetPaymentSource(v GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) {
	o.PaymentSource = &v
}

func (o GETCustomerPaymentSources200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerPaymentSources200ResponseDataInnerRelationships struct {
	value *GETCustomerPaymentSources200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationships) Get() *GETCustomerPaymentSources200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationships) Set(val *GETCustomerPaymentSources200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPaymentSources200ResponseDataInnerRelationships(val *GETCustomerPaymentSources200ResponseDataInnerRelationships) *NullableGETCustomerPaymentSources200ResponseDataInnerRelationships {
	return &NullableGETCustomerPaymentSources200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
