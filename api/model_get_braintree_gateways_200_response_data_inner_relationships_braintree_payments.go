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

// GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments struct for GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments
type GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments instantiates a new GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments(type_ string, id string) *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	this := GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsWithDefaults instantiates a new GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsWithDefaults() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	this := GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments{}
	var type_ string = "braintree_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) SetId(v string) {
	o.Id = v
}

func (o GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments struct {
	value *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments
	isSet bool
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) Get() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	return v.value
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) Set(val *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments(val *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	return &NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments{value: val, isSet: true}
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
