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

// GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime struct for GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime
type GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime instantiates a new GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime(type_ string, id string) *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime {
	this := GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTimeWithDefaults() *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime {
	this := GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime{}
	var type_ string = "delivery_lead_times"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) SetId(v string) {
	o.Id = v
}

func (o GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime struct {
	value *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) Get() *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) Set(val *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime(val *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime {
	return &NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}