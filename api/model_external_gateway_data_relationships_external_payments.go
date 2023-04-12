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

// checks if the ExternalGatewayDataRelationshipsExternalPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalGatewayDataRelationshipsExternalPayments{}

// ExternalGatewayDataRelationshipsExternalPayments struct for ExternalGatewayDataRelationshipsExternalPayments
type ExternalGatewayDataRelationshipsExternalPayments struct {
	Data *ExternalGatewayDataRelationshipsExternalPaymentsData `json:"data,omitempty"`
}

// NewExternalGatewayDataRelationshipsExternalPayments instantiates a new ExternalGatewayDataRelationshipsExternalPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGatewayDataRelationshipsExternalPayments() *ExternalGatewayDataRelationshipsExternalPayments {
	this := ExternalGatewayDataRelationshipsExternalPayments{}
	return &this
}

// NewExternalGatewayDataRelationshipsExternalPaymentsWithDefaults instantiates a new ExternalGatewayDataRelationshipsExternalPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGatewayDataRelationshipsExternalPaymentsWithDefaults() *ExternalGatewayDataRelationshipsExternalPayments {
	this := ExternalGatewayDataRelationshipsExternalPayments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExternalGatewayDataRelationshipsExternalPayments) GetData() ExternalGatewayDataRelationshipsExternalPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret ExternalGatewayDataRelationshipsExternalPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataRelationshipsExternalPayments) GetDataOk() (*ExternalGatewayDataRelationshipsExternalPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExternalGatewayDataRelationshipsExternalPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ExternalGatewayDataRelationshipsExternalPaymentsData and assigns it to the Data field.
func (o *ExternalGatewayDataRelationshipsExternalPayments) SetData(v ExternalGatewayDataRelationshipsExternalPaymentsData) {
	o.Data = &v
}

func (o ExternalGatewayDataRelationshipsExternalPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalGatewayDataRelationshipsExternalPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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
