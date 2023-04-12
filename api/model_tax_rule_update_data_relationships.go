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

// checks if the TaxRuleUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRuleUpdateDataRelationships{}

// TaxRuleUpdateDataRelationships struct for TaxRuleUpdateDataRelationships
type TaxRuleUpdateDataRelationships struct {
	ManualTaxCalculator *TaxRuleCreateDataRelationshipsManualTaxCalculator `json:"manual_tax_calculator,omitempty"`
}

// NewTaxRuleUpdateDataRelationships instantiates a new TaxRuleUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleUpdateDataRelationships() *TaxRuleUpdateDataRelationships {
	this := TaxRuleUpdateDataRelationships{}
	return &this
}

// NewTaxRuleUpdateDataRelationshipsWithDefaults instantiates a new TaxRuleUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleUpdateDataRelationshipsWithDefaults() *TaxRuleUpdateDataRelationships {
	this := TaxRuleUpdateDataRelationships{}
	return &this
}

// GetManualTaxCalculator returns the ManualTaxCalculator field value if set, zero value otherwise.
func (o *TaxRuleUpdateDataRelationships) GetManualTaxCalculator() TaxRuleCreateDataRelationshipsManualTaxCalculator {
	if o == nil || IsNil(o.ManualTaxCalculator) {
		var ret TaxRuleCreateDataRelationshipsManualTaxCalculator
		return ret
	}
	return *o.ManualTaxCalculator
}

// GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleUpdateDataRelationships) GetManualTaxCalculatorOk() (*TaxRuleCreateDataRelationshipsManualTaxCalculator, bool) {
	if o == nil || IsNil(o.ManualTaxCalculator) {
		return nil, false
	}
	return o.ManualTaxCalculator, true
}

// HasManualTaxCalculator returns a boolean if a field has been set.
func (o *TaxRuleUpdateDataRelationships) HasManualTaxCalculator() bool {
	if o != nil && !IsNil(o.ManualTaxCalculator) {
		return true
	}

	return false
}

// SetManualTaxCalculator gets a reference to the given TaxRuleCreateDataRelationshipsManualTaxCalculator and assigns it to the ManualTaxCalculator field.
func (o *TaxRuleUpdateDataRelationships) SetManualTaxCalculator(v TaxRuleCreateDataRelationshipsManualTaxCalculator) {
	o.ManualTaxCalculator = &v
}

func (o TaxRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRuleUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManualTaxCalculator) {
		toSerialize["manual_tax_calculator"] = o.ManualTaxCalculator
	}
	return toSerialize, nil
}

type NullableTaxRuleUpdateDataRelationships struct {
	value *TaxRuleUpdateDataRelationships
	isSet bool
}

func (v NullableTaxRuleUpdateDataRelationships) Get() *TaxRuleUpdateDataRelationships {
	return v.value
}

func (v *NullableTaxRuleUpdateDataRelationships) Set(val *TaxRuleUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleUpdateDataRelationships(val *TaxRuleUpdateDataRelationships) *NullableTaxRuleUpdateDataRelationships {
	return &NullableTaxRuleUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableTaxRuleUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}