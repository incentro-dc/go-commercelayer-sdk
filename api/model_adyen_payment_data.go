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

// AdyenPaymentData struct for AdyenPaymentData
type AdyenPaymentData struct {
	// The resource's type
	Type          interface{}                                    `json:"type"`
	Attributes    GETAdyenPayments200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AdyenPaymentDataRelationships                 `json:"relationships,omitempty"`
}

// NewAdyenPaymentData instantiates a new AdyenPaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentData(type_ interface{}, attributes GETAdyenPayments200ResponseDataInnerAttributes) *AdyenPaymentData {
	this := AdyenPaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdyenPaymentDataWithDefaults instantiates a new AdyenPaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentDataWithDefaults() *AdyenPaymentData {
	this := AdyenPaymentData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdyenPaymentData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdyenPaymentData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenPaymentData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenPaymentData) GetAttributes() GETAdyenPayments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETAdyenPayments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenPaymentData) GetAttributesOk() (*GETAdyenPayments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenPaymentData) SetAttributes(v GETAdyenPayments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenPaymentData) GetRelationships() AdyenPaymentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenPaymentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentDataRelationships and assigns it to the Relationships field.
func (o *AdyenPaymentData) SetRelationships(v AdyenPaymentDataRelationships) {
	o.Relationships = &v
}

func (o AdyenPaymentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenPaymentData struct {
	value *AdyenPaymentData
	isSet bool
}

func (v NullableAdyenPaymentData) Get() *AdyenPaymentData {
	return v.value
}

func (v *NullableAdyenPaymentData) Set(val *AdyenPaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentData(val *AdyenPaymentData) *NullableAdyenPaymentData {
	return &NullableAdyenPaymentData{value: val, isSet: true}
}

func (v NullableAdyenPaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
