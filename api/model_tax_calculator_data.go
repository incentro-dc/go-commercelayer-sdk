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

// TaxCalculatorData struct for TaxCalculatorData
type TaxCalculatorData struct {
	// The resource's type
	Type          string                            `json:"type"`
	Attributes    ManualTaxCalculatorDataAttributes `json:"attributes"`
	Relationships *AvalaraAccountDataRelationships  `json:"relationships,omitempty"`
}

// NewTaxCalculatorData instantiates a new TaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCalculatorData(type_ string, attributes ManualTaxCalculatorDataAttributes) *TaxCalculatorData {
	this := TaxCalculatorData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTaxCalculatorDataWithDefaults instantiates a new TaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCalculatorDataWithDefaults() *TaxCalculatorData {
	this := TaxCalculatorData{}
	var type_ string = "tax_calculators"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxCalculatorData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxCalculatorData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxCalculatorData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TaxCalculatorData) GetAttributes() ManualTaxCalculatorDataAttributes {
	if o == nil {
		var ret ManualTaxCalculatorDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxCalculatorData) GetAttributesOk() (*ManualTaxCalculatorDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxCalculatorData) SetAttributes(v ManualTaxCalculatorDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxCalculatorData) GetRelationships() AvalaraAccountDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AvalaraAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCalculatorData) GetRelationshipsOk() (*AvalaraAccountDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxCalculatorData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AvalaraAccountDataRelationships and assigns it to the Relationships field.
func (o *TaxCalculatorData) SetRelationships(v AvalaraAccountDataRelationships) {
	o.Relationships = &v
}

func (o TaxCalculatorData) MarshalJSON() ([]byte, error) {
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

type NullableTaxCalculatorData struct {
	value *TaxCalculatorData
	isSet bool
}

func (v NullableTaxCalculatorData) Get() *TaxCalculatorData {
	return v.value
}

func (v *NullableTaxCalculatorData) Set(val *TaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCalculatorData(val *TaxCalculatorData) *NullableTaxCalculatorData {
	return &NullableTaxCalculatorData{value: val, isSet: true}
}

func (v NullableTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
