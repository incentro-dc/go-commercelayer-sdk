# FixedPricePromotionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "fixed_price_promotions"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**FixedPricePromotionUpdateDataAttributes**](FixedPricePromotionUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionUpdateDataRelationships**](FixedPricePromotionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewFixedPricePromotionUpdateData

`func NewFixedPricePromotionUpdateData(type_ string, id string, attributes FixedPricePromotionUpdateDataAttributes, ) *FixedPricePromotionUpdateData`

NewFixedPricePromotionUpdateData instantiates a new FixedPricePromotionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedPricePromotionUpdateDataWithDefaults

`func NewFixedPricePromotionUpdateDataWithDefaults() *FixedPricePromotionUpdateData`

NewFixedPricePromotionUpdateDataWithDefaults instantiates a new FixedPricePromotionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedPricePromotionUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedPricePromotionUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedPricePromotionUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FixedPricePromotionUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedPricePromotionUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedPricePromotionUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FixedPricePromotionUpdateData) GetAttributes() FixedPricePromotionUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FixedPricePromotionUpdateData) GetAttributesOk() (*FixedPricePromotionUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FixedPricePromotionUpdateData) SetAttributes(v FixedPricePromotionUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FixedPricePromotionUpdateData) GetRelationships() FixedPricePromotionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FixedPricePromotionUpdateData) GetRelationshipsOk() (*FixedPricePromotionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FixedPricePromotionUpdateData) SetRelationships(v FixedPricePromotionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FixedPricePromotionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


