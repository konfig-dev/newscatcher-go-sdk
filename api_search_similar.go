/*
NewsCatcher-V3 Production API

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>

API version: 3.2.16
Contact: maksym@newscatcherapi.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package newscatcherapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// SearchSimilarApiService SearchSimilarApi service
type SearchSimilarApiService service

type SearchSimilarApiGetRequest struct {
	ctx context.Context
	ApiService *SearchSimilarApiService
	q string
	searchIn *string
	includeSimilarDocuments *bool
	similarDocumentsNumber *int32
	similarDocumentsFields *string
	predefinedSources *interface{}
	sources *interface{}
	notSources *interface{}
	lang *interface{}
	notLang *interface{}
	countries *interface{}
	notCountries *interface{}
	from *From
	to *To
	byParseDate *bool
	publishedDatePrecision *string
	sortBy *string
	rankedOnly *string
	fromRank *int32
	toRank *int32
	isHeadline *bool
	isPaidContent *bool
	parentUrl *interface{}
	allLinks *interface{}
	allDomainLinks *interface{}
	wordCountMin *int32
	wordCountMax *int32
	page *int32
	pageSize *int32
	includeNlpData *bool
	hasNlp *bool
	theme *string
	notTheme *string
	nerName *string
	titleSentimentMin *float32
	titleSentimentMax *float32
	contentSentimentMin *float32
	contentSentimentMax *float32
	iptcTags *interface{}
	notIptcTags *interface{}
}

func (r *SearchSimilarApiGetRequest) SearchIn(searchIn string) *SearchSimilarApiGetRequest {
	r.searchIn = &searchIn
	return r
}

func (r *SearchSimilarApiGetRequest) IncludeSimilarDocuments(includeSimilarDocuments bool) *SearchSimilarApiGetRequest {
	r.includeSimilarDocuments = &includeSimilarDocuments
	return r
}

func (r *SearchSimilarApiGetRequest) SimilarDocumentsNumber(similarDocumentsNumber int32) *SearchSimilarApiGetRequest {
	r.similarDocumentsNumber = &similarDocumentsNumber
	return r
}

func (r *SearchSimilarApiGetRequest) SimilarDocumentsFields(similarDocumentsFields string) *SearchSimilarApiGetRequest {
	r.similarDocumentsFields = &similarDocumentsFields
	return r
}

func (r *SearchSimilarApiGetRequest) PredefinedSources(predefinedSources interface{}) *SearchSimilarApiGetRequest {
	r.predefinedSources = &predefinedSources
	return r
}

func (r *SearchSimilarApiGetRequest) Sources(sources interface{}) *SearchSimilarApiGetRequest {
	r.sources = &sources
	return r
}

func (r *SearchSimilarApiGetRequest) NotSources(notSources interface{}) *SearchSimilarApiGetRequest {
	r.notSources = &notSources
	return r
}

func (r *SearchSimilarApiGetRequest) Lang(lang interface{}) *SearchSimilarApiGetRequest {
	r.lang = &lang
	return r
}

func (r *SearchSimilarApiGetRequest) NotLang(notLang interface{}) *SearchSimilarApiGetRequest {
	r.notLang = &notLang
	return r
}

func (r *SearchSimilarApiGetRequest) Countries(countries interface{}) *SearchSimilarApiGetRequest {
	r.countries = &countries
	return r
}

func (r *SearchSimilarApiGetRequest) NotCountries(notCountries interface{}) *SearchSimilarApiGetRequest {
	r.notCountries = &notCountries
	return r
}

func (r *SearchSimilarApiGetRequest) From(from From) *SearchSimilarApiGetRequest {
	r.from = &from
	return r
}

func (r *SearchSimilarApiGetRequest) To(to To) *SearchSimilarApiGetRequest {
	r.to = &to
	return r
}

func (r *SearchSimilarApiGetRequest) ByParseDate(byParseDate bool) *SearchSimilarApiGetRequest {
	r.byParseDate = &byParseDate
	return r
}

func (r *SearchSimilarApiGetRequest) PublishedDatePrecision(publishedDatePrecision string) *SearchSimilarApiGetRequest {
	r.publishedDatePrecision = &publishedDatePrecision
	return r
}

func (r *SearchSimilarApiGetRequest) SortBy(sortBy string) *SearchSimilarApiGetRequest {
	r.sortBy = &sortBy
	return r
}

func (r *SearchSimilarApiGetRequest) RankedOnly(rankedOnly string) *SearchSimilarApiGetRequest {
	r.rankedOnly = &rankedOnly
	return r
}

func (r *SearchSimilarApiGetRequest) FromRank(fromRank int32) *SearchSimilarApiGetRequest {
	r.fromRank = &fromRank
	return r
}

func (r *SearchSimilarApiGetRequest) ToRank(toRank int32) *SearchSimilarApiGetRequest {
	r.toRank = &toRank
	return r
}

func (r *SearchSimilarApiGetRequest) IsHeadline(isHeadline bool) *SearchSimilarApiGetRequest {
	r.isHeadline = &isHeadline
	return r
}

func (r *SearchSimilarApiGetRequest) IsPaidContent(isPaidContent bool) *SearchSimilarApiGetRequest {
	r.isPaidContent = &isPaidContent
	return r
}

func (r *SearchSimilarApiGetRequest) ParentUrl(parentUrl interface{}) *SearchSimilarApiGetRequest {
	r.parentUrl = &parentUrl
	return r
}

func (r *SearchSimilarApiGetRequest) AllLinks(allLinks interface{}) *SearchSimilarApiGetRequest {
	r.allLinks = &allLinks
	return r
}

func (r *SearchSimilarApiGetRequest) AllDomainLinks(allDomainLinks interface{}) *SearchSimilarApiGetRequest {
	r.allDomainLinks = &allDomainLinks
	return r
}

func (r *SearchSimilarApiGetRequest) WordCountMin(wordCountMin int32) *SearchSimilarApiGetRequest {
	r.wordCountMin = &wordCountMin
	return r
}

func (r *SearchSimilarApiGetRequest) WordCountMax(wordCountMax int32) *SearchSimilarApiGetRequest {
	r.wordCountMax = &wordCountMax
	return r
}

func (r *SearchSimilarApiGetRequest) Page(page int32) *SearchSimilarApiGetRequest {
	r.page = &page
	return r
}

func (r *SearchSimilarApiGetRequest) PageSize(pageSize int32) *SearchSimilarApiGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *SearchSimilarApiGetRequest) IncludeNlpData(includeNlpData bool) *SearchSimilarApiGetRequest {
	r.includeNlpData = &includeNlpData
	return r
}

func (r *SearchSimilarApiGetRequest) HasNlp(hasNlp bool) *SearchSimilarApiGetRequest {
	r.hasNlp = &hasNlp
	return r
}

func (r *SearchSimilarApiGetRequest) Theme(theme string) *SearchSimilarApiGetRequest {
	r.theme = &theme
	return r
}

func (r *SearchSimilarApiGetRequest) NotTheme(notTheme string) *SearchSimilarApiGetRequest {
	r.notTheme = &notTheme
	return r
}

func (r *SearchSimilarApiGetRequest) NerName(nerName string) *SearchSimilarApiGetRequest {
	r.nerName = &nerName
	return r
}

func (r *SearchSimilarApiGetRequest) TitleSentimentMin(titleSentimentMin float32) *SearchSimilarApiGetRequest {
	r.titleSentimentMin = &titleSentimentMin
	return r
}

func (r *SearchSimilarApiGetRequest) TitleSentimentMax(titleSentimentMax float32) *SearchSimilarApiGetRequest {
	r.titleSentimentMax = &titleSentimentMax
	return r
}

func (r *SearchSimilarApiGetRequest) ContentSentimentMin(contentSentimentMin float32) *SearchSimilarApiGetRequest {
	r.contentSentimentMin = &contentSentimentMin
	return r
}

func (r *SearchSimilarApiGetRequest) ContentSentimentMax(contentSentimentMax float32) *SearchSimilarApiGetRequest {
	r.contentSentimentMax = &contentSentimentMax
	return r
}

func (r *SearchSimilarApiGetRequest) IptcTags(iptcTags interface{}) *SearchSimilarApiGetRequest {
	r.iptcTags = &iptcTags
	return r
}

func (r *SearchSimilarApiGetRequest) NotIptcTags(notIptcTags interface{}) *SearchSimilarApiGetRequest {
	r.notIptcTags = &notIptcTags
	return r
}

func (r SearchSimilarApiGetRequest) Execute() (*SearchSimilarGetResponse, *http.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get [Get] Search For Similar Articles Request

This endpoint returns a list of articles that are similar to the query provided. You also have the option to get similar articles for the results of a search.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param q
 @return SearchSimilarApiGetRequest
*/
func (a *SearchSimilarApiService) Get(q string) SearchSimilarApiGetRequest {
	return SearchSimilarApiGetRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		q: q,
	}
}

