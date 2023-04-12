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

// checks if the PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{}

// PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes struct for PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes
type PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes struct {
	// The url to redirect the customer after the payment flow is completed.
	RedirectUrl interface{} `json:"redirect_url,omitempty"`
	// Send this attribute if you want to refresh all the pending transactions, can be used as webhooks fallback logic.
	Refresh interface{} `json:"_refresh,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes instantiates a new PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes() *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	this := PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults() *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	this := PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{}
	return &this
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRedirectUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRedirectUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasRedirectUrl() bool {
	if o != nil && IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given interface{} and assigns it to the RedirectUrl field.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetRedirectUrl(v interface{}) {
	o.RedirectUrl = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRefresh() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRefreshOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Refresh) {
		return nil, false
	}
	return &o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasRefresh() bool {
	if o != nil && IsNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given interface{} and assigns it to the Refresh field.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetRefresh(v interface{}) {
	o.Refresh = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if o.Refresh != nil {
		toSerialize["_refresh"] = o.Refresh
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
	return toSerialize, nil
}

type NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes struct {
	value *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) Get() *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) Set(val *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes(val *PATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) *NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes {
	return &NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}