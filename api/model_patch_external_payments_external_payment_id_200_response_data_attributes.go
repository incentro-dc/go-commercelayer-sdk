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

// PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes struct for PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes
type PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes struct {
	// External payment options.
	Options interface{} `json:"options,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes instantiates a new PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes() *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	this := PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHExternalPaymentsExternalPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalPaymentsExternalPaymentId200ResponseDataAttributesWithDefaults() *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	this := PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes struct {
	value *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) Get() *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) Set(val *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes(val *PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) *NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	return &NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
