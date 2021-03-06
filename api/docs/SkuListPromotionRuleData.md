# SkuListPromotionRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "sku_list_promotion_rules"]
**Attributes** | [**SkuListPromotionRuleDataAttributes**](SkuListPromotionRuleDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuListPromotionRuleDataRelationships**](SkuListPromotionRuleDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListPromotionRuleData

`func NewSkuListPromotionRuleData(type_ string, attributes SkuListPromotionRuleDataAttributes, ) *SkuListPromotionRuleData`

NewSkuListPromotionRuleData instantiates a new SkuListPromotionRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListPromotionRuleDataWithDefaults

`func NewSkuListPromotionRuleDataWithDefaults() *SkuListPromotionRuleData`

NewSkuListPromotionRuleDataWithDefaults instantiates a new SkuListPromotionRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListPromotionRuleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListPromotionRuleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListPromotionRuleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuListPromotionRuleData) GetAttributes() SkuListPromotionRuleDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListPromotionRuleData) GetAttributesOk() (*SkuListPromotionRuleDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListPromotionRuleData) SetAttributes(v SkuListPromotionRuleDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListPromotionRuleData) GetRelationships() SkuListPromotionRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListPromotionRuleData) GetRelationshipsOk() (*SkuListPromotionRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListPromotionRuleData) SetRelationships(v SkuListPromotionRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListPromotionRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


