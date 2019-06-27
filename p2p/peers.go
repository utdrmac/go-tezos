package p2p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/pkg/errors"
)

//PeerURL is a helper struct to represent a peer's URL
type PeerURL struct {
	URL  string
	Port string
}

// Peers represents peers.json
type Peers []Peer

// Peer represents a single peer in peers.json
type Peer struct {
	PeerID       string `json:"peer_id"`
	Created      string `json:"created"`
	PeerMetadata struct {
		Responses struct {
			Sent struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"sent"`
			Failed struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"failed"`
			Received struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"received"`
			Unexpected string `json:"unexpected"`
			Outdated   string `json:"outdated"`
		} `json:"responses"`
		Requests struct {
			Sent struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"sent"`
			Received struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"received"`
			Failed struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"failed"`
			Scheduled struct {
				Branch                  string `json:"branch"`
				Head                    string `json:"head"`
				BlockHeader             string `json:"block_header"`
				Operations              string `json:"operations"`
				Protocols               string `json:"protocols"`
				OperationHashesForBlock string `json:"operation_hashes_for_block"`
				OperationsForBlock      string `json:"operations_for_block"`
				Other                   string `json:"other"`
			} `json:"scheduled"`
		} `json:"requests"`
		ValidBlocks         string `json:"valid_blocks"`
		OldHeads            string `json:"old_heads"`
		PrevalidatorResults struct {
			CannotDownload      string `json:"cannot_download"`
			CannotParse         string `json:"cannot_parse"`
			RefusedByPrefilter  string `json:"refused_by_prefilter"`
			RefusedByPostfilter string `json:"refused_by_postfilter"`
			Applied             string `json:"applied"`
			BranchDelayed       string `json:"branch_delayed"`
			BranchRefused       string `json:"branch_refused"`
			Refused             string `json:"refused"`
			Duplicate           string `json:"duplicate"`
			Outdated            string `json:"outdated"`
		} `json:"prevalidator_results"`
		UnactivatedChains      string `json:"unactivated_chains"`
		InactiveChains         string `json:"inactive_chains"`
		FutureBlocksAdvertised string `json:"future_blocks_advertised"`
		Unadvertised           struct {
			Block      string `json:"block"`
			Operations string `json:"operations"`
			Protocol   string `json:"protocol"`
		} `json:"unadvertised"`
		Advertisements struct {
			Sent struct {
				Head   string `json:"head"`
				Branch string `json:"branch"`
			} `json:"sent"`
			Received struct {
				Head   string `json:"head"`
				Branch string `json:"branch"`
			} `json:"received"`
		} `json:"advertisements"`
	} `json:"peer_metadata"`
	LastRejectedConnection    []interface{} `json:"last_rejected_connection"`
	LastEstablishedConnection []interface{} `json:"last_established_connection"`
	LastDisconnection         []interface{} `json:"last_disconnection"`
}

// GetPeers reads in peers.json from the tezos node directory
func GetPeers(nodedir string) (*Peers, error) {
	var peers *Peers

	var file string
	if _, err := os.Stat(nodedir); !os.IsNotExist(err) {
		file = fmt.Sprintf("%s/peers.json", nodedir)
	} else {
		usr, err := user.Current()
		if err != nil {
			return nil, errors.Wrap(err, "could not get peers, home directory not found")
		}

		file = fmt.Sprintf("%s/.tezos-node/peers.json", usr.HomeDir)
	}

	fbyts, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "could not get peers")
	}

	peers, err = peers.unmarshalJSON(fbyts)
	if err != nil {
		return nil, errors.Wrap(err, "could not get peers")
	}

	return peers, nil
}

// GetPeerURLs gets all known peer URL's to Peers
func (p *Peers) GetPeerURLs() []PeerURL {
	var peerURLs []PeerURL
	for _, peer := range *p {
		connections := peer.LastEstablishedConnection
		if len(connections) > 0 {
			conByts, err := json.Marshal(connections[0])
			if err != nil {
				continue
			}
			var peerURL PeerURL
			err = json.Unmarshal(conByts, &peerURL)
			if err != nil {
				continue
			}

			peerURLs = append(peerURLs, peerURL)
		}
	}

	return peerURLs
}

func (p *Peers) unmarshalJSON(v []byte) (*Peers, error) {
	var peers *Peers
	err := json.Unmarshal(v, peers)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal peers")
	}

	return peers, nil
}
