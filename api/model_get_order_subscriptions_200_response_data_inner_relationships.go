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

// GETOrderSubscriptions200ResponseDataInnerRelationships struct for GETOrderSubscriptions200ResponseDataInnerRelationships
type GETOrderSubscriptions200ResponseDataInnerRelationships struct {
	Market      *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket `json:"market,omitempty"`
	SourceOrder *GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder           `json:"source_order,omitempty"`
	Customer    *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer         `json:"customer,omitempty"`
	OrderCopies *GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies    `json:"order_copies,omitempty"`
	Orders      *GETCustomers200ResponseDataInnerRelationshipsOrders                  `json:"orders,omitempty"`
	Events      *GETCleanups200ResponseDataInnerRelationshipsEvents                   `json:"events,omitempty"`
}

// NewGETOrderSubscriptions200ResponseDataInnerRelationships instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderSubscriptions200ResponseDataInnerRelationships() *GETOrderSubscriptions200ResponseDataInnerRelationships {
	this := GETOrderSubscriptions200ResponseDataInnerRelationships{}
	return &this
}

// NewGETOrderSubscriptions200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderSubscriptions200ResponseDataInnerRelationshipsWithDefaults() *GETOrderSubscriptions200ResponseDataInnerRelationships {
	this := GETOrderSubscriptions200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetSourceOrder returns the SourceOrder field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetSourceOrder() GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder {
	if o == nil || o.SourceOrder == nil {
		var ret GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder
		return ret
	}
	return *o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetSourceOrderOk() (*GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder, bool) {
	if o == nil || o.SourceOrder == nil {
		return nil, false
	}
	return o.SourceOrder, true
}

// HasSourceOrder returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasSourceOrder() bool {
	if o != nil && o.SourceOrder != nil {
		return true
	}

	return false
}

// SetSourceOrder gets a reference to the given GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder and assigns it to the SourceOrder field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetSourceOrder(v GETOrderCopies200ResponseDataInnerRelationshipsSourceOrder) {
	o.SourceOrder = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetOrderCopies returns the OrderCopies field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderCopies() GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies {
	if o == nil || o.OrderCopies == nil {
		var ret GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies
		return ret
	}
	return *o.OrderCopies
}

// GetOrderCopiesOk returns a tuple with the OrderCopies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrderCopiesOk() (*GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies, bool) {
	if o == nil || o.OrderCopies == nil {
		return nil, false
	}
	return o.OrderCopies, true
}

// HasOrderCopies returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasOrderCopies() bool {
	if o != nil && o.OrderCopies != nil {
		return true
	}

	return false
}

// SetOrderCopies gets a reference to the given GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies and assigns it to the OrderCopies field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetOrderCopies(v GETOrderSubscriptions200ResponseDataInnerRelationshipsOrderCopies) {
	o.OrderCopies = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrders() GETCustomers200ResponseDataInnerRelationshipsOrders {
	if o == nil || o.Orders == nil {
		var ret GETCustomers200ResponseDataInnerRelationshipsOrders
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetOrdersOk() (*GETCustomers200ResponseDataInnerRelationshipsOrders, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given GETCustomers200ResponseDataInnerRelationshipsOrders and assigns it to the Orders field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetOrders(v GETCustomers200ResponseDataInnerRelationshipsOrders) {
	o.Orders = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetEvents() GETCleanups200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCleanups200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) GetEventsOk() (*GETCleanups200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCleanups200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationships) SetEvents(v GETCleanups200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETOrderSubscriptions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.SourceOrder != nil {
		toSerialize["source_order"] = o.SourceOrder
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.OrderCopies != nil {
		toSerialize["order_copies"] = o.OrderCopies
	}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderSubscriptions200ResponseDataInnerRelationships struct {
	value *GETOrderSubscriptions200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationships) Get() *GETOrderSubscriptions200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationships) Set(val *GETOrderSubscriptions200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderSubscriptions200ResponseDataInnerRelationships(val *GETOrderSubscriptions200ResponseDataInnerRelationships) *NullableGETOrderSubscriptions200ResponseDataInnerRelationships {
	return &NullableGETOrderSubscriptions200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
