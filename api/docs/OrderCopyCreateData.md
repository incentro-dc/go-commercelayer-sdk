# OrderCopyCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrderCopies201ResponseDataAttributes**](POSTOrderCopies201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCopyCreateDataRelationships**](OrderCopyCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderCopyCreateData

`func NewOrderCopyCreateData(type_ interface{}, attributes POSTOrderCopies201ResponseDataAttributes, ) *OrderCopyCreateData`

NewOrderCopyCreateData instantiates a new OrderCopyCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCopyCreateDataWithDefaults

`func NewOrderCopyCreateDataWithDefaults() *OrderCopyCreateData`

NewOrderCopyCreateDataWithDefaults instantiates a new OrderCopyCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderCopyCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCopyCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCopyCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderCopyCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderCopyCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderCopyCreateData) GetAttributes() POSTOrderCopies201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderCopyCreateData) GetAttributesOk() (*POSTOrderCopies201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderCopyCreateData) SetAttributes(v POSTOrderCopies201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderCopyCreateData) GetRelationships() OrderCopyCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderCopyCreateData) GetRelationshipsOk() (*OrderCopyCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderCopyCreateData) SetRelationships(v OrderCopyCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderCopyCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


