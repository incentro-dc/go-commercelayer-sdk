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

// ExternalGatewayUpdateData struct for ExternalGatewayUpdateData
type ExternalGatewayUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                          `json:"id"`
	Attributes    PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                                          `json:"relationships,omitempty"`
}

// NewExternalGatewayUpdateData instantiates a new ExternalGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGatewayUpdateData(type_ string, id string, attributes PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) *ExternalGatewayUpdateData {
	this := ExternalGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewExternalGatewayUpdateDataWithDefaults instantiates a new ExternalGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGatewayUpdateDataWithDefaults() *ExternalGatewayUpdateData {
	this := ExternalGatewayUpdateData{}
	var type_ string = "external_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ExternalGatewayUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalGatewayUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalGatewayUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ExternalGatewayUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalGatewayUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalGatewayUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalGatewayUpdateData) GetAttributes() PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalGatewayUpdateData) GetAttributesOk() (*PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalGatewayUpdateData) SetAttributes(v PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalGatewayUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalGatewayUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ExternalGatewayUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ExternalGatewayUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableExternalGatewayUpdateData struct {
	value *ExternalGatewayUpdateData
	isSet bool
}

func (v NullableExternalGatewayUpdateData) Get() *ExternalGatewayUpdateData {
	return v.value
}

func (v *NullableExternalGatewayUpdateData) Set(val *ExternalGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGatewayUpdateData(val *ExternalGatewayUpdateData) *NullableExternalGatewayUpdateData {
	return &NullableExternalGatewayUpdateData{value: val, isSet: true}
}

func (v NullableExternalGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
