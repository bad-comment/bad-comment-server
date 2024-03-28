package main

type User struct {
	Id       int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Account  string `json:"account" gorm:"column:account;type:string;not null"`
	Password string `json:"password" gorm:"column:password;type:string;not null"`
}

func (User) TableName() string {
	return "user"
}

type Subject struct {
	Id        int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"column:name;type:string;not null"`
	CreatedBy int64  `json:"created_by" gorm:"column:created_by;type:bigint;not null"`
}

func (Subject) TableName() string {
	return "subject"
}

type SubjectComment struct {
	Id               int64 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserId           int64 `json:"user_id" gorm:"column:user_id;type:bigint;not null"`
	SubjectId        int64 `json:"subject_id" gorm:"column:subject_id;type:bigint;not null"`
	SubjectIdCreator int64 `json:"subject_id_creator" gorm:"column:subject_id_creator;type:bigint;not null"`
	Score            int8  `json:"score" gorm:"column:score;type:tinyint;not null"`
	C                int32 `json:"c" gorm:"column:c;type:int;not null"`
}

func (SubjectComment) TableName() string {
	return "subject_comment"
}

type UserScore struct {
	UserId int64 `json:"user_id" gorm:"column:user_id;primaryKey;type:bigint"`
	Score  int64 `json:"score" gorm:"column:score;type:bigint;not null"`
}

func (UserScore) TableName() string {
	return "user_score"
}

type SubjectBloom struct {
	UserId int64  `json:"user_id" gorm:"column:user_id;primaryKey;type:bigint"`
	Bloom  []byte `json:"score" gorm:"column:bloom;type:blob;not null"`
}

func (SubjectBloom) TableName() string {
	return "subject_bloom"
}
