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

// POSTBraintreeGateways201ResponseDataAttributes struct for POSTBraintreeGateways201ResponseDataAttributes
type POSTBraintreeGateways201ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway merchant account ID.
	MerchantAccountId string `json:"merchant_account_id"`
	// The gateway merchant ID.
	MerchantId string `json:"merchant_id"`
	// The gateway API public key.
	PublicKey string `json:"public_key"`
	// The gateway API private key.
	PrivateKey string `json:"private_key"`
	// The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars.
	DescriptorName *string `json:"descriptor_name,omitempty"`
	// The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods.
	DescriptorPhone *string `json:"descriptor_phone,omitempty"`
	// The dynamic descriptor URL.
	DescriptorUrl *string `json:"descriptor_url,omitempty"`
}

// NewPOSTBraintreeGateways201ResponseDataAttributes instantiates a new POSTBraintreeGateways201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTBraintreeGateways201ResponseDataAttributes(name string, merchantAccountId string, merchantId string, publicKey string, privateKey string) *POSTBraintreeGateways201ResponseDataAttributes {
	this := POSTBraintreeGateways201ResponseDataAttributes{}
	this.Name = name
	this.MerchantAccountId = merchantAccountId
	this.MerchantId = merchantId
	this.PublicKey = publicKey
	this.PrivateKey = privateKey
	return &this
}

// NewPOSTBraintreeGateways201ResponseDataAttributesWithDefaults instantiates a new POSTBraintreeGateways201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTBraintreeGateways201ResponseDataAttributesWithDefaults() *POSTBraintreeGateways201ResponseDataAttributes {
	this := POSTBraintreeGateways201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMerchantAccountId returns the MerchantAccountId field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAccountId, true
}

// SetMerchantAccountId sets field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetMerchantAccountId(v string) {
	o.MerchantAccountId = v
}

// GetMerchantId returns the MerchantId field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetPublicKey returns the PublicKey field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetDescriptorName returns the DescriptorName field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorName() string {
	if o == nil || o.DescriptorName == nil {
		var ret string
		return ret
	}
	return *o.DescriptorName
}

// GetDescriptorNameOk returns a tuple with the DescriptorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorNameOk() (*string, bool) {
	if o == nil || o.DescriptorName == nil {
		return nil, false
	}
	return o.DescriptorName, true
}

// HasDescriptorName returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) HasDescriptorName() bool {
	if o != nil && o.DescriptorName != nil {
		return true
	}

	return false
}

// SetDescriptorName gets a reference to the given string and assigns it to the DescriptorName field.
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetDescriptorName(v string) {
	o.DescriptorName = &v
}

// GetDescriptorPhone returns the DescriptorPhone field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorPhone() string {
	if o == nil || o.DescriptorPhone == nil {
		var ret string
		return ret
	}
	return *o.DescriptorPhone
}

// GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorPhoneOk() (*string, bool) {
	if o == nil || o.DescriptorPhone == nil {
		return nil, false
	}
	return o.DescriptorPhone, true
}

// HasDescriptorPhone returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) HasDescriptorPhone() bool {
	if o != nil && o.DescriptorPhone != nil {
		return true
	}

	return false
}

// SetDescriptorPhone gets a reference to the given string and assigns it to the DescriptorPhone field.
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetDescriptorPhone(v string) {
	o.DescriptorPhone = &v
}

// GetDescriptorUrl returns the DescriptorUrl field value if set, zero value otherwise.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorUrl() string {
	if o == nil || o.DescriptorUrl == nil {
		var ret string
		return ret
	}
	return *o.DescriptorUrl
}

// GetDescriptorUrlOk returns a tuple with the DescriptorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) GetDescriptorUrlOk() (*string, bool) {
	if o == nil || o.DescriptorUrl == nil {
		return nil, false
	}
	return o.DescriptorUrl, true
}

// HasDescriptorUrl returns a boolean if a field has been set.
func (o *POSTBraintreeGateways201ResponseDataAttributes) HasDescriptorUrl() bool {
	if o != nil && o.DescriptorUrl != nil {
		return true
	}

	return false
}

// SetDescriptorUrl gets a reference to the given string and assigns it to the DescriptorUrl field.
func (o *POSTBraintreeGateways201ResponseDataAttributes) SetDescriptorUrl(v string) {
	o.DescriptorUrl = &v
}

func (o POSTBraintreeGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if true {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
	}
	if true {
		toSerialize["private_key"] = o.PrivateKey
	}
	if o.DescriptorName != nil {
		toSerialize["descriptor_name"] = o.DescriptorName
	}
	if o.DescriptorPhone != nil {
		toSerialize["descriptor_phone"] = o.DescriptorPhone
	}
	if o.DescriptorUrl != nil {
		toSerialize["descriptor_url"] = o.DescriptorUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTBraintreeGateways201ResponseDataAttributes struct {
	value *POSTBraintreeGateways201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTBraintreeGateways201ResponseDataAttributes) Get() *POSTBraintreeGateways201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTBraintreeGateways201ResponseDataAttributes) Set(val *POSTBraintreeGateways201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTBraintreeGateways201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTBraintreeGateways201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTBraintreeGateways201ResponseDataAttributes(val *POSTBraintreeGateways201ResponseDataAttributes) *NullablePOSTBraintreeGateways201ResponseDataAttributes {
	return &NullablePOSTBraintreeGateways201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTBraintreeGateways201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTBraintreeGateways201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
