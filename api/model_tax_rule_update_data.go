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

// TaxRuleUpdateData struct for TaxRuleUpdateData
type TaxRuleUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                      `json:"id"`
	Attributes    TaxRuleUpdateDataAttributes `json:"attributes"`
	Relationships *TaxRuleDataRelationships   `json:"relationships,omitempty"`
}

// NewTaxRuleUpdateData instantiates a new TaxRuleUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleUpdateData(type_ string, id string, attributes TaxRuleUpdateDataAttributes) *TaxRuleUpdateData {
	this := TaxRuleUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewTaxRuleUpdateDataWithDefaults instantiates a new TaxRuleUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleUpdateDataWithDefaults() *TaxRuleUpdateData {
	this := TaxRuleUpdateData{}
	var type_ string = "tax_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxRuleUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxRuleUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxRuleUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TaxRuleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaxRuleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxRuleUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TaxRuleUpdateData) GetAttributes() TaxRuleUpdateDataAttributes {
	if o == nil {
		var ret TaxRuleUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxRuleUpdateData) GetAttributesOk() (*TaxRuleUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxRuleUpdateData) SetAttributes(v TaxRuleUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxRuleUpdateData) GetRelationships() TaxRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TaxRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleUpdateData) GetRelationshipsOk() (*TaxRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxRuleUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxRuleDataRelationships and assigns it to the Relationships field.
func (o *TaxRuleUpdateData) SetRelationships(v TaxRuleDataRelationships) {
	o.Relationships = &v
}

func (o TaxRuleUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableTaxRuleUpdateData struct {
	value *TaxRuleUpdateData
	isSet bool
}

func (v NullableTaxRuleUpdateData) Get() *TaxRuleUpdateData {
	return v.value
}

func (v *NullableTaxRuleUpdateData) Set(val *TaxRuleUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleUpdateData(val *TaxRuleUpdateData) *NullableTaxRuleUpdateData {
	return &NullableTaxRuleUpdateData{value: val, isSet: true}
}

func (v NullableTaxRuleUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
