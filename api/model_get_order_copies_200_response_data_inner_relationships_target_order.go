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

// GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder struct for GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder
type GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrderData `json:"data,omitempty"`
}

// NewGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder instantiates a new GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder() *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder {
	this := GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder{}
	return &this
}

// NewGETOrderCopies200ResponseDataInnerRelationshipsTargetOrderWithDefaults instantiates a new GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderCopies200ResponseDataInnerRelationshipsTargetOrderWithDefaults() *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder {
	this := GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) GetData() GETOrderCopies200ResponseDataInnerRelationshipsTargetOrderData {
	if o == nil || o.Data == nil {
		var ret GETOrderCopies200ResponseDataInnerRelationshipsTargetOrderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) GetDataOk() (*GETOrderCopies200ResponseDataInnerRelationshipsTargetOrderData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETOrderCopies200ResponseDataInnerRelationshipsTargetOrderData and assigns it to the Data field.
func (o *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) SetData(v GETOrderCopies200ResponseDataInnerRelationshipsTargetOrderData) {
	o.Data = &v
}

func (o GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder struct {
	value *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder
	isSet bool
}

func (v NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) Get() *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder {
	return v.value
}

func (v *NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) Set(val *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder(val *GETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) *NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder {
	return &NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder{value: val, isSet: true}
}

func (v NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderCopies200ResponseDataInnerRelationshipsTargetOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
