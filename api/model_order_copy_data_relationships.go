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

// OrderCopyDataRelationships struct for OrderCopyDataRelationships
type OrderCopyDataRelationships struct {
	SourceOrder       *AdyenPaymentDataRelationshipsOrder          `json:"source_order,omitempty"`
	TargetOrder       *AdyenPaymentDataRelationshipsOrder          `json:"target_order,omitempty"`
	OrderSubscription *CustomerDataRelationshipsOrderSubscriptions `json:"order_subscription,omitempty"`
	Events            *CustomerAddressDataRelationshipsEvents      `json:"events,omitempty"`
}

// NewOrderCopyDataRelationships instantiates a new OrderCopyDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyDataRelationships() *OrderCopyDataRelationships {
	this := OrderCopyDataRelationships{}
	return &this
}

// NewOrderCopyDataRelationshipsWithDefaults instantiates a new OrderCopyDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyDataRelationshipsWithDefaults() *OrderCopyDataRelationships {
	this := OrderCopyDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetSourceOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.SourceOrder == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetSourceOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.SourceOrder == nil {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasSourceOrder() bool {
	if o != nil && o.SourceOrder != nil {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the SourceOrder field.
func (o *OrderCopyDataRelationships) SetSourceOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.SourceOrder = &v
}

// GetTargetOrder returns the TargetOrder field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetTargetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.TargetOrder == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.TargetOrder
}

// GetTargetOrderOk returns a tuple with the TargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetTargetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.TargetOrder == nil {
		return nil, false
	}
	return o.TargetOrder, true
}

// HasTargetOrder returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasTargetOrder() bool {
	if o != nil && o.TargetOrder != nil {
		return true
	}

	return false
}

// SetTargetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the TargetOrder field.
func (o *OrderCopyDataRelationships) SetTargetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.TargetOrder = &v
}

// GetOrderSubscription returns the OrderSubscription field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetOrderSubscription() CustomerDataRelationshipsOrderSubscriptions {
	if o == nil || o.OrderSubscription == nil {
		var ret CustomerDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscription
}

// GetOrderSubscriptionOk returns a tuple with the OrderSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetOrderSubscriptionOk() (*CustomerDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || o.OrderSubscription == nil {
		return nil, false
	}
	return o.OrderSubscription, true
}

// HasOrderSubscription returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasOrderSubscription() bool {
	if o != nil && o.OrderSubscription != nil {
		return true
	}

	return false
}

// SetOrderSubscription gets a reference to the given CustomerDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscription field.
func (o *OrderCopyDataRelationships) SetOrderSubscription(v CustomerDataRelationshipsOrderSubscriptions) {
	o.OrderSubscription = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrderCopyDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrderCopyDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *OrderCopyDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o OrderCopyDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceOrder != nil {
		toSerialize["source_order"] = o.SourceOrder
	}
	if o.TargetOrder != nil {
		toSerialize["target_order"] = o.TargetOrder
	}
	if o.OrderSubscription != nil {
		toSerialize["order_subscription"] = o.OrderSubscription
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCopyDataRelationships struct {
	value *OrderCopyDataRelationships
	isSet bool
}

func (v NullableOrderCopyDataRelationships) Get() *OrderCopyDataRelationships {
	return v.value
}

func (v *NullableOrderCopyDataRelationships) Set(val *OrderCopyDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyDataRelationships(val *OrderCopyDataRelationships) *NullableOrderCopyDataRelationships {
	return &NullableOrderCopyDataRelationships{value: val, isSet: true}
}

func (v NullableOrderCopyDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
