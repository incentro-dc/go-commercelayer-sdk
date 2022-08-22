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

// GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments struct for GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments
type GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments(type_ string, id string) *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	this := GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsWithDefaults instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPaymentsWithDefaults() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	this := GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments{}
	var type_ string = "checkout_com_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) SetId(v string) {
	o.Id = v
}

func (o GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments struct {
	value *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments
	isSet bool
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) Get() *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	return v.value
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) Set(val *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments(val *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	return &NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
