# BraintreeGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes**](PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BraintreeGatewayCreateDataRelationships**](BraintreeGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreeGatewayUpdateData

`func NewBraintreeGatewayUpdateData(type_ interface{}, id interface{}, attributes PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, ) *BraintreeGatewayUpdateData`

NewBraintreeGatewayUpdateData instantiates a new BraintreeGatewayUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreeGatewayUpdateDataWithDefaults

`func NewBraintreeGatewayUpdateDataWithDefaults() *BraintreeGatewayUpdateData`

NewBraintreeGatewayUpdateDataWithDefaults instantiates a new BraintreeGatewayUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BraintreeGatewayUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreeGatewayUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreeGatewayUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *BraintreeGatewayUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BraintreeGatewayUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *BraintreeGatewayUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BraintreeGatewayUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BraintreeGatewayUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *BraintreeGatewayUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BraintreeGatewayUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *BraintreeGatewayUpdateData) GetAttributes() PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreeGatewayUpdateData) GetAttributesOk() (*PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreeGatewayUpdateData) SetAttributes(v PATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BraintreeGatewayUpdateData) GetRelationships() BraintreeGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreeGatewayUpdateData) GetRelationshipsOk() (*BraintreeGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreeGatewayUpdateData) SetRelationships(v BraintreeGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreeGatewayUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


