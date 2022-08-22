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

// GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway struct for GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway
type GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway instantiates a new GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway(type_ string, id string) *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway {
	this := GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayWithDefaults instantiates a new GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayWithDefaults() *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway {
	this := GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway{}
	var type_ string = "payment_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) SetId(v string) {
	o.Id = v
}

func (o GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway struct {
	value *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway
	isSet bool
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) Get() *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway {
	return v.value
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) Set(val *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway(val *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway {
	return &NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway{value: val, isSet: true}
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}