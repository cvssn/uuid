package uuid

import (
	"encoding/binary"
)

// newuuid retorna um uuid versão 1 baseado no nodeid atual e
// sequência clock, e tempo atual. se o nodeid não foi setado
// pelo setnodeid ou setnodeinterface, isso será então
// definido automaticamente. se o nodeid não pode ser
// definido, newuuid retornará nil. se a sequência clock não
// foi definida por setclocksequence então será definido
// automaticamente. se gettime falhar ao retornar o newuuid
// atual, retorna nil e um erro
//
// na maioria dos casos, "new" deve ser utilizado
func NewUUID() (UUID, error) {
	var uuid UUID
	now, seq, err := GetTime()

	if err != nil {
		return uuid, err
	}

	timeLow := uint32(now & 0xffffffff)
	timeMid := uint16((now >> 32) & 0xffff)
	timeHi := uint16((now >> 48) & 0x0fff)
	timeHi |= 0x1000 // versão 1

	binary.BigEndian.PutUint32(uuid[0:], timeLow)
	binary.BigEndian.PutUint16(uuid[4:], timeMid)
	binary.BigEndian.PutUint16(uuid[6:], timeHi)
	binary.BigEndian.PutUint16(uuid[8:], seq)

	nodeMu.Lock()

	if nodeID == zeroID {
		setNodeInterface("")
	}

	copy(uuid[10:], nodeID[:])

	nodeMu.Unlock()

	return uuid, nil
}
