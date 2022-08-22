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

// PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes struct for PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes
type PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes struct {
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
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The URL to the service that will compute the discount.
	PromotionUrl *string `json:"promotion_url,omitempty"`
}

// NewPATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes instantiates a new PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes() *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	this := PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHExternalPromotionsExternalPromotionId200ResponseDataAttributesWithDefaults instantiates a new PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalPromotionsExternalPromotionId200ResponseDataAttributesWithDefaults() *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	this := PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetStartsAt() string {
	if o == nil || o.StartsAt == nil {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetStartsAtOk() (*string, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetStartsAt(v string) {
	o.StartsAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTotalUsageLimit() int32 {
	if o == nil || o.TotalUsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*int32, bool) {
	if o == nil || o.TotalUsageLimit == nil {
		return nil, false
	}
	return o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && o.TotalUsageLimit != nil {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given int32 and assigns it to the TotalUsageLimit field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v int32) {
	o.TotalUsageLimit = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPromotionUrl returns the PromotionUrl field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetPromotionUrl() string {
	if o == nil || o.PromotionUrl == nil {
		var ret string
		return ret
	}
	return *o.PromotionUrl
}

// GetPromotionUrlOk returns a tuple with the PromotionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) GetPromotionUrlOk() (*string, bool) {
	if o == nil || o.PromotionUrl == nil {
		return nil, false
	}
	return o.PromotionUrl, true
}

// HasPromotionUrl returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) HasPromotionUrl() bool {
	if o != nil && o.PromotionUrl != nil {
		return true
	}

	return false
}

// SetPromotionUrl gets a reference to the given string and assigns it to the PromotionUrl field.
func (o *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) SetPromotionUrl(v string) {
	o.PromotionUrl = &v
}

func (o PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes struct {
	value *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) Get() *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) Set(val *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes(val *PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) *NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes {
	return &NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
