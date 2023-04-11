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

// PATCHAddressesAddressId200ResponseData struct for PATCHAddressesAddressId200ResponseData
type PATCHAddressesAddressId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                       `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks            `json:"links,omitempty"`
	Attributes    *PATCHAddressesAddressId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *GETAddresses200ResponseDataInnerRelationships    `json:"relationships,omitempty"`
}

// NewPATCHAddressesAddressId200ResponseData instantiates a new PATCHAddressesAddressId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAddressesAddressId200ResponseData() *PATCHAddressesAddressId200ResponseData {
	this := PATCHAddressesAddressId200ResponseData{}
	return &this
}

// NewPATCHAddressesAddressId200ResponseDataWithDefaults instantiates a new PATCHAddressesAddressId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAddressesAddressId200ResponseDataWithDefaults() *PATCHAddressesAddressId200ResponseData {
	this := PATCHAddressesAddressId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHAddressesAddressId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHAddressesAddressId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *PATCHAddressesAddressId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseData) GetAttributes() PATCHAddressesAddressId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PATCHAddressesAddressId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseData) GetAttributesOk() (*PATCHAddressesAddressId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHAddressesAddressId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHAddressesAddressId200ResponseData) SetAttributes(v PATCHAddressesAddressId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseData) GetRelationships() GETAddresses200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETAddresses200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseData) GetRelationshipsOk() (*GETAddresses200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETAddresses200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *PATCHAddressesAddressId200ResponseData) SetRelationships(v GETAddresses200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o PATCHAddressesAddressId200ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePATCHAddressesAddressId200ResponseData struct {
	value *PATCHAddressesAddressId200ResponseData
	isSet bool
}

func (v NullablePATCHAddressesAddressId200ResponseData) Get() *PATCHAddressesAddressId200ResponseData {
	return v.value
}

func (v *NullablePATCHAddressesAddressId200ResponseData) Set(val *PATCHAddressesAddressId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAddressesAddressId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAddressesAddressId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAddressesAddressId200ResponseData(val *PATCHAddressesAddressId200ResponseData) *NullablePATCHAddressesAddressId200ResponseData {
	return &NullablePATCHAddressesAddressId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHAddressesAddressId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAddressesAddressId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
