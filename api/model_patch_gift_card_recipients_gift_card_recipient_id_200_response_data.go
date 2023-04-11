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

// PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData struct for PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData
type PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                                      `json:"type,omitempty"`
	Links         *GETAddresses200ResponseDataInnerLinks                           `json:"links,omitempty"`
	Attributes    *PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *GETCouponRecipients200ResponseDataInnerRelationships            `json:"relationships,omitempty"`
}

// NewPATCHGiftCardRecipientsGiftCardRecipientId200ResponseData instantiates a new PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHGiftCardRecipientsGiftCardRecipientId200ResponseData() *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData {
	this := PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData{}
	return &this
}

// NewPATCHGiftCardRecipientsGiftCardRecipientId200ResponseDataWithDefaults instantiates a new PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHGiftCardRecipientsGiftCardRecipientId200ResponseDataWithDefaults() *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData {
	this := PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerLinks and assigns it to the Links field.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetAttributes() PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetAttributesOk() (*PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) SetAttributes(v PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetRelationships() GETCouponRecipients200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) GetRelationshipsOk() (*GETCouponRecipients200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) SetRelationships(v GETCouponRecipients200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) MarshalJSON() ([]byte, error) {
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

type NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData struct {
	value *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData
	isSet bool
}

func (v NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) Get() *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData {
	return v.value
}

func (v *NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) Set(val *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData(val *PATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) *NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData {
	return &NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHGiftCardRecipientsGiftCardRecipientId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
