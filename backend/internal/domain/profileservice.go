package domain

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProfileService interface {
	GetProfile(ctx context.Context, profileId uint64) (Profile, error)
	UpdateProfile(ctx context.Context, profileId uint64, profile Profile) error
	GetFriendsList(ctx context.Context, id uint64) ([]Profile, error)
	MakeFriendship(ctx context.Context, sourceProfileId uint64, targetProfileId uint64) error
	CancelFriendship(ctx context.Context, sourceProfileId uint64, targetProfileId uint64) error
}

type ProfileServiceImplementation struct {
	db *sqlx.DB
}

const selectQuery = `SELECT Id, Email, FirstName, LastName, Age, Gender, City, Hobby FROM Profile`

func (s ProfileServiceImplementation) GetProfile(ctx context.Context, id uint64) (Profile, error) {
	var profile Profile
	row := s.db.QueryRowContext(ctx, selectQuery)
	err := row.Scan(&profile.Id, &profile.Email, &profile.FirstName, &profile.LastName, &profile.Age, &profile.Gender, &profile.City, &profile.Hobby)
	return profile, err
}

const updateQuery = `UPDATE Profile SET Email = ?, FirstName = ?, LastName = ?, Age = ?, Gender = ?, City = ?, Hobby = ? WHERE Id = ?`

func (s ProfileServiceImplementation) UpdateProfile(ctx context.Context, id uint64, p Profile) error {
	_, err := s.db.ExecContext(ctx, updateQuery, p.Email, p.FirstName, p.LastName, p.Age, p.Gender, p.City, p.Hobby, id)
	return err
}


const getFriendsQuery = "SELECT p.Id, p.Email, p.FirstName, p.LastName, p.Age, p.Gender, p.City, p.Hobby FROM Relationship r INNER JOIN Profile p ON p.Id = r.TargetId WHERE r.SourceId = ?"
func (s ProfileServiceImplementation) GetFriendsList(ctx context.Context, id uint64) ([]Profile, error) {
	row, err := s.db.QueryxContext(ctx, getFriendsQuery, id)
	if err != nil {
		return nil, err
	}

	profiles := make([]Profile, 0)
	for row.Next() {
		var profile Profile
		err = row.Scan(&profile.Id, &profile.Email, &profile.FirstName, &profile.LastName, &profile.Age, &profile.Gender, &profile.City, &profile.Hobby)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

const insertRelationQuery = "INSERT INTO Relationship (Id, SourceId, TargetId) VALUES (UUID(), ? ,? )"
func (s ProfileServiceImplementation) MakeFriendship(ctx context.Context, sourceProfileId uint64, targetProfileId uint64) error {
	_, err := s.db.ExecContext(ctx, insertRelationQuery, sourceProfileId, targetProfileId)
	return err
}

const deleteRelationQuery = "DELETE FROM Relationship WHERE SourceId = ? AND TargetId = ?"
func (s ProfileServiceImplementation) CancelFriendship(ctx context.Context, sourceProfileId uint64, targetProfileId uint64) error {
	_, err := s.db.ExecContext(ctx, deleteRelationQuery, sourceProfileId, targetProfileId)
	return err
}

func NewProfileService(db *sqlx.DB) ProfileService {
	return ProfileServiceImplementation{
		db: db,
	}
}
