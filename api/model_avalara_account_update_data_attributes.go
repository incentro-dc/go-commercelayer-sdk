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

// AvalaraAccountUpdateDataAttributes struct for AvalaraAccountUpdateDataAttributes
type AvalaraAccountUpdateDataAttributes struct {
	// The tax calculator's internal name.
	Name *string `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The Avalara account username.
	Username *string `json:"username,omitempty"`
	// The Avalara account password.
	Password *string `json:"password,omitempty"`
	// The Avalara company code.
	CompanyCode *string `json:"company_code,omitempty"`
	// Indicates if the transaction will be recorded and visible on the Avalara website.
	CommitInvoice *string `json:"commit_invoice,omitempty"`
	// Indicates if the seller is responsible for paying/remitting the customs duty & import tax to the customs authorities.
	Ddp *string `json:"ddp,omitempty"`
}

// NewAvalaraAccountUpdateDataAttributes instantiates a new AvalaraAccountUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountUpdateDataAttributes() *AvalaraAccountUpdateDataAttributes {
	this := AvalaraAccountUpdateDataAttributes{}
	return &this
}

// NewAvalaraAccountUpdateDataAttributesWithDefaults instantiates a new AvalaraAccountUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountUpdateDataAttributesWithDefaults() *AvalaraAccountUpdateDataAttributes {
	this := AvalaraAccountUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AvalaraAccountUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AvalaraAccountUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *AvalaraAccountUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AvalaraAccountUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AvalaraAccountUpdateDataAttributes) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AvalaraAccountUpdateDataAttributes) SetPassword(v string) {
	o.Password = &v
}

// GetCompanyCode returns the CompanyCode field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetCompanyCode() string {
	if o == nil || o.CompanyCode == nil {
		var ret string
		return ret
	}
	return *o.CompanyCode
}

// GetCompanyCodeOk returns a tuple with the CompanyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetCompanyCodeOk() (*string, bool) {
	if o == nil || o.CompanyCode == nil {
		return nil, false
	}
	return o.CompanyCode, true
}

// HasCompanyCode returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasCompanyCode() bool {
	if o != nil && o.CompanyCode != nil {
		return true
	}

	return false
}

// SetCompanyCode gets a reference to the given string and assigns it to the CompanyCode field.
func (o *AvalaraAccountUpdateDataAttributes) SetCompanyCode(v string) {
	o.CompanyCode = &v
}

// GetCommitInvoice returns the CommitInvoice field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetCommitInvoice() string {
	if o == nil || o.CommitInvoice == nil {
		var ret string
		return ret
	}
	return *o.CommitInvoice
}

// GetCommitInvoiceOk returns a tuple with the CommitInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetCommitInvoiceOk() (*string, bool) {
	if o == nil || o.CommitInvoice == nil {
		return nil, false
	}
	return o.CommitInvoice, true
}

// HasCommitInvoice returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasCommitInvoice() bool {
	if o != nil && o.CommitInvoice != nil {
		return true
	}

	return false
}

// SetCommitInvoice gets a reference to the given string and assigns it to the CommitInvoice field.
func (o *AvalaraAccountUpdateDataAttributes) SetCommitInvoice(v string) {
	o.CommitInvoice = &v
}

// GetDdp returns the Ddp field value if set, zero value otherwise.
func (o *AvalaraAccountUpdateDataAttributes) GetDdp() string {
	if o == nil || o.Ddp == nil {
		var ret string
		return ret
	}
	return *o.Ddp
}

// GetDdpOk returns a tuple with the Ddp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountUpdateDataAttributes) GetDdpOk() (*string, bool) {
	if o == nil || o.Ddp == nil {
		return nil, false
	}
	return o.Ddp, true
}

// HasDdp returns a boolean if a field has been set.
func (o *AvalaraAccountUpdateDataAttributes) HasDdp() bool {
	if o != nil && o.Ddp != nil {
		return true
	}

	return false
}

// SetDdp gets a reference to the given string and assigns it to the Ddp field.
func (o *AvalaraAccountUpdateDataAttributes) SetDdp(v string) {
	o.Ddp = &v
}

func (o AvalaraAccountUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.CompanyCode != nil {
		toSerialize["company_code"] = o.CompanyCode
	}
	if o.CommitInvoice != nil {
		toSerialize["commit_invoice"] = o.CommitInvoice
	}
	if o.Ddp != nil {
		toSerialize["ddp"] = o.Ddp
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountUpdateDataAttributes struct {
	value *AvalaraAccountUpdateDataAttributes
	isSet bool
}

func (v NullableAvalaraAccountUpdateDataAttributes) Get() *AvalaraAccountUpdateDataAttributes {
	return v.value
}

func (v *NullableAvalaraAccountUpdateDataAttributes) Set(val *AvalaraAccountUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountUpdateDataAttributes(val *AvalaraAccountUpdateDataAttributes) *NullableAvalaraAccountUpdateDataAttributes {
	return &NullableAvalaraAccountUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableAvalaraAccountUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
