# \OpenapiAppsApi

All URIs are relative to *https://cloud.goodrain.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1TeamsRegionsAppsCopyCreate**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsCopyCreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
[**OpenapiV1TeamsRegionsAppsCopyList**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsCopyList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/copy | 
[**OpenapiV1TeamsRegionsAppsCreate**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsCreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps | 
[**OpenapiV1TeamsRegionsAppsHttpdomainsCreate**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsHttpdomainsCreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/httpdomains | 
[**OpenapiV1TeamsRegionsAppsList**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps | 
[**OpenapiV1TeamsRegionsAppsOperationsCreate**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsOperationsCreate) | **Post** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/operations | 
[**OpenapiV1TeamsRegionsAppsRead**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsRead) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id} | 
[**OpenapiV1TeamsRegionsAppsServicesList**](OpenapiAppsApi.md#OpenapiV1TeamsRegionsAppsServicesList) | **Get** /openapi/v1/teams/{team_id}/regions/{region_name}/apps/{app_id}/services | 


# **OpenapiV1TeamsRegionsAppsCopyCreate**
> AppCopyCRes OpenapiV1TeamsRegionsAppsCopyCreate(ctx, appId, regionName, teamId, data)


复制应用

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**AppCopyC**](AppCopyC.md)|  | 

### Return type

[**AppCopyCRes**](AppCopyCRes.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsCopyList**
> []AppCopyL OpenapiV1TeamsRegionsAppsCopyList(ctx, appId, regionName, teamId)


获取需要复制的应用组件信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**[]AppCopyL**](AppCopyL.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsCreate**
> AppBaseInfo OpenapiV1TeamsRegionsAppsCreate(ctx, regionName, teamId, data)


创建应用

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**AppPostInfo**](AppPostInfo.md)|  | 

### Return type

[**AppBaseInfo**](AppBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsHttpdomainsCreate**
> HttpGatewayRule OpenapiV1TeamsRegionsAppsHttpdomainsCreate(ctx, appId, regionName, teamId, data)


创建HTTP网关策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**PostHttpGatewayRule**](PostHttpGatewayRule.md)|  | 

### Return type

[**HttpGatewayRule**](HTTPGatewayRule.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsList**
> []AppBaseInfo OpenapiV1TeamsRegionsAppsList(ctx, regionName, teamId, optional)


团队应用列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 
 **optional** | ***OpenapiAppsApiOpenapiV1TeamsRegionsAppsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiAppsApiOpenapiV1TeamsRegionsAppsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**| 搜索查询应用名称，团队名称 | 

### Return type

[**[]AppBaseInfo**](AppBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsOperationsCreate**
> Success OpenapiV1TeamsRegionsAppsOperationsCreate(ctx, appId, regionName, teamId, data)


操作应用

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**ServiceGroupOperations**](ServiceGroupOperations.md)|  | 

### Return type

[**Success**](Success.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsRead**
> AppInfo OpenapiV1TeamsRegionsAppsRead(ctx, appId, regionName, teamId)


应用详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**AppInfo**](AppInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsRegionsAppsServicesList**
> []ServiceBaseInfo OpenapiV1TeamsRegionsAppsServicesList(ctx, appId, regionName, teamId)


查询应用下组件列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **regionName** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**[]ServiceBaseInfo**](ServiceBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

