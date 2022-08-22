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

// BundleData struct for BundleData
type BundleData struct {
	// The resource's type
	Type          string                                       `json:"type"`
	Attributes    GETBundles200ResponseDataInnerAttributes     `json:"attributes"`
	Relationships *GETBundles200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewBundleData instantiates a new BundleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleData(type_ string, attributes GETBundles200ResponseDataInnerAttributes) *BundleData {
	this := BundleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBundleDataWithDefaults instantiates a new BundleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataWithDefaults() *BundleData {
	this := BundleData{}
	var type_ string = "bundles"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BundleData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BundleData) GetAttributes() GETBundles200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETBundles200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BundleData) GetAttributesOk() (*GETBundles200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BundleData) SetAttributes(v GETBundles200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BundleData) GetRelationships() GETBundles200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETBundles200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleData) GetRelationshipsOk() (*GETBundles200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BundleData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETBundles200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *BundleData) SetRelationships(v GETBundles200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o BundleData) MarshalJSON() ([]byte, error) {
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

type NullableBundleData struct {
	value *BundleData
	isSet bool
}

func (v NullableBundleData) Get() *BundleData {
	return v.value
}

func (v *NullableBundleData) Set(val *BundleData) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleData) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleData(val *BundleData) *NullableBundleData {
	return &NullableBundleData{value: val, isSet: true}
}

func (v NullableBundleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
