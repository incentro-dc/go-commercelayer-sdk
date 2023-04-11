# POSTStockItemsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStockItemsRequestDataAttributes**](POSTStockItemsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTStockItemsRequestDataRelationships**](POSTStockItemsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTStockItemsRequestData

`func NewPOSTStockItemsRequestData(type_ interface{}, attributes POSTStockItemsRequestDataAttributes, ) *POSTStockItemsRequestData`

NewPOSTStockItemsRequestData instantiates a new POSTStockItemsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockItemsRequestDataWithDefaults

`func NewPOSTStockItemsRequestDataWithDefaults() *POSTStockItemsRequestData`

NewPOSTStockItemsRequestDataWithDefaults instantiates a new POSTStockItemsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTStockItemsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStockItemsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStockItemsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTStockItemsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTStockItemsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTStockItemsRequestData) GetAttributes() POSTStockItemsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStockItemsRequestData) GetAttributesOk() (*POSTStockItemsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStockItemsRequestData) SetAttributes(v POSTStockItemsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTStockItemsRequestData) GetRelationships() POSTStockItemsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStockItemsRequestData) GetRelationshipsOk() (*POSTStockItemsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStockItemsRequestData) SetRelationships(v POSTStockItemsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStockItemsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

