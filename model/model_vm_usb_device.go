/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VmUsbDevice struct {
	Index *Number `json:"index,omitempty"`
	Connected string `json:"connected,omitempty"`
	BackingInfo string `json:"backingInfo,omitempty"`
	BackingType *Number `json:"BackingType,omitempty"`
}
