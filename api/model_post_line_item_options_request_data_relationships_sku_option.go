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

// checks if the POSTLineItemOptionsRequestDataRelationshipsSkuOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItemOptionsRequestDataRelationshipsSkuOption{}

// POSTLineItemOptionsRequestDataRelationshipsSkuOption struct for POSTLineItemOptionsRequestDataRelationshipsSkuOption
type POSTLineItemOptionsRequestDataRelationshipsSkuOption struct {
	Data POSTLineItemOptionsRequestDataRelationshipsSkuOptionData `json:"data"`
}

// NewPOSTLineItemOptionsRequestDataRelationshipsSkuOption instantiates a new POSTLineItemOptionsRequestDataRelationshipsSkuOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItemOptionsRequestDataRelationshipsSkuOption(data POSTLineItemOptionsRequestDataRelationshipsSkuOptionData) *POSTLineItemOptionsRequestDataRelationshipsSkuOption {
	this := POSTLineItemOptionsRequestDataRelationshipsSkuOption{}
	this.Data = data
	return &this
}

// NewPOSTLineItemOptionsRequestDataRelationshipsSkuOptionWithDefaults instantiates a new POSTLineItemOptionsRequestDataRelationshipsSkuOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItemOptionsRequestDataRelationshipsSkuOptionWithDefaults() *POSTLineItemOptionsRequestDataRelationshipsSkuOption {
	this := POSTLineItemOptionsRequestDataRelationshipsSkuOption{}
	return &this
}

// GetData returns the Data field value
func (o *POSTLineItemOptionsRequestDataRelationshipsSkuOption) GetData() POSTLineItemOptionsRequestDataRelationshipsSkuOptionData {
	if o == nil {
		var ret POSTLineItemOptionsRequestDataRelationshipsSkuOptionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTLineItemOptionsRequestDataRelationshipsSkuOption) GetDataOk() (*POSTLineItemOptionsRequestDataRelationshipsSkuOptionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTLineItemOptionsRequestDataRelationshipsSkuOption) SetData(v POSTLineItemOptionsRequestDataRelationshipsSkuOptionData) {
	o.Data = v
}

func (o POSTLineItemOptionsRequestDataRelationshipsSkuOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItemOptionsRequestDataRelationshipsSkuOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption struct {
	value *POSTLineItemOptionsRequestDataRelationshipsSkuOption
	isSet bool
}

func (v NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption) Get() *POSTLineItemOptionsRequestDataRelationshipsSkuOption {
	return v.value
}

func (v *NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption) Set(val *POSTLineItemOptionsRequestDataRelationshipsSkuOption) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption(val *POSTLineItemOptionsRequestDataRelationshipsSkuOption) *NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption {
	return &NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption{value: val, isSet: true}
}

func (v NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItemOptionsRequestDataRelationshipsSkuOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}