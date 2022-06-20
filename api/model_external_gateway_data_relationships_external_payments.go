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

// ExternalGatewayDataRelationshipsExternalPayments struct for ExternalGatewayDataRelationshipsExternalPayments
type ExternalGatewayDataRelationshipsExternalPayments struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewExternalGatewayDataRelationshipsExternalPayments instantiates a new ExternalGatewayDataRelationshipsExternalPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGatewayDataRelationshipsExternalPayments(type_ string, id string) *ExternalGatewayDataRelationshipsExternalPayments {
	this := ExternalGatewayDataRelationshipsExternalPayments{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewExternalGatewayDataRelationshipsExternalPaymentsWithDefaults instantiates a new ExternalGatewayDataRelationshipsExternalPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGatewayDataRelationshipsExternalPaymentsWithDefaults() *ExternalGatewayDataRelationshipsExternalPayments {
	this := ExternalGatewayDataRelationshipsExternalPayments{}
	var type_ string = "external_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ExternalGatewayDataRelationshipsExternalPayments) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataRelationshipsExternalPayments) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalGatewayDataRelationshipsExternalPayments) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ExternalGatewayDataRelationshipsExternalPayments) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataRelationshipsExternalPayments) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalGatewayDataRelationshipsExternalPayments) SetId(v string) {
	o.Id = v
}

func (o ExternalGatewayDataRelationshipsExternalPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGatewayDataRelationshipsExternalPayments struct {
	value *ExternalGatewayDataRelationshipsExternalPayments
	isSet bool
}

func (v NullableExternalGatewayDataRelationshipsExternalPayments) Get() *ExternalGatewayDataRelationshipsExternalPayments {
	return v.value
}

func (v *NullableExternalGatewayDataRelationshipsExternalPayments) Set(val *ExternalGatewayDataRelationshipsExternalPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGatewayDataRelationshipsExternalPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGatewayDataRelationshipsExternalPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGatewayDataRelationshipsExternalPayments(val *ExternalGatewayDataRelationshipsExternalPayments) *NullableExternalGatewayDataRelationshipsExternalPayments {
	return &NullableExternalGatewayDataRelationshipsExternalPayments{value: val, isSet: true}
}

func (v NullableExternalGatewayDataRelationshipsExternalPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGatewayDataRelationshipsExternalPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
