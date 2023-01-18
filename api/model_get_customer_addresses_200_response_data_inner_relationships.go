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

// GETCustomerAddresses200ResponseDataInnerRelationships struct for GETCustomerAddresses200ResponseDataInnerRelationships
type GETCustomerAddresses200ResponseDataInnerRelationships struct {
	Customer *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer `json:"customer,omitempty"`
	Address  *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress `json:"address,omitempty"`
	Events   *GETCleanups200ResponseDataInnerRelationshipsEvents           `json:"events,omitempty"`
}

// NewGETCustomerAddresses200ResponseDataInnerRelationships instantiates a new GETCustomerAddresses200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerAddresses200ResponseDataInnerRelationships() *GETCustomerAddresses200ResponseDataInnerRelationships {
	this := GETCustomerAddresses200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCustomerAddresses200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCustomerAddresses200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerAddresses200ResponseDataInnerRelationshipsWithDefaults() *GETCustomerAddresses200ResponseDataInnerRelationships {
	this := GETCustomerAddresses200ResponseDataInnerRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) GetAddress() GETCustomerAddresses200ResponseDataInnerRelationshipsAddress {
	if o == nil || o.Address == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) GetAddressOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsAddress and assigns it to the Address field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) SetAddress(v GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) {
	o.Address = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) GetEvents() GETCleanups200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCleanups200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) GetEventsOk() (*GETCleanups200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCleanups200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationships) SetEvents(v GETCleanups200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETCustomerAddresses200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerAddresses200ResponseDataInnerRelationships struct {
	value *GETCustomerAddresses200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationships) Get() *GETCustomerAddresses200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationships) Set(val *GETCustomerAddresses200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerAddresses200ResponseDataInnerRelationships(val *GETCustomerAddresses200ResponseDataInnerRelationships) *NullableGETCustomerAddresses200ResponseDataInnerRelationships {
	return &NullableGETCustomerAddresses200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
