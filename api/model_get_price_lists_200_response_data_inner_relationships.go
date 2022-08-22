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

// GETPriceLists200ResponseDataInnerRelationships struct for GETPriceLists200ResponseDataInnerRelationships
type GETPriceLists200ResponseDataInnerRelationships struct {
	Prices      *GETPriceLists200ResponseDataInnerRelationshipsPrices           `json:"prices,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETPriceLists200ResponseDataInnerRelationships instantiates a new GETPriceLists200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceLists200ResponseDataInnerRelationships() *GETPriceLists200ResponseDataInnerRelationships {
	this := GETPriceLists200ResponseDataInnerRelationships{}
	return &this
}

// NewGETPriceLists200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETPriceLists200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceLists200ResponseDataInnerRelationshipsWithDefaults() *GETPriceLists200ResponseDataInnerRelationships {
	this := GETPriceLists200ResponseDataInnerRelationships{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInnerRelationships) GetPrices() GETPriceLists200ResponseDataInnerRelationshipsPrices {
	if o == nil || o.Prices == nil {
		var ret GETPriceLists200ResponseDataInnerRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInnerRelationships) GetPricesOk() (*GETPriceLists200ResponseDataInnerRelationshipsPrices, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInnerRelationships) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given GETPriceLists200ResponseDataInnerRelationshipsPrices and assigns it to the Prices field.
func (o *GETPriceLists200ResponseDataInnerRelationships) SetPrices(v GETPriceLists200ResponseDataInnerRelationshipsPrices) {
	o.Prices = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETPriceLists200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETPriceLists200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceLists200ResponseDataInnerRelationships struct {
	value *GETPriceLists200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETPriceLists200ResponseDataInnerRelationships) Get() *GETPriceLists200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationships) Set(val *GETPriceLists200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceLists200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceLists200ResponseDataInnerRelationships(val *GETPriceLists200ResponseDataInnerRelationships) *NullableGETPriceLists200ResponseDataInnerRelationships {
	return &NullableGETPriceLists200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETPriceLists200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
