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

// checks if the PATCHExternalGatewaysExternalGatewayIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHExternalGatewaysExternalGatewayIdRequest{}

// PATCHExternalGatewaysExternalGatewayIdRequest struct for PATCHExternalGatewaysExternalGatewayIdRequest
type PATCHExternalGatewaysExternalGatewayIdRequest struct {
	Data PATCHExternalGatewaysExternalGatewayIdRequestData `json:"data"`
}

// NewPATCHExternalGatewaysExternalGatewayIdRequest instantiates a new PATCHExternalGatewaysExternalGatewayIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalGatewaysExternalGatewayIdRequest(data PATCHExternalGatewaysExternalGatewayIdRequestData) *PATCHExternalGatewaysExternalGatewayIdRequest {
	this := PATCHExternalGatewaysExternalGatewayIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHExternalGatewaysExternalGatewayIdRequestWithDefaults instantiates a new PATCHExternalGatewaysExternalGatewayIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalGatewaysExternalGatewayIdRequestWithDefaults() *PATCHExternalGatewaysExternalGatewayIdRequest {
	this := PATCHExternalGatewaysExternalGatewayIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHExternalGatewaysExternalGatewayIdRequest) GetData() PATCHExternalGatewaysExternalGatewayIdRequestData {
	if o == nil {
		var ret PATCHExternalGatewaysExternalGatewayIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHExternalGatewaysExternalGatewayIdRequest) GetDataOk() (*PATCHExternalGatewaysExternalGatewayIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHExternalGatewaysExternalGatewayIdRequest) SetData(v PATCHExternalGatewaysExternalGatewayIdRequestData) {
	o.Data = v
}

func (o PATCHExternalGatewaysExternalGatewayIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHExternalGatewaysExternalGatewayIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHExternalGatewaysExternalGatewayIdRequest struct {
	value *PATCHExternalGatewaysExternalGatewayIdRequest
	isSet bool
}

func (v NullablePATCHExternalGatewaysExternalGatewayIdRequest) Get() *PATCHExternalGatewaysExternalGatewayIdRequest {
	return v.value
}

func (v *NullablePATCHExternalGatewaysExternalGatewayIdRequest) Set(val *PATCHExternalGatewaysExternalGatewayIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalGatewaysExternalGatewayIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalGatewaysExternalGatewayIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalGatewaysExternalGatewayIdRequest(val *PATCHExternalGatewaysExternalGatewayIdRequest) *NullablePATCHExternalGatewaysExternalGatewayIdRequest {
	return &NullablePATCHExternalGatewaysExternalGatewayIdRequest{value: val, isSet: true}
}

func (v NullablePATCHExternalGatewaysExternalGatewayIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalGatewaysExternalGatewayIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}