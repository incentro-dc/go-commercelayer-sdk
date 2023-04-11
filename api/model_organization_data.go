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

// OrganizationData struct for OrganizationData
type OrganizationData struct {
	// The resource's type
	Type          interface{}                                            `json:"type"`
	Attributes    GETOrganizationOrganizationId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                            `json:"relationships,omitempty"`
}

// NewOrganizationData instantiates a new OrganizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationData(type_ interface{}, attributes GETOrganizationOrganizationId200ResponseDataAttributes) *OrganizationData {
	this := OrganizationData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrganizationDataWithDefaults instantiates a new OrganizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDataWithDefaults() *OrganizationData {
	this := OrganizationData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrganizationData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrganizationData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrganizationData) GetAttributes() GETOrganizationOrganizationId200ResponseDataAttributes {
	if o == nil {
		var ret GETOrganizationOrganizationId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrganizationData) GetAttributesOk() (*GETOrganizationOrganizationId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrganizationData) SetAttributes(v GETOrganizationOrganizationId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrganizationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *OrganizationData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o OrganizationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
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

type NullableOrganizationData struct {
	value *OrganizationData
	isSet bool
}

func (v NullableOrganizationData) Get() *OrganizationData {
	return v.value
}

func (v *NullableOrganizationData) Set(val *OrganizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationData(val *OrganizationData) *NullableOrganizationData {
	return &NullableOrganizationData{value: val, isSet: true}
}

func (v NullableOrganizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
