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

// CustomerPasswordResetDataRelationships struct for CustomerPasswordResetDataRelationships
type CustomerPasswordResetDataRelationships struct {
	Customer *CouponRecipientDataRelationshipsCustomer `json:"customer,omitempty"`
	Events   *CustomerAddressDataRelationshipsEvents   `json:"events,omitempty"`
}

// NewCustomerPasswordResetDataRelationships instantiates a new CustomerPasswordResetDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetDataRelationships() *CustomerPasswordResetDataRelationships {
	this := CustomerPasswordResetDataRelationships{}
	return &this
}

// NewCustomerPasswordResetDataRelationshipsWithDefaults instantiates a new CustomerPasswordResetDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetDataRelationshipsWithDefaults() *CustomerPasswordResetDataRelationships {
	this := CustomerPasswordResetDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CustomerPasswordResetDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CustomerPasswordResetDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CustomerPasswordResetDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *CustomerPasswordResetDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o CustomerPasswordResetDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPasswordResetDataRelationships struct {
	value *CustomerPasswordResetDataRelationships
	isSet bool
}

func (v NullableCustomerPasswordResetDataRelationships) Get() *CustomerPasswordResetDataRelationships {
	return v.value
}

func (v *NullableCustomerPasswordResetDataRelationships) Set(val *CustomerPasswordResetDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetDataRelationships(val *CustomerPasswordResetDataRelationships) *NullableCustomerPasswordResetDataRelationships {
	return &NullableCustomerPasswordResetDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
