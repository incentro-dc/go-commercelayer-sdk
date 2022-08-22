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

// GETStripeGatewaysStripeGatewayId200Response struct for GETStripeGatewaysStripeGatewayId200Response
type GETStripeGatewaysStripeGatewayId200Response struct {
	Data *GETStripeGateways200ResponseDataInner `json:"data,omitempty"`
}

// NewGETStripeGatewaysStripeGatewayId200Response instantiates a new GETStripeGatewaysStripeGatewayId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStripeGatewaysStripeGatewayId200Response() *GETStripeGatewaysStripeGatewayId200Response {
	this := GETStripeGatewaysStripeGatewayId200Response{}
	return &this
}

// NewGETStripeGatewaysStripeGatewayId200ResponseWithDefaults instantiates a new GETStripeGatewaysStripeGatewayId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStripeGatewaysStripeGatewayId200ResponseWithDefaults() *GETStripeGatewaysStripeGatewayId200Response {
	this := GETStripeGatewaysStripeGatewayId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStripeGatewaysStripeGatewayId200Response) GetData() GETStripeGateways200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETStripeGateways200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGatewaysStripeGatewayId200Response) GetDataOk() (*GETStripeGateways200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStripeGatewaysStripeGatewayId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETStripeGateways200ResponseDataInner and assigns it to the Data field.
func (o *GETStripeGatewaysStripeGatewayId200Response) SetData(v GETStripeGateways200ResponseDataInner) {
	o.Data = &v
}

func (o GETStripeGatewaysStripeGatewayId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETStripeGatewaysStripeGatewayId200Response struct {
	value *GETStripeGatewaysStripeGatewayId200Response
	isSet bool
}

func (v NullableGETStripeGatewaysStripeGatewayId200Response) Get() *GETStripeGatewaysStripeGatewayId200Response {
	return v.value
}

func (v *NullableGETStripeGatewaysStripeGatewayId200Response) Set(val *GETStripeGatewaysStripeGatewayId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStripeGatewaysStripeGatewayId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStripeGatewaysStripeGatewayId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStripeGatewaysStripeGatewayId200Response(val *GETStripeGatewaysStripeGatewayId200Response) *NullableGETStripeGatewaysStripeGatewayId200Response {
	return &NullableGETStripeGatewaysStripeGatewayId200Response{value: val, isSet: true}
}

func (v NullableGETStripeGatewaysStripeGatewayId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStripeGatewaysStripeGatewayId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}