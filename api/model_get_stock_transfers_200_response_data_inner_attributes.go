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

// GETStockTransfers200ResponseDataInnerAttributes struct for GETStockTransfers200ResponseDataInnerAttributes
type GETStockTransfers200ResponseDataInnerAttributes struct {
	// The code of the associated SKU.
	SkuCode *string `json:"sku_code,omitempty"`
	// The stock transfer status, one of 'draft', 'upcoming', 'picking', 'in_transit', 'completed', or 'cancelled'
	Status *string `json:"status,omitempty"`
	// The stock quantity to be transferred from the origin stock location to destination one
	Quantity *int32 `json:"quantity,omitempty"`
	// Time at which the stock transfer was completed.
	CompletedAt *string `json:"completed_at,omitempty"`
	// Time at which the stock transfer was cancelled.
	CancelledAt *string `json:"cancelled_at,omitempty"`
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

// NewGETStockTransfers200ResponseDataInnerAttributes instantiates a new GETStockTransfers200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfers200ResponseDataInnerAttributes() *GETStockTransfers200ResponseDataInnerAttributes {
	this := GETStockTransfers200ResponseDataInnerAttributes{}
	return &this
}

// NewGETStockTransfers200ResponseDataInnerAttributesWithDefaults instantiates a new GETStockTransfers200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfers200ResponseDataInnerAttributesWithDefaults() *GETStockTransfers200ResponseDataInnerAttributes {
	this := GETStockTransfers200ResponseDataInnerAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetCompletedAt() string {
	if o == nil || o.CompletedAt == nil {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetCompletedAtOk() (*string, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetCancelledAt returns the CancelledAt field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetCancelledAt() string {
	if o == nil || o.CancelledAt == nil {
		var ret string
		return ret
	}
	return *o.CancelledAt
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetCancelledAtOk() (*string, bool) {
	if o == nil || o.CancelledAt == nil {
		return nil, false
	}
	return o.CancelledAt, true
}

// HasCancelledAt returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasCancelledAt() bool {
	if o != nil && o.CancelledAt != nil {
		return true
	}

	return false
}

// SetCancelledAt gets a reference to the given string and assigns it to the CancelledAt field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetCancelledAt(v string) {
	o.CancelledAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETStockTransfers200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETStockTransfers200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.CompletedAt != nil {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if o.CancelledAt != nil {
		toSerialize["cancelled_at"] = o.CancelledAt
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

type NullableGETStockTransfers200ResponseDataInnerAttributes struct {
	value *GETStockTransfers200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETStockTransfers200ResponseDataInnerAttributes) Get() *GETStockTransfers200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETStockTransfers200ResponseDataInnerAttributes) Set(val *GETStockTransfers200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfers200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfers200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfers200ResponseDataInnerAttributes(val *GETStockTransfers200ResponseDataInnerAttributes) *NullableGETStockTransfers200ResponseDataInnerAttributes {
	return &NullableGETStockTransfers200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETStockTransfers200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfers200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
