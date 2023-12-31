package server

import (
	"github.com/sigmachain/sigmachain/consensus"
	consensusDev "github.com/sigmachain/sigmachain/consensus/dev"
	consensusDummy "github.com/sigmachain/sigmachain/consensus/dummy"
	consensusIBFT "github.com/sigmachain/sigmachain/consensus/ibft"
	"github.com/sigmachain/sigmachain/secrets"
	"github.com/sigmachain/sigmachain/secrets/awsssm"
	"github.com/sigmachain/sigmachain/secrets/gcpssm"
	"github.com/sigmachain/sigmachain/secrets/hashicorpvault"
	"github.com/sigmachain/sigmachain/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
