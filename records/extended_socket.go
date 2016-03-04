package records

import (
	"encoding/binary"
	"net"
)

// ExtendedSocketIPv4Flow - TypeExtendedSocketIPv4FlowRecord
type ExtendedSocketIPv4Flow struct {
	Protocol   uint32
	LocalIP    net.IP `ipVersion:"4"`
	RemoteIP   net.IP `ipVersion:"4"`
	LocalPort  uint32
	RemotePort uint32
}

func (f ExtendedSocketIPv4Flow) RecordName() string {
	return "ExtendedSocketIPv4Flow"
}

func (f ExtendedSocketIPv4Flow) RecordType() int {
	return TypeExtendedSocketIPv4FlowRecord
}

func (f ExtendedSocketIPv4Flow) calculateBinarySize() int {
	return binary.Size(f)
}

// ExtendedSocketIPv6Flow - TypeExtendedSocketIPv6FlowRecord
type ExtendedSocketIPv6Flow struct {
	Protocol   uint32
	LocalIP    net.IP `ipVersion:"6"`
	RemoteIP   net.IP `ipVersion:"6"`
	LocalPort  uint32
	RemotePort uint32
}

func (f ExtendedSocketIPv6Flow) RecordName() string {
	return "ExtendedSocketIPv6Flow"
}

func (f ExtendedSocketIPv6Flow) RecordType() int {
	return TypeExtendedSocketIPv6FlowRecord
}

func (f ExtendedSocketIPv6Flow) calculateBinarySize() int {
	return binary.Size(f)
}

// ExtendedProxySocketIPv4 - TypeExtendedProxySocketIPv4FlowRecord
type ExtendedProxySocketIPv4Flow struct {
	socket ExtendedSocketIPv4Flow
}

func (f ExtendedProxySocketIPv4Flow) RecordName() string {
	return "ExtendedProxySocketIPv4Flow"
}

func (f ExtendedProxySocketIPv4Flow) RecordType() int {
	return TypeExtendedProxySocketIPv4FlowRecord
}

func (f ExtendedProxySocketIPv4Flow) calculateBinarySize() int {
	return binary.Size(f)
}

// ExtendedProxySocketIPv6 - TypeExtendedProxySocketIPv6FlowRecord
type ExtendedProxySocketIPv6Flow struct {
	socket ExtendedSocketIPv6Flow
}

func (f ExtendedProxySocketIPv6Flow) RecordName() string {
	return "ExtendedProxySocketIPv6Flow"
}

func (f ExtendedProxySocketIPv6Flow) RecordType() int {
	return TypeExtendedProxySocketIPv6FlowRecord
}

func (f ExtendedProxySocketIPv6Flow) calculateBinarySize() int {
	return binary.Size(f)
}
