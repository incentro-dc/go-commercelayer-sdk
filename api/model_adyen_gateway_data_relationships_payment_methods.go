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

// checks if the AdyenGatewayDataRelationshipsPaymentMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenGatewayDataRelationshipsPaymentMethods{}

// AdyenGatewayDataRelationshipsPaymentMethods struct for AdyenGatewayDataRelationshipsPaymentMethods
type AdyenGatewayDataRelationshipsPaymentMethods struct {
	Data *AdyenGatewayDataRelationshipsPaymentMethodsData `json:"data,omitempty"`
}

// NewAdyenGatewayDataRelationshipsPaymentMethods instantiates a new AdyenGatewayDataRelationshipsPaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayDataRelationshipsPaymentMethods() *AdyenGatewayDataRelationshipsPaymentMethods {
	this := AdyenGatewayDataRelationshipsPaymentMethods{}
	return &this
}

// NewAdyenGatewayDataRelationshipsPaymentMethodsWithDefaults instantiates a new AdyenGatewayDataRelationshipsPaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayDataRelationshipsPaymentMethodsWithDefaults() *AdyenGatewayDataRelationshipsPaymentMethods {
	this := AdyenGatewayDataRelationshipsPaymentMethods{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AdyenGatewayDataRelationshipsPaymentMethods) GetData() AdyenGatewayDataRelationshipsPaymentMethodsData {
	if o == nil || IsNil(o.Data) {
		var ret AdyenGatewayDataRelationshipsPaymentMethodsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayDataRelationshipsPaymentMethods) GetDataOk() (*AdyenGatewayDataRelationshipsPaymentMethodsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdyenGatewayDataRelationshipsPaymentMethods) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethodsData and assigns it to the Data field.
func (o *AdyenGatewayDataRelationshipsPaymentMethods) SetData(v AdyenGatewayDataRelationshipsPaymentMethodsData) {
	o.Data = &v
}

func (o AdyenGatewayDataRelationshipsPaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenGatewayDataRelationshipsPaymentMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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
