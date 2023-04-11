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

// checks if the PATCHTaxCategoriesTaxCategoryIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHTaxCategoriesTaxCategoryIdRequestData{}

// PATCHTaxCategoriesTaxCategoryIdRequestData struct for PATCHTaxCategoriesTaxCategoryIdRequestData
type PATCHTaxCategoriesTaxCategoryIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                              `json:"id"`
	Attributes    PATCHTaxCategoriesTaxCategoryIdRequestDataAttributes     `json:"attributes"`
	Relationships *PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHTaxCategoriesTaxCategoryIdRequestData instantiates a new PATCHTaxCategoriesTaxCategoryIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHTaxCategoriesTaxCategoryIdRequestData(type_ interface{}, id interface{}, attributes PATCHTaxCategoriesTaxCategoryIdRequestDataAttributes) *PATCHTaxCategoriesTaxCategoryIdRequestData {
	this := PATCHTaxCategoriesTaxCategoryIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHTaxCategoriesTaxCategoryIdRequestDataWithDefaults instantiates a new PATCHTaxCategoriesTaxCategoryIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHTaxCategoriesTaxCategoryIdRequestDataWithDefaults() *PATCHTaxCategoriesTaxCategoryIdRequestData {
	this := PATCHTaxCategoriesTaxCategoryIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetAttributes() PATCHTaxCategoriesTaxCategoryIdRequestDataAttributes {
	if o == nil {
		var ret PATCHTaxCategoriesTaxCategoryIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetAttributesOk() (*PATCHTaxCategoriesTaxCategoryIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) SetAttributes(v PATCHTaxCategoriesTaxCategoryIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetRelationships() PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) GetRelationshipsOk() (*PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHTaxCategoriesTaxCategoryIdRequestData) SetRelationships(v PATCHTaxCategoriesTaxCategoryIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHTaxCategoriesTaxCategoryIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHTaxCategoriesTaxCategoryIdRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHTaxCategoriesTaxCategoryIdRequestData struct {
	value *PATCHTaxCategoriesTaxCategoryIdRequestData
	isSet bool
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequestData) Get() *PATCHTaxCategoriesTaxCategoryIdRequestData {
	return v.value
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequestData) Set(val *PATCHTaxCategoriesTaxCategoryIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHTaxCategoriesTaxCategoryIdRequestData(val *PATCHTaxCategoriesTaxCategoryIdRequestData) *NullablePATCHTaxCategoriesTaxCategoryIdRequestData {
	return &NullablePATCHTaxCategoriesTaxCategoryIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHTaxCategoriesTaxCategoryIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHTaxCategoriesTaxCategoryIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}