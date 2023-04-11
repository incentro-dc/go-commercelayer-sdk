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

// PATCHCouponsCouponId200ResponseData struct for PATCHCouponsCouponId200ResponseData
type PATCHCouponsCouponId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                    `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks         `json:"links,omitempty"`
	Attributes    *PATCHCouponsCouponId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *GETCoupons200ResponseDataInnerRelationships   `json:"relationships,omitempty"`
}

// NewPATCHCouponsCouponId200ResponseData instantiates a new PATCHCouponsCouponId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCouponsCouponId200ResponseData() *PATCHCouponsCouponId200ResponseData {
	this := PATCHCouponsCouponId200ResponseData{}
	return &this
}

// NewPATCHCouponsCouponId200ResponseDataWithDefaults instantiates a new PATCHCouponsCouponId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCouponsCouponId200ResponseDataWithDefaults() *PATCHCouponsCouponId200ResponseData {
	this := PATCHCouponsCouponId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponsCouponId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponsCouponId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHCouponsCouponId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponsCouponId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponsCouponId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHCouponsCouponId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *PATCHCouponsCouponId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseData) GetAttributes() PATCHCouponsCouponId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PATCHCouponsCouponId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseData) GetAttributesOk() (*PATCHCouponsCouponId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHCouponsCouponId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHCouponsCouponId200ResponseData) SetAttributes(v PATCHCouponsCouponId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseData) GetRelationships() GETCoupons200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETCoupons200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseData) GetRelationshipsOk() (*GETCoupons200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETCoupons200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *PATCHCouponsCouponId200ResponseData) SetRelationships(v GETCoupons200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o PATCHCouponsCouponId200ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePATCHCouponsCouponId200ResponseData struct {
	value *PATCHCouponsCouponId200ResponseData
	isSet bool
}

func (v NullablePATCHCouponsCouponId200ResponseData) Get() *PATCHCouponsCouponId200ResponseData {
	return v.value
}

func (v *NullablePATCHCouponsCouponId200ResponseData) Set(val *PATCHCouponsCouponId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCouponsCouponId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCouponsCouponId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCouponsCouponId200ResponseData(val *PATCHCouponsCouponId200ResponseData) *NullablePATCHCouponsCouponId200ResponseData {
	return &NullablePATCHCouponsCouponId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHCouponsCouponId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCouponsCouponId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
