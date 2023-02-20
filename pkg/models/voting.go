package models

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
	"votingMicroservicesApp/pkg/config"
)

var ctx = context.Background()

//Candidate model representing candidates
type Candidate struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Votes int    `json:"votes"`
}

//NewCandidate :constructor for Candidate model
func NewCandidate(UUID string, name string, votes int) *Candidate {
	return &Candidate{UUID: UUID, Name: name, Votes: votes}
}

//Voting model representing voting sessions
type Voting struct {
	UUID       string      `json:"uuid"`
	Name       string      `json:"name"`
	Candidates []Candidate `json:"candidates"`
}

//Voting model representing voting sessions
type VotingAction struct {
	VotingUUID    string `json:"votingUuid"`
	CandidateUuid string `json:"candidateUuid"`
}

func UnmarshalVotingAction(vote []byte) (*VotingAction, error) {
	var votingAction VotingAction
	err := json.Unmarshal(vote, &votingAction)
	if err != nil {
		return nil, err
	}
	return &votingAction, nil
}

//NewVoting :constructor for Voting sessions
func NewVoting(UUID string, name string, candidates []Candidate) *Voting {
	return &Voting{UUID: UUID, Name: name, Candidates: candidates}
}

func (v *Voting) PersistOnRedis() {
	marshal, err := json.Marshal(v)
	if err != nil {
		panic(err)
		return
	}
	config.AppContext.RedisConfig.Client.Set(ctx, v.UUID, string(marshal), time.Hour*1)
}
func (v *Voting) DeleteOnRedis() {
	config.AppContext.RedisConfig.Client.Del(ctx, v.UUID)
}

func FetchVotingSessions() {
	res := config.AppContext.RedisConfig.Client.Keys(ctx, "*")
	log.Println(res)
}
func FetchVotingSession(uuid string) (Voting, error) {
	var voting Voting
	res, err := config.AppContext.RedisConfig.Client.Get(ctx, uuid).Result()
	if err != nil {
		return Voting{}, err
	}
	err = json.Unmarshal([]byte(res), &voting)
	if err != nil {
		return Voting{}, err
	}
	return voting, err
}
func (v *Voting) SendMessageOnRabbit(queue ...string) {
	marshal, err := json.Marshal(v)
	if err != nil {
		panic(err)
		return
	}
	if len(queue) > 0 {
		config.AppContext.RabbitConfig.Send(string(marshal), "application/json", queue[0])
	} else {
		config.AppContext.RabbitConfig.Send(string(marshal), "application/json")
	}

}

//VoteOnCandidateAction send vote to rabbitmq where the worker will take care of the counting part
func (v *Voting) VoteOnCandidateAction(uuid string) error {
	for i := range v.Candidates {
		if v.Candidates[i].UUID == uuid {
			votingAction := VotingAction{CandidateUuid: uuid, VotingUUID: v.UUID}
			marshal, err := json.Marshal(votingAction)
			if err != nil {
				return err
			}
			config.AppContext.RabbitConfig.Send(string(marshal), "application/json")
			return nil
		}
	}
	return errors.New("candidate not found on voting session")
}
func (v *Voting) VoteOnCandidate(uuid string) error {
	for i := range v.Candidates {
		if v.Candidates[i].UUID == uuid {
			v.Candidates[i].Votes += 1
			v.PersistOnRedis()
			v.SendMessageOnRabbit("events")
			return nil
		}
	}
	return errors.New("candidate not found on voting session")
}
func CountVote(voting *VotingAction) error {
	session, err := FetchVotingSession(voting.VotingUUID)
	if err != nil {
		return err
	}
	err = session.VoteOnCandidate(voting.CandidateUuid)
	if err != nil {
		return err
	}
	return nil
}
