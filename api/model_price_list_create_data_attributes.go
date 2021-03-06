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

// PriceListCreateDataAttributes struct for PriceListCreateDataAttributes
type PriceListCreateDataAttributes struct {
	// The price list's internal name
	Name string `json:"name"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode string `json:"currency_code"`
	// Indicates if the associated prices include taxes.
	TaxIncluded *bool `json:"tax_included,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPriceListCreateDataAttributes instantiates a new PriceListCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListCreateDataAttributes(name string, currencyCode string) *PriceListCreateDataAttributes {
	this := PriceListCreateDataAttributes{}
	this.Name = name
	this.CurrencyCode = currencyCode
	return &this
}

// NewPriceListCreateDataAttributesWithDefaults instantiates a new PriceListCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListCreateDataAttributesWithDefaults() *PriceListCreateDataAttributes {
	this := PriceListCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *PriceListCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PriceListCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PriceListCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *PriceListCreateDataAttributes) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *PriceListCreateDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *PriceListCreateDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetTaxIncluded returns the TaxIncluded field value if set, zero value otherwise.
func (o *PriceListCreateDataAttributes) GetTaxIncluded() bool {
	if o == nil || o.TaxIncluded == nil {
		var ret bool
		return ret
	}
	return *o.TaxIncluded
}

// GetTaxIncludedOk returns a tuple with the TaxIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListCreateDataAttributes) GetTaxIncludedOk() (*bool, bool) {
	if o == nil || o.TaxIncluded == nil {
		return nil, false
	}
	return o.TaxIncluded, true
}

// HasTaxIncluded returns a boolean if a field has been set.
func (o *PriceListCreateDataAttributes) HasTaxIncluded() bool {
	if o != nil && o.TaxIncluded != nil {
		return true
	}

	return false
}

// SetTaxIncluded gets a reference to the given bool and assigns it to the TaxIncluded field.
func (o *PriceListCreateDataAttributes) SetTaxIncluded(v bool) {
	o.TaxIncluded = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PriceListCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PriceListCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PriceListCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PriceListCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PriceListCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PriceListCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PriceListCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PriceListCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PriceListCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PriceListCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.TaxIncluded != nil {
		toSerialize["tax_included"] = o.TaxIncluded
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePriceListCreateDataAttributes struct {
	value *PriceListCreateDataAttributes
	isSet bool
}

func (v NullablePriceListCreateDataAttributes) Get() *PriceListCreateDataAttributes {
	return v.value
}

func (v *NullablePriceListCreateDataAttributes) Set(val *PriceListCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListCreateDataAttributes(val *PriceListCreateDataAttributes) *NullablePriceListCreateDataAttributes {
	return &NullablePriceListCreateDataAttributes{value: val, isSet: true}
}

func (v NullablePriceListCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
