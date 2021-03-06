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

// CustomerDataRelationships struct for CustomerDataRelationships
type CustomerDataRelationships struct {
	CustomerGroup          *CustomerDataRelationshipsCustomerGroup          `json:"customer_group,omitempty"`
	CustomerAddresses      *CustomerDataRelationshipsCustomerAddresses      `json:"customer_addresses,omitempty"`
	CustomerPaymentSources *CustomerDataRelationshipsCustomerPaymentSources `json:"customer_payment_sources,omitempty"`
	CustomerSubscriptions  *CustomerDataRelationshipsCustomerSubscriptions  `json:"customer_subscriptions,omitempty"`
	Orders                 *AdyenPaymentDataRelationshipsOrder              `json:"orders,omitempty"`
	OrderSubscriptions     *CustomerDataRelationshipsOrderSubscriptions     `json:"order_subscriptions,omitempty"`
	Returns                *CustomerDataRelationshipsReturns                `json:"returns,omitempty"`
	Attachments            *AvalaraAccountDataRelationshipsAttachments      `json:"attachments,omitempty"`
	Events                 *CustomerAddressDataRelationshipsEvents          `json:"events,omitempty"`
}

// NewCustomerDataRelationships instantiates a new CustomerDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationships() *CustomerDataRelationships {
	this := CustomerDataRelationships{}
	return &this
}

// NewCustomerDataRelationshipsWithDefaults instantiates a new CustomerDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsWithDefaults() *CustomerDataRelationships {
	this := CustomerDataRelationships{}
	return &this
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetCustomerGroup() CustomerDataRelationshipsCustomerGroup {
	if o == nil || o.CustomerGroup == nil {
		var ret CustomerDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetCustomerGroupOk() (*CustomerDataRelationshipsCustomerGroup, bool) {
	if o == nil || o.CustomerGroup == nil {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasCustomerGroup() bool {
	if o != nil && o.CustomerGroup != nil {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given CustomerDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *CustomerDataRelationships) SetCustomerGroup(v CustomerDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

// GetCustomerAddresses returns the CustomerAddresses field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetCustomerAddresses() CustomerDataRelationshipsCustomerAddresses {
	if o == nil || o.CustomerAddresses == nil {
		var ret CustomerDataRelationshipsCustomerAddresses
		return ret
	}
	return *o.CustomerAddresses
}

// GetCustomerAddressesOk returns a tuple with the CustomerAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetCustomerAddressesOk() (*CustomerDataRelationshipsCustomerAddresses, bool) {
	if o == nil || o.CustomerAddresses == nil {
		return nil, false
	}
	return o.CustomerAddresses, true
}

// HasCustomerAddresses returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasCustomerAddresses() bool {
	if o != nil && o.CustomerAddresses != nil {
		return true
	}

	return false
}

// SetCustomerAddresses gets a reference to the given CustomerDataRelationshipsCustomerAddresses and assigns it to the CustomerAddresses field.
func (o *CustomerDataRelationships) SetCustomerAddresses(v CustomerDataRelationshipsCustomerAddresses) {
	o.CustomerAddresses = &v
}

// GetCustomerPaymentSources returns the CustomerPaymentSources field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetCustomerPaymentSources() CustomerDataRelationshipsCustomerPaymentSources {
	if o == nil || o.CustomerPaymentSources == nil {
		var ret CustomerDataRelationshipsCustomerPaymentSources
		return ret
	}
	return *o.CustomerPaymentSources
}

// GetCustomerPaymentSourcesOk returns a tuple with the CustomerPaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetCustomerPaymentSourcesOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool) {
	if o == nil || o.CustomerPaymentSources == nil {
		return nil, false
	}
	return o.CustomerPaymentSources, true
}

// HasCustomerPaymentSources returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasCustomerPaymentSources() bool {
	if o != nil && o.CustomerPaymentSources != nil {
		return true
	}

	return false
}

// SetCustomerPaymentSources gets a reference to the given CustomerDataRelationshipsCustomerPaymentSources and assigns it to the CustomerPaymentSources field.
func (o *CustomerDataRelationships) SetCustomerPaymentSources(v CustomerDataRelationshipsCustomerPaymentSources) {
	o.CustomerPaymentSources = &v
}