// Execute executes the request
//  @return SearchSimilarGetResponse
func (a *SearchSimilarApiService) GetExecute(r SearchSimilarApiGetRequest) (*SearchSimilarGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchSimilarGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchSimilarApiService.Get")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search_similar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if *r.similarDocumentsNumber < 0 {
		return localVarReturnValue, nil, reportError("similarDocumentsNumber must be greater than 0")
	}
	if *r.wordCountMin < 0 {
		return localVarReturnValue, nil, reportError("wordCountMin must be greater than 0")
	}
	if *r.wordCountMax < 0 {
		return localVarReturnValue, nil, reportError("wordCountMax must be greater than 0")
	}
	if *r.page < 0 {
		return localVarReturnValue, nil, reportError("page must be greater than 0")
	}
	if *r.pageSize < 0 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 0")
	}

	localVarQueryParams.Add("q", parameterToString(r.q, ""))
	if r.searchIn != nil {
		localVarQueryParams.Add("search_in", parameterToString(*r.searchIn, ""))
	}
	if r.includeSimilarDocuments != nil {
		localVarQueryParams.Add("include_similar_documents", parameterToString(*r.includeSimilarDocuments, ""))
	}
	if r.similarDocumentsNumber != nil {
		localVarQueryParams.Add("similar_documents_number", parameterToString(*r.similarDocumentsNumber, ""))
	}
	if r.similarDocumentsFields != nil {
		localVarQueryParams.Add("similar_documents_fields", parameterToString(*r.similarDocumentsFields, ""))
	}
	if r.predefinedSources != nil {
		localVarQueryParams.Add("predefined_sources", parameterToString(*r.predefinedSources, ""))
	}
	if r.sources != nil {
		localVarQueryParams.Add("sources", parameterToString(*r.sources, ""))
	}
	if r.notSources != nil {
		localVarQueryParams.Add("not_sources", parameterToString(*r.notSources, ""))
	}
	if r.lang != nil {
		localVarQueryParams.Add("lang", parameterToString(*r.lang, ""))
	}
	if r.notLang != nil {
		localVarQueryParams.Add("not_lang", parameterToString(*r.notLang, ""))
	}
	if r.countries != nil {
		localVarQueryParams.Add("countries", parameterToString(*r.countries, ""))
	}
	if r.notCountries != nil {
		localVarQueryParams.Add("not_countries", parameterToString(*r.notCountries, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from_", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to_", parameterToString(*r.to, ""))
	}
	if r.byParseDate != nil {
		localVarQueryParams.Add("by_parse_date", parameterToString(*r.byParseDate, ""))
	}
	if r.publishedDatePrecision != nil {
		localVarQueryParams.Add("published_date_precision", parameterToString(*r.publishedDatePrecision, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.rankedOnly != nil {
		localVarQueryParams.Add("ranked_only", parameterToString(*r.rankedOnly, ""))
	}
	if r.fromRank != nil {
		localVarQueryParams.Add("from_rank", parameterToString(*r.fromRank, ""))
	}
	if r.toRank != nil {
		localVarQueryParams.Add("to_rank", parameterToString(*r.toRank, ""))
	}
	if r.isHeadline != nil {
		localVarQueryParams.Add("is_headline", parameterToString(*r.isHeadline, ""))
	}
	if r.isPaidContent != nil {
		localVarQueryParams.Add("is_paid_content", parameterToString(*r.isPaidContent, ""))
	}
	if r.parentUrl != nil {
		localVarQueryParams.Add("parent_url", parameterToString(*r.parentUrl, ""))
	}
	if r.allLinks != nil {
		localVarQueryParams.Add("all_links", parameterToString(*r.allLinks, ""))
	}
	if r.allDomainLinks != nil {
		localVarQueryParams.Add("all_domain_links", parameterToString(*r.allDomainLinks, ""))
	}
	if r.wordCountMin != nil {
		localVarQueryParams.Add("word_count_min", parameterToString(*r.wordCountMin, ""))
	}
	if r.wordCountMax != nil {
		localVarQueryParams.Add("word_count_max", parameterToString(*r.wordCountMax, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.includeNlpData != nil {
		localVarQueryParams.Add("include_nlp_data", parameterToString(*r.includeNlpData, ""))
	}
	if r.hasNlp != nil {
		localVarQueryParams.Add("has_nlp", parameterToString(*r.hasNlp, ""))
	}
	if r.theme != nil {
		localVarQueryParams.Add("theme", parameterToString(*r.theme, ""))
	}
	if r.notTheme != nil {
		localVarQueryParams.Add("not_theme", parameterToString(*r.notTheme, ""))
	}
	if r.nerName != nil {
		localVarQueryParams.Add("ner_name", parameterToString(*r.nerName, ""))
	}
	if r.titleSentimentMin != nil {
		localVarQueryParams.Add("title_sentiment_min", parameterToString(*r.titleSentimentMin, ""))
	}
	if r.titleSentimentMax != nil {
		localVarQueryParams.Add("title_sentiment_max", parameterToString(*r.titleSentimentMax, ""))
	}
	if r.contentSentimentMin != nil {
		localVarQueryParams.Add("content_sentiment_min", parameterToString(*r.contentSentimentMin, ""))
	}
	if r.contentSentimentMax != nil {
		localVarQueryParams.Add("content_sentiment_max", parameterToString(*r.contentSentimentMax, ""))
	}
	if r.iptcTags != nil {
		localVarQueryParams.Add("iptc_tags", parameterToString(*r.iptcTags, ""))
	}
	if r.notIptcTags != nil {
		localVarQueryParams.Add("not_iptc_tags", parameterToString(*r.notIptcTags, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SearchSimilarApiPostRequest struct {
	ctx context.Context
	ApiService *SearchSimilarApiService
	moreLikeThisRequest MoreLikeThisRequest
}

func (r SearchSimilarApiPostRequest) Execute() (*SearchSimilarPostResponse, *http.Response, error) {
	return r.ApiService.PostExecute(r)
}

/*
Post [Post] Search For Similar Articles Request

This endpoint returns a list of articles that are similar to the query provided. You also have the option to get similar articles for the results of a search.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param moreLikeThisRequest
 @return SearchSimilarApiPostRequest
*/
func (a *SearchSimilarApiService) Post(moreLikeThisRequest MoreLikeThisRequest) SearchSimilarApiPostRequest {
	return SearchSimilarApiPostRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		moreLikeThisRequest: moreLikeThisRequest,
	}
}

// Execute executes the request
//  @return SearchSimilarPostResponse
func (a *SearchSimilarApiService) PostExecute(r SearchSimilarApiPostRequest) (*SearchSimilarPostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchSimilarPostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchSimilarApiService.Post")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search_similar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.moreLikeThisRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
