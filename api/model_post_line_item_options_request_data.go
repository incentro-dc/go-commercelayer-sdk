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

// checks if the POSTLineItemOptionsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItemOptionsRequestData{}

// POSTLineItemOptionsRequestData struct for POSTLineItemOptionsRequestData
type POSTLineItemOptionsRequestData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTLineItemOptionsRequestDataAttributes     `json:"attributes"`
	Relationships *POSTLineItemOptionsRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTLineItemOptionsRequestData instantiates a new POSTLineItemOptionsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItemOptionsRequestData(type_ interface{}, attributes POSTLineItemOptionsRequestDataAttributes) *POSTLineItemOptionsRequestData {
	this := POSTLineItemOptionsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTLineItemOptionsRequestDataWithDefaults instantiates a new POSTLineItemOptionsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItemOptionsRequestDataWithDefaults() *POSTLineItemOptionsRequestData {
	this := POSTLineItemOptionsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTLineItemOptionsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTLineItemOptionsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTLineItemOptionsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTLineItemOptionsRequestData) GetAttributes() POSTLineItemOptionsRequestDataAttributes {
	if o == nil {
		var ret POSTLineItemOptionsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTLineItemOptionsRequestData) GetAttributesOk() (*POSTLineItemOptionsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTLineItemOptionsRequestData) SetAttributes(v POSTLineItemOptionsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTLineItemOptionsRequestData) GetRelationships() POSTLineItemOptionsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTLineItemOptionsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItemOptionsRequestData) GetRelationshipsOk() (*POSTLineItemOptionsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTLineItemOptionsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTLineItemOptionsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTLineItemOptionsRequestData) SetRelationships(v POSTLineItemOptionsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTLineItemOptionsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItemOptionsRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePOSTLineItemOptionsRequestData struct {
	value *POSTLineItemOptionsRequestData
	isSet bool
}

func (v NullablePOSTLineItemOptionsRequestData) Get() *POSTLineItemOptionsRequestData {
	return v.value
}

func (v *NullablePOSTLineItemOptionsRequestData) Set(val *POSTLineItemOptionsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItemOptionsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItemOptionsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItemOptionsRequestData(val *POSTLineItemOptionsRequestData) *NullablePOSTLineItemOptionsRequestData {
	return &NullablePOSTLineItemOptionsRequestData{value: val, isSet: true}
}

func (v NullablePOSTLineItemOptionsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItemOptionsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}