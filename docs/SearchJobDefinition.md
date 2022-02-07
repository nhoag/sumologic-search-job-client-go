# SearchJobDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Search query. | [optional] 
**From** | Pointer to **string** | Search start time. | [optional] 
**To** | Pointer to **string** | Search end time. | [optional] 
**TimeZone** | Pointer to **string** | Timezone context for search. | [optional] 
**ByReceiptTime** | Pointer to **bool** | Use the receipt time rather than timestamps log messages. | [optional] [default to false]

## Methods

### NewSearchJobDefinition

`func NewSearchJobDefinition() *SearchJobDefinition`

NewSearchJobDefinition instantiates a new SearchJobDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchJobDefinitionWithDefaults

`func NewSearchJobDefinitionWithDefaults() *SearchJobDefinition`

NewSearchJobDefinitionWithDefaults instantiates a new SearchJobDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchJobDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchJobDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchJobDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchJobDefinition) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFrom

`func (o *SearchJobDefinition) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchJobDefinition) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchJobDefinition) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchJobDefinition) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SearchJobDefinition) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SearchJobDefinition) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SearchJobDefinition) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SearchJobDefinition) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTimeZone

`func (o *SearchJobDefinition) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SearchJobDefinition) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SearchJobDefinition) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SearchJobDefinition) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetByReceiptTime

`func (o *SearchJobDefinition) GetByReceiptTime() bool`

GetByReceiptTime returns the ByReceiptTime field if non-nil, zero value otherwise.

### GetByReceiptTimeOk

`func (o *SearchJobDefinition) GetByReceiptTimeOk() (*bool, bool)`

GetByReceiptTimeOk returns a tuple with the ByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReceiptTime

`func (o *SearchJobDefinition) SetByReceiptTime(v bool)`

SetByReceiptTime sets ByReceiptTime field to given value.

### HasByReceiptTime

`func (o *SearchJobDefinition) HasByReceiptTime() bool`

HasByReceiptTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


