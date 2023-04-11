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

// checks if the PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes{}

// PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes struct for PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes
type PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The gateway client ID.
	ClientId interface{} `json:"client_id,omitempty"`
	// The gateway client secret.
	ClientSecret interface{} `json:"client_secret,omitempty"`
}

// NewPATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes instantiates a new PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes() *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes {
	this := PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes{}
	return &this
}

// NewPATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributesWithDefaults instantiates a new PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributesWithDefaults() *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes {
	this := PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetClientId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetClientIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return &o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) HasClientId() bool {
	if o != nil && IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given interface{} and assigns it to the ClientId field.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) SetClientId(v interface{}) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetClientSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) GetClientSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return &o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) HasClientSecret() bool {
	if o != nil && IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given interface{} and assigns it to the ClientSecret field.
func (o *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) SetClientSecret(v interface{}) {
	o.ClientSecret = v
}

func (o PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	return toSerialize, nil
}

type NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes struct {
	value *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) Get() *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) Set(val *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes(val *PATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) *NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes {
	return &NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaypalGatewaysPaypalGatewayIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}