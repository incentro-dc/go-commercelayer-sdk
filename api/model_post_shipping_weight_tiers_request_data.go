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

// checks if the POSTShippingWeightTiersRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingWeightTiersRequestData{}

// POSTShippingWeightTiersRequestData struct for POSTShippingWeightTiersRequestData
type POSTShippingWeightTiersRequestData struct {
	// The resource's type
	Type          interface{}                                      `json:"type"`
	Attributes    POSTShippingWeightTiersRequestDataAttributes     `json:"attributes"`
	Relationships *POSTShippingWeightTiersRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTShippingWeightTiersRequestData instantiates a new POSTShippingWeightTiersRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingWeightTiersRequestData(type_ interface{}, attributes POSTShippingWeightTiersRequestDataAttributes) *POSTShippingWeightTiersRequestData {
	this := POSTShippingWeightTiersRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTShippingWeightTiersRequestDataWithDefaults instantiates a new POSTShippingWeightTiersRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingWeightTiersRequestDataWithDefaults() *POSTShippingWeightTiersRequestData {
	this := POSTShippingWeightTiersRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTShippingWeightTiersRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingWeightTiersRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTShippingWeightTiersRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTShippingWeightTiersRequestData) GetAttributes() POSTShippingWeightTiersRequestDataAttributes {
	if o == nil {
		var ret POSTShippingWeightTiersRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTShippingWeightTiersRequestData) GetAttributesOk() (*POSTShippingWeightTiersRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTShippingWeightTiersRequestData) SetAttributes(v POSTShippingWeightTiersRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTShippingWeightTiersRequestData) GetRelationships() POSTShippingWeightTiersRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTShippingWeightTiersRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingWeightTiersRequestData) GetRelationshipsOk() (*POSTShippingWeightTiersRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTShippingWeightTiersRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTShippingWeightTiersRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTShippingWeightTiersRequestData) SetRelationships(v POSTShippingWeightTiersRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTShippingWeightTiersRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingWeightTiersRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTShippingWeightTiersRequestData struct {
	value *POSTShippingWeightTiersRequestData
	isSet bool
}

func (v NullablePOSTShippingWeightTiersRequestData) Get() *POSTShippingWeightTiersRequestData {
	return v.value
}

func (v *NullablePOSTShippingWeightTiersRequestData) Set(val *POSTShippingWeightTiersRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingWeightTiersRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingWeightTiersRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingWeightTiersRequestData(val *POSTShippingWeightTiersRequestData) *NullablePOSTShippingWeightTiersRequestData {
	return &NullablePOSTShippingWeightTiersRequestData{value: val, isSet: true}
}

func (v NullablePOSTShippingWeightTiersRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingWeightTiersRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}