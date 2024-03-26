package main

import (
	"bytes"
	"github.com/bits-and-blooms/bloom/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func loginService(userId int64) (AuthToken, error) {
	token, err := makeToken(userId)
	if err != nil {
		return AuthToken{}, err
	}

	return AuthToken{
		Token:       token,
		ExpiredTime: 1000,
	}, nil
}

func getInitBloom(db *gorm.DB, userId int64) *bloom.BloomFilter {
	var subjectBloom SubjectBloom
	result := db.Where(&SubjectBloom{UserId: userId}).First(&subjectBloom)
	filter := bloom.NewWithEstimates(10000, 0.01)
	if result.Error != nil {
		return filter
	} else {
		var stream = bytes.NewBuffer(subjectBloom.Bloom)
		_, _ = filter.ReadFrom(stream)
		return filter
	}
}

func upsertBloom(db *gorm.DB, filter *bloom.BloomFilter, userId int64) {
	var w = &bytes.Buffer{}
	_, _ = filter.WriteTo(w)

	var subjectBloom = SubjectBloom{
		UserId: userId,
		Bloom:  w.Bytes(),
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"bloom"}),
	}).Create(&subjectBloom)
}

func calcAndSave(db *gorm.DB) {
	m := calc(db)
	saveScore(db, &m)
}

func saveScore(db *gorm.DB, m *map[int64]float64) {
	for anyUserId, score := range *m {
		var userScore UserScore
		userScore.UserId = anyUserId
		userScore.Score = int64(int(score * 10000))
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"score"}),
		}).Create(&userScore)
	}
}

func calc(db *gorm.DB) map[int64]float64 {
	var comments []Comment
	db.Table("subject_comment").
		Select("subject_comment.user_id as RaterId, `subject`.created_by as UserId, subject_comment.score as score").
		Joins("inner join `subject` on `subject`.id = `subject_comment`.subject_id").
		Find(&comments)

	var users []User
	db.Table("user").Find(&users)

	return calcScore(&users, &comments)
}
