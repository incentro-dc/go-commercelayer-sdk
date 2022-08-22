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

// POSTFixedAmountPromotions201ResponseData struct for POSTFixedAmountPromotions201ResponseData
type POSTFixedAmountPromotions201ResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                             `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks              `json:"links,omitempty"`
	Attributes    *POSTFixedAmountPromotions201ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTExternalPromotions201ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTFixedAmountPromotions201ResponseData instantiates a new POSTFixedAmountPromotions201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFixedAmountPromotions201ResponseData() *POSTFixedAmountPromotions201ResponseData {
	this := POSTFixedAmountPromotions201ResponseData{}
	var type_ string = "fixed_amount_promotions"
	this.Type = &type_
	return &this
}

// NewPOSTFixedAmountPromotions201ResponseDataWithDefaults instantiates a new POSTFixedAmountPromotions201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFixedAmountPromotions201ResponseDataWithDefaults() *POSTFixedAmountPromotions201ResponseData {
	this := POSTFixedAmountPromotions201ResponseData{}
	var type_ string = "fixed_amount_promotions"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *POSTFixedAmountPromotions201ResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFixedAmountPromotions201ResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotions201ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *POSTFixedAmountPromotions201ResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *POSTFixedAmountPromotions201ResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFixedAmountPromotions201ResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotions201ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *POSTFixedAmountPromotions201ResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTFixedAmountPromotions201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFixedAmountPromotions201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotions201ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *POSTFixedAmountPromotions201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTFixedAmountPromotions201ResponseData) GetAttributes() POSTFixedAmountPromotions201ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret POSTFixedAmountPromotions201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFixedAmountPromotions201ResponseData) GetAttributesOk() (*POSTFixedAmountPromotions201ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotions201ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTFixedAmountPromotions201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTFixedAmountPromotions201ResponseData) SetAttributes(v POSTFixedAmountPromotions201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTFixedAmountPromotions201ResponseData) GetRelationships() POSTExternalPromotions201ResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret POSTExternalPromotions201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTFixedAmountPromotions201ResponseData) GetRelationshipsOk() (*POSTExternalPromotions201ResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTFixedAmountPromotions201ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTExternalPromotions201ResponseDataRelationships and assigns it to the Relationships field.
func (o *POSTFixedAmountPromotions201ResponseData) SetRelationships(v POSTExternalPromotions201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o POSTFixedAmountPromotions201ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePOSTFixedAmountPromotions201ResponseData struct {
	value *POSTFixedAmountPromotions201ResponseData
	isSet bool
}

func (v NullablePOSTFixedAmountPromotions201ResponseData) Get() *POSTFixedAmountPromotions201ResponseData {
	return v.value
}

func (v *NullablePOSTFixedAmountPromotions201ResponseData) Set(val *POSTFixedAmountPromotions201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFixedAmountPromotions201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFixedAmountPromotions201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFixedAmountPromotions201ResponseData(val *POSTFixedAmountPromotions201ResponseData) *NullablePOSTFixedAmountPromotions201ResponseData {
	return &NullablePOSTFixedAmountPromotions201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTFixedAmountPromotions201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFixedAmountPromotions201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
