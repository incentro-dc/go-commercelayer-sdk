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

// StripeGatewayDataRelationshipsStripePayments struct for StripeGatewayDataRelationshipsStripePayments
type StripeGatewayDataRelationshipsStripePayments struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewStripeGatewayDataRelationshipsStripePayments instantiates a new StripeGatewayDataRelationshipsStripePayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayDataRelationshipsStripePayments(type_ string, id string) *StripeGatewayDataRelationshipsStripePayments {
	this := StripeGatewayDataRelationshipsStripePayments{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewStripeGatewayDataRelationshipsStripePaymentsWithDefaults instantiates a new StripeGatewayDataRelationshipsStripePayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayDataRelationshipsStripePaymentsWithDefaults() *StripeGatewayDataRelationshipsStripePayments {
	this := StripeGatewayDataRelationshipsStripePayments{}
	var type_ string = "stripe_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StripeGatewayDataRelationshipsStripePayments) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayDataRelationshipsStripePayments) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripeGatewayDataRelationshipsStripePayments) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *StripeGatewayDataRelationshipsStripePayments) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayDataRelationshipsStripePayments) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StripeGatewayDataRelationshipsStripePayments) SetId(v string) {
	o.Id = v
}

func (o StripeGatewayDataRelationshipsStripePayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableStripeGatewayDataRelationshipsStripePayments struct {
	value *StripeGatewayDataRelationshipsStripePayments
	isSet bool
}

func (v NullableStripeGatewayDataRelationshipsStripePayments) Get() *StripeGatewayDataRelationshipsStripePayments {
	return v.value
}

func (v *NullableStripeGatewayDataRelationshipsStripePayments) Set(val *StripeGatewayDataRelationshipsStripePayments) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayDataRelationshipsStripePayments) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayDataRelationshipsStripePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayDataRelationshipsStripePayments(val *StripeGatewayDataRelationshipsStripePayments) *NullableStripeGatewayDataRelationshipsStripePayments {
	return &NullableStripeGatewayDataRelationshipsStripePayments{value: val, isSet: true}
}

func (v NullableStripeGatewayDataRelationshipsStripePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayDataRelationshipsStripePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
