package main

import (
	"crypto/sha256"
	"fmt"
)

// Vote represents a single vote cast by a voter.
type Vote struct {
	VoterID   int
	Candidate string
}

// Block represents a block in the blockchain, containing multiple votes.
type Block struct {
	PrevHash    string
	CurrentHash string // Add CurrentHash field
	Votes       []Vote
}

// Blockchain is a slice of Block elements.
var Blockchain []Block

// Candidates is a map to store candidate vote counts.
var Candidates map[string]int

func init() {
	Candidates = make(map[string]int)
}

var VoterReg map[int]bool

func init() {
	VoterReg = make(map[int]bool)
}

// // // calculateHash calculates the hash of a block.
func calculateHash(block Block, vote Vote) string {

	hash := sha256.Sum256([]byte(fmt.Sprintf("%v%v", block, vote)))
	hex := fmt.Sprintf("%x", hash)
	return hex
}

// RegisterVoter adds a new voter to the system.
func RegisterVoter(voterID int) {

	if _, exists := VoterReg[voterID]; exists {
		fmt.Printf("Voter %d is already registered.\n", voterID)
		return
	}

	VoterReg[voterID] = true
	fmt.Printf("Voter %d registered.\n", voterID)
}

// CastVote allows a registered voter to cast a vote.
func CastVote(voterID int, candidate string) {
	// Check if the voter is registered
	// voterRegistered := false
	// for i := 0; i < len(Blockchain); i++ {
	// 	for j := 0; j < len(Blockchain[i].Votes); j++ {
	// 		if Blockchain[i].Votes[j].VoterID == voterID {
	// 			voterRegistered = true
	// 			break
	// 		}
	// 	}
	// 	if !voterRegistered {
	// 		break
	// 	}
	// }

	// if !voterRegistered {
	// 	fmt.Printf("Voter with ID %d is not registered.\n", voterID)
	// 	return
	// }
	if _, exists := VoterReg[voterID]; !exists {
		fmt.Printf("Voter with ID %d is not registered.\n", voterID)
		return
	}
	// Check if the candidate exists
	if _, exists := Candidates[candidate]; !exists {
		fmt.Printf("Candidate %s does not exist.\n", candidate)
		return
	}

	// Check if the voter has already cast a vote
	for i := 0; i < len(Blockchain); i++ {
		for j := 0; j < len(Blockchain[i].Votes); j++ {
			if Blockchain[i].Votes[j].VoterID == voterID {
				fmt.Printf("Voter %d has already cast a vote.\n", voterID)
				return
			}
		}
	}

	// Add the vote to the current block
	vote := Vote{VoterID: voterID, Candidate: candidate}
	lastBlock := Blockchain[len(Blockchain)-1]
	newBlock := Block{
		PrevHash:    lastBlock.CurrentHash,
		CurrentHash: "",
		Votes:       []Vote{vote},
	}
	newBlock.CurrentHash = calculateHash(newBlock, vote)
	Blockchain = append(Blockchain, newBlock)

	// Update candidate vote count
	Candidates[candidate]++

	fmt.Printf("Vote cast by Voter %d for %s is recorded.\n", voterID, candidate)
}

// CalculateElectionResults calculates and displays the winner of the election.
func CalculateElectionResults() {
	fmt.Println("\nTotal Candidates: ")
	fmt.Printf("%d", len(Candidates))
	fmt.Println("\n\t\t\t\t\t\t-------------------------------\n")
	fmt.Println("\n\nElection Results:\n")
	for i, count := range Candidates {
		fmt.Printf("%s :\t%d\n", i, count)
	}
	// for i := 0; i < len(Blockchain); i++ {
	// 	for j := 0; j < len(Blockchain[i].Votes); j++ {
	// 		fmt.Printf("%d", Candidates[Blockchain[i].Candidate])
	// 	}
	// }

	var winner string
	maxVotes := -1

	// Write your logic for calculating the winner of the election
	for candidate, votes := range Candidates {
		if votes > maxVotes {
			maxVotes = votes
			winner = candidate
		} else if votes == maxVotes {
			winner = "Tie"
		}
	}
	if winner != "Tie" {
		fmt.Printf("\nWinner: %s\n", winner)
	} else {
		fmt.Println("\nElection resulted in a tie.")
	}
}

func main() {
	// Initialize the blockchain with a genesis block.
	genesisBlock := Block{PrevHash: "", CurrentHash: "", Votes: nil}
	Blockchain = append(Blockchain, genesisBlock)

	// Register candidates
	Candidates = make(map[string]int)
	Candidates["Candidate A"] = 0
	Candidates["Candidate B"] = 0

	// Register voters
	for i := 1; i <= 10; i++ {
		RegisterVoter(i)
	}
	fmt.Println("\n\n")
	// Simulate voting process
	CastVote(1, "Candidate A")
	CastVote(2, "Candidate B")
	CastVote(3, "Candidate A")
	CastVote(3, "Candidate B") // Attempted Double Voting
	CastVote(4, "Candidate B")
	CastVote(5, "Candidate A")
	CastVote(5, "Candidate A")  // Attempted Double Voting
	CastVote(6, "Candidate B")  // Should Print in case of tie
	CastVote(7, "Candidate C")  // Invalid Candidate ID
	CastVote(11, "Candidate B") // Invalid Voter ID

	// Calculate and display election results
	CalculateElectionResults()

	// Display the blockchain
	fmt.Println("\nBlockchain:")
	for i, block := range Blockchain {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("CurrentHash: %s\n", block.CurrentHash)
		fmt.Printf("Votes: %v\n\n", block.Votes)
	}
}
