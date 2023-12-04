/*
NewsCatcher-V3 Production API

Testing SearchSimilarApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package newscatcherapi

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // newscatcherapi "github.com/konfig-dev/newscatcher-sdks/v3/go"
)

func Test_newscatcherapi_SearchSimilarApiService(t *testing.T) {

    // configuration := newscatcherapi.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetApiKey("X_API_TOKEN")
    client := newscatcherapi.NewAPIClient(configuration)
    */

    t.Run("Test SearchSimilarApiService Get", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.SearchSimilarApi.Get(
            "q_example",
        )
        request.SearchIn(""title_content"")
        request.IncludeSimilarDocuments(false)
        request.SimilarDocumentsNumber(5)
        request.SimilarDocumentsFields(""title,content"")
        request.PredefinedSources("predefinedSources_example")
        request.Sources("sources_example")
        request.NotSources("notSources_example")
        request.Lang("lang_example")
        request.NotLang("notLang_example")
        request.Countries("countries_example")
        request.NotCountries("notCountries_example")
        request.From(from)
        request.To(to)
        request.ByParseDate(false)
        request.PublishedDatePrecision("publishedDatePrecision_example")
        request.SortBy(""relevancy"")
        request.RankedOnly("rankedOnly_example")
        request.FromRank(56)
        request.ToRank(56)
        request.IsHeadline(true)
        request.IsPaidContent(true)
        request.ParentUrl("parentUrl_example")
        request.AllLinks("allLinks_example")
        request.AllDomainLinks("allDomainLinks_example")
        request.WordCountMin(56)
        request.WordCountMax(56)
        request.Page(1)
        request.PageSize(100)
        request.IncludeNlpData(true)
        request.HasNlp(true)
        request.Theme("theme_example")
        request.NerName("nerName_example")
        request.TitleSentimentMin(8.14)
        request.TitleSentimentMax(8.14)
        request.ContentSentimentMin(8.14)
        request.ContentSentimentMax(8.14)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test SearchSimilarApiService Post", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        from := *newscatcherapi.NewFrom()
        to := *newscatcherapi.NewTo()
        
        moreLikeThisRequest := *newscatcherapi.NewMoreLikeThisRequest(
            "null",
        )
        moreLikeThisRequest.SetSearchIn("title_content")
        moreLikeThisRequest.SetIncludeSimilarDocuments(false)
        moreLikeThisRequest.SetSimilarDocumentsNumber(5)
        moreLikeThisRequest.SetSimilarDocumentsFields("title,content")
        moreLikeThisRequest.SetPredefinedSources("null")
        moreLikeThisRequest.SetSources("null")
        moreLikeThisRequest.SetNotSources("null")
        moreLikeThisRequest.SetLang("null")
        moreLikeThisRequest.SetNotLang("null")
        moreLikeThisRequest.SetCountries("null")
        moreLikeThisRequest.SetNotCountries("null")
        moreLikeThisRequest.SetFrom(from)
        moreLikeThisRequest.SetTo(to)
        moreLikeThisRequest.SetByParseDate(false)
        moreLikeThisRequest.SetPublishedDatePrecision("null")
        moreLikeThisRequest.SetSortBy("relevancy")
        moreLikeThisRequest.SetRankedOnly("null")
        moreLikeThisRequest.SetFromRank(null)
        moreLikeThisRequest.SetToRank(null)
        moreLikeThisRequest.SetIsHeadline(null)
        moreLikeThisRequest.SetIsPaidContent(null)
        moreLikeThisRequest.SetParentUrl("null")
        moreLikeThisRequest.SetAllLinks("null")
        moreLikeThisRequest.SetAllDomainLinks("null")
        moreLikeThisRequest.SetWordCountMin(null)
        moreLikeThisRequest.SetWordCountMax(null)
        moreLikeThisRequest.SetPage(1)
        moreLikeThisRequest.SetPageSize(100)
        moreLikeThisRequest.SetIncludeNlpData(null)
        moreLikeThisRequest.SetHasNlp(null)
        moreLikeThisRequest.SetTheme("null")
        moreLikeThisRequest.SetNerName("null")
        moreLikeThisRequest.SetTitleSentimentMin(null)
        moreLikeThisRequest.SetTitleSentimentMax(null)
        moreLikeThisRequest.SetContentSentimentMin(null)
        moreLikeThisRequest.SetContentSentimentMax(null)
        
        request := client.SearchSimilarApi.Post(
            moreLikeThisRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}