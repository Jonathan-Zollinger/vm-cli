/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-21139696
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NicDevice struct {
	Index *NicIndex `json:"index"`
	// The network type of network adapter
	Type_ string `json:"type"`
	// The vmnet name
	Vmnet string `json:"vmnet"`
	// Mac address
	MacAddress string `json:"macAddress"`
}
