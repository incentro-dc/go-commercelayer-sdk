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

// TaxRuleCreateData struct for TaxRuleCreateData
type TaxRuleCreateData struct {
	// The resource's type
	Type          string                          `json:"type"`
	Attributes    TaxRuleCreateDataAttributes     `json:"attributes"`
	Relationships *TaxRuleCreateDataRelationships `json:"relationships,omitempty"`
}

// NewTaxRuleCreateData instantiates a new TaxRuleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleCreateData(type_ string, attributes TaxRuleCreateDataAttributes) *TaxRuleCreateData {
	this := TaxRuleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxRuleCreateDataWithDefaults instantiates a new TaxRuleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleCreateDataWithDefaults() *TaxRuleCreateData {
	this := TaxRuleCreateData{}
	var type_ string = "tax_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxRuleCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxRuleCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxRuleCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxRuleCreateData) GetAttributes() TaxRuleCreateDataAttributes {
	if o == nil {
		var ret TaxRuleCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxRuleCreateData) GetAttributesOk() (*TaxRuleCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxRuleCreateData) SetAttributes(v TaxRuleCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxRuleCreateData) GetRelationships() TaxRuleCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TaxRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleCreateData) GetRelationshipsOk() (*TaxRuleCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxRuleCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *TaxRuleCreateData) SetRelationships(v TaxRuleCreateDataRelationships) {
	o.Relationships = &v
}

func (o TaxRuleCreateData) MarshalJSON() ([]byte, error) {
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

type NullableTaxRuleCreateData struct {
	value *TaxRuleCreateData
	isSet bool
}

func (v NullableTaxRuleCreateData) Get() *TaxRuleCreateData {
	return v.value
}

func (v *NullableTaxRuleCreateData) Set(val *TaxRuleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleCreateData(val *TaxRuleCreateData) *NullableTaxRuleCreateData {
	return &NullableTaxRuleCreateData{value: val, isSet: true}
}

func (v NullableTaxRuleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
