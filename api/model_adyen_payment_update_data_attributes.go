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

// AdyenPaymentUpdateDataAttributes struct for AdyenPaymentUpdateDataAttributes
type AdyenPaymentUpdateDataAttributes struct {
	// The Adyen payment request data, collected by client.
	PaymentRequestData map[string]interface{} `json:"payment_request_data,omitempty"`
	// The Adyen additional details request data, collected by client.
	PaymentRequestDetails map[string]interface{} `json:"payment_request_details,omitempty"`
	// The Adyen payment response, used by client (includes 'resultCode' and 'action').
	PaymentResponse map[string]interface{} `json:"payment_response,omitempty"`
	// Send this attribute if you want to send additional details the payment request.
	Details *bool `json:"_details,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewAdyenPaymentUpdateDataAttributes instantiates a new AdyenPaymentUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenPaymentUpdateDataAttributes() *AdyenPaymentUpdateDataAttributes {
	this := AdyenPaymentUpdateDataAttributes{}
	return &this
}

// NewAdyenPaymentUpdateDataAttributesWithDefaults instantiates a new AdyenPaymentUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenPaymentUpdateDataAttributesWithDefaults() *AdyenPaymentUpdateDataAttributes {
	this := AdyenPaymentUpdateDataAttributes{}
	return &this
}

// GetPaymentRequestData returns the PaymentRequestData field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestData() map[string]interface{} {
	if o == nil || o.PaymentRequestData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentRequestData
}

// GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestDataOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentRequestData == nil {
		return nil, false
	}
	return o.PaymentRequestData, true
}

// HasPaymentRequestData returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasPaymentRequestData() bool {
	if o != nil && o.PaymentRequestData != nil {
		return true
	}

	return false
}

// SetPaymentRequestData gets a reference to the given map[string]interface{} and assigns it to the PaymentRequestData field.
func (o *AdyenPaymentUpdateDataAttributes) SetPaymentRequestData(v map[string]interface{}) {
	o.PaymentRequestData = v
}

// GetPaymentRequestDetails returns the PaymentRequestDetails field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestDetails() map[string]interface{} {
	if o == nil || o.PaymentRequestDetails == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentRequestDetails
}

// GetPaymentRequestDetailsOk returns a tuple with the PaymentRequestDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestDetailsOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentRequestDetails == nil {
		return nil, false
	}
	return o.PaymentRequestDetails, true
}

// HasPaymentRequestDetails returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasPaymentRequestDetails() bool {
	if o != nil && o.PaymentRequestDetails != nil {
		return true
	}

	return false
}

// SetPaymentRequestDetails gets a reference to the given map[string]interface{} and assigns it to the PaymentRequestDetails field.
func (o *AdyenPaymentUpdateDataAttributes) SetPaymentRequestDetails(v map[string]interface{}) {
	o.PaymentRequestDetails = v
}

// GetPaymentResponse returns the PaymentResponse field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetPaymentResponse() map[string]interface{} {
	if o == nil || o.PaymentResponse == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentResponse
}

// GetPaymentResponseOk returns a tuple with the PaymentResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetPaymentResponseOk() (map[string]interface{}, bool) {
	if o == nil || o.PaymentResponse == nil {
		return nil, false
	}
	return o.PaymentResponse, true
}

// HasPaymentResponse returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasPaymentResponse() bool {
	if o != nil && o.PaymentResponse != nil {
		return true
	}

	return false
}

// SetPaymentResponse gets a reference to the given map[string]interface{} and assigns it to the PaymentResponse field.
func (o *AdyenPaymentUpdateDataAttributes) SetPaymentResponse(v map[string]interface{}) {
	o.PaymentResponse = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetDetails() bool {
	if o == nil || o.Details == nil {
		var ret bool
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetDetailsOk() (*bool, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given bool and assigns it to the Details field.
func (o *AdyenPaymentUpdateDataAttributes) SetDetails(v bool) {
	o.Details = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AdyenPaymentUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *AdyenPaymentUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AdyenPaymentUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenPaymentUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AdyenPaymentUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AdyenPaymentUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o AdyenPaymentUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentRequestData != nil {
		toSerialize["payment_request_data"] = o.PaymentRequestData
	}
	if o.PaymentRequestDetails != nil {
		toSerialize["payment_request_details"] = o.PaymentRequestDetails
	}
	if o.PaymentResponse != nil {
		toSerialize["payment_response"] = o.PaymentResponse
	}
	if o.Details != nil {
		toSerialize["_details"] = o.Details
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

type NullableAdyenPaymentUpdateDataAttributes struct {
	value *AdyenPaymentUpdateDataAttributes
	isSet bool
}

func (v NullableAdyenPaymentUpdateDataAttributes) Get() *AdyenPaymentUpdateDataAttributes {
	return v.value
}

func (v *NullableAdyenPaymentUpdateDataAttributes) Set(val *AdyenPaymentUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenPaymentUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenPaymentUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenPaymentUpdateDataAttributes(val *AdyenPaymentUpdateDataAttributes) *NullableAdyenPaymentUpdateDataAttributes {
	return &NullableAdyenPaymentUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableAdyenPaymentUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenPaymentUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}