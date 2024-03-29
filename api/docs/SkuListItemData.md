# SkuListItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETSkuListItemsSkuListItemId200ResponseDataAttributes**](GETSkuListItemsSkuListItemId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuListItemDataRelationships**](SkuListItemDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListItemData

`func NewSkuListItemData(type_ interface{}, attributes GETSkuListItemsSkuListItemId200ResponseDataAttributes, ) *SkuListItemData`

NewSkuListItemData instantiates a new SkuListItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListItemDataWithDefaults

`func NewSkuListItemDataWithDefaults() *SkuListItemData`

NewSkuListItemDataWithDefaults instantiates a new SkuListItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListItemData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListItemData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListItemData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SkuListItemData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SkuListItemData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SkuListItemData) GetAttributes() GETSkuListItemsSkuListItemId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListItemData) GetAttributesOk() (*GETSkuListItemsSkuListItemId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListItemData) SetAttributes(v GETSkuListItemsSkuListItemId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListItemData) GetRelationships() SkuListItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListItemData) GetRelationshipsOk() (*SkuListItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListItemData) SetRelationships(v SkuListItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


