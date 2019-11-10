package models

import "github.com/jinzhu/gorm"

type Participant struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255)"`
}

type Tournament struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(255)"`
}

type Rule struct {
	gorm.Model
	Name            string `gorm:"column:name;type:varchar(255)"`
	AdditionalScore int    `gorm:"column:additional_score;type:int"`
}

type RuleMatch struct {
	RuleID       uint   `gorm:"column:rule_id;not null"`
	MatchID      uint   `gorm:"column:match_id;not null"`
	LocalVisitor string `gorm:"column:local_visitor;not null"`
}

type Participation struct {
	gorm.Model
	ParticipantID uint `gorm:"column:participant_id"`
	TournamentID  uint `gorm:"column:tournament_id"`
}

type Match struct {
	gorm.Model
	LocalParticipantOne     uint `gorm:"column:local_participant_one;not null"`
	LocalParticipantTwo     uint `gorm:"column:local_participant_two;null"`
	LocalParticipantThree   uint `gorm:"column:local_participant_three;null"`
	LocalParticipantFour    uint `gorm:"column:local_participant_four;null"`
	VisitorParticipantOne   uint `gorm:"column:visitor_participant_one;not null"`
	VisitorParticipantTwo   uint `gorm:"column:visitor_participant_two;null"`
	VisitorParticipantThree uint `gorm:"column:visitor_participant_three;null"`
	VisitorParticipantFour  uint `gorm:"column:visitor_participant_four;null"`
	LocalResult             int
	VisitorResult           int
}
