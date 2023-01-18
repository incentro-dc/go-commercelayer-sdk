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

// PriceListUpdateData struct for PriceListUpdateData
type PriceListUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                              `json:"id"`
	Attributes    PATCHPriceListsPriceListId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                              `json:"relationships,omitempty"`
}

// NewPriceListUpdateData instantiates a new PriceListUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListUpdateData(type_ string, id string, attributes PATCHPriceListsPriceListId200ResponseDataAttributes) *PriceListUpdateData {
	this := PriceListUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPriceListUpdateDataWithDefaults instantiates a new PriceListUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListUpdateDataWithDefaults() *PriceListUpdateData {
	this := PriceListUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *PriceListUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceListUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceListUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PriceListUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PriceListUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PriceListUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PriceListUpdateData) GetAttributes() PATCHPriceListsPriceListId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHPriceListsPriceListId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceListUpdateData) GetAttributesOk() (*PATCHPriceListsPriceListId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceListUpdateData) SetAttributes(v PATCHPriceListsPriceListId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceListUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceListUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *PriceListUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o PriceListUpdateData) MarshalJSON() ([]byte, error) {
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

type NullablePriceListUpdateData struct {
	value *PriceListUpdateData
	isSet bool
}

func (v NullablePriceListUpdateData) Get() *PriceListUpdateData {
	return v.value
}

func (v *NullablePriceListUpdateData) Set(val *PriceListUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListUpdateData(val *PriceListUpdateData) *NullablePriceListUpdateData {
	return &NullablePriceListUpdateData{value: val, isSet: true}
}

func (v NullablePriceListUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
