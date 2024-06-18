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


// SearchApiService SearchApi service
type SearchApiService service

type SearchApiGetRequest struct {
	ctx context.Context
	ApiService *SearchApiService
	q string
	searchIn *string
	predefinedSources *interface{}
	sources *interface{}
	notSources *interface{}
	lang *interface{}
	notLang *interface{}
	countries *interface{}
	notCountries *interface{}
	notAuthorName *interface{}
	from *string
	to *string
	publishedDatePrecision *string
	byParseDate *bool
	sortBy *string
	rankedOnly *string
	fromRank *int32
	toRank *int32
	isHeadline *bool
	isOpinion *bool
	isPaidContent *bool
	parentUrl *interface{}
	allLinks *interface{}
	allDomainLinks *interface{}
	wordCountMin *int32
	wordCountMax *int32
	page *int32
	pageSize *int32
	clusteringVariable *string
	clusteringEnabled *bool
	clusteringThreshold *float32
	includeNlpData *bool
	hasNlp *bool
	theme *string
	notTheme *string
	oRGEntityName *string
	pEREntityName *string
	lOCEntityName *string
	mISCEntityName *string
	titleSentimentMin *float32
	titleSentimentMax *float32
	contentSentimentMin *float32
	contentSentimentMax *float32
	iptcTags *interface{}
	notIptcTags *interface{}
	sourceName *interface{}
	iabTags *interface{}
	notIabTags *interface{}
	excludeDuplicates *bool
}

func (r *SearchApiGetRequest) SearchIn(searchIn string) *SearchApiGetRequest {
	r.searchIn = &searchIn
	return r
}

func (r *SearchApiGetRequest) PredefinedSources(predefinedSources interface{}) *SearchApiGetRequest {
	r.predefinedSources = &predefinedSources
	return r
}

func (r *SearchApiGetRequest) Sources(sources interface{}) *SearchApiGetRequest {
	r.sources = &sources
	return r
}

func (r *SearchApiGetRequest) NotSources(notSources interface{}) *SearchApiGetRequest {
	r.notSources = &notSources
	return r
}

func (r *SearchApiGetRequest) Lang(lang interface{}) *SearchApiGetRequest {
	r.lang = &lang
	return r
}

func (r *SearchApiGetRequest) NotLang(notLang interface{}) *SearchApiGetRequest {
	r.notLang = &notLang
	return r
}

func (r *SearchApiGetRequest) Countries(countries interface{}) *SearchApiGetRequest {
	r.countries = &countries
	return r
}

func (r *SearchApiGetRequest) NotCountries(notCountries interface{}) *SearchApiGetRequest {
	r.notCountries = &notCountries
	return r
}

func (r *SearchApiGetRequest) NotAuthorName(notAuthorName interface{}) *SearchApiGetRequest {
	r.notAuthorName = &notAuthorName
	return r
}

func (r *SearchApiGetRequest) From(from string) *SearchApiGetRequest {
	r.from = &from
	return r
}

func (r *SearchApiGetRequest) To(to string) *SearchApiGetRequest {
	r.to = &to
	return r
}

func (r *SearchApiGetRequest) PublishedDatePrecision(publishedDatePrecision string) *SearchApiGetRequest {
	r.publishedDatePrecision = &publishedDatePrecision
	return r
}

func (r *SearchApiGetRequest) ByParseDate(byParseDate bool) *SearchApiGetRequest {
	r.byParseDate = &byParseDate
	return r
}

func (r *SearchApiGetRequest) SortBy(sortBy string) *SearchApiGetRequest {
	r.sortBy = &sortBy
	return r
}

func (r *SearchApiGetRequest) RankedOnly(rankedOnly string) *SearchApiGetRequest {
	r.rankedOnly = &rankedOnly
	return r
}

func (r *SearchApiGetRequest) FromRank(fromRank int32) *SearchApiGetRequest {
	r.fromRank = &fromRank
	return r
}

func (r *SearchApiGetRequest) ToRank(toRank int32) *SearchApiGetRequest {
	r.toRank = &toRank
	return r
}

func (r *SearchApiGetRequest) IsHeadline(isHeadline bool) *SearchApiGetRequest {
	r.isHeadline = &isHeadline
	return r
}

func (r *SearchApiGetRequest) IsOpinion(isOpinion bool) *SearchApiGetRequest {
	r.isOpinion = &isOpinion
	return r
}

func (r *SearchApiGetRequest) IsPaidContent(isPaidContent bool) *SearchApiGetRequest {
	r.isPaidContent = &isPaidContent
	return r
}

