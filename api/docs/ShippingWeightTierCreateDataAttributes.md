# ShippingWeightTierCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The shipping method tier&#39;s name | 
**UpTo** | Pointer to **float32** | The tier upper limit. When &#39;null&#39; it means infinity (useful to have an always matching tier). | [optional] 
**PriceAmountCents** | **int32** | The price of this shipping method tier, in cents. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShippingWeightTierCreateDataAttributes

`func NewShippingWeightTierCreateDataAttributes(name string, priceAmountCents int32, ) *ShippingWeightTierCreateDataAttributes`

NewShippingWeightTierCreateDataAttributes instantiates a new ShippingWeightTierCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingWeightTierCreateDataAttributesWithDefaults

`func NewShippingWeightTierCreateDataAttributesWithDefaults() *ShippingWeightTierCreateDataAttributes`

NewShippingWeightTierCreateDataAttributesWithDefaults instantiates a new ShippingWeightTierCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShippingWeightTierCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingWeightTierCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingWeightTierCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetUpTo

`func (o *ShippingWeightTierCreateDataAttributes) GetUpTo() float32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *ShippingWeightTierCreateDataAttributes) GetUpToOk() (*float32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *ShippingWeightTierCreateDataAttributes) SetUpTo(v float32)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *ShippingWeightTierCreateDataAttributes) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### GetPriceAmountCents

`func (o *ShippingWeightTierCreateDataAttributes) GetPriceAmountCents() int32`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *ShippingWeightTierCreateDataAttributes) GetPriceAmountCentsOk() (*int32, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *ShippingWeightTierCreateDataAttributes) SetPriceAmountCents(v int32)`

SetPriceAmountCents sets PriceAmountCents field to given value.


### GetReference

`func (o *ShippingWeightTierCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShippingWeightTierCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShippingWeightTierCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShippingWeightTierCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShippingWeightTierCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShippingWeightTierCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShippingWeightTierCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShippingWeightTierCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingWeightTierCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingWeightTierCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingWeightTierCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingWeightTierCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


