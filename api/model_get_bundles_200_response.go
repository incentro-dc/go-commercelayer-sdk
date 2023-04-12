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

// checks if the GETBundles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETBundles200Response{}

// GETBundles200Response struct for GETBundles200Response
type GETBundles200Response struct {
	Data interface{} `json:"data,omitempty"`
}

// NewGETBundles200Response instantiates a new GETBundles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBundles200Response() *GETBundles200Response {
	this := GETBundles200Response{}
	return &this
}

// NewGETBundles200ResponseWithDefaults instantiates a new GETBundles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBundles200ResponseWithDefaults() *GETBundles200Response {
	this := GETBundles200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBundles200Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBundles200Response) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBundles200Response) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *GETBundles200Response) SetData(v interface{}) {
	o.Data = v
}

func (o GETBundles200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETBundles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETBundles200Response struct {
	value *GETBundles200Response
	isSet bool
}

func (v NullableGETBundles200Response) Get() *GETBundles200Response {
	return v.value
}

func (v *NullableGETBundles200Response) Set(val *GETBundles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBundles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBundles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBundles200Response(val *GETBundles200Response) *NullableGETBundles200Response {
	return &NullableGETBundles200Response{value: val, isSet: true}
}

func (v NullableGETBundles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBundles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
