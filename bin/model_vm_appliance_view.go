/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VmApplianceView struct {
	Author string `json:"author,omitempty"`
	Version string `json:"version,omitempty"`
	Port *Port `json:"port,omitempty"`
	ShowAtPowerOn string `json:"showAtPowerOn,omitempty"`
}
