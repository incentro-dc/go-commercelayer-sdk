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

// TaxCategoryDataRelationships struct for TaxCategoryDataRelationships
type TaxCategoryDataRelationships struct {
	Sku           *BundleDataRelationshipsSkus                `json:"sku,omitempty"`
	TaxCalculator *TaxCategoryDataRelationshipsTaxCalculator  `json:"tax_calculator,omitempty"`
	Attachments   *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewTaxCategoryDataRelationships instantiates a new TaxCategoryDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryDataRelationships() *TaxCategoryDataRelationships {
	this := TaxCategoryDataRelationships{}
	return &this
}

// NewTaxCategoryDataRelationshipsWithDefaults instantiates a new TaxCategoryDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryDataRelationshipsWithDefaults() *TaxCategoryDataRelationships {
	this := TaxCategoryDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || o.Sku == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *TaxCategoryDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetTaxCalculator() TaxCategoryDataRelationshipsTaxCalculator {
	if o == nil || o.TaxCalculator == nil {
		var ret TaxCategoryDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetTaxCalculatorOk() (*TaxCategoryDataRelationshipsTaxCalculator, bool) {
	if o == nil || o.TaxCalculator == nil {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasTaxCalculator() bool {
	if o != nil && o.TaxCalculator != nil {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given TaxCategoryDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *TaxCategoryDataRelationships) SetTaxCalculator(v TaxCategoryDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *TaxCategoryDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *TaxCategoryDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *TaxCategoryDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o TaxCategoryDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.TaxCalculator != nil {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableTaxCategoryDataRelationships struct {
	value *TaxCategoryDataRelationships
	isSet bool
}

func (v NullableTaxCategoryDataRelationships) Get() *TaxCategoryDataRelationships {
	return v.value
}

func (v *NullableTaxCategoryDataRelationships) Set(val *TaxCategoryDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryDataRelationships(val *TaxCategoryDataRelationships) *NullableTaxCategoryDataRelationships {
	return &NullableTaxCategoryDataRelationships{value: val, isSet: true}
}

func (v NullableTaxCategoryDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