func (r *SearchApiGetRequest) ParentUrl(parentUrl interface{}) *SearchApiGetRequest {
	r.parentUrl = &parentUrl
	return r
}

func (r *SearchApiGetRequest) AllLinks(allLinks interface{}) *SearchApiGetRequest {
	r.allLinks = &allLinks
	return r
}

func (r *SearchApiGetRequest) AllDomainLinks(allDomainLinks interface{}) *SearchApiGetRequest {
	r.allDomainLinks = &allDomainLinks
	return r
}

func (r *SearchApiGetRequest) WordCountMin(wordCountMin int32) *SearchApiGetRequest {
	r.wordCountMin = &wordCountMin
	return r
}

func (r *SearchApiGetRequest) WordCountMax(wordCountMax int32) *SearchApiGetRequest {
	r.wordCountMax = &wordCountMax
	return r
}

func (r *SearchApiGetRequest) Page(page int32) *SearchApiGetRequest {
	r.page = &page
	return r
}

func (r *SearchApiGetRequest) PageSize(pageSize int32) *SearchApiGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *SearchApiGetRequest) ClusteringVariable(clusteringVariable string) *SearchApiGetRequest {
	r.clusteringVariable = &clusteringVariable
	return r
}

func (r *SearchApiGetRequest) ClusteringEnabled(clusteringEnabled bool) *SearchApiGetRequest {
	r.clusteringEnabled = &clusteringEnabled
	return r
}

func (r *SearchApiGetRequest) ClusteringThreshold(clusteringThreshold float32) *SearchApiGetRequest {
	r.clusteringThreshold = &clusteringThreshold
	return r
}

func (r *SearchApiGetRequest) IncludeNlpData(includeNlpData bool) *SearchApiGetRequest {
	r.includeNlpData = &includeNlpData
	return r
}

func (r *SearchApiGetRequest) HasNlp(hasNlp bool) *SearchApiGetRequest {
	r.hasNlp = &hasNlp
	return r
}

func (r *SearchApiGetRequest) Theme(theme string) *SearchApiGetRequest {
	r.theme = &theme
	return r
}

func (r *SearchApiGetRequest) NotTheme(notTheme string) *SearchApiGetRequest {
	r.notTheme = &notTheme
	return r
}

func (r *SearchApiGetRequest) ORGEntityName(oRGEntityName string) *SearchApiGetRequest {
	r.oRGEntityName = &oRGEntityName
	return r
}

func (r *SearchApiGetRequest) PEREntityName(pEREntityName string) *SearchApiGetRequest {
	r.pEREntityName = &pEREntityName
	return r
}

func (r *SearchApiGetRequest) LOCEntityName(lOCEntityName string) *SearchApiGetRequest {
	r.lOCEntityName = &lOCEntityName
	return r
}

func (r *SearchApiGetRequest) MISCEntityName(mISCEntityName string) *SearchApiGetRequest {
	r.mISCEntityName = &mISCEntityName
	return r
}

func (r *SearchApiGetRequest) TitleSentimentMin(titleSentimentMin float32) *SearchApiGetRequest {
	r.titleSentimentMin = &titleSentimentMin
	return r
}

func (r *SearchApiGetRequest) TitleSentimentMax(titleSentimentMax float32) *SearchApiGetRequest {
	r.titleSentimentMax = &titleSentimentMax
	return r
}

func (r *SearchApiGetRequest) ContentSentimentMin(contentSentimentMin float32) *SearchApiGetRequest {
	r.contentSentimentMin = &contentSentimentMin
	return r
}

func (r *SearchApiGetRequest) ContentSentimentMax(contentSentimentMax float32) *SearchApiGetRequest {
	r.contentSentimentMax = &contentSentimentMax
	return r
}

func (r *SearchApiGetRequest) IptcTags(iptcTags interface{}) *SearchApiGetRequest {
	r.iptcTags = &iptcTags
	return r
}

func (r *SearchApiGetRequest) NotIptcTags(notIptcTags interface{}) *SearchApiGetRequest {
	r.notIptcTags = &notIptcTags
	return r
}

func (r *SearchApiGetRequest) SourceName(sourceName interface{}) *SearchApiGetRequest {
	r.sourceName = &sourceName
	return r
}

func (r *SearchApiGetRequest) IabTags(iabTags interface{}) *SearchApiGetRequest {
	r.iabTags = &iabTags
	return r
}

func (r *SearchApiGetRequest) NotIabTags(notIabTags interface{}) *SearchApiGetRequest {
	r.notIabTags = &notIabTags
	return r
}

