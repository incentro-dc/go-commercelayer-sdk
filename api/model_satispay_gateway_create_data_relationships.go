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

// checks if the SatispayGatewayCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayGatewayCreateDataRelationships{}

// SatispayGatewayCreateDataRelationships struct for SatispayGatewayCreateDataRelationships
type SatispayGatewayCreateDataRelationships struct {
	SatispayPayments *SatispayGatewayCreateDataRelationshipsSatispayPayments `json:"satispay_payments,omitempty"`
}

// NewSatispayGatewayCreateDataRelationships instantiates a new SatispayGatewayCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGatewayCreateDataRelationships() *SatispayGatewayCreateDataRelationships {
	this := SatispayGatewayCreateDataRelationships{}
	return &this
}

// NewSatispayGatewayCreateDataRelationshipsWithDefaults instantiates a new SatispayGatewayCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayCreateDataRelationshipsWithDefaults() *SatispayGatewayCreateDataRelationships {
	this := SatispayGatewayCreateDataRelationships{}
	return &this
}

// GetSatispayPayments returns the SatispayPayments field value if set, zero value otherwise.
func (o *SatispayGatewayCreateDataRelationships) GetSatispayPayments() SatispayGatewayCreateDataRelationshipsSatispayPayments {
	if o == nil || IsNil(o.SatispayPayments) {
		var ret SatispayGatewayCreateDataRelationshipsSatispayPayments
		return ret
	}
	return *o.SatispayPayments
}

// GetSatispayPaymentsOk returns a tuple with the SatispayPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayCreateDataRelationships) GetSatispayPaymentsOk() (*SatispayGatewayCreateDataRelationshipsSatispayPayments, bool) {
	if o == nil || IsNil(o.SatispayPayments) {
		return nil, false
	}
	return o.SatispayPayments, true
}

// HasSatispayPayments returns a boolean if a field has been set.
func (o *SatispayGatewayCreateDataRelationships) HasSatispayPayments() bool {
	if o != nil && !IsNil(o.SatispayPayments) {
		return true
	}

	return false
}

// SetSatispayPayments gets a reference to the given SatispayGatewayCreateDataRelationshipsSatispayPayments and assigns it to the SatispayPayments field.
func (o *SatispayGatewayCreateDataRelationships) SetSatispayPayments(v SatispayGatewayCreateDataRelationshipsSatispayPayments) {
	o.SatispayPayments = &v
}

func (o SatispayGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayGatewayCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SatispayPayments) {
		toSerialize["satispay_payments"] = o.SatispayPayments
	}
	return toSerialize, nil
}

type NullableSatispayGatewayCreateDataRelationships struct {
	value *SatispayGatewayCreateDataRelationships
	isSet bool
}

func (v NullableSatispayGatewayCreateDataRelationships) Get() *SatispayGatewayCreateDataRelationships {
	return v.value
}

func (v *NullableSatispayGatewayCreateDataRelationships) Set(val *SatispayGatewayCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGatewayCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGatewayCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGatewayCreateDataRelationships(val *SatispayGatewayCreateDataRelationships) *NullableSatispayGatewayCreateDataRelationships {
	return &NullableSatispayGatewayCreateDataRelationships{value: val, isSet: true}
}

func (v NullableSatispayGatewayCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGatewayCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}