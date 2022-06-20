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

// ShippingMethodDataAttributes struct for ShippingMethodDataAttributes
type ShippingMethodDataAttributes struct {
	// The shipping method's name
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Time at which the shipping method was disabled.
	DisabledAt *string `json:"disabled_at,omitempty"`
	// The price of this shipping method, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// The price of this shipping method, float.
	PriceAmountFloat *float32 `json:"price_amount_float,omitempty"`
	// The price of this shipping method, formatted.
	FormattedPriceAmount *string `json:"formatted_price_amount,omitempty"`
	// Apply free shipping if the order amount is over this value, in cents.
	FreeOverAmountCents *int32 `json:"free_over_amount_cents,omitempty"`
	// Apply free shipping if the order amount is over this value, float.
	FreeOverAmountFloat *float32 `json:"free_over_amount_float,omitempty"`
	// Apply free shipping if the order amount is over this value, formatted.
	FormattedFreeOverAmount *string `json:"formatted_free_over_amount,omitempty"`
	// The calculated price (zero or price amount) when associated to a shipment, in cents.
	PriceAmountForShipmentCents *int32 `json:"price_amount_for_shipment_cents,omitempty"`
	// The calculated price (zero or price amount) when associated to a shipment, float.
	PriceAmountForShipmentFloat *float32 `json:"price_amount_for_shipment_float,omitempty"`
	// The calculated price (zero or price amount) when associated to a shipment, formatted.
	FormattedPriceAmountForShipment *string `json:"formatted_price_amount_for_shipment,omitempty"`
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

// NewShippingMethodDataAttributes instantiates a new ShippingMethodDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodDataAttributes() *ShippingMethodDataAttributes {
	this := ShippingMethodDataAttributes{}
	return &this
}

// NewShippingMethodDataAttributesWithDefaults instantiates a new ShippingMethodDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodDataAttributesWithDefaults() *ShippingMethodDataAttributes {
	this := ShippingMethodDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingMethodDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ShippingMethodDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetDisabledAt() string {
	if o == nil || o.DisabledAt == nil {
		var ret string
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetDisabledAtOk() (*string, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasDisabledAt() bool {
	if o != nil && o.DisabledAt != nil {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given string and assigns it to the DisabledAt field.
func (o *ShippingMethodDataAttributes) SetDisabledAt(v string) {
	o.DisabledAt = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *ShippingMethodDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetPriceAmountFloat returns the PriceAmountFloat field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetPriceAmountFloat() float32 {
	if o == nil || o.PriceAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.PriceAmountFloat
}

// GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetPriceAmountFloatOk() (*float32, bool) {
	if o == nil || o.PriceAmountFloat == nil {
		return nil, false
	}
	return o.PriceAmountFloat, true
}

// HasPriceAmountFloat returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasPriceAmountFloat() bool {
	if o != nil && o.PriceAmountFloat != nil {
		return true
	}

	return false
}

// SetPriceAmountFloat gets a reference to the given float32 and assigns it to the PriceAmountFloat field.
func (o *ShippingMethodDataAttributes) SetPriceAmountFloat(v float32) {
	o.PriceAmountFloat = &v
}

// GetFormattedPriceAmount returns the FormattedPriceAmount field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetFormattedPriceAmount() string {
	if o == nil || o.FormattedPriceAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedPriceAmount
}

// GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetFormattedPriceAmountOk() (*string, bool) {
	if o == nil || o.FormattedPriceAmount == nil {
		return nil, false
	}
	return o.FormattedPriceAmount, true
}

// HasFormattedPriceAmount returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasFormattedPriceAmount() bool {
	if o != nil && o.FormattedPriceAmount != nil {
		return true
	}

	return false
}

// SetFormattedPriceAmount gets a reference to the given string and assigns it to the FormattedPriceAmount field.
func (o *ShippingMethodDataAttributes) SetFormattedPriceAmount(v string) {
	o.FormattedPriceAmount = &v
}

// GetFreeOverAmountCents returns the FreeOverAmountCents field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetFreeOverAmountCents() int32 {
	if o == nil || o.FreeOverAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.FreeOverAmountCents
}

// GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetFreeOverAmountCentsOk() (*int32, bool) {
	if o == nil || o.FreeOverAmountCents == nil {
		return nil, false
	}
	return o.FreeOverAmountCents, true
}

// HasFreeOverAmountCents returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasFreeOverAmountCents() bool {
	if o != nil && o.FreeOverAmountCents != nil {
		return true
	}

	return false
}

// SetFreeOverAmountCents gets a reference to the given int32 and assigns it to the FreeOverAmountCents field.
func (o *ShippingMethodDataAttributes) SetFreeOverAmountCents(v int32) {
	o.FreeOverAmountCents = &v
}

// GetFreeOverAmountFloat returns the FreeOverAmountFloat field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetFreeOverAmountFloat() float32 {
	if o == nil || o.FreeOverAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.FreeOverAmountFloat
}

// GetFreeOverAmountFloatOk returns a tuple with the FreeOverAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetFreeOverAmountFloatOk() (*float32, bool) {
	if o == nil || o.FreeOverAmountFloat == nil {
		return nil, false
	}
	return o.FreeOverAmountFloat, true
}

// HasFreeOverAmountFloat returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasFreeOverAmountFloat() bool {
	if o != nil && o.FreeOverAmountFloat != nil {
		return true
	}

	return false
}

// SetFreeOverAmountFloat gets a reference to the given float32 and assigns it to the FreeOverAmountFloat field.
func (o *ShippingMethodDataAttributes) SetFreeOverAmountFloat(v float32) {
	o.FreeOverAmountFloat = &v
}

// GetFormattedFreeOverAmount returns the FormattedFreeOverAmount field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetFormattedFreeOverAmount() string {
	if o == nil || o.FormattedFreeOverAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedFreeOverAmount
}

// GetFormattedFreeOverAmountOk returns a tuple with the FormattedFreeOverAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetFormattedFreeOverAmountOk() (*string, bool) {
	if o == nil || o.FormattedFreeOverAmount == nil {
		return nil, false
	}
	return o.FormattedFreeOverAmount, true
}

// HasFormattedFreeOverAmount returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasFormattedFreeOverAmount() bool {
	if o != nil && o.FormattedFreeOverAmount != nil {
		return true
	}

	return false
}

// SetFormattedFreeOverAmount gets a reference to the given string and assigns it to the FormattedFreeOverAmount field.
func (o *ShippingMethodDataAttributes) SetFormattedFreeOverAmount(v string) {
	o.FormattedFreeOverAmount = &v
}

// GetPriceAmountForShipmentCents returns the PriceAmountForShipmentCents field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentCents() int32 {
	if o == nil || o.PriceAmountForShipmentCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountForShipmentCents
}

// GetPriceAmountForShipmentCentsOk returns a tuple with the PriceAmountForShipmentCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountForShipmentCents == nil {
		return nil, false
	}
	return o.PriceAmountForShipmentCents, true
}

// HasPriceAmountForShipmentCents returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasPriceAmountForShipmentCents() bool {
	if o != nil && o.PriceAmountForShipmentCents != nil {
		return true
	}

	return false
}

// SetPriceAmountForShipmentCents gets a reference to the given int32 and assigns it to the PriceAmountForShipmentCents field.
func (o *ShippingMethodDataAttributes) SetPriceAmountForShipmentCents(v int32) {
	o.PriceAmountForShipmentCents = &v
}

// GetPriceAmountForShipmentFloat returns the PriceAmountForShipmentFloat field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentFloat() float32 {
	if o == nil || o.PriceAmountForShipmentFloat == nil {
		var ret float32
		return ret
	}
	return *o.PriceAmountForShipmentFloat
}

