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

// TaxRuleDataRelationships struct for TaxRuleDataRelationships
type TaxRuleDataRelationships struct {
	ManualTaxCalculator *TaxRuleDataRelationshipsManualTaxCalculator `json:"manual_tax_calculator,omitempty"`
}

// NewTaxRuleDataRelationships instantiates a new TaxRuleDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleDataRelationships() *TaxRuleDataRelationships {
	this := TaxRuleDataRelationships{}
	return &this
}

// NewTaxRuleDataRelationshipsWithDefaults instantiates a new TaxRuleDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleDataRelationshipsWithDefaults() *TaxRuleDataRelationships {
	this := TaxRuleDataRelationships{}
	return &this
}

// GetManualTaxCalculator returns the ManualTaxCalculator field value if set, zero value otherwise.
func (o *TaxRuleDataRelationships) GetManualTaxCalculator() TaxRuleDataRelationshipsManualTaxCalculator {
	if o == nil || o.ManualTaxCalculator == nil {
		var ret TaxRuleDataRelationshipsManualTaxCalculator
		return ret
	}
	return *o.ManualTaxCalculator
}

// GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationships) GetManualTaxCalculatorOk() (*TaxRuleDataRelationshipsManualTaxCalculator, bool) {
	if o == nil || o.ManualTaxCalculator == nil {
		return nil, false
	}
	return o.ManualTaxCalculator, true
}

// HasManualTaxCalculator returns a boolean if a field has been set.
func (o *TaxRuleDataRelationships) HasManualTaxCalculator() bool {
	if o != nil && o.ManualTaxCalculator != nil {
		return true
	}

	return false
}

// SetManualTaxCalculator gets a reference to the given TaxRuleDataRelationshipsManualTaxCalculator and assigns it to the ManualTaxCalculator field.
func (o *TaxRuleDataRelationships) SetManualTaxCalculator(v TaxRuleDataRelationshipsManualTaxCalculator) {
	o.ManualTaxCalculator = &v
}

func (o TaxRuleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ManualTaxCalculator != nil {
		toSerialize["manual_tax_calculator"] = o.ManualTaxCalculator
	}
	return json.Marshal(toSerialize)
}

type NullableTaxRuleDataRelationships struct {
	value *TaxRuleDataRelationships
	isSet bool
}

func (v NullableTaxRuleDataRelationships) Get() *TaxRuleDataRelationships {
	return v.value
}

func (v *NullableTaxRuleDataRelationships) Set(val *TaxRuleDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleDataRelationships(val *TaxRuleDataRelationships) *NullableTaxRuleDataRelationships {
	return &NullableTaxRuleDataRelationships{value: val, isSet: true}
}

func (v NullableTaxRuleDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
