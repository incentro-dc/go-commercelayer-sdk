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

// GETPrices200ResponseDataInnerAttributes struct for GETPrices200ResponseDataInnerAttributes
type GETPrices200ResponseDataInnerAttributes struct {
	// The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated price list.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The SKU price amount for the associated price list, in cents.
	AmountCents interface{} `json:"amount_cents,omitempty"`
	// The SKU price amount for the associated price list, float.
	AmountFloat interface{} `json:"amount_float,omitempty"`
	// The SKU price amount for the associated price list, formatted.
	FormattedAmount interface{} `json:"formatted_amount,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents interface{} `json:"compare_at_amount_cents,omitempty"`
	// The compared price amount, float.
	CompareAtAmountFloat interface{} `json:"compare_at_amount_float,omitempty"`
	// The compared price amount, formatted.
	FormattedCompareAtAmount interface{} `json:"formatted_compare_at_amount,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETPrices200ResponseDataInnerAttributes instantiates a new GETPrices200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPrices200ResponseDataInnerAttributes() *GETPrices200ResponseDataInnerAttributes {
	this := GETPrices200ResponseDataInnerAttributes{}
	return &this
}

// NewGETPrices200ResponseDataInnerAttributesWithDefaults instantiates a new GETPrices200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPrices200ResponseDataInnerAttributesWithDefaults() *GETPrices200ResponseDataInnerAttributes {
	this := GETPrices200ResponseDataInnerAttributes{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETPrices200ResponseDataInnerAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *GETPrices200ResponseDataInnerAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetAmountCentsOk() (*interface{}, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given interface{} and assigns it to the AmountCents field.
func (o *GETPrices200ResponseDataInnerAttributes) SetAmountCents(v interface{}) {
	o.AmountCents = v
}

// GetAmountFloat returns the AmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountFloat
}

// GetAmountFloatOk returns a tuple with the AmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetAmountFloatOk() (*interface{}, bool) {
	if o == nil || o.AmountFloat == nil {
		return nil, false
	}
	return &o.AmountFloat, true
}

// HasAmountFloat returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasAmountFloat() bool {
	if o != nil && o.AmountFloat != nil {
		return true
	}

	return false
}

// SetAmountFloat gets a reference to the given interface{} and assigns it to the AmountFloat field.
func (o *GETPrices200ResponseDataInnerAttributes) SetAmountFloat(v interface{}) {
	o.AmountFloat = v
}

// GetFormattedAmount returns the FormattedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedAmount
}

// GetFormattedAmountOk returns a tuple with the FormattedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedAmountOk() (*interface{}, bool) {
	if o == nil || o.FormattedAmount == nil {
		return nil, false
	}
	return &o.FormattedAmount, true
}

// HasFormattedAmount returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasFormattedAmount() bool {
	if o != nil && o.FormattedAmount != nil {
		return true
	}

	return false
}

// SetFormattedAmount gets a reference to the given interface{} and assigns it to the FormattedAmount field.
func (o *GETPrices200ResponseDataInnerAttributes) SetFormattedAmount(v interface{}) {
	o.FormattedAmount = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool) {
	if o == nil || o.CompareAtAmountCents == nil {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasCompareAtAmountCents() bool {
	if o != nil && o.CompareAtAmountCents != nil {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given interface{} and assigns it to the CompareAtAmountCents field.
func (o *GETPrices200ResponseDataInnerAttributes) SetCompareAtAmountCents(v interface{}) {
	o.CompareAtAmountCents = v
}

// GetCompareAtAmountFloat returns the CompareAtAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompareAtAmountFloat
}

// GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool) {
	if o == nil || o.CompareAtAmountFloat == nil {
		return nil, false
	}
	return &o.CompareAtAmountFloat, true
}

// HasCompareAtAmountFloat returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasCompareAtAmountFloat() bool {
	if o != nil && o.CompareAtAmountFloat != nil {
		return true
	}

	return false
}

// SetCompareAtAmountFloat gets a reference to the given interface{} and assigns it to the CompareAtAmountFloat field.
func (o *GETPrices200ResponseDataInnerAttributes) SetCompareAtAmountFloat(v interface{}) {
	o.CompareAtAmountFloat = v
}

// GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedCompareAtAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedCompareAtAmount
}

// GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool) {
	if o == nil || o.FormattedCompareAtAmount == nil {
		return nil, false
	}
	return &o.FormattedCompareAtAmount, true
}

// HasFormattedCompareAtAmount returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasFormattedCompareAtAmount() bool {
	if o != nil && o.FormattedCompareAtAmount != nil {
		return true
	}

	return false
}

// SetFormattedCompareAtAmount gets a reference to the given interface{} and assigns it to the FormattedCompareAtAmount field.
func (o *GETPrices200ResponseDataInnerAttributes) SetFormattedCompareAtAmount(v interface{}) {
	o.FormattedCompareAtAmount = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETPrices200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETPrices200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETPrices200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETPrices200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPrices200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPrices200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETPrices200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETPrices200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETPrices200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.AmountFloat != nil {
		toSerialize["amount_float"] = o.AmountFloat
	}
	if o.FormattedAmount != nil {
		toSerialize["formatted_amount"] = o.FormattedAmount
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
	}
	if o.CompareAtAmountFloat != nil {
		toSerialize["compare_at_amount_float"] = o.CompareAtAmountFloat
	}
	if o.FormattedCompareAtAmount != nil {
		toSerialize["formatted_compare_at_amount"] = o.FormattedCompareAtAmount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETPrices200ResponseDataInnerAttributes struct {
	value *GETPrices200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETPrices200ResponseDataInnerAttributes) Get() *GETPrices200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETPrices200ResponseDataInnerAttributes) Set(val *GETPrices200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPrices200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPrices200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPrices200ResponseDataInnerAttributes(val *GETPrices200ResponseDataInnerAttributes) *NullableGETPrices200ResponseDataInnerAttributes {
	return &NullableGETPrices200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETPrices200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPrices200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
