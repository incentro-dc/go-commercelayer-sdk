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

// GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData struct for GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData
type GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData instantiates a new GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData() *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData {
	this := GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData{}
	return &this
}

// NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayDataWithDefaults instantiates a new GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayDataWithDefaults() *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData {
	this := GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) SetId(v interface{}) {
	o.Id = v
}

func (o GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData struct {
	value *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData
	isSet bool
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) Get() *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData {
	return v.value
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) Set(val *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData(val *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData {
	return &NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData{value: val, isSet: true}
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsPaymentGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
