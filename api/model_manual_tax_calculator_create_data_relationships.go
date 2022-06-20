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

// ManualTaxCalculatorCreateDataRelationships struct for ManualTaxCalculatorCreateDataRelationships
type ManualTaxCalculatorCreateDataRelationships struct {
	TaxCategories *AvalaraAccountDataRelationshipsTaxCategories `json:"tax_categories,omitempty"`
	TaxRules      *ManualTaxCalculatorDataRelationshipsTaxRules `json:"tax_rules,omitempty"`
}

// NewManualTaxCalculatorCreateDataRelationships instantiates a new ManualTaxCalculatorCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorCreateDataRelationships() *ManualTaxCalculatorCreateDataRelationships {
	this := ManualTaxCalculatorCreateDataRelationships{}
	return &this
}

// NewManualTaxCalculatorCreateDataRelationshipsWithDefaults instantiates a new ManualTaxCalculatorCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorCreateDataRelationshipsWithDefaults() *ManualTaxCalculatorCreateDataRelationships {
	this := ManualTaxCalculatorCreateDataRelationships{}
	return &this
}

// GetTaxCategories returns the TaxCategories field value if set, zero value otherwise.
func (o *ManualTaxCalculatorCreateDataRelationships) GetTaxCategories() AvalaraAccountDataRelationshipsTaxCategories {
	if o == nil || o.TaxCategories == nil {
		var ret AvalaraAccountDataRelationshipsTaxCategories
		return ret
	}
	return *o.TaxCategories
}

// GetTaxCategoriesOk returns a tuple with the TaxCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateDataRelationships) GetTaxCategoriesOk() (*AvalaraAccountDataRelationshipsTaxCategories, bool) {
	if o == nil || o.TaxCategories == nil {
		return nil, false
	}
	return o.TaxCategories, true
}

// HasTaxCategories returns a boolean if a field has been set.
func (o *ManualTaxCalculatorCreateDataRelationships) HasTaxCategories() bool {
	if o != nil && o.TaxCategories != nil {
		return true
	}

	return false
}

// SetTaxCategories gets a reference to the given AvalaraAccountDataRelationshipsTaxCategories and assigns it to the TaxCategories field.
func (o *ManualTaxCalculatorCreateDataRelationships) SetTaxCategories(v AvalaraAccountDataRelationshipsTaxCategories) {
	o.TaxCategories = &v
}

// GetTaxRules returns the TaxRules field value if set, zero value otherwise.
func (o *ManualTaxCalculatorCreateDataRelationships) GetTaxRules() ManualTaxCalculatorDataRelationshipsTaxRules {
	if o == nil || o.TaxRules == nil {
		var ret ManualTaxCalculatorDataRelationshipsTaxRules
		return ret
	}
	return *o.TaxRules
}

// GetTaxRulesOk returns a tuple with the TaxRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorCreateDataRelationships) GetTaxRulesOk() (*ManualTaxCalculatorDataRelationshipsTaxRules, bool) {
	if o == nil || o.TaxRules == nil {
		return nil, false
	}
	return o.TaxRules, true
}

// HasTaxRules returns a boolean if a field has been set.
func (o *ManualTaxCalculatorCreateDataRelationships) HasTaxRules() bool {
	if o != nil && o.TaxRules != nil {
		return true
	}

	return false
}

// SetTaxRules gets a reference to the given ManualTaxCalculatorDataRelationshipsTaxRules and assigns it to the TaxRules field.
func (o *ManualTaxCalculatorCreateDataRelationships) SetTaxRules(v ManualTaxCalculatorDataRelationshipsTaxRules) {
	o.TaxRules = &v
}

func (o ManualTaxCalculatorCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxCategories != nil {
		toSerialize["tax_categories"] = o.TaxCategories
	}
	if o.TaxRules != nil {
		toSerialize["tax_rules"] = o.TaxRules
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculatorCreateDataRelationships struct {
	value *ManualTaxCalculatorCreateDataRelationships
	isSet bool
}

func (v NullableManualTaxCalculatorCreateDataRelationships) Get() *ManualTaxCalculatorCreateDataRelationships {
	return v.value
}

func (v *NullableManualTaxCalculatorCreateDataRelationships) Set(val *ManualTaxCalculatorCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorCreateDataRelationships(val *ManualTaxCalculatorCreateDataRelationships) *NullableManualTaxCalculatorCreateDataRelationships {
	return &NullableManualTaxCalculatorCreateDataRelationships{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
