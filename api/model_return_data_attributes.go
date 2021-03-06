/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReturnDataAttributes struct for ReturnDataAttributes
type ReturnDataAttributes struct {
	// Unique identifier for the return
	Number *string `json:"number,omitempty"`
	// The return status, one of 'draft', 'requested', 'approved', 'cancelled', 'shipped', 'rejected' or 'received'
	Status *string `json:"status,omitempty"`
	// The email address of the associated customer.
	CustomerEmail *string `json:"customer_email,omitempty"`
	// The total number of skus in the return's line items. This can be useful to display a preview of the return content.
	SkusCount *int32 `json:"skus_count,omitempty"`
	// Time at which the return was approved.
	ApprovedAt *string `json:"approved_at,omitempty"`
	// Time at which the return was cancelled.
	CancelledAt *string `json:"cancelled_at,omitempty"`
	// Time at which the return was shipped.
	ShippedAt *string `json:"shipped_at,omitempty"`
	// Time at which the return was rejected.
	RejectedAt *string `json:"rejected_at,omitempty"`
	// Time at which the return was received.
	ReceivedAt *string `json:"received_at,omitempty"`
	// Time at which the resource has been archived.
	ArchivedAt *string `json:"archived_at,omitempty"`
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

// NewReturnDataAttributes instantiates a new ReturnDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnDataAttributes() *ReturnDataAttributes {
	this := ReturnDataAttributes{}
	return &this
}

// NewReturnDataAttributesWithDefaults instantiates a new ReturnDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnDataAttributesWithDefaults() *ReturnDataAttributes {
	this := ReturnDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ReturnDataAttributes) SetNumber(v string) {
	o.Number = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReturnDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetCustomerEmail() string {
	if o == nil || o.CustomerEmail == nil {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetCustomerEmailOk() (*string, bool) {
	if o == nil || o.CustomerEmail == nil {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail != nil {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *ReturnDataAttributes) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetSkusCount returns the SkusCount field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetSkusCount() int32 {
	if o == nil || o.SkusCount == nil {
		var ret int32
		return ret
	}
	return *o.SkusCount
}

// GetSkusCountOk returns a tuple with the SkusCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetSkusCountOk() (*int32, bool) {
	if o == nil || o.SkusCount == nil {
		return nil, false
	}
	return o.SkusCount, true
}

// HasSkusCount returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasSkusCount() bool {
	if o != nil && o.SkusCount != nil {
		return true
	}

	return false
}

// SetSkusCount gets a reference to the given int32 and assigns it to the SkusCount field.
func (o *ReturnDataAttributes) SetSkusCount(v int32) {
	o.SkusCount = &v
}

// GetApprovedAt returns the ApprovedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetApprovedAt() string {
	if o == nil || o.ApprovedAt == nil {
		var ret string
		return ret
	}
	return *o.ApprovedAt
}

// GetApprovedAtOk returns a tuple with the ApprovedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetApprovedAtOk() (*string, bool) {
	if o == nil || o.ApprovedAt == nil {
		return nil, false
	}
	return o.ApprovedAt, true
}

// HasApprovedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasApprovedAt() bool {
	if o != nil && o.ApprovedAt != nil {
		return true
	}

	return false
}

// SetApprovedAt gets a reference to the given string and assigns it to the ApprovedAt field.
func (o *ReturnDataAttributes) SetApprovedAt(v string) {
	o.ApprovedAt = &v
}

// GetCancelledAt returns the CancelledAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetCancelledAt() string {
	if o == nil || o.CancelledAt == nil {
		var ret string
		return ret
	}
	return *o.CancelledAt
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetCancelledAtOk() (*string, bool) {
	if o == nil || o.CancelledAt == nil {
		return nil, false
	}
	return o.CancelledAt, true
}

// HasCancelledAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasCancelledAt() bool {
	if o != nil && o.CancelledAt != nil {
		return true
	}

	return false
}

