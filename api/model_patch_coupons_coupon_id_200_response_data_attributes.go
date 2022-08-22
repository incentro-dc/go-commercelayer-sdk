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

// PATCHCouponsCouponId200ResponseDataAttributes struct for PATCHCouponsCouponId200ResponseDataAttributes
type PATCHCouponsCouponId200ResponseDataAttributes struct {
	// The coupon code, that uniquely identifies the coupon within the promotion rule.
	Code *string `json:"code,omitempty"`
	// Indicates if the coupon can be used just once per customer.
	CustomerSingleUse *bool `json:"customer_single_use,omitempty"`
	// The total number of times this coupon can be used.
	UsageLimit *int32 `json:"usage_limit,omitempty"`
	// The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail *string `json:"recipient_email,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHCouponsCouponId200ResponseDataAttributes instantiates a new PATCHCouponsCouponId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCouponsCouponId200ResponseDataAttributes() *PATCHCouponsCouponId200ResponseDataAttributes {
	this := PATCHCouponsCouponId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCouponsCouponId200ResponseDataAttributesWithDefaults instantiates a new PATCHCouponsCouponId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCouponsCouponId200ResponseDataAttributesWithDefaults() *PATCHCouponsCouponId200ResponseDataAttributes {
	this := PATCHCouponsCouponId200ResponseDataAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetCode(v string) {
	o.Code = &v
}

// GetCustomerSingleUse returns the CustomerSingleUse field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCustomerSingleUse() bool {
	if o == nil || o.CustomerSingleUse == nil {
		var ret bool
		return ret
	}
	return *o.CustomerSingleUse
}

// GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCustomerSingleUseOk() (*bool, bool) {
	if o == nil || o.CustomerSingleUse == nil {
		return nil, false
	}
	return o.CustomerSingleUse, true
}

// HasCustomerSingleUse returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasCustomerSingleUse() bool {
	if o != nil && o.CustomerSingleUse != nil {
		return true
	}

	return false
}

// SetCustomerSingleUse gets a reference to the given bool and assigns it to the CustomerSingleUse field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetCustomerSingleUse(v bool) {
	o.CustomerSingleUse = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetUsageLimit() int32 {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetUsageLimitOk() (*int32, bool) {
	if o == nil || o.UsageLimit == nil {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetRecipientEmail() string {
	if o == nil || o.RecipientEmail == nil {
		var ret string
		return ret
	}
	return *o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetRecipientEmailOk() (*string, bool) {
	if o == nil || o.RecipientEmail == nil {
		return nil, false
	}
	return o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail != nil {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given string and assigns it to the RecipientEmail field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetRecipientEmail(v string) {
	o.RecipientEmail = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHCouponsCouponId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.RecipientEmail != nil {
		toSerialize["recipient_email"] = o.RecipientEmail
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

type NullablePATCHCouponsCouponId200ResponseDataAttributes struct {
	value *PATCHCouponsCouponId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCouponsCouponId200ResponseDataAttributes) Get() *PATCHCouponsCouponId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCouponsCouponId200ResponseDataAttributes) Set(val *PATCHCouponsCouponId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCouponsCouponId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCouponsCouponId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCouponsCouponId200ResponseDataAttributes(val *PATCHCouponsCouponId200ResponseDataAttributes) *NullablePATCHCouponsCouponId200ResponseDataAttributes {
	return &NullablePATCHCouponsCouponId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCouponsCouponId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCouponsCouponId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
