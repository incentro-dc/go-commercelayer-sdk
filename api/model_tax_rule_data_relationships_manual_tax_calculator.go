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

// TaxRuleDataRelationshipsManualTaxCalculator struct for TaxRuleDataRelationshipsManualTaxCalculator
type TaxRuleDataRelationshipsManualTaxCalculator struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewTaxRuleDataRelationshipsManualTaxCalculator instantiates a new TaxRuleDataRelationshipsManualTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleDataRelationshipsManualTaxCalculator(type_ string, id string) *TaxRuleDataRelationshipsManualTaxCalculator {
	this := TaxRuleDataRelationshipsManualTaxCalculator{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewTaxRuleDataRelationshipsManualTaxCalculatorWithDefaults instantiates a new TaxRuleDataRelationshipsManualTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleDataRelationshipsManualTaxCalculatorWithDefaults() *TaxRuleDataRelationshipsManualTaxCalculator {
	this := TaxRuleDataRelationshipsManualTaxCalculator{}
	var type_ string = "manual_tax_calculators"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TaxRuleDataRelationshipsManualTaxCalculator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationshipsManualTaxCalculator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxRuleDataRelationshipsManualTaxCalculator) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TaxRuleDataRelationshipsManualTaxCalculator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationshipsManualTaxCalculator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxRuleDataRelationshipsManualTaxCalculator) SetId(v string) {
	o.Id = v
}

func (o TaxRuleDataRelationshipsManualTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTaxRuleDataRelationshipsManualTaxCalculator struct {
	value *TaxRuleDataRelationshipsManualTaxCalculator
	isSet bool
}

func (v NullableTaxRuleDataRelationshipsManualTaxCalculator) Get() *TaxRuleDataRelationshipsManualTaxCalculator {
	return v.value
}

func (v *NullableTaxRuleDataRelationshipsManualTaxCalculator) Set(val *TaxRuleDataRelationshipsManualTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleDataRelationshipsManualTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleDataRelationshipsManualTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleDataRelationshipsManualTaxCalculator(val *TaxRuleDataRelationshipsManualTaxCalculator) *NullableTaxRuleDataRelationshipsManualTaxCalculator {
	return &NullableTaxRuleDataRelationshipsManualTaxCalculator{value: val, isSet: true}
}

func (v NullableTaxRuleDataRelationshipsManualTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleDataRelationshipsManualTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
