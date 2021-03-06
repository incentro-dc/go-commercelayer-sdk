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

// AdyenGatewayDataRelationshipsPaymentMethods struct for AdyenGatewayDataRelationshipsPaymentMethods
type AdyenGatewayDataRelationshipsPaymentMethods struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewAdyenGatewayDataRelationshipsPaymentMethods instantiates a new AdyenGatewayDataRelationshipsPaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayDataRelationshipsPaymentMethods(type_ string, id string) *AdyenGatewayDataRelationshipsPaymentMethods {
	this := AdyenGatewayDataRelationshipsPaymentMethods{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAdyenGatewayDataRelationshipsPaymentMethodsWithDefaults instantiates a new AdyenGatewayDataRelationshipsPaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayDataRelationshipsPaymentMethodsWithDefaults() *AdyenGatewayDataRelationshipsPaymentMethods {
	this := AdyenGatewayDataRelationshipsPaymentMethods{}
	var type_ string = "payment_methods"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AdyenGatewayDataRelationshipsPaymentMethods) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationshipsPaymentMethods) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenGatewayDataRelationshipsPaymentMethods) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AdyenGatewayDataRelationshipsPaymentMethods) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationshipsPaymentMethods) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdyenGatewayDataRelationshipsPaymentMethods) SetId(v string) {
	o.Id = v
}

func (o AdyenGatewayDataRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenGatewayDataRelationshipsPaymentMethods struct {
	value *AdyenGatewayDataRelationshipsPaymentMethods
	isSet bool
}

func (v NullableAdyenGatewayDataRelationshipsPaymentMethods) Get() *AdyenGatewayDataRelationshipsPaymentMethods {
	return v.value
}

func (v *NullableAdyenGatewayDataRelationshipsPaymentMethods) Set(val *AdyenGatewayDataRelationshipsPaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayDataRelationshipsPaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayDataRelationshipsPaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayDataRelationshipsPaymentMethods(val *AdyenGatewayDataRelationshipsPaymentMethods) *NullableAdyenGatewayDataRelationshipsPaymentMethods {
	return &NullableAdyenGatewayDataRelationshipsPaymentMethods{value: val, isSet: true}
}

func (v NullableAdyenGatewayDataRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayDataRelationshipsPaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
