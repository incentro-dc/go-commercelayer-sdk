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

// checks if the POSTLineItems201ResponseDataRelationshipsStockTransfers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsStockTransfers{}

// POSTLineItems201ResponseDataRelationshipsStockTransfers struct for POSTLineItems201ResponseDataRelationshipsStockTransfers
type POSTLineItems201ResponseDataRelationshipsStockTransfers struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *POSTLineItems201ResponseDataRelationshipsStockTransfersData `json:"data,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsStockTransfers instantiates a new POSTLineItems201ResponseDataRelationshipsStockTransfers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsStockTransfers() *POSTLineItems201ResponseDataRelationshipsStockTransfers {
	this := POSTLineItems201ResponseDataRelationshipsStockTransfers{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsStockTransfersWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsStockTransfers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsStockTransfersWithDefaults() *POSTLineItems201ResponseDataRelationshipsStockTransfers {
	this := POSTLineItems201ResponseDataRelationshipsStockTransfers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) GetData() POSTLineItems201ResponseDataRelationshipsStockTransfersData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseDataRelationshipsStockTransfersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) GetDataOk() (*POSTLineItems201ResponseDataRelationshipsStockTransfersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseDataRelationshipsStockTransfersData and assigns it to the Data field.
func (o *POSTLineItems201ResponseDataRelationshipsStockTransfers) SetData(v POSTLineItems201ResponseDataRelationshipsStockTransfersData) {
	o.Data = &v
}

func (o POSTLineItems201ResponseDataRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsStockTransfers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers struct {
	value *POSTLineItems201ResponseDataRelationshipsStockTransfers
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers) Get() *POSTLineItems201ResponseDataRelationshipsStockTransfers {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers) Set(val *POSTLineItems201ResponseDataRelationshipsStockTransfers) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsStockTransfers(val *POSTLineItems201ResponseDataRelationshipsStockTransfers) *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers {
	return &NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsStockTransfers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
