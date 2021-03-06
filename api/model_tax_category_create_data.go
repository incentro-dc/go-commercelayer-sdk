/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TaxCategoryCreateData struct for TaxCategoryCreateData
type TaxCategoryCreateData struct {
	// The resource's type
	Type          string                              `json:"type"`
	Attributes    TaxCategoryCreateDataAttributes     `json:"attributes"`
	Relationships *TaxCategoryCreateDataRelationships `json:"relationships,omitempty"`
}

// NewTaxCategoryCreateData instantiates a new TaxCategoryCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryCreateData(type_ string, attributes TaxCategoryCreateDataAttributes) *TaxCategoryCreateData {
	this := TaxCategoryCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxCategoryCreateDataWithDefaults instantiates a new TaxCategoryCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryCreateDataWithDefaults() *TaxCategoryCreateData {
	this := TaxCategoryCreateData{}
	var type_ string = "tax_categories"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxCategoryCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxCategoryCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxCategoryCreateData) GetAttributes() TaxCategoryCreateDataAttributes {
	if o == nil {
		var ret TaxCategoryCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryCreateData) GetAttributesOk() (*TaxCategoryCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxCategoryCreateData) SetAttributes(v TaxCategoryCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxCategoryCreateData) GetRelationships() TaxCategoryCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TaxCategoryCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryCreateData) GetRelationshipsOk() (*TaxCategoryCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxCategoryCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxCategoryCreateDataRelationships and assigns it to the Relationships field.
func (o *TaxCategoryCreateData) SetRelationships(v TaxCategoryCreateDataRelationships) {
	o.Relationships = &v
}

func (o TaxCategoryCreateData) MarshalJSON() ([]byte, error) {
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

type NullableTaxCategoryCreateData struct {
	value *TaxCategoryCreateData
	isSet bool
}

func (v NullableTaxCategoryCreateData) Get() *TaxCategoryCreateData {
	return v.value
}

func (v *NullableTaxCategoryCreateData) Set(val *TaxCategoryCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryCreateData(val *TaxCategoryCreateData) *NullableTaxCategoryCreateData {
	return &NullableTaxCategoryCreateData{value: val, isSet: true}
}

func (v NullableTaxCategoryCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
