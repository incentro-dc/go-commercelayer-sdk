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

// GETStripeGateways200ResponseDataInner struct for GETStripeGateways200ResponseDataInner
type GETStripeGateways200ResponseDataInner struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                             `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks              `json:"links,omitempty"`
	Attributes    *GETStripeGateways200ResponseDataInnerAttributes    `json:"attributes,omitempty"`
	Relationships *GETStripeGateways200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewGETStripeGateways200ResponseDataInner instantiates a new GETStripeGateways200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStripeGateways200ResponseDataInner() *GETStripeGateways200ResponseDataInner {
	this := GETStripeGateways200ResponseDataInner{}
	return &this
}

// NewGETStripeGateways200ResponseDataInnerWithDefaults instantiates a new GETStripeGateways200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStripeGateways200ResponseDataInnerWithDefaults() *GETStripeGateways200ResponseDataInner {
	this := GETStripeGateways200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETStripeGateways200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETStripeGateways200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInner) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETStripeGateways200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInner) GetAttributes() GETStripeGateways200ResponseDataInnerAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETStripeGateways200ResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInner) GetAttributesOk() (*GETStripeGateways200ResponseDataInnerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInner) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETStripeGateways200ResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GETStripeGateways200ResponseDataInner) SetAttributes(v GETStripeGateways200ResponseDataInnerAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInner) GetRelationships() GETStripeGateways200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETStripeGateways200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInner) GetRelationshipsOk() (*GETStripeGateways200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInner) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETStripeGateways200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *GETStripeGateways200ResponseDataInner) SetRelationships(v GETStripeGateways200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o GETStripeGateways200ResponseDataInner) MarshalJSON() ([]byte, error) {
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

type NullableGETStripeGateways200ResponseDataInner struct {
	value *GETStripeGateways200ResponseDataInner
	isSet bool
}

func (v NullableGETStripeGateways200ResponseDataInner) Get() *GETStripeGateways200ResponseDataInner {
	return v.value
}

func (v *NullableGETStripeGateways200ResponseDataInner) Set(val *GETStripeGateways200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStripeGateways200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStripeGateways200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStripeGateways200ResponseDataInner(val *GETStripeGateways200ResponseDataInner) *NullableGETStripeGateways200ResponseDataInner {
	return &NullableGETStripeGateways200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGETStripeGateways200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStripeGateways200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
