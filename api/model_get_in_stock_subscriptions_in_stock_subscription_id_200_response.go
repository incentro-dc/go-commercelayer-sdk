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

// GETInStockSubscriptionsInStockSubscriptionId200Response struct for GETInStockSubscriptionsInStockSubscriptionId200Response
type GETInStockSubscriptionsInStockSubscriptionId200Response struct {
	Data *GETInStockSubscriptions200ResponseDataInner `json:"data,omitempty"`
}

// NewGETInStockSubscriptionsInStockSubscriptionId200Response instantiates a new GETInStockSubscriptionsInStockSubscriptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInStockSubscriptionsInStockSubscriptionId200Response() *GETInStockSubscriptionsInStockSubscriptionId200Response {
	this := GETInStockSubscriptionsInStockSubscriptionId200Response{}
	return &this
}

// NewGETInStockSubscriptionsInStockSubscriptionId200ResponseWithDefaults instantiates a new GETInStockSubscriptionsInStockSubscriptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInStockSubscriptionsInStockSubscriptionId200ResponseWithDefaults() *GETInStockSubscriptionsInStockSubscriptionId200Response {
	this := GETInStockSubscriptionsInStockSubscriptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInStockSubscriptionsInStockSubscriptionId200Response) GetData() GETInStockSubscriptions200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETInStockSubscriptions200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptionsInStockSubscriptionId200Response) GetDataOk() (*GETInStockSubscriptions200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInStockSubscriptionsInStockSubscriptionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInStockSubscriptions200ResponseDataInner and assigns it to the Data field.
func (o *GETInStockSubscriptionsInStockSubscriptionId200Response) SetData(v GETInStockSubscriptions200ResponseDataInner) {
	o.Data = &v
}

func (o GETInStockSubscriptionsInStockSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInStockSubscriptionsInStockSubscriptionId200Response struct {
	value *GETInStockSubscriptionsInStockSubscriptionId200Response
	isSet bool
}

func (v NullableGETInStockSubscriptionsInStockSubscriptionId200Response) Get() *GETInStockSubscriptionsInStockSubscriptionId200Response {
	return v.value
}

func (v *NullableGETInStockSubscriptionsInStockSubscriptionId200Response) Set(val *GETInStockSubscriptionsInStockSubscriptionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInStockSubscriptionsInStockSubscriptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInStockSubscriptionsInStockSubscriptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInStockSubscriptionsInStockSubscriptionId200Response(val *GETInStockSubscriptionsInStockSubscriptionId200Response) *NullableGETInStockSubscriptionsInStockSubscriptionId200Response {
	return &NullableGETInStockSubscriptionsInStockSubscriptionId200Response{value: val, isSet: true}
}

func (v NullableGETInStockSubscriptionsInStockSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInStockSubscriptionsInStockSubscriptionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
