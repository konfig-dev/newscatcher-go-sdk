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

// DtoResponsesSearchResponseArticleResult struct for DtoResponsesSearchResponseArticleResult
type DtoResponsesSearchResponseArticleResult struct {
	Title string `json:"title"`
	Description *string `json:"description,omitempty"`
	Author *string `json:"author,omitempty"`
	Authors *AuthorsProperty `json:"authors,omitempty"`
	Journalists *JournalistsProperty `json:"journalists,omitempty"`
	PublishedDate *string `json:"published_date,omitempty"`
	PublishedDatePrecision *string `json:"published_date_precision,omitempty"`
	UpdatedDate *string `json:"updated_date,omitempty"`
	UpdatedDatePrecision *string `json:"updated_date_precision,omitempty"`
	ParseDate *string `json:"parse_date,omitempty"`
	Link string `json:"link"`
	DomainUrl string `json:"domain_url"`
	FullDomainUrl string `json:"full_domain_url"`
	NameSource *string `json:"name_source,omitempty"`
	IsHeadline *string `json:"is_headline,omitempty"`
	PaidContent *bool `json:"paid_content,omitempty"`
	ExtractionData string `json:"extraction_data"`
	Country *string `json:"country,omitempty"`
	Rights *string `json:"rights,omitempty"`
	Rank int32 `json:"rank"`
	Media *string `json:"media,omitempty"`
	Language *string `json:"language,omitempty"`
	Content *string `json:"content,omitempty"`
	WordCount *int32 `json:"word_count,omitempty"`
	IsOpinion *bool `json:"is_opinion,omitempty"`
	TwitterAccount *string `json:"twitter_account,omitempty"`
	AllLinks *AllLinksProperty `json:"all_links,omitempty"`
	AllDomainLinks *AllDomainLinksProperty `json:"all_domain_links,omitempty"`
	Nlp map[string]interface{} `json:"nlp,omitempty"`
	Id string `json:"id"`
	Score float32 `json:"score"`
}

// NewDtoResponsesSearchResponseArticleResult instantiates a new DtoResponsesSearchResponseArticleResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoResponsesSearchResponseArticleResult(title string, link string, domainUrl string, fullDomainUrl string, extractionData string, rank int32, id string, score float32) *DtoResponsesSearchResponseArticleResult {
	this := DtoResponsesSearchResponseArticleResult{}
	this.Title = title
	this.Link = link
	this.DomainUrl = domainUrl
	this.FullDomainUrl = fullDomainUrl
	this.ExtractionData = extractionData
	this.Rank = rank
	var wordCount int32 = 0
	this.WordCount = &wordCount
	this.Id = id
	this.Score = score
	return &this
}

// NewDtoResponsesSearchResponseArticleResultWithDefaults instantiates a new DtoResponsesSearchResponseArticleResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoResponsesSearchResponseArticleResultWithDefaults() *DtoResponsesSearchResponseArticleResult {
	this := DtoResponsesSearchResponseArticleResult{}
	var wordCount int32 = 0
	this.WordCount = &wordCount
	return &this
}

