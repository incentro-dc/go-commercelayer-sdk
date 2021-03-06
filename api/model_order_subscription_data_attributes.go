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

// OrderSubscriptionDataAttributes struct for OrderSubscriptionDataAttributes
type OrderSubscriptionDataAttributes struct {
	// Unique identifier for the subscription (numeric)
	Number *string `json:"number,omitempty"`
	// The subscription status. One of 'draft' (default), 'inactive', 'active', or 'cancelled'.
	Status *string `json:"status,omitempty"`
	// The frequency of the subscription. One of 'hourly', 'daily', 'weekly', 'monthly', 'two-month', 'three-month', 'four-month', 'six-month', or 'yearly'.
	Frequency *string `json:"frequency,omitempty"`
	// Indicates if the subscription will be actvated considering the placed source order as its first run, default to true.
	ActivateBySourceOrder *bool `json:"activate_by_source_order,omitempty"`
	// The email address of the customer, if any, associated to the source order.
	CustomerEmail *string `json:"customer_email,omitempty"`
	// The activation date/time of this subscription.
	StartsAt *string `json:"starts_at,omitempty"`
	// The expiration date/time of this subscription (must be after starts_at).
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The date/time of the subscription next run.
	NextRunAt *string `json:"next_run_at,omitempty"`
	// The number of times this subscription has run.
	Occurrencies *int32 `json:"occurrencies,omitempty"`
	// Indicates the number of subscription errors, if any.
	ErrorsCount *int32 `json:"errors_count,omitempty"`
	// Indicates if the subscription has succeeded on its last run.
	SucceededOnLastRun *bool `json:"succeeded_on_last_run,omitempty"`
	// The subscription options used to create the order copy (check order_copies for more information). For subscriptions the `place_target_order` is enabled by default, specify custom options to overwrite it.
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

// NewOrderSubscriptionDataAttributes instantiates a new OrderSubscriptionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionDataAttributes() *OrderSubscriptionDataAttributes {
	this := OrderSubscriptionDataAttributes{}
	return &this
}

// NewOrderSubscriptionDataAttributesWithDefaults instantiates a new OrderSubscriptionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataAttributesWithDefaults() *OrderSubscriptionDataAttributes {
	this := OrderSubscriptionDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *OrderSubscriptionDataAttributes) SetNumber(v string) {
	o.Number = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderSubscriptionDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetFrequency() string {
	if o == nil || o.Frequency == nil {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetFrequencyOk() (*string, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *OrderSubscriptionDataAttributes) SetFrequency(v string) {
	o.Frequency = &v
}

// GetActivateBySourceOrder returns the ActivateBySourceOrder field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetActivateBySourceOrder() bool {
	if o == nil || o.ActivateBySourceOrder == nil {
		var ret bool
		return ret
	}
	return *o.ActivateBySourceOrder
}

// GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetActivateBySourceOrderOk() (*bool, bool) {
	if o == nil || o.ActivateBySourceOrder == nil {
		return nil, false
	}
	return o.ActivateBySourceOrder, true
}

// HasActivateBySourceOrder returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasActivateBySourceOrder() bool {
	if o != nil && o.ActivateBySourceOrder != nil {
		return true
	}

	return false
}

