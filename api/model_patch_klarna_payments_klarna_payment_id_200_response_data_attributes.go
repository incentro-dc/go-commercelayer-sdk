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

// checks if the PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{}

// PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes struct for PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes
type PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes struct {
	// The token returned by a successful client authorization, mandatory to place the order.
	AuthToken interface{} `json:"auth_token,omitempty"`
	// Send this attribute if you want to update the payment session with fresh order data.
	Update interface{} `json:"_update,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	this := PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	this := PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{}
	return &this
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AuthToken) {
		return nil, false
	}
	return &o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasAuthToken() bool {
	if o != nil && IsNil(o.AuthToken) {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given interface{} and assigns it to the AuthToken field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetAuthToken(v interface{}) {
	o.AuthToken = v
}

// GetUpdate returns the Update field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return &o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasUpdate() bool {
	if o != nil && IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given interface{} and assigns it to the Update field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetUpdate(v interface{}) {
	o.Update = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthToken != nil {
		toSerialize["auth_token"] = o.AuthToken
	}
	if o.Update != nil {
		toSerialize["_update"] = o.Update
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

type NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes struct {
	value *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) Get() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) Set(val *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes(val *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes {
	return &NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
