/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AdyenPaymentDataAttributes struct for AdyenPaymentDataAttributes
type AdyenPaymentDataAttributes struct {
	// The public key linked to your API credential.
	PublicKey *string `json:"public_key,omitempty"`
	// The merchant available payment methods for the assoiated order (i.e. country and amount). Required by the Adyen JS SDK.
	PaymentMethods map[string]interface{} `json:"payment_methods,omitempty"`
	// The Adyen payment request data, collected by client.
	PaymentRequestData map[string]interface{} `json:"payment_request_data,omitempty"`
	// The Adyen additional details request data, collected by client.
	PaymentRequestDetails map[string]interface{} `json:"payment_request_details,omitempty"`
	// The Adyen payment response, used by client (includes 'resultCode' and 'action').
	PaymentResponse map[string]interface{} `json:"payment_response,omitempty"`
	// Indicates if the order current amount differs form the one of the associated authorization.
	MismatchedAmounts *bool `json:"mismatched_amounts,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
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

// NewAdyenPaymentDataAttributes instantiates a new AdyenPaymentDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentDataAttributes() *AdyenPaymentDataAttributes {
	this := AdyenPaymentDataAttributes{}
	return &this
}

// NewAdyenPaymentDataAttributesWithDefaults instantiates a new AdyenPaymentDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentDataAttributesWithDefaults() *AdyenPaymentDataAttributes {
	this := AdyenPaymentDataAttributes{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *AdyenPaymentDataAttributes) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetPaymentMethods() map[string]interface{} {
	if o == nil || o.PaymentMethods == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetPaymentMethodsOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given map[string]interface{} and assigns it to the PaymentMethods field.
func (o *AdyenPaymentDataAttributes) SetPaymentMethods(v map[string]interface{}) {
	o.PaymentMethods = v
}

// GetPaymentRequestData returns the PaymentRequestData field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetPaymentRequestData() map[string]interface{} {
	if o == nil || o.PaymentRequestData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentRequestData
}

// GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetPaymentRequestDataOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentRequestData == nil {
		return nil, false
	}
	return o.PaymentRequestData, true
}

// HasPaymentRequestData returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasPaymentRequestData() bool {
	if o != nil && o.PaymentRequestData != nil {
		return true
	}

	return false
}

// SetPaymentRequestData gets a reference to the given map[string]interface{} and assigns it to the PaymentRequestData field.
func (o *AdyenPaymentDataAttributes) SetPaymentRequestData(v map[string]interface{}) {
	o.PaymentRequestData = v
}

// GetPaymentRequestDetails returns the PaymentRequestDetails field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetPaymentRequestDetails() map[string]interface{} {
	if o == nil || o.PaymentRequestDetails == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentRequestDetails
}

// GetPaymentRequestDetailsOk returns a tuple with the PaymentRequestDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetPaymentRequestDetailsOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentRequestDetails == nil {
		return nil, false
	}
	return o.PaymentRequestDetails, true
}

// HasPaymentRequestDetails returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasPaymentRequestDetails() bool {
	if o != nil && o.PaymentRequestDetails != nil {
		return true
	}

	return false
}

// SetPaymentRequestDetails gets a reference to the given map[string]interface{} and assigns it to the PaymentRequestDetails field.
func (o *AdyenPaymentDataAttributes) SetPaymentRequestDetails(v map[string]interface{}) {
	o.PaymentRequestDetails = v
}

// GetPaymentResponse returns the PaymentResponse field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetPaymentResponse() map[string]interface{} {
	if o == nil || o.PaymentResponse == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentResponse
}

// GetPaymentResponseOk returns a tuple with the PaymentResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetPaymentResponseOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentResponse == nil {
		return nil, false
	}
	return o.PaymentResponse, true
}

// HasPaymentResponse returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasPaymentResponse() bool {
	if o != nil && o.PaymentResponse != nil {
		return true
	}

	return false
}

// SetPaymentResponse gets a reference to the given map[string]interface{} and assigns it to the PaymentResponse field.
func (o *AdyenPaymentDataAttributes) SetPaymentResponse(v map[string]interface{}) {
	o.PaymentResponse = v
}

// GetMismatchedAmounts returns the MismatchedAmounts field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetMismatchedAmounts() bool {
	if o == nil || o.MismatchedAmounts == nil {
		var ret bool
		return ret
	}
	return *o.MismatchedAmounts
}

// GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetMismatchedAmountsOk() (*bool, bool) {
	if o == nil || o.MismatchedAmounts == nil {
		return nil, false
	}
	return o.MismatchedAmounts, true
}

// HasMismatchedAmounts returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasMismatchedAmounts() bool {
	if o != nil && o.MismatchedAmounts != nil {
		return true
	}

	return false
}

// SetMismatchedAmounts gets a reference to the given bool and assigns it to the MismatchedAmounts field.
func (o *AdyenPaymentDataAttributes) SetMismatchedAmounts(v bool) {
	o.MismatchedAmounts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdyenPaymentDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AdyenPaymentDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AdyenPaymentDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AdyenPaymentDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *AdyenPaymentDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AdyenPaymentDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AdyenPaymentDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AdyenPaymentDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o AdyenPaymentDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.PaymentRequestData != nil {
		toSerialize["payment_request_data"] = o.PaymentRequestData
	}
	if o.PaymentRequestDetails != nil {
		toSerialize["payment_request_details"] = o.PaymentRequestDetails
	}
	if o.PaymentResponse != nil {
		toSerialize["payment_response"] = o.PaymentResponse
	}
	if o.MismatchedAmounts != nil {
		toSerialize["mismatched_amounts"] = o.MismatchedAmounts
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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

type NullableAdyenPaymentDataAttributes struct {
	value *AdyenPaymentDataAttributes
	isSet bool
}

func (v NullableAdyenPaymentDataAttributes) Get() *AdyenPaymentDataAttributes {
	return v.value
}

func (v *NullableAdyenPaymentDataAttributes) Set(val *AdyenPaymentDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentDataAttributes(val *AdyenPaymentDataAttributes) *NullableAdyenPaymentDataAttributes {
	return &NullableAdyenPaymentDataAttributes{value: val, isSet: true}
}

func (v NullableAdyenPaymentDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
