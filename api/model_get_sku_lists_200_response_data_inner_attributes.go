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

// GETSkuLists200ResponseDataInnerAttributes struct for GETSkuLists200ResponseDataInnerAttributes
type GETSkuLists200ResponseDataInnerAttributes struct {
	// The SKU list's internal name.
	Name interface{} `json:"name,omitempty"`
	// The SKU list's internal slug.
	Slug interface{} `json:"slug,omitempty"`
	// An internal description of the SKU list.
	Description interface{} `json:"description,omitempty"`
	// The URL of an image that represents the SKU list.
	ImageUrl interface{} `json:"image_url,omitempty"`
	// Indicates if the SKU list is populated manually.
	Manual interface{} `json:"manual,omitempty"`
	// The regex that will be evaluated to match the SKU codes.
	SkuCodeRegex interface{} `json:"sku_code_regex,omitempty"`
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

// NewGETSkuLists200ResponseDataInnerAttributes instantiates a new GETSkuLists200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuLists200ResponseDataInnerAttributes() *GETSkuLists200ResponseDataInnerAttributes {
	this := GETSkuLists200ResponseDataInnerAttributes{}
	return &this
}

// NewGETSkuLists200ResponseDataInnerAttributesWithDefaults instantiates a new GETSkuLists200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuLists200ResponseDataInnerAttributesWithDefaults() *GETSkuLists200ResponseDataInnerAttributes {
	this := GETSkuLists200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetSlug() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetSlugOk() (*interface{}, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return &o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given interface{} and assigns it to the Slug field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetSlug(v interface{}) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetDescription(v interface{}) {
	o.Description = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetManual returns the Manual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetManual() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetManualOk() (*interface{}, bool) {
	if o == nil || o.Manual == nil {
		return nil, false
	}
	return &o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasManual() bool {
	if o != nil && o.Manual != nil {
		return true
	}

	return false
}

// SetManual gets a reference to the given interface{} and assigns it to the Manual field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetManual(v interface{}) {
	o.Manual = v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetSkuCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetSkuCodeRegexOk() (*interface{}, bool) {
	if o == nil || o.SkuCodeRegex == nil {
		return nil, false
	}
	return &o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasSkuCodeRegex() bool {
	if o != nil && o.SkuCodeRegex != nil {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given interface{} and assigns it to the SkuCodeRegex field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetSkuCodeRegex(v interface{}) {
	o.SkuCodeRegex = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkuLists200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkuLists200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETSkuLists200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETSkuLists200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Manual != nil {
		toSerialize["manual"] = o.Manual
	}
	if o.SkuCodeRegex != nil {
		toSerialize["sku_code_regex"] = o.SkuCodeRegex
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

type NullableGETSkuLists200ResponseDataInnerAttributes struct {
	value *GETSkuLists200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETSkuLists200ResponseDataInnerAttributes) Get() *GETSkuLists200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETSkuLists200ResponseDataInnerAttributes) Set(val *GETSkuLists200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuLists200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuLists200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuLists200ResponseDataInnerAttributes(val *GETSkuLists200ResponseDataInnerAttributes) *NullableGETSkuLists200ResponseDataInnerAttributes {
	return &NullableGETSkuLists200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETSkuLists200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuLists200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
