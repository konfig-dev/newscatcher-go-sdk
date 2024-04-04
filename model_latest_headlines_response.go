/*
NewsCatcher-V3 Production API

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>

API version: 3.2.16
Contact: maksym@newscatcherapi.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package newscatcherapi

import (
	"encoding/json"
)

// LatestHeadlinesResponse struct for LatestHeadlinesResponse
type LatestHeadlinesResponse struct {
	Status *string `json:"status,omitempty"`
	TotalHits int32 `json:"total_hits"`
	Page int32 `json:"page"`
	TotalPages int32 `json:"total_pages"`
	PageSize int32 `json:"page_size"`
	Articles []map[string]interface{} `json:"articles"`
	UserInput map[string]interface{} `json:"user_input"`
}

// NewLatestHeadlinesResponse instantiates a new LatestHeadlinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLatestHeadlinesResponse(totalHits int32, page int32, totalPages int32, pageSize int32, articles []map[string]interface{}, userInput map[string]interface{}) *LatestHeadlinesResponse {
	this := LatestHeadlinesResponse{}
	var status string = "ok"
	this.Status = &status
	this.TotalHits = totalHits
	this.Page = page
	this.TotalPages = totalPages
	this.PageSize = pageSize
	this.Articles = articles
	this.UserInput = userInput
	return &this
}

// NewLatestHeadlinesResponseWithDefaults instantiates a new LatestHeadlinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLatestHeadlinesResponseWithDefaults() *LatestHeadlinesResponse {
	this := LatestHeadlinesResponse{}
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LatestHeadlinesResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LatestHeadlinesResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LatestHeadlinesResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTotalHits returns the TotalHits field value
func (o *LatestHeadlinesResponse) GetTotalHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetTotalHitsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalHits, true
}

// SetTotalHits sets field value
func (o *LatestHeadlinesResponse) SetTotalHits(v int32) {
	o.TotalHits = v
}

// GetPage returns the Page field value
func (o *LatestHeadlinesResponse) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetPageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *LatestHeadlinesResponse) SetPage(v int32) {
	o.Page = v
}

// GetTotalPages returns the TotalPages field value
func (o *LatestHeadlinesResponse) GetTotalPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *LatestHeadlinesResponse) SetTotalPages(v int32) {
	o.TotalPages = v
}

// GetPageSize returns the PageSize field value
func (o *LatestHeadlinesResponse) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *LatestHeadlinesResponse) SetPageSize(v int32) {
	o.PageSize = v
}

// GetArticles returns the Articles field value
func (o *LatestHeadlinesResponse) GetArticles() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetArticlesOk() ([]map[string]interface{}, bool) {
	if o == nil {
    return nil, false
	}
	return o.Articles, true
}

// SetArticles sets field value
func (o *LatestHeadlinesResponse) SetArticles(v []map[string]interface{}) {
	o.Articles = v
}

// GetUserInput returns the UserInput field value
func (o *LatestHeadlinesResponse) GetUserInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserInput
}

// GetUserInputOk returns a tuple with the UserInput field value
// and a boolean to check if the value has been set.
func (o *LatestHeadlinesResponse) GetUserInputOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.UserInput, true
}

// SetUserInput sets field value
func (o *LatestHeadlinesResponse) SetUserInput(v map[string]interface{}) {
	o.UserInput = v
}

func (o LatestHeadlinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["total_hits"] = o.TotalHits
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["total_pages"] = o.TotalPages
	}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if true {
		toSerialize["articles"] = o.Articles
	}
	if true {
		toSerialize["user_input"] = o.UserInput
	}
	return json.Marshal(toSerialize)
}

type NullableLatestHeadlinesResponse struct {
	value *LatestHeadlinesResponse
	isSet bool
}

func (v NullableLatestHeadlinesResponse) Get() *LatestHeadlinesResponse {
	return v.value
}

func (v *NullableLatestHeadlinesResponse) Set(val *LatestHeadlinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLatestHeadlinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLatestHeadlinesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatestHeadlinesResponse(val *LatestHeadlinesResponse) *NullableLatestHeadlinesResponse {
	return &NullableLatestHeadlinesResponse{value: val, isSet: true}
}

func (v NullableLatestHeadlinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatestHeadlinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


