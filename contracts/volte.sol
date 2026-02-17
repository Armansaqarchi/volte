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
            uint256[4] Input;
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

        error NullifierProofInvalid();
        error MembershipProofInvalid();
        error BallotProofInvalid();

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

        mapping (string /* eventID */ => uint256 /* VoteRootHash */)      public voteMerkleRoots;
        mapping (string /* eventID */ => bytes /* EventDetailsHash */)  public eventHashes;
        mapping (string /* eventID */ => TallyScore) public tallyScores;
        mapping (string /* eventID */ => uint256) public totalEventVotes;

        error EventRootMismatch(uint256 expectedRoot, uint256 gotRoot);

        // Next step, store a hash of all the nullifiers for consistency!

        modifier onlyOwner() {
            require(msg.sender == admin, "Only owner is allowed to execute this transaction.");
            _;
        }

        function SetVoteMerkleRoot(string calldata eventID, uint256 value) external {
            // check before to make sure its a valid hash
            voteMerkleRoots[eventID] = value;
        }

        function SetEventHash(string calldata eventID, bytes calldata value) external {
            // A change in event details changes the event hash completely.
            eventHashes[eventID] = value;
        }

        function GetVoteMerkleRoot(string calldata eventID) external view returns (uint256) {
            return voteMerkleRoots[eventID];
        }

        function GetEventHash(string calldata eventID) external view returns (bytes memory) {
            return eventHashes[eventID];
        }

        function GetTallyScore(string memory eventID) external view returns (uint256[4] memory){
            TallyScore storage score = tallyScores[eventID];
            return [score.C1x, score.C1y, score.C2x, score.C2y];
        }

        function GetTotalEventVotes(string memory eventID) external view returns (uint256) {
            return totalEventVotes[eventID];
        }

        // ---- BabyJubJub point addition (twisted Edwards, affine) ----

        // circomlib babyjub constants (make sure these match your circuit!)
        uint256 constant BABYJUB_Q =
            21888242871839275222246405745257275088548364400416034343698204186575808495617;
        uint256 constant BABYJUB_A = 168700;
        uint256 constant BABYJUB_D = 168696;

        function _submodQ(uint256 a, uint256 b) internal pure returns (uint256) {
            unchecked { return addmod(a, BABYJUB_Q - (b % BABYJUB_Q), BABYJUB_Q); }
        }

        function _modExp(uint256 base, uint256 exp, uint256 mod_) internal view returns (uint256 result) {
            // bigModExp precompile 0x05
            bytes memory input = abi.encodePacked(
                uint256(32), uint256(32), uint256(32),
                base, exp, mod_
            );
            bytes memory output = new bytes(32);
            bool ok;
            assembly ("memory-safe") {
                ok := staticcall(gas(), 0x05, add(input, 32), mload(input), add(output, 32), 32)
            }
            require(ok, "modexp failed");
            result = abi.decode(output, (uint256));
        }

        function _invQ(uint256 x) internal view returns (uint256) {
            require(x != 0, "inv(0)");
            return _modExp(x, BABYJUB_Q - 2, BABYJUB_Q);
        }

        /// @notice Adds two BabyJubJub affine points (x1,y1) + (x2,y2)
        function _babyJubAdd(
            uint256 x1, uint256 y1,
            uint256 x2, uint256 y2
        ) internal view returns (uint256 x3, uint256 y3) {
            // x_num = x1*y2 + y1*x2
            uint256 x_num = addmod(mulmod(x1, y2, BABYJUB_Q), mulmod(y1, x2, BABYJUB_Q), BABYJUB_Q);

            // y_num = y1*y2 - a*x1*x2
            uint256 y_num = _submodQ(
                mulmod(y1, y2, BABYJUB_Q),
                mulmod(BABYJUB_A, mulmod(x1, x2, BABYJUB_Q), BABYJUB_Q)
            );

            // t = d*x1*x2*y1*y2
            uint256 t = mulmod(
                BABYJUB_D,
                mulmod(mulmod(x1, x2, BABYJUB_Q), mulmod(y1, y2, BABYJUB_Q), BABYJUB_Q),
                BABYJUB_Q
            );

            // x_den = 1 + t
            uint256 x_den = addmod(1, t, BABYJUB_Q);

            // y_den = 1 - t
            uint256 y_den = _submodQ(1, t);

            // x3 = x_num / x_den
            x3 = mulmod(x_num, _invQ(x_den), BABYJUB_Q);

            // y3 = y_num / y_den
            y3 = mulmod(y_num, _invQ(y_den), BABYJUB_Q);
        }


        // No revert means the proof has been verified and the vote has been submitted.
        function Vote(VoteSubmission calldata proof) external{
            if (proof.proofs.membership.Input[0] != voteMerkleRoots[proof.eventID])
                revert EventRootMismatch(voteMerkleRoots[proof.eventID], proof.proofs.membership.Input[0]);

            bool nullifier_accepted = nullifier.verifyProof(
                [proof.proofs.nullifier.Proof.Arx, proof.proofs.nullifier.Proof.Ary],
                [
                    [proof.proofs.nullifier.Proof.Brx1, proof.proofs.nullifier.Proof.Brx0],
                    [proof.proofs.nullifier.Proof.Bry1, proof.proofs.nullifier.Proof.Bry0]
                ],
                [proof.proofs.nullifier.Proof.Cx, proof.proofs.nullifier.Proof.Cy],
                proof.proofs.nullifier.Input
            );
            if (!nullifier_accepted) revert NullifierProofInvalid();

            bool membership_accepted = membership.verifyProof(
                [proof.proofs.membership.Proof.Arx, proof.proofs.membership.Proof.Ary],
                [
                    [proof.proofs.membership.Proof.Brx1, proof.proofs.membership.Proof.Brx0],
                    [proof.proofs.membership.Proof.Bry1, proof.proofs.membership.Proof.Bry0]
                ],
                [proof.proofs.membership.Proof.Cx, proof.proofs.membership.Proof.Cy],
                proof.proofs.membership.Input
            );
            if (!membership_accepted) revert MembershipProofInvalid();

            bool ballot_accepted = ballot.verifyProof(
                [proof.proofs.ballot.Proof.Arx, proof.proofs.ballot.Proof.Ary],
                [
                    [proof.proofs.ballot.Proof.Brx1, proof.proofs.ballot.Proof.Brx0],
                    [proof.proofs.ballot.Proof.Bry1, proof.proofs.ballot.Proof.Bry0]
                ],
                [proof.proofs.ballot.Proof.Cx, proof.proofs.ballot.Proof.Cy],
                proof.proofs.ballot.Input
            );
            if (!ballot_accepted) revert BallotProofInvalid();


            // Preparing input parameters to perform tally + C.
            string memory eventID = proof.eventID;
            TallyScore storage tallyScore = tallyScores[eventID];

            uint256[2] memory C1 = [proof.proofs.ballot.Input[0], proof.proofs.ballot.Input[1]];
            uint256[2] memory C2 = [proof.proofs.ballot.Input[2], proof.proofs.ballot.Input[3]];

            // IMPORTANT: BabyJubJub identity is (0,1). If you never initialized tallyScore,
            // the default is (0,0) which is NOT a valid point. Treat (0,0) as "unset" and
            // initialize to identity on first vote.
            if (tallyScore.C1x == 0 && tallyScore.C1y == 0) {
                tallyScore.C1x = 0;
                tallyScore.C1y = 1;
            }
            if (tallyScore.C2x == 0 && tallyScore.C2y == 0) {
                tallyScore.C2x = 0;
                tallyScore.C2y = 1;
            }

            uint256[2] memory tallyC1 = [tallyScore.C1x, tallyScore.C1y];
            uint256[2] memory tallyC2 = [tallyScore.C2x, tallyScore.C2y];

            // Add BabyJubJub points: result = C + tallyC
            (uint256 r1x, uint256 r1y) = _babyJubAdd(C1[0], C1[1], tallyC1[0], tallyC1[1]);
            (uint256 r2x, uint256 r2y) = _babyJubAdd(C2[0], C2[1], tallyC2[0], tallyC2[1]);

            // Update tally score.
            tallyScore.C1x = r1x;
            tallyScore.C1y = r1y;
            tallyScore.C2x = r2x;
            tallyScore.C2y = r2y;

            totalEventVotes[eventID] = totalEventVotes[eventID] + 1;

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