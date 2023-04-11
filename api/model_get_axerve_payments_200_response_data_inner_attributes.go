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

// GETAxervePayments200ResponseDataInnerAttributes struct for GETAxervePayments200ResponseDataInnerAttributes
type GETAxervePayments200ResponseDataInnerAttributes struct {
	// The merchant login code.
	Login interface{} `json:"login,omitempty"`
	// The URL where the payer is redirected after they approve the payment.
	ReturnUrl interface{} `json:"return_url,omitempty"`
	// The Axerve payment request data, collected by client.
	PaymentRequestData interface{} `json:"payment_request_data,omitempty"`
	// Indicates if the order current amount differs form the one of the associated authorization.
	MismatchedAmounts interface{} `json:"mismatched_amounts,omitempty"`
	// The amount of the associated payment intent, in cents.
	IntentAmountCents interface{} `json:"intent_amount_cents,omitempty"`
	// The amount of the associated payment intent, float.
	IntentAmountFloat interface{} `json:"intent_amount_float,omitempty"`
	// The amount of the associated payment intent, formatted.
	FormattedIntentAmount interface{} `json:"formatted_intent_amount,omitempty"`
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

// NewGETAxervePayments200ResponseDataInnerAttributes instantiates a new GETAxervePayments200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAxervePayments200ResponseDataInnerAttributes() *GETAxervePayments200ResponseDataInnerAttributes {
	this := GETAxervePayments200ResponseDataInnerAttributes{}
	return &this
}

// NewGETAxervePayments200ResponseDataInnerAttributesWithDefaults instantiates a new GETAxervePayments200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAxervePayments200ResponseDataInnerAttributesWithDefaults() *GETAxervePayments200ResponseDataInnerAttributes {
	this := GETAxervePayments200ResponseDataInnerAttributes{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetLogin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetLoginOk() (*interface{}, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return &o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given interface{} and assigns it to the Login field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetLogin(v interface{}) {
	o.Login = v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || o.ReturnUrl == nil {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasReturnUrl() bool {
	if o != nil && o.ReturnUrl != nil {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given interface{} and assigns it to the ReturnUrl field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetPaymentRequestData returns the PaymentRequestData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentRequestData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentRequestData
}

// GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentRequestDataOk() (*interface{}, bool) {
	if o == nil || o.PaymentRequestData == nil {
		return nil, false
	}
	return &o.PaymentRequestData, true
}

// HasPaymentRequestData returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasPaymentRequestData() bool {
	if o != nil && o.PaymentRequestData != nil {
		return true
	}

	return false
}

// SetPaymentRequestData gets a reference to the given interface{} and assigns it to the PaymentRequestData field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetPaymentRequestData(v interface{}) {
	o.PaymentRequestData = v
}

// GetMismatchedAmounts returns the MismatchedAmounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMismatchedAmounts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MismatchedAmounts
}

// GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMismatchedAmountsOk() (*interface{}, bool) {
	if o == nil || o.MismatchedAmounts == nil {
		return nil, false
	}
	return &o.MismatchedAmounts, true
}

// HasMismatchedAmounts returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasMismatchedAmounts() bool {
	if o != nil && o.MismatchedAmounts != nil {
		return true
	}

	return false
}

// SetMismatchedAmounts gets a reference to the given interface{} and assigns it to the MismatchedAmounts field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetMismatchedAmounts(v interface{}) {
	o.MismatchedAmounts = v
}

// GetIntentAmountCents returns the IntentAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IntentAmountCents
}

// GetIntentAmountCentsOk returns a tuple with the IntentAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountCentsOk() (*interface{}, bool) {
	if o == nil || o.IntentAmountCents == nil {
		return nil, false
	}
	return &o.IntentAmountCents, true
}

// HasIntentAmountCents returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasIntentAmountCents() bool {
	if o != nil && o.IntentAmountCents != nil {
		return true
	}

	return false
}

// SetIntentAmountCents gets a reference to the given interface{} and assigns it to the IntentAmountCents field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetIntentAmountCents(v interface{}) {
	o.IntentAmountCents = v
}

// GetIntentAmountFloat returns the IntentAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IntentAmountFloat
}

// GetIntentAmountFloatOk returns a tuple with the IntentAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetIntentAmountFloatOk() (*interface{}, bool) {
	if o == nil || o.IntentAmountFloat == nil {
		return nil, false
	}
	return &o.IntentAmountFloat, true
}

// HasIntentAmountFloat returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasIntentAmountFloat() bool {
	if o != nil && o.IntentAmountFloat != nil {
		return true
	}

	return false
}

// SetIntentAmountFloat gets a reference to the given interface{} and assigns it to the IntentAmountFloat field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetIntentAmountFloat(v interface{}) {
	o.IntentAmountFloat = v
}

// GetFormattedIntentAmount returns the FormattedIntentAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetFormattedIntentAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedIntentAmount
}

// GetFormattedIntentAmountOk returns a tuple with the FormattedIntentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetFormattedIntentAmountOk() (*interface{}, bool) {
	if o == nil || o.FormattedIntentAmount == nil {
		return nil, false
	}
	return &o.FormattedIntentAmount, true
}

// HasFormattedIntentAmount returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasFormattedIntentAmount() bool {
	if o != nil && o.FormattedIntentAmount != nil {
		return true
	}

	return false
}

// SetFormattedIntentAmount gets a reference to the given interface{} and assigns it to the FormattedIntentAmount field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetFormattedIntentAmount(v interface{}) {
	o.FormattedIntentAmount = v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentInstrument() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetPaymentInstrumentOk() (*interface{}, bool) {
	if o == nil || o.PaymentInstrument == nil {
		return nil, false
	}
	return &o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasPaymentInstrument() bool {
	if o != nil && o.PaymentInstrument != nil {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given interface{} and assigns it to the PaymentInstrument field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetPaymentInstrument(v interface{}) {
	o.PaymentInstrument = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAxervePayments200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETAxervePayments200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETAxervePayments200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETAxervePayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.PaymentRequestData != nil {
		toSerialize["payment_request_data"] = o.PaymentRequestData
	}
	if o.MismatchedAmounts != nil {
		toSerialize["mismatched_amounts"] = o.MismatchedAmounts
	}
	if o.IntentAmountCents != nil {
		toSerialize["intent_amount_cents"] = o.IntentAmountCents
	}
	if o.IntentAmountFloat != nil {
		toSerialize["intent_amount_float"] = o.IntentAmountFloat
	}
	if o.FormattedIntentAmount != nil {
		toSerialize["formatted_intent_amount"] = o.FormattedIntentAmount
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
	return json.Marshal(toSerialize)
}

type NullableGETAxervePayments200ResponseDataInnerAttributes struct {
	value *GETAxervePayments200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETAxervePayments200ResponseDataInnerAttributes) Get() *GETAxervePayments200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETAxervePayments200ResponseDataInnerAttributes) Set(val *GETAxervePayments200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAxervePayments200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAxervePayments200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAxervePayments200ResponseDataInnerAttributes(val *GETAxervePayments200ResponseDataInnerAttributes) *NullableGETAxervePayments200ResponseDataInnerAttributes {
	return &NullableGETAxervePayments200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETAxervePayments200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAxervePayments200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
