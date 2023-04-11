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

// checks if the POSTKlarnaGatewaysRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTKlarnaGatewaysRequestDataRelationships{}

// POSTKlarnaGatewaysRequestDataRelationships struct for POSTKlarnaGatewaysRequestDataRelationships
type POSTKlarnaGatewaysRequestDataRelationships struct {
	KlarnaPayments *POSTKlarnaGatewaysRequestDataRelationshipsKlarnaPayments `json:"klarna_payments,omitempty"`
}

// NewPOSTKlarnaGatewaysRequestDataRelationships instantiates a new POSTKlarnaGatewaysRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTKlarnaGatewaysRequestDataRelationships() *POSTKlarnaGatewaysRequestDataRelationships {
	this := POSTKlarnaGatewaysRequestDataRelationships{}
	return &this
}

// NewPOSTKlarnaGatewaysRequestDataRelationshipsWithDefaults instantiates a new POSTKlarnaGatewaysRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTKlarnaGatewaysRequestDataRelationshipsWithDefaults() *POSTKlarnaGatewaysRequestDataRelationships {
	this := POSTKlarnaGatewaysRequestDataRelationships{}
	return &this
}

// GetKlarnaPayments returns the KlarnaPayments field value if set, zero value otherwise.
func (o *POSTKlarnaGatewaysRequestDataRelationships) GetKlarnaPayments() POSTKlarnaGatewaysRequestDataRelationshipsKlarnaPayments {
	if o == nil || IsNil(o.KlarnaPayments) {
		var ret POSTKlarnaGatewaysRequestDataRelationshipsKlarnaPayments
		return ret
	}
	return *o.KlarnaPayments
}

// GetKlarnaPaymentsOk returns a tuple with the KlarnaPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTKlarnaGatewaysRequestDataRelationships) GetKlarnaPaymentsOk() (*POSTKlarnaGatewaysRequestDataRelationshipsKlarnaPayments, bool) {
	if o == nil || IsNil(o.KlarnaPayments) {
		return nil, false
	}
	return o.KlarnaPayments, true
}

// HasKlarnaPayments returns a boolean if a field has been set.
func (o *POSTKlarnaGatewaysRequestDataRelationships) HasKlarnaPayments() bool {
	if o != nil && !IsNil(o.KlarnaPayments) {
		return true
	}

	return false
}

// SetKlarnaPayments gets a reference to the given POSTKlarnaGatewaysRequestDataRelationshipsKlarnaPayments and assigns it to the KlarnaPayments field.
func (o *POSTKlarnaGatewaysRequestDataRelationships) SetKlarnaPayments(v POSTKlarnaGatewaysRequestDataRelationshipsKlarnaPayments) {
	o.KlarnaPayments = &v
}

func (o POSTKlarnaGatewaysRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTKlarnaGatewaysRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KlarnaPayments) {
		toSerialize["klarna_payments"] = o.KlarnaPayments
	}
	return toSerialize, nil
}

type NullablePOSTKlarnaGatewaysRequestDataRelationships struct {
	value *POSTKlarnaGatewaysRequestDataRelationships
	isSet bool
}

func (v NullablePOSTKlarnaGatewaysRequestDataRelationships) Get() *POSTKlarnaGatewaysRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTKlarnaGatewaysRequestDataRelationships) Set(val *POSTKlarnaGatewaysRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTKlarnaGatewaysRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTKlarnaGatewaysRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTKlarnaGatewaysRequestDataRelationships(val *POSTKlarnaGatewaysRequestDataRelationships) *NullablePOSTKlarnaGatewaysRequestDataRelationships {
	return &NullablePOSTKlarnaGatewaysRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTKlarnaGatewaysRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTKlarnaGatewaysRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}