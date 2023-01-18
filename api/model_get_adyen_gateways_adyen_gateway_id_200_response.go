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

// GETAdyenGatewaysAdyenGatewayId200Response struct for GETAdyenGatewaysAdyenGatewayId200Response
type GETAdyenGatewaysAdyenGatewayId200Response struct {
	Data *GETAdyenGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETAdyenGatewaysAdyenGatewayId200Response instantiates a new GETAdyenGatewaysAdyenGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenGatewaysAdyenGatewayId200Response() *GETAdyenGatewaysAdyenGatewayId200Response {
	this := GETAdyenGatewaysAdyenGatewayId200Response{}
	return &this
}

// NewGETAdyenGatewaysAdyenGatewayId200ResponseWithDefaults instantiates a new GETAdyenGatewaysAdyenGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenGatewaysAdyenGatewayId200ResponseWithDefaults() *GETAdyenGatewaysAdyenGatewayId200Response {
	this := GETAdyenGatewaysAdyenGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdyenGatewaysAdyenGatewayId200Response) GetData() GETAdyenGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETAdyenGateways200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenGatewaysAdyenGatewayId200Response) GetDataOk() (*GETAdyenGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdyenGatewaysAdyenGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAdyenGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETAdyenGatewaysAdyenGatewayId200Response) SetData(v GETAdyenGateways200ResponseDataInner) {
	o.Data = &v
}

func (o GETAdyenGatewaysAdyenGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenGatewaysAdyenGatewayId200Response struct {
	value *GETAdyenGatewaysAdyenGatewayId200Response
	isSet bool
}

func (v NullableGETAdyenGatewaysAdyenGatewayId200Response) Get() *GETAdyenGatewaysAdyenGatewayId200Response {
	return v.value
}

func (v *NullableGETAdyenGatewaysAdyenGatewayId200Response) Set(val *GETAdyenGatewaysAdyenGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenGatewaysAdyenGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenGatewaysAdyenGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenGatewaysAdyenGatewayId200Response(val *GETAdyenGatewaysAdyenGatewayId200Response) *NullableGETAdyenGatewaysAdyenGatewayId200Response {
	return &NullableGETAdyenGatewaysAdyenGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETAdyenGatewaysAdyenGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenGatewaysAdyenGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
