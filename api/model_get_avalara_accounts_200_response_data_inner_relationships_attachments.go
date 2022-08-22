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

// GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments struct for GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
type GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments instantiates a new GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments(type_ string, id string) *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	this := GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsWithDefaults instantiates a new GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsWithDefaults() *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	this := GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments{}
	var type_ string = "attachments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) SetId(v string) {
	o.Id = v
}

func (o GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments struct {
	value *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
	isSet bool
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) Get() *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	return v.value
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) Set(val *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments(val *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	return &NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments{value: val, isSet: true}
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}