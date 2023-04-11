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

// GETSubscriptionModels200ResponseDataInnerRelationships struct for GETSubscriptionModels200ResponseDataInnerRelationships
type GETSubscriptionModels200ResponseDataInnerRelationships struct {
	Markets            *GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets      `json:"markets,omitempty"`
	OrderSubscriptions *GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions `json:"order_subscriptions,omitempty"`
	Attachments        *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments  `json:"attachments,omitempty"`
}

// NewGETSubscriptionModels200ResponseDataInnerRelationships instantiates a new GETSubscriptionModels200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSubscriptionModels200ResponseDataInnerRelationships() *GETSubscriptionModels200ResponseDataInnerRelationships {
	this := GETSubscriptionModels200ResponseDataInnerRelationships{}
	return &this
}

// NewGETSubscriptionModels200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETSubscriptionModels200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSubscriptionModels200ResponseDataInnerRelationshipsWithDefaults() *GETSubscriptionModels200ResponseDataInnerRelationships {
	this := GETSubscriptionModels200ResponseDataInnerRelationships{}
	return &this
}

// GetMarkets returns the Markets field value if set, zero value otherwise.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) GetMarkets() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets {
	if o == nil || o.Markets == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets
		return ret
	}
	return *o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) GetMarketsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool) {
	if o == nil || o.Markets == nil {
		return nil, false
	}
	return o.Markets, true
}

// HasMarkets returns a boolean if a field has been set.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) HasMarkets() bool {
	if o != nil && o.Markets != nil {
		return true
	}

	return false
}

// SetMarkets gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets and assigns it to the Markets field.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) SetMarkets(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets) {
	o.Markets = &v
}

// GetOrderSubscriptions returns the OrderSubscriptions field value if set, zero value otherwise.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) GetOrderSubscriptions() GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions {
	if o == nil || o.OrderSubscriptions == nil {
		var ret GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscriptions
}

// GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) GetOrderSubscriptionsOk() (*GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions, bool) {
	if o == nil || o.OrderSubscriptions == nil {
		return nil, false
	}
	return o.OrderSubscriptions, true
}

// HasOrderSubscriptions returns a boolean if a field has been set.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) HasOrderSubscriptions() bool {
	if o != nil && o.OrderSubscriptions != nil {
		return true
	}

	return false
}

// SetOrderSubscriptions gets a reference to the given GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions and assigns it to the OrderSubscriptions field.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) SetOrderSubscriptions(v GETCustomers200ResponseDataInnerRelationshipsOrderSubscriptions) {
	o.OrderSubscriptions = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETSubscriptionModels200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETSubscriptionModels200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Markets != nil {
		toSerialize["markets"] = o.Markets
	}
	if o.OrderSubscriptions != nil {
		toSerialize["order_subscriptions"] = o.OrderSubscriptions
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETSubscriptionModels200ResponseDataInnerRelationships struct {
	value *GETSubscriptionModels200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETSubscriptionModels200ResponseDataInnerRelationships) Get() *GETSubscriptionModels200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETSubscriptionModels200ResponseDataInnerRelationships) Set(val *GETSubscriptionModels200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSubscriptionModels200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSubscriptionModels200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSubscriptionModels200ResponseDataInnerRelationships(val *GETSubscriptionModels200ResponseDataInnerRelationships) *NullableGETSubscriptionModels200ResponseDataInnerRelationships {
	return &NullableGETSubscriptionModels200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETSubscriptionModels200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSubscriptionModels200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
