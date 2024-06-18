# ArticlesPropertyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Authors** | Pointer to [**ArticlesPropertyInnerAuthors**](ArticlesPropertyInnerAuthors.md) |  | [optional] 
**Journalists** | Pointer to [**ArticlesPropertyInnerJournalists**](ArticlesPropertyInnerJournalists.md) |  | [optional] 
**PublishedDate** | Pointer to **string** |  | [optional] 
**PublishedDatePrecision** | Pointer to **string** |  | [optional] 
**UpdatedDate** | Pointer to **string** |  | [optional] 
**UpdatedDatePrecision** | Pointer to **string** |  | [optional] 
**ParseDate** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**DomainUrl** | Pointer to **string** |  | [optional] 
**FullDomainUrl** | Pointer to **string** |  | [optional] 
**NameSource** | Pointer to **string** |  | [optional] 
**IsHeadline** | Pointer to **string** |  | [optional] 
**PaidContent** | Pointer to **bool** |  | [optional] 
**ExtractionData** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Rights** | Pointer to **string** |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Media** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**WordCount** | Pointer to **int32** |  | [optional] [default to 0]
**IsOpinion** | Pointer to **bool** |  | [optional] 
**TwitterAccount** | Pointer to **string** |  | [optional] 
**AllLinks** | Pointer to [**ArticlesPropertyInnerAllLinks**](ArticlesPropertyInnerAllLinks.md) |  | [optional] 
**AllDomainLinks** | Pointer to [**ArticlesPropertyInnerAllDomainLinks**](ArticlesPropertyInnerAllDomainLinks.md) |  | [optional] 
**Nlp** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Id** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 

## Methods

### NewArticlesPropertyInner

`func NewArticlesPropertyInner() *ArticlesPropertyInner`

NewArticlesPropertyInner instantiates a new ArticlesPropertyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticlesPropertyInnerWithDefaults

`func NewArticlesPropertyInnerWithDefaults() *ArticlesPropertyInner`

NewArticlesPropertyInnerWithDefaults instantiates a new ArticlesPropertyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ArticlesPropertyInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArticlesPropertyInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArticlesPropertyInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArticlesPropertyInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ArticlesPropertyInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticlesPropertyInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticlesPropertyInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArticlesPropertyInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthor

`func (o *ArticlesPropertyInner) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ArticlesPropertyInner) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ArticlesPropertyInner) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ArticlesPropertyInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthors

`func (o *ArticlesPropertyInner) GetAuthors() ArticlesPropertyInnerAuthors`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ArticlesPropertyInner) GetAuthorsOk() (*ArticlesPropertyInnerAuthors, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ArticlesPropertyInner) SetAuthors(v ArticlesPropertyInnerAuthors)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *ArticlesPropertyInner) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetJournalists

`func (o *ArticlesPropertyInner) GetJournalists() ArticlesPropertyInnerJournalists`

GetJournalists returns the Journalists field if non-nil, zero value otherwise.

### GetJournalistsOk

`func (o *ArticlesPropertyInner) GetJournalistsOk() (*ArticlesPropertyInnerJournalists, bool)`

GetJournalistsOk returns a tuple with the Journalists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalists

`func (o *ArticlesPropertyInner) SetJournalists(v ArticlesPropertyInnerJournalists)`

SetJournalists sets Journalists field to given value.

### HasJournalists

`func (o *ArticlesPropertyInner) HasJournalists() bool`

HasJournalists returns a boolean if a field has been set.

### GetPublishedDate

`func (o *ArticlesPropertyInner) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ArticlesPropertyInner) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ArticlesPropertyInner) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *ArticlesPropertyInner) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetPublishedDatePrecision

`func (o *ArticlesPropertyInner) GetPublishedDatePrecision() string`

GetPublishedDatePrecision returns the PublishedDatePrecision field if non-nil, zero value otherwise.

### GetPublishedDatePrecisionOk

`func (o *ArticlesPropertyInner) GetPublishedDatePrecisionOk() (*string, bool)`

GetPublishedDatePrecisionOk returns a tuple with the PublishedDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDatePrecision

`func (o *ArticlesPropertyInner) SetPublishedDatePrecision(v string)`

SetPublishedDatePrecision sets PublishedDatePrecision field to given value.

### HasPublishedDatePrecision

`func (o *ArticlesPropertyInner) HasPublishedDatePrecision() bool`

HasPublishedDatePrecision returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *ArticlesPropertyInner) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *ArticlesPropertyInner) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *ArticlesPropertyInner) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *ArticlesPropertyInner) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetUpdatedDatePrecision

`func (o *ArticlesPropertyInner) GetUpdatedDatePrecision() string`

GetUpdatedDatePrecision returns the UpdatedDatePrecision field if non-nil, zero value otherwise.

### GetUpdatedDatePrecisionOk

`func (o *ArticlesPropertyInner) GetUpdatedDatePrecisionOk() (*string, bool)`

GetUpdatedDatePrecisionOk returns a tuple with the UpdatedDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDatePrecision

`func (o *ArticlesPropertyInner) SetUpdatedDatePrecision(v string)`

