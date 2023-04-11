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

// Export struct for Export
type Export struct {
	Data *ExportData `json:"data,omitempty"`
}

// NewExport instantiates a new Export object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport() *Export {
	this := Export{}
	return &this
}

// NewExportWithDefaults instantiates a new Export object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportWithDefaults() *Export {
	this := Export{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Export) GetData() ExportData {
	if o == nil || o.Data == nil {
		var ret ExportData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Export) GetDataOk() (*ExportData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Export) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ExportData and assigns it to the Data field.
func (o *Export) SetData(v ExportData) {
	o.Data = &v
}

func (o Export) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExport struct {
	value *Export
	isSet bool
}

func (v NullableExport) Get() *Export {
	return v.value
}

func (v *NullableExport) Set(val *Export) {
	v.value = val
	v.isSet = true
}

func (v NullableExport) IsSet() bool {
	return v.isSet
}

func (v *NullableExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport(val *Export) *NullableExport {
	return &NullableExport{value: val, isSet: true}
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
