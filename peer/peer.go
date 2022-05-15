package peer

import uuid "github.com/satori/go.uuid"

var ct = 0

type Peer struct {
	Nr int
	Id uuid.UUID
	PeerMap map[uuid.UUID]int
}

func New() Peer {
	p := Peer{
		ct,
		uuid.Must(uuid.NewV4(), nil),
		make(map[uuid.UUID]int),
	}
	ct++
	return p
}

func (p Peer) SetPeersToPeerMap(distances map[uuid.UUID]int) {
	for key, value := range distances {
		p.PeerMap[key] = value
	}
}
