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

// POSTGiftCards201ResponseDataAttributes struct for POSTGiftCards201ResponseDataAttributes
type POSTGiftCards201ResponseDataAttributes struct {
	// The gift card code UUID. If not set, it's automatically generated.
	Code *string `json:"code,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The gift card balance, in cents.
	BalanceCents int32 `json:"balance_cents"`
	// The gift card balance max, in cents.
	BalanceMaxCents *string `json:"balance_max_cents,omitempty"`
	// Indicates if the gift card can be used only one.
	SingleUse *bool `json:"single_use,omitempty"`
	// Indicates if the gift card can be recharged.
	Rechargeable *bool `json:"rechargeable,omitempty"`
	// The URL of an image that represents the gift card.
	ImageUrl *string `json:"image_url,omitempty"`
	// Time at which the gift card will expire.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail *string `json:"recipient_email,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTGiftCards201ResponseDataAttributes instantiates a new POSTGiftCards201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGiftCards201ResponseDataAttributes(balanceCents int32) *POSTGiftCards201ResponseDataAttributes {
	this := POSTGiftCards201ResponseDataAttributes{}
	this.BalanceCents = balanceCents
	return &this
}

// NewPOSTGiftCards201ResponseDataAttributesWithDefaults instantiates a new POSTGiftCards201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGiftCards201ResponseDataAttributesWithDefaults() *POSTGiftCards201ResponseDataAttributes {
	this := POSTGiftCards201ResponseDataAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *POSTGiftCards201ResponseDataAttributes) SetCode(v string) {
	o.Code = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *POSTGiftCards201ResponseDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetBalanceCents returns the BalanceCents field value
func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BalanceCents
}

// GetBalanceCentsOk returns a tuple with the BalanceCents field value
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceCents, true
}

// SetBalanceCents sets field value
func (o *POSTGiftCards201ResponseDataAttributes) SetBalanceCents(v int32) {
	o.BalanceCents = v
}

// GetBalanceMaxCents returns the BalanceMaxCents field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceMaxCents() string {
	if o == nil || o.BalanceMaxCents == nil {
		var ret string
		return ret
	}
	return *o.BalanceMaxCents
}

// GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetBalanceMaxCentsOk() (*string, bool) {
	if o == nil || o.BalanceMaxCents == nil {
		return nil, false
	}
	return o.BalanceMaxCents, true
}

// HasBalanceMaxCents returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasBalanceMaxCents() bool {
	if o != nil && o.BalanceMaxCents != nil {
		return true
	}

	return false
}

// SetBalanceMaxCents gets a reference to the given string and assigns it to the BalanceMaxCents field.
func (o *POSTGiftCards201ResponseDataAttributes) SetBalanceMaxCents(v string) {
	o.BalanceMaxCents = &v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetSingleUse() bool {
	if o == nil || o.SingleUse == nil {
		var ret bool
		return ret
	}
	return *o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetSingleUseOk() (*bool, bool) {
	if o == nil || o.SingleUse == nil {
		return nil, false
	}
	return o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasSingleUse() bool {
	if o != nil && o.SingleUse != nil {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given bool and assigns it to the SingleUse field.
func (o *POSTGiftCards201ResponseDataAttributes) SetSingleUse(v bool) {
	o.SingleUse = &v
}

// GetRechargeable returns the Rechargeable field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetRechargeable() bool {
	if o == nil || o.Rechargeable == nil {
		var ret bool
		return ret
	}
	return *o.Rechargeable
}

// GetRechargeableOk returns a tuple with the Rechargeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetRechargeableOk() (*bool, bool) {
	if o == nil || o.Rechargeable == nil {
		return nil, false
	}
	return o.Rechargeable, true
}

// HasRechargeable returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasRechargeable() bool {
	if o != nil && o.Rechargeable != nil {
		return true
	}

	return false
}

// SetRechargeable gets a reference to the given bool and assigns it to the Rechargeable field.
func (o *POSTGiftCards201ResponseDataAttributes) SetRechargeable(v bool) {
	o.Rechargeable = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *POSTGiftCards201ResponseDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *POSTGiftCards201ResponseDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetRecipientEmail() string {
	if o == nil || o.RecipientEmail == nil {
		var ret string
		return ret
	}
	return *o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetRecipientEmailOk() (*string, bool) {
	if o == nil || o.RecipientEmail == nil {
		return nil, false
	}
	return o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail != nil {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given string and assigns it to the RecipientEmail field.
func (o *POSTGiftCards201ResponseDataAttributes) SetRecipientEmail(v string) {
	o.RecipientEmail = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTGiftCards201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTGiftCards201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTGiftCards201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTGiftCards201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if true {
		toSerialize["balance_cents"] = o.BalanceCents
	}
	if o.BalanceMaxCents != nil {
		toSerialize["balance_max_cents"] = o.BalanceMaxCents
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	if o.Rechargeable != nil {
		toSerialize["rechargeable"] = o.Rechargeable
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
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

type NullablePOSTGiftCards201ResponseDataAttributes struct {
	value *POSTGiftCards201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTGiftCards201ResponseDataAttributes) Get() *POSTGiftCards201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTGiftCards201ResponseDataAttributes) Set(val *POSTGiftCards201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGiftCards201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGiftCards201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGiftCards201ResponseDataAttributes(val *POSTGiftCards201ResponseDataAttributes) *NullablePOSTGiftCards201ResponseDataAttributes {
	return &NullablePOSTGiftCards201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTGiftCards201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGiftCards201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
