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

// checks if the POSTReturnLineItemsRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTReturnLineItemsRequestData{}

// POSTReturnLineItemsRequestData struct for POSTReturnLineItemsRequestData
type POSTReturnLineItemsRequestData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTReturnLineItemsRequestDataAttributes     `json:"attributes"`
	Relationships *POSTReturnLineItemsRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTReturnLineItemsRequestData instantiates a new POSTReturnLineItemsRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTReturnLineItemsRequestData(type_ interface{}, attributes POSTReturnLineItemsRequestDataAttributes) *POSTReturnLineItemsRequestData {
	this := POSTReturnLineItemsRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPOSTReturnLineItemsRequestDataWithDefaults instantiates a new POSTReturnLineItemsRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTReturnLineItemsRequestDataWithDefaults() *POSTReturnLineItemsRequestData {
	this := POSTReturnLineItemsRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTReturnLineItemsRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTReturnLineItemsRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *POSTReturnLineItemsRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *POSTReturnLineItemsRequestData) GetAttributes() POSTReturnLineItemsRequestDataAttributes {
	if o == nil {
		var ret POSTReturnLineItemsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItemsRequestData) GetAttributesOk() (*POSTReturnLineItemsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *POSTReturnLineItemsRequestData) SetAttributes(v POSTReturnLineItemsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTReturnLineItemsRequestData) GetRelationships() POSTReturnLineItemsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTReturnLineItemsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTReturnLineItemsRequestData) GetRelationshipsOk() (*POSTReturnLineItemsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTReturnLineItemsRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTReturnLineItemsRequestDataRelationships and assigns it to the Relationships field.
func (o *POSTReturnLineItemsRequestData) SetRelationships(v POSTReturnLineItemsRequestDataRelationships) {
	o.Relationships = &v
}

func (o POSTReturnLineItemsRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTReturnLineItemsRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTReturnLineItemsRequestData struct {
	value *POSTReturnLineItemsRequestData
	isSet bool
}

func (v NullablePOSTReturnLineItemsRequestData) Get() *POSTReturnLineItemsRequestData {
	return v.value
}

func (v *NullablePOSTReturnLineItemsRequestData) Set(val *POSTReturnLineItemsRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTReturnLineItemsRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTReturnLineItemsRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTReturnLineItemsRequestData(val *POSTReturnLineItemsRequestData) *NullablePOSTReturnLineItemsRequestData {
	return &NullablePOSTReturnLineItemsRequestData{value: val, isSet: true}
}

func (v NullablePOSTReturnLineItemsRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTReturnLineItemsRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}