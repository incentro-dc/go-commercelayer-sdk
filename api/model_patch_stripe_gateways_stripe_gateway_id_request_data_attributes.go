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

// checks if the PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes{}

// PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes struct for PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes
type PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// Indicates if the gateway will accept payment methods enabled in the Stripe dashboard.
	AutoPayments interface{} `json:"auto_payments,omitempty"`
}

// NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributes instantiates a new PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributes() *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes {
	this := PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes{}
	return &this
}

// NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributesWithDefaults instantiates a new PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStripeGatewaysStripeGatewayIdRequestDataAttributesWithDefaults() *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes {
	this := PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetAutoPayments returns the AutoPayments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetAutoPayments() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoPayments
}

// GetAutoPaymentsOk returns a tuple with the AutoPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) GetAutoPaymentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoPayments) {
		return nil, false
	}
	return &o.AutoPayments, true
}

// HasAutoPayments returns a boolean if a field has been set.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) HasAutoPayments() bool {
	if o != nil && IsNil(o.AutoPayments) {
		return true
	}

	return false
}

// SetAutoPayments gets a reference to the given interface{} and assigns it to the AutoPayments field.
func (o *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) SetAutoPayments(v interface{}) {
	o.AutoPayments = v
}

func (o PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.AutoPayments != nil {
		toSerialize["auto_payments"] = o.AutoPayments
	}
	return toSerialize, nil
}

type NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes struct {
	value *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) Get() *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) Set(val *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes(val *PATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) *NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes {
	return &NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStripeGatewaysStripeGatewayIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}