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

// CustomerPasswordResetDataAttributes struct for CustomerPasswordResetDataAttributes
type CustomerPasswordResetDataAttributes struct {
	// The email of the customer that requires a password reset
	CustomerEmail *string `json:"customer_email,omitempty"`
	// Automatically generated on create. Send its value as the '_reset_password_token' argument when updating the customer password.
	ResetPasswordToken *string `json:"reset_password_token,omitempty"`
	// Time at which the password was reset.
	ResetPasswordAt *string `json:"reset_password_at,omitempty"`
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

// NewCustomerPasswordResetDataAttributes instantiates a new CustomerPasswordResetDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetDataAttributes() *CustomerPasswordResetDataAttributes {
	this := CustomerPasswordResetDataAttributes{}
	return &this
}

// NewCustomerPasswordResetDataAttributesWithDefaults instantiates a new CustomerPasswordResetDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetDataAttributesWithDefaults() *CustomerPasswordResetDataAttributes {
	this := CustomerPasswordResetDataAttributes{}
	return &this
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetCustomerEmail() string {
	if o == nil || o.CustomerEmail == nil {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetCustomerEmailOk() (*string, bool) {
	if o == nil || o.CustomerEmail == nil {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail != nil {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *CustomerPasswordResetDataAttributes) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetResetPasswordToken returns the ResetPasswordToken field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetResetPasswordToken() string {
	if o == nil || o.ResetPasswordToken == nil {
		var ret string
		return ret
	}
	return *o.ResetPasswordToken
}

// GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetResetPasswordTokenOk() (*string, bool) {
	if o == nil || o.ResetPasswordToken == nil {
		return nil, false
	}
	return o.ResetPasswordToken, true
}

// HasResetPasswordToken returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasResetPasswordToken() bool {
	if o != nil && o.ResetPasswordToken != nil {
		return true
	}

	return false
}

// SetResetPasswordToken gets a reference to the given string and assigns it to the ResetPasswordToken field.
func (o *CustomerPasswordResetDataAttributes) SetResetPasswordToken(v string) {
	o.ResetPasswordToken = &v
}

// GetResetPasswordAt returns the ResetPasswordAt field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetResetPasswordAt() string {
	if o == nil || o.ResetPasswordAt == nil {
		var ret string
		return ret
	}
	return *o.ResetPasswordAt
}

// GetResetPasswordAtOk returns a tuple with the ResetPasswordAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetResetPasswordAtOk() (*string, bool) {
	if o == nil || o.ResetPasswordAt == nil {
		return nil, false
	}
	return o.ResetPasswordAt, true
}

// HasResetPasswordAt returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasResetPasswordAt() bool {
	if o != nil && o.ResetPasswordAt != nil {
		return true
	}

	return false
}

// SetResetPasswordAt gets a reference to the given string and assigns it to the ResetPasswordAt field.
func (o *CustomerPasswordResetDataAttributes) SetResetPasswordAt(v string) {
	o.ResetPasswordAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerPasswordResetDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CustomerPasswordResetDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CustomerPasswordResetDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *CustomerPasswordResetDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *CustomerPasswordResetDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerPasswordResetDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CustomerPasswordResetDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.ResetPasswordToken != nil {
		toSerialize["reset_password_token"] = o.ResetPasswordToken
	}
	if o.ResetPasswordAt != nil {
		toSerialize["reset_password_at"] = o.ResetPasswordAt
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

type NullableCustomerPasswordResetDataAttributes struct {
	value *CustomerPasswordResetDataAttributes
	isSet bool
}

func (v NullableCustomerPasswordResetDataAttributes) Get() *CustomerPasswordResetDataAttributes {
	return v.value
}

func (v *NullableCustomerPasswordResetDataAttributes) Set(val *CustomerPasswordResetDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetDataAttributes(val *CustomerPasswordResetDataAttributes) *NullableCustomerPasswordResetDataAttributes {
	return &NullableCustomerPasswordResetDataAttributes{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}