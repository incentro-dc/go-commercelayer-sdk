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

// ImportDataAttributes struct for ImportDataAttributes
type ImportDataAttributes struct {
	// The type of resource being imported.
	ResourceType *string `json:"resource_type,omitempty"`
	// The ID of the parent resource to be associated with imported data.
	ParentResourceId *string `json:"parent_resource_id,omitempty"`
	// The import job status. One of 'pending' (default), 'in_progress', 'interrputed', or 'completed'.
	Status *string `json:"status,omitempty"`
	// Time at which the import was started.
	StartedAt *string `json:"started_at,omitempty"`
	// Time at which the import was completed.
	CompletedAt *string `json:"completed_at,omitempty"`
	// Time at which the import was interrupted.
	InterruptedAt *string `json:"interrupted_at,omitempty"`
	// Array of objects representing the resources that are being imported.
	Inputs []map[string]interface{} `json:"inputs,omitempty"`
	// Indicates the size of the objects to be imported.
	InputsSize *int32 `json:"inputs_size,omitempty"`
	// Indicates the number of import errors, if any.
	ErrorsCount *int32 `json:"errors_count,omitempty"`
	// Indicates the number of import warnings, if any.
	WarningsCount *int32 `json:"warnings_count,omitempty"`
	// Indicates the number of records that have been destroyed, if any.
	DestroyedCount *int32 `json:"destroyed_count,omitempty"`
	// Indicates the number of records that have been processed (created or updated).
	ProcessedCount *int32 `json:"processed_count,omitempty"`
	// Contains the import errors, if any.
	ErrorsLog map[string]interface{} `json:"errors_log,omitempty"`
	// Contains the import warnings, if any.
	WarningsLog map[string]interface{} `json:"warnings_log,omitempty"`
	// Indicates if the import should cleanup records that are not included in the inputs array.
	CleanupRecords *bool `json:"cleanup_records,omitempty"`
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

// NewImportDataAttributes instantiates a new ImportDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportDataAttributes() *ImportDataAttributes {
	this := ImportDataAttributes{}
	return &this
}

// NewImportDataAttributesWithDefaults instantiates a new ImportDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportDataAttributesWithDefaults() *ImportDataAttributes {
	this := ImportDataAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ImportDataAttributes) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetParentResourceId returns the ParentResourceId field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetParentResourceId() string {
	if o == nil || o.ParentResourceId == nil {
		var ret string
		return ret
	}
	return *o.ParentResourceId
}

// GetParentResourceIdOk returns a tuple with the ParentResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetParentResourceIdOk() (*string, bool) {
	if o == nil || o.ParentResourceId == nil {
		return nil, false
	}
	return o.ParentResourceId, true
}

// HasParentResourceId returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasParentResourceId() bool {
	if o != nil && o.ParentResourceId != nil {
		return true
	}

	return false
}