// GetCustomerSubscriptions returns the CustomerSubscriptions field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetCustomerSubscriptions() CustomerDataRelationshipsCustomerSubscriptions {
	if o == nil || o.CustomerSubscriptions == nil {
		var ret CustomerDataRelationshipsCustomerSubscriptions
		return ret
	}
	return *o.CustomerSubscriptions
}

// GetCustomerSubscriptionsOk returns a tuple with the CustomerSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetCustomerSubscriptionsOk() (*CustomerDataRelationshipsCustomerSubscriptions, bool) {
	if o == nil || o.CustomerSubscriptions == nil {
		return nil, false
	}
	return o.CustomerSubscriptions, true
}

// HasCustomerSubscriptions returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasCustomerSubscriptions() bool {
	if o != nil && o.CustomerSubscriptions != nil {
		return true
	}

	return false
}

// SetCustomerSubscriptions gets a reference to the given CustomerDataRelationshipsCustomerSubscriptions and assigns it to the CustomerSubscriptions field.
func (o *CustomerDataRelationships) SetCustomerSubscriptions(v CustomerDataRelationshipsCustomerSubscriptions) {
	o.CustomerSubscriptions = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetOrders() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Orders == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetOrdersOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Orders field.
func (o *CustomerDataRelationships) SetOrders(v AdyenPaymentDataRelationshipsOrder) {
	o.Orders = &v
}

// GetOrderSubscriptions returns the OrderSubscriptions field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetOrderSubscriptions() CustomerDataRelationshipsOrderSubscriptions {
	if o == nil || o.OrderSubscriptions == nil {
		var ret CustomerDataRelationshipsOrderSubscriptions
		return ret
	}
	return *o.OrderSubscriptions
}

// GetOrderSubscriptionsOk returns a tuple with the OrderSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetOrderSubscriptionsOk() (*CustomerDataRelationshipsOrderSubscriptions, bool) {
	if o == nil || o.OrderSubscriptions == nil {
		return nil, false
	}
	return o.OrderSubscriptions, true
}

// HasOrderSubscriptions returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasOrderSubscriptions() bool {
	if o != nil && o.OrderSubscriptions != nil {
		return true
	}

	return false
}

// SetOrderSubscriptions gets a reference to the given CustomerDataRelationshipsOrderSubscriptions and assigns it to the OrderSubscriptions field.
func (o *CustomerDataRelationships) SetOrderSubscriptions(v CustomerDataRelationshipsOrderSubscriptions) {
	o.OrderSubscriptions = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetReturns() CustomerDataRelationshipsReturns {
	if o == nil || o.Returns == nil {
		var ret CustomerDataRelationshipsReturns
		return ret
	}
	return *o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetReturnsOk() (*CustomerDataRelationshipsReturns, bool) {
	if o == nil || o.Returns == nil {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasReturns() bool {
	if o != nil && o.Returns != nil {
		return true
	}

	return false
}

// SetReturns gets a reference to the given CustomerDataRelationshipsReturns and assigns it to the Returns field.
func (o *CustomerDataRelationships) SetReturns(v CustomerDataRelationshipsReturns) {
	o.Returns = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *CustomerDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *CustomerDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *CustomerDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *CustomerDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o CustomerDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerGroup != nil {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	if o.CustomerAddresses != nil {
		toSerialize["customer_addresses"] = o.CustomerAddresses
	}
	if o.CustomerPaymentSources != nil {
		toSerialize["customer_payment_sources"] = o.CustomerPaymentSources
	}
	if o.CustomerSubscriptions != nil {
		toSerialize["customer_subscriptions"] = o.CustomerSubscriptions
	}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if o.OrderSubscriptions != nil {
		toSerialize["order_subscriptions"] = o.OrderSubscriptions
	}
	if o.Returns != nil {
		toSerialize["returns"] = o.Returns
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationships struct {
	value *CustomerDataRelationships
	isSet bool
}

func (v NullableCustomerDataRelationships) Get() *CustomerDataRelationships {
	return v.value
}

func (v *NullableCustomerDataRelationships) Set(val *CustomerDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationships(val *CustomerDataRelationships) *NullableCustomerDataRelationships {
	return &NullableCustomerDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
