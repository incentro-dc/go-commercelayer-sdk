/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETLineItemOptions200ResponseDataInner struct for GETLineItemOptions200ResponseDataInner
type GETLineItemOptions200ResponseDataInner struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                              `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks               `json:"links,omitempty"`
	Attributes    *GETLineItemOptions200ResponseDataInnerAttributes    `json:"attributes,omitempty"`
	Relationships *GETLineItemOptions200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewGETLineItemOptions200ResponseDataInner instantiates a new GETLineItemOptions200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItemOptions200ResponseDataInner() *GETLineItemOptions200ResponseDataInner {
	this := GETLineItemOptions200ResponseDataInner{}
	return &this
}

// NewGETLineItemOptions200ResponseDataInnerWithDefaults instantiates a new GETLineItemOptions200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItemOptions200ResponseDataInnerWithDefaults() *GETLineItemOptions200ResponseDataInner {
	this := GETLineItemOptions200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETLineItemOptions200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETLineItemOptions200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInner) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETLineItemOptions200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInner) GetAttributes() GETLineItemOptions200ResponseDataInnerAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETLineItemOptions200ResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInner) GetAttributesOk() (*GETLineItemOptions200ResponseDataInnerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInner) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETLineItemOptions200ResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GETLineItemOptions200ResponseDataInner) SetAttributes(v GETLineItemOptions200ResponseDataInnerAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETLineItemOptions200ResponseDataInner) GetRelationships() GETLineItemOptions200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItemOptions200ResponseDataInner) GetRelationshipsOk() (*GETLineItemOptions200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETLineItemOptions200ResponseDataInner) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *GETLineItemOptions200ResponseDataInner) SetRelationships(v GETLineItemOptions200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o GETLineItemOptions200ResponseDataInner) MarshalJSON() ([]byte, error) {
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

type NullableGETLineItemOptions200ResponseDataInner struct {
	value *GETLineItemOptions200ResponseDataInner
	isSet bool
}

func (v NullableGETLineItemOptions200ResponseDataInner) Get() *GETLineItemOptions200ResponseDataInner {
	return v.value
}

func (v *NullableGETLineItemOptions200ResponseDataInner) Set(val *GETLineItemOptions200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItemOptions200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItemOptions200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItemOptions200ResponseDataInner(val *GETLineItemOptions200ResponseDataInner) *NullableGETLineItemOptions200ResponseDataInner {
	return &NullableGETLineItemOptions200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGETLineItemOptions200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItemOptions200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
