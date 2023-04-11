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

// checks if the POSTInventoryReturnLocationsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTInventoryReturnLocationsRequestData{}

// POSTInventoryReturnLocationsRequestData struct for POSTInventoryReturnLocationsRequestData
type POSTInventoryReturnLocationsRequestData struct {
	// The resource's type
	Type          interface{}                                           `json:"type"`
	Attributes    POSTInventoryReturnLocationsRequestDataAttributes     `json:"attributes"`
	Relationships *POSTInventoryReturnLocationsRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTInventoryReturnLocationsRequestData instantiates a new POSTInventoryReturnLocationsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryReturnLocationsRequestData(type_ interface{}, attributes POSTInventoryReturnLocationsRequestDataAttributes) *POSTInventoryReturnLocationsRequestData {
	this := POSTInventoryReturnLocationsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTInventoryReturnLocationsRequestDataWithDefaults instantiates a new POSTInventoryReturnLocationsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryReturnLocationsRequestDataWithDefaults() *POSTInventoryReturnLocationsRequestData {
	this := POSTInventoryReturnLocationsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTInventoryReturnLocationsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTInventoryReturnLocationsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTInventoryReturnLocationsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTInventoryReturnLocationsRequestData) GetAttributes() POSTInventoryReturnLocationsRequestDataAttributes {
	if o == nil {
		var ret POSTInventoryReturnLocationsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTInventoryReturnLocationsRequestData) GetAttributesOk() (*POSTInventoryReturnLocationsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTInventoryReturnLocationsRequestData) SetAttributes(v POSTInventoryReturnLocationsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTInventoryReturnLocationsRequestData) GetRelationships() POSTInventoryReturnLocationsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTInventoryReturnLocationsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTInventoryReturnLocationsRequestData) GetRelationshipsOk() (*POSTInventoryReturnLocationsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTInventoryReturnLocationsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTInventoryReturnLocationsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTInventoryReturnLocationsRequestData) SetRelationships(v POSTInventoryReturnLocationsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTInventoryReturnLocationsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTInventoryReturnLocationsRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTInventoryReturnLocationsRequestData struct {
	value *POSTInventoryReturnLocationsRequestData
	isSet bool
}

func (v NullablePOSTInventoryReturnLocationsRequestData) Get() *POSTInventoryReturnLocationsRequestData {
	return v.value
}

func (v *NullablePOSTInventoryReturnLocationsRequestData) Set(val *POSTInventoryReturnLocationsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryReturnLocationsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryReturnLocationsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryReturnLocationsRequestData(val *POSTInventoryReturnLocationsRequestData) *NullablePOSTInventoryReturnLocationsRequestData {
	return &NullablePOSTInventoryReturnLocationsRequestData{value: val, isSet: true}
}

func (v NullablePOSTInventoryReturnLocationsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryReturnLocationsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}