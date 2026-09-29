package uuid

import (
	"encoding/binary"
	"fmt"
	"os"
)

// um domínio representa um domínio de versão 2
type Domain byte

// constantes de domínio para uuids de dce de segurança (versão 2)
const (
	Person = Domain(0)
	Group  = Domain(1)
	Org    = Domain(2)
)

func NewDCESecurity(domain Domain, id uint32) (UUID, error) {
	uuid, err := NewUUID()

	if err == nil {
		uuid[6] = (uuid[6] & 0x0f) | 0x20 // versão 2
		uuid[9] = byte(domain)

		binary.BigEndian.PutUint32(uuid[0:], id)
	}

	return uuid, err
}

func NewDCEPerson() (UUID, error) {
	return NewDCESecurity(Person, uint32(os.Getuid()))
}

func NewDCEGroup() (UUID, error) {
	return NewDCESecurity(Group, uint32(os.Getgid()))
}

func (uuid UUID) Domain() Domain {
	return Domain(uuid[9])
}

func (uuid UUID) ID() uint32 {
	return binary.BigEndian.Uint32(uuid[0:4])
}

func (d Domain) String() string {
	switch d {
		case Person:
			return "Person"
		
		case Group:
			return "Group"
		
		case Org:
			return "Org"
	}
	
	return fmt.Sprintf("Domain%d", int(d))
}
