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

// GETPriceLists200ResponseDataInner struct for GETPriceLists200ResponseDataInner
type GETPriceLists200ResponseDataInner struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                     `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks          `json:"links,omitempty"`
	Attributes    *GETPriceLists200ResponseDataInnerAttributes    `json:"attributes,omitempty"`
	Relationships *GETPriceLists200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewGETPriceLists200ResponseDataInner instantiates a new GETPriceLists200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceLists200ResponseDataInner() *GETPriceLists200ResponseDataInner {
	this := GETPriceLists200ResponseDataInner{}
	return &this
}

// NewGETPriceLists200ResponseDataInnerWithDefaults instantiates a new GETPriceLists200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceLists200ResponseDataInnerWithDefaults() *GETPriceLists200ResponseDataInner {
	this := GETPriceLists200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceLists200ResponseDataInner) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceLists200ResponseDataInner) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETPriceLists200ResponseDataInner) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceLists200ResponseDataInner) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceLists200ResponseDataInner) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETPriceLists200ResponseDataInner) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInner) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETPriceLists200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInner) GetAttributes() GETPriceLists200ResponseDataInnerAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETPriceLists200ResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInner) GetAttributesOk() (*GETPriceLists200ResponseDataInnerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInner) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETPriceLists200ResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GETPriceLists200ResponseDataInner) SetAttributes(v GETPriceLists200ResponseDataInnerAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInner) GetRelationships() GETPriceLists200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETPriceLists200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInner) GetRelationshipsOk() (*GETPriceLists200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInner) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETPriceLists200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *GETPriceLists200ResponseDataInner) SetRelationships(v GETPriceLists200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o GETPriceLists200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceLists200ResponseDataInner struct {
	value *GETPriceLists200ResponseDataInner
	isSet bool
}

func (v NullableGETPriceLists200ResponseDataInner) Get() *GETPriceLists200ResponseDataInner {
	return v.value
}

func (v *NullableGETPriceLists200ResponseDataInner) Set(val *GETPriceLists200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceLists200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceLists200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceLists200ResponseDataInner(val *GETPriceLists200ResponseDataInner) *NullableGETPriceLists200ResponseDataInner {
	return &NullableGETPriceLists200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGETPriceLists200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceLists200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
