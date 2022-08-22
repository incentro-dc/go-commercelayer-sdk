/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaypalPaymentCreateData struct for PaypalPaymentCreateData
type PaypalPaymentCreateData struct {
	// The resource's type
	Type          string                                         `json:"type"`
	Attributes    POSTPaypalPayments201ResponseDataAttributes    `json:"attributes"`
	Relationships *POSTAdyenPayments201ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewPaypalPaymentCreateData instantiates a new PaypalPaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalPaymentCreateData(type_ string, attributes POSTPaypalPayments201ResponseDataAttributes) *PaypalPaymentCreateData {
	this := PaypalPaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaypalPaymentCreateDataWithDefaults instantiates a new PaypalPaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalPaymentCreateDataWithDefaults() *PaypalPaymentCreateData {
	this := PaypalPaymentCreateData{}
	var type_ string = "paypal_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PaypalPaymentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaypalPaymentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaypalPaymentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaypalPaymentCreateData) GetAttributes() POSTPaypalPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTPaypalPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaypalPaymentCreateData) GetAttributesOk() (*POSTPaypalPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaypalPaymentCreateData) SetAttributes(v POSTPaypalPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PaypalPaymentCreateData) GetRelationships() POSTAdyenPayments201ResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret POSTAdyenPayments201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalPaymentCreateData) GetRelationshipsOk() (*POSTAdyenPayments201ResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PaypalPaymentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTAdyenPayments201ResponseDataRelationships and assigns it to the Relationships field.
func (o *PaypalPaymentCreateData) SetRelationships(v POSTAdyenPayments201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o PaypalPaymentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullablePaypalPaymentCreateData struct {
	value *PaypalPaymentCreateData
	isSet bool
}

func (v NullablePaypalPaymentCreateData) Get() *PaypalPaymentCreateData {
	return v.value
}

func (v *NullablePaypalPaymentCreateData) Set(val *PaypalPaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalPaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalPaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalPaymentCreateData(val *PaypalPaymentCreateData) *NullablePaypalPaymentCreateData {
	return &NullablePaypalPaymentCreateData{value: val, isSet: true}
}

func (v NullablePaypalPaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalPaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
