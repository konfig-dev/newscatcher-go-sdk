# newscatcherapi - Maksym Sugonyaka's Go SDK

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>


## Installation

Add to your project:

```shell
go get github.com/konfig-dev/newscatcher-go-sdk/tree/main
```

## Getting Started

```golang
package main

import (
    "fmt"
    "os"
    newscatcherapi "github.com/konfig-dev/newscatcher-go-sdk/tree/main"
)

func main() {
    configuration := newscatcherapi.NewConfiguration()
    configuration.SetApiKey("X_API_TOKEN")
    client := newscatcherapi.NewAPIClient(configuration)

    request := client.AuthorsApi.Get(
        "authorName_example",
    )
    request.Sources("sources_example")
    request.PredefinedSources("predefinedSources_example")
    request.NotSources("notSources_example")
    request.Lang("lang_example")
    request.NotLang("notLang_example")
    request.Countries("countries_example")
    request.NotCountries("notCountries_example")
    request.From(from)
    request.To(to)
    request.PublishedDatePrecision("publishedDatePrecision_example")
    request.ByParseDate(false)
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

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Get`: AuthorsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorsApi.Get`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.TotalHits`: %v\n", resp.TotalHits)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.Page`: %v\n", resp.Page)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.TotalPages`: %v\n", resp.TotalPages)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.PageSize`: %v\n", resp.PageSize)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.Articles`: %v\n", resp.Articles)
    fmt.Fprintf(os.Stdout, "Response from `AuthorsGetResponse.Get.UserInput`: %v\n", resp.UserInput)
}

```

## Documentation for API Endpoints

All URIs are relative to *https://v3-api.newscatcherapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorsApi* | [**Get**](docs/AuthorsApi.md#get) | **Get** /api/authors | [Get] Search By Author Request
*AuthorsApi* | [**Post**](docs/AuthorsApi.md#post) | **Post** /api/authors | [Post] Search By Author Request
*LatestHeadlinesApi* | [**Get**](docs/LatestHeadlinesApi.md#get) | **Get** /api/latest_headlines | [Get] Search For Latest Headlines Request
*LatestHeadlinesApi* | [**Post**](docs/LatestHeadlinesApi.md#post) | **Post** /api/latest_headlines | [Post] Search For Latest Headlines Request
*SearchApi* | [**Get**](docs/SearchApi.md#get) | **Get** /api/search | [Get] Search For Articles Request
*SearchApi* | [**Post**](docs/SearchApi.md#post) | **Post** /api/search | [Post] Search For Articles Request
*SearchSimilarApi* | [**Get**](docs/SearchSimilarApi.md#get) | **Get** /api/search_similar | [Get] Search For Similar Articles Request
*SearchSimilarApi* | [**Post**](docs/SearchSimilarApi.md#post) | **Post** /api/search_similar | [Post] Search For Similar Articles Request
*SourcesApi* | [**Get**](docs/SourcesApi.md#get) | **Get** /api/sources | [Get] Search For Sources Request
*SourcesApi* | [**Post**](docs/SourcesApi.md#post) | **Post** /api/sources | [Post] Search For Sources Request
*SubscriptionApi* | [**Get**](docs/SubscriptionApi.md#get) | **Get** /api/subscription | [Get] Get My Plan Info
*SubscriptionApi* | [**Post**](docs/SubscriptionApi.md#post) | **Post** /api/subscription | [Post] Get My Plan Info


## Documentation For Models

 - [AllDomainLinksProperty](docs/AllDomainLinksProperty.md)
 - [AllLinksProperty](docs/AllLinksProperty.md)
 - [AuthorSearchRequest](docs/AuthorSearchRequest.md)
 - [AuthorsGetResponse](docs/AuthorsGetResponse.md)
 - [AuthorsPostResponse](docs/AuthorsPostResponse.md)
 - [AuthorsProperty](docs/AuthorsProperty.md)
 - [Cluster](docs/Cluster.md)
 - [ClusteringSearchResponse](docs/ClusteringSearchResponse.md)
 - [CountriesProperty](docs/CountriesProperty.md)
 - [DtoResponsesAuthorSearchResponseArticleResult](docs/DtoResponsesAuthorSearchResponseArticleResult.md)
 - [DtoResponsesAuthorSearchResponseFailedSearchResponse](docs/DtoResponsesAuthorSearchResponseFailedSearchResponse.md)
 - [DtoResponsesAuthorSearchResponseSearchResponse](docs/DtoResponsesAuthorSearchResponseSearchResponse.md)
 - [DtoResponsesLatestHeadlinesResponseArticleResult](docs/DtoResponsesLatestHeadlinesResponseArticleResult.md)
 - [DtoResponsesMoreLikeThisResponseArticleResult](docs/DtoResponsesMoreLikeThisResponseArticleResult.md)
 - [DtoResponsesMoreLikeThisResponseFailedSearchResponse](docs/DtoResponsesMoreLikeThisResponseFailedSearchResponse.md)
 - [DtoResponsesMoreLikeThisResponseSearchResponse](docs/DtoResponsesMoreLikeThisResponseSearchResponse.md)
 - [DtoResponsesSearchResponseArticleResult](docs/DtoResponsesSearchResponseArticleResult.md)
 - [DtoResponsesSearchResponseFailedSearchResponse](docs/DtoResponsesSearchResponseFailedSearchResponse.md)
 - [DtoResponsesSearchResponseSearchResponse](docs/DtoResponsesSearchResponseSearchResponse.md)
 - [FailedLatestHeadlinesResponse](docs/FailedLatestHeadlinesResponse.md)
 - [From](docs/From.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [JournalistsProperty](docs/JournalistsProperty.md)
 - [LangProperty](docs/LangProperty.md)
 - [LatestHeadlinesGetResponse](docs/LatestHeadlinesGetResponse.md)
 - [LatestHeadlinesPostResponse](docs/LatestHeadlinesPostResponse.md)
 - [LatestHeadlinesRequest](docs/LatestHeadlinesRequest.md)
 - [LatestHeadlinesResponse](docs/LatestHeadlinesResponse.md)
 - [LocationPropertyInner](docs/LocationPropertyInner.md)
 - [MoreLikeThisRequest](docs/MoreLikeThisRequest.md)
 - [SearchGetResponse](docs/SearchGetResponse.md)
 - [SearchPostResponse](docs/SearchPostResponse.md)
 - [SearchRequest](docs/SearchRequest.md)
 - [SearchSimilarGetResponse](docs/SearchSimilarGetResponse.md)
 - [SearchSimilarPostResponse](docs/SearchSimilarPostResponse.md)
 - [SimilarDocument](docs/SimilarDocument.md)
 - [SourceResponse](docs/SourceResponse.md)
 - [SourcesRequest](docs/SourcesRequest.md)
 - [SubscriptionResponse](docs/SubscriptionResponse.md)
 - [To](docs/To.md)
 - [UserInput](docs/UserInput.md)
 - [ValidationError](docs/ValidationError.md)
