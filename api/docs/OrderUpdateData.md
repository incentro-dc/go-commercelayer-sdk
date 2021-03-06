# OrderUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "orders"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**OrderUpdateDataAttributes**](OrderUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCreateDataRelationships**](OrderCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderUpdateData

`func NewOrderUpdateData(type_ string, id string, attributes OrderUpdateDataAttributes, ) *OrderUpdateData`

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

`func (o *OrderUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *OrderUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *OrderUpdateData) GetAttributes() OrderUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderUpdateData) GetAttributesOk() (*OrderUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderUpdateData) SetAttributes(v OrderUpdateDataAttributes)`

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


