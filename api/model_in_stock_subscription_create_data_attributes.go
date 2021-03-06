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

// InStockSubscriptionCreateDataAttributes struct for InStockSubscriptionCreateDataAttributes
type InStockSubscriptionCreateDataAttributes struct {
	// The email of the associated customer, replace the relationship
	CustomerEmail *string `json:"customer_email,omitempty"`
	// The code of the associated sku, replace the relationship
	SkuCode *string `json:"sku_code,omitempty"`
	// The threshold at which to trigger the back in stock notification, default to 1.
	StockThreshold *int32 `json:"stock_threshold,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewInStockSubscriptionCreateDataAttributes instantiates a new InStockSubscriptionCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionCreateDataAttributes() *InStockSubscriptionCreateDataAttributes {
	this := InStockSubscriptionCreateDataAttributes{}
	return &this
}

// NewInStockSubscriptionCreateDataAttributesWithDefaults instantiates a new InStockSubscriptionCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionCreateDataAttributesWithDefaults() *InStockSubscriptionCreateDataAttributes {
	this := InStockSubscriptionCreateDataAttributes{}
	return &this
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateDataAttributes) GetCustomerEmail() string {
	if o == nil || o.CustomerEmail == nil {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataAttributes) GetCustomerEmailOk() (*string, bool) {
	if o == nil || o.CustomerEmail == nil {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateDataAttributes) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail != nil {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *InStockSubscriptionCreateDataAttributes) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateDataAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *InStockSubscriptionCreateDataAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetStockThreshold returns the StockThreshold field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateDataAttributes) GetStockThreshold() int32 {
	if o == nil || o.StockThreshold == nil {
		var ret int32
		return ret
	}
	return *o.StockThreshold
}

// GetStockThresholdOk returns a tuple with the StockThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataAttributes) GetStockThresholdOk() (*int32, bool) {
	if o == nil || o.StockThreshold == nil {
		return nil, false
	}
	return o.StockThreshold, true
}

// HasStockThreshold returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateDataAttributes) HasStockThreshold() bool {
	if o != nil && o.StockThreshold != nil {
		return true
	}

	return false
}

// SetStockThreshold gets a reference to the given int32 and assigns it to the StockThreshold field.
func (o *InStockSubscriptionCreateDataAttributes) SetStockThreshold(v int32) {
	o.StockThreshold = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *InStockSubscriptionCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *InStockSubscriptionCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InStockSubscriptionCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o InStockSubscriptionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.StockThreshold != nil {
		toSerialize["stock_threshold"] = o.StockThreshold
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

type NullableInStockSubscriptionCreateDataAttributes struct {
	value *InStockSubscriptionCreateDataAttributes
	isSet bool
}

func (v NullableInStockSubscriptionCreateDataAttributes) Get() *InStockSubscriptionCreateDataAttributes {
	return v.value
}

func (v *NullableInStockSubscriptionCreateDataAttributes) Set(val *InStockSubscriptionCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionCreateDataAttributes(val *InStockSubscriptionCreateDataAttributes) *NullableInStockSubscriptionCreateDataAttributes {
	return &NullableInStockSubscriptionCreateDataAttributes{value: val, isSet: true}
}

func (v NullableInStockSubscriptionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
