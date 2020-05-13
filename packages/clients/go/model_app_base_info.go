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

type AppBaseInfo struct {
	ID int32 `json:"ID,omitempty"`
	// 组名
	GroupName string `json:"group_name"`
	// 默认组件
	IsDefault bool `json:"is_default,omitempty"`
	// 备注
	Note string `json:"note,omitempty"`
	// 应用排序
	OrderIndex int32 `json:"order_index,omitempty"`
	// 区域中心名称
	RegionName string `json:"region_name"`
	// 租户id
	TenantId string `json:"tenant_id"`
}