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

// GETFixedPricePromotions200ResponseDataInnerAttributes struct for GETFixedPricePromotions200ResponseDataInnerAttributes
type GETFixedPricePromotions200ResponseDataInnerAttributes struct {
	// The promotion's internal name.
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The activation date/time of this promotion.
	StartsAt *string `json:"starts_at,omitempty"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The total number of times this promotion can be applied.
	TotalUsageLimit *int32 `json:"total_usage_limit,omitempty"`
	// The number of times this promotion has been applied.
	TotalUsageCount *int32 `json:"total_usage_count,omitempty"`
	// Indicates if the promotion is active.
	Active *bool `json:"active,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The price fixed amount to be applied on matching SKUs, in cents
	FixedAmountCents *int32 `json:"fixed_amount_cents,omitempty"`
	// The discount fixed amount to be applied, float.
	FixedAmountFloat *float32 `json:"fixed_amount_float,omitempty"`
	// The discount fixed amount to be applied, formatted.
	FormattedFixedAmount *string `json:"formatted_fixed_amount,omitempty"`
}

// NewGETFixedPricePromotions200ResponseDataInnerAttributes instantiates a new GETFixedPricePromotions200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFixedPricePromotions200ResponseDataInnerAttributes() *GETFixedPricePromotions200ResponseDataInnerAttributes {
	this := GETFixedPricePromotions200ResponseDataInnerAttributes{}
	return &this
}

// NewGETFixedPricePromotions200ResponseDataInnerAttributesWithDefaults instantiates a new GETFixedPricePromotions200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFixedPricePromotions200ResponseDataInnerAttributesWithDefaults() *GETFixedPricePromotions200ResponseDataInnerAttributes {
	this := GETFixedPricePromotions200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetStartsAt() string {
	if o == nil || o.StartsAt == nil {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetStartsAtOk() (*string, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetStartsAt(v string) {
	o.StartsAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetTotalUsageLimit() int32 {
	if o == nil || o.TotalUsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetTotalUsageLimitOk() (*int32, bool) {
	if o == nil || o.TotalUsageLimit == nil {
		return nil, false
	}
	return o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasTotalUsageLimit() bool {
	if o != nil && o.TotalUsageLimit != nil {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given int32 and assigns it to the TotalUsageLimit field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetTotalUsageLimit(v int32) {
	o.TotalUsageLimit = &v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetTotalUsageCount() int32 {
	if o == nil || o.TotalUsageCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetTotalUsageCountOk() (*int32, bool) {
	if o == nil || o.TotalUsageCount == nil {
		return nil, false
	}
	return o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasTotalUsageCount() bool {
	if o != nil && o.TotalUsageCount != nil {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given int32 and assigns it to the TotalUsageCount field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetTotalUsageCount(v int32) {
	o.TotalUsageCount = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetFixedAmountCents returns the FixedAmountCents field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetFixedAmountCents() int32 {
	if o == nil || o.FixedAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.FixedAmountCents
}

// GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetFixedAmountCentsOk() (*int32, bool) {
	if o == nil || o.FixedAmountCents == nil {
		return nil, false
	}
	return o.FixedAmountCents, true
}

// HasFixedAmountCents returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasFixedAmountCents() bool {
	if o != nil && o.FixedAmountCents != nil {
		return true
	}

	return false
}

// SetFixedAmountCents gets a reference to the given int32 and assigns it to the FixedAmountCents field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetFixedAmountCents(v int32) {
	o.FixedAmountCents = &v
}

// GetFixedAmountFloat returns the FixedAmountFloat field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetFixedAmountFloat() float32 {
	if o == nil || o.FixedAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.FixedAmountFloat
}

// GetFixedAmountFloatOk returns a tuple with the FixedAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetFixedAmountFloatOk() (*float32, bool) {
	if o == nil || o.FixedAmountFloat == nil {
		return nil, false
	}
	return o.FixedAmountFloat, true
}

// HasFixedAmountFloat returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasFixedAmountFloat() bool {
	if o != nil && o.FixedAmountFloat != nil {
		return true
	}

	return false
}

// SetFixedAmountFloat gets a reference to the given float32 and assigns it to the FixedAmountFloat field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetFixedAmountFloat(v float32) {
	o.FixedAmountFloat = &v
}

// GetFormattedFixedAmount returns the FormattedFixedAmount field value if set, zero value otherwise.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetFormattedFixedAmount() string {
	if o == nil || o.FormattedFixedAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedFixedAmount
}

// GetFormattedFixedAmountOk returns a tuple with the FormattedFixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) GetFormattedFixedAmountOk() (*string, bool) {
	if o == nil || o.FormattedFixedAmount == nil {
		return nil, false
	}
	return o.FormattedFixedAmount, true
}

// HasFormattedFixedAmount returns a boolean if a field has been set.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) HasFormattedFixedAmount() bool {
	if o != nil && o.FormattedFixedAmount != nil {
		return true
	}

	return false
}

// SetFormattedFixedAmount gets a reference to the given string and assigns it to the FormattedFixedAmount field.
func (o *GETFixedPricePromotions200ResponseDataInnerAttributes) SetFormattedFixedAmount(v string) {
	o.FormattedFixedAmount = &v
}

func (o GETFixedPricePromotions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.TotalUsageLimit != nil {
		toSerialize["total_usage_limit"] = o.TotalUsageLimit
	}
	if o.TotalUsageCount != nil {
		toSerialize["total_usage_count"] = o.TotalUsageCount
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
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
	if o.FixedAmountCents != nil {
		toSerialize["fixed_amount_cents"] = o.FixedAmountCents
	}
	if o.FixedAmountFloat != nil {
		toSerialize["fixed_amount_float"] = o.FixedAmountFloat
	}
	if o.FormattedFixedAmount != nil {
		toSerialize["formatted_fixed_amount"] = o.FormattedFixedAmount
	}
	return json.Marshal(toSerialize)
}

type NullableGETFixedPricePromotions200ResponseDataInnerAttributes struct {
	value *GETFixedPricePromotions200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETFixedPricePromotions200ResponseDataInnerAttributes) Get() *GETFixedPricePromotions200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETFixedPricePromotions200ResponseDataInnerAttributes) Set(val *GETFixedPricePromotions200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFixedPricePromotions200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFixedPricePromotions200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFixedPricePromotions200ResponseDataInnerAttributes(val *GETFixedPricePromotions200ResponseDataInnerAttributes) *NullableGETFixedPricePromotions200ResponseDataInnerAttributes {
	return &NullableGETFixedPricePromotions200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETFixedPricePromotions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFixedPricePromotions200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
