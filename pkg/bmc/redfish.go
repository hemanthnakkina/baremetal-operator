package bmc

import (
        "strings"
)

type redfishAccessDetails struct {
        bmcType  string
        portNum  string
        hostname string
        path     string
}

func (a *redfishAccessDetails) Type() string {
        return a.bmcType
}

// NeedsMAC returns true when the host is going to need a separate
// port created rather than having it discovered.
func (a *redfishAccessDetails) NeedsMAC() bool {
        return false
}

func (a *redfishAccessDetails) Driver() string {
        return "redfish"
}

// DriverInfo returns a data structure to pass as the DriverInfo
// parameter when creating a node in Ironic. The structure is
// pre-populated with the access information, and the caller is
// expected to add any other information that might be needed (such as
// the kernel and ramdisk locations).
func (a *redfishAccessDetails) DriverInfo(bmcCreds Credentials) map[string]interface{} {
        // TODO(nh863p): Hardcoded redfish_verify_ca, redfish_auth_type
        result := map[string]interface{}{
                "redfish_username": bmcCreds.Username,
                "redfish_password": bmcCreds.Password,
                "redfish_verify_ca": false,
                "redfish_system_id": a.path,
                "redfish_auth_type": "basic",
        }

        schemes := strings.Split(a.bmcType, "+")
        if len(schemes) > 1 {
                var redfish_address string
                redfish_address = schemes[1] + "://" + a.hostname
                if a.portNum != "" {
                        redfish_address = redfish_address + ":" + a.portNum
                }
                result["redfish_address"] = redfish_address
        }

        return result
}

func (a *redfishAccessDetails) BootInterface() string {
        return "fake"
}

func (a *redfishAccessDetails) InspectInterface() string {
        return "redfish"
}