// GetTitle returns the Title field value
func (o *DtoResponsesSearchResponseArticleResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoResponsesSearchResponseArticleResult) SetDescription(v string) {
	o.Description = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetAuthor() string {
	if o == nil || isNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetAuthorOk() (*string, bool) {
	if o == nil || isNil(o.Author) {
    return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasAuthor() bool {
	if o != nil && !isNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *DtoResponsesSearchResponseArticleResult) SetAuthor(v string) {
	o.Author = &v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetAuthors() AuthorsProperty {
	if o == nil || isNil(o.Authors) {
		var ret AuthorsProperty
		return ret
	}
	return *o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetAuthorsOk() (*AuthorsProperty, bool) {
	if o == nil || isNil(o.Authors) {
    return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasAuthors() bool {
	if o != nil && !isNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given AuthorsProperty and assigns it to the Authors field.
func (o *DtoResponsesSearchResponseArticleResult) SetAuthors(v AuthorsProperty) {
	o.Authors = &v
}

// GetJournalists returns the Journalists field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetJournalists() JournalistsProperty {
	if o == nil || isNil(o.Journalists) {
		var ret JournalistsProperty
		return ret
	}
	return *o.Journalists
}

// GetJournalistsOk returns a tuple with the Journalists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetJournalistsOk() (*JournalistsProperty, bool) {
	if o == nil || isNil(o.Journalists) {
    return nil, false
	}
	return o.Journalists, true
}

// HasJournalists returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasJournalists() bool {
	if o != nil && !isNil(o.Journalists) {
		return true
	}

	return false
}

// SetJournalists gets a reference to the given JournalistsProperty and assigns it to the Journalists field.
func (o *DtoResponsesSearchResponseArticleResult) SetJournalists(v JournalistsProperty) {
	o.Journalists = &v
}

// GetPublishedDate returns the PublishedDate field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetPublishedDate() string {
	if o == nil || isNil(o.PublishedDate) {
		var ret string
		return ret
	}
	return *o.PublishedDate
}

// GetPublishedDateOk returns a tuple with the PublishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetPublishedDateOk() (*string, bool) {
	if o == nil || isNil(o.PublishedDate) {
    return nil, false
	}
	return o.PublishedDate, true
}

// HasPublishedDate returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasPublishedDate() bool {
	if o != nil && !isNil(o.PublishedDate) {
		return true
	}

	return false
}

// SetPublishedDate gets a reference to the given string and assigns it to the PublishedDate field.
func (o *DtoResponsesSearchResponseArticleResult) SetPublishedDate(v string) {
	o.PublishedDate = &v
}

// GetPublishedDatePrecision returns the PublishedDatePrecision field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetPublishedDatePrecision() string {
	if o == nil || isNil(o.PublishedDatePrecision) {
		var ret string
		return ret
	}
	return *o.PublishedDatePrecision
}

// GetPublishedDatePrecisionOk returns a tuple with the PublishedDatePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetPublishedDatePrecisionOk() (*string, bool) {
	if o == nil || isNil(o.PublishedDatePrecision) {
    return nil, false
	}
	return o.PublishedDatePrecision, true
}

// HasPublishedDatePrecision returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasPublishedDatePrecision() bool {
	if o != nil && !isNil(o.PublishedDatePrecision) {
		return true
	}

	return false
}

// SetPublishedDatePrecision gets a reference to the given string and assigns it to the PublishedDatePrecision field.
func (o *DtoResponsesSearchResponseArticleResult) SetPublishedDatePrecision(v string) {
	o.PublishedDatePrecision = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetUpdatedDate() string {
	if o == nil || isNil(o.UpdatedDate) {
		var ret string
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetUpdatedDateOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedDate) {
    return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasUpdatedDate() bool {
	if o != nil && !isNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given string and assigns it to the UpdatedDate field.
func (o *DtoResponsesSearchResponseArticleResult) SetUpdatedDate(v string) {
	o.UpdatedDate = &v
}

// GetUpdatedDatePrecision returns the UpdatedDatePrecision field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetUpdatedDatePrecision() string {
	if o == nil || isNil(o.UpdatedDatePrecision) {
		var ret string
		return ret
	}
	return *o.UpdatedDatePrecision
}

// GetUpdatedDatePrecisionOk returns a tuple with the UpdatedDatePrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetUpdatedDatePrecisionOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedDatePrecision) {
    return nil, false
	}
	return o.UpdatedDatePrecision, true
}

// HasUpdatedDatePrecision returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasUpdatedDatePrecision() bool {
	if o != nil && !isNil(o.UpdatedDatePrecision) {
		return true
	}

	return false
}

// SetUpdatedDatePrecision gets a reference to the given string and assigns it to the UpdatedDatePrecision field.
func (o *DtoResponsesSearchResponseArticleResult) SetUpdatedDatePrecision(v string) {
	o.UpdatedDatePrecision = &v
}

// GetParseDate returns the ParseDate field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetParseDate() string {
	if o == nil || isNil(o.ParseDate) {
		var ret string
		return ret
	}
	return *o.ParseDate
}

// GetParseDateOk returns a tuple with the ParseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetParseDateOk() (*string, bool) {
	if o == nil || isNil(o.ParseDate) {
    return nil, false
	}
	return o.ParseDate, true
}

// HasParseDate returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasParseDate() bool {
	if o != nil && !isNil(o.ParseDate) {
		return true
	}

	return false
}

// SetParseDate gets a reference to the given string and assigns it to the ParseDate field.
func (o *DtoResponsesSearchResponseArticleResult) SetParseDate(v string) {
	o.ParseDate = &v
}

// GetLink returns the Link field value
func (o *DtoResponsesSearchResponseArticleResult) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetLinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetLink(v string) {
	o.Link = v
}

