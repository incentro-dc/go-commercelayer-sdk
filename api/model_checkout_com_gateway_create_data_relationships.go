/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CheckoutComGatewayCreateDataRelationships struct for CheckoutComGatewayCreateDataRelationships
type CheckoutComGatewayCreateDataRelationships struct {
	CheckoutComPayments *CheckoutComGatewayDataRelationshipsCheckoutComPayments `json:"checkout_com_payments,omitempty"`
}

// NewCheckoutComGatewayCreateDataRelationships instantiates a new CheckoutComGatewayCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayCreateDataRelationships() *CheckoutComGatewayCreateDataRelationships {
	this := CheckoutComGatewayCreateDataRelationships{}
	return &this
}

// NewCheckoutComGatewayCreateDataRelationshipsWithDefaults instantiates a new CheckoutComGatewayCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayCreateDataRelationshipsWithDefaults() *CheckoutComGatewayCreateDataRelationships {
	this := CheckoutComGatewayCreateDataRelationships{}
	return &this
}

// GetCheckoutComPayments returns the CheckoutComPayments field value if set, zero value otherwise.
func (o *CheckoutComGatewayCreateDataRelationships) GetCheckoutComPayments() CheckoutComGatewayDataRelationshipsCheckoutComPayments {
	if o == nil || o.CheckoutComPayments == nil {
		var ret CheckoutComGatewayDataRelationshipsCheckoutComPayments
		return ret
	}
	return *o.CheckoutComPayments
}

// GetCheckoutComPaymentsOk returns a tuple with the CheckoutComPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayCreateDataRelationships) GetCheckoutComPaymentsOk() (*CheckoutComGatewayDataRelationshipsCheckoutComPayments, bool) {
	if o == nil || o.CheckoutComPayments == nil {
		return nil, false
	}
	return o.CheckoutComPayments, true
}

// HasCheckoutComPayments returns a boolean if a field has been set.
func (o *CheckoutComGatewayCreateDataRelationships) HasCheckoutComPayments() bool {
	if o != nil && o.CheckoutComPayments != nil {
		return true
	}

	return false
}

// SetCheckoutComPayments gets a reference to the given CheckoutComGatewayDataRelationshipsCheckoutComPayments and assigns it to the CheckoutComPayments field.
func (o *CheckoutComGatewayCreateDataRelationships) SetCheckoutComPayments(v CheckoutComGatewayDataRelationshipsCheckoutComPayments) {
	o.CheckoutComPayments = &v
}

func (o CheckoutComGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckoutComPayments != nil {
		toSerialize["checkout_com_payments"] = o.CheckoutComPayments
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComGatewayCreateDataRelationships struct {
	value *CheckoutComGatewayCreateDataRelationships
	isSet bool
}

func (v NullableCheckoutComGatewayCreateDataRelationships) Get() *CheckoutComGatewayCreateDataRelationships {
	return v.value
}

func (v *NullableCheckoutComGatewayCreateDataRelationships) Set(val *CheckoutComGatewayCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayCreateDataRelationships(val *CheckoutComGatewayCreateDataRelationships) *NullableCheckoutComGatewayCreateDataRelationships {
	return &NullableCheckoutComGatewayCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
