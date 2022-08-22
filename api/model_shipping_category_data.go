/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShippingCategoryData struct for ShippingCategoryData
type ShippingCategoryData struct {
	// The resource's type
	Type          string                                                  `json:"type"`
	Attributes    GETShippingCategories200ResponseDataInnerAttributes     `json:"attributes"`
	Relationships *GETShippingCategories200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewShippingCategoryData instantiates a new ShippingCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategoryData(type_ string, attributes GETShippingCategories200ResponseDataInnerAttributes) *ShippingCategoryData {
	this := ShippingCategoryData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShippingCategoryDataWithDefaults instantiates a new ShippingCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryDataWithDefaults() *ShippingCategoryData {
	this := ShippingCategoryData{}
	var type_ string = "shipping_categories"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ShippingCategoryData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingCategoryData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingCategoryData) GetAttributes() GETShippingCategories200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETShippingCategories200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryData) GetAttributesOk() (*GETShippingCategories200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingCategoryData) SetAttributes(v GETShippingCategories200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingCategoryData) GetRelationships() GETShippingCategories200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETShippingCategories200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCategoryData) GetRelationshipsOk() (*GETShippingCategories200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingCategoryData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETShippingCategories200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *ShippingCategoryData) SetRelationships(v GETShippingCategories200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o ShippingCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableShippingCategoryData struct {
	value *ShippingCategoryData
	isSet bool
}

func (v NullableShippingCategoryData) Get() *ShippingCategoryData {
	return v.value
}

func (v *NullableShippingCategoryData) Set(val *ShippingCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategoryData(val *ShippingCategoryData) *NullableShippingCategoryData {
	return &NullableShippingCategoryData{value: val, isSet: true}
}

func (v NullableShippingCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
