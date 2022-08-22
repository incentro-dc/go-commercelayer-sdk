/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTTaxRules201ResponseDataRelationships struct for POSTTaxRules201ResponseDataRelationships
type POSTTaxRules201ResponseDataRelationships struct {
	ManualTaxCalculator GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculator `json:"manual_tax_calculator"`
}

// NewPOSTTaxRules201ResponseDataRelationships instantiates a new POSTTaxRules201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTTaxRules201ResponseDataRelationships(manualTaxCalculator GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculator) *POSTTaxRules201ResponseDataRelationships {
	this := POSTTaxRules201ResponseDataRelationships{}
	this.ManualTaxCalculator = manualTaxCalculator
	return &this
}

// NewPOSTTaxRules201ResponseDataRelationshipsWithDefaults instantiates a new POSTTaxRules201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTTaxRules201ResponseDataRelationshipsWithDefaults() *POSTTaxRules201ResponseDataRelationships {
	this := POSTTaxRules201ResponseDataRelationships{}
	return &this
}

// GetManualTaxCalculator returns the ManualTaxCalculator field value
func (o *POSTTaxRules201ResponseDataRelationships) GetManualTaxCalculator() GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculator {
	if o == nil {
		var ret GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculator
		return ret
	}

	return o.ManualTaxCalculator
}

// GetManualTaxCalculatorOk returns a tuple with the ManualTaxCalculator field value
// and a boolean to check if the value has been set.
func (o *POSTTaxRules201ResponseDataRelationships) GetManualTaxCalculatorOk() (*GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualTaxCalculator, true
}

// SetManualTaxCalculator sets field value
func (o *POSTTaxRules201ResponseDataRelationships) SetManualTaxCalculator(v GETTaxRules200ResponseDataInnerRelationshipsManualTaxCalculator) {
	o.ManualTaxCalculator = v
}

func (o POSTTaxRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["manual_tax_calculator"] = o.ManualTaxCalculator
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTTaxRules201ResponseDataRelationships struct {
	value *POSTTaxRules201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTTaxRules201ResponseDataRelationships) Get() *POSTTaxRules201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTTaxRules201ResponseDataRelationships) Set(val *POSTTaxRules201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTTaxRules201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTTaxRules201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTTaxRules201ResponseDataRelationships(val *POSTTaxRules201ResponseDataRelationships) *NullablePOSTTaxRules201ResponseDataRelationships {
	return &NullablePOSTTaxRules201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTTaxRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTTaxRules201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}