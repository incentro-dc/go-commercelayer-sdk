# MarketCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The market&#39;s internal name | 
**FacebookPixelId** | Pointer to **string** | The Facebook Pixed ID | [optional] 
**CheckoutUrl** | Pointer to **string** | The checkout URL for this market | [optional] 
**ExternalPricesUrl** | Pointer to **string** | The URL used to fetch prices from an external source | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewMarketCreateDataAttributes

`func NewMarketCreateDataAttributes(name string, ) *MarketCreateDataAttributes`

NewMarketCreateDataAttributes instantiates a new MarketCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketCreateDataAttributesWithDefaults

`func NewMarketCreateDataAttributesWithDefaults() *MarketCreateDataAttributes`

NewMarketCreateDataAttributesWithDefaults instantiates a new MarketCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetFacebookPixelId

`func (o *MarketCreateDataAttributes) GetFacebookPixelId() string`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *MarketCreateDataAttributes) GetFacebookPixelIdOk() (*string, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *MarketCreateDataAttributes) SetFacebookPixelId(v string)`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *MarketCreateDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *MarketCreateDataAttributes) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *MarketCreateDataAttributes) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *MarketCreateDataAttributes) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *MarketCreateDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetExternalPricesUrl

`func (o *MarketCreateDataAttributes) GetExternalPricesUrl() string`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *MarketCreateDataAttributes) GetExternalPricesUrlOk() (*string, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *MarketCreateDataAttributes) SetExternalPricesUrl(v string)`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *MarketCreateDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### GetReference

`func (o *MarketCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MarketCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MarketCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MarketCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *MarketCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *MarketCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *MarketCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *MarketCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *MarketCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MarketCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MarketCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MarketCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


