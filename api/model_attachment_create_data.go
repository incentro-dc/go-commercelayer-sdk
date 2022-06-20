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

// AttachmentCreateData struct for AttachmentCreateData
type AttachmentCreateData struct {
	// The resource's type
	Type          string                             `json:"type"`
	Attributes    AttachmentCreateDataAttributes     `json:"attributes"`
	Relationships *AttachmentCreateDataRelationships `json:"relationships,omitempty"`
}

// NewAttachmentCreateData instantiates a new AttachmentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentCreateData(type_ string, attributes AttachmentCreateDataAttributes) *AttachmentCreateData {
	this := AttachmentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAttachmentCreateDataWithDefaults instantiates a new AttachmentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentCreateDataWithDefaults() *AttachmentCreateData {
	this := AttachmentCreateData{}
	var type_ string = "attachments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AttachmentCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AttachmentCreateData) GetAttributes() AttachmentCreateDataAttributes {
	if o == nil {
		var ret AttachmentCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentCreateData) GetAttributesOk() (*AttachmentCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AttachmentCreateData) SetAttributes(v AttachmentCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AttachmentCreateData) GetRelationships() AttachmentCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AttachmentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentCreateData) GetRelationshipsOk() (*AttachmentCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AttachmentCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AttachmentCreateDataRelationships and assigns it to the Relationships field.
func (o *AttachmentCreateData) SetRelationships(v AttachmentCreateDataRelationships) {
	o.Relationships = &v
}

func (o AttachmentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableAttachmentCreateData struct {
	value *AttachmentCreateData
	isSet bool
}

func (v NullableAttachmentCreateData) Get() *AttachmentCreateData {
	return v.value
}

func (v *NullableAttachmentCreateData) Set(val *AttachmentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentCreateData(val *AttachmentCreateData) *NullableAttachmentCreateData {
	return &NullableAttachmentCreateData{value: val, isSet: true}
}

func (v NullableAttachmentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
