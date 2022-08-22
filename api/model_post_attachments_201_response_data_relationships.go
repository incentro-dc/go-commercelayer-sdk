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

// POSTAttachments201ResponseDataRelationships struct for POSTAttachments201ResponseDataRelationships
type POSTAttachments201ResponseDataRelationships struct {
	Attachable GETAttachments200ResponseDataInnerRelationshipsAttachable `json:"attachable"`
}

// NewPOSTAttachments201ResponseDataRelationships instantiates a new POSTAttachments201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAttachments201ResponseDataRelationships(attachable GETAttachments200ResponseDataInnerRelationshipsAttachable) *POSTAttachments201ResponseDataRelationships {
	this := POSTAttachments201ResponseDataRelationships{}
	this.Attachable = attachable
	return &this
}

// NewPOSTAttachments201ResponseDataRelationshipsWithDefaults instantiates a new POSTAttachments201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAttachments201ResponseDataRelationshipsWithDefaults() *POSTAttachments201ResponseDataRelationships {
	this := POSTAttachments201ResponseDataRelationships{}
	return &this
}

// GetAttachable returns the Attachable field value
func (o *POSTAttachments201ResponseDataRelationships) GetAttachable() GETAttachments200ResponseDataInnerRelationshipsAttachable {
	if o == nil {
		var ret GETAttachments200ResponseDataInnerRelationshipsAttachable
		return ret
	}

	return o.Attachable
}

// GetAttachableOk returns a tuple with the Attachable field value
// and a boolean to check if the value has been set.
func (o *POSTAttachments201ResponseDataRelationships) GetAttachableOk() (*GETAttachments200ResponseDataInnerRelationshipsAttachable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachable, true
}

// SetAttachable sets field value
func (o *POSTAttachments201ResponseDataRelationships) SetAttachable(v GETAttachments200ResponseDataInnerRelationshipsAttachable) {
	o.Attachable = v
}

func (o POSTAttachments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attachable"] = o.Attachable
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTAttachments201ResponseDataRelationships struct {
	value *POSTAttachments201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTAttachments201ResponseDataRelationships) Get() *POSTAttachments201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTAttachments201ResponseDataRelationships) Set(val *POSTAttachments201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAttachments201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAttachments201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAttachments201ResponseDataRelationships(val *POSTAttachments201ResponseDataRelationships) *NullablePOSTAttachments201ResponseDataRelationships {
	return &NullablePOSTAttachments201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTAttachments201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAttachments201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
