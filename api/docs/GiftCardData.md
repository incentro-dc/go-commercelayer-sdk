# GiftCardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "gift_cards"]
**Attributes** | [**GiftCardDataAttributes**](GiftCardDataAttributes.md) |  | 
**Relationships** | Pointer to [**GiftCardDataRelationships**](GiftCardDataRelationships.md) |  | [optional] 

## Methods

### NewGiftCardData

`func NewGiftCardData(type_ string, attributes GiftCardDataAttributes, ) *GiftCardData`

NewGiftCardData instantiates a new GiftCardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardDataWithDefaults

`func NewGiftCardDataWithDefaults() *GiftCardData`

NewGiftCardDataWithDefaults instantiates a new GiftCardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GiftCardData) GetAttributes() GiftCardDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GiftCardData) GetAttributesOk() (*GiftCardDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GiftCardData) SetAttributes(v GiftCardDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GiftCardData) GetRelationships() GiftCardDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GiftCardData) GetRelationshipsOk() (*GiftCardDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GiftCardData) SetRelationships(v GiftCardDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GiftCardData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


