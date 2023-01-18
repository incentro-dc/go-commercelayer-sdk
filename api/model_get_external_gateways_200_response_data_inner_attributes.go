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

// GETExternalGateways200ResponseDataInnerAttributes struct for GETExternalGateways200ResponseDataInnerAttributes
type GETExternalGateways200ResponseDataInnerAttributes struct {
	// The payment gateway's internal name.
	Name *string `json:"name,omitempty"`
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
	// The shared secret used to sign the external request payload.
	SharedSecret *string `json:"shared_secret,omitempty"`
	// The endpoint used by the external gateway to authorize payments.
	AuthorizeUrl *string `json:"authorize_url,omitempty"`
	// The endpoint used by the external gateway to capture payments.
	CaptureUrl *string `json:"capture_url,omitempty"`
	// The endpoint used by the external gateway to void payments.
	VoidUrl *string `json:"void_url,omitempty"`
	// The endpoint used by the external gateway to refund payments.
	RefundUrl *string `json:"refund_url,omitempty"`
	// The endpoint used by the external gateway to create a customer payment token.
	TokenUrl *string `json:"token_url,omitempty"`
}

// NewGETExternalGateways200ResponseDataInnerAttributes instantiates a new GETExternalGateways200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalGateways200ResponseDataInnerAttributes() *GETExternalGateways200ResponseDataInnerAttributes {
	this := GETExternalGateways200ResponseDataInnerAttributes{}
	return &this
}

// NewGETExternalGateways200ResponseDataInnerAttributesWithDefaults instantiates a new GETExternalGateways200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalGateways200ResponseDataInnerAttributesWithDefaults() *GETExternalGateways200ResponseDataInnerAttributes {
	this := GETExternalGateways200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetAuthorizeUrl returns the AuthorizeUrl field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetAuthorizeUrl() string {
	if o == nil || o.AuthorizeUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthorizeUrl
}

// GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetAuthorizeUrlOk() (*string, bool) {
	if o == nil || o.AuthorizeUrl == nil {
		return nil, false
	}
	return o.AuthorizeUrl, true
}

// HasAuthorizeUrl returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasAuthorizeUrl() bool {
	if o != nil && o.AuthorizeUrl != nil {
		return true
	}

	return false
}

// SetAuthorizeUrl gets a reference to the given string and assigns it to the AuthorizeUrl field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetAuthorizeUrl(v string) {
	o.AuthorizeUrl = &v
}

// GetCaptureUrl returns the CaptureUrl field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCaptureUrl() string {
	if o == nil || o.CaptureUrl == nil {
		var ret string
		return ret
	}
	return *o.CaptureUrl
}

// GetCaptureUrlOk returns a tuple with the CaptureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetCaptureUrlOk() (*string, bool) {
	if o == nil || o.CaptureUrl == nil {
		return nil, false
	}
	return o.CaptureUrl, true
}

// HasCaptureUrl returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasCaptureUrl() bool {
	if o != nil && o.CaptureUrl != nil {
		return true
	}

	return false
}

// SetCaptureUrl gets a reference to the given string and assigns it to the CaptureUrl field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetCaptureUrl(v string) {
	o.CaptureUrl = &v
}

// GetVoidUrl returns the VoidUrl field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetVoidUrl() string {
	if o == nil || o.VoidUrl == nil {
		var ret string
		return ret
	}
	return *o.VoidUrl
}

// GetVoidUrlOk returns a tuple with the VoidUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetVoidUrlOk() (*string, bool) {
	if o == nil || o.VoidUrl == nil {
		return nil, false
	}
	return o.VoidUrl, true
}

// HasVoidUrl returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasVoidUrl() bool {
	if o != nil && o.VoidUrl != nil {
		return true
	}

	return false
}

// SetVoidUrl gets a reference to the given string and assigns it to the VoidUrl field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetVoidUrl(v string) {
	o.VoidUrl = &v
}

// GetRefundUrl returns the RefundUrl field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetRefundUrl() string {
	if o == nil || o.RefundUrl == nil {
		var ret string
		return ret
	}
	return *o.RefundUrl
}

// GetRefundUrlOk returns a tuple with the RefundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetRefundUrlOk() (*string, bool) {
	if o == nil || o.RefundUrl == nil {
		return nil, false
	}
	return o.RefundUrl, true
}

// HasRefundUrl returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasRefundUrl() bool {
	if o != nil && o.RefundUrl != nil {
		return true
	}

	return false
}

// SetRefundUrl gets a reference to the given string and assigns it to the RefundUrl field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetRefundUrl(v string) {
	o.RefundUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *GETExternalGateways200ResponseDataInnerAttributes) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *GETExternalGateways200ResponseDataInnerAttributes) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

func (o GETExternalGateways200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	if o.AuthorizeUrl != nil {
		toSerialize["authorize_url"] = o.AuthorizeUrl
	}
	if o.CaptureUrl != nil {
		toSerialize["capture_url"] = o.CaptureUrl
	}
	if o.VoidUrl != nil {
		toSerialize["void_url"] = o.VoidUrl
	}
	if o.RefundUrl != nil {
		toSerialize["refund_url"] = o.RefundUrl
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalGateways200ResponseDataInnerAttributes struct {
	value *GETExternalGateways200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETExternalGateways200ResponseDataInnerAttributes) Get() *GETExternalGateways200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETExternalGateways200ResponseDataInnerAttributes) Set(val *GETExternalGateways200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalGateways200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalGateways200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalGateways200ResponseDataInnerAttributes(val *GETExternalGateways200ResponseDataInnerAttributes) *NullableGETExternalGateways200ResponseDataInnerAttributes {
	return &NullableGETExternalGateways200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETExternalGateways200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalGateways200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
