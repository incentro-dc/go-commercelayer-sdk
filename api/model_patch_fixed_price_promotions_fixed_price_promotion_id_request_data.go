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

// checks if the PATCHFixedPricePromotionsFixedPricePromotionIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHFixedPricePromotionsFixedPricePromotionIdRequestData{}

// PATCHFixedPricePromotionsFixedPricePromotionIdRequestData struct for PATCHFixedPricePromotionsFixedPricePromotionIdRequestData
type PATCHFixedPricePromotionsFixedPricePromotionIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                             `json:"id"`
	Attributes    PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes     `json:"attributes"`
	Relationships *PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestData instantiates a new PATCHFixedPricePromotionsFixedPricePromotionIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestData(type_ interface{}, id interface{}, attributes PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes) *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData {
	this := PATCHFixedPricePromotionsFixedPricePromotionIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestDataWithDefaults instantiates a new PATCHFixedPricePromotionsFixedPricePromotionIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFixedPricePromotionsFixedPricePromotionIdRequestDataWithDefaults() *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData {
	this := PATCHFixedPricePromotionsFixedPricePromotionIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetAttributes() PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes {
	if o == nil {
		var ret PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetAttributesOk() (*PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) SetAttributes(v PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetRelationships() PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) GetRelationshipsOk() (*PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) SetRelationships(v PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData struct {
	value *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData
	isSet bool
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData) Get() *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData {
	return v.value
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData) Set(val *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData(val *PATCHFixedPricePromotionsFixedPricePromotionIdRequestData) *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData {
	return &NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFixedPricePromotionsFixedPricePromotionIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}