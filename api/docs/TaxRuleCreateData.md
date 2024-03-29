# TaxRuleCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTTaxRules201ResponseDataAttributes**](POSTTaxRules201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**TaxRuleCreateDataRelationships**](TaxRuleCreateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxRuleCreateData

`func NewTaxRuleCreateData(type_ interface{}, attributes POSTTaxRules201ResponseDataAttributes, ) *TaxRuleCreateData`

NewTaxRuleCreateData instantiates a new TaxRuleCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRuleCreateDataWithDefaults

`func NewTaxRuleCreateDataWithDefaults() *TaxRuleCreateData`

NewTaxRuleCreateDataWithDefaults instantiates a new TaxRuleCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxRuleCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxRuleCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxRuleCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxRuleCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxRuleCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *TaxRuleCreateData) GetAttributes() POSTTaxRules201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxRuleCreateData) GetAttributesOk() (*POSTTaxRules201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxRuleCreateData) SetAttributes(v POSTTaxRules201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxRuleCreateData) GetRelationships() TaxRuleCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxRuleCreateData) GetRelationshipsOk() (*TaxRuleCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxRuleCreateData) SetRelationships(v TaxRuleCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxRuleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


