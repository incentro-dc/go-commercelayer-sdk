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

// checks if the PATCHAxerveGatewaysAxerveGatewayId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAxerveGatewaysAxerveGatewayId200Response{}

// PATCHAxerveGatewaysAxerveGatewayId200Response struct for PATCHAxerveGatewaysAxerveGatewayId200Response
type PATCHAxerveGatewaysAxerveGatewayId200Response struct {
	Data *PATCHAxerveGatewaysAxerveGatewayId200ResponseData `json:"data,omitempty"`
}

// NewPATCHAxerveGatewaysAxerveGatewayId200Response instantiates a new PATCHAxerveGatewaysAxerveGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAxerveGatewaysAxerveGatewayId200Response() *PATCHAxerveGatewaysAxerveGatewayId200Response {
	this := PATCHAxerveGatewaysAxerveGatewayId200Response{}
	return &this
}

// NewPATCHAxerveGatewaysAxerveGatewayId200ResponseWithDefaults instantiates a new PATCHAxerveGatewaysAxerveGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAxerveGatewaysAxerveGatewayId200ResponseWithDefaults() *PATCHAxerveGatewaysAxerveGatewayId200Response {
	this := PATCHAxerveGatewaysAxerveGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHAxerveGatewaysAxerveGatewayId200Response) GetData() PATCHAxerveGatewaysAxerveGatewayId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHAxerveGatewaysAxerveGatewayId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200Response) GetDataOk() (*PATCHAxerveGatewaysAxerveGatewayId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHAxerveGatewaysAxerveGatewayId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHAxerveGatewaysAxerveGatewayId200ResponseData and assigns it to the Data field.
func (o *PATCHAxerveGatewaysAxerveGatewayId200Response) SetData(v PATCHAxerveGatewaysAxerveGatewayId200ResponseData) {
	o.Data = &v
}

func (o PATCHAxerveGatewaysAxerveGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAxerveGatewaysAxerveGatewayId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHAxerveGatewaysAxerveGatewayId200Response struct {
	value *PATCHAxerveGatewaysAxerveGatewayId200Response
	isSet bool
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayId200Response) Get() *PATCHAxerveGatewaysAxerveGatewayId200Response {
	return v.value
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayId200Response) Set(val *PATCHAxerveGatewaysAxerveGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAxerveGatewaysAxerveGatewayId200Response(val *PATCHAxerveGatewaysAxerveGatewayId200Response) *NullablePATCHAxerveGatewaysAxerveGatewayId200Response {
	return &NullablePATCHAxerveGatewaysAxerveGatewayId200Response{value: val, isSet: true}
}

func (v NullablePATCHAxerveGatewaysAxerveGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAxerveGatewaysAxerveGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
