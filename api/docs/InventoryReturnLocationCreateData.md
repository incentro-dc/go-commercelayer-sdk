# InventoryReturnLocationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "inventory_return_locations"]
**Attributes** | [**InventoryReturnLocationCreateDataAttributes**](InventoryReturnLocationCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationCreateDataRelationships**](InventoryReturnLocationCreateDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryReturnLocationCreateData

`func NewInventoryReturnLocationCreateData(type_ string, attributes InventoryReturnLocationCreateDataAttributes, ) *InventoryReturnLocationCreateData`

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

`func (o *InventoryReturnLocationCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryReturnLocationCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryReturnLocationCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InventoryReturnLocationCreateData) GetAttributes() InventoryReturnLocationCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryReturnLocationCreateData) GetAttributesOk() (*InventoryReturnLocationCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryReturnLocationCreateData) SetAttributes(v InventoryReturnLocationCreateDataAttributes)`

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


