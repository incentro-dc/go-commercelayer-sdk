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

// GETBraintreePayments200ResponseDataInnerAttributes struct for GETBraintreePayments200ResponseDataInnerAttributes
type GETBraintreePayments200ResponseDataInnerAttributes struct {
	// The Braintree payment client token. Required by the Braintree JS SDK.
	ClientToken *string `json:"client_token,omitempty"`
	// The Braintree payment method nonce. Sent by the Braintree JS SDK.
	PaymentMethodNonce *string `json:"payment_method_nonce,omitempty"`
	// The Braintree payment ID used by local payment and sent by the Braintree JS SDK.
	PaymentId *string `json:"payment_id,omitempty"`
	// Indicates if the payment is local, in such case Braintree will trigger a webhook call passing the \"payment_id\" and \"payment_method_nonce\" in order to complete the transaction.
	Local *bool `json:"local,omitempty"`
	// Braintree payment options: 'customer_id' and 'payment_method_token'
	Options map[string]interface{} `json:"options,omitempty"`
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
}

// NewGETBraintreePayments200ResponseDataInnerAttributes instantiates a new GETBraintreePayments200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreePayments200ResponseDataInnerAttributes() *GETBraintreePayments200ResponseDataInnerAttributes {
	this := GETBraintreePayments200ResponseDataInnerAttributes{}
	return &this
}

// NewGETBraintreePayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETBraintreePayments200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreePayments200ResponseDataInnerAttributesWithDefaults() *GETBraintreePayments200ResponseDataInnerAttributes {
	this := GETBraintreePayments200ResponseDataInnerAttributes{}
	return &this
}

// GetClientToken returns the ClientToken field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetClientToken() string {
	if o == nil || o.ClientToken == nil {
		var ret string
		return ret
	}
	return *o.ClientToken
}

// GetClientTokenOk returns a tuple with the ClientToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetClientTokenOk() (*string, bool) {
	if o == nil || o.ClientToken == nil {
		return nil, false
	}
	return o.ClientToken, true
}

// HasClientToken returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasClientToken() bool {
	if o != nil && o.ClientToken != nil {
		return true
	}

	return false
}

// SetClientToken gets a reference to the given string and assigns it to the ClientToken field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetClientToken(v string) {
	o.ClientToken = &v
}

// GetPaymentMethodNonce returns the PaymentMethodNonce field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetPaymentMethodNonce() string {
	if o == nil || o.PaymentMethodNonce == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodNonce
}

// GetPaymentMethodNonceOk returns a tuple with the PaymentMethodNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetPaymentMethodNonceOk() (*string, bool) {
	if o == nil || o.PaymentMethodNonce == nil {
		return nil, false
	}
	return o.PaymentMethodNonce, true
}

// HasPaymentMethodNonce returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasPaymentMethodNonce() bool {
	if o != nil && o.PaymentMethodNonce != nil {
		return true
	}

	return false
}

// SetPaymentMethodNonce gets a reference to the given string and assigns it to the PaymentMethodNonce field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetPaymentMethodNonce(v string) {
	o.PaymentMethodNonce = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetLocal(v bool) {
	o.Local = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETBraintreePayments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETBraintreePayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientToken != nil {
		toSerialize["client_token"] = o.ClientToken
	}
	if o.PaymentMethodNonce != nil {
		toSerialize["payment_method_nonce"] = o.PaymentMethodNonce
	}
	if o.PaymentId != nil {
		toSerialize["payment_id"] = o.PaymentId
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullableGETBraintreePayments200ResponseDataInnerAttributes struct {
	value *GETBraintreePayments200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETBraintreePayments200ResponseDataInnerAttributes) Get() *GETBraintreePayments200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETBraintreePayments200ResponseDataInnerAttributes) Set(val *GETBraintreePayments200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreePayments200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreePayments200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreePayments200ResponseDataInnerAttributes(val *GETBraintreePayments200ResponseDataInnerAttributes) *NullableGETBraintreePayments200ResponseDataInnerAttributes {
	return &NullableGETBraintreePayments200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETBraintreePayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreePayments200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