// GetDomainUrl returns the DomainUrl field value
func (o *DtoResponsesSearchResponseArticleResult) GetDomainUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainUrl
}

// GetDomainUrlOk returns a tuple with the DomainUrl field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetDomainUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DomainUrl, true
}

// SetDomainUrl sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetDomainUrl(v string) {
	o.DomainUrl = v
}

// GetFullDomainUrl returns the FullDomainUrl field value
func (o *DtoResponsesSearchResponseArticleResult) GetFullDomainUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullDomainUrl
}

// GetFullDomainUrlOk returns a tuple with the FullDomainUrl field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetFullDomainUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FullDomainUrl, true
}

// SetFullDomainUrl sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetFullDomainUrl(v string) {
	o.FullDomainUrl = v
}

// GetNameSource returns the NameSource field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetNameSource() string {
	if o == nil || isNil(o.NameSource) {
		var ret string
		return ret
	}
	return *o.NameSource
}

// GetNameSourceOk returns a tuple with the NameSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetNameSourceOk() (*string, bool) {
	if o == nil || isNil(o.NameSource) {
    return nil, false
	}
	return o.NameSource, true
}

// HasNameSource returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasNameSource() bool {
	if o != nil && !isNil(o.NameSource) {
		return true
	}

	return false
}

// SetNameSource gets a reference to the given string and assigns it to the NameSource field.
func (o *DtoResponsesSearchResponseArticleResult) SetNameSource(v string) {
	o.NameSource = &v
}

// GetIsHeadline returns the IsHeadline field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetIsHeadline() string {
	if o == nil || isNil(o.IsHeadline) {
		var ret string
		return ret
	}
	return *o.IsHeadline
}

// GetIsHeadlineOk returns a tuple with the IsHeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetIsHeadlineOk() (*string, bool) {
	if o == nil || isNil(o.IsHeadline) {
    return nil, false
	}
	return o.IsHeadline, true
}

// HasIsHeadline returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasIsHeadline() bool {
	if o != nil && !isNil(o.IsHeadline) {
		return true
	}

	return false
}

// SetIsHeadline gets a reference to the given string and assigns it to the IsHeadline field.
func (o *DtoResponsesSearchResponseArticleResult) SetIsHeadline(v string) {
	o.IsHeadline = &v
}

// GetPaidContent returns the PaidContent field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetPaidContent() bool {
	if o == nil || isNil(o.PaidContent) {
		var ret bool
		return ret
	}
	return *o.PaidContent
}

// GetPaidContentOk returns a tuple with the PaidContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetPaidContentOk() (*bool, bool) {
	if o == nil || isNil(o.PaidContent) {
    return nil, false
	}
	return o.PaidContent, true
}

// HasPaidContent returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasPaidContent() bool {
	if o != nil && !isNil(o.PaidContent) {
		return true
	}

	return false
}

// SetPaidContent gets a reference to the given bool and assigns it to the PaidContent field.
func (o *DtoResponsesSearchResponseArticleResult) SetPaidContent(v bool) {
	o.PaidContent = &v
}

// GetExtractionData returns the ExtractionData field value
func (o *DtoResponsesSearchResponseArticleResult) GetExtractionData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtractionData
}

// GetExtractionDataOk returns a tuple with the ExtractionData field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetExtractionDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtractionData, true
}