// SetParentResourceId gets a reference to the given string and assigns it to the ParentResourceId field.
func (o *ImportDataAttributes) SetParentResourceId(v string) {
	o.ParentResourceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImportDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetStartedAt() string {
	if o == nil || o.StartedAt == nil {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetStartedAtOk() (*string, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *ImportDataAttributes) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetCompletedAt() string {
	if o == nil || o.CompletedAt == nil {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetCompletedAtOk() (*string, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *ImportDataAttributes) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetInterruptedAt returns the InterruptedAt field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetInterruptedAt() string {
	if o == nil || o.InterruptedAt == nil {
		var ret string
		return ret
	}
	return *o.InterruptedAt
}

// GetInterruptedAtOk returns a tuple with the InterruptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetInterruptedAtOk() (*string, bool) {
	if o == nil || o.InterruptedAt == nil {
		return nil, false
	}
	return o.InterruptedAt, true
}

// HasInterruptedAt returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasInterruptedAt() bool {
	if o != nil && o.InterruptedAt != nil {
		return true
	}

	return false
}

// SetInterruptedAt gets a reference to the given string and assigns it to the InterruptedAt field.
func (o *ImportDataAttributes) SetInterruptedAt(v string) {
	o.InterruptedAt = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetInputs() []map[string]interface{} {
	if o == nil || o.Inputs == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetInputsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []map[string]interface{} and assigns it to the Inputs field.
func (o *ImportDataAttributes) SetInputs(v []map[string]interface{}) {
	o.Inputs = v
}

// GetInputsSize returns the InputsSize field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetInputsSize() int32 {
	if o == nil || o.InputsSize == nil {
		var ret int32
		return ret
	}
	return *o.InputsSize
}

// GetInputsSizeOk returns a tuple with the InputsSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetInputsSizeOk() (*int32, bool) {
	if o == nil || o.InputsSize == nil {
		return nil, false
	}
	return o.InputsSize, true
}

// HasInputsSize returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasInputsSize() bool {
	if o != nil && o.InputsSize != nil {
		return true
	}

	return false
}

// SetInputsSize gets a reference to the given int32 and assigns it to the InputsSize field.
func (o *ImportDataAttributes) SetInputsSize(v int32) {
	o.InputsSize = &v
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetErrorsCount() int32 {
	if o == nil || o.ErrorsCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorsCount
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetErrorsCountOk() (*int32, bool) {
	if o == nil || o.ErrorsCount == nil {
		return nil, false
	}
	return o.ErrorsCount, true
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasErrorsCount() bool {
	if o != nil && o.ErrorsCount != nil {
		return true
	}

	return false
}

// SetErrorsCount gets a reference to the given int32 and assigns it to the ErrorsCount field.
func (o *ImportDataAttributes) SetErrorsCount(v int32) {
	o.ErrorsCount = &v
}

// GetWarningsCount returns the WarningsCount field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetWarningsCount() int32 {
	if o == nil || o.WarningsCount == nil {
		var ret int32
		return ret
	}
	return *o.WarningsCount
}

// GetWarningsCountOk returns a tuple with the WarningsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetWarningsCountOk() (*int32, bool) {
	if o == nil || o.WarningsCount == nil {
		return nil, false
	}
	return o.WarningsCount, true
}

// HasWarningsCount returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasWarningsCount() bool {
	if o != nil && o.WarningsCount != nil {
		return true
	}

	return false
}

// SetWarningsCount gets a reference to the given int32 and assigns it to the WarningsCount field.
func (o *ImportDataAttributes) SetWarningsCount(v int32) {
	o.WarningsCount = &v
}

// GetDestroyedCount returns the DestroyedCount field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetDestroyedCount() int32 {
	if o == nil || o.DestroyedCount == nil {
		var ret int32
		return ret
	}
	return *o.DestroyedCount
}

// GetDestroyedCountOk returns a tuple with the DestroyedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetDestroyedCountOk() (*int32, bool) {
	if o == nil || o.DestroyedCount == nil {
		return nil, false
	}
	return o.DestroyedCount, true
}

// HasDestroyedCount returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasDestroyedCount() bool {
	if o != nil && o.DestroyedCount != nil {
		return true
	}

	return false
}

// SetDestroyedCount gets a reference to the given int32 and assigns it to the DestroyedCount field.
func (o *ImportDataAttributes) SetDestroyedCount(v int32) {
	o.DestroyedCount = &v
}

// GetProcessedCount returns the ProcessedCount field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetProcessedCount() int32 {
	if o == nil || o.ProcessedCount == nil {
		var ret int32
		return ret
	}
	return *o.ProcessedCount
}

// GetProcessedCountOk returns a tuple with the ProcessedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetProcessedCountOk() (*int32, bool) {
	if o == nil || o.ProcessedCount == nil {
		return nil, false
	}
	return o.ProcessedCount, true
}

// HasProcessedCount returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasProcessedCount() bool {
	if o != nil && o.ProcessedCount != nil {
		return true
	}

	return false
}

// SetProcessedCount gets a reference to the given int32 and assigns it to the ProcessedCount field.
func (o *ImportDataAttributes) SetProcessedCount(v int32) {
	o.ProcessedCount = &v
}

