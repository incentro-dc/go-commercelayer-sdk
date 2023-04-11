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

// GETCheckoutComGateways200ResponseDataInner struct for GETCheckoutComGateways200ResponseDataInner
type GETCheckoutComGateways200ResponseDataInner struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                              `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                   `json:"links,omitempty"`
	Attributes    *GETCheckoutComGateways200ResponseDataInnerAttributes    `json:"attributes,omitempty"`
	Relationships *GETCheckoutComGateways200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewGETCheckoutComGateways200ResponseDataInner instantiates a new GETCheckoutComGateways200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200ResponseDataInner() *GETCheckoutComGateways200ResponseDataInner {
	this := GETCheckoutComGateways200ResponseDataInner{}
	return &this
}

// NewGETCheckoutComGateways200ResponseDataInnerWithDefaults instantiates a new GETCheckoutComGateways200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseDataInnerWithDefaults() *GETCheckoutComGateways200ResponseDataInner {
	this := GETCheckoutComGateways200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGateways200ResponseDataInner) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGateways200ResponseDataInner) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETCheckoutComGateways200ResponseDataInner) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCheckoutComGateways200ResponseDataInner) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCheckoutComGateways200ResponseDataInner) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETCheckoutComGateways200ResponseDataInner) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *GETCheckoutComGateways200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInner) GetAttributes() GETCheckoutComGateways200ResponseDataInnerAttributes {
	if o == nil || o.Attributes == nil {
		var ret GETCheckoutComGateways200ResponseDataInnerAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) GetAttributesOk() (*GETCheckoutComGateways200ResponseDataInnerAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GETCheckoutComGateways200ResponseDataInnerAttributes and assigns it to the Attributes field.
func (o *GETCheckoutComGateways200ResponseDataInner) SetAttributes(v GETCheckoutComGateways200ResponseDataInnerAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInner) GetRelationships() GETCheckoutComGateways200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETCheckoutComGateways200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) GetRelationshipsOk() (*GETCheckoutComGateways200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInner) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETCheckoutComGateways200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *GETCheckoutComGateways200ResponseDataInner) SetRelationships(v GETCheckoutComGateways200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o GETCheckoutComGateways200ResponseDataInner) MarshalJSON() ([]byte, error) {
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

type NullableGETCheckoutComGateways200ResponseDataInner struct {
	value *GETCheckoutComGateways200ResponseDataInner
	isSet bool
}

func (v NullableGETCheckoutComGateways200ResponseDataInner) Get() *GETCheckoutComGateways200ResponseDataInner {
	return v.value
}

func (v *NullableGETCheckoutComGateways200ResponseDataInner) Set(val *GETCheckoutComGateways200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200ResponseDataInner(val *GETCheckoutComGateways200ResponseDataInner) *NullableGETCheckoutComGateways200ResponseDataInner {
	return &NullableGETCheckoutComGateways200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
