# SkuOptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "sku_options"]
**Attributes** | [**SkuOptionCreateDataAttributes**](SkuOptionCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**BillingInfoValidationRuleDataRelationships**](BillingInfoValidationRuleDataRelationships.md) |  | [optional] 

## Methods

### NewSkuOptionCreateData

`func NewSkuOptionCreateData(type_ string, attributes SkuOptionCreateDataAttributes, ) *SkuOptionCreateData`

NewSkuOptionCreateData instantiates a new SkuOptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuOptionCreateDataWithDefaults

`func NewSkuOptionCreateDataWithDefaults() *SkuOptionCreateData`

NewSkuOptionCreateDataWithDefaults instantiates a new SkuOptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuOptionCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuOptionCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuOptionCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuOptionCreateData) GetAttributes() SkuOptionCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuOptionCreateData) GetAttributesOk() (*SkuOptionCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuOptionCreateData) SetAttributes(v SkuOptionCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuOptionCreateData) GetRelationships() BillingInfoValidationRuleDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuOptionCreateData) GetRelationshipsOk() (*BillingInfoValidationRuleDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuOptionCreateData) SetRelationships(v BillingInfoValidationRuleDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuOptionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


