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

// checks if the POSTCheckoutComGatewaysRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCheckoutComGatewaysRequestDataAttributes{}

// POSTCheckoutComGatewaysRequestDataAttributes struct for POSTCheckoutComGatewaysRequestDataAttributes
type POSTCheckoutComGatewaysRequestDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The gateway secret key.
	SecretKey interface{} `json:"secret_key"`
	// The gateway public key.
	PublicKey interface{} `json:"public_key"`
}

// NewPOSTCheckoutComGatewaysRequestDataAttributes instantiates a new POSTCheckoutComGatewaysRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComGatewaysRequestDataAttributes(name interface{}, secretKey interface{}, publicKey interface{}) *POSTCheckoutComGatewaysRequestDataAttributes {
	this := POSTCheckoutComGatewaysRequestDataAttributes{}
	this.Name = name
	this.SecretKey = secretKey
	this.PublicKey = publicKey
	return &this
}

// NewPOSTCheckoutComGatewaysRequestDataAttributesWithDefaults instantiates a new POSTCheckoutComGatewaysRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComGatewaysRequestDataAttributesWithDefaults() *POSTCheckoutComGatewaysRequestDataAttributes {
	this := POSTCheckoutComGatewaysRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTCheckoutComGatewaysRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCheckoutComGatewaysRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTCheckoutComGatewaysRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCheckoutComGatewaysRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTCheckoutComGatewaysRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCheckoutComGatewaysRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTCheckoutComGatewaysRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetSecretKey returns the SecretKey field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetSecretKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetSecretKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *POSTCheckoutComGatewaysRequestDataAttributes) SetSecretKey(v interface{}) {
	o.SecretKey = v
}

// GetPublicKey returns the PublicKey field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetPublicKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCheckoutComGatewaysRequestDataAttributes) GetPublicKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *POSTCheckoutComGatewaysRequestDataAttributes) SetPublicKey(v interface{}) {
	o.PublicKey = v
}

func (o POSTCheckoutComGatewaysRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCheckoutComGatewaysRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullablePOSTCheckoutComGatewaysRequestDataAttributes struct {
	value *POSTCheckoutComGatewaysRequestDataAttributes
	isSet bool
}

func (v NullablePOSTCheckoutComGatewaysRequestDataAttributes) Get() *POSTCheckoutComGatewaysRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTCheckoutComGatewaysRequestDataAttributes) Set(val *POSTCheckoutComGatewaysRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComGatewaysRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComGatewaysRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComGatewaysRequestDataAttributes(val *POSTCheckoutComGatewaysRequestDataAttributes) *NullablePOSTCheckoutComGatewaysRequestDataAttributes {
	return &NullablePOSTCheckoutComGatewaysRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComGatewaysRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComGatewaysRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}