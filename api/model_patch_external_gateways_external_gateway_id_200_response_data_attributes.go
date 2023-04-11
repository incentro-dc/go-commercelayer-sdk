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

// PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes struct for PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes
type PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The endpoint used by the external gateway to authorize payments.
	AuthorizeUrl interface{} `json:"authorize_url,omitempty"`
	// The endpoint used by the external gateway to capture payments.
	CaptureUrl interface{} `json:"capture_url,omitempty"`
	// The endpoint used by the external gateway to void payments.
	VoidUrl interface{} `json:"void_url,omitempty"`
	// The endpoint used by the external gateway to refund payments.
	RefundUrl interface{} `json:"refund_url,omitempty"`
	// The endpoint used by the external gateway to create a customer payment token.
	TokenUrl interface{} `json:"token_url,omitempty"`
}

// NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes instantiates a new PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes() *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	this := PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributesWithDefaults() *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	this := PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetAuthorizeUrl returns the AuthorizeUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AuthorizeUrl
}

// GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetAuthorizeUrlOk() (*interface{}, bool) {
	if o == nil || o.AuthorizeUrl == nil {
		return nil, false
	}
	return &o.AuthorizeUrl, true
}

// HasAuthorizeUrl returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasAuthorizeUrl() bool {
	if o != nil && o.AuthorizeUrl != nil {
		return true
	}

	return false
}

// SetAuthorizeUrl gets a reference to the given interface{} and assigns it to the AuthorizeUrl field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetAuthorizeUrl(v interface{}) {
	o.AuthorizeUrl = v
}

// GetCaptureUrl returns the CaptureUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CaptureUrl
}

// GetCaptureUrlOk returns a tuple with the CaptureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetCaptureUrlOk() (*interface{}, bool) {
	if o == nil || o.CaptureUrl == nil {
		return nil, false
	}
	return &o.CaptureUrl, true
}

// HasCaptureUrl returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasCaptureUrl() bool {
	if o != nil && o.CaptureUrl != nil {
		return true
	}

	return false
}

// SetCaptureUrl gets a reference to the given interface{} and assigns it to the CaptureUrl field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetCaptureUrl(v interface{}) {
	o.CaptureUrl = v
}

// GetVoidUrl returns the VoidUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.VoidUrl
}

// GetVoidUrlOk returns a tuple with the VoidUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetVoidUrlOk() (*interface{}, bool) {
	if o == nil || o.VoidUrl == nil {
		return nil, false
	}
	return &o.VoidUrl, true
}

// HasVoidUrl returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasVoidUrl() bool {
	if o != nil && o.VoidUrl != nil {
		return true
	}

	return false
}

// SetVoidUrl gets a reference to the given interface{} and assigns it to the VoidUrl field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetVoidUrl(v interface{}) {
	o.VoidUrl = v
}

// GetRefundUrl returns the RefundUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RefundUrl
}

// GetRefundUrlOk returns a tuple with the RefundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetRefundUrlOk() (*interface{}, bool) {
	if o == nil || o.RefundUrl == nil {
		return nil, false
	}
	return &o.RefundUrl, true
}

// HasRefundUrl returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasRefundUrl() bool {
	if o != nil && o.RefundUrl != nil {
		return true
	}

	return false
}

// SetRefundUrl gets a reference to the given interface{} and assigns it to the RefundUrl field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetRefundUrl(v interface{}) {
	o.RefundUrl = v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) GetTokenUrlOk() (*interface{}, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return &o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given interface{} and assigns it to the TokenUrl field.
func (o *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) SetTokenUrl(v interface{}) {
	o.TokenUrl = v
}

func (o PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.AuthorizeUrl != nil {
		toSerialize["authorize_url"] = o.AuthorizeUrl
	}
	if o.CaptureUrl != nil {
		toSerialize["capture_url"] = o.CaptureUrl
	}
	if o.VoidUrl != nil {
		toSerialize["void_url"] = o.VoidUrl
	}
	if o.RefundUrl != nil {
		toSerialize["refund_url"] = o.RefundUrl
	}
	if o.TokenUrl != nil {
		toSerialize["token_url"] = o.TokenUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes struct {
	value *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) Get() *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) Set(val *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes(val *PATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) *NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes {
	return &NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
