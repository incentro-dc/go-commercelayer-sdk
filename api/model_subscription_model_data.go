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

// checks if the SubscriptionModelData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionModelData{}

// SubscriptionModelData struct for SubscriptionModelData
type SubscriptionModelData struct {
	// The resource's type
	Type          interface{}                                                       `json:"type"`
	Attributes    GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes `json:"attributes"`
	Relationships *SubscriptionModelDataRelationships                               `json:"relationships,omitempty"`
}

// NewSubscriptionModelData instantiates a new SubscriptionModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionModelData(type_ interface{}, attributes GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) *SubscriptionModelData {
	this := SubscriptionModelData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSubscriptionModelDataWithDefaults instantiates a new SubscriptionModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionModelDataWithDefaults() *SubscriptionModelData {
	this := SubscriptionModelData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SubscriptionModelData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionModelData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionModelData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionModelData) GetAttributes() GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes {
	if o == nil {
		var ret GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionModelData) GetAttributesOk() (*GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionModelData) SetAttributes(v GETSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionModelData) GetRelationships() SubscriptionModelDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionModelDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionModelData) GetRelationshipsOk() (*SubscriptionModelDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionModelData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionModelDataRelationships and assigns it to the Relationships field.
func (o *SubscriptionModelData) SetRelationships(v SubscriptionModelDataRelationships) {
	o.Relationships = &v
}

func (o SubscriptionModelData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionModelData) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionModelData struct {
	value *SubscriptionModelData
	isSet bool
}

func (v NullableSubscriptionModelData) Get() *SubscriptionModelData {
	return v.value
}

func (v *NullableSubscriptionModelData) Set(val *SubscriptionModelData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionModelData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionModelData(val *SubscriptionModelData) *NullableSubscriptionModelData {
	return &NullableSubscriptionModelData{value: val, isSet: true}
}

func (v NullableSubscriptionModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
