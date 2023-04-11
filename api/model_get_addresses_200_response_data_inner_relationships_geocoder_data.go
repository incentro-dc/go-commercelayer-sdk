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

// GETAddresses200ResponseDataInnerRelationshipsGeocoderData struct for GETAddresses200ResponseDataInnerRelationshipsGeocoderData
type GETAddresses200ResponseDataInnerRelationshipsGeocoderData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETAddresses200ResponseDataInnerRelationshipsGeocoderData instantiates a new GETAddresses200ResponseDataInnerRelationshipsGeocoderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAddresses200ResponseDataInnerRelationshipsGeocoderData() *GETAddresses200ResponseDataInnerRelationshipsGeocoderData {
	this := GETAddresses200ResponseDataInnerRelationshipsGeocoderData{}
	return &this
}

// NewGETAddresses200ResponseDataInnerRelationshipsGeocoderDataWithDefaults instantiates a new GETAddresses200ResponseDataInnerRelationshipsGeocoderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAddresses200ResponseDataInnerRelationshipsGeocoderDataWithDefaults() *GETAddresses200ResponseDataInnerRelationshipsGeocoderData {
	this := GETAddresses200ResponseDataInnerRelationshipsGeocoderData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) SetId(v interface{}) {
	o.Id = v
}

func (o GETAddresses200ResponseDataInnerRelationshipsGeocoderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData struct {
	value *GETAddresses200ResponseDataInnerRelationshipsGeocoderData
	isSet bool
}

func (v NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData) Get() *GETAddresses200ResponseDataInnerRelationshipsGeocoderData {
	return v.value
}

func (v *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData) Set(val *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData(val *GETAddresses200ResponseDataInnerRelationshipsGeocoderData) *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData {
	return &NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData{value: val, isSet: true}
}

func (v NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAddresses200ResponseDataInnerRelationshipsGeocoderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
