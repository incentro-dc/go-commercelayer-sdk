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

// GETLineItems200ResponseDataInnerRelationshipsStockTransfers struct for GETLineItems200ResponseDataInnerRelationshipsStockTransfers
type GETLineItems200ResponseDataInnerRelationshipsStockTransfers struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockTransfers instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockTransfers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsStockTransfers(type_ string, id string) *GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	this := GETLineItems200ResponseDataInnerRelationshipsStockTransfers{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockTransfers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	this := GETLineItems200ResponseDataInnerRelationshipsStockTransfers{}
	var type_ string = "stock_transfers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) SetId(v string) {
	o.Id = v
}

func (o GETLineItems200ResponseDataInnerRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers struct {
	value *GETLineItems200ResponseDataInnerRelationshipsStockTransfers
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) Get() *GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) Set(val *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers(val *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
