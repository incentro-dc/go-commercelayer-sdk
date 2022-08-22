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

// GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories struct for GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories
type GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories instantiates a new GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories(type_ string, id string) *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories {
	this := GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategoriesWithDefaults instantiates a new GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategoriesWithDefaults() *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories {
	this := GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories{}
	var type_ string = "tax_categories"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) SetId(v string) {
	o.Id = v
}

func (o GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories struct {
	value *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories
	isSet bool
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) Get() *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories {
	return v.value
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) Set(val *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories(val *GETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories {
	return &NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories{value: val, isSet: true}
}

func (v NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAvalaraAccounts200ResponseDataInnerRelationshipsTaxCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
