/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeliveryLeadTimeDataRelationshipsStockLocation struct for DeliveryLeadTimeDataRelationshipsStockLocation
type DeliveryLeadTimeDataRelationshipsStockLocation struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewDeliveryLeadTimeDataRelationshipsStockLocation instantiates a new DeliveryLeadTimeDataRelationshipsStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeDataRelationshipsStockLocation(type_ string, id string) *DeliveryLeadTimeDataRelationshipsStockLocation {
	this := DeliveryLeadTimeDataRelationshipsStockLocation{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewDeliveryLeadTimeDataRelationshipsStockLocationWithDefaults instantiates a new DeliveryLeadTimeDataRelationshipsStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataRelationshipsStockLocationWithDefaults() *DeliveryLeadTimeDataRelationshipsStockLocation {
	this := DeliveryLeadTimeDataRelationshipsStockLocation{}
	var type_ string = "stock_locations"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeliveryLeadTimeDataRelationshipsStockLocation) SetId(v string) {
	o.Id = v
}

func (o DeliveryLeadTimeDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeDataRelationshipsStockLocation struct {
	value *DeliveryLeadTimeDataRelationshipsStockLocation
	isSet bool
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocation) Get() *DeliveryLeadTimeDataRelationshipsStockLocation {
	return v.value
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocation) Set(val *DeliveryLeadTimeDataRelationshipsStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeDataRelationshipsStockLocation(val *DeliveryLeadTimeDataRelationshipsStockLocation) *NullableDeliveryLeadTimeDataRelationshipsStockLocation {
	return &NullableDeliveryLeadTimeDataRelationshipsStockLocation{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeDataRelationshipsStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeDataRelationshipsStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
