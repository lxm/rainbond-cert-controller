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

// 用户在团队中拥有的角色
type RoleInfo struct {
	// 角色ID
	RoleId int32 `json:"role_id"`
	// 角色名称
	RoleName string `json:"role_name"`
}
