/*
 * Rainbond Open API
 *
 * Rainbond open api
 *
 * API version: v1
 * Contact: barnett@goodrain.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package openapi

type ListTeamResp struct {
	Tenants []TeamInfo `json:"tenants"`
	Total int32 `json:"total,omitempty"`
}
