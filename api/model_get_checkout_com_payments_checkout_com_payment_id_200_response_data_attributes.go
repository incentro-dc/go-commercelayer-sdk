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

// checks if the GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes{}

// GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes struct for GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes
type GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes struct {
	// The Checkout.com publishable API key.
	PublicKey interface{} `json:"public_key,omitempty"`
	// The payment source type.
	PaymentType interface{} `json:"payment_type,omitempty"`
	// The Checkout.com card or digital wallet token.
	Token interface{} `json:"token,omitempty"`
	// A payment session ID used to obtain the details.
	SessionId interface{} `json:"session_id,omitempty"`
	// The URL to redirect your customer upon 3DS succeeded authentication.
	SuccessUrl interface{} `json:"success_url,omitempty"`
	// The URL to redirect your customer upon 3DS failed authentication.
	FailureUrl interface{} `json:"failure_url,omitempty"`
	// The payment source identifier that can be used for subsequent payments.
	SourceId interface{} `json:"source_id,omitempty"`
	// The customer's unique identifier. This can be passed as a source when making a payment.
	CustomerToken interface{} `json:"customer_token,omitempty"`
	// The URI that the customer should be redirected to in order to complete the payment.
	RedirectUri interface{} `json:"redirect_uri,omitempty"`
	// The Checkout.com payment response, used to fetch internal data.
	PaymentResponse interface{} `json:"payment_response,omitempty"`
	// Indicates if the order current amount differs form the one of the associated authorization.
	MismatchedAmounts interface{} `json:"mismatched_amounts,omitempty"`
	// Information about the payment instrument used in the transaction
	PaymentInstrument interface{} `json:"payment_instrument,omitempty"`
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

// NewGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes instantiates a new GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes() *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes {
	this := GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes{}
	return &this
}

// NewGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributesWithDefaults() *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes {
	this := GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPublicKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPublicKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return &o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasPublicKey() bool {
	if o != nil && IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given interface{} and assigns it to the PublicKey field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPublicKey(v interface{}) {
	o.PublicKey = v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return &o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasPaymentType() bool {
	if o != nil && IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given interface{} and assigns it to the PaymentType field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPaymentType(v interface{}) {
	o.PaymentType = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return &o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasToken() bool {
	if o != nil && IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given interface{} and assigns it to the Token field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetToken(v interface{}) {
	o.Token = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSessionId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSessionIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return &o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSessionId() bool {
	if o != nil && IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given interface{} and assigns it to the SessionId field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSessionId(v interface{}) {
	o.SessionId = v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSuccessUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSuccessUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return &o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSuccessUrl() bool {
	if o != nil && IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given interface{} and assigns it to the SuccessUrl field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSuccessUrl(v interface{}) {
	o.SuccessUrl = v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetFailureUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetFailureUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FailureUrl) {
		return nil, false
	}
	return &o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasFailureUrl() bool {
	if o != nil && IsNil(o.FailureUrl) {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given interface{} and assigns it to the FailureUrl field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetFailureUrl(v interface{}) {
	o.FailureUrl = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSourceId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSourceIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return &o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSourceId() bool {
	if o != nil && IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given interface{} and assigns it to the SourceId field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSourceId(v interface{}) {
	o.SourceId = v
}

// GetCustomerToken returns the CustomerToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetCustomerToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerToken
}

// GetCustomerTokenOk returns a tuple with the CustomerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetCustomerTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerToken) {
		return nil, false
	}
	return &o.CustomerToken, true
}

// HasCustomerToken returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasCustomerToken() bool {
	if o != nil && IsNil(o.CustomerToken) {
		return true
	}

	return false
}

// SetCustomerToken gets a reference to the given interface{} and assigns it to the CustomerToken field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetCustomerToken(v interface{}) {
	o.CustomerToken = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetRedirectUri() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetRedirectUriOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return &o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasRedirectUri() bool {
	if o != nil && IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given interface{} and assigns it to the RedirectUri field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetRedirectUri(v interface{}) {
	o.RedirectUri = v
}

// GetPaymentResponse returns the PaymentResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentResponse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentResponse
}

// GetPaymentResponseOk returns a tuple with the PaymentResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentResponseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentResponse) {
		return nil, false
	}
	return &o.PaymentResponse, true
}

// HasPaymentResponse returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasPaymentResponse() bool {
	if o != nil && IsNil(o.PaymentResponse) {
		return true
	}

	return false
}

// SetPaymentResponse gets a reference to the given interface{} and assigns it to the PaymentResponse field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPaymentResponse(v interface{}) {
	o.PaymentResponse = v
}

// GetMismatchedAmounts returns the MismatchedAmounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MismatchedAmounts
}

// GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MismatchedAmounts) {
		return nil, false
	}
	return &o.MismatchedAmounts, true
}

// HasMismatchedAmounts returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool {
	if o != nil && IsNil(o.MismatchedAmounts) {
		return true
	}

	return false
}

// SetMismatchedAmounts gets a reference to the given interface{} and assigns it to the MismatchedAmounts field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{}) {
	o.MismatchedAmounts = v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return &o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasPaymentInstrument() bool {
	if o != nil && IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given interface{} and assigns it to the PaymentInstrument field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{}) {
	o.PaymentInstrument = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.PaymentInstrument != nil {
		toSerialize["payment_instrument"] = o.PaymentInstrument
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

type NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes struct {
	value *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) Get() *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) Set(val *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes(val *GETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) *NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes {
	return &NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}