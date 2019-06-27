package p2p

import (
	"github.com/boltdb/bolt"
)

// #[derive(Clone)]
// pub struct P2pClient {
//     listener_port: u16,
//     init_chain_id: Vec<u8>,
//     identity: Identity,
//     versions: Vec<Version>,
//     db: Arc<RwLock<Db>>,
// }

// Client represents a p2p client
type Client struct {
	ListenPort  uint16
	InitChainID []byte
	Identity    Identity
	Versions    []interface{} //TODO
	DB          *bolt.DB
}

type ConnectionMessage struct {
	ListenPort       uint16
	PublicKey        string
	ProofOfWorkStamp string
	Nonce            []byte
	Versions         []interface{}
}

// NewClient returns a new p2p client
func NewClient(identity Identity, versions []interface{}, DB *bolt.DB) *Client {
	return &Client{ListenPort: 8732, InitChainID: getGenesisChainID(), Identity: identity, Versions: versions, DB: DB}
}

func (c *Client) prepareConnectionMessage() ConnectionMessage {

	return ConnectionMessage{
		ListenPort:       c.ListenPort,
		PublicKey:        c.Identity.PublicKey,
		ProofOfWorkStamp: c.Identity.ProofOfWorkStamp,
		Nonce:            []byte{},
		Versions:         nil,
	}
}

// fn prepare_connection_message(&self) -> ConnectionMessage {
// 	// generate init random nonce
// 	let nonce = Nonce::random();
// 	let connection_message = ConnectionMessage::new(
// 		self.listener_port,
// 		&self.identity.public_key,
// 		&self.identity.proof_of_work_stamp,
// 		&nonce.get_bytes(),
// 		self.versions.iter().map(|v| v.into()).collect()
// 	);
// 	connection_message
// }

func getGenesisChainID() []byte {
	return []byte("8eceda2f")
}
