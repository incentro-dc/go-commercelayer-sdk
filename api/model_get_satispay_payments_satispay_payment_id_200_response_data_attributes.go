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

// checks if the GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{}

// GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes struct for GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes
type GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes struct {
	// Activation code generated from the Satispay Dashboard.
	Token interface{} `json:"token,omitempty"`
	// The Satispay API key auto generated basing on activation code.
	KeyId interface{} `json:"key_id,omitempty"`
	// The payment unique identifier.
	PaymentId interface{} `json:"payment_id,omitempty"`
	// The url to redirect the customer after the payment flow is completed.
	RedirectUrl interface{} `json:"redirect_url,omitempty"`
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

// NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes instantiates a new GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes() *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	this := GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{}
	return &this
}

// NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults() *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	this := GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return &o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasToken() bool {
	if o != nil && IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given interface{} and assigns it to the Token field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetToken(v interface{}) {
	o.Token = v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetKeyId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetKeyIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return &o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasKeyId() bool {
	if o != nil && IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given interface{} and assigns it to the KeyId field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetKeyId(v interface{}) {
	o.KeyId = v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return &o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasPaymentId() bool {
	if o != nil && IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given interface{} and assigns it to the PaymentId field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentId(v interface{}) {
	o.PaymentId = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRedirectUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRedirectUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasRedirectUrl() bool {
	if o != nil && IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given interface{} and assigns it to the RedirectUrl field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetRedirectUrl(v interface{}) {
	o.RedirectUrl = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.KeyId != nil {
		toSerialize["key_id"] = o.KeyId
	}
	if o.PaymentId != nil {
		toSerialize["payment_id"] = o.PaymentId
	}
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
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

type NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes struct {
	value *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) Get() *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) Set(val *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes(val *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) *NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	return &NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}