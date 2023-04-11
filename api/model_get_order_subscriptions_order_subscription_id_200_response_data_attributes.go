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

// checks if the GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{}

// GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes struct for GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes
type GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes struct {
	// Unique identifier for the subscription (numeric)
	Number interface{} `json:"number,omitempty"`
	// The subscription status. One of 'draft' (default), 'inactive', 'active', or 'cancelled'.
	Status interface{} `json:"status,omitempty"`
	// The frequency of the subscription. Use one of the supported within 'hourly', 'daily', 'weekly', 'monthly', 'two-month', 'three-month', 'four-month', 'six-month', 'yearly', or provide your custom crontab expression (min unit is hour). Must be supported by existing associated subscription_model.
	Frequency interface{} `json:"frequency,omitempty"`
	// Indicates if the subscription will be activated considering the placed source order as its first run.
	ActivateBySourceOrder interface{} `json:"activate_by_source_order,omitempty"`
	// The email address of the customer, if any, associated to the source order.
	CustomerEmail interface{} `json:"customer_email,omitempty"`
	// The activation date/time of this subscription.
	StartsAt interface{} `json:"starts_at,omitempty"`
	// The expiration date/time of this subscription (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// The date/time of the subscription next run. Can be updated but only in the future, to copy with frequency changes.
	NextRunAt interface{} `json:"next_run_at,omitempty"`
	// The number of times this subscription has run.
	Occurrencies interface{} `json:"occurrencies,omitempty"`
	// Indicates the number of subscription errors, if any.
	ErrorsCount interface{} `json:"errors_count,omitempty"`
	// Indicates if the subscription has succeeded on its last run.
	SucceededOnLastRun interface{} `json:"succeeded_on_last_run,omitempty"`
	// The subscription options used to create the order (check order_factories for more information). For subscriptions the `place_target_order` is enabled by default, specify custom options to overwrite it.
	Options interface{} `json:"options,omitempty"`
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

// NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes instantiates a new GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes() *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	this := GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{}
	return &this
}

// NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults instantiates a new GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults() *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	this := GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasNumber() bool {
	if o != nil && IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given interface{} and assigns it to the Number field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetNumber(v interface{}) {
	o.Number = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetFrequency() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return &o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasFrequency() bool {
	if o != nil && IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given interface{} and assigns it to the Frequency field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetFrequency(v interface{}) {
	o.Frequency = v
}

// GetActivateBySourceOrder returns the ActivateBySourceOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetActivateBySourceOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ActivateBySourceOrder
}

// GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetActivateBySourceOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ActivateBySourceOrder) {
		return nil, false
	}
	return &o.ActivateBySourceOrder, true
}

// HasActivateBySourceOrder returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasActivateBySourceOrder() bool {
	if o != nil && IsNil(o.ActivateBySourceOrder) {
		return true
	}

	return false
}

// SetActivateBySourceOrder gets a reference to the given interface{} and assigns it to the ActivateBySourceOrder field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetActivateBySourceOrder(v interface{}) {
	o.ActivateBySourceOrder = v
}

// GetCustomerEmail returns the CustomerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCustomerEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerEmail) {
		return nil, false
	}
	return &o.CustomerEmail, true
}

// HasCustomerEmail returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasCustomerEmail() bool {
	if o != nil && IsNil(o.CustomerEmail) {
		return true
	}

	return false
}

// SetCustomerEmail gets a reference to the given interface{} and assigns it to the CustomerEmail field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCustomerEmail(v interface{}) {
	o.CustomerEmail = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetNextRunAt returns the NextRunAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNextRunAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NextRunAt
}

// GetNextRunAtOk returns a tuple with the NextRunAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNextRunAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NextRunAt) {
		return nil, false
	}
	return &o.NextRunAt, true
}

// HasNextRunAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasNextRunAt() bool {
	if o != nil && IsNil(o.NextRunAt) {
		return true
	}

	return false
}

// SetNextRunAt gets a reference to the given interface{} and assigns it to the NextRunAt field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetNextRunAt(v interface{}) {
	o.NextRunAt = v
}

// GetOccurrencies returns the Occurrencies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetOccurrencies() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Occurrencies
}

// GetOccurrenciesOk returns a tuple with the Occurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetOccurrenciesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Occurrencies) {
		return nil, false
	}
	return &o.Occurrencies, true
}

// HasOccurrencies returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasOccurrencies() bool {
	if o != nil && IsNil(o.Occurrencies) {
		return true
	}

	return false
}

// SetOccurrencies gets a reference to the given interface{} and assigns it to the Occurrencies field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetOccurrencies(v interface{}) {
	o.Occurrencies = v
}

// GetErrorsCount returns the ErrorsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetErrorsCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorsCount
}

// GetErrorsCountOk returns a tuple with the ErrorsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetErrorsCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ErrorsCount) {
		return nil, false
	}
	return &o.ErrorsCount, true
}

// HasErrorsCount returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasErrorsCount() bool {
	if o != nil && IsNil(o.ErrorsCount) {
		return true
	}

	return false
}

// SetErrorsCount gets a reference to the given interface{} and assigns it to the ErrorsCount field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetErrorsCount(v interface{}) {
	o.ErrorsCount = v
}

// GetSucceededOnLastRun returns the SucceededOnLastRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetSucceededOnLastRun() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SucceededOnLastRun
}

// GetSucceededOnLastRunOk returns a tuple with the SucceededOnLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetSucceededOnLastRunOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SucceededOnLastRun) {
		return nil, false
	}
	return &o.SucceededOnLastRun, true
}

// HasSucceededOnLastRun returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasSucceededOnLastRun() bool {
	if o != nil && IsNil(o.SucceededOnLastRun) {
		return true
	}

	return false
}

// SetSucceededOnLastRun gets a reference to the given interface{} and assigns it to the SucceededOnLastRun field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetSucceededOnLastRun(v interface{}) {
	o.SucceededOnLastRun = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes struct {
	value *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) Get() *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) Set(val *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes(val *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) *NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	return &NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}