# StripePaymentCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **map[string]interface{}** | Stripe payment options: &#39;setup_future_usage&#39;, &#39;customer&#39;, and &#39;payment_method&#39; | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewStripePaymentCreateDataAttributes

`func NewStripePaymentCreateDataAttributes() *StripePaymentCreateDataAttributes`

NewStripePaymentCreateDataAttributes instantiates a new StripePaymentCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentCreateDataAttributesWithDefaults

`func NewStripePaymentCreateDataAttributesWithDefaults() *StripePaymentCreateDataAttributes`

NewStripePaymentCreateDataAttributesWithDefaults instantiates a new StripePaymentCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *StripePaymentCreateDataAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StripePaymentCreateDataAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StripePaymentCreateDataAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StripePaymentCreateDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReference

`func (o *StripePaymentCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StripePaymentCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StripePaymentCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StripePaymentCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *StripePaymentCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *StripePaymentCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *StripePaymentCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *StripePaymentCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *StripePaymentCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StripePaymentCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StripePaymentCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StripePaymentCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


