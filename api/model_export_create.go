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

// checks if the ExportCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportCreate{}

// ExportCreate struct for ExportCreate
type ExportCreate struct {
	Data ExportCreateData `json:"data"`
}

// NewExportCreate instantiates a new ExportCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportCreate(data ExportCreateData) *ExportCreate {
	this := ExportCreate{}
	this.Data = data
	return &this
}

// NewExportCreateWithDefaults instantiates a new ExportCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportCreateWithDefaults() *ExportCreate {
	this := ExportCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ExportCreate) GetData() ExportCreateData {
	if o == nil {
		var ret ExportCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExportCreate) GetDataOk() (*ExportCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExportCreate) SetData(v ExportCreateData) {
	o.Data = v
}

func (o ExportCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableExportCreate struct {
	value *ExportCreate
	isSet bool
}

func (v NullableExportCreate) Get() *ExportCreate {
	return v.value
}

func (v *NullableExportCreate) Set(val *ExportCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableExportCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableExportCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportCreate(val *ExportCreate) *NullableExportCreate {
	return &NullableExportCreate{value: val, isSet: true}
}

func (v NullableExportCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
