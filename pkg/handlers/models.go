package handlers

type CandidateInfo struct {
	SessionUuid   string `uri:"sessionUuid" binding:"required,uuid"`
	CandidateUuid string `uri:"candidateUuid" binding:"required,uuid"`
}
type SessionInfo struct {
	SessionUuid string `uri:"sessionUuid" binding:"required,uuid"`
}