// GetPriceAmountForShipmentFloatOk returns a tuple with the PriceAmountForShipmentFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetPriceAmountForShipmentFloatOk() (*float32, bool) {
	if o == nil || o.PriceAmountForShipmentFloat == nil {
		return nil, false
	}
	return o.PriceAmountForShipmentFloat, true
}

// HasPriceAmountForShipmentFloat returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasPriceAmountForShipmentFloat() bool {
	if o != nil && o.PriceAmountForShipmentFloat != nil {
		return true
	}

	return false
}

// SetPriceAmountForShipmentFloat gets a reference to the given float32 and assigns it to the PriceAmountForShipmentFloat field.
func (o *ShippingMethodDataAttributes) SetPriceAmountForShipmentFloat(v float32) {
	o.PriceAmountForShipmentFloat = &v
}

// GetFormattedPriceAmountForShipment returns the FormattedPriceAmountForShipment field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetFormattedPriceAmountForShipment() string {
	if o == nil || o.FormattedPriceAmountForShipment == nil {
		var ret string
		return ret
	}
	return *o.FormattedPriceAmountForShipment
}

// GetFormattedPriceAmountForShipmentOk returns a tuple with the FormattedPriceAmountForShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetFormattedPriceAmountForShipmentOk() (*string, bool) {
	if o == nil || o.FormattedPriceAmountForShipment == nil {
		return nil, false
	}
	return o.FormattedPriceAmountForShipment, true
}

// HasFormattedPriceAmountForShipment returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasFormattedPriceAmountForShipment() bool {
	if o != nil && o.FormattedPriceAmountForShipment != nil {
		return true
	}

	return false
}

// SetFormattedPriceAmountForShipment gets a reference to the given string and assigns it to the FormattedPriceAmountForShipment field.
func (o *ShippingMethodDataAttributes) SetFormattedPriceAmountForShipment(v string) {
	o.FormattedPriceAmountForShipment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShippingMethodDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ShippingMethodDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ShippingMethodDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ShippingMethodDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ShippingMethodDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShippingMethodDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShippingMethodDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ShippingMethodDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ShippingMethodDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.PriceAmountFloat != nil {
		toSerialize["price_amount_float"] = o.PriceAmountFloat
	}
	if o.FormattedPriceAmount != nil {
		toSerialize["formatted_price_amount"] = o.FormattedPriceAmount
	}
	if o.FreeOverAmountCents != nil {
		toSerialize["free_over_amount_cents"] = o.FreeOverAmountCents
	}
	if o.FreeOverAmountFloat != nil {
		toSerialize["free_over_amount_float"] = o.FreeOverAmountFloat
	}
	if o.FormattedFreeOverAmount != nil {
		toSerialize["formatted_free_over_amount"] = o.FormattedFreeOverAmount
	}
	if o.PriceAmountForShipmentCents != nil {
		toSerialize["price_amount_for_shipment_cents"] = o.PriceAmountForShipmentCents
	}
	if o.PriceAmountForShipmentFloat != nil {
		toSerialize["price_amount_for_shipment_float"] = o.PriceAmountForShipmentFloat
	}
	if o.FormattedPriceAmountForShipment != nil {
		toSerialize["formatted_price_amount_for_shipment"] = o.FormattedPriceAmountForShipment
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

type NullableShippingMethodDataAttributes struct {
	value *ShippingMethodDataAttributes
	isSet bool
}

func (v NullableShippingMethodDataAttributes) Get() *ShippingMethodDataAttributes {
	return v.value
}

func (v *NullableShippingMethodDataAttributes) Set(val *ShippingMethodDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodDataAttributes(val *ShippingMethodDataAttributes) *NullableShippingMethodDataAttributes {
	return &NullableShippingMethodDataAttributes{value: val, isSet: true}
}

func (v NullableShippingMethodDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
