# OrderUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHOrdersOrderId200ResponseDataAttributes**](PATCHOrdersOrderId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCreateDataRelationships**](OrderCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderUpdateData

`func NewOrderUpdateData(type_ interface{}, id interface{}, attributes PATCHOrdersOrderId200ResponseDataAttributes, ) *OrderUpdateData`

NewOrderUpdateData instantiates a new OrderUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateDataWithDefaults

`func NewOrderUpdateDataWithDefaults() *OrderUpdateData`

NewOrderUpdateDataWithDefaults instantiates a new OrderUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *OrderUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *OrderUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OrderUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *OrderUpdateData) GetAttributes() PATCHOrdersOrderId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderUpdateData) GetAttributesOk() (*PATCHOrdersOrderId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderUpdateData) SetAttributes(v PATCHOrdersOrderId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderUpdateData) GetRelationships() OrderCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderUpdateData) GetRelationshipsOk() (*OrderCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderUpdateData) SetRelationships(v OrderCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


