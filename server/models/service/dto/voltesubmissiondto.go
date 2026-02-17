package dto

import (
	"encoding/json"
	"fmt"
	"math/big"
	"volte/backend/chain/contracts"
)

type BigInt big.Int

func (b *BigInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		// leave as zero value; caller can treat separately if needed
		*b = BigInt{}
		return nil
	}

	// string?
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		i, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return fmt.Errorf("invalid big int string: %q", s)
		}
		*b = BigInt(*i)
		return nil
	}

	// number?
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		i, ok := new(big.Int).SetString(n.String(), 10)
		if !ok {
			return fmt.Errorf("invalid big int number: %q", n.String())
		}
		*b = BigInt(*i)
		return nil
	}

	return fmt.Errorf("bigint must be string or number, got: %s", string(data))
}

func (b *BigInt) ToBig() *big.Int {
	if b == nil {
		return nil
	}
	// convert value -> *big.Int
	i := big.Int(*b)
	return &i
}

type VolteContractProofDTO struct {
	Arx  BigInt `json:"Arx"`
	Ary  BigInt `json:"Ary"`
	Brx0 BigInt `json:"Brx0"`
	Brx1 BigInt `json:"Brx1"`
	Bry0 BigInt `json:"Bry0"`
	Bry1 BigInt `json:"Bry1"`
	Cx   BigInt `json:"Cx"`
	Cy   BigInt `json:"Cy"`
}

func ConvertVolteContractProofDTO(dto VolteContractProofDTO) contracts.VolteContractProof {
	return contracts.VolteContractProof{
		Arx:  dto.Arx.ToBig(),
		Ary:  dto.Ary.ToBig(),
		Brx0: dto.Brx0.ToBig(),
		Brx1: dto.Brx1.ToBig(),
		Bry0: dto.Bry0.ToBig(),
		Bry1: dto.Bry1.ToBig(),
		Cx:   dto.Cx.ToBig(),
		Cy:   dto.Cy.ToBig(),
	}
}

type VolteContractBallotProofDTO struct {
	Proof VolteContractProofDTO `json:"Proof"`
	Input [4]BigInt             `json:"Input"`
}

func ConvertVolteContractBallotProofDTO(dto VolteContractBallotProofDTO) contracts.VolteContractBallotProof {
	proof := contracts.VolteContractBallotProof{
		Proof: ConvertVolteContractProofDTO(dto.Proof),
	}
	for i, in := range dto.Input {
		proof.Input[i] = in.ToBig()
	}
	return proof
}

type VolteContractMembershipProofDTO struct {
	Proof VolteContractProofDTO `json:"Proof"`
	Input [2]BigInt             `json:"Input"`
}

func ConvertVolteContractMembershipProofDTO(dto VolteContractMembershipProofDTO) contracts.VolteContractMembershipProof {
	proof := contracts.VolteContractMembershipProof{
		Proof: ConvertVolteContractProofDTO(dto.Proof),
	}
	for i, in := range dto.Input {
		proof.Input[i] = in.ToBig()
	}
	return proof
}

type VolteContractNullifierProofDTO struct {
	Proof VolteContractProofDTO `json:"Proof"`
	Input [3]BigInt             `json:"Input"`
}

func ConvertVolteContractNullifierProofDTO(dto VolteContractNullifierProofDTO) contracts.VolteContractNullifierProof {
	proof := contracts.VolteContractNullifierProof{
		Proof: ConvertVolteContractProofDTO(dto.Proof),
	}
	for i, in := range dto.Input {
		proof.Input[i] = in.ToBig()
	}
	return proof
}

type VolteContractProofsDTO struct {
	Ballot     VolteContractBallotProofDTO     `json:"Ballot"`
	Membership VolteContractMembershipProofDTO `json:"Membership"`
	Nullifier  VolteContractNullifierProofDTO  `json:"Nullifier"`
}

func ConvertVolteContractProofsDTO(dto VolteContractProofsDTO) contracts.VolteContractProofs {
	return contracts.VolteContractProofs{
		Ballot:     ConvertVolteContractBallotProofDTO(dto.Ballot),
		Nullifier:  ConvertVolteContractNullifierProofDTO(dto.Nullifier),
		Membership: ConvertVolteContractMembershipProofDTO(dto.Membership),
	}
}

// VolteContractVoteSubmissionDTO facilitates the data validation of vote.
// Can't directly use bigint in json because it would automatically reformat the value in
// exponential notation; therefore ruins the data valdation.
type VolteContractVoteSubmissionDTO struct {
	EventID string                 `json:"EventID"`
	Proofs  VolteContractProofsDTO `json:"Proofs"`
}

func ConvertVolteContractVoteSubmissionDTO(dto VolteContractVoteSubmissionDTO) *contracts.VolteContractVoteSubmission {
	return &contracts.VolteContractVoteSubmission{
		EventID: dto.EventID,
		Proofs:  ConvertVolteContractProofsDTO(dto.Proofs),
	}
}
