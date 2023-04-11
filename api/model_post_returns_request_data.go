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

// checks if the POSTReturnsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTReturnsRequestData{}

// POSTReturnsRequestData struct for POSTReturnsRequestData
type POSTReturnsRequestData struct {
	// The resource's type
	Type          interface{}                            `json:"type"`
	Attributes    POSTAdyenPaymentsRequestDataAttributes `json:"attributes"`
	Relationships *POSTReturnsRequestDataRelationships   `json:"relationships,omitempty"`
}

// NewPOSTReturnsRequestData instantiates a new POSTReturnsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturnsRequestData(type_ interface{}, attributes POSTAdyenPaymentsRequestDataAttributes) *POSTReturnsRequestData {
	this := POSTReturnsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTReturnsRequestDataWithDefaults instantiates a new POSTReturnsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturnsRequestDataWithDefaults() *POSTReturnsRequestData {
	this := POSTReturnsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTReturnsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTReturnsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTReturnsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTReturnsRequestData) GetAttributes() POSTAdyenPaymentsRequestDataAttributes {
	if o == nil {
		var ret POSTAdyenPaymentsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTReturnsRequestData) GetAttributesOk() (*POSTAdyenPaymentsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTReturnsRequestData) SetAttributes(v POSTAdyenPaymentsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTReturnsRequestData) GetRelationships() POSTReturnsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTReturnsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnsRequestData) GetRelationshipsOk() (*POSTReturnsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTReturnsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTReturnsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTReturnsRequestData) SetRelationships(v POSTReturnsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTReturnsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTReturnsRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTReturnsRequestData struct {
	value *POSTReturnsRequestData
	isSet bool
}

func (v NullablePOSTReturnsRequestData) Get() *POSTReturnsRequestData {
	return v.value
}

func (v *NullablePOSTReturnsRequestData) Set(val *POSTReturnsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturnsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturnsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturnsRequestData(val *POSTReturnsRequestData) *NullablePOSTReturnsRequestData {
	return &NullablePOSTReturnsRequestData{value: val, isSet: true}
}

func (v NullablePOSTReturnsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturnsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}