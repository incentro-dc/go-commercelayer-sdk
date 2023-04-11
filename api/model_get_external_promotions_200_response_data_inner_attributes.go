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

// GETExternalPromotions200ResponseDataInnerAttributes struct for GETExternalPromotions200ResponseDataInnerAttributes
type GETExternalPromotions200ResponseDataInnerAttributes struct {
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
	// The URL to the service that will compute the discount.
	PromotionUrl interface{} `json:"promotion_url,omitempty"`
	// The shared secret used to sign the external request payload.
	SharedSecret interface{} `json:"shared_secret,omitempty"`
}

// NewGETExternalPromotions200ResponseDataInnerAttributes instantiates a new GETExternalPromotions200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200ResponseDataInnerAttributes() *GETExternalPromotions200ResponseDataInnerAttributes {
	this := GETExternalPromotions200ResponseDataInnerAttributes{}
	return &this
}

// NewGETExternalPromotions200ResponseDataInnerAttributesWithDefaults instantiates a new GETExternalPromotions200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseDataInnerAttributesWithDefaults() *GETExternalPromotions200ResponseDataInnerAttributes {
	this := GETExternalPromotions200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || o.TotalUsageLimit == nil {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasTotalUsageLimit() bool {
	if o != nil && o.TotalUsageLimit != nil {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetTotalUsageCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetTotalUsageCountOk() (*interface{}, bool) {
	if o == nil || o.TotalUsageCount == nil {
		return nil, false
	}
	return &o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasTotalUsageCount() bool {
	if o != nil && o.TotalUsageCount != nil {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given interface{} and assigns it to the TotalUsageCount field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetTotalUsageCount(v interface{}) {
	o.TotalUsageCount = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetPromotionUrl returns the PromotionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetPromotionUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PromotionUrl
}

// GetPromotionUrlOk returns a tuple with the PromotionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetPromotionUrlOk() (*interface{}, bool) {
	if o == nil || o.PromotionUrl == nil {
		return nil, false
	}
	return &o.PromotionUrl, true
}

// HasPromotionUrl returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasPromotionUrl() bool {
	if o != nil && o.PromotionUrl != nil {
		return true
	}

	return false
}

// SetPromotionUrl gets a reference to the given interface{} and assigns it to the PromotionUrl field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetPromotionUrl(v interface{}) {
	o.PromotionUrl = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetSharedSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPromotions200ResponseDataInnerAttributes) GetSharedSecretOk() (*interface{}, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return &o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given interface{} and assigns it to the SharedSecret field.
func (o *GETExternalPromotions200ResponseDataInnerAttributes) SetSharedSecret(v interface{}) {
	o.SharedSecret = v
}

func (o GETExternalPromotions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
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
	if o.PromotionUrl != nil {
		toSerialize["promotion_url"] = o.PromotionUrl
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200ResponseDataInnerAttributes struct {
	value *GETExternalPromotions200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETExternalPromotions200ResponseDataInnerAttributes) Get() *GETExternalPromotions200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETExternalPromotions200ResponseDataInnerAttributes) Set(val *GETExternalPromotions200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200ResponseDataInnerAttributes(val *GETExternalPromotions200ResponseDataInnerAttributes) *NullableGETExternalPromotions200ResponseDataInnerAttributes {
	return &NullableGETExternalPromotions200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
