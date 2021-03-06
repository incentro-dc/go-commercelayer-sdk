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

// OrderAmountPromotionRuleDataAttributes struct for OrderAmountPromotionRuleDataAttributes
type OrderAmountPromotionRuleDataAttributes struct {
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
	// Apply the promotion only when order is over this amount, in cents.
	OrderAmountCents *int32 `json:"order_amount_cents,omitempty"`
	// Apply the promotion only when order is over this amount, float.
	OrderAmountFloat *float32 `json:"order_amount_float,omitempty"`
	// Apply the promotion only when order is over this amount, formatted.
	FormattedOrderAmount *string `json:"formatted_order_amount,omitempty"`
}

// NewOrderAmountPromotionRuleDataAttributes instantiates a new OrderAmountPromotionRuleDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleDataAttributes() *OrderAmountPromotionRuleDataAttributes {
	this := OrderAmountPromotionRuleDataAttributes{}
	return &this
}

// NewOrderAmountPromotionRuleDataAttributesWithDefaults instantiates a new OrderAmountPromotionRuleDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleDataAttributesWithDefaults() *OrderAmountPromotionRuleDataAttributes {
	this := OrderAmountPromotionRuleDataAttributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderAmountPromotionRuleDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OrderAmountPromotionRuleDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OrderAmountPromotionRuleDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *OrderAmountPromotionRuleDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *OrderAmountPromotionRuleDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *OrderAmountPromotionRuleDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetOrderAmountCents returns the OrderAmountCents field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountCents() int32 {
	if o == nil || o.OrderAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.OrderAmountCents
}

// GetOrderAmountCentsOk returns a tuple with the OrderAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountCentsOk() (*int32, bool) {
	if o == nil || o.OrderAmountCents == nil {
		return nil, false
	}
	return o.OrderAmountCents, true
}

// HasOrderAmountCents returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasOrderAmountCents() bool {
	if o != nil && o.OrderAmountCents != nil {
		return true
	}

	return false
}

// SetOrderAmountCents gets a reference to the given int32 and assigns it to the OrderAmountCents field.
func (o *OrderAmountPromotionRuleDataAttributes) SetOrderAmountCents(v int32) {
	o.OrderAmountCents = &v
}

// GetOrderAmountFloat returns the OrderAmountFloat field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountFloat() float32 {
	if o == nil || o.OrderAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.OrderAmountFloat
}

// GetOrderAmountFloatOk returns a tuple with the OrderAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetOrderAmountFloatOk() (*float32, bool) {
	if o == nil || o.OrderAmountFloat == nil {
		return nil, false
	}
	return o.OrderAmountFloat, true
}

// HasOrderAmountFloat returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasOrderAmountFloat() bool {
	if o != nil && o.OrderAmountFloat != nil {
		return true
	}

	return false
}

// SetOrderAmountFloat gets a reference to the given float32 and assigns it to the OrderAmountFloat field.
func (o *OrderAmountPromotionRuleDataAttributes) SetOrderAmountFloat(v float32) {
	o.OrderAmountFloat = &v
}

// GetFormattedOrderAmount returns the FormattedOrderAmount field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleDataAttributes) GetFormattedOrderAmount() string {
	if o == nil || o.FormattedOrderAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedOrderAmount
}

// GetFormattedOrderAmountOk returns a tuple with the FormattedOrderAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleDataAttributes) GetFormattedOrderAmountOk() (*string, bool) {
	if o == nil || o.FormattedOrderAmount == nil {
		return nil, false
	}
	return o.FormattedOrderAmount, true
}

// HasFormattedOrderAmount returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleDataAttributes) HasFormattedOrderAmount() bool {
	if o != nil && o.FormattedOrderAmount != nil {
		return true
	}

	return false
}

// SetFormattedOrderAmount gets a reference to the given string and assigns it to the FormattedOrderAmount field.
func (o *OrderAmountPromotionRuleDataAttributes) SetFormattedOrderAmount(v string) {
	o.FormattedOrderAmount = &v
}

func (o OrderAmountPromotionRuleDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.OrderAmountCents != nil {
		toSerialize["order_amount_cents"] = o.OrderAmountCents
	}
	if o.OrderAmountFloat != nil {
		toSerialize["order_amount_float"] = o.OrderAmountFloat
	}
	if o.FormattedOrderAmount != nil {
		toSerialize["formatted_order_amount"] = o.FormattedOrderAmount
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleDataAttributes struct {
	value *OrderAmountPromotionRuleDataAttributes
	isSet bool
}

func (v NullableOrderAmountPromotionRuleDataAttributes) Get() *OrderAmountPromotionRuleDataAttributes {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleDataAttributes) Set(val *OrderAmountPromotionRuleDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleDataAttributes(val *OrderAmountPromotionRuleDataAttributes) *NullableOrderAmountPromotionRuleDataAttributes {
	return &NullableOrderAmountPromotionRuleDataAttributes{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
