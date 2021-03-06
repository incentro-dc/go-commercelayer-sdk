# FreeGiftPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "free_gift_promotions"]
**Attributes** | [**FreeGiftPromotionCreateDataAttributes**](FreeGiftPromotionCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionCreateDataRelationships**](FixedPricePromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFreeGiftPromotionCreateData

`func NewFreeGiftPromotionCreateData(type_ string, attributes FreeGiftPromotionCreateDataAttributes, ) *FreeGiftPromotionCreateData`

NewFreeGiftPromotionCreateData instantiates a new FreeGiftPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeGiftPromotionCreateDataWithDefaults

`func NewFreeGiftPromotionCreateDataWithDefaults() *FreeGiftPromotionCreateData`

NewFreeGiftPromotionCreateDataWithDefaults instantiates a new FreeGiftPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FreeGiftPromotionCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeGiftPromotionCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeGiftPromotionCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *FreeGiftPromotionCreateData) GetAttributes() FreeGiftPromotionCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FreeGiftPromotionCreateData) GetAttributesOk() (*FreeGiftPromotionCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FreeGiftPromotionCreateData) SetAttributes(v FreeGiftPromotionCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FreeGiftPromotionCreateData) GetRelationships() FixedPricePromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FreeGiftPromotionCreateData) GetRelationshipsOk() (*FixedPricePromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FreeGiftPromotionCreateData) SetRelationships(v FixedPricePromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FreeGiftPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


