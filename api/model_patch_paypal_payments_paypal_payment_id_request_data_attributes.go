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

// checks if the PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes{}

// PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes struct for PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes
type PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes struct {
	// The id of the payer that PayPal passes in the return_url.
	PaypalPayerId interface{} `json:"paypal_payer_id,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes instantiates a new PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes() *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes {
	this := PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes{}
	return &this
}

// NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributesWithDefaults instantiates a new PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributesWithDefaults() *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes {
	this := PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes{}
	return &this
}

// GetPaypalPayerId returns the PaypalPayerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetPaypalPayerId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaypalPayerId
}

// GetPaypalPayerIdOk returns a tuple with the PaypalPayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetPaypalPayerIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaypalPayerId) {
		return nil, false
	}
	return &o.PaypalPayerId, true
}

// HasPaypalPayerId returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) HasPaypalPayerId() bool {
	if o != nil && IsNil(o.PaypalPayerId) {
		return true
	}

	return false
}

// SetPaypalPayerId gets a reference to the given interface{} and assigns it to the PaypalPayerId field.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) SetPaypalPayerId(v interface{}) {
	o.PaypalPayerId = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaypalPayerId != nil {
		toSerialize["paypal_payer_id"] = o.PaypalPayerId
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

type NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes struct {
	value *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) Get() *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) Set(val *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes(val *PATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) *NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes {
	return &NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}