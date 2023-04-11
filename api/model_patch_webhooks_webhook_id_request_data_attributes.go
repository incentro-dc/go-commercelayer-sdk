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

// checks if the PATCHWebhooksWebhookIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHWebhooksWebhookIdRequestDataAttributes{}

// PATCHWebhooksWebhookIdRequestDataAttributes struct for PATCHWebhooksWebhookIdRequestDataAttributes
type PATCHWebhooksWebhookIdRequestDataAttributes struct {
	// Unique name for the webhook.
	Name interface{} `json:"name,omitempty"`
	// The identifier of the resource/event that will trigger the webhook.
	Topic interface{} `json:"topic,omitempty"`
	// URI where the webhook subscription should send the POST request when the event occurs.
	CallbackUrl interface{} `json:"callback_url,omitempty"`
	// List of related resources that should be included in the webhook body.
	IncludeResources interface{} `json:"include_resources,omitempty"`
	// Send this attribute if you want to reset the circuit breaker associated to this webhook to 'closed' state and zero failures count.
	ResetCircuit interface{} `json:"_reset_circuit,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHWebhooksWebhookIdRequestDataAttributes instantiates a new PATCHWebhooksWebhookIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHWebhooksWebhookIdRequestDataAttributes() *PATCHWebhooksWebhookIdRequestDataAttributes {
	this := PATCHWebhooksWebhookIdRequestDataAttributes{}
	return &this
}

// NewPATCHWebhooksWebhookIdRequestDataAttributesWithDefaults instantiates a new PATCHWebhooksWebhookIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHWebhooksWebhookIdRequestDataAttributesWithDefaults() *PATCHWebhooksWebhookIdRequestDataAttributes {
	this := PATCHWebhooksWebhookIdRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetTopic() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetTopicOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return &o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasTopic() bool {
	if o != nil && IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given interface{} and assigns it to the Topic field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetTopic(v interface{}) {
	o.Topic = v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetCallbackUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetCallbackUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasCallbackUrl() bool {
	if o != nil && IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given interface{} and assigns it to the CallbackUrl field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetCallbackUrl(v interface{}) {
	o.CallbackUrl = v
}

// GetIncludeResources returns the IncludeResources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetIncludeResources() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IncludeResources
}

// GetIncludeResourcesOk returns a tuple with the IncludeResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetIncludeResourcesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.IncludeResources) {
		return nil, false
	}
	return &o.IncludeResources, true
}

// HasIncludeResources returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasIncludeResources() bool {
	if o != nil && IsNil(o.IncludeResources) {
		return true
	}

	return false
}

// SetIncludeResources gets a reference to the given interface{} and assigns it to the IncludeResources field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetIncludeResources(v interface{}) {
	o.IncludeResources = v
}

// GetResetCircuit returns the ResetCircuit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetResetCircuit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetCircuit
}

// GetResetCircuitOk returns a tuple with the ResetCircuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetResetCircuitOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetCircuit) {
		return nil, false
	}
	return &o.ResetCircuit, true
}

// HasResetCircuit returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasResetCircuit() bool {
	if o != nil && IsNil(o.ResetCircuit) {
		return true
	}

	return false
}

// SetResetCircuit gets a reference to the given interface{} and assigns it to the ResetCircuit field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetResetCircuit(v interface{}) {
	o.ResetCircuit = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHWebhooksWebhookIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHWebhooksWebhookIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.IncludeResources != nil {
		toSerialize["include_resources"] = o.IncludeResources
	}
	if o.ResetCircuit != nil {
		toSerialize["_reset_circuit"] = o.ResetCircuit
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

type NullablePATCHWebhooksWebhookIdRequestDataAttributes struct {
	value *PATCHWebhooksWebhookIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHWebhooksWebhookIdRequestDataAttributes) Get() *PATCHWebhooksWebhookIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHWebhooksWebhookIdRequestDataAttributes) Set(val *PATCHWebhooksWebhookIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHWebhooksWebhookIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHWebhooksWebhookIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHWebhooksWebhookIdRequestDataAttributes(val *PATCHWebhooksWebhookIdRequestDataAttributes) *NullablePATCHWebhooksWebhookIdRequestDataAttributes {
	return &NullablePATCHWebhooksWebhookIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHWebhooksWebhookIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHWebhooksWebhookIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}