// GetErrorsLog returns the ErrorsLog field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetErrorsLog() map[string]interface{} {
	if o == nil || o.ErrorsLog == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ErrorsLog
}

// GetErrorsLogOk returns a tuple with the ErrorsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetErrorsLogOk() (map[string]interface{}, bool) {
	if o == nil || o.ErrorsLog == nil {
		return nil, false
	}
	return o.ErrorsLog, true
}

// HasErrorsLog returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasErrorsLog() bool {
	if o != nil && o.ErrorsLog != nil {
		return true
	}

	return false
}

// SetErrorsLog gets a reference to the given map[string]interface{} and assigns it to the ErrorsLog field.
func (o *ImportDataAttributes) SetErrorsLog(v map[string]interface{}) {
	o.ErrorsLog = v
}

// GetWarningsLog returns the WarningsLog field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetWarningsLog() map[string]interface{} {
	if o == nil || o.WarningsLog == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.WarningsLog
}

// GetWarningsLogOk returns a tuple with the WarningsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetWarningsLogOk() (map[string]interface{}, bool) {
	if o == nil || o.WarningsLog == nil {
		return nil, false
	}
	return o.WarningsLog, true
}

// HasWarningsLog returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasWarningsLog() bool {
	if o != nil && o.WarningsLog != nil {
		return true
	}

	return false
}

// SetWarningsLog gets a reference to the given map[string]interface{} and assigns it to the WarningsLog field.
func (o *ImportDataAttributes) SetWarningsLog(v map[string]interface{}) {
	o.WarningsLog = v
}

// GetCleanupRecords returns the CleanupRecords field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetCleanupRecords() bool {
	if o == nil || o.CleanupRecords == nil {
		var ret bool
		return ret
	}
	return *o.CleanupRecords
}

// GetCleanupRecordsOk returns a tuple with the CleanupRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetCleanupRecordsOk() (*bool, bool) {
	if o == nil || o.CleanupRecords == nil {
		return nil, false
	}
	return o.CleanupRecords, true
}

// HasCleanupRecords returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasCleanupRecords() bool {
	if o != nil && o.CleanupRecords != nil {
		return true
	}

	return false
}

// SetCleanupRecords gets a reference to the given bool and assigns it to the CleanupRecords field.
func (o *ImportDataAttributes) SetCleanupRecords(v bool) {
	o.CleanupRecords = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImportDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ImportDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ImportDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ImportDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ImportDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImportDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImportDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ImportDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ImportDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.ParentResourceId != nil {
		toSerialize["parent_resource_id"] = o.ParentResourceId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.CompletedAt != nil {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if o.InterruptedAt != nil {
		toSerialize["interrupted_at"] = o.InterruptedAt
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if o.InputsSize != nil {
		toSerialize["inputs_size"] = o.InputsSize
	}
	if o.ErrorsCount != nil {
		toSerialize["errors_count"] = o.ErrorsCount
	}
	if o.WarningsCount != nil {
		toSerialize["warnings_count"] = o.WarningsCount
	}
	if o.DestroyedCount != nil {
		toSerialize["destroyed_count"] = o.DestroyedCount
	}
	if o.ProcessedCount != nil {
		toSerialize["processed_count"] = o.ProcessedCount
	}
	if o.ErrorsLog != nil {
		toSerialize["errors_log"] = o.ErrorsLog
	}
	if o.WarningsLog != nil {
		toSerialize["warnings_log"] = o.WarningsLog
	}
	if o.CleanupRecords != nil {
		toSerialize["cleanup_records"] = o.CleanupRecords
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

type NullableImportDataAttributes struct {
	value *ImportDataAttributes
	isSet bool
}

func (v NullableImportDataAttributes) Get() *ImportDataAttributes {
	return v.value
}

func (v *NullableImportDataAttributes) Set(val *ImportDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableImportDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableImportDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportDataAttributes(val *ImportDataAttributes) *NullableImportDataAttributes {
	return &NullableImportDataAttributes{value: val, isSet: true}
}

func (v NullableImportDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