SetUpdatedDatePrecision sets UpdatedDatePrecision field to given value.

### HasUpdatedDatePrecision

`func (o *ArticlesPropertyInner) HasUpdatedDatePrecision() bool`

HasUpdatedDatePrecision returns a boolean if a field has been set.

### GetParseDate

`func (o *ArticlesPropertyInner) GetParseDate() string`

GetParseDate returns the ParseDate field if non-nil, zero value otherwise.

### GetParseDateOk

`func (o *ArticlesPropertyInner) GetParseDateOk() (*string, bool)`

GetParseDateOk returns a tuple with the ParseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseDate

`func (o *ArticlesPropertyInner) SetParseDate(v string)`

SetParseDate sets ParseDate field to given value.

### HasParseDate

`func (o *ArticlesPropertyInner) HasParseDate() bool`

HasParseDate returns a boolean if a field has been set.

### GetLink

`func (o *ArticlesPropertyInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ArticlesPropertyInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ArticlesPropertyInner) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ArticlesPropertyInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDomainUrl

`func (o *ArticlesPropertyInner) GetDomainUrl() string`

GetDomainUrl returns the DomainUrl field if non-nil, zero value otherwise.

### GetDomainUrlOk

`func (o *ArticlesPropertyInner) GetDomainUrlOk() (*string, bool)`

GetDomainUrlOk returns a tuple with the DomainUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUrl

`func (o *ArticlesPropertyInner) SetDomainUrl(v string)`

SetDomainUrl sets DomainUrl field to given value.

### HasDomainUrl

`func (o *ArticlesPropertyInner) HasDomainUrl() bool`

HasDomainUrl returns a boolean if a field has been set.

### GetFullDomainUrl

`func (o *ArticlesPropertyInner) GetFullDomainUrl() string`

GetFullDomainUrl returns the FullDomainUrl field if non-nil, zero value otherwise.

### GetFullDomainUrlOk

`func (o *ArticlesPropertyInner) GetFullDomainUrlOk() (*string, bool)`

GetFullDomainUrlOk returns a tuple with the FullDomainUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDomainUrl

`func (o *ArticlesPropertyInner) SetFullDomainUrl(v string)`

SetFullDomainUrl sets FullDomainUrl field to given value.

### HasFullDomainUrl

`func (o *ArticlesPropertyInner) HasFullDomainUrl() bool`

HasFullDomainUrl returns a boolean if a field has been set.

### GetNameSource

`func (o *ArticlesPropertyInner) GetNameSource() string`

GetNameSource returns the NameSource field if non-nil, zero value otherwise.

### GetNameSourceOk

`func (o *ArticlesPropertyInner) GetNameSourceOk() (*string, bool)`

GetNameSourceOk returns a tuple with the NameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSource

`func (o *ArticlesPropertyInner) SetNameSource(v string)`

SetNameSource sets NameSource field to given value.

### HasNameSource

`func (o *ArticlesPropertyInner) HasNameSource() bool`

HasNameSource returns a boolean if a field has been set.

### GetIsHeadline

`func (o *ArticlesPropertyInner) GetIsHeadline() string`

GetIsHeadline returns the IsHeadline field if non-nil, zero value otherwise.

### GetIsHeadlineOk

`func (o *ArticlesPropertyInner) GetIsHeadlineOk() (*string, bool)`

GetIsHeadlineOk returns a tuple with the IsHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHeadline

`func (o *ArticlesPropertyInner) SetIsHeadline(v string)`

SetIsHeadline sets IsHeadline field to given value.

### HasIsHeadline

`func (o *ArticlesPropertyInner) HasIsHeadline() bool`

HasIsHeadline returns a boolean if a field has been set.

### GetPaidContent

`func (o *ArticlesPropertyInner) GetPaidContent() bool`

GetPaidContent returns the PaidContent field if non-nil, zero value otherwise.

### GetPaidContentOk

`func (o *ArticlesPropertyInner) GetPaidContentOk() (*bool, bool)`

GetPaidContentOk returns a tuple with the PaidContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidContent

`func (o *ArticlesPropertyInner) SetPaidContent(v bool)`

SetPaidContent sets PaidContent field to given value.

### HasPaidContent

`func (o *ArticlesPropertyInner) HasPaidContent() bool`

HasPaidContent returns a boolean if a field has been set.

### GetExtractionData

`func (o *ArticlesPropertyInner) GetExtractionData() string`

GetExtractionData returns the ExtractionData field if non-nil, zero value otherwise.

### GetExtractionDataOk

`func (o *ArticlesPropertyInner) GetExtractionDataOk() (*string, bool)`

GetExtractionDataOk returns a tuple with the ExtractionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionData

`func (o *ArticlesPropertyInner) SetExtractionData(v string)`

SetExtractionData sets ExtractionData field to given value.

### HasExtractionData

`func (o *ArticlesPropertyInner) HasExtractionData() bool`

HasExtractionData returns a boolean if a field has been set.

### GetCountry

`func (o *ArticlesPropertyInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ArticlesPropertyInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ArticlesPropertyInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ArticlesPropertyInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRights

`func (o *ArticlesPropertyInner) GetRights() string`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *ArticlesPropertyInner) GetRightsOk() (*string, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *ArticlesPropertyInner) SetRights(v string)`

SetRights sets Rights field to given value.

### HasRights

`func (o *ArticlesPropertyInner) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetRank

`func (o *ArticlesPropertyInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ArticlesPropertyInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ArticlesPropertyInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ArticlesPropertyInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetMedia

`func (o *ArticlesPropertyInner) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ArticlesPropertyInner) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ArticlesPropertyInner) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ArticlesPropertyInner) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetLanguage

`func (o *ArticlesPropertyInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ArticlesPropertyInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ArticlesPropertyInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ArticlesPropertyInner) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetContent

`func (o *ArticlesPropertyInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ArticlesPropertyInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ArticlesPropertyInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ArticlesPropertyInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetWordCount

`func (o *ArticlesPropertyInner) GetWordCount() int32`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *ArticlesPropertyInner) GetWordCountOk() (*int32, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *ArticlesPropertyInner) SetWordCount(v int32)`

SetWordCount sets WordCount field to given value.

### HasWordCount

`func (o *ArticlesPropertyInner) HasWordCount() bool`

HasWordCount returns a boolean if a field has been set.

### GetIsOpinion

`func (o *ArticlesPropertyInner) GetIsOpinion() bool`

GetIsOpinion returns the IsOpinion field if non-nil, zero value otherwise.

### GetIsOpinionOk

`func (o *ArticlesPropertyInner) GetIsOpinionOk() (*bool, bool)`

GetIsOpinionOk returns a tuple with the IsOpinion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpinion

`func (o *ArticlesPropertyInner) SetIsOpinion(v bool)`

SetIsOpinion sets IsOpinion field to given value.

### HasIsOpinion

`func (o *ArticlesPropertyInner) HasIsOpinion() bool`

HasIsOpinion returns a boolean if a field has been set.

### GetTwitterAccount

`func (o *ArticlesPropertyInner) GetTwitterAccount() string`

GetTwitterAccount returns the TwitterAccount field if non-nil, zero value otherwise.

### GetTwitterAccountOk

`func (o *ArticlesPropertyInner) GetTwitterAccountOk() (*string, bool)`

GetTwitterAccountOk returns a tuple with the TwitterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterAccount

`func (o *ArticlesPropertyInner) SetTwitterAccount(v string)`

SetTwitterAccount sets TwitterAccount field to given value.

### HasTwitterAccount

`func (o *ArticlesPropertyInner) HasTwitterAccount() bool`

HasTwitterAccount returns a boolean if a field has been set.

### GetAllLinks

`func (o *ArticlesPropertyInner) GetAllLinks() ArticlesPropertyInnerAllLinks`

GetAllLinks returns the AllLinks field if non-nil, zero value otherwise.

### GetAllLinksOk

`func (o *ArticlesPropertyInner) GetAllLinksOk() (*ArticlesPropertyInnerAllLinks, bool)`

GetAllLinksOk returns a tuple with the AllLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLinks

`func (o *ArticlesPropertyInner) SetAllLinks(v ArticlesPropertyInnerAllLinks)`

SetAllLinks sets AllLinks field to given value.

### HasAllLinks

`func (o *ArticlesPropertyInner) HasAllLinks() bool`

HasAllLinks returns a boolean if a field has been set.

### GetAllDomainLinks

`func (o *ArticlesPropertyInner) GetAllDomainLinks() ArticlesPropertyInnerAllDomainLinks`

GetAllDomainLinks returns the AllDomainLinks field if non-nil, zero value otherwise.

### GetAllDomainLinksOk

`func (o *ArticlesPropertyInner) GetAllDomainLinksOk() (*ArticlesPropertyInnerAllDomainLinks, bool)`

GetAllDomainLinksOk returns a tuple with the AllDomainLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomainLinks

`func (o *ArticlesPropertyInner) SetAllDomainLinks(v ArticlesPropertyInnerAllDomainLinks)`

SetAllDomainLinks sets AllDomainLinks field to given value.

### HasAllDomainLinks

`func (o *ArticlesPropertyInner) HasAllDomainLinks() bool`

HasAllDomainLinks returns a boolean if a field has been set.

### GetNlp

`func (o *ArticlesPropertyInner) GetNlp() map[string]interface{}`

GetNlp returns the Nlp field if non-nil, zero value otherwise.

### GetNlpOk

`func (o *ArticlesPropertyInner) GetNlpOk() (*map[string]interface{}, bool)`

GetNlpOk returns a tuple with the Nlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNlp

`func (o *ArticlesPropertyInner) SetNlp(v map[string]interface{})`

SetNlp sets Nlp field to given value.

### HasNlp

`func (o *ArticlesPropertyInner) HasNlp() bool`

HasNlp returns a boolean if a field has been set.

### GetId

`func (o *ArticlesPropertyInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticlesPropertyInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticlesPropertyInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArticlesPropertyInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScore

`func (o *ArticlesPropertyInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ArticlesPropertyInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ArticlesPropertyInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ArticlesPropertyInner) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


