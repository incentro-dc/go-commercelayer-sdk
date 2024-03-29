# TaxjarAccountCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTTaxjarAccounts201ResponseDataAttributes**](POSTTaxjarAccounts201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountCreateDataRelationships**](AvalaraAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxjarAccountCreateData

`func NewTaxjarAccountCreateData(type_ interface{}, attributes POSTTaxjarAccounts201ResponseDataAttributes, ) *TaxjarAccountCreateData`

NewTaxjarAccountCreateData instantiates a new TaxjarAccountCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxjarAccountCreateDataWithDefaults

`func NewTaxjarAccountCreateDataWithDefaults() *TaxjarAccountCreateData`

NewTaxjarAccountCreateDataWithDefaults instantiates a new TaxjarAccountCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxjarAccountCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxjarAccountCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxjarAccountCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxjarAccountCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxjarAccountCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *TaxjarAccountCreateData) GetAttributes() POSTTaxjarAccounts201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxjarAccountCreateData) GetAttributesOk() (*POSTTaxjarAccounts201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxjarAccountCreateData) SetAttributes(v POSTTaxjarAccounts201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxjarAccountCreateData) GetRelationships() AvalaraAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxjarAccountCreateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxjarAccountCreateData) SetRelationships(v AvalaraAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxjarAccountCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


