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

// checks if the StockLocationCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLocationCreateData{}

// StockLocationCreateData struct for StockLocationCreateData
type StockLocationCreateData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    POSTStockLocations201ResponseDataAttributes `json:"attributes"`
	Relationships *MerchantCreateDataRelationships            `json:"relationships,omitempty"`
}

// NewStockLocationCreateData instantiates a new StockLocationCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLocationCreateData(type_ interface{}, attributes POSTStockLocations201ResponseDataAttributes) *StockLocationCreateData {
	this := StockLocationCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStockLocationCreateDataWithDefaults instantiates a new StockLocationCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLocationCreateDataWithDefaults() *StockLocationCreateData {
	this := StockLocationCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *StockLocationCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StockLocationCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockLocationCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StockLocationCreateData) GetAttributes() POSTStockLocations201ResponseDataAttributes {
	if o == nil {
		var ret POSTStockLocations201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockLocationCreateData) GetAttributesOk() (*POSTStockLocations201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockLocationCreateData) SetAttributes(v POSTStockLocations201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockLocationCreateData) GetRelationships() MerchantCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret MerchantCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockLocationCreateData) GetRelationshipsOk() (*MerchantCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockLocationCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given MerchantCreateDataRelationships and assigns it to the Relationships field.
func (o *StockLocationCreateData) SetRelationships(v MerchantCreateDataRelationships) {
	o.Relationships = &v
}

func (o StockLocationCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLocationCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableStockLocationCreateData struct {
	value *StockLocationCreateData
	isSet bool
}

func (v NullableStockLocationCreateData) Get() *StockLocationCreateData {
	return v.value
}

func (v *NullableStockLocationCreateData) Set(val *StockLocationCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLocationCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLocationCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLocationCreateData(val *StockLocationCreateData) *NullableStockLocationCreateData {
	return &NullableStockLocationCreateData{value: val, isSet: true}
}

func (v NullableStockLocationCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLocationCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
