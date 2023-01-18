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

// WebhookUpdateData struct for WebhookUpdateData
type WebhookUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                          `json:"id"`
	Attributes    PATCHWebhooksWebhookId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                          `json:"relationships,omitempty"`
}

// NewWebhookUpdateData instantiates a new WebhookUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUpdateData(type_ string, id string, attributes PATCHWebhooksWebhookId200ResponseDataAttributes) *WebhookUpdateData {
	this := WebhookUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewWebhookUpdateDataWithDefaults instantiates a new WebhookUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUpdateDataWithDefaults() *WebhookUpdateData {
	this := WebhookUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *WebhookUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *WebhookUpdateData) GetAttributes() PATCHWebhooksWebhookId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHWebhooksWebhookId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateData) GetAttributesOk() (*PATCHWebhooksWebhookId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WebhookUpdateData) SetAttributes(v PATCHWebhooksWebhookId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WebhookUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WebhookUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *WebhookUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o WebhookUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookUpdateData struct {
	value *WebhookUpdateData
	isSet bool
}

func (v NullableWebhookUpdateData) Get() *WebhookUpdateData {
	return v.value
}

func (v *NullableWebhookUpdateData) Set(val *WebhookUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUpdateData(val *WebhookUpdateData) *NullableWebhookUpdateData {
	return &NullableWebhookUpdateData{value: val, isSet: true}
}

func (v NullableWebhookUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
