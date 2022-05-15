package broadcast

import (
	"fmt"

	"github.com/minhkhiemm/ring/peer"
	uuid "github.com/satori/go.uuid"
)

var peers []peer.Peer

func Broadcast(nodes int, latency int, peersPerNode int, bandwidth int, blocksize int) (int, int) {
	for i := 0; i < nodes; i++ {
		peers = append(peers, peer.New())
	}

	for i := range peers {
		computeDistance(peers[i], peersPerNode)
	}
	SendMessage(peers[0], nodes)
	return 10, 20
}

func SendMessage(user peer.Peer, nodes int) {
	visited := make([]int, nodes)
	var queue []peer.Peer

	visited[user.Nr] = 1

	queue = append(queue, user)

	for len(queue) != 0 {
		var neighbor peer.Peer

		node := queue[0]
		queue = append(queue[:0], queue[1:]...)

		for key := range node.PeerMap {
			for _, user := range peers {
				if user.Id == key {
					neighbor = user
				}
			}

			if visited[neighbor.Nr] == 0 {
				visited[neighbor.Nr] = 1

				queue = append(queue, neighbor)
			}
		}
	}
}

func computeDistance(node peer.Peer, peersPerNode int) {
	distanceToPeers := make([]int, len(peers))
	distanceMap := make(map[uuid.UUID]int)
	var nodeID []byte = (node.Id).Bytes()
	var peerID []byte

	for i, user := range peers {
		peerID = (user.Id).Bytes()

		XOR, err := XORBytes(nodeID, peerID)
		if err != nil {
			fmt.Println(err)
		}

		distanceToPeers[i] = CountSetBits(XOR)
	}

	for i := 0; i < peersPerNode; i++ {
		min := 200
		indexForMin := 0
		for j := 0; j < len(distanceToPeers); j++ {
			if distanceToPeers[j] < min && distanceToPeers[j] != 0 {
				min = distanceToPeers[j]
				indexForMin = j
			}
		}
		distanceMap[peers[indexForMin].Id] = min
		distanceToPeers[indexForMin] = 0
	}

	node.SetPeersToPeerMap(distanceMap)
}

func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of byte slices is not equivalent: %d != %d", len(a), len(b))
	}

	buf := make([]byte, len(a))

	for i := range a {
		buf[i] = a[i] ^ b[i]
	}

	return buf, nil
}

func CountSetBits(nr []byte) int {
	count := 0

	for _, byteNo := range nr {
		for i := 0; i < 8; i++ {
			if byteNo&1 == 1 {
				count++
			}
		}
		byteNo >>= 1
	}
	return count
}