// SetActivateBySourceOrder gets a reference to the given bool and assigns it to the ActivateBySourceOrder field.
func (o *OrderSubscriptionDataAttributes) SetActivateBySourceOrder(v bool) {
	o.ActivateBySourceOrder = &v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetCustomerEmail() string {
	if o == nil || o.CustomerEmail == nil {
		var ret string
		return ret
	}
	return *o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetCustomerEmailOk() (*string, bool) {
	if o == nil || o.CustomerEmail == nil {
		return nil, false
	}
	return o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasCustomerEmail() bool {
	if o != nil && o.CustomerEmail != nil {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given string and assigns it to the CustomerEmail field.
func (o *OrderSubscriptionDataAttributes) SetCustomerEmail(v string) {
	o.CustomerEmail = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetStartsAt() string {
	if o == nil || o.StartsAt == nil {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetStartsAtOk() (*string, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *OrderSubscriptionDataAttributes) SetStartsAt(v string) {
	o.StartsAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *OrderSubscriptionDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetNextRunAt returns the NextRunAt field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetNextRunAt() string {
	if o == nil || o.NextRunAt == nil {
		var ret string
		return ret
	}
	return *o.NextRunAt
}

// GetNextRunAtOk returns a tuple with the NextRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetNextRunAtOk() (*string, bool) {
	if o == nil || o.NextRunAt == nil {
		return nil, false
	}
	return o.NextRunAt, true
}

// HasNextRunAt returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasNextRunAt() bool {
	if o != nil && o.NextRunAt != nil {
		return true
	}

	return false
}

// SetNextRunAt gets a reference to the given string and assigns it to the NextRunAt field.
func (o *OrderSubscriptionDataAttributes) SetNextRunAt(v string) {
	o.NextRunAt = &v
}

// GetOccurrencies returns the Occurrencies field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetOccurrencies() int32 {
	if o == nil || o.Occurrencies == nil {
		var ret int32
		return ret
	}
	return *o.Occurrencies
}

// GetOccurrenciesOk returns a tuple with the Occurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetOccurrenciesOk() (*int32, bool) {
	if o == nil || o.Occurrencies == nil {
		return nil, false
	}
	return o.Occurrencies, true
}

// HasOccurrencies returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasOccurrencies() bool {
	if o != nil && o.Occurrencies != nil {
		return true
	}

	return false
}

// SetOccurrencies gets a reference to the given int32 and assigns it to the Occurrencies field.
func (o *OrderSubscriptionDataAttributes) SetOccurrencies(v int32) {
	o.Occurrencies = &v
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetErrorsCount() int32 {
	if o == nil || o.ErrorsCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorsCount
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetErrorsCountOk() (*int32, bool) {
	if o == nil || o.ErrorsCount == nil {
		return nil, false
	}
	return o.ErrorsCount, true
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasErrorsCount() bool {
	if o != nil && o.ErrorsCount != nil {
		return true
	}

	return false
}

// SetErrorsCount gets a reference to the given int32 and assigns it to the ErrorsCount field.
func (o *OrderSubscriptionDataAttributes) SetErrorsCount(v int32) {
	o.ErrorsCount = &v
}

// GetSucceededOnLastRun returns the SucceededOnLastRun field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetSucceededOnLastRun() bool {
	if o == nil || o.SucceededOnLastRun == nil {
		var ret bool
		return ret
	}
	return *o.SucceededOnLastRun
}

// GetSucceededOnLastRunOk returns a tuple with the SucceededOnLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetSucceededOnLastRunOk() (*bool, bool) {
	if o == nil || o.SucceededOnLastRun == nil {
		return nil, false
	}
	return o.SucceededOnLastRun, true
}

// HasSucceededOnLastRun returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasSucceededOnLastRun() bool {
	if o != nil && o.SucceededOnLastRun != nil {
		return true
	}

	return false
}

// SetSucceededOnLastRun gets a reference to the given bool and assigns it to the SucceededOnLastRun field.
func (o *OrderSubscriptionDataAttributes) SetSucceededOnLastRun(v bool) {
	o.SucceededOnLastRun = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *OrderSubscriptionDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderSubscriptionDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OrderSubscriptionDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OrderSubscriptionDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *OrderSubscriptionDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *OrderSubscriptionDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderSubscriptionDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderSubscriptionDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *OrderSubscriptionDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o OrderSubscriptionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.ActivateBySourceOrder != nil {
		toSerialize["activate_by_source_order"] = o.ActivateBySourceOrder
	}
	if o.CustomerEmail != nil {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.NextRunAt != nil {
		toSerialize["next_run_at"] = o.NextRunAt
	}
	if o.Occurrencies != nil {
		toSerialize["occurrencies"] = o.Occurrencies
	}
	if o.ErrorsCount != nil {
		toSerialize["errors_count"] = o.ErrorsCount
	}
	if o.SucceededOnLastRun != nil {
		toSerialize["succeeded_on_last_run"] = o.SucceededOnLastRun
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

type NullableOrderSubscriptionDataAttributes struct {
	value *OrderSubscriptionDataAttributes
	isSet bool
}

func (v NullableOrderSubscriptionDataAttributes) Get() *OrderSubscriptionDataAttributes {
	return v.value
}

func (v *NullableOrderSubscriptionDataAttributes) Set(val *OrderSubscriptionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionDataAttributes(val *OrderSubscriptionDataAttributes) *NullableOrderSubscriptionDataAttributes {
	return &NullableOrderSubscriptionDataAttributes{value: val, isSet: true}
}

func (v NullableOrderSubscriptionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
