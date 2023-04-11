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

// checks if the OrderSubscriptionItemCreateDataRelationshipsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionItemCreateDataRelationshipsItem{}

// OrderSubscriptionItemCreateDataRelationshipsItem struct for OrderSubscriptionItemCreateDataRelationshipsItem
type OrderSubscriptionItemCreateDataRelationshipsItem struct {
	Data POSTOrderSubscriptionItemsRequestDataRelationshipsItemData `json:"data"`
}

// NewOrderSubscriptionItemCreateDataRelationshipsItem instantiates a new OrderSubscriptionItemCreateDataRelationshipsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionItemCreateDataRelationshipsItem(data POSTOrderSubscriptionItemsRequestDataRelationshipsItemData) *OrderSubscriptionItemCreateDataRelationshipsItem {
	this := OrderSubscriptionItemCreateDataRelationshipsItem{}
	this.Data = data
	return &this
}

// NewOrderSubscriptionItemCreateDataRelationshipsItemWithDefaults instantiates a new OrderSubscriptionItemCreateDataRelationshipsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionItemCreateDataRelationshipsItemWithDefaults() *OrderSubscriptionItemCreateDataRelationshipsItem {
	this := OrderSubscriptionItemCreateDataRelationshipsItem{}
	return &this
}

// GetData returns the Data field value
func (o *OrderSubscriptionItemCreateDataRelationshipsItem) GetData() POSTOrderSubscriptionItemsRequestDataRelationshipsItemData {
	if o == nil {
		var ret POSTOrderSubscriptionItemsRequestDataRelationshipsItemData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemCreateDataRelationshipsItem) GetDataOk() (*POSTOrderSubscriptionItemsRequestDataRelationshipsItemData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderSubscriptionItemCreateDataRelationshipsItem) SetData(v POSTOrderSubscriptionItemsRequestDataRelationshipsItemData) {
	o.Data = v
}

func (o OrderSubscriptionItemCreateDataRelationshipsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionItemCreateDataRelationshipsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrderSubscriptionItemCreateDataRelationshipsItem struct {
	value *OrderSubscriptionItemCreateDataRelationshipsItem
	isSet bool
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsItem) Get() *OrderSubscriptionItemCreateDataRelationshipsItem {
	return v.value
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsItem) Set(val *OrderSubscriptionItemCreateDataRelationshipsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionItemCreateDataRelationshipsItem(val *OrderSubscriptionItemCreateDataRelationshipsItem) *NullableOrderSubscriptionItemCreateDataRelationshipsItem {
	return &NullableOrderSubscriptionItemCreateDataRelationshipsItem{value: val, isSet: true}
}

func (v NullableOrderSubscriptionItemCreateDataRelationshipsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionItemCreateDataRelationshipsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}