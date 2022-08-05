# AlertsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**[]Alert**](Alert.md) |  | [optional] 

## Methods

### NewAlertsListResponse

`func NewAlertsListResponse() *AlertsListResponse`

NewAlertsListResponse instantiates a new AlertsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsListResponseWithDefaults

`func NewAlertsListResponseWithDefaults() *AlertsListResponse`

NewAlertsListResponseWithDefaults instantiates a new AlertsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *AlertsListResponse) GetResponse() []Alert`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AlertsListResponse) GetResponseOk() (*[]Alert, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AlertsListResponse) SetResponse(v []Alert)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AlertsListResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


