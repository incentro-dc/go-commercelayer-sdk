/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OrderSubscriptionUpdateData struct for OrderSubscriptionUpdateData
type OrderSubscriptionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                              `json:"id"`
	Attributes    PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                                              `json:"relationships,omitempty"`
}

// NewOrderSubscriptionUpdateData instantiates a new OrderSubscriptionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionUpdateData(type_ string, id string, attributes PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) *OrderSubscriptionUpdateData {
	this := OrderSubscriptionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewOrderSubscriptionUpdateDataWithDefaults instantiates a new OrderSubscriptionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionUpdateDataWithDefaults() *OrderSubscriptionUpdateData {
	this := OrderSubscriptionUpdateData{}
	var type_ string = "order_subscriptions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *OrderSubscriptionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderSubscriptionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *OrderSubscriptionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderSubscriptionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *OrderSubscriptionUpdateData) GetAttributes() PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionUpdateData) GetAttributesOk() (*PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderSubscriptionUpdateData) SetAttributes(v PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderSubscriptionUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderSubscriptionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *OrderSubscriptionUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o OrderSubscriptionUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableOrderSubscriptionUpdateData struct {
	value *OrderSubscriptionUpdateData
	isSet bool
}

func (v NullableOrderSubscriptionUpdateData) Get() *OrderSubscriptionUpdateData {
	return v.value
}

func (v *NullableOrderSubscriptionUpdateData) Set(val *OrderSubscriptionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionUpdateData(val *OrderSubscriptionUpdateData) *NullableOrderSubscriptionUpdateData {
	return &NullableOrderSubscriptionUpdateData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
