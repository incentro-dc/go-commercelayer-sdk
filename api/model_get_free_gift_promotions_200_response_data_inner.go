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

// GETFreeGiftPromotions200ResponseDataInner struct for GETFreeGiftPromotions200ResponseDataInner
type GETFreeGiftPromotions200ResponseDataInner struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                                   `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                    `json:"links,omitempty"`
	Attributes    *GETFreeGiftPromotions200ResponseDataInnerAttributes      `json:"attributes,omitempty"`
	Relationships *GETFixedPricePromotions200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewGETFreeGiftPromotions200ResponseDataInner instantiates a new GETFreeGiftPromotions200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETFreeGiftPromotions200ResponseDataInner() *GETFreeGiftPromotions200ResponseDataInner {
	this := GETFreeGiftPromotions200ResponseDataInner{}
	return &this
}

// NewGETFreeGiftPromotions200ResponseDataInnerWithDefaults instantiates a new GETFreeGiftPromotions200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETFreeGiftPromotions200ResponseDataInnerWithDefaults() *GETFreeGiftPromotions200ResponseDataInner {
	this := GETFreeGiftPromotions200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETFreeGiftPromotions200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETFreeGiftPromotions200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETFreeGiftPromotions200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetAttributes() GETFreeGiftPromotions200ResponseDataInnerAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETFreeGiftPromotions200ResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetAttributesOk() (*GETFreeGiftPromotions200ResponseDataInnerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETFreeGiftPromotions200ResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GETFreeGiftPromotions200ResponseDataInner) SetAttributes(v GETFreeGiftPromotions200ResponseDataInnerAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetRelationships() GETFixedPricePromotions200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETFixedPricePromotions200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) GetRelationshipsOk() (*GETFixedPricePromotions200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETFreeGiftPromotions200ResponseDataInner) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETFixedPricePromotions200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *GETFreeGiftPromotions200ResponseDataInner) SetRelationships(v GETFixedPricePromotions200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o GETFreeGiftPromotions200ResponseDataInner) MarshalJSON() ([]byte, error) {
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

type NullableGETFreeGiftPromotions200ResponseDataInner struct {
	value *GETFreeGiftPromotions200ResponseDataInner
	isSet bool
}

func (v NullableGETFreeGiftPromotions200ResponseDataInner) Get() *GETFreeGiftPromotions200ResponseDataInner {
	return v.value
}

func (v *NullableGETFreeGiftPromotions200ResponseDataInner) Set(val *GETFreeGiftPromotions200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETFreeGiftPromotions200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETFreeGiftPromotions200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETFreeGiftPromotions200ResponseDataInner(val *GETFreeGiftPromotions200ResponseDataInner) *NullableGETFreeGiftPromotions200ResponseDataInner {
	return &NullableGETFreeGiftPromotions200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGETFreeGiftPromotions200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETFreeGiftPromotions200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
