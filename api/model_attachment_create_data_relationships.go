/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AttachmentCreateDataRelationships struct for AttachmentCreateDataRelationships
type AttachmentCreateDataRelationships struct {
	Attachable AttachmentDataRelationshipsAttachable `json:"attachable"`
}

// NewAttachmentCreateDataRelationships instantiates a new AttachmentCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentCreateDataRelationships(attachable AttachmentDataRelationshipsAttachable) *AttachmentCreateDataRelationships {
	this := AttachmentCreateDataRelationships{}
	this.Attachable = attachable
	return &this
}

// NewAttachmentCreateDataRelationshipsWithDefaults instantiates a new AttachmentCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentCreateDataRelationshipsWithDefaults() *AttachmentCreateDataRelationships {
	this := AttachmentCreateDataRelationships{}
	return &this
}

// GetAttachable returns the Attachable field value
func (o *AttachmentCreateDataRelationships) GetAttachable() AttachmentDataRelationshipsAttachable {
	if o == nil {
		var ret AttachmentDataRelationshipsAttachable
		return ret
	}

	return o.Attachable
}

// GetAttachableOk returns a tuple with the Attachable field value
// and a boolean to check if the value has been set.
func (o *AttachmentCreateDataRelationships) GetAttachableOk() (*AttachmentDataRelationshipsAttachable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachable, true
}

// SetAttachable sets field value
func (o *AttachmentCreateDataRelationships) SetAttachable(v AttachmentDataRelationshipsAttachable) {
	o.Attachable = v
}

func (o AttachmentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attachable"] = o.Attachable
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentCreateDataRelationships struct {
	value *AttachmentCreateDataRelationships
	isSet bool
}

func (v NullableAttachmentCreateDataRelationships) Get() *AttachmentCreateDataRelationships {
	return v.value
}

func (v *NullableAttachmentCreateDataRelationships) Set(val *AttachmentCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentCreateDataRelationships(val *AttachmentCreateDataRelationships) *NullableAttachmentCreateDataRelationships {
	return &NullableAttachmentCreateDataRelationships{value: val, isSet: true}
}

func (v NullableAttachmentCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
