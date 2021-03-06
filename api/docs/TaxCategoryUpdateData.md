# TaxCategoryUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "tax_categories"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**TaxCategoryUpdateDataAttributes**](TaxCategoryUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**TaxCategoryUpdateDataRelationships**](TaxCategoryUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewTaxCategoryUpdateData

`func NewTaxCategoryUpdateData(type_ string, id string, attributes TaxCategoryUpdateDataAttributes, ) *TaxCategoryUpdateData`

NewTaxCategoryUpdateData instantiates a new TaxCategoryUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCategoryUpdateDataWithDefaults

`func NewTaxCategoryUpdateDataWithDefaults() *TaxCategoryUpdateData`

NewTaxCategoryUpdateDataWithDefaults instantiates a new TaxCategoryUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxCategoryUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCategoryUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCategoryUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TaxCategoryUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCategoryUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCategoryUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *TaxCategoryUpdateData) GetAttributes() TaxCategoryUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxCategoryUpdateData) GetAttributesOk() (*TaxCategoryUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxCategoryUpdateData) SetAttributes(v TaxCategoryUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxCategoryUpdateData) GetRelationships() TaxCategoryUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxCategoryUpdateData) GetRelationshipsOk() (*TaxCategoryUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxCategoryUpdateData) SetRelationships(v TaxCategoryUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxCategoryUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


