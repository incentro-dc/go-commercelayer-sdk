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

// POSTPriceLists201ResponseData struct for POSTPriceLists201ResponseData
type POSTPriceLists201ResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                  `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks   `json:"links,omitempty"`
	Attributes    *POSTPriceLists201ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships map[string]interface{}                   `json:"relationships,omitempty"`
}

// NewPOSTPriceLists201ResponseData instantiates a new POSTPriceLists201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceLists201ResponseData() *POSTPriceLists201ResponseData {
	this := POSTPriceLists201ResponseData{}
	var type_ string = "price_lists"
	this.Type = &type_
	return &this
}

// NewPOSTPriceLists201ResponseDataWithDefaults instantiates a new POSTPriceLists201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceLists201ResponseDataWithDefaults() *POSTPriceLists201ResponseData {
	this := POSTPriceLists201ResponseData{}
	var type_ string = "price_lists"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *POSTPriceLists201ResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *POSTPriceLists201ResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *POSTPriceLists201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseData) GetAttributes() POSTPriceLists201ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret POSTPriceLists201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseData) GetAttributesOk() (*POSTPriceLists201ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTPriceLists201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTPriceLists201ResponseData) SetAttributes(v POSTPriceLists201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTPriceLists201ResponseData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPriceLists201ResponseData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *POSTPriceLists201ResponseData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o POSTPriceLists201ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePOSTPriceLists201ResponseData struct {
	value *POSTPriceLists201ResponseData
	isSet bool
}

func (v NullablePOSTPriceLists201ResponseData) Get() *POSTPriceLists201ResponseData {
	return v.value
}

func (v *NullablePOSTPriceLists201ResponseData) Set(val *POSTPriceLists201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceLists201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceLists201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceLists201ResponseData(val *POSTPriceLists201ResponseData) *NullablePOSTPriceLists201ResponseData {
	return &NullablePOSTPriceLists201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTPriceLists201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceLists201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
