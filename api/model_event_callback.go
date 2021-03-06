/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EventCallback struct for EventCallback
type EventCallback struct {
	Data EventCallbackData `json:"data"`
}

// NewEventCallback instantiates a new EventCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCallback(data EventCallbackData) *EventCallback {
	this := EventCallback{}
	this.Data = data
	return &this
}

// NewEventCallbackWithDefaults instantiates a new EventCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCallbackWithDefaults() *EventCallback {
	this := EventCallback{}
	return &this
}

// GetData returns the Data field value
func (o *EventCallback) GetData() EventCallbackData {
	if o == nil {
		var ret EventCallbackData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EventCallback) GetDataOk() (*EventCallbackData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EventCallback) SetData(v EventCallbackData) {
	o.Data = v
}

func (o EventCallback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableEventCallback struct {
	value *EventCallback
	isSet bool
}

func (v NullableEventCallback) Get() *EventCallback {
	return v.value
}

func (v *NullableEventCallback) Set(val *EventCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCallback(val *EventCallback) *NullableEventCallback {
	return &NullableEventCallback{value: val, isSet: true}
}

func (v NullableEventCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
