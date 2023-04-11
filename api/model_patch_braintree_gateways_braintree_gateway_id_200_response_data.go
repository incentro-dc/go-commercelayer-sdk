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

// PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData struct for PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData
type PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                        `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                             `json:"links,omitempty"`
	Attributes    *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *GETBraintreeGateways200ResponseDataInnerRelationships             `json:"relationships,omitempty"`
}

// NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseData instantiates a new PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseData() *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData {
	this := PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData{}
	return &this
}

// NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataWithDefaults instantiates a new PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataWithDefaults() *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData {
	this := PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetAttributes() PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetAttributesOk() (*PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) SetAttributes(v PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetRelationships() GETBraintreeGateways200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETBraintreeGateways200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) GetRelationshipsOk() (*GETBraintreeGateways200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETBraintreeGateways200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) SetRelationships(v GETBraintreeGateways200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData struct {
	value *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData
	isSet bool
}

func (v NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) Get() *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData {
	return v.value
}

func (v *NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) Set(val *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData(val *PATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) *NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData {
	return &NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBraintreeGatewaysBraintreeGatewayId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
