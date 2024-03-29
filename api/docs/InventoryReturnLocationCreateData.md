# InventoryReturnLocationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTInventoryReturnLocations201ResponseDataAttributes**](POSTInventoryReturnLocations201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationCreateDataRelationships**](InventoryReturnLocationCreateDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryReturnLocationCreateData

`func NewInventoryReturnLocationCreateData(type_ interface{}, attributes POSTInventoryReturnLocations201ResponseDataAttributes, ) *InventoryReturnLocationCreateData`

NewInventoryReturnLocationCreateData instantiates a new InventoryReturnLocationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryReturnLocationCreateDataWithDefaults

`func NewInventoryReturnLocationCreateDataWithDefaults() *InventoryReturnLocationCreateData`

NewInventoryReturnLocationCreateDataWithDefaults instantiates a new InventoryReturnLocationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryReturnLocationCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryReturnLocationCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryReturnLocationCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InventoryReturnLocationCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InventoryReturnLocationCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *InventoryReturnLocationCreateData) GetAttributes() POSTInventoryReturnLocations201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryReturnLocationCreateData) GetAttributesOk() (*POSTInventoryReturnLocations201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryReturnLocationCreateData) SetAttributes(v POSTInventoryReturnLocations201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryReturnLocationCreateData) GetRelationships() InventoryReturnLocationCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryReturnLocationCreateData) GetRelationshipsOk() (*InventoryReturnLocationCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryReturnLocationCreateData) SetRelationships(v InventoryReturnLocationCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryReturnLocationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


