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

// PATCHShippingMethodsShippingMethodId200ResponseData struct for PATCHShippingMethodsShippingMethodId200ResponseData
type PATCHShippingMethodsShippingMethodId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                    `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                         `json:"links,omitempty"`
	Attributes    *PATCHShippingMethodsShippingMethodId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *GETShippingMethods200ResponseDataInnerRelationships           `json:"relationships,omitempty"`
}

// NewPATCHShippingMethodsShippingMethodId200ResponseData instantiates a new PATCHShippingMethodsShippingMethodId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingMethodsShippingMethodId200ResponseData() *PATCHShippingMethodsShippingMethodId200ResponseData {
	this := PATCHShippingMethodsShippingMethodId200ResponseData{}
	return &this
}

// NewPATCHShippingMethodsShippingMethodId200ResponseDataWithDefaults instantiates a new PATCHShippingMethodsShippingMethodId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingMethodsShippingMethodId200ResponseDataWithDefaults() *PATCHShippingMethodsShippingMethodId200ResponseData {
	this := PATCHShippingMethodsShippingMethodId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetAttributes() PATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PATCHShippingMethodsShippingMethodId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetAttributesOk() (*PATCHShippingMethodsShippingMethodId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHShippingMethodsShippingMethodId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) SetAttributes(v PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetRelationships() GETShippingMethods200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETShippingMethods200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) GetRelationshipsOk() (*GETShippingMethods200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETShippingMethods200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *PATCHShippingMethodsShippingMethodId200ResponseData) SetRelationships(v GETShippingMethods200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o PATCHShippingMethodsShippingMethodId200ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePATCHShippingMethodsShippingMethodId200ResponseData struct {
	value *PATCHShippingMethodsShippingMethodId200ResponseData
	isSet bool
}

func (v NullablePATCHShippingMethodsShippingMethodId200ResponseData) Get() *PATCHShippingMethodsShippingMethodId200ResponseData {
	return v.value
}

func (v *NullablePATCHShippingMethodsShippingMethodId200ResponseData) Set(val *PATCHShippingMethodsShippingMethodId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingMethodsShippingMethodId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingMethodsShippingMethodId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingMethodsShippingMethodId200ResponseData(val *PATCHShippingMethodsShippingMethodId200ResponseData) *NullablePATCHShippingMethodsShippingMethodId200ResponseData {
	return &NullablePATCHShippingMethodsShippingMethodId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHShippingMethodsShippingMethodId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingMethodsShippingMethodId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
