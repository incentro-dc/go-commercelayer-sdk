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

// checks if the POSTReturnLineItems201ResponseDataRelationshipsReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTReturnLineItems201ResponseDataRelationshipsReturn{}

// POSTReturnLineItems201ResponseDataRelationshipsReturn struct for POSTReturnLineItems201ResponseDataRelationshipsReturn
type POSTReturnLineItems201ResponseDataRelationshipsReturn struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *POSTReturnLineItems201ResponseDataRelationshipsReturnData `json:"data,omitempty"`
}

// NewPOSTReturnLineItems201ResponseDataRelationshipsReturn instantiates a new POSTReturnLineItems201ResponseDataRelationshipsReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturnLineItems201ResponseDataRelationshipsReturn() *POSTReturnLineItems201ResponseDataRelationshipsReturn {
	this := POSTReturnLineItems201ResponseDataRelationshipsReturn{}
	return &this
}

// NewPOSTReturnLineItems201ResponseDataRelationshipsReturnWithDefaults instantiates a new POSTReturnLineItems201ResponseDataRelationshipsReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturnLineItems201ResponseDataRelationshipsReturnWithDefaults() *POSTReturnLineItems201ResponseDataRelationshipsReturn {
	this := POSTReturnLineItems201ResponseDataRelationshipsReturn{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) GetData() POSTReturnLineItems201ResponseDataRelationshipsReturnData {
	if o == nil || IsNil(o.Data) {
		var ret POSTReturnLineItems201ResponseDataRelationshipsReturnData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) GetDataOk() (*POSTReturnLineItems201ResponseDataRelationshipsReturnData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTReturnLineItems201ResponseDataRelationshipsReturnData and assigns it to the Data field.
func (o *POSTReturnLineItems201ResponseDataRelationshipsReturn) SetData(v POSTReturnLineItems201ResponseDataRelationshipsReturnData) {
	o.Data = &v
}

func (o POSTReturnLineItems201ResponseDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTReturnLineItems201ResponseDataRelationshipsReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn struct {
	value *POSTReturnLineItems201ResponseDataRelationshipsReturn
	isSet bool
}

func (v NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn) Get() *POSTReturnLineItems201ResponseDataRelationshipsReturn {
	return v.value
}

func (v *NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn) Set(val *POSTReturnLineItems201ResponseDataRelationshipsReturn) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturnLineItems201ResponseDataRelationshipsReturn(val *POSTReturnLineItems201ResponseDataRelationshipsReturn) *NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn {
	return &NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn{value: val, isSet: true}
}

func (v NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturnLineItems201ResponseDataRelationshipsReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
