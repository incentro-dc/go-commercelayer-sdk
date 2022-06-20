/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerPasswordResetData struct for CustomerPasswordResetData
type CustomerPasswordResetData struct {
	// The resource's type
	Type          string                                  `json:"type"`
	Attributes    CustomerPasswordResetDataAttributes     `json:"attributes"`
	Relationships *CouponRecipientCreateDataRelationships `json:"relationships,omitempty"`
}

// NewCustomerPasswordResetData instantiates a new CustomerPasswordResetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetData(type_ string, attributes CustomerPasswordResetDataAttributes) *CustomerPasswordResetData {
	this := CustomerPasswordResetData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerPasswordResetDataWithDefaults instantiates a new CustomerPasswordResetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetDataWithDefaults() *CustomerPasswordResetData {
	this := CustomerPasswordResetData{}
	var type_ string = "customer_password_resets"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerPasswordResetData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPasswordResetData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerPasswordResetData) GetAttributes() CustomerPasswordResetDataAttributes {
	if o == nil {
		var ret CustomerPasswordResetDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetData) GetAttributesOk() (*CustomerPasswordResetDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerPasswordResetData) SetAttributes(v CustomerPasswordResetDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPasswordResetData) GetRelationships() CouponRecipientCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponRecipientCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPasswordResetData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerPasswordResetData) SetRelationships(v CouponRecipientCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerPasswordResetData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerPasswordResetData struct {
	value *CustomerPasswordResetData
	isSet bool
}

func (v NullableCustomerPasswordResetData) Get() *CustomerPasswordResetData {
	return v.value
}

func (v *NullableCustomerPasswordResetData) Set(val *CustomerPasswordResetData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetData(val *CustomerPasswordResetData) *NullableCustomerPasswordResetData {
	return &NullableCustomerPasswordResetData{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
