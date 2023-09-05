// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package srv6map

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"

	srv6Types "github.com/cilium/cilium/enterprise/pkg/srv6/types"
	"github.com/cilium/cilium/pkg/ebpf"
	"github.com/cilium/cilium/pkg/types"
)

const (
	SIDMapName    = "cilium_srv6_sid"
	MaxSIDEntries = 16384
)

var (
	SRv6SIDMap *srv6SIDMap
)

type SIDKey struct {
	SID types.IPv6
}

func (a *SIDKey) Equal(b *SIDKey) bool {
	if (a != nil) != (b != nil) {
		return false
	}
	return a.SID.IP().Equal(b.SID.IP())
}

func (k *SIDKey) String() string {
	return fmt.Sprintf("%s", k.SID)
}

func NewSIDKey(sid types.IPv6) SIDKey {
	result := SIDKey{}
	result.SID = sid
	return result
}

func NewSIDKeyFromSID(sid *srv6Types.SID) (*SIDKey, error) {
	result := &SIDKey{}
	copy(result.SID[:], sid.Addr.AsSlice())
	return result, nil
}

type SIDValue struct {
	VRFID uint32
}

func (a *SIDValue) Equal(b *SIDValue) bool {
	if (a != nil) != (b != nil) {
		return false
	}
	return a.VRFID == b.VRFID
}

// srv6SIDMap is the internal representation of an SRv6 SID map.
type srv6SIDMap struct {
	*ebpf.Map
}

func initSIDMap(create bool) error {
	var m *ebpf.Map
	var err error

	if create {
		m = ebpf.NewMap(&ebpf.MapSpec{
			Name:       SIDMapName,
			Type:       ebpf.Hash,
			KeySize:    uint32(unsafe.Sizeof(SIDKey{})),
			ValueSize:  uint32(unsafe.Sizeof(SIDValue{})),
			MaxEntries: uint32(MaxSIDEntries),
			Flags:      unix.BPF_F_NO_PREALLOC,
			Pinning:    ebpf.PinByName,
		})
		if err = m.OpenOrCreate(); err != nil {
			return err
		}
	} else {
		if m, err = ebpf.LoadRegisterMap(SIDMapName); err != nil {
			return err
		}
	}

	SRv6SIDMap = &srv6SIDMap{
		m,
	}

	return nil
}

func CreateSIDMap() error {
	return initSIDMap(true)
}

func OpenSIDMap() error {
	return initSIDMap(false)
}

func (m *srv6SIDMap) Lookup(key SIDKey, val *SIDValue) error {
	return m.Map.Lookup(key, val)
}

func (m *srv6SIDMap) Update(key SIDKey, vrfID uint32) error {
	val := SIDValue{VRFID: vrfID}
	return m.Map.Update(key, val, 0)
}

func (m *srv6SIDMap) Delete(key SIDKey) error {
	return m.Map.Delete(key)
}

// SRv6SIDIterateCallback represents the signature of the callback function
// expected by the IterateWithCallback method, which in turn is used to iterate
// all the keys/values of an SRv6 SID map.
type SRv6SIDIterateCallback func(*SIDKey, *SIDValue)

// IterateWithCallback iterates through all the keys/values of an SRv6 SID
// map, passing each key/value pair to the cb callback.
func (m srv6SIDMap) IterateWithCallback(cb SRv6SIDIterateCallback) error {
	return m.Map.IterateWithCallback(&SIDKey{}, &SIDValue{},
		func(k, v interface{}) {
			key := k.(*SIDKey)
			value := v.(*SIDValue)

			cb(key, value)
		})
}
