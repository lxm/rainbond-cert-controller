# \OpenapiGatewayApi

All URIs are relative to *https://cloud.goodrain.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1TeamsRegionsAppsHttpdomainsList**](OpenapiGatewayApi.md#OpenapiV1TeamsRegionsAppsHttpdomainsList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 


# **OpenapiV1TeamsRegionsAppsHttpdomainsList**
> []HttpGatewayRule OpenapiV1TeamsRegionsAppsHttpdomainsList(ctx, appId, regionName, teamId)


获取应用http访问策略列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**[]HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

