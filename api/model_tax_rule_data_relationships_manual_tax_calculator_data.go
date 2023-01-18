/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TaxRuleDataRelationshipsManualTaxCalculatorData struct for TaxRuleDataRelationshipsManualTaxCalculatorData
type TaxRuleDataRelationshipsManualTaxCalculatorData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource's id
	Id *string `json:"id,omitempty"`
}

// NewTaxRuleDataRelationshipsManualTaxCalculatorData instantiates a new TaxRuleDataRelationshipsManualTaxCalculatorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleDataRelationshipsManualTaxCalculatorData() *TaxRuleDataRelationshipsManualTaxCalculatorData {
	this := TaxRuleDataRelationshipsManualTaxCalculatorData{}
	return &this
}

// NewTaxRuleDataRelationshipsManualTaxCalculatorDataWithDefaults instantiates a new TaxRuleDataRelationshipsManualTaxCalculatorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleDataRelationshipsManualTaxCalculatorDataWithDefaults() *TaxRuleDataRelationshipsManualTaxCalculatorData {
	this := TaxRuleDataRelationshipsManualTaxCalculatorData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaxRuleDataRelationshipsManualTaxCalculatorData) SetId(v string) {
	o.Id = &v
}

func (o TaxRuleDataRelationshipsManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableTaxRuleDataRelationshipsManualTaxCalculatorData struct {
	value *TaxRuleDataRelationshipsManualTaxCalculatorData
	isSet bool
}

func (v NullableTaxRuleDataRelationshipsManualTaxCalculatorData) Get() *TaxRuleDataRelationshipsManualTaxCalculatorData {
	return v.value
}

func (v *NullableTaxRuleDataRelationshipsManualTaxCalculatorData) Set(val *TaxRuleDataRelationshipsManualTaxCalculatorData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleDataRelationshipsManualTaxCalculatorData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleDataRelationshipsManualTaxCalculatorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleDataRelationshipsManualTaxCalculatorData(val *TaxRuleDataRelationshipsManualTaxCalculatorData) *NullableTaxRuleDataRelationshipsManualTaxCalculatorData {
	return &NullableTaxRuleDataRelationshipsManualTaxCalculatorData{value: val, isSet: true}
}

func (v NullableTaxRuleDataRelationshipsManualTaxCalculatorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleDataRelationshipsManualTaxCalculatorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
