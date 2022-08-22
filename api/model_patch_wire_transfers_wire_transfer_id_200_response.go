/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHWireTransfersWireTransferId200Response struct for PATCHWireTransfersWireTransferId200Response
type PATCHWireTransfersWireTransferId200Response struct {
	Data *PATCHWireTransfersWireTransferId200ResponseData `json:"data,omitempty"`
}

// NewPATCHWireTransfersWireTransferId200Response instantiates a new PATCHWireTransfersWireTransferId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHWireTransfersWireTransferId200Response() *PATCHWireTransfersWireTransferId200Response {
	this := PATCHWireTransfersWireTransferId200Response{}
	return &this
}

// NewPATCHWireTransfersWireTransferId200ResponseWithDefaults instantiates a new PATCHWireTransfersWireTransferId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHWireTransfersWireTransferId200ResponseWithDefaults() *PATCHWireTransfersWireTransferId200Response {
	this := PATCHWireTransfersWireTransferId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHWireTransfersWireTransferId200Response) GetData() PATCHWireTransfersWireTransferId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHWireTransfersWireTransferId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWireTransfersWireTransferId200Response) GetDataOk() (*PATCHWireTransfersWireTransferId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHWireTransfersWireTransferId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHWireTransfersWireTransferId200ResponseData and assigns it to the Data field.
func (o *PATCHWireTransfersWireTransferId200Response) SetData(v PATCHWireTransfersWireTransferId200ResponseData) {
	o.Data = &v
}

func (o PATCHWireTransfersWireTransferId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHWireTransfersWireTransferId200Response struct {
	value *PATCHWireTransfersWireTransferId200Response
	isSet bool
}

func (v NullablePATCHWireTransfersWireTransferId200Response) Get() *PATCHWireTransfersWireTransferId200Response {
	return v.value
}

func (v *NullablePATCHWireTransfersWireTransferId200Response) Set(val *PATCHWireTransfersWireTransferId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHWireTransfersWireTransferId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHWireTransfersWireTransferId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHWireTransfersWireTransferId200Response(val *PATCHWireTransfersWireTransferId200Response) *NullablePATCHWireTransfersWireTransferId200Response {
	return &NullablePATCHWireTransfersWireTransferId200Response{value: val, isSet: true}
}

func (v NullablePATCHWireTransfersWireTransferId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHWireTransfersWireTransferId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
