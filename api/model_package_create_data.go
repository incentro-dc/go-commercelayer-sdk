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

// PackageCreateData struct for PackageCreateData
type PackageCreateData struct {
	// The resource's type
	Type          interface{}                           `json:"type"`
	Attributes    POSTPackages201ResponseDataAttributes `json:"attributes"`
	Relationships *PackageCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewPackageCreateData instantiates a new PackageCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageCreateData(type_ interface{}, attributes POSTPackages201ResponseDataAttributes) *PackageCreateData {
	this := PackageCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPackageCreateDataWithDefaults instantiates a new PackageCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageCreateDataWithDefaults() *PackageCreateData {
	this := PackageCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PackageCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PackageCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PackageCreateData) GetAttributes() POSTPackages201ResponseDataAttributes {
	if o == nil {
		var ret POSTPackages201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PackageCreateData) GetAttributesOk() (*POSTPackages201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PackageCreateData) SetAttributes(v POSTPackages201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PackageCreateData) GetRelationships() PackageCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PackageCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageCreateData) GetRelationshipsOk() (*PackageCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PackageCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PackageCreateDataRelationships and assigns it to the Relationships field.
func (o *PackageCreateData) SetRelationships(v PackageCreateDataRelationships) {
	o.Relationships = &v
}

func (o PackageCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
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

type NullablePackageCreateData struct {
	value *PackageCreateData
	isSet bool
}

func (v NullablePackageCreateData) Get() *PackageCreateData {
	return v.value
}

func (v *NullablePackageCreateData) Set(val *PackageCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageCreateData(val *PackageCreateData) *NullablePackageCreateData {
	return &NullablePackageCreateData{value: val, isSet: true}
}

func (v NullablePackageCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
