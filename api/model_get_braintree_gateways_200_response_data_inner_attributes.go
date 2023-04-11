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

// GETBraintreeGateways200ResponseDataInnerAttributes struct for GETBraintreeGateways200ResponseDataInnerAttributes
type GETBraintreeGateways200ResponseDataInnerAttributes struct {
	// The payment gateway's internal name.
	Name interface{} `json:"name,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The dynamic descriptor name. Must be composed by business name (3, 7 or 12 chars), an asterisk (*) and the product name (18, 14 or 9 chars), for a total length of 22 chars.
	DescriptorName interface{} `json:"descriptor_name,omitempty"`
	// The dynamic descriptor phone number. Must be 10-14 characters and can only contain numbers, dashes, parentheses and periods.
	DescriptorPhone interface{} `json:"descriptor_phone,omitempty"`
	// The dynamic descriptor URL.
	DescriptorUrl interface{} `json:"descriptor_url,omitempty"`
	// The gateway webhook URL, generated automatically.
	WebhookEndpointUrl interface{} `json:"webhook_endpoint_url,omitempty"`
}

// NewGETBraintreeGateways200ResponseDataInnerAttributes instantiates a new GETBraintreeGateways200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreeGateways200ResponseDataInnerAttributes() *GETBraintreeGateways200ResponseDataInnerAttributes {
	this := GETBraintreeGateways200ResponseDataInnerAttributes{}
	return &this
}

// NewGETBraintreeGateways200ResponseDataInnerAttributesWithDefaults instantiates a new GETBraintreeGateways200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreeGateways200ResponseDataInnerAttributesWithDefaults() *GETBraintreeGateways200ResponseDataInnerAttributes {
	this := GETBraintreeGateways200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetDescriptorName returns the DescriptorName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DescriptorName
}

// GetDescriptorNameOk returns a tuple with the DescriptorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorNameOk() (*interface{}, bool) {
	if o == nil || o.DescriptorName == nil {
		return nil, false
	}
	return &o.DescriptorName, true
}

// HasDescriptorName returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasDescriptorName() bool {
	if o != nil && o.DescriptorName != nil {
		return true
	}

	return false
}

// SetDescriptorName gets a reference to the given interface{} and assigns it to the DescriptorName field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetDescriptorName(v interface{}) {
	o.DescriptorName = v
}

// GetDescriptorPhone returns the DescriptorPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorPhone() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DescriptorPhone
}

// GetDescriptorPhoneOk returns a tuple with the DescriptorPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorPhoneOk() (*interface{}, bool) {
	if o == nil || o.DescriptorPhone == nil {
		return nil, false
	}
	return &o.DescriptorPhone, true
}

// HasDescriptorPhone returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasDescriptorPhone() bool {
	if o != nil && o.DescriptorPhone != nil {
		return true
	}

	return false
}

// SetDescriptorPhone gets a reference to the given interface{} and assigns it to the DescriptorPhone field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetDescriptorPhone(v interface{}) {
	o.DescriptorPhone = v
}

// GetDescriptorUrl returns the DescriptorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DescriptorUrl
}

// GetDescriptorUrlOk returns a tuple with the DescriptorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetDescriptorUrlOk() (*interface{}, bool) {
	if o == nil || o.DescriptorUrl == nil {
		return nil, false
	}
	return &o.DescriptorUrl, true
}

// HasDescriptorUrl returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasDescriptorUrl() bool {
	if o != nil && o.DescriptorUrl != nil {
		return true
	}

	return false
}

// SetDescriptorUrl gets a reference to the given interface{} and assigns it to the DescriptorUrl field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetDescriptorUrl(v interface{}) {
	o.DescriptorUrl = v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetWebhookEndpointUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) GetWebhookEndpointUrlOk() (*interface{}, bool) {
	if o == nil || o.WebhookEndpointUrl == nil {
		return nil, false
	}
	return &o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) HasWebhookEndpointUrl() bool {
	if o != nil && o.WebhookEndpointUrl != nil {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given interface{} and assigns it to the WebhookEndpointUrl field.
func (o *GETBraintreeGateways200ResponseDataInnerAttributes) SetWebhookEndpointUrl(v interface{}) {
	o.WebhookEndpointUrl = v
}

func (o GETBraintreeGateways200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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
	if o.DescriptorName != nil {
		toSerialize["descriptor_name"] = o.DescriptorName
	}
	if o.DescriptorPhone != nil {
		toSerialize["descriptor_phone"] = o.DescriptorPhone
	}
	if o.DescriptorUrl != nil {
		toSerialize["descriptor_url"] = o.DescriptorUrl
	}
	if o.WebhookEndpointUrl != nil {
		toSerialize["webhook_endpoint_url"] = o.WebhookEndpointUrl
	}
	return json.Marshal(toSerialize)
}

type NullableGETBraintreeGateways200ResponseDataInnerAttributes struct {
	value *GETBraintreeGateways200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETBraintreeGateways200ResponseDataInnerAttributes) Get() *GETBraintreeGateways200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerAttributes) Set(val *GETBraintreeGateways200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreeGateways200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreeGateways200ResponseDataInnerAttributes(val *GETBraintreeGateways200ResponseDataInnerAttributes) *NullableGETBraintreeGateways200ResponseDataInnerAttributes {
	return &NullableGETBraintreeGateways200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETBraintreeGateways200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
