# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TransactionData**](TransactionData.md) |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Transaction) GetData() TransactionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Transaction) GetDataOk() (*TransactionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Transaction) SetData(v TransactionData)`

SetData sets Data field to given value.

### HasData

`func (o *Transaction) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


