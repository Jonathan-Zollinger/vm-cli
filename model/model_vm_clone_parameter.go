/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VmCloneParameter struct {
	// New VM name
	Name string `json:"name"`
	// Existing VM ID to clone.
	ParentId string `json:"parentId"`
}
