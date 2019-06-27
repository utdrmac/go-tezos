package p2p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/pkg/errors"
)

// Identity represents a node identity
type Identity struct {
	PeerID           string `json:"peer_id"`
	PublicKey        string `json:"public_key"`
	SecretKey        string `json:"secret_key"`
	ProofOfWorkStamp string `json:"proof_of_work_stamp"`
}

// NewIdentity generates a new identity with the set difficulty
func NewIdentity(difficulty int) (*Identity, error) {
	return nil, nil
}

// GetIdentity reads in the identity of the tezos node from the specified path or in
// the default $HOME/.tezos-node directory.
func GetIdentity(nodedir string) (Identity, error) {

	var file string
	if _, err := os.Stat(nodedir); !os.IsNotExist(err) {
		file = fmt.Sprintf("%s/identity.json", nodedir)
	} else {
		usr, err := user.Current()
		if err != nil {
			return Identity{}, errors.Wrap(err, "could not get identity, home directory not found")
		}

		file = fmt.Sprintf("%s/.tezos-node/identity.json", usr.HomeDir)
	}

	fbyts, err := ioutil.ReadFile(file)
	if err != nil {
		return Identity{}, errors.Wrap(err, "could not get identity")
	}

	var identity Identity
	identity, err = identity.unmarshalJSON(fbyts)
	if err != nil {
		return Identity{}, errors.Wrap(err, "could not get identity")
	}

	return identity, nil
}

func (i *Identity) unmarshalJSON(v []byte) (Identity, error) {
	var identity Identity
	err := json.Unmarshal(v, &identity)
	if err != nil {
		return identity, errors.Wrap(err, "could not unmarshal identity")
	}

	return identity, nil
}