// SetExtractionData sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetExtractionData(v string) {
	o.ExtractionData = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *DtoResponsesSearchResponseArticleResult) SetCountry(v string) {
	o.Country = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetRights() string {
	if o == nil || isNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetRightsOk() (*string, bool) {
	if o == nil || isNil(o.Rights) {
    return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasRights() bool {
	if o != nil && !isNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *DtoResponsesSearchResponseArticleResult) SetRights(v string) {
	o.Rights = &v
}

// GetRank returns the Rank field value
func (o *DtoResponsesSearchResponseArticleResult) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetRankOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetRank(v int32) {
	o.Rank = v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetMedia() string {
	if o == nil || isNil(o.Media) {
		var ret string
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetMediaOk() (*string, bool) {
	if o == nil || isNil(o.Media) {
    return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasMedia() bool {
	if o != nil && !isNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given string and assigns it to the Media field.
func (o *DtoResponsesSearchResponseArticleResult) SetMedia(v string) {
	o.Media = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetLanguage() string {
	if o == nil || isNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetLanguageOk() (*string, bool) {
	if o == nil || isNil(o.Language) {
    return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasLanguage() bool {
	if o != nil && !isNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *DtoResponsesSearchResponseArticleResult) SetLanguage(v string) {
	o.Language = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetContent() string {
	if o == nil || isNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetContentOk() (*string, bool) {
	if o == nil || isNil(o.Content) {
    return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasContent() bool {
	if o != nil && !isNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DtoResponsesSearchResponseArticleResult) SetContent(v string) {
	o.Content = &v
}

// GetWordCount returns the WordCount field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetWordCount() int32 {
	if o == nil || isNil(o.WordCount) {
		var ret int32
		return ret
	}
	return *o.WordCount
}

// GetWordCountOk returns a tuple with the WordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetWordCountOk() (*int32, bool) {
	if o == nil || isNil(o.WordCount) {
    return nil, false
	}
	return o.WordCount, true
}

// HasWordCount returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasWordCount() bool {
	if o != nil && !isNil(o.WordCount) {
		return true
	}

	return false
}

// SetWordCount gets a reference to the given int32 and assigns it to the WordCount field.
func (o *DtoResponsesSearchResponseArticleResult) SetWordCount(v int32) {
	o.WordCount = &v
}

// GetIsOpinion returns the IsOpinion field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetIsOpinion() bool {
	if o == nil || isNil(o.IsOpinion) {
		var ret bool
		return ret
	}
	return *o.IsOpinion
}

// GetIsOpinionOk returns a tuple with the IsOpinion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetIsOpinionOk() (*bool, bool) {
	if o == nil || isNil(o.IsOpinion) {
    return nil, false
	}
	return o.IsOpinion, true
}

// HasIsOpinion returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasIsOpinion() bool {
	if o != nil && !isNil(o.IsOpinion) {
		return true
	}

	return false
}

// SetIsOpinion gets a reference to the given bool and assigns it to the IsOpinion field.
func (o *DtoResponsesSearchResponseArticleResult) SetIsOpinion(v bool) {
	o.IsOpinion = &v
}

// GetTwitterAccount returns the TwitterAccount field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetTwitterAccount() string {
	if o == nil || isNil(o.TwitterAccount) {
		var ret string
		return ret
	}
	return *o.TwitterAccount
}

// GetTwitterAccountOk returns a tuple with the TwitterAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetTwitterAccountOk() (*string, bool) {
	if o == nil || isNil(o.TwitterAccount) {
    return nil, false
	}
	return o.TwitterAccount, true
}

// HasTwitterAccount returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasTwitterAccount() bool {
	if o != nil && !isNil(o.TwitterAccount) {
		return true
	}

	return false
}

// SetTwitterAccount gets a reference to the given string and assigns it to the TwitterAccount field.
func (o *DtoResponsesSearchResponseArticleResult) SetTwitterAccount(v string) {
	o.TwitterAccount = &v
}

// GetAllLinks returns the AllLinks field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetAllLinks() AllLinksProperty {
	if o == nil || isNil(o.AllLinks) {
		var ret AllLinksProperty
		return ret
	}
	return *o.AllLinks
}

// GetAllLinksOk returns a tuple with the AllLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetAllLinksOk() (*AllLinksProperty, bool) {
	if o == nil || isNil(o.AllLinks) {
    return nil, false
	}
	return o.AllLinks, true
}

// HasAllLinks returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasAllLinks() bool {
	if o != nil && !isNil(o.AllLinks) {
		return true
	}

	return false
}

// SetAllLinks gets a reference to the given AllLinksProperty and assigns it to the AllLinks field.
func (o *DtoResponsesSearchResponseArticleResult) SetAllLinks(v AllLinksProperty) {
	o.AllLinks = &v
}

// GetAllDomainLinks returns the AllDomainLinks field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetAllDomainLinks() AllDomainLinksProperty {
	if o == nil || isNil(o.AllDomainLinks) {
		var ret AllDomainLinksProperty
		return ret
	}
	return *o.AllDomainLinks
}

// GetAllDomainLinksOk returns a tuple with the AllDomainLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetAllDomainLinksOk() (*AllDomainLinksProperty, bool) {
	if o == nil || isNil(o.AllDomainLinks) {
    return nil, false
	}
	return o.AllDomainLinks, true
}

// HasAllDomainLinks returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasAllDomainLinks() bool {
	if o != nil && !isNil(o.AllDomainLinks) {
		return true
	}

	return false
}

// SetAllDomainLinks gets a reference to the given AllDomainLinksProperty and assigns it to the AllDomainLinks field.
func (o *DtoResponsesSearchResponseArticleResult) SetAllDomainLinks(v AllDomainLinksProperty) {
	o.AllDomainLinks = &v
}

// GetNlp returns the Nlp field value if set, zero value otherwise.
func (o *DtoResponsesSearchResponseArticleResult) GetNlp() map[string]interface{} {
	if o == nil || isNil(o.Nlp) {
		var ret map[string]interface{}
		return ret
	}
	return o.Nlp
}

// GetNlpOk returns a tuple with the Nlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetNlpOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Nlp) {
    return map[string]interface{}{}, false
	}
	return o.Nlp, true
}

// HasNlp returns a boolean if a field has been set.
func (o *DtoResponsesSearchResponseArticleResult) HasNlp() bool {
	if o != nil && !isNil(o.Nlp) {
		return true
	}

	return false
}

// SetNlp gets a reference to the given map[string]interface{} and assigns it to the Nlp field.
func (o *DtoResponsesSearchResponseArticleResult) SetNlp(v map[string]interface{}) {
	o.Nlp = v
}

// GetId returns the Id field value
func (o *DtoResponsesSearchResponseArticleResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetId(v string) {
	o.Id = v
}

// GetScore returns the Score field value
func (o *DtoResponsesSearchResponseArticleResult) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *DtoResponsesSearchResponseArticleResult) GetScoreOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *DtoResponsesSearchResponseArticleResult) SetScore(v float32) {
	o.Score = v
}

func (o DtoResponsesSearchResponseArticleResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !isNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !isNil(o.Journalists) {
		toSerialize["journalists"] = o.Journalists
	}
	if !isNil(o.PublishedDate) {
		toSerialize["published_date"] = o.PublishedDate
	}
	if !isNil(o.PublishedDatePrecision) {
		toSerialize["published_date_precision"] = o.PublishedDatePrecision
	}
	if !isNil(o.UpdatedDate) {
		toSerialize["updated_date"] = o.UpdatedDate
	}
	if !isNil(o.UpdatedDatePrecision) {
		toSerialize["updated_date_precision"] = o.UpdatedDatePrecision
	}
	if !isNil(o.ParseDate) {
		toSerialize["parse_date"] = o.ParseDate
	}
	if true {
		toSerialize["link"] = o.Link
	}
	if true {
		toSerialize["domain_url"] = o.DomainUrl
	}
	if true {
		toSerialize["full_domain_url"] = o.FullDomainUrl
	}
	if !isNil(o.NameSource) {
		toSerialize["name_source"] = o.NameSource
	}
	if !isNil(o.IsHeadline) {
		toSerialize["is_headline"] = o.IsHeadline
	}
	if !isNil(o.PaidContent) {
		toSerialize["paid_content"] = o.PaidContent
	}
	if true {
		toSerialize["extraction_data"] = o.ExtractionData
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	if true {
		toSerialize["rank"] = o.Rank
	}
	if !isNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if !isNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !isNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !isNil(o.WordCount) {
		toSerialize["word_count"] = o.WordCount
	}
	if !isNil(o.IsOpinion) {
		toSerialize["is_opinion"] = o.IsOpinion
	}
	if !isNil(o.TwitterAccount) {
		toSerialize["twitter_account"] = o.TwitterAccount
	}
	if !isNil(o.AllLinks) {
		toSerialize["all_links"] = o.AllLinks
	}
	if !isNil(o.AllDomainLinks) {
		toSerialize["all_domain_links"] = o.AllDomainLinks
	}
	if !isNil(o.Nlp) {
		toSerialize["nlp"] = o.Nlp
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableDtoResponsesSearchResponseArticleResult struct {
	value *DtoResponsesSearchResponseArticleResult
	isSet bool
}

func (v NullableDtoResponsesSearchResponseArticleResult) Get() *DtoResponsesSearchResponseArticleResult {
	return v.value
}

func (v *NullableDtoResponsesSearchResponseArticleResult) Set(val *DtoResponsesSearchResponseArticleResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoResponsesSearchResponseArticleResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoResponsesSearchResponseArticleResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoResponsesSearchResponseArticleResult(val *DtoResponsesSearchResponseArticleResult) *NullableDtoResponsesSearchResponseArticleResult {
	return &NullableDtoResponsesSearchResponseArticleResult{value: val, isSet: true}
}

func (v NullableDtoResponsesSearchResponseArticleResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoResponsesSearchResponseArticleResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


