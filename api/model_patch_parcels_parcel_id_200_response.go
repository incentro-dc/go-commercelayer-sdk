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

// PATCHParcelsParcelId200Response struct for PATCHParcelsParcelId200Response
type PATCHParcelsParcelId200Response struct {
	Data *PATCHParcelsParcelId200ResponseData `json:"data,omitempty"`
}

// NewPATCHParcelsParcelId200Response instantiates a new PATCHParcelsParcelId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHParcelsParcelId200Response() *PATCHParcelsParcelId200Response {
	this := PATCHParcelsParcelId200Response{}
	return &this
}

// NewPATCHParcelsParcelId200ResponseWithDefaults instantiates a new PATCHParcelsParcelId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHParcelsParcelId200ResponseWithDefaults() *PATCHParcelsParcelId200Response {
	this := PATCHParcelsParcelId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHParcelsParcelId200Response) GetData() PATCHParcelsParcelId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHParcelsParcelId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelsParcelId200Response) GetDataOk() (*PATCHParcelsParcelId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHParcelsParcelId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHParcelsParcelId200ResponseData and assigns it to the Data field.
func (o *PATCHParcelsParcelId200Response) SetData(v PATCHParcelsParcelId200ResponseData) {
	o.Data = &v
}

func (o PATCHParcelsParcelId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHParcelsParcelId200Response struct {
	value *PATCHParcelsParcelId200Response
	isSet bool
}

func (v NullablePATCHParcelsParcelId200Response) Get() *PATCHParcelsParcelId200Response {
	return v.value
}

func (v *NullablePATCHParcelsParcelId200Response) Set(val *PATCHParcelsParcelId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHParcelsParcelId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHParcelsParcelId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHParcelsParcelId200Response(val *PATCHParcelsParcelId200Response) *NullablePATCHParcelsParcelId200Response {
	return &NullablePATCHParcelsParcelId200Response{value: val, isSet: true}
}

func (v NullablePATCHParcelsParcelId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHParcelsParcelId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
