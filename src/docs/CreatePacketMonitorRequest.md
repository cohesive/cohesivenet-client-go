# CreatePacketMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of packet monitor. Must be conntrack for type&#x3D;conntrack as only one conntrack monitor can run at a time | 
**Type** | **string** | conntrack, netflow or pcap | 
**Interface** | **string** |  | 
**Filter** | **NullableString** | filter strings are particular to the type of packet monitor. For instance, \&quot;-p 8000\&quot; for pcap. Can be empty string. | 
**Duration** | **string** | Indicates length of time to run capture for. Can be forever or some string parsable by the Linux date command | 
**Destination** | **string** | must be file if pcap or conntrack. Otherwise a host should be specified with the prefix \&quot;host\&quot;. E.g. \&quot;host:10.0.3.2:4000\&quot; | 

## Methods

### NewCreatePacketMonitorRequest

`func NewCreatePacketMonitorRequest(name string, type_ string, interface_ string, filter NullableString, duration string, destination string, ) *CreatePacketMonitorRequest`

NewCreatePacketMonitorRequest instantiates a new CreatePacketMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePacketMonitorRequestWithDefaults

`func NewCreatePacketMonitorRequestWithDefaults() *CreatePacketMonitorRequest`

NewCreatePacketMonitorRequestWithDefaults instantiates a new CreatePacketMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePacketMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePacketMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePacketMonitorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreatePacketMonitorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePacketMonitorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePacketMonitorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetInterface

`func (o *CreatePacketMonitorRequest) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CreatePacketMonitorRequest) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CreatePacketMonitorRequest) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetFilter

`func (o *CreatePacketMonitorRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreatePacketMonitorRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreatePacketMonitorRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### SetFilterNil

`func (o *CreatePacketMonitorRequest) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CreatePacketMonitorRequest) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetDuration

`func (o *CreatePacketMonitorRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreatePacketMonitorRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreatePacketMonitorRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetDestination

`func (o *CreatePacketMonitorRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreatePacketMonitorRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreatePacketMonitorRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


