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

// AttachmentData struct for AttachmentData
type AttachmentData struct {
	// The resource's type
	Type          string                                           `json:"type"`
	Attributes    GETAttachments200ResponseDataInnerAttributes     `json:"attributes"`
	Relationships *GETAttachments200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewAttachmentData instantiates a new AttachmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentData(type_ string, attributes GETAttachments200ResponseDataInnerAttributes) *AttachmentData {
	this := AttachmentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAttachmentDataWithDefaults instantiates a new AttachmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentDataWithDefaults() *AttachmentData {
	this := AttachmentData{}
	var type_ string = "attachments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *AttachmentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AttachmentData) GetAttributes() GETAttachments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETAttachments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetAttributesOk() (*GETAttachments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AttachmentData) SetAttributes(v GETAttachments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AttachmentData) GetRelationships() GETAttachments200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETAttachments200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetRelationshipsOk() (*GETAttachments200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AttachmentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETAttachments200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *AttachmentData) SetRelationships(v GETAttachments200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o AttachmentData) MarshalJSON() ([]byte, error) {
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

type NullableAttachmentData struct {
	value *AttachmentData
	isSet bool
}

func (v NullableAttachmentData) Get() *AttachmentData {
	return v.value
}

func (v *NullableAttachmentData) Set(val *AttachmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentData(val *AttachmentData) *NullableAttachmentData {
	return &NullableAttachmentData{value: val, isSet: true}
}

func (v NullableAttachmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
