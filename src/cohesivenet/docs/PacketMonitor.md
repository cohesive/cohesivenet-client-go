# PacketMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of packet monitor. Must be conntrack for type&#x3D;conntrack as only one conntrack monitor can run at a time | [optional] 
**Type** | Pointer to **string** | conntrack, netflow or pcap | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **string** | filter strings are particular to the type of packet monitor. For instance, \&quot;-p 8000\&quot; for pcap | [optional] 
**Duration** | Pointer to **string** | Indicates length of time to run capture for. Can be forever or some string parsable by the Linux date command | [optional] 
**Destination** | Pointer to **string** | must be file if pcap or conntrack. Otherwise a host should be specified with the prefix \&quot;host\&quot;. E.g. \&quot;host:10.0.3.2:4000\&quot; | [optional] 

## Methods

### NewPacketMonitor

`func NewPacketMonitor() *PacketMonitor`

NewPacketMonitor instantiates a new PacketMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketMonitorWithDefaults

`func NewPacketMonitorWithDefaults() *PacketMonitor`

NewPacketMonitorWithDefaults instantiates a new PacketMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PacketMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PacketMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PacketMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PacketMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PacketMonitor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PacketMonitor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PacketMonitor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PacketMonitor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInterface

`func (o *PacketMonitor) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PacketMonitor) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PacketMonitor) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PacketMonitor) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetFilter

`func (o *PacketMonitor) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PacketMonitor) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PacketMonitor) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PacketMonitor) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDuration

`func (o *PacketMonitor) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PacketMonitor) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PacketMonitor) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PacketMonitor) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDestination

`func (o *PacketMonitor) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PacketMonitor) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PacketMonitor) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PacketMonitor) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


