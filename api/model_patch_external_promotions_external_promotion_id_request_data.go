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

// checks if the PATCHExternalPromotionsExternalPromotionIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHExternalPromotionsExternalPromotionIdRequestData{}

// PATCHExternalPromotionsExternalPromotionIdRequestData struct for PATCHExternalPromotionsExternalPromotionIdRequestData
type PATCHExternalPromotionsExternalPromotionIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                     `json:"id"`
	Attributes    PATCHExternalPromotionsExternalPromotionIdRequestDataAttributes `json:"attributes"`
	Relationships *POSTExternalPromotionsRequestDataRelationships                 `json:"relationships,omitempty"`
}

// NewPATCHExternalPromotionsExternalPromotionIdRequestData instantiates a new PATCHExternalPromotionsExternalPromotionIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalPromotionsExternalPromotionIdRequestData(type_ interface{}, id interface{}, attributes PATCHExternalPromotionsExternalPromotionIdRequestDataAttributes) *PATCHExternalPromotionsExternalPromotionIdRequestData {
	this := PATCHExternalPromotionsExternalPromotionIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHExternalPromotionsExternalPromotionIdRequestDataWithDefaults instantiates a new PATCHExternalPromotionsExternalPromotionIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalPromotionsExternalPromotionIdRequestDataWithDefaults() *PATCHExternalPromotionsExternalPromotionIdRequestData {
	this := PATCHExternalPromotionsExternalPromotionIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetAttributes() PATCHExternalPromotionsExternalPromotionIdRequestDataAttributes {
	if o == nil {
		var ret PATCHExternalPromotionsExternalPromotionIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetAttributesOk() (*PATCHExternalPromotionsExternalPromotionIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) SetAttributes(v PATCHExternalPromotionsExternalPromotionIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetRelationships() POSTExternalPromotionsRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTExternalPromotionsRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) GetRelationshipsOk() (*POSTExternalPromotionsRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTExternalPromotionsRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHExternalPromotionsExternalPromotionIdRequestData) SetRelationships(v POSTExternalPromotionsRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHExternalPromotionsExternalPromotionIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHExternalPromotionsExternalPromotionIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHExternalPromotionsExternalPromotionIdRequestData struct {
	value *PATCHExternalPromotionsExternalPromotionIdRequestData
	isSet bool
}

func (v NullablePATCHExternalPromotionsExternalPromotionIdRequestData) Get() *PATCHExternalPromotionsExternalPromotionIdRequestData {
	return v.value
}

func (v *NullablePATCHExternalPromotionsExternalPromotionIdRequestData) Set(val *PATCHExternalPromotionsExternalPromotionIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalPromotionsExternalPromotionIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalPromotionsExternalPromotionIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalPromotionsExternalPromotionIdRequestData(val *PATCHExternalPromotionsExternalPromotionIdRequestData) *NullablePATCHExternalPromotionsExternalPromotionIdRequestData {
	return &NullablePATCHExternalPromotionsExternalPromotionIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHExternalPromotionsExternalPromotionIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalPromotionsExternalPromotionIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}