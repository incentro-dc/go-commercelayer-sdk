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

// GETFreeShippingPromotions200ResponseDataInnerAttributes struct for GETFreeShippingPromotions200ResponseDataInnerAttributes
type GETFreeShippingPromotions200ResponseDataInnerAttributes struct {
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
}

// NewGETFreeShippingPromotions200ResponseDataInnerAttributes instantiates a new GETFreeShippingPromotions200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFreeShippingPromotions200ResponseDataInnerAttributes() *GETFreeShippingPromotions200ResponseDataInnerAttributes {
	this := GETFreeShippingPromotions200ResponseDataInnerAttributes{}
	return &this
}

// NewGETFreeShippingPromotions200ResponseDataInnerAttributesWithDefaults instantiates a new GETFreeShippingPromotions200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFreeShippingPromotions200ResponseDataInnerAttributesWithDefaults() *GETFreeShippingPromotions200ResponseDataInnerAttributes {
	this := GETFreeShippingPromotions200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || o.TotalUsageLimit == nil {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasTotalUsageLimit() bool {
	if o != nil && o.TotalUsageLimit != nil {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given interface{} and assigns it to the TotalUsageLimit field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetTotalUsageCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetTotalUsageCountOk() (*interface{}, bool) {
	if o == nil || o.TotalUsageCount == nil {
		return nil, false
	}
	return &o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasTotalUsageCount() bool {
	if o != nil && o.TotalUsageCount != nil {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given interface{} and assigns it to the TotalUsageCount field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetTotalUsageCount(v interface{}) {
	o.TotalUsageCount = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETFreeShippingPromotions200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETFreeShippingPromotions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableGETFreeShippingPromotions200ResponseDataInnerAttributes struct {
	value *GETFreeShippingPromotions200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETFreeShippingPromotions200ResponseDataInnerAttributes) Get() *GETFreeShippingPromotions200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETFreeShippingPromotions200ResponseDataInnerAttributes) Set(val *GETFreeShippingPromotions200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFreeShippingPromotions200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFreeShippingPromotions200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFreeShippingPromotions200ResponseDataInnerAttributes(val *GETFreeShippingPromotions200ResponseDataInnerAttributes) *NullableGETFreeShippingPromotions200ResponseDataInnerAttributes {
	return &NullableGETFreeShippingPromotions200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETFreeShippingPromotions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFreeShippingPromotions200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
