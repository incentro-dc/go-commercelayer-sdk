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

// checks if the POSTSatispayGatewaysRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSatispayGatewaysRequestDataRelationships{}

// POSTSatispayGatewaysRequestDataRelationships struct for POSTSatispayGatewaysRequestDataRelationships
type POSTSatispayGatewaysRequestDataRelationships struct {
	SatispayPayments *POSTSatispayGatewaysRequestDataRelationshipsSatispayPayments `json:"satispay_payments,omitempty"`
}

// NewPOSTSatispayGatewaysRequestDataRelationships instantiates a new POSTSatispayGatewaysRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSatispayGatewaysRequestDataRelationships() *POSTSatispayGatewaysRequestDataRelationships {
	this := POSTSatispayGatewaysRequestDataRelationships{}
	return &this
}

// NewPOSTSatispayGatewaysRequestDataRelationshipsWithDefaults instantiates a new POSTSatispayGatewaysRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSatispayGatewaysRequestDataRelationshipsWithDefaults() *POSTSatispayGatewaysRequestDataRelationships {
	this := POSTSatispayGatewaysRequestDataRelationships{}
	return &this
}

// GetSatispayPayments returns the SatispayPayments field value if set, zero value otherwise.
func (o *POSTSatispayGatewaysRequestDataRelationships) GetSatispayPayments() POSTSatispayGatewaysRequestDataRelationshipsSatispayPayments {
	if o == nil || IsNil(o.SatispayPayments) {
		var ret POSTSatispayGatewaysRequestDataRelationshipsSatispayPayments
		return ret
	}
	return *o.SatispayPayments
}

// GetSatispayPaymentsOk returns a tuple with the SatispayPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSatispayGatewaysRequestDataRelationships) GetSatispayPaymentsOk() (*POSTSatispayGatewaysRequestDataRelationshipsSatispayPayments, bool) {
	if o == nil || IsNil(o.SatispayPayments) {
		return nil, false
	}
	return o.SatispayPayments, true
}

// HasSatispayPayments returns a boolean if a field has been set.
func (o *POSTSatispayGatewaysRequestDataRelationships) HasSatispayPayments() bool {
	if o != nil && !IsNil(o.SatispayPayments) {
		return true
	}

	return false
}

// SetSatispayPayments gets a reference to the given POSTSatispayGatewaysRequestDataRelationshipsSatispayPayments and assigns it to the SatispayPayments field.
func (o *POSTSatispayGatewaysRequestDataRelationships) SetSatispayPayments(v POSTSatispayGatewaysRequestDataRelationshipsSatispayPayments) {
	o.SatispayPayments = &v
}

func (o POSTSatispayGatewaysRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSatispayGatewaysRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SatispayPayments) {
		toSerialize["satispay_payments"] = o.SatispayPayments
	}
	return toSerialize, nil
}

type NullablePOSTSatispayGatewaysRequestDataRelationships struct {
	value *POSTSatispayGatewaysRequestDataRelationships
	isSet bool
}

func (v NullablePOSTSatispayGatewaysRequestDataRelationships) Get() *POSTSatispayGatewaysRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTSatispayGatewaysRequestDataRelationships) Set(val *POSTSatispayGatewaysRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSatispayGatewaysRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSatispayGatewaysRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSatispayGatewaysRequestDataRelationships(val *POSTSatispayGatewaysRequestDataRelationships) *NullablePOSTSatispayGatewaysRequestDataRelationships {
	return &NullablePOSTSatispayGatewaysRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTSatispayGatewaysRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSatispayGatewaysRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}