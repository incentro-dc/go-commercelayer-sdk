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

// GETExports200ResponseDataInnerAttributes struct for GETExports200ResponseDataInnerAttributes
type GETExports200ResponseDataInnerAttributes struct {
	// The type of resource being exported.
	ResourceType interface{} `json:"resource_type,omitempty"`
	// The format of the export one of 'json' (default) or 'csv'.
	Format interface{} `json:"format,omitempty"`
	// The export job status. One of 'pending' (default), 'in_progress', or 'completed'.
	Status interface{} `json:"status,omitempty"`
	// List of related resources that should be included in the export.
	Includes []interface{} `json:"includes,omitempty"`
	// The filters used to select the records to be exported.
	Filters interface{} `json:"filters,omitempty"`
	// Send this attribute if you want to skip exporting redundant attributes (IDs, timestamps, blanks, etc.), useful when combining export and import to duplicate your dataset.
	DryData interface{} `json:"dry_data,omitempty"`
	// Time at which the export was started.
	StartedAt interface{} `json:"started_at,omitempty"`
	// Time at which the export was completed.
	CompletedAt interface{} `json:"completed_at,omitempty"`
	// Time at which the export was interrupted.
	InterruptedAt interface{} `json:"interrupted_at,omitempty"`
	// Indicates the number of records to be exported.
	RecordsCount interface{} `json:"records_count,omitempty"`
	// The URL to the output file, which will be generated upon export completion.
	AttachmentUrl interface{} `json:"attachment_url,omitempty"`
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

// NewGETExports200ResponseDataInnerAttributes instantiates a new GETExports200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExports200ResponseDataInnerAttributes() *GETExports200ResponseDataInnerAttributes {
	this := GETExports200ResponseDataInnerAttributes{}
	return &this
}

// NewGETExports200ResponseDataInnerAttributesWithDefaults instantiates a new GETExports200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExports200ResponseDataInnerAttributesWithDefaults() *GETExports200ResponseDataInnerAttributes {
	this := GETExports200ResponseDataInnerAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetResourceType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetResourceTypeOk() (*interface{}, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given interface{} and assigns it to the ResourceType field.
func (o *GETExports200ResponseDataInnerAttributes) SetResourceType(v interface{}) {
	o.ResourceType = v
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetFormat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetFormatOk() (*interface{}, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return &o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given interface{} and assigns it to the Format field.
func (o *GETExports200ResponseDataInnerAttributes) SetFormat(v interface{}) {
	o.Format = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETExports200ResponseDataInnerAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetIncludes returns the Includes field value if set, zero value otherwise.
func (o *GETExports200ResponseDataInnerAttributes) GetIncludes() []interface{} {
	if o == nil || o.Includes == nil {
		var ret []interface{}
		return ret
	}
	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExports200ResponseDataInnerAttributes) GetIncludesOk() ([]interface{}, bool) {
	if o == nil || o.Includes == nil {
		return nil, false
	}
	return o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasIncludes() bool {
	if o != nil && o.Includes != nil {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given []interface{} and assigns it to the Includes field.
func (o *GETExports200ResponseDataInnerAttributes) SetIncludes(v []interface{}) {
	o.Includes = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetFilters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetFiltersOk() (*interface{}, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *GETExports200ResponseDataInnerAttributes) SetFilters(v interface{}) {
	o.Filters = v
}

// GetDryData returns the DryData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetDryData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DryData
}

// GetDryDataOk returns a tuple with the DryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetDryDataOk() (*interface{}, bool) {
	if o == nil || o.DryData == nil {
		return nil, false
	}
	return &o.DryData, true
}

// HasDryData returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasDryData() bool {
	if o != nil && o.DryData != nil {
		return true
	}

	return false
}

// SetDryData gets a reference to the given interface{} and assigns it to the DryData field.
func (o *GETExports200ResponseDataInnerAttributes) SetDryData(v interface{}) {
	o.DryData = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetStartedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetStartedAtOk() (*interface{}, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given interface{} and assigns it to the StartedAt field.
func (o *GETExports200ResponseDataInnerAttributes) SetStartedAt(v interface{}) {
	o.StartedAt = v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetCompletedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetCompletedAtOk() (*interface{}, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given interface{} and assigns it to the CompletedAt field.
func (o *GETExports200ResponseDataInnerAttributes) SetCompletedAt(v interface{}) {
	o.CompletedAt = v
}

// GetInterruptedAt returns the InterruptedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetInterruptedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InterruptedAt
}

// GetInterruptedAtOk returns a tuple with the InterruptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetInterruptedAtOk() (*interface{}, bool) {
	if o == nil || o.InterruptedAt == nil {
		return nil, false
	}
	return &o.InterruptedAt, true
}

// HasInterruptedAt returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasInterruptedAt() bool {
	if o != nil && o.InterruptedAt != nil {
		return true
	}

	return false
}

// SetInterruptedAt gets a reference to the given interface{} and assigns it to the InterruptedAt field.
func (o *GETExports200ResponseDataInnerAttributes) SetInterruptedAt(v interface{}) {
	o.InterruptedAt = v
}

// GetRecordsCount returns the RecordsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetRecordsCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RecordsCount
}

// GetRecordsCountOk returns a tuple with the RecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetRecordsCountOk() (*interface{}, bool) {
	if o == nil || o.RecordsCount == nil {
		return nil, false
	}
	return &o.RecordsCount, true
}

// HasRecordsCount returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasRecordsCount() bool {
	if o != nil && o.RecordsCount != nil {
		return true
	}

	return false
}

// SetRecordsCount gets a reference to the given interface{} and assigns it to the RecordsCount field.
func (o *GETExports200ResponseDataInnerAttributes) SetRecordsCount(v interface{}) {
	o.RecordsCount = v
}

// GetAttachmentUrl returns the AttachmentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetAttachmentUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AttachmentUrl
}

// GetAttachmentUrlOk returns a tuple with the AttachmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetAttachmentUrlOk() (*interface{}, bool) {
	if o == nil || o.AttachmentUrl == nil {
		return nil, false
	}
	return &o.AttachmentUrl, true
}

// HasAttachmentUrl returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasAttachmentUrl() bool {
	if o != nil && o.AttachmentUrl != nil {
		return true
	}

	return false
}

// SetAttachmentUrl gets a reference to the given interface{} and assigns it to the AttachmentUrl field.
func (o *GETExports200ResponseDataInnerAttributes) SetAttachmentUrl(v interface{}) {
	o.AttachmentUrl = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETExports200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETExports200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETExports200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETExports200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExports200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExports200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETExports200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETExports200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETExports200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Includes != nil {
		toSerialize["includes"] = o.Includes
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.DryData != nil {
		toSerialize["dry_data"] = o.DryData
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
	if o.RecordsCount != nil {
		toSerialize["records_count"] = o.RecordsCount
	}
	if o.AttachmentUrl != nil {
		toSerialize["attachment_url"] = o.AttachmentUrl
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

type NullableGETExports200ResponseDataInnerAttributes struct {
	value *GETExports200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETExports200ResponseDataInnerAttributes) Get() *GETExports200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETExports200ResponseDataInnerAttributes) Set(val *GETExports200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExports200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExports200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExports200ResponseDataInnerAttributes(val *GETExports200ResponseDataInnerAttributes) *NullableGETExports200ResponseDataInnerAttributes {
	return &NullableGETExports200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETExports200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExports200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
