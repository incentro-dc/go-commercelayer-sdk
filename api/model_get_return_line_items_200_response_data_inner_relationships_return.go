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

// GETReturnLineItems200ResponseDataInnerRelationshipsReturn struct for GETReturnLineItems200ResponseDataInnerRelationshipsReturn
type GETReturnLineItems200ResponseDataInnerRelationshipsReturn struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *GETReturnLineItems200ResponseDataInnerRelationshipsReturnData `json:"data,omitempty"`
}

// NewGETReturnLineItems200ResponseDataInnerRelationshipsReturn instantiates a new GETReturnLineItems200ResponseDataInnerRelationshipsReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETReturnLineItems200ResponseDataInnerRelationshipsReturn() *GETReturnLineItems200ResponseDataInnerRelationshipsReturn {
	this := GETReturnLineItems200ResponseDataInnerRelationshipsReturn{}
	return &this
}

// NewGETReturnLineItems200ResponseDataInnerRelationshipsReturnWithDefaults instantiates a new GETReturnLineItems200ResponseDataInnerRelationshipsReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETReturnLineItems200ResponseDataInnerRelationshipsReturnWithDefaults() *GETReturnLineItems200ResponseDataInnerRelationshipsReturn {
	this := GETReturnLineItems200ResponseDataInnerRelationshipsReturn{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) GetData() GETReturnLineItems200ResponseDataInnerRelationshipsReturnData {
	if o == nil || o.Data == nil {
		var ret GETReturnLineItems200ResponseDataInnerRelationshipsReturnData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) GetDataOk() (*GETReturnLineItems200ResponseDataInnerRelationshipsReturnData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETReturnLineItems200ResponseDataInnerRelationshipsReturnData and assigns it to the Data field.
func (o *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) SetData(v GETReturnLineItems200ResponseDataInnerRelationshipsReturnData) {
	o.Data = &v
}

func (o GETReturnLineItems200ResponseDataInnerRelationshipsReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn struct {
	value *GETReturnLineItems200ResponseDataInnerRelationshipsReturn
	isSet bool
}

func (v NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn) Get() *GETReturnLineItems200ResponseDataInnerRelationshipsReturn {
	return v.value
}

func (v *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn) Set(val *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn(val *GETReturnLineItems200ResponseDataInnerRelationshipsReturn) *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn {
	return &NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn{value: val, isSet: true}
}

func (v NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETReturnLineItems200ResponseDataInnerRelationshipsReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
