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

// checks if the POSTAdyenGatewaysRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAdyenGatewaysRequestData{}

// POSTAdyenGatewaysRequestData struct for POSTAdyenGatewaysRequestData
type POSTAdyenGatewaysRequestData struct {
	// The resource's type
	Type          interface{}                                `json:"type"`
	Attributes    POSTAdyenGatewaysRequestDataAttributes     `json:"attributes"`
	Relationships *POSTAdyenGatewaysRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTAdyenGatewaysRequestData instantiates a new POSTAdyenGatewaysRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAdyenGatewaysRequestData(type_ interface{}, attributes POSTAdyenGatewaysRequestDataAttributes) *POSTAdyenGatewaysRequestData {
	this := POSTAdyenGatewaysRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTAdyenGatewaysRequestDataWithDefaults instantiates a new POSTAdyenGatewaysRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAdyenGatewaysRequestDataWithDefaults() *POSTAdyenGatewaysRequestData {
	this := POSTAdyenGatewaysRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAdyenGatewaysRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAdyenGatewaysRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTAdyenGatewaysRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTAdyenGatewaysRequestData) GetAttributes() POSTAdyenGatewaysRequestDataAttributes {
	if o == nil {
		var ret POSTAdyenGatewaysRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTAdyenGatewaysRequestData) GetAttributesOk() (*POSTAdyenGatewaysRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTAdyenGatewaysRequestData) SetAttributes(v POSTAdyenGatewaysRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTAdyenGatewaysRequestData) GetRelationships() POSTAdyenGatewaysRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTAdyenGatewaysRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAdyenGatewaysRequestData) GetRelationshipsOk() (*POSTAdyenGatewaysRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTAdyenGatewaysRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTAdyenGatewaysRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTAdyenGatewaysRequestData) SetRelationships(v POSTAdyenGatewaysRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTAdyenGatewaysRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAdyenGatewaysRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTAdyenGatewaysRequestData struct {
	value *POSTAdyenGatewaysRequestData
	isSet bool
}

func (v NullablePOSTAdyenGatewaysRequestData) Get() *POSTAdyenGatewaysRequestData {
	return v.value
}

func (v *NullablePOSTAdyenGatewaysRequestData) Set(val *POSTAdyenGatewaysRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAdyenGatewaysRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAdyenGatewaysRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAdyenGatewaysRequestData(val *POSTAdyenGatewaysRequestData) *NullablePOSTAdyenGatewaysRequestData {
	return &NullablePOSTAdyenGatewaysRequestData{value: val, isSet: true}
}

func (v NullablePOSTAdyenGatewaysRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAdyenGatewaysRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}