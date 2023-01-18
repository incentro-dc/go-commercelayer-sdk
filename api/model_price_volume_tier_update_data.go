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

// PriceVolumeTierUpdateData struct for PriceVolumeTierUpdateData
type PriceVolumeTierUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                          `json:"id"`
	Attributes    PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes `json:"attributes"`
	Relationships *PriceVolumeTierUpdateDataRelationships                         `json:"relationships,omitempty"`
}

// NewPriceVolumeTierUpdateData instantiates a new PriceVolumeTierUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierUpdateData(type_ string, id string, attributes PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) *PriceVolumeTierUpdateData {
	this := PriceVolumeTierUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPriceVolumeTierUpdateDataWithDefaults instantiates a new PriceVolumeTierUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierUpdateDataWithDefaults() *PriceVolumeTierUpdateData {
	this := PriceVolumeTierUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *PriceVolumeTierUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceVolumeTierUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PriceVolumeTierUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PriceVolumeTierUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PriceVolumeTierUpdateData) GetAttributes() PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierUpdateData) GetAttributesOk() (*PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceVolumeTierUpdateData) SetAttributes(v PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceVolumeTierUpdateData) GetRelationships() PriceVolumeTierUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceVolumeTierUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierUpdateData) GetRelationshipsOk() (*PriceVolumeTierUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceVolumeTierUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceVolumeTierUpdateDataRelationships and assigns it to the Relationships field.
func (o *PriceVolumeTierUpdateData) SetRelationships(v PriceVolumeTierUpdateDataRelationships) {
	o.Relationships = &v
}

func (o PriceVolumeTierUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePriceVolumeTierUpdateData struct {
	value *PriceVolumeTierUpdateData
	isSet bool
}

func (v NullablePriceVolumeTierUpdateData) Get() *PriceVolumeTierUpdateData {
	return v.value
}

func (v *NullablePriceVolumeTierUpdateData) Set(val *PriceVolumeTierUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierUpdateData(val *PriceVolumeTierUpdateData) *NullablePriceVolumeTierUpdateData {
	return &NullablePriceVolumeTierUpdateData{value: val, isSet: true}
}

func (v NullablePriceVolumeTierUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
