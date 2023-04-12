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

// checks if the GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes{}

// GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes struct for GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes
type GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes struct {
	// The promotion's internal name.
	Name interface{} `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The activation date/time of this promotion.
	StartsAt interface{} `json:"starts_at,omitempty"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// The total number of times this promotion can be applied.
	TotalUsageLimit interface{} `json:"total_usage_limit,omitempty"`
	// The number of times this promotion has been applied.
	TotalUsageCount interface{} `json:"total_usage_count,omitempty"`
	// Indicates if the promotion is active.
	Active interface{} `json:"active,omitempty"`
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
	// The price fixed amount to be applied on matching SKUs, in cents
	FixedAmountCents interface{} `json:"fixed_amount_cents,omitempty"`
	// The discount fixed amount to be applied, float.
	FixedAmountFloat interface{} `json:"fixed_amount_float,omitempty"`
	// The discount fixed amount to be applied, formatted.
	FormattedFixedAmount interface{} `json:"formatted_fixed_amount,omitempty"`
}

// NewGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes instantiates a new GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes() *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes {
	this := GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes{}
	return &this
}

// NewGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributesWithDefaults instantiates a new GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributesWithDefaults() *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes {
	this := GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageLimit) {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && IsNil(o.TotalUsageLimit) {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetTotalUsageCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetTotalUsageCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageCount) {
		return nil, false
	}
	return &o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasTotalUsageCount() bool {
	if o != nil && IsNil(o.TotalUsageCount) {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given interface{} and assigns it to the TotalUsageCount field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetTotalUsageCount(v interface{}) {
	o.TotalUsageCount = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasActive() bool {
	if o != nil && IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetFixedAmountCents returns the FixedAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetFixedAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FixedAmountCents
}

// GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetFixedAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FixedAmountCents) {
		return nil, false
	}
	return &o.FixedAmountCents, true
}

// HasFixedAmountCents returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasFixedAmountCents() bool {
	if o != nil && IsNil(o.FixedAmountCents) {
		return true
	}

	return false
}

// SetFixedAmountCents gets a reference to the given interface{} and assigns it to the FixedAmountCents field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetFixedAmountCents(v interface{}) {
	o.FixedAmountCents = v
}

// GetFixedAmountFloat returns the FixedAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetFixedAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FixedAmountFloat
}

// GetFixedAmountFloatOk returns a tuple with the FixedAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetFixedAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FixedAmountFloat) {
		return nil, false
	}
	return &o.FixedAmountFloat, true
}

// HasFixedAmountFloat returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasFixedAmountFloat() bool {
	if o != nil && IsNil(o.FixedAmountFloat) {
		return true
	}

	return false
}

// SetFixedAmountFloat gets a reference to the given interface{} and assigns it to the FixedAmountFloat field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetFixedAmountFloat(v interface{}) {
	o.FixedAmountFloat = v
}

// GetFormattedFixedAmount returns the FormattedFixedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetFormattedFixedAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedFixedAmount
}

// GetFormattedFixedAmountOk returns a tuple with the FormattedFixedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) GetFormattedFixedAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedFixedAmount) {
		return nil, false
	}
	return &o.FormattedFixedAmount, true
}

// HasFormattedFixedAmount returns a boolean if a field has been set.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) HasFormattedFixedAmount() bool {
	if o != nil && IsNil(o.FormattedFixedAmount) {
		return true
	}

	return false
}

// SetFormattedFixedAmount gets a reference to the given interface{} and assigns it to the FormattedFixedAmount field.
func (o *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) SetFormattedFixedAmount(v interface{}) {
	o.FormattedFixedAmount = v
}

func (o GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes struct {
	value *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) Get() *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) Set(val *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes(val *GETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) *NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes {
	return &NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
