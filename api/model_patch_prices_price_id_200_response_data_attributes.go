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

// PATCHPricesPriceId200ResponseDataAttributes struct for PATCHPricesPriceId200ResponseDataAttributes
type PATCHPricesPriceId200ResponseDataAttributes struct {
	// The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present.
	SkuCode *string `json:"sku_code,omitempty"`
	// The SKU price amount for the associated price list, in cents.
	AmountCents *int32 `json:"amount_cents,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents *int32 `json:"compare_at_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHPricesPriceId200ResponseDataAttributes instantiates a new PATCHPricesPriceId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPricesPriceId200ResponseDataAttributes() *PATCHPricesPriceId200ResponseDataAttributes {
	this := PATCHPricesPriceId200ResponseDataAttributes{}
	return &this
}

// NewPATCHPricesPriceId200ResponseDataAttributesWithDefaults instantiates a new PATCHPricesPriceId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPricesPriceId200ResponseDataAttributesWithDefaults() *PATCHPricesPriceId200ResponseDataAttributes {
	this := PATCHPricesPriceId200ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *PATCHPricesPriceId200ResponseDataAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetAmountCents() int32 {
	if o == nil || o.AmountCents == nil {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetAmountCentsOk() (*int32, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *PATCHPricesPriceId200ResponseDataAttributes) SetAmountCents(v int32) {
	o.AmountCents = &v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetCompareAtAmountCents() int32 {
	if o == nil || o.CompareAtAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*int32, bool) {
	if o == nil || o.CompareAtAmountCents == nil {
		return nil, false
	}
	return o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) HasCompareAtAmountCents() bool {
	if o != nil && o.CompareAtAmountCents != nil {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given int32 and assigns it to the CompareAtAmountCents field.
func (o *PATCHPricesPriceId200ResponseDataAttributes) SetCompareAtAmountCents(v int32) {
	o.CompareAtAmountCents = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHPricesPriceId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHPricesPriceId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHPricesPriceId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHPricesPriceId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
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

type NullablePATCHPricesPriceId200ResponseDataAttributes struct {
	value *PATCHPricesPriceId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHPricesPriceId200ResponseDataAttributes) Get() *PATCHPricesPriceId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHPricesPriceId200ResponseDataAttributes) Set(val *PATCHPricesPriceId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPricesPriceId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPricesPriceId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPricesPriceId200ResponseDataAttributes(val *PATCHPricesPriceId200ResponseDataAttributes) *NullablePATCHPricesPriceId200ResponseDataAttributes {
	return &NullablePATCHPricesPriceId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPricesPriceId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPricesPriceId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
