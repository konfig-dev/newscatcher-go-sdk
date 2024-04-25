# DuplicatesByOriginalIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalArticleId** | **string** |  | 
**Page** | Pointer to **int32** |  | [optional] [default to 1]
**PageSize** | Pointer to **int32** |  | [optional] [default to 100]

## Methods

### NewDuplicatesByOriginalIdRequest

`func NewDuplicatesByOriginalIdRequest(originalArticleId string, ) *DuplicatesByOriginalIdRequest`

NewDuplicatesByOriginalIdRequest instantiates a new DuplicatesByOriginalIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicatesByOriginalIdRequestWithDefaults

`func NewDuplicatesByOriginalIdRequestWithDefaults() *DuplicatesByOriginalIdRequest`

NewDuplicatesByOriginalIdRequestWithDefaults instantiates a new DuplicatesByOriginalIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalArticleId

`func (o *DuplicatesByOriginalIdRequest) GetOriginalArticleId() string`

GetOriginalArticleId returns the OriginalArticleId field if non-nil, zero value otherwise.

### GetOriginalArticleIdOk

`func (o *DuplicatesByOriginalIdRequest) GetOriginalArticleIdOk() (*string, bool)`

GetOriginalArticleIdOk returns a tuple with the OriginalArticleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalArticleId

`func (o *DuplicatesByOriginalIdRequest) SetOriginalArticleId(v string)`

SetOriginalArticleId sets OriginalArticleId field to given value.


### GetPage

`func (o *DuplicatesByOriginalIdRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DuplicatesByOriginalIdRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DuplicatesByOriginalIdRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *DuplicatesByOriginalIdRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *DuplicatesByOriginalIdRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DuplicatesByOriginalIdRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DuplicatesByOriginalIdRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DuplicatesByOriginalIdRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


