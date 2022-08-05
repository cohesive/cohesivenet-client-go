# LinkEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | Tunnel event, up or down | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**TimestampI** | Pointer to **int32** |  | [optional] 

## Methods

### NewLinkEvent

`func NewLinkEvent() *LinkEvent`

NewLinkEvent instantiates a new LinkEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkEventWithDefaults

`func NewLinkEventWithDefaults() *LinkEvent`

NewLinkEventWithDefaults instantiates a new LinkEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *LinkEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LinkEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LinkEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *LinkEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *LinkEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LinkEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LinkEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LinkEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampI

`func (o *LinkEvent) GetTimestampI() int32`

GetTimestampI returns the TimestampI field if non-nil, zero value otherwise.

### GetTimestampIOk

`func (o *LinkEvent) GetTimestampIOk() (*int32, bool)`

GetTimestampIOk returns a tuple with the TimestampI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampI

`func (o *LinkEvent) SetTimestampI(v int32)`

SetTimestampI sets TimestampI field to given value.

### HasTimestampI

`func (o *LinkEvent) HasTimestampI() bool`

HasTimestampI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


