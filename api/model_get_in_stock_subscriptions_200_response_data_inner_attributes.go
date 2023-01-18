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

// GETInStockSubscriptions200ResponseDataInnerAttributes struct for GETInStockSubscriptions200ResponseDataInnerAttributes
type GETInStockSubscriptions200ResponseDataInnerAttributes struct {
	// The subscription status. One of 'active' (default), 'inactive', or 'notified'
	Status *string `json:"status,omitempty"`
	// The email of the associated customer, replace the relationship
	CustomerEmail *string `json:"customer_email,omitempty"`
	// The code of the associated SKU, replace the relationship
	SkuCode *string `json:"sku_code,omitempty"`
	// The threshold at which to trigger the back in stock notification, default to 1.
	StockThreshold *int32 `json:"stock_threshold,omitempty"`
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

// NewGETInStockSubscriptions200ResponseDataInnerAttributes instantiates a new GETInStockSubscriptions200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInStockSubscriptions200ResponseDataInnerAttributes() *GETInStockSubscriptions200ResponseDataInnerAttributes {
	this := GETInStockSubscriptions200ResponseDataInnerAttributes{}
	return &this
}

// NewGETInStockSubscriptions200ResponseDataInnerAttributesWithDefaults instantiates a new GETInStockSubscriptions200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInStockSubscriptions200ResponseDataInnerAttributesWithDefaults() *GETInStockSubscriptions200ResponseDataInnerAttributes {
	this := GETInStockSubscriptions200ResponseDataInnerAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetCustomerEmail() string {
	if o == nil || o.CustomerEmail == nil {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetCustomerEmailOk() (*string, bool) {
	if o == nil || o.CustomerEmail == nil {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail != nil {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetStockThreshold returns the StockThreshold field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetStockThreshold() int32 {
	if o == nil || o.StockThreshold == nil {
		var ret int32
		return ret
	}
	return *o.StockThreshold
}

// GetStockThresholdOk returns a tuple with the StockThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetStockThresholdOk() (*int32, bool) {
	if o == nil || o.StockThreshold == nil {
		return nil, false
	}
	return o.StockThreshold, true
}

// HasStockThreshold returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasStockThreshold() bool {
	if o != nil && o.StockThreshold != nil {
		return true
	}

	return false
}

// SetStockThreshold gets a reference to the given int32 and assigns it to the StockThreshold field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetStockThreshold(v int32) {
	o.StockThreshold = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETInStockSubscriptions200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETInStockSubscriptions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.StockThreshold != nil {
		toSerialize["stock_threshold"] = o.StockThreshold
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

type NullableGETInStockSubscriptions200ResponseDataInnerAttributes struct {
	value *GETInStockSubscriptions200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerAttributes) Get() *GETInStockSubscriptions200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerAttributes) Set(val *GETInStockSubscriptions200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInStockSubscriptions200ResponseDataInnerAttributes(val *GETInStockSubscriptions200ResponseDataInnerAttributes) *NullableGETInStockSubscriptions200ResponseDataInnerAttributes {
	return &NullableGETInStockSubscriptions200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
