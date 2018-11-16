package converter

import "github.com/ligato/networkservicemesh/controlplane/pkg/apis/local/connection"

type LocalConnectionConverter struct {
	*connection.Connection
	name         string
	ipAddressKey string
}

func NewLocalConnectionConverter(c *connection.Connection, name string, ipAddressKey string) Converter {
	if c.GetMechanism().GetType() == connection.MechanismType_KERNEL_INTERFACE {
		return NewKernelConnectionConverter(c, name, ipAddressKey)
	}
	return nil
}