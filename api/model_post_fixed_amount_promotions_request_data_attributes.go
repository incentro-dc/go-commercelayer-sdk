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

// checks if the POSTFixedAmountPromotionsRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTFixedAmountPromotionsRequestDataAttributes{}

// POSTFixedAmountPromotionsRequestDataAttributes struct for POSTFixedAmountPromotionsRequestDataAttributes
type POSTFixedAmountPromotionsRequestDataAttributes struct {
	// The promotion's internal name.
	Name interface{} `json:"name"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The activation date/time of this promotion.
	StartsAt interface{} `json:"starts_at"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at"`
	// The total number of times this promotion can be applied.
	TotalUsageLimit interface{} `json:"total_usage_limit"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The discount fixed amount to be applied, in cents
	FixedAmountCents interface{} `json:"fixed_amount_cents"`
}

// NewPOSTFixedAmountPromotionsRequestDataAttributes instantiates a new POSTFixedAmountPromotionsRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFixedAmountPromotionsRequestDataAttributes(name interface{}, startsAt interface{}, expiresAt interface{}, totalUsageLimit interface{}, fixedAmountCents interface{}) *POSTFixedAmountPromotionsRequestDataAttributes {
	this := POSTFixedAmountPromotionsRequestDataAttributes{}
	this.Name = name
	this.StartsAt = startsAt
	this.ExpiresAt = expiresAt
	this.TotalUsageLimit = totalUsageLimit
	this.FixedAmountCents = fixedAmountCents
	return &this
}

// NewPOSTFixedAmountPromotionsRequestDataAttributesWithDefaults instantiates a new POSTFixedAmountPromotionsRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFixedAmountPromotionsRequestDataAttributesWithDefaults() *POSTFixedAmountPromotionsRequestDataAttributes {
	this := POSTFixedAmountPromotionsRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetStartsAt returns the StartsAt field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetTotalUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalUsageLimit) {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// SetTotalUsageLimit sets field value
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetTotalUsageLimit(v interface{}) {
	o.TotalUsageLimit = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetFixedAmountCents returns the FixedAmountCents field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetFixedAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.FixedAmountCents
}

// GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTFixedAmountPromotionsRequestDataAttributes) GetFixedAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FixedAmountCents) {
		return nil, false
	}
	return &o.FixedAmountCents, true
}

// SetFixedAmountCents sets field value
func (o *POSTFixedAmountPromotionsRequestDataAttributes) SetFixedAmountCents(v interface{}) {
	o.FixedAmountCents = v
}

func (o POSTFixedAmountPromotionsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTFixedAmountPromotionsRequestDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.FixedAmountCents != nil {
		toSerialize["fixed_amount_cents"] = o.FixedAmountCents
	}
	return toSerialize, nil
}

type NullablePOSTFixedAmountPromotionsRequestDataAttributes struct {
	value *POSTFixedAmountPromotionsRequestDataAttributes
	isSet bool
}

func (v NullablePOSTFixedAmountPromotionsRequestDataAttributes) Get() *POSTFixedAmountPromotionsRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTFixedAmountPromotionsRequestDataAttributes) Set(val *POSTFixedAmountPromotionsRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFixedAmountPromotionsRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFixedAmountPromotionsRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFixedAmountPromotionsRequestDataAttributes(val *POSTFixedAmountPromotionsRequestDataAttributes) *NullablePOSTFixedAmountPromotionsRequestDataAttributes {
	return &NullablePOSTFixedAmountPromotionsRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTFixedAmountPromotionsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFixedAmountPromotionsRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}