func (r *SearchApiGetRequest) ExcludeDuplicates(excludeDuplicates bool) *SearchApiGetRequest {
	r.excludeDuplicates = &excludeDuplicates
	return r
}

func (r SearchApiGetRequest) Execute() (*CSearchResponse, *http.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get [Get] Search For Articles Request

This endpoint allows you to search for articles. You can search for articles by keyword, language, country, source, and more.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param q
 @return SearchApiGetRequest
*/
func (a *SearchApiService) Get(q string) SearchApiGetRequest {
	return SearchApiGetRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		q: q,
	}
}

// Execute executes the request
//  @return CSearchResponse
func (a *SearchApiService) GetExecute(r SearchApiGetRequest) (*CSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.Get")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.wordCountMin != nil && *r.wordCountMin < 0 {
		return localVarReturnValue, nil, reportError("wordCountMin must be greater than 0")
	}
	if r.wordCountMax != nil && *r.wordCountMax < 0 {
		return localVarReturnValue, nil, reportError("wordCountMax must be greater than 0")
	}
	if r.page != nil && *r.page < 0 {
		return localVarReturnValue, nil, reportError("page must be greater than 0")
	}
	if r.pageSize != nil && *r.pageSize < 0 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 0")
	}

	localVarQueryParams.Add("q", parameterToString(r.q, ""))
	if r.searchIn != nil {
		localVarQueryParams.Add("search_in", parameterToString(*r.searchIn, ""))
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
	if r.notAuthorName != nil {
		localVarQueryParams.Add("not_author_name", parameterToString(*r.notAuthorName, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from_", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to_", parameterToString(*r.to, ""))
	}
	if r.publishedDatePrecision != nil {
		localVarQueryParams.Add("published_date_precision", parameterToString(*r.publishedDatePrecision, ""))
	}
	if r.byParseDate != nil {
		localVarQueryParams.Add("by_parse_date", parameterToString(*r.byParseDate, ""))
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
	if r.isOpinion != nil {
		localVarQueryParams.Add("is_opinion", parameterToString(*r.isOpinion, ""))
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
	if r.clusteringVariable != nil {
		localVarQueryParams.Add("clustering_variable", parameterToString(*r.clusteringVariable, ""))
	}
	if r.clusteringEnabled != nil {
		localVarQueryParams.Add("clustering_enabled", parameterToString(*r.clusteringEnabled, ""))
	}
	if r.clusteringThreshold != nil {
		localVarQueryParams.Add("clustering_threshold", parameterToString(*r.clusteringThreshold, ""))
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
	if r.oRGEntityName != nil {
		localVarQueryParams.Add("ORG_entity_name", parameterToString(*r.oRGEntityName, ""))
	}
	if r.pEREntityName != nil {
		localVarQueryParams.Add("PER_entity_name", parameterToString(*r.pEREntityName, ""))
	}
	if r.lOCEntityName != nil {
		localVarQueryParams.Add("LOC_entity_name", parameterToString(*r.lOCEntityName, ""))
	}
	if r.mISCEntityName != nil {
		localVarQueryParams.Add("MISC_entity_name", parameterToString(*r.mISCEntityName, ""))
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
	if r.sourceName != nil {
		localVarQueryParams.Add("source_name", parameterToString(*r.sourceName, ""))
	}
	if r.iabTags != nil {
		localVarQueryParams.Add("iab_tags", parameterToString(*r.iabTags, ""))
	}
	if r.notIabTags != nil {
		localVarQueryParams.Add("not_iab_tags", parameterToString(*r.notIabTags, ""))
	}
	if r.excludeDuplicates != nil {
		localVarQueryParams.Add("exclude_duplicates", parameterToString(*r.excludeDuplicates, ""))
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

type SearchApiPostRequest struct {
	ctx context.Context
	ApiService *SearchApiService
	searchRequest SearchRequest
}

func (r SearchApiPostRequest) Execute() (*CSearchResponse1, *http.Response, error) {
	return r.ApiService.PostExecute(r)
}

/*
Post [Post] Search For Articles Request

This endpoint allows you to search for articles. You can search for articles by keyword, language, country, source, and more.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param searchRequest
 @return SearchApiPostRequest
*/
func (a *SearchApiService) Post(searchRequest SearchRequest) SearchApiPostRequest {
	return SearchApiPostRequest{
		ApiService: a,
		ctx: a.client.cfg.Context,
		searchRequest: searchRequest,
	}
}

// Execute executes the request
//  @return CSearchResponse1
func (a *SearchApiService) PostExecute(r SearchApiPostRequest) (*CSearchResponse1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CSearchResponse1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchApiService.Post")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/search"

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
	localVarPostBody = r.searchRequest
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
