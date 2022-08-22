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

// GETEvents200ResponseDataInnerRelationships struct for GETEvents200ResponseDataInnerRelationships
type GETEvents200ResponseDataInnerRelationships struct {
	LastEventCallbacks *GETEvents200ResponseDataInnerRelationshipsLastEventCallbacks `json:"last_event_callbacks,omitempty"`
	Webhooks           *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook    `json:"webhooks,omitempty"`
}

// NewGETEvents200ResponseDataInnerRelationships instantiates a new GETEvents200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEvents200ResponseDataInnerRelationships() *GETEvents200ResponseDataInnerRelationships {
	this := GETEvents200ResponseDataInnerRelationships{}
	return &this
}

// NewGETEvents200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETEvents200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEvents200ResponseDataInnerRelationshipsWithDefaults() *GETEvents200ResponseDataInnerRelationships {
	this := GETEvents200ResponseDataInnerRelationships{}
	return &this
}

// GetLastEventCallbacks returns the LastEventCallbacks field value if set, zero value otherwise.
func (o *GETEvents200ResponseDataInnerRelationships) GetLastEventCallbacks() GETEvents200ResponseDataInnerRelationshipsLastEventCallbacks {
	if o == nil || o.LastEventCallbacks == nil {
		var ret GETEvents200ResponseDataInnerRelationshipsLastEventCallbacks
		return ret
	}
	return *o.LastEventCallbacks
}

// GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200ResponseDataInnerRelationships) GetLastEventCallbacksOk() (*GETEvents200ResponseDataInnerRelationshipsLastEventCallbacks, bool) {
	if o == nil || o.LastEventCallbacks == nil {
		return nil, false
	}
	return o.LastEventCallbacks, true
}

// HasLastEventCallbacks returns a boolean if a field has been set.
func (o *GETEvents200ResponseDataInnerRelationships) HasLastEventCallbacks() bool {
	if o != nil && o.LastEventCallbacks != nil {
		return true
	}

	return false
}

// SetLastEventCallbacks gets a reference to the given GETEvents200ResponseDataInnerRelationshipsLastEventCallbacks and assigns it to the LastEventCallbacks field.
func (o *GETEvents200ResponseDataInnerRelationships) SetLastEventCallbacks(v GETEvents200ResponseDataInnerRelationshipsLastEventCallbacks) {
	o.LastEventCallbacks = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *GETEvents200ResponseDataInnerRelationships) GetWebhooks() GETEventCallbacks200ResponseDataInnerRelationshipsWebhook {
	if o == nil || o.Webhooks == nil {
		var ret GETEventCallbacks200ResponseDataInnerRelationshipsWebhook
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEvents200ResponseDataInnerRelationships) GetWebhooksOk() (*GETEventCallbacks200ResponseDataInnerRelationshipsWebhook, bool) {
	if o == nil || o.Webhooks == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *GETEvents200ResponseDataInnerRelationships) HasWebhooks() bool {
	if o != nil && o.Webhooks != nil {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given GETEventCallbacks200ResponseDataInnerRelationshipsWebhook and assigns it to the Webhooks field.
func (o *GETEvents200ResponseDataInnerRelationships) SetWebhooks(v GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) {
	o.Webhooks = &v
}

func (o GETEvents200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastEventCallbacks != nil {
		toSerialize["last_event_callbacks"] = o.LastEventCallbacks
	}
	if o.Webhooks != nil {
		toSerialize["webhooks"] = o.Webhooks
	}
	return json.Marshal(toSerialize)
}

type NullableGETEvents200ResponseDataInnerRelationships struct {
	value *GETEvents200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETEvents200ResponseDataInnerRelationships) Get() *GETEvents200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETEvents200ResponseDataInnerRelationships) Set(val *GETEvents200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEvents200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEvents200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEvents200ResponseDataInnerRelationships(val *GETEvents200ResponseDataInnerRelationships) *NullableGETEvents200ResponseDataInnerRelationships {
	return &NullableGETEvents200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETEvents200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEvents200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
