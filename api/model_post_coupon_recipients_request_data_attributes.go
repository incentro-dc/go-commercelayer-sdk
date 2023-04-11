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

// checks if the POSTCouponRecipientsRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponRecipientsRequestDataAttributes{}

// POSTCouponRecipientsRequestDataAttributes struct for POSTCouponRecipientsRequestDataAttributes
type POSTCouponRecipientsRequestDataAttributes struct {
	// The recipient email address
	Email interface{} `json:"email"`
	// The recipient first name
	FirstName interface{} `json:"first_name,omitempty"`
	// The recipient last name
	LastName interface{} `json:"last_name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTCouponRecipientsRequestDataAttributes instantiates a new POSTCouponRecipientsRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponRecipientsRequestDataAttributes(email interface{}) *POSTCouponRecipientsRequestDataAttributes {
	this := POSTCouponRecipientsRequestDataAttributes{}
	this.Email = email
	return &this
}

// NewPOSTCouponRecipientsRequestDataAttributesWithDefaults instantiates a new POSTCouponRecipientsRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponRecipientsRequestDataAttributesWithDefaults() *POSTCouponRecipientsRequestDataAttributes {
	this := POSTCouponRecipientsRequestDataAttributes{}
	return &this
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *POSTCouponRecipientsRequestDataAttributes) SetEmail(v interface{}) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponRecipientsRequestDataAttributes) GetFirstName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetFirstNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return &o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *POSTCouponRecipientsRequestDataAttributes) HasFirstName() bool {
	if o != nil && IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given interface{} and assigns it to the FirstName field.
func (o *POSTCouponRecipientsRequestDataAttributes) SetFirstName(v interface{}) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponRecipientsRequestDataAttributes) GetLastName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetLastNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return &o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *POSTCouponRecipientsRequestDataAttributes) HasLastName() bool {
	if o != nil && IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given interface{} and assigns it to the LastName field.
func (o *POSTCouponRecipientsRequestDataAttributes) SetLastName(v interface{}) {
	o.LastName = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponRecipientsRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTCouponRecipientsRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTCouponRecipientsRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponRecipientsRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTCouponRecipientsRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTCouponRecipientsRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponRecipientsRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponRecipientsRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTCouponRecipientsRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTCouponRecipientsRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTCouponRecipientsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponRecipientsRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
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

type NullablePOSTCouponRecipientsRequestDataAttributes struct {
	value *POSTCouponRecipientsRequestDataAttributes
	isSet bool
}

func (v NullablePOSTCouponRecipientsRequestDataAttributes) Get() *POSTCouponRecipientsRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTCouponRecipientsRequestDataAttributes) Set(val *POSTCouponRecipientsRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponRecipientsRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponRecipientsRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponRecipientsRequestDataAttributes(val *POSTCouponRecipientsRequestDataAttributes) *NullablePOSTCouponRecipientsRequestDataAttributes {
	return &NullablePOSTCouponRecipientsRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTCouponRecipientsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponRecipientsRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}