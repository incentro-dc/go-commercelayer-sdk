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

// PATCHStripePaymentsStripePaymentId200ResponseDataAttributes struct for PATCHStripePaymentsStripePaymentId200ResponseDataAttributes
type PATCHStripePaymentsStripePaymentId200ResponseDataAttributes struct {
	// Stripe payment options: 'setup_future_usage', 'customer', and 'payment_method'
	Options map[string]interface{} `json:"options,omitempty"`
	// Send this attribute if you want to refresh the payment status, can be used as webhooks fallback logic.
	Refresh *bool `json:"_refresh,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes instantiates a new PATCHStripePaymentsStripePaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	this := PATCHStripePaymentsStripePaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHStripePaymentsStripePaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	this := PATCHStripePaymentsStripePaymentId200ResponseDataAttributes{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetRefresh() bool {
	if o == nil || o.Refresh == nil {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetRefreshOk() (*bool, bool) {
	if o == nil || o.Refresh == nil {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasRefresh() bool {
	if o != nil && o.Refresh != nil {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes struct {
	value *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) Get() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) Set(val *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes(val *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes {
	return &NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
