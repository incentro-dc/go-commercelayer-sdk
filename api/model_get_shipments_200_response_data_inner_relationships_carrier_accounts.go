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

// GETShipments200ResponseDataInnerRelationshipsCarrierAccounts struct for GETShipments200ResponseDataInnerRelationshipsCarrierAccounts
type GETShipments200ResponseDataInnerRelationshipsCarrierAccounts struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETShipments200ResponseDataInnerRelationshipsCarrierAccounts instantiates a new GETShipments200ResponseDataInnerRelationshipsCarrierAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerRelationshipsCarrierAccounts(type_ string, id string) *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts {
	this := GETShipments200ResponseDataInnerRelationshipsCarrierAccounts{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETShipments200ResponseDataInnerRelationshipsCarrierAccountsWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationshipsCarrierAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerRelationshipsCarrierAccountsWithDefaults() *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts {
	this := GETShipments200ResponseDataInnerRelationshipsCarrierAccounts{}
	var type_ string = "carrier_accounts"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) SetId(v string) {
	o.Id = v
}

func (o GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts struct {
	value *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts) Get() *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts) Set(val *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts(val *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) *NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts {
	return &NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerRelationshipsCarrierAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
