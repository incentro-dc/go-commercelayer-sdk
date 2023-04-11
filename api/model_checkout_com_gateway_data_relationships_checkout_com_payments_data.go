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

// CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData struct for CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData
type CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData instantiates a new CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData() *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData {
	this := CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData{}
	return &this
}

// NewCheckoutComGatewayDataRelationshipsCheckoutComPaymentsDataWithDefaults instantiates a new CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayDataRelationshipsCheckoutComPaymentsDataWithDefaults() *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData {
	this := CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) SetId(v interface{}) {
	o.Id = v
}

func (o CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData struct {
	value *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData
	isSet bool
}

func (v NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) Get() *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData {
	return v.value
}

func (v *NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) Set(val *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData(val *CheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) *NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData {
	return &NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayDataRelationshipsCheckoutComPaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
