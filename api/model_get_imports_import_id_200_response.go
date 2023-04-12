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

// checks if the GETImportsImportId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETImportsImportId200Response{}

// GETImportsImportId200Response struct for GETImportsImportId200Response
type GETImportsImportId200Response struct {
	Data *GETImportsImportId200ResponseData `json:"data,omitempty"`
}

// NewGETImportsImportId200Response instantiates a new GETImportsImportId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETImportsImportId200Response() *GETImportsImportId200Response {
	this := GETImportsImportId200Response{}
	return &this
}

// NewGETImportsImportId200ResponseWithDefaults instantiates a new GETImportsImportId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETImportsImportId200ResponseWithDefaults() *GETImportsImportId200Response {
	this := GETImportsImportId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETImportsImportId200Response) GetData() GETImportsImportId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETImportsImportId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETImportsImportId200Response) GetDataOk() (*GETImportsImportId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETImportsImportId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETImportsImportId200ResponseData and assigns it to the Data field.
func (o *GETImportsImportId200Response) SetData(v GETImportsImportId200ResponseData) {
	o.Data = &v
}

func (o GETImportsImportId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETImportsImportId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETImportsImportId200Response struct {
	value *GETImportsImportId200Response
	isSet bool
}

func (v NullableGETImportsImportId200Response) Get() *GETImportsImportId200Response {
	return v.value
}

func (v *NullableGETImportsImportId200Response) Set(val *GETImportsImportId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETImportsImportId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETImportsImportId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETImportsImportId200Response(val *GETImportsImportId200Response) *NullableGETImportsImportId200Response {
	return &NullableGETImportsImportId200Response{value: val, isSet: true}
}

func (v NullableGETImportsImportId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETImportsImportId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