// SetCancelledAt gets a reference to the given string and assigns it to the CancelledAt field.
func (o *ReturnDataAttributes) SetCancelledAt(v string) {
	o.CancelledAt = &v
}

// GetShippedAt returns the ShippedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetShippedAt() string {
	if o == nil || o.ShippedAt == nil {
		var ret string
		return ret
	}
	return *o.ShippedAt
}

// GetShippedAtOk returns a tuple with the ShippedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetShippedAtOk() (*string, bool) {
	if o == nil || o.ShippedAt == nil {
		return nil, false
	}
	return o.ShippedAt, true
}

// HasShippedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasShippedAt() bool {
	if o != nil && o.ShippedAt != nil {
		return true
	}

	return false
}

// SetShippedAt gets a reference to the given string and assigns it to the ShippedAt field.
func (o *ReturnDataAttributes) SetShippedAt(v string) {
	o.ShippedAt = &v
}

// GetRejectedAt returns the RejectedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetRejectedAt() string {
	if o == nil || o.RejectedAt == nil {
		var ret string
		return ret
	}
	return *o.RejectedAt
}

// GetRejectedAtOk returns a tuple with the RejectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetRejectedAtOk() (*string, bool) {
	if o == nil || o.RejectedAt == nil {
		return nil, false
	}
	return o.RejectedAt, true
}

// HasRejectedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasRejectedAt() bool {
	if o != nil && o.RejectedAt != nil {
		return true
	}

	return false
}

// SetRejectedAt gets a reference to the given string and assigns it to the RejectedAt field.
func (o *ReturnDataAttributes) SetRejectedAt(v string) {
	o.RejectedAt = &v
}

// GetReceivedAt returns the ReceivedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetReceivedAt() string {
	if o == nil || o.ReceivedAt == nil {
		var ret string
		return ret
	}
	return *o.ReceivedAt
}

// GetReceivedAtOk returns a tuple with the ReceivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetReceivedAtOk() (*string, bool) {
	if o == nil || o.ReceivedAt == nil {
		return nil, false
	}
	return o.ReceivedAt, true
}

// HasReceivedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasReceivedAt() bool {
	if o != nil && o.ReceivedAt != nil {
		return true
	}

	return false
}

// SetReceivedAt gets a reference to the given string and assigns it to the ReceivedAt field.
func (o *ReturnDataAttributes) SetReceivedAt(v string) {
	o.ReceivedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetArchivedAt() string {
	if o == nil || o.ArchivedAt == nil {
		var ret string
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetArchivedAtOk() (*string, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given string and assigns it to the ArchivedAt field.
func (o *ReturnDataAttributes) SetArchivedAt(v string) {
	o.ArchivedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReturnDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ReturnDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ReturnDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ReturnDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ReturnDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ReturnDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ReturnDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ReturnDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ReturnDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.SkusCount != nil {
		toSerialize["skus_count"] = o.SkusCount
	}
	if o.ApprovedAt != nil {
		toSerialize["approved_at"] = o.ApprovedAt
	}
	if o.CancelledAt != nil {
		toSerialize["cancelled_at"] = o.CancelledAt
	}
	if o.ShippedAt != nil {
		toSerialize["shipped_at"] = o.ShippedAt
	}
	if o.RejectedAt != nil {
		toSerialize["rejected_at"] = o.RejectedAt
	}
	if o.ReceivedAt != nil {
		toSerialize["received_at"] = o.ReceivedAt
	}
	if o.ArchivedAt != nil {
		toSerialize["archived_at"] = o.ArchivedAt
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

type NullableReturnDataAttributes struct {
	value *ReturnDataAttributes
	isSet bool
}

func (v NullableReturnDataAttributes) Get() *ReturnDataAttributes {
	return v.value
}

func (v *NullableReturnDataAttributes) Set(val *ReturnDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnDataAttributes(val *ReturnDataAttributes) *NullableReturnDataAttributes {
	return &NullableReturnDataAttributes{value: val, isSet: true}
}

func (v NullableReturnDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
