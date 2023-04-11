# POSTSkuListItemsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTSkuListItemsRequestDataAttributes**](POSTSkuListItemsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTSkuListItemsRequestDataRelationships**](POSTSkuListItemsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTSkuListItemsRequestData

`func NewPOSTSkuListItemsRequestData(type_ interface{}, attributes POSTSkuListItemsRequestDataAttributes, ) *POSTSkuListItemsRequestData`

NewPOSTSkuListItemsRequestData instantiates a new POSTSkuListItemsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuListItemsRequestDataWithDefaults

`func NewPOSTSkuListItemsRequestDataWithDefaults() *POSTSkuListItemsRequestData`

NewPOSTSkuListItemsRequestDataWithDefaults instantiates a new POSTSkuListItemsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTSkuListItemsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTSkuListItemsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTSkuListItemsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTSkuListItemsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTSkuListItemsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTSkuListItemsRequestData) GetAttributes() POSTSkuListItemsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTSkuListItemsRequestData) GetAttributesOk() (*POSTSkuListItemsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTSkuListItemsRequestData) SetAttributes(v POSTSkuListItemsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTSkuListItemsRequestData) GetRelationships() POSTSkuListItemsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTSkuListItemsRequestData) GetRelationshipsOk() (*POSTSkuListItemsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTSkuListItemsRequestData) SetRelationships(v POSTSkuListItemsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTSkuListItemsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

