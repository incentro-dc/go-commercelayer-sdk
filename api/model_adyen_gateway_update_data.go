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

// AdyenGatewayUpdateData struct for AdyenGatewayUpdateData
type AdyenGatewayUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                    `json:"id"`
	Attributes    PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenGatewayCreateDataRelationships                      `json:"relationships,omitempty"`
}

// NewAdyenGatewayUpdateData instantiates a new AdyenGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayUpdateData(type_ string, id string, attributes PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) *AdyenGatewayUpdateData {
	this := AdyenGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAdyenGatewayUpdateDataWithDefaults instantiates a new AdyenGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayUpdateDataWithDefaults() *AdyenGatewayUpdateData {
	this := AdyenGatewayUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *AdyenGatewayUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenGatewayUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AdyenGatewayUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdyenGatewayUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenGatewayUpdateData) GetAttributes() PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayUpdateData) GetAttributesOk() (*PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenGatewayUpdateData) SetAttributes(v PATCHAdyenGatewaysAdyenGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenGatewayUpdateData) GetRelationships() AdyenGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayUpdateData) GetRelationshipsOk() (*AdyenGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenGatewayUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *AdyenGatewayUpdateData) SetRelationships(v AdyenGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o AdyenGatewayUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableAdyenGatewayUpdateData struct {
	value *AdyenGatewayUpdateData
	isSet bool
}

func (v NullableAdyenGatewayUpdateData) Get() *AdyenGatewayUpdateData {
	return v.value
}

func (v *NullableAdyenGatewayUpdateData) Set(val *AdyenGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayUpdateData(val *AdyenGatewayUpdateData) *NullableAdyenGatewayUpdateData {
	return &NullableAdyenGatewayUpdateData{value: val, isSet: true}
}

func (v NullableAdyenGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
