# PromotionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "promotions"]
**Attributes** | [**FreeShippingPromotionDataAttributes**](FreeShippingPromotionDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPromotionDataRelationships**](ExternalPromotionDataRelationships.md) |  | [optional] 

## Methods

### NewPromotionData

`func NewPromotionData(type_ string, attributes FreeShippingPromotionDataAttributes, ) *PromotionData`

NewPromotionData instantiates a new PromotionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionDataWithDefaults

`func NewPromotionDataWithDefaults() *PromotionData`

NewPromotionDataWithDefaults instantiates a new PromotionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PromotionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PromotionData) GetAttributes() FreeShippingPromotionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PromotionData) GetAttributesOk() (*FreeShippingPromotionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PromotionData) SetAttributes(v FreeShippingPromotionDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PromotionData) GetRelationships() ExternalPromotionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PromotionData) GetRelationshipsOk() (*ExternalPromotionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PromotionData) SetRelationships(v ExternalPromotionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PromotionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


