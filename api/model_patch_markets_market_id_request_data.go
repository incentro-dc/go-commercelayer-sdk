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

// checks if the PATCHMarketsMarketIdRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHMarketsMarketIdRequestData{}

// PATCHMarketsMarketIdRequestData struct for PATCHMarketsMarketIdRequestData
type PATCHMarketsMarketIdRequestData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                   `json:"id"`
	Attributes    PATCHMarketsMarketIdRequestDataAttributes     `json:"attributes"`
	Relationships *PATCHMarketsMarketIdRequestDataRelationships `json:"relationships,omitempty"`
}

// NewPATCHMarketsMarketIdRequestData instantiates a new PATCHMarketsMarketIdRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHMarketsMarketIdRequestData(type_ interface{}, id interface{}, attributes PATCHMarketsMarketIdRequestDataAttributes) *PATCHMarketsMarketIdRequestData {
	this := PATCHMarketsMarketIdRequestData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPATCHMarketsMarketIdRequestDataWithDefaults instantiates a new PATCHMarketsMarketIdRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHMarketsMarketIdRequestDataWithDefaults() *PATCHMarketsMarketIdRequestData {
	this := PATCHMarketsMarketIdRequestData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHMarketsMarketIdRequestData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHMarketsMarketIdRequestData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PATCHMarketsMarketIdRequestData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PATCHMarketsMarketIdRequestData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHMarketsMarketIdRequestData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PATCHMarketsMarketIdRequestData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PATCHMarketsMarketIdRequestData) GetAttributes() PATCHMarketsMarketIdRequestDataAttributes {
	if o == nil {
		var ret PATCHMarketsMarketIdRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PATCHMarketsMarketIdRequestData) GetAttributesOk() (*PATCHMarketsMarketIdRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PATCHMarketsMarketIdRequestData) SetAttributes(v PATCHMarketsMarketIdRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHMarketsMarketIdRequestData) GetRelationships() PATCHMarketsMarketIdRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PATCHMarketsMarketIdRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHMarketsMarketIdRequestData) GetRelationshipsOk() (*PATCHMarketsMarketIdRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHMarketsMarketIdRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHMarketsMarketIdRequestDataRelationships and assigns it to the Relationships field.
func (o *PATCHMarketsMarketIdRequestData) SetRelationships(v PATCHMarketsMarketIdRequestDataRelationships) {
	o.Relationships = &v
}

func (o PATCHMarketsMarketIdRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHMarketsMarketIdRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePATCHMarketsMarketIdRequestData struct {
	value *PATCHMarketsMarketIdRequestData
	isSet bool
}

func (v NullablePATCHMarketsMarketIdRequestData) Get() *PATCHMarketsMarketIdRequestData {
	return v.value
}

func (v *NullablePATCHMarketsMarketIdRequestData) Set(val *PATCHMarketsMarketIdRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHMarketsMarketIdRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHMarketsMarketIdRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHMarketsMarketIdRequestData(val *PATCHMarketsMarketIdRequestData) *NullablePATCHMarketsMarketIdRequestData {
	return &NullablePATCHMarketsMarketIdRequestData{value: val, isSet: true}
}

func (v NullablePATCHMarketsMarketIdRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHMarketsMarketIdRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}