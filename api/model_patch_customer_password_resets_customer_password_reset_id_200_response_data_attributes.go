/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct for PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes
type PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct {
	// The customer new password. This will be accepted only if a valid '_reset_password_token' is sent with the request.
	CustomerPassword *string `json:"customer_password,omitempty"`
	// Send the 'reset_password_token' that you got on create when updating the customer password.
	ResetPasswordToken *string `json:"_reset_password_token,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	this := PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults instantiates a new PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributesWithDefaults() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	this := PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{}
	return &this
}

// GetCustomerPassword returns the CustomerPassword field value if set, zero value otherwise.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPassword() string {
	if o == nil || o.CustomerPassword == nil {
		var ret string
		return ret
	}
	return *o.CustomerPassword
}

// GetCustomerPasswordOk returns a tuple with the CustomerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetCustomerPasswordOk() (*string, bool) {
	if o == nil || o.CustomerPassword == nil {
		return nil, false
	}
	return o.CustomerPassword, true
}

// HasCustomerPassword returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasCustomerPassword() bool {
	if o != nil && o.CustomerPassword != nil {
		return true
	}

	return false
}

// SetCustomerPassword gets a reference to the given string and assigns it to the CustomerPassword field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetCustomerPassword(v string) {
	o.CustomerPassword = &v
}

// GetResetPasswordToken returns the ResetPasswordToken field value if set, zero value otherwise.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordToken() string {
	if o == nil || o.ResetPasswordToken == nil {
		var ret string
		return ret
	}
	return *o.ResetPasswordToken
}

// GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetResetPasswordTokenOk() (*string, bool) {
	if o == nil || o.ResetPasswordToken == nil {
		return nil, false
	}
	return o.ResetPasswordToken, true
}

// HasResetPasswordToken returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasResetPasswordToken() bool {
	if o != nil && o.ResetPasswordToken != nil {
		return true
	}

	return false
}

// SetResetPasswordToken gets a reference to the given string and assigns it to the ResetPasswordToken field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetResetPasswordToken(v string) {
	o.ResetPasswordToken = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerPassword != nil {
		toSerialize["customer_password"] = o.CustomerPassword
	}
	if o.ResetPasswordToken != nil {
		toSerialize["_reset_password_token"] = o.ResetPasswordToken
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

type NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes struct {
	value *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Get() *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Set(val *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes(val *PATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes {
	return &NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
