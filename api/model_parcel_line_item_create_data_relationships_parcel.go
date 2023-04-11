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

// ParcelLineItemCreateDataRelationshipsParcel struct for ParcelLineItemCreateDataRelationshipsParcel
type ParcelLineItemCreateDataRelationshipsParcel struct {
	Data PackageDataRelationshipsParcelsData `json:"data"`
}

// NewParcelLineItemCreateDataRelationshipsParcel instantiates a new ParcelLineItemCreateDataRelationshipsParcel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemCreateDataRelationshipsParcel(data PackageDataRelationshipsParcelsData) *ParcelLineItemCreateDataRelationshipsParcel {
	this := ParcelLineItemCreateDataRelationshipsParcel{}
	this.Data = data
	return &this
}

// NewParcelLineItemCreateDataRelationshipsParcelWithDefaults instantiates a new ParcelLineItemCreateDataRelationshipsParcel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemCreateDataRelationshipsParcelWithDefaults() *ParcelLineItemCreateDataRelationshipsParcel {
	this := ParcelLineItemCreateDataRelationshipsParcel{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelLineItemCreateDataRelationshipsParcel) GetData() PackageDataRelationshipsParcelsData {
	if o == nil {
		var ret PackageDataRelationshipsParcelsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemCreateDataRelationshipsParcel) GetDataOk() (*PackageDataRelationshipsParcelsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelLineItemCreateDataRelationshipsParcel) SetData(v PackageDataRelationshipsParcelsData) {
	o.Data = v
}

func (o ParcelLineItemCreateDataRelationshipsParcel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableParcelLineItemCreateDataRelationshipsParcel struct {
	value *ParcelLineItemCreateDataRelationshipsParcel
	isSet bool
}

func (v NullableParcelLineItemCreateDataRelationshipsParcel) Get() *ParcelLineItemCreateDataRelationshipsParcel {
	return v.value
}

func (v *NullableParcelLineItemCreateDataRelationshipsParcel) Set(val *ParcelLineItemCreateDataRelationshipsParcel) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemCreateDataRelationshipsParcel) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemCreateDataRelationshipsParcel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemCreateDataRelationshipsParcel(val *ParcelLineItemCreateDataRelationshipsParcel) *NullableParcelLineItemCreateDataRelationshipsParcel {
	return &NullableParcelLineItemCreateDataRelationshipsParcel{value: val, isSet: true}
}

func (v NullableParcelLineItemCreateDataRelationshipsParcel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemCreateDataRelationshipsParcel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
