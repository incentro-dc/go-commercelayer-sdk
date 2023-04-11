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

// checks if the POSTShippingMethodsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethodsRequestData{}

// POSTShippingMethodsRequestData struct for POSTShippingMethodsRequestData
type POSTShippingMethodsRequestData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTShippingMethodsRequestDataAttributes     `json:"attributes"`
	Relationships *POSTShippingMethodsRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTShippingMethodsRequestData instantiates a new POSTShippingMethodsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethodsRequestData(type_ interface{}, attributes POSTShippingMethodsRequestDataAttributes) *POSTShippingMethodsRequestData {
	this := POSTShippingMethodsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTShippingMethodsRequestDataWithDefaults instantiates a new POSTShippingMethodsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethodsRequestDataWithDefaults() *POSTShippingMethodsRequestData {
	this := POSTShippingMethodsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTShippingMethodsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingMethodsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTShippingMethodsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTShippingMethodsRequestData) GetAttributes() POSTShippingMethodsRequestDataAttributes {
	if o == nil {
		var ret POSTShippingMethodsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestData) GetAttributesOk() (*POSTShippingMethodsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTShippingMethodsRequestData) SetAttributes(v POSTShippingMethodsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTShippingMethodsRequestData) GetRelationships() POSTShippingMethodsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTShippingMethodsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethodsRequestData) GetRelationshipsOk() (*POSTShippingMethodsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTShippingMethodsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTShippingMethodsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTShippingMethodsRequestData) SetRelationships(v POSTShippingMethodsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTShippingMethodsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethodsRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTShippingMethodsRequestData struct {
	value *POSTShippingMethodsRequestData
	isSet bool
}

func (v NullablePOSTShippingMethodsRequestData) Get() *POSTShippingMethodsRequestData {
	return v.value
}

func (v *NullablePOSTShippingMethodsRequestData) Set(val *POSTShippingMethodsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethodsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethodsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethodsRequestData(val *POSTShippingMethodsRequestData) *NullablePOSTShippingMethodsRequestData {
	return &NullablePOSTShippingMethodsRequestData{value: val, isSet: true}
}

func (v NullablePOSTShippingMethodsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethodsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}