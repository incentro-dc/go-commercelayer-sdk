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

// checks if the GETCouponsCouponId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCouponsCouponId200ResponseDataAttributes{}

// GETCouponsCouponId200ResponseDataAttributes struct for GETCouponsCouponId200ResponseDataAttributes
type GETCouponsCouponId200ResponseDataAttributes struct {
	// The coupon code, that uniquely identifies the coupon within the promotion rule.
	Code interface{} `json:"code,omitempty"`
	// Indicates if the coupon can be used just once per customer.
	CustomerSingleUse interface{} `json:"customer_single_use,omitempty"`
	// The total number of times this coupon can be used.
	UsageLimit interface{} `json:"usage_limit,omitempty"`
	// The number of times this coupon has been used.
	UsageCount interface{} `json:"usage_count,omitempty"`
	// The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail interface{} `json:"recipient_email,omitempty"`
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

// NewGETCouponsCouponId200ResponseDataAttributes instantiates a new GETCouponsCouponId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponsCouponId200ResponseDataAttributes() *GETCouponsCouponId200ResponseDataAttributes {
	this := GETCouponsCouponId200ResponseDataAttributes{}
	return &this
}

// NewGETCouponsCouponId200ResponseDataAttributesWithDefaults instantiates a new GETCouponsCouponId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponsCouponId200ResponseDataAttributesWithDefaults() *GETCouponsCouponId200ResponseDataAttributes {
	this := GETCouponsCouponId200ResponseDataAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetCustomerSingleUse returns the CustomerSingleUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetCustomerSingleUse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerSingleUse
}

// GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetCustomerSingleUseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerSingleUse) {
		return nil, false
	}
	return &o.CustomerSingleUse, true
}

// HasCustomerSingleUse returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasCustomerSingleUse() bool {
	if o != nil && IsNil(o.CustomerSingleUse) {
		return true
	}

	return false
}

// SetCustomerSingleUse gets a reference to the given interface{} and assigns it to the CustomerSingleUse field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetCustomerSingleUse(v interface{}) {
	o.CustomerSingleUse = v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetUsageLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetUsageLimitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UsageLimit) {
		return nil, false
	}
	return &o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasUsageLimit() bool {
	if o != nil && IsNil(o.UsageLimit) {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given interface{} and assigns it to the UsageLimit field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetUsageLimit(v interface{}) {
	o.UsageLimit = v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetUsageCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetUsageCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return &o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasUsageCount() bool {
	if o != nil && IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given interface{} and assigns it to the UsageCount field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetUsageCount(v interface{}) {
	o.UsageCount = v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetRecipientEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RecipientEmail) {
		return nil, false
	}
	return &o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasRecipientEmail() bool {
	if o != nil && IsNil(o.RecipientEmail) {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given interface{} and assigns it to the RecipientEmail field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetRecipientEmail(v interface{}) {
	o.RecipientEmail = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCouponsCouponId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCouponsCouponId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCouponsCouponId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCouponsCouponId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETCouponsCouponId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCouponsCouponId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CustomerSingleUse != nil {
		toSerialize["customer_single_use"] = o.CustomerSingleUse
	}
	if o.UsageLimit != nil {
		toSerialize["usage_limit"] = o.UsageLimit
	}
	if o.UsageCount != nil {
		toSerialize["usage_count"] = o.UsageCount
	}
	if o.RecipientEmail != nil {
		toSerialize["recipient_email"] = o.RecipientEmail
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
	return toSerialize, nil
}

type NullableGETCouponsCouponId200ResponseDataAttributes struct {
	value *GETCouponsCouponId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCouponsCouponId200ResponseDataAttributes) Get() *GETCouponsCouponId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCouponsCouponId200ResponseDataAttributes) Set(val *GETCouponsCouponId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponsCouponId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponsCouponId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponsCouponId200ResponseDataAttributes(val *GETCouponsCouponId200ResponseDataAttributes) *NullableGETCouponsCouponId200ResponseDataAttributes {
	return &NullableGETCouponsCouponId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCouponsCouponId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponsCouponId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
