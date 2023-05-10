/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The virtual network
type Network struct {
	// Name of virtual network
	Name string `json:"name"`
	Type_ string `json:"type"`
	Dhcp string `json:"dhcp"`
	Subnet string `json:"subnet"`
	Mask string `json:"mask"`
}
