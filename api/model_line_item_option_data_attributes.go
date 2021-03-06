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

// LineItemOptionDataAttributes struct for LineItemOptionDataAttributes
type LineItemOptionDataAttributes struct {
	// The name of the line item option. When blank, it gets populated with the name of the associated sku option.
	Name *string `json:"name,omitempty"`
	// The line item option's quantity
	Quantity *int32 `json:"quantity,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order's market.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The unit amount of the line item option, in cents. When you add a line item option to an order, this is automatically populated from associated sku option's price.
	UnitAmountCents *int32 `json:"unit_amount_cents,omitempty"`
	// The unit amount of the line item option, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce.
	UnitAmountFloat *float32 `json:"unit_amount_float,omitempty"`
	// The unit amount of the line item option, formatted. This can be useful to display the amount with currency in you views.
	FormattedUnitAmount *string `json:"formatted_unit_amount,omitempty"`
	// The unit amount x quantity, in cents.
	TotalAmountCents *int32 `json:"total_amount_cents,omitempty"`
	// The unit amount x quantity, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce.
	TotalAmountFloat *float32 `json:"total_amount_float,omitempty"`
	// The unit amount x quantity, formatted. This can be useful to display the amount with currency in you views.
	FormattedTotalAmount *string `json:"formatted_total_amount,omitempty"`
	// The shipping delay that the customer can expect when adding this option (hours). Inherited from the associated sku option.
	DelayHours *int32 `json:"delay_hours,omitempty"`
	// The shipping delay that the customer can expect when adding this option (days, rounded).
	DelayDays *int32 `json:"delay_days,omitempty"`
	// Set of key-value pairs that represent the selected options.
	Options map[string]interface{} `json:"options,omitempty"`
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

// NewLineItemOptionDataAttributes instantiates a new LineItemOptionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionDataAttributes() *LineItemOptionDataAttributes {
	this := LineItemOptionDataAttributes{}
	return &this
}

// NewLineItemOptionDataAttributesWithDefaults instantiates a new LineItemOptionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionDataAttributesWithDefaults() *LineItemOptionDataAttributes {
	this := LineItemOptionDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LineItemOptionDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *LineItemOptionDataAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *LineItemOptionDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetUnitAmountCents returns the UnitAmountCents field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetUnitAmountCents() int32 {
	if o == nil || o.UnitAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.UnitAmountCents
}

// GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetUnitAmountCentsOk() (*int32, bool) {
	if o == nil || o.UnitAmountCents == nil {
		return nil, false
	}
	return o.UnitAmountCents, true
}

// HasUnitAmountCents returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasUnitAmountCents() bool {
	if o != nil && o.UnitAmountCents != nil {
		return true
	}

	return false
}

// SetUnitAmountCents gets a reference to the given int32 and assigns it to the UnitAmountCents field.
func (o *LineItemOptionDataAttributes) SetUnitAmountCents(v int32) {
	o.UnitAmountCents = &v
}

// GetUnitAmountFloat returns the UnitAmountFloat field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetUnitAmountFloat() float32 {
	if o == nil || o.UnitAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.UnitAmountFloat
}

// GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetUnitAmountFloatOk() (*float32, bool) {
	if o == nil || o.UnitAmountFloat == nil {
		return nil, false
	}
	return o.UnitAmountFloat, true
}

// HasUnitAmountFloat returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasUnitAmountFloat() bool {
	if o != nil && o.UnitAmountFloat != nil {
		return true
	}

	return false
}

// SetUnitAmountFloat gets a reference to the given float32 and assigns it to the UnitAmountFloat field.
func (o *LineItemOptionDataAttributes) SetUnitAmountFloat(v float32) {
	o.UnitAmountFloat = &v
}

// GetFormattedUnitAmount returns the FormattedUnitAmount field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetFormattedUnitAmount() string {
	if o == nil || o.FormattedUnitAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedUnitAmount
}

// GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetFormattedUnitAmountOk() (*string, bool) {
	if o == nil || o.FormattedUnitAmount == nil {
		return nil, false
	}
	return o.FormattedUnitAmount, true
}

// HasFormattedUnitAmount returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasFormattedUnitAmount() bool {
	if o != nil && o.FormattedUnitAmount != nil {
		return true
	}

	return false
}

// SetFormattedUnitAmount gets a reference to the given string and assigns it to the FormattedUnitAmount field.
func (o *LineItemOptionDataAttributes) SetFormattedUnitAmount(v string) {
	o.FormattedUnitAmount = &v
}

// GetTotalAmountCents returns the TotalAmountCents field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetTotalAmountCents() int32 {
	if o == nil || o.TotalAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.TotalAmountCents
}

// GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetTotalAmountCentsOk() (*int32, bool) {
	if o == nil || o.TotalAmountCents == nil {
		return nil, false
	}
	return o.TotalAmountCents, true
}

// HasTotalAmountCents returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasTotalAmountCents() bool {
	if o != nil && o.TotalAmountCents != nil {
		return true
	}

	return false
}

// SetTotalAmountCents gets a reference to the given int32 and assigns it to the TotalAmountCents field.
func (o *LineItemOptionDataAttributes) SetTotalAmountCents(v int32) {
	o.TotalAmountCents = &v
}

// GetTotalAmountFloat returns the TotalAmountFloat field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetTotalAmountFloat() float32 {
	if o == nil || o.TotalAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmountFloat
}

// GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetTotalAmountFloatOk() (*float32, bool) {
	if o == nil || o.TotalAmountFloat == nil {
		return nil, false
	}
	return o.TotalAmountFloat, true
}

// HasTotalAmountFloat returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasTotalAmountFloat() bool {
	if o != nil && o.TotalAmountFloat != nil {
		return true
	}

	return false
}

// SetTotalAmountFloat gets a reference to the given float32 and assigns it to the TotalAmountFloat field.
func (o *LineItemOptionDataAttributes) SetTotalAmountFloat(v float32) {
	o.TotalAmountFloat = &v
}

// GetFormattedTotalAmount returns the FormattedTotalAmount field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetFormattedTotalAmount() string {
	if o == nil || o.FormattedTotalAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedTotalAmount
}

// GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetFormattedTotalAmountOk() (*string, bool) {
	if o == nil || o.FormattedTotalAmount == nil {
		return nil, false
	}
	return o.FormattedTotalAmount, true
}

// HasFormattedTotalAmount returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasFormattedTotalAmount() bool {
	if o != nil && o.FormattedTotalAmount != nil {
		return true
	}

	return false
}

// SetFormattedTotalAmount gets a reference to the given string and assigns it to the FormattedTotalAmount field.
func (o *LineItemOptionDataAttributes) SetFormattedTotalAmount(v string) {
	o.FormattedTotalAmount = &v
}

// GetDelayHours returns the DelayHours field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetDelayHours() int32 {
	if o == nil || o.DelayHours == nil {
		var ret int32
		return ret
	}
	return *o.DelayHours
}

// GetDelayHoursOk returns a tuple with the DelayHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetDelayHoursOk() (*int32, bool) {
	if o == nil || o.DelayHours == nil {
		return nil, false
	}
	return o.DelayHours, true
}

// HasDelayHours returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasDelayHours() bool {
	if o != nil && o.DelayHours != nil {
		return true
	}

	return false
}

// SetDelayHours gets a reference to the given int32 and assigns it to the DelayHours field.
func (o *LineItemOptionDataAttributes) SetDelayHours(v int32) {
	o.DelayHours = &v
}

// GetDelayDays returns the DelayDays field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetDelayDays() int32 {
	if o == nil || o.DelayDays == nil {
		var ret int32
		return ret
	}
	return *o.DelayDays
}

// GetDelayDaysOk returns a tuple with the DelayDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetDelayDaysOk() (*int32, bool) {
	if o == nil || o.DelayDays == nil {
		return nil, false
	}
	return o.DelayDays, true
}

// HasDelayDays returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasDelayDays() bool {
	if o != nil && o.DelayDays != nil {
		return true
	}

	return false
}

// SetDelayDays gets a reference to the given int32 and assigns it to the DelayDays field.
func (o *LineItemOptionDataAttributes) SetDelayDays(v int32) {
	o.DelayDays = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *LineItemOptionDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LineItemOptionDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LineItemOptionDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LineItemOptionDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LineItemOptionDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *LineItemOptionDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LineItemOptionDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LineItemOptionDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LineItemOptionDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o LineItemOptionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.UnitAmountCents != nil {
		toSerialize["unit_amount_cents"] = o.UnitAmountCents
	}
	if o.UnitAmountFloat != nil {
		toSerialize["unit_amount_float"] = o.UnitAmountFloat
	}
	if o.FormattedUnitAmount != nil {
		toSerialize["formatted_unit_amount"] = o.FormattedUnitAmount
	}
	if o.TotalAmountCents != nil {
		toSerialize["total_amount_cents"] = o.TotalAmountCents
	}
	if o.TotalAmountFloat != nil {
		toSerialize["total_amount_float"] = o.TotalAmountFloat
	}
	if o.FormattedTotalAmount != nil {
		toSerialize["formatted_total_amount"] = o.FormattedTotalAmount
	}
	if o.DelayHours != nil {
		toSerialize["delay_hours"] = o.DelayHours
	}
	if o.DelayDays != nil {
		toSerialize["delay_days"] = o.DelayDays
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
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

type NullableLineItemOptionDataAttributes struct {
	value *LineItemOptionDataAttributes
	isSet bool
}

func (v NullableLineItemOptionDataAttributes) Get() *LineItemOptionDataAttributes {
	return v.value
}

func (v *NullableLineItemOptionDataAttributes) Set(val *LineItemOptionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionDataAttributes(val *LineItemOptionDataAttributes) *NullableLineItemOptionDataAttributes {
	return &NullableLineItemOptionDataAttributes{value: val, isSet: true}
}

func (v NullableLineItemOptionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
