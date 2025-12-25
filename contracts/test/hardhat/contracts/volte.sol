    // SPDX-License-Identifier: GPL-3.0
    pragma solidity >=0.4.16 <0.9.0;

    import "./groth16/ballot.sol" as Ballot;
    import "./groth16/nullifier.sol" as Nullifier;
    import "./groth16/membership.sol" as Membership;

    contract VolteContract{

        uint256 constant P = 21888242871839275222246405745257275088548364400416034343698204186575808495617;

        uint256 constant LIMB_SIZE = 64;

        struct TallyScore {
            uint256 C1x;
            uint256 C1y;
            uint256 C2x;
            uint256 C2y;
        }

        struct Proof {
            uint256 Arx;
            uint256 Ary;
            uint256 Brx0;
            uint256 Brx1;
            uint256 Bry0;
            uint256 Bry1;
            uint256 Cx;
            uint256 Cy;
        }

        struct BallotProof {
            Proof       Proof;
            // The public inputs to the circuit.
            // [C1x0, C1x1, C1x2, C1y3, C1y0, C1y1, C1y2, C1y3, C2x0, C2x1, C2x2, C2y3, C2y0, C2y1, C2y2, C2y3]
            uint256[16] Input;
            uint256     CommitmentX;
            uint256     CommitmentY;
            uint256     CommitmentPokX;
            uint256     CommitmentPokY;
        }

        struct MembershipProof {
            Proof      Proof;
            // The public inputs to the circuit.
            // [merkleRoot, leafValue]
            uint256[2] Input;
        }

        struct NullifierProof {
            Proof     Proof;
            // The public inputs to the circuit.
            // [eventID, nullifier]
            uint256[3] Input;
        }

        struct Proofs {
            BallotProof     ballot;
            MembershipProof membership;
            NullifierProof  nullifier;
        }

        struct VoteSubmission {
            string  eventID;
            Proofs  proofs;
        }

        address public admin;

        Ballot.BallotVerifier public ballot;
        Nullifier.NullifierVerifier public nullifier;
        Membership.MembershipVerifier public membership;

        constructor(
            address _ballot,
            address _membership,
            address _nullifier
        ) {
            admin = msg.sender;
            ballot = Ballot.BallotVerifier(_ballot);
            nullifier = Nullifier.NullifierVerifier(_nullifier);
            membership = Membership.MembershipVerifier(_membership);
        }

        mapping (string /* eventID */ => bytes /* NullifierRootHash */) public membershipMerkleRoots;
        mapping (string /* eventID */ => bytes /* VoteRootHash */)      public voteMerkleRoots;
        mapping (string /* eventID */ => bytes /* EventDetailsHash */)  public eventHashes;
        mapping (uint256 /* eventID */ => TallyScore) public tallyScores;

        modifier onlyOwner() {
            require(msg.sender == admin, "Only owner is allowed to execute this transaction.");
            _;
        }

        function SetVoteMerkleRoot(string calldata eventID, bytes calldata value) external {
            // check before to make sure its a valid hash
            voteMerkleRoots[eventID] = value;
        }

        function SetEventHash(string calldata eventID, bytes calldata value) external {
            // A change in event details changes the event hash completely.
            eventHashes[eventID] = value;
        }

        function GetVoteMerkleRoot(string calldata eventID) external view returns (bytes memory) {
            return voteMerkleRoots[eventID];
        }

        function GetEventHash(string calldata eventID) external view returns (bytes memory) {
            return eventHashes[eventID];
        }

        function GetTallyScore(uint256 eventID) external view returns (uint256[4] memory){
            TallyScore storage score = tallyScores[eventID];
            return [score.C1x, score.C1y, score.C2x, score.C2y];
        }

        function addCiphertexts(
            uint256[2] memory C1, // [x1, y1]
            uint256[2] memory C2  // [x2, y2]
        ) public view returns (uint256[2] memory Csum) {
            bool success;
            assembly ("memory-safe") {
                let ptr := mload(0x40)

                mstore(ptr,         mload(C1))             // x1
                mstore(add(ptr,32), mload(add(C1,32)))    // y1
                mstore(add(ptr,64), mload(C2))            // x2
                mstore(add(ptr,96), mload(add(C2,32)))    // y2

                success := staticcall(gas(), 0x06, ptr, 0x80, ptr, 0x40)

                let csumPtr := Csum
                mstore(csumPtr,        mload(ptr))          // x3
                mstore(add(csumPtr,32), mload(add(ptr,32))) // y3
                mstore(0x40, add(ptr, 0x80))
            }

            require(success, "ECADD failed");
            return Csum;
        }

        // No revert means the proof has been verified and the vote has been submitted.
        function Vote(VoteSubmission calldata proof) external{
            nullifier.verifyProof(
                [
                    proof.proofs.nullifier.Proof.Arx,
                    proof.proofs.nullifier.Proof.Ary,
                    // Notice the ordering of B point as the verifier expects them in big-endian order.
                    proof.proofs.nullifier.Proof.Brx1,
                    proof.proofs.nullifier.Proof.Brx0,
                    proof.proofs.nullifier.Proof.Bry1,
                    proof.proofs.nullifier.Proof.Bry0,
                    proof.proofs.nullifier.Proof.Cx,
                    proof.proofs.nullifier.Proof.Cy
                ],
                proof.proofs.nullifier.Input
            );
            membership.verifyProof(
                [
                    proof.proofs.membership.Proof.Arx,
                    proof.proofs.membership.Proof.Ary,
                    // Notice the ordering of B point as the verifier expects them in big-endian order.
                    proof.proofs.membership.Proof.Brx1,
                    proof.proofs.membership.Proof.Brx0,
                    proof.proofs.membership.Proof.Bry1,
                    proof.proofs.membership.Proof.Bry0,
                    proof.proofs.membership.Proof.Cx,
                    proof.proofs.membership.Proof.Cy
                ],
                proof.proofs.membership.Input
            );

            ballot.verifyProof(
                [
                    proof.proofs.ballot.Proof.Arx,
                    proof.proofs.ballot.Proof.Ary,
                    // Notice the ordering of B point as the verifier expects them in big-endian order.
                    proof.proofs.ballot.Proof.Brx1,
                    proof.proofs.ballot.Proof.Brx0,
                    proof.proofs.ballot.Proof.Bry1,
                    proof.proofs.ballot.Proof.Bry0,
                    proof.proofs.ballot.Proof.Cx,
                    proof.proofs.ballot.Proof.Cy
                ],
                [proof.proofs.ballot.CommitmentX, proof.proofs.ballot.CommitmentY],
                [proof.proofs.ballot.CommitmentPokX, proof.proofs.ballot.CommitmentPokY],
                proof.proofs.ballot.Input
            );
            // Recombining into 256bit format for each coordination.
            uint256 C1x = recombine([
                uint64(proof.proofs.ballot.Input[0]),
                uint64(proof.proofs.ballot.Input[1]),
                uint64(proof.proofs.ballot.Input[2]),
                uint64(proof.proofs.ballot.Input[3])
            ]);
            uint256 C1y = recombine([
                uint64(proof.proofs.ballot.Input[4]),
                uint64(proof.proofs.ballot.Input[5]),
                uint64(proof.proofs.ballot.Input[6]),
                uint64(proof.proofs.ballot.Input[7])
            ]);
            uint256 C2x = recombine([
                uint64(proof.proofs.ballot.Input[8]),
                uint64(proof.proofs.ballot.Input[9]),
                uint64(proof.proofs.ballot.Input[10]),
                uint64(proof.proofs.ballot.Input[11])
            ]);
            uint256 C2y = recombine([
                uint64(proof.proofs.ballot.Input[12]),
                uint64(proof.proofs.ballot.Input[13]),
                uint64(proof.proofs.ballot.Input[14]),
                uint64(proof.proofs.ballot.Input[15])
            ]);

            // Preparing input parameters to perform tally + C.
            uint256 eventID = proof.proofs.nullifier.Input[0];
            TallyScore storage tallyScore = tallyScores[eventID];

            uint256[2] memory C1 = [C1x, C1y];
            uint256[2] memory C2 = [C2x, C2y];
            uint256[2] memory tallyC1 = [tallyScore.C1x, tallyScore.C1y];
            uint256[2] memory tallyC2 = [tallyScore.C2x, tallyScore.C2y];
            uint256[2] memory resultC1 = addCiphertexts(C1, tallyC1);
            uint256[2] memory resultC2 = addCiphertexts(C2, tallyC2);

            // Update tally score.
            tallyScore.C1x = resultC1[0];
            tallyScore.C1y = resultC1[1];
            tallyScore.C2x = resultC2[0];
            tallyScore.C2y = resultC2[1];
        }

        function recombine(uint64[4] memory limbs) internal pure returns (uint256 x) {

            x = uint256(limbs[0])
                | (uint256(limbs[1]) << 64)
                | (uint256(limbs[2]) << 128)
                | (uint256(limbs[3]) << 192);

            // reduce modulo P to ensure valid field element
            x = addmod(x, 0, P);
            return x;
        }

    }