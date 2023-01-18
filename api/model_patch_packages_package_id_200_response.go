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

// PATCHPackagesPackageId200Response struct for PATCHPackagesPackageId200Response
type PATCHPackagesPackageId200Response struct {
	Data *PATCHPackagesPackageId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPackagesPackageId200Response instantiates a new PATCHPackagesPackageId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPackagesPackageId200Response() *PATCHPackagesPackageId200Response {
	this := PATCHPackagesPackageId200Response{}
	return &this
}

// NewPATCHPackagesPackageId200ResponseWithDefaults instantiates a new PATCHPackagesPackageId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPackagesPackageId200ResponseWithDefaults() *PATCHPackagesPackageId200Response {
	this := PATCHPackagesPackageId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPackagesPackageId200Response) GetData() PATCHPackagesPackageId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHPackagesPackageId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPackagesPackageId200Response) GetDataOk() (*PATCHPackagesPackageId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPackagesPackageId200ResponseData and assigns it to the Data field.
func (o *PATCHPackagesPackageId200Response) SetData(v PATCHPackagesPackageId200ResponseData) {
	o.Data = &v
}

func (o PATCHPackagesPackageId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPackagesPackageId200Response struct {
	value *PATCHPackagesPackageId200Response
	isSet bool
}

func (v NullablePATCHPackagesPackageId200Response) Get() *PATCHPackagesPackageId200Response {
	return v.value
}

func (v *NullablePATCHPackagesPackageId200Response) Set(val *PATCHPackagesPackageId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPackagesPackageId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPackagesPackageId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPackagesPackageId200Response(val *PATCHPackagesPackageId200Response) *NullablePATCHPackagesPackageId200Response {
	return &NullablePATCHPackagesPackageId200Response{value: val, isSet: true}
}

func (v NullablePATCHPackagesPackageId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPackagesPackageId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
