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

// GETCheckoutComPayments200ResponseDataInnerAttributes struct for GETCheckoutComPayments200ResponseDataInnerAttributes
type GETCheckoutComPayments200ResponseDataInnerAttributes struct {
	// The Checkout.com publishable API key.
	PublicKey *string `json:"public_key,omitempty"`
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
	// The payment source identifier that can be used for subsequent payments.
	SourceId *string `json:"source_id,omitempty"`
	// The customer's unique identifier. This can be passed as a source when making a payment.
	CustomerToken *string `json:"customer_token,omitempty"`
	// The URI that the customer should be redirected to in order to complete the payment.
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The Checkout.com payment response, used to fetch internal data.
	PaymentResponse map[string]interface{} `json:"payment_response,omitempty"`
	// Indicates if the order current amount differs form the one of the associated authorization.
	MismatchedAmounts *bool `json:"mismatched_amounts,omitempty"`
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

// NewGETCheckoutComPayments200ResponseDataInnerAttributes instantiates a new GETCheckoutComPayments200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComPayments200ResponseDataInnerAttributes() *GETCheckoutComPayments200ResponseDataInnerAttributes {
	this := GETCheckoutComPayments200ResponseDataInnerAttributes{}
	return &this
}

// NewGETCheckoutComPayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETCheckoutComPayments200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComPayments200ResponseDataInnerAttributesWithDefaults() *GETCheckoutComPayments200ResponseDataInnerAttributes {
	this := GETCheckoutComPayments200ResponseDataInnerAttributes{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentType() string {
	if o == nil || o.PaymentType == nil {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentTypeOk() (*string, bool) {
	if o == nil || o.PaymentType == nil {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentType() bool {
	if o != nil && o.PaymentType != nil {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetToken(v string) {
	o.Token = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSuccessUrl() string {
	if o == nil || o.SuccessUrl == nil {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSuccessUrlOk() (*string, bool) {
	if o == nil || o.SuccessUrl == nil {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSuccessUrl() bool {
	if o != nil && o.SuccessUrl != nil {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetFailureUrl() string {
	if o == nil || o.FailureUrl == nil {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetFailureUrlOk() (*string, bool) {
	if o == nil || o.FailureUrl == nil {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasFailureUrl() bool {
	if o != nil && o.FailureUrl != nil {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetSourceId(v string) {
	o.SourceId = &v
}

// GetCustomerToken returns the CustomerToken field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCustomerToken() string {
	if o == nil || o.CustomerToken == nil {
		var ret string
		return ret
	}
	return *o.CustomerToken
}

// GetCustomerTokenOk returns a tuple with the CustomerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCustomerTokenOk() (*string, bool) {
	if o == nil || o.CustomerToken == nil {
		return nil, false
	}
	return o.CustomerToken, true
}

// HasCustomerToken returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasCustomerToken() bool {
	if o != nil && o.CustomerToken != nil {
		return true
	}

	return false
}

// SetCustomerToken gets a reference to the given string and assigns it to the CustomerToken field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCustomerToken(v string) {
	o.CustomerToken = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetPaymentResponse returns the PaymentResponse field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentResponse() map[string]interface{} {
	if o == nil || o.PaymentResponse == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentResponse
}

// GetPaymentResponseOk returns a tuple with the PaymentResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetPaymentResponseOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentResponse == nil {
		return nil, false
	}
	return o.PaymentResponse, true
}

// HasPaymentResponse returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasPaymentResponse() bool {
	if o != nil && o.PaymentResponse != nil {
		return true
	}

	return false
}

// SetPaymentResponse gets a reference to the given map[string]interface{} and assigns it to the PaymentResponse field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetPaymentResponse(v map[string]interface{}) {
	o.PaymentResponse = v
}

// GetMismatchedAmounts returns the MismatchedAmounts field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMismatchedAmounts() bool {
	if o == nil || o.MismatchedAmounts == nil {
		var ret bool
		return ret
	}
	return *o.MismatchedAmounts
}

// GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*bool, bool) {
	if o == nil || o.MismatchedAmounts == nil {
		return nil, false
	}
	return o.MismatchedAmounts, true
}

// HasMismatchedAmounts returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool {
	if o != nil && o.MismatchedAmounts != nil {
		return true
	}

	return false
}

// SetMismatchedAmounts gets a reference to the given bool and assigns it to the MismatchedAmounts field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v bool) {
	o.MismatchedAmounts = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETCheckoutComPayments200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETCheckoutComPayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
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
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.CustomerToken != nil {
		toSerialize["customer_token"] = o.CustomerToken
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.PaymentResponse != nil {
		toSerialize["payment_response"] = o.PaymentResponse
	}
	if o.MismatchedAmounts != nil {
		toSerialize["mismatched_amounts"] = o.MismatchedAmounts
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

type NullableGETCheckoutComPayments200ResponseDataInnerAttributes struct {
	value *GETCheckoutComPayments200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETCheckoutComPayments200ResponseDataInnerAttributes) Get() *GETCheckoutComPayments200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETCheckoutComPayments200ResponseDataInnerAttributes) Set(val *GETCheckoutComPayments200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComPayments200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComPayments200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComPayments200ResponseDataInnerAttributes(val *GETCheckoutComPayments200ResponseDataInnerAttributes) *NullableGETCheckoutComPayments200ResponseDataInnerAttributes {
	return &NullableGETCheckoutComPayments200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETCheckoutComPayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComPayments200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
