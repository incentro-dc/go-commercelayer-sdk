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

// CustomerData struct for CustomerData
type CustomerData struct {
	// The resource's type
	Type          string                     `json:"type"`
	Attributes    CustomerDataAttributes     `json:"attributes"`
	Relationships *CustomerDataRelationships `json:"relationships,omitempty"`
}

// NewCustomerData instantiates a new CustomerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerData(type_ string, attributes CustomerDataAttributes) *CustomerData {
	this := CustomerData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerDataWithDefaults instantiates a new CustomerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataWithDefaults() *CustomerData {
	this := CustomerData{}
	var type_ string = "customers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerData) GetAttributes() CustomerDataAttributes {
	if o == nil {
		var ret CustomerDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerData) GetAttributesOk() (*CustomerDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerData) SetAttributes(v CustomerDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerData) GetRelationships() CustomerDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CustomerDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerData) GetRelationshipsOk() (*CustomerDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerDataRelationships and assigns it to the Relationships field.
func (o *CustomerData) SetRelationships(v CustomerDataRelationships) {
	o.Relationships = &v
}

func (o CustomerData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerData struct {
	value *CustomerData
	isSet bool
}

func (v NullableCustomerData) Get() *CustomerData {
	return v.value
}

func (v *NullableCustomerData) Set(val *CustomerData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerData(val *CustomerData) *NullableCustomerData {
	return &NullableCustomerData{value: val, isSet: true}
}

func (v NullableCustomerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
