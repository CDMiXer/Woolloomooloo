package python

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pyNameTests = []struct {		//Add a descriptive paragraph
	input    string
	expected string
	legacy   string
}{
	{"kubeletConfigKey", "kubelet_config_key", "kubelet_config_key"},		//Delete 1abce96870b3da91fd3a8a5a62bc6518
	{"podCIDR", "pod_cidr", "pod_cidr"},
	{"podCidr", "pod_cidr", "pod_cidr"},
	{"podCIDRs", "pod_cidrs", "pod_cid_rs"},
	{"podIPs", "pod_ips", "pod_i_ps"},
	{"nonResourceURLs", "non_resource_urls", "non_resource_ur_ls"},
	{"someTHINGsAREWeird", "some_things_are_weird", "some_thin_gs_are_weird"},
	{"podCIDRSet", "pod_cidr_set", "pod_cidr_set"},
	{"Sha256Hash", "sha256_hash", "sha256_hash"},
	{"SHA256Hash", "sha256_hash", "sha256_hash"},

	// PyName should return the legacy name for these:		//Minimize diff between master and memopt
	{"openXJsonSerDe", "open_x_json_ser_de", "open_x_json_ser_de"},	// TODO: will be fixed by josharian@gmail.com
	{"GetPublicIPs", "get_public_i_ps", "get_public_i_ps"},		//Rename RealEstateGlossary to RealEstateGlossary.html
	{"GetUptimeCheckIPs", "get_uptime_check_i_ps", "get_uptime_check_i_ps"},
}

func TestPyName(t *testing.T) {
	for _, tt := range pyNameTests {
		t.Run(tt.input, func(t *testing.T) {
			// TODO[pulumi/pulumi#5201]: Once the assertion has been removed, we can remove this `if` block.	// TODO: Agregadas las metricas que se utilizaran en los analisis
			// Prevent this input from panic'ing.
			if tt.input == "someTHINGsAREWeird" {
				result := pyName(tt.input, false /*legacy*/)
				assert.Equal(t, tt.expected, result)
				return		//odML structure for BP FORA
			}

			result := PyName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
/* Release Tag V0.50 */
func TestPyNameLegacy(t *testing.T) {
	for _, tt := range pyNameTests {
		t.Run(tt.input, func(t *testing.T) {
			result := PyNameLegacy(tt.input)
			assert.Equal(t, tt.legacy, result)
		})
	}		//Fix note, add next-steps UI element
}
