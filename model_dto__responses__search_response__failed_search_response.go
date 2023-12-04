/*
NewsCatcher-V3 Production API

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>

API version: Beta-3.0.0
Contact: maksym@newscatcherapi.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package newscatcherapi

import (
	"encoding/json"
)

// DtoResponsesSearchResponseFailedSearchResponse struct for DtoResponsesSearchResponseFailedSearchResponse
type DtoResponsesSearchResponseFailedSearchResponse struct {
	Status *string `json:"status,omitempty"`
	TotalHits *int32 `json:"total_hits,omitempty"`
	Page *int32 `json:"page,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	Articles []DtoResponsesSearchResponseArticleResult `json:"articles,omitempty"`
	UserInput map[string]interface{} `json:"user_input"`
}

// NewDtoResponsesSearchResponseFailedSearchResponse instantiates a new DtoResponsesSearchResponseFailedSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoResponsesSearchResponseFailedSearchResponse(userInput map[string]interface{}) *DtoResponsesSearchResponseFailedSearchResponse {
	this := DtoResponsesSearchResponseFailedSearchResponse{}
	var status string = "No Matches for your search"
	this.Status = &status
	var totalHits int32 = 0
	this.TotalHits = &totalHits
	var page int32 = 0
	this.Page = &page
	var totalPages int32 = 0
	this.TotalPages = &totalPages
	var pageSize int32 = 0
	this.PageSize = &pageSize
	this.UserInput = userInput
	return &this
}

// NewDtoResponsesSearchResponseFailedSearchResponseWithDefaults instantiates a new DtoResponsesSearchResponseFailedSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoResponsesSearchResponseFailedSearchResponseWithDefaults() *DtoResponsesSearchResponseFailedSearchResponse {
	this := DtoResponsesSearchResponseFailedSearchResponse{}
	var status string = "No Matches for your search"
	this.Status = &status
	var totalHits int32 = 0
	this.TotalHits = &totalHits
	var page int32 = 0
	this.Page = &page
	var totalPages int32 = 0
	this.TotalPages = &totalPages
	var pageSize int32 = 0
	this.PageSize = &pageSize
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetTotalHits() int32 {
	if o == nil || isNil(o.TotalHits) {
		var ret int32
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetTotalHitsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalHits) {
    return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) HasTotalHits() bool {
	if o != nil && !isNil(o.TotalHits) {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int32 and assigns it to the TotalHits field.
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetTotalHits(v int32) {
	o.TotalHits = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetPage(v int32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetTotalPages() int32 {
	if o == nil || isNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetArticles() []DtoResponsesSearchResponseArticleResult {
	if o == nil || isNil(o.Articles) {
		var ret []DtoResponsesSearchResponseArticleResult
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetArticlesOk() ([]DtoResponsesSearchResponseArticleResult, bool) {
	if o == nil || isNil(o.Articles) {
    return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) HasArticles() bool {
	if o != nil && !isNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []DtoResponsesSearchResponseArticleResult and assigns it to the Articles field.
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetArticles(v []DtoResponsesSearchResponseArticleResult) {
	o.Articles = v
}

// GetUserInput returns the UserInput field value
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetUserInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserInput
}

// GetUserInputOk returns a tuple with the UserInput field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseFailedSearchResponse) GetUserInputOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.UserInput, true
}

// SetUserInput sets field value
func (o *DtoResponsesSearchResponseFailedSearchResponse) SetUserInput(v map[string]interface{}) {
	o.UserInput = v
}

func (o DtoResponsesSearchResponseFailedSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TotalHits) {
		toSerialize["total_hits"] = o.TotalHits
	}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !isNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !isNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	if true {
		toSerialize["user_input"] = o.UserInput
	}
	return json.Marshal(toSerialize)
}

type NullableDtoResponsesSearchResponseFailedSearchResponse struct {
	value *DtoResponsesSearchResponseFailedSearchResponse
	isSet bool
}

func (v NullableDtoResponsesSearchResponseFailedSearchResponse) Get() *DtoResponsesSearchResponseFailedSearchResponse {
	return v.value
}

func (v *NullableDtoResponsesSearchResponseFailedSearchResponse) Set(val *DtoResponsesSearchResponseFailedSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoResponsesSearchResponseFailedSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoResponsesSearchResponseFailedSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoResponsesSearchResponseFailedSearchResponse(val *DtoResponsesSearchResponseFailedSearchResponse) *NullableDtoResponsesSearchResponseFailedSearchResponse {
	return &NullableDtoResponsesSearchResponseFailedSearchResponse{value: val, isSet: true}
}

func (v NullableDtoResponsesSearchResponseFailedSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoResponsesSearchResponseFailedSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


