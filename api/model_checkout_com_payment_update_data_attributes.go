/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CheckoutComPaymentUpdateDataAttributes struct for CheckoutComPaymentUpdateDataAttributes
type CheckoutComPaymentUpdateDataAttributes struct {
	// The payment source type.
	PaymentType *string `json:"payment_type,omitempty"`
	// The Checkout.com card or digital wallet token.
	Token *string `json:"token,omitempty"`
	// A payment session ID used to obtain the details.
	SessionId *string `json:"session_id,omitempty"`
	// The URL to redirect your customer upon 3DS succeeded authentication.
	SuccessUrl *string `json:"success_url,omitempty"`
	// The URL to redirect your customer upon 3DS failed authentication.
	FailureUrl *string `json:"failure_url,omitempty"`
	// Send this attribute if you want to send additional details the payment request (i.e. upon 3DS check).
	Details *bool `json:"_details,omitempty"`
	// Send this attribute if you want to refresh all the pending transactions, can be used as webhooks fallback logic.
	Refresh *bool `json:"_refresh,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewCheckoutComPaymentUpdateDataAttributes instantiates a new CheckoutComPaymentUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPaymentUpdateDataAttributes() *CheckoutComPaymentUpdateDataAttributes {
	this := CheckoutComPaymentUpdateDataAttributes{}
	return &this
}

// NewCheckoutComPaymentUpdateDataAttributesWithDefaults instantiates a new CheckoutComPaymentUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentUpdateDataAttributesWithDefaults() *CheckoutComPaymentUpdateDataAttributes {
	this := CheckoutComPaymentUpdateDataAttributes{}
	return &this
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetPaymentType() string {
	if o == nil || o.PaymentType == nil {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetPaymentTypeOk() (*string, bool) {
	if o == nil || o.PaymentType == nil {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasPaymentType() bool {
	if o != nil && o.PaymentType != nil {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetToken(v string) {
	o.Token = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetSuccessUrl() string {
	if o == nil || o.SuccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetSuccessUrlOk() (*string, bool) {
	if o == nil || o.SuccessUrl == nil {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasSuccessUrl() bool {
	if o != nil && o.SuccessUrl != nil {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetFailureUrl() string {
	if o == nil || o.FailureUrl == nil {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetFailureUrlOk() (*string, bool) {
	if o == nil || o.FailureUrl == nil {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasFailureUrl() bool {
	if o != nil && o.FailureUrl != nil {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetDetails() bool {
	if o == nil || o.Details == nil {
		var ret bool
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetDetailsOk() (*bool, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given bool and assigns it to the Details field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetDetails(v bool) {
	o.Details = &v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetRefresh() bool {
	if o == nil || o.Refresh == nil {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetRefreshOk() (*bool, bool) {
	if o == nil || o.Refresh == nil {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasRefresh() bool {
	if o != nil && o.Refresh != nil {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CheckoutComPaymentUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckoutComPaymentUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CheckoutComPaymentUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CheckoutComPaymentUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentType != nil {
		toSerialize["payment_type"] = o.PaymentType
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	if o.SuccessUrl != nil {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if o.FailureUrl != nil {
		toSerialize["failure_url"] = o.FailureUrl
	}
	if o.Details != nil {
		toSerialize["_details"] = o.Details
	}
	if o.Refresh != nil {
		toSerialize["_refresh"] = o.Refresh
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

type NullableCheckoutComPaymentUpdateDataAttributes struct {
	value *CheckoutComPaymentUpdateDataAttributes
	isSet bool
}

func (v NullableCheckoutComPaymentUpdateDataAttributes) Get() *CheckoutComPaymentUpdateDataAttributes {
	return v.value
}

func (v *NullableCheckoutComPaymentUpdateDataAttributes) Set(val *CheckoutComPaymentUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPaymentUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPaymentUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPaymentUpdateDataAttributes(val *CheckoutComPaymentUpdateDataAttributes) *NullableCheckoutComPaymentUpdateDataAttributes {
	return &NullableCheckoutComPaymentUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableCheckoutComPaymentUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPaymentUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
