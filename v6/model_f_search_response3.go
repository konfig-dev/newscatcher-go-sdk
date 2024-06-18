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

// FSearchResponse3 SearchResponse DTO class.
type FSearchResponse3 struct {
	Status *string `json:"status,omitempty"`
	TotalHits *int32 `json:"total_hits,omitempty"`
	Page *int32 `json:"page,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	Articles []ArticlesPropertyInner1 `json:"articles,omitempty"`
	UserInput map[string]interface{} `json:"user_input"`
}

// NewFSearchResponse3 instantiates a new FSearchResponse3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFSearchResponse3(userInput map[string]interface{}) *FSearchResponse3 {
	this := FSearchResponse3{}
	var status string = "ok"
	this.Status = &status
	this.UserInput = userInput
	return &this
}

// NewFSearchResponse3WithDefaults instantiates a new FSearchResponse3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFSearchResponse3WithDefaults() *FSearchResponse3 {
	this := FSearchResponse3{}
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FSearchResponse3) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FSearchResponse3) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FSearchResponse3) SetStatus(v string) {
	o.Status = &v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *FSearchResponse3) GetTotalHits() int32 {
	if o == nil || isNil(o.TotalHits) {
		var ret int32
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetTotalHitsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalHits) {
    return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *FSearchResponse3) HasTotalHits() bool {
	if o != nil && !isNil(o.TotalHits) {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int32 and assigns it to the TotalHits field.
func (o *FSearchResponse3) SetTotalHits(v int32) {
	o.TotalHits = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *FSearchResponse3) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *FSearchResponse3) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *FSearchResponse3) SetPage(v int32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *FSearchResponse3) GetTotalPages() int32 {
	if o == nil || isNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetTotalPagesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *FSearchResponse3) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *FSearchResponse3) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *FSearchResponse3) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *FSearchResponse3) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *FSearchResponse3) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *FSearchResponse3) GetArticles() []ArticlesPropertyInner1 {
	if o == nil || isNil(o.Articles) {
		var ret []ArticlesPropertyInner1
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetArticlesOk() ([]ArticlesPropertyInner1, bool) {
	if o == nil || isNil(o.Articles) {
    return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *FSearchResponse3) HasArticles() bool {
	if o != nil && !isNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []ArticlesPropertyInner1 and assigns it to the Articles field.
func (o *FSearchResponse3) SetArticles(v []ArticlesPropertyInner1) {
	o.Articles = v
}

// GetUserInput returns the UserInput field value
func (o *FSearchResponse3) GetUserInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserInput
}

// GetUserInputOk returns a tuple with the UserInput field value
// and a boolean to check if the value has been set.
func (o *FSearchResponse3) GetUserInputOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.UserInput, true
}

// SetUserInput sets field value
func (o *FSearchResponse3) SetUserInput(v map[string]interface{}) {
	o.UserInput = v
}

func (o FSearchResponse3) MarshalJSON() ([]byte, error) {
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

type NullableFSearchResponse3 struct {
	value *FSearchResponse3
	isSet bool
}

func (v NullableFSearchResponse3) Get() *FSearchResponse3 {
	return v.value
}

func (v *NullableFSearchResponse3) Set(val *FSearchResponse3) {
	v.value = val
	v.isSet = true
}

func (v NullableFSearchResponse3) IsSet() bool {
	return v.isSet
}

func (v *NullableFSearchResponse3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFSearchResponse3(val *FSearchResponse3) *NullableFSearchResponse3 {
	return &NullableFSearchResponse3{value: val, isSet: true}
}

func (v NullableFSearchResponse3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFSearchResponse3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


