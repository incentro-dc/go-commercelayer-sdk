# OrderSubscriptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrderSubscriptions201ResponseDataAttributes**](POSTOrderSubscriptions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderSubscriptionCreateDataRelationships**](OrderSubscriptionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderSubscriptionCreateData

`func NewOrderSubscriptionCreateData(type_ interface{}, attributes POSTOrderSubscriptions201ResponseDataAttributes, ) *OrderSubscriptionCreateData`

NewOrderSubscriptionCreateData instantiates a new OrderSubscriptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionCreateDataWithDefaults

`func NewOrderSubscriptionCreateDataWithDefaults() *OrderSubscriptionCreateData`

NewOrderSubscriptionCreateDataWithDefaults instantiates a new OrderSubscriptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderSubscriptionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderSubscriptionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderSubscriptionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderSubscriptionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderSubscriptionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderSubscriptionCreateData) GetAttributes() POSTOrderSubscriptions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderSubscriptionCreateData) GetAttributesOk() (*POSTOrderSubscriptions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderSubscriptionCreateData) SetAttributes(v POSTOrderSubscriptions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderSubscriptionCreateData) GetRelationships() OrderSubscriptionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderSubscriptionCreateData) GetRelationshipsOk() (*OrderSubscriptionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderSubscriptionCreateData) SetRelationships(v OrderSubscriptionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderSubscriptionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


