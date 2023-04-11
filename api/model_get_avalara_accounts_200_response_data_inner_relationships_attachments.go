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

// GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments struct for GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
type GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks         `json:"links,omitempty"`
	Data  *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsData `json:"data,omitempty"`
}

// NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments instantiates a new GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments() *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	this := GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments{}
	return &this
}

// NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsWithDefaults instantiates a new GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsWithDefaults() *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	this := GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetData() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsData {
	if o == nil || o.Data == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) GetDataOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsData and assigns it to the Data field.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) SetData(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachmentsData) {
	o.Data = &v
}

func (o GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
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
