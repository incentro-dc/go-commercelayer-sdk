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

// checks if the POSTExportsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExportsRequestData{}

// POSTExportsRequestData struct for POSTExportsRequestData
type POSTExportsRequestData struct {
	// The resource's type
	Type          interface{}                      `json:"type"`
	Attributes    POSTExportsRequestDataAttributes `json:"attributes"`
	Relationships interface{}                      `json:"relationships,omitempty"`
}

// NewPOSTExportsRequestData instantiates a new POSTExportsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExportsRequestData(type_ interface{}, attributes POSTExportsRequestDataAttributes) *POSTExportsRequestData {
	this := POSTExportsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTExportsRequestDataWithDefaults instantiates a new POSTExportsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExportsRequestDataWithDefaults() *POSTExportsRequestData {
	this := POSTExportsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTExportsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExportsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTExportsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTExportsRequestData) GetAttributes() POSTExportsRequestDataAttributes {
	if o == nil {
		var ret POSTExportsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTExportsRequestData) GetAttributesOk() (*POSTExportsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTExportsRequestData) SetAttributes(v POSTExportsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTExportsRequestData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTExportsRequestData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTExportsRequestData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *POSTExportsRequestData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o POSTExportsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExportsRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePOSTExportsRequestData struct {
	value *POSTExportsRequestData
	isSet bool
}

func (v NullablePOSTExportsRequestData) Get() *POSTExportsRequestData {
	return v.value
}

func (v *NullablePOSTExportsRequestData) Set(val *POSTExportsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExportsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExportsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExportsRequestData(val *POSTExportsRequestData) *NullablePOSTExportsRequestData {
	return &NullablePOSTExportsRequestData{value: val, isSet: true}
}

func (v NullablePOSTExportsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExportsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}