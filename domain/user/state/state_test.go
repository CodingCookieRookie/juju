// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"crypto/rand"
	"database/sql"
	"time"

	"github.com/canonical/sqlair"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"golang.org/x/net/context"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/user"
	schematesting "github.com/juju/juju/domain/schema/testing"
	usererrors "github.com/juju/juju/domain/user/errors"
	"github.com/juju/juju/internal/auth"
	"github.com/juju/juju/internal/database"
)

type stateSuite struct {
	schematesting.ControllerSuite
}

var _ = gc.Suite(&stateSuite{})

// TestSingletonActiveUser asserts the idx_singleton_active_user unique index
// in the DDL. What we need in the DDL is the ability to have multiple users
// with the same username. However, only one username can exist in the table
// where the username has not been removed. We are free to have as many removed
// identical usernames as we want.
//
// This test will make 3 users called "bob" that have all been removed. This
// check asserts that we can have more than one removed bob.
// We will then add another user called "bob" that is not removed
// (an active user). This should not fail.
// We will then try and add a 5 user called "bob" that is also not removed and
// this will produce a unique index constraint error.
func (s *stateSuite) TestSingletonActiveUser(c *gc.C) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, "123", "bob", "Bob", true, "123", time.Now())
		return err
	})
	c.Assert(err, jc.ErrorIsNil)

	err = s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, "124", "bob", "Bob", true, "123", time.Now())
		return err
	})
	c.Assert(err, jc.ErrorIsNil)

	err = s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, "125", "bob", "Bob", true, "123", time.Now())
		return err
	})
	c.Assert(err, jc.ErrorIsNil)

	// Insert the first non-removed (active) Bob user.
	err = s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, "126", "bob", "Bob", false, "123", time.Now())
		return err
	})
	c.Assert(err, jc.ErrorIsNil)

	// Try and insert the second non-removed (active) Bob user. This should blow
	// up the constraint.
	err = s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, "127", "bob", "Bob", false, "123", time.Now())
		return err
	})
	c.Assert(database.IsErrConstraintUnique(err), jc.IsTrue)
}

func generateActivationKey() ([]byte, error) {
	var activationKey [32]byte
	if _, err := rand.Read(activationKey[:]); err != nil {
		return nil, errors.Annotate(err, "generating activation key")
	}
	return activationKey[:], nil
}

// AddUserWithPassword asserts that we can add a user with no
// password authorization.
func (s *stateSuite) TestBootstrapAddUserWithPassword(c *gc.C) {
	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Add user with no password authorization.
	err = s.TxnRunner().Txn(context.Background(), func(ctx context.Context, tx *sqlair.TX) error {
		err = AddUserWithPassword(context.Background(), tx, adminUUID, adminUser, adminUUID, "passwordHash", salt)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)

	// Check that the user was added correctly.
	db := s.DB()

	row := db.QueryRow(`
SELECT uuid, name, display_name, removed, created_by_uuid, created_at
FROM user
WHERE uuid = ?
	`, adminUUID)

	c.Assert(row.Err(), jc.ErrorIsNil)

	var uuid, name, displayName string
	var creatorUUID user.UUID
	var removed bool
	var createdAt time.Time
	err = row.Scan(&uuid, &name, &displayName, &removed, &creatorUUID, &createdAt)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(name, gc.Equals, adminUser.Name)
	c.Check(removed, gc.Equals, false)
	c.Check(displayName, gc.Equals, adminUser.DisplayName)
	c.Check(creatorUUID, gc.Equals, adminUUID)
	c.Check(createdAt, gc.NotNil)
}

// TestAddUserAlreadyExists asserts that we get an error when we try to add a
// user that already exists.
func (s *stateSuite) TestAddUserAlreadyExists(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Try and add admin user again.
	adminCloneUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUser(context.Background(), adminCloneUUID, adminUser, adminCloneUUID)
	c.Assert(err, jc.ErrorIs, usererrors.AlreadyExists)
}

// TestAddUserCreatorNotFound asserts that we get an error when we try
// to add a user that has a creator that does not exist.
func (s *stateSuite) TestAddUserCreatorNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	// Try and add admin user with a creator that does not exist.
	nonExistingUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUser(context.Background(), adminUUID, adminUser, nonExistingUUID)
	c.Assert(err, jc.ErrorIs, usererrors.UserCreatorUUIDNotFound)
}

// TestGetUser asserts that we can get a user from the database.
func (s *stateSuite) TestGetUser(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	u, err := st.GetUser(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(u.Name, gc.Equals, adminUser.Name)
	c.Check(u.DisplayName, gc.Equals, adminUser.DisplayName)
	c.Check(u.CreatorUUID, gc.Equals, adminUUID)
	c.Check(u.CreatedAt, gc.NotNil)
}

// TestGetRemovedUser asserts that we can get a removed user from the database.
func (s *stateSuite) TestGetRemovedUser(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Add userToRemove.
	userToRemoveUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	userToRemove := user.User{
		Name:        "userToRemove",
		DisplayName: "userToRemove",
	}
	err = st.AddUser(context.Background(), userToRemoveUUID, userToRemove, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Remove userToRemove.
	err = st.RemoveUser(context.Background(), userToRemoveUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	u, err := st.GetUser(context.Background(), userToRemoveUUID)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(u.Name, gc.Equals, userToRemove.Name)
	c.Check(u.DisplayName, gc.Equals, userToRemove.DisplayName)
	c.Check(u.CreatorUUID, gc.Equals, adminUUID)
	c.Check(u.CreatedAt, gc.NotNil)
}

// TestGetUserNotFound asserts that we get an error when we try to get a user
// that does not exist.
func (s *stateSuite) TestGetUserNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Generate a random UUID.
	userUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	_, err = st.GetUser(context.Background(), userUUID)
	c.Assert(err, jc.ErrorIs, usererrors.NotFound)
}

// TestGetUserByName asserts that we can get a user by name from the database.
func (s *stateSuite) TestGetUserByName(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	u, err := st.GetUserByName(context.Background(), adminUser.Name)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(u.Name, gc.Equals, adminUser.Name)
	c.Check(u.DisplayName, gc.Equals, adminUser.DisplayName)
	c.Check(u.CreatorUUID, gc.Equals, adminUUID)
	c.Check(u.CreatedAt, gc.NotNil)
}

// TestGetRemovedUserByName asserts that we can get only non-removed user by name.
func (s *stateSuite) TestGetRemovedUserByName(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Add userToRemove.
	userToRemoveUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	userToRemove := user.User{
		Name:        "userToRemove",
		DisplayName: "userToRemove",
	}
	err = st.AddUser(context.Background(), userToRemoveUUID, userToRemove, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Remove userToRemove.
	err = st.RemoveUser(context.Background(), userToRemoveUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	_, err = st.GetUserByName(context.Background(), userToRemove.Name)
	c.Assert(err, jc.ErrorIs, usererrors.NotFound)
}

// TestGetUserByNameMultipleUsers asserts that we get a non-removed user when we try to
// get a user by name that has multiple users with the same name.
func (s *stateSuite) TestGetUserByNameMultipleUsers(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Remove admin user.
	err = st.RemoveUser(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Add admin2 user.
	admin2UUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	admin2User := user.User{
		Name:        "admin",
		DisplayName: "admin2",
	}
	err = st.AddUser(context.Background(), admin2UUID, admin2User, admin2UUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	u, err := st.GetUserByName(context.Background(), "admin")
	c.Assert(err, jc.ErrorIsNil)

	c.Check(u.Name, gc.Equals, adminUser.Name)
	c.Check(u.DisplayName, gc.Equals, admin2User.DisplayName)
	c.Check(u.CreatorUUID, gc.Equals, admin2UUID)
	c.Check(u.CreatedAt, gc.NotNil)
}

// TestGetUserByNameNotFound asserts that we get an error when we try to get a
// user by name that does not exist.
func (s *stateSuite) TestGetUserByNameNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Get the user.
	_, err := st.GetUserByName(context.Background(), "admin")
	c.Assert(err, jc.ErrorIs, usererrors.NotFound)
}

// TestGetUserWithAuthInfoByName asserts that we can get a user with auth info
// by name from the database.
func (s *stateSuite) TestGetUserWithAuthInfoByName(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user with password hash.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Get the user.
	u, err := st.GetUserWithAuthInfoByName(context.Background(), adminUser.Name)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(u.Name, gc.Equals, adminUser.Name)
	c.Check(u.DisplayName, gc.Equals, adminUser.DisplayName)
	c.Check(u.CreatorUUID, gc.Equals, adminUUID)
	c.Check(u.CreatedAt, gc.NotNil)
	c.Check(u.LastLogin, gc.NotNil)
	c.Check(u.Disabled, gc.Equals, false)
}

// TestRemoveUser asserts that we can remove a user from the database.
func (s *stateSuite) TestRemoveUser(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	err = st.AddUser(context.Background(), adminUUID, adminUser, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Add userToRemove.
	userToRemoveUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	userToRemove := user.User{
		Name:        "userToRemove",
		DisplayName: "userToRemove",
	}
	err = st.AddUser(context.Background(), userToRemoveUUID, userToRemove, adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Remove userToRemove.
	err = st.RemoveUser(context.Background(), userToRemoveUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the user was removed correctly.
	db := s.DB()

	row := db.QueryRow(`
SELECT removed
FROM user
WHERE uuid = ?
	`, userToRemoveUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var removed bool
	err = row.Scan(&removed)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(removed, gc.Equals, true)
}

// GetAllUsersWihAuthInfo asserts that we can get all users with auth info from
// the database.
func (s *stateSuite) TestGetAllUsersWihAuthInfo(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin1 user with password hash.
	admin1UUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	admin1User := user.User{
		Name:        "admin1",
		DisplayName: "admin1",
	}
	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUserWithPasswordHash(context.Background(), admin1UUID, admin1User, admin1UUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Add admin2 user with activation key.
	admin2UUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	admin2User := user.User{
		Name:        "admin2",
		DisplayName: "admin2",
	}
	admin2ActivationKey, err := generateActivationKey()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUserWithActivationKey(context.Background(), admin2UUID, admin2User, admin2UUID, admin2ActivationKey)
	c.Assert(err, jc.ErrorIsNil)

	// Disable admin2 user.
	err = st.DisableUserAuthentication(context.Background(), admin2UUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get all users with auth info.
	users, err := st.GetAllUsersWithAuthInfo(context.Background())
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(users, gc.HasLen, 2)

	c.Check(users[0].Name, gc.Equals, admin1User.Name)
	c.Check(users[0].DisplayName, gc.Equals, admin1User.DisplayName)
	c.Check(users[0].CreatorUUID, gc.Equals, admin1UUID)
	c.Check(users[0].CreatedAt, gc.NotNil)
	c.Check(users[0].LastLogin, gc.NotNil)
	c.Check(users[0].Disabled, gc.Equals, false)

	c.Check(users[1].Name, gc.Equals, admin2User.Name)
	c.Check(users[1].DisplayName, gc.Equals, admin2User.DisplayName)
	c.Check(users[1].CreatorUUID, gc.Equals, admin2UUID)
	c.Check(users[1].CreatedAt, gc.NotNil)
	c.Check(users[1].LastLogin, gc.NotNil)
	c.Check(users[1].Disabled, gc.Equals, true)
}

// TestUserWithAuthInfo asserts that we can get a user with auth info from the
// database.
func (s *stateSuite) TestUserWithAuthInfo(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin1 user with password hash.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Disable admin1 user.
	err = st.DisableUserAuthentication(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Get user with auth info.
	u, err := st.GetUserWithAuthInfo(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(u.Name, gc.Equals, adminUser.Name)
	c.Check(u.DisplayName, gc.Equals, adminUser.DisplayName)
	c.Check(u.CreatorUUID, gc.Equals, adminUUID)
	c.Check(u.CreatedAt, gc.NotNil)
	c.Check(u.LastLogin, gc.NotNil)
	c.Check(u.Disabled, gc.Equals, true)
}

// TestSetPasswordHash asserts that we can set a password hash for a user.
func (s *stateSuite) TestSetPasswordHash(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user with activation key.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	newActivationKey, err := generateActivationKey()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUserWithActivationKey(context.Background(), adminUUID, adminUser, adminUUID, newActivationKey)
	c.Assert(err, jc.ErrorIsNil)

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Set password hash.
	err = st.SetPasswordHash(context.Background(), adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the password hash was set correctly.
	db := s.DB()

	rowAuth := db.QueryRow(`
SELECT disabled
FROM user_authentication
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(rowAuth.Err(), jc.ErrorIsNil)

	var disabled bool
	err = rowAuth.Scan(&disabled)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(disabled, gc.Equals, false)

	row := db.QueryRow(`
SELECT password_hash
FROM user_password
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var passwordHash string
	err = row.Scan(&passwordHash)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(passwordHash, gc.Equals, "passwordHash")

	row = db.QueryRow(`
SELECT activation_key
FROM user_activation_key
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var activationKey string
	err = row.Scan(&activationKey)
	c.Assert(err, jc.ErrorIs, sql.ErrNoRows)
}

// TestAddUserWithPasswordHash asserts that we can add a user with a password
// hash.
func (s *stateSuite) TestAddUserWithPasswordHash(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Add user with password hash.
	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the password hash was set correctly.
	db := s.DB()

	row := db.QueryRow(`
SELECT password_hash
FROM user_password
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var passwordHash string
	err = row.Scan(&passwordHash)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(passwordHash, gc.Equals, "passwordHash")
}

// TestAddUserWithPasswordWhichCreatorDoesNotExist asserts that we get an error
// when we try to add a user with a password that has a creator that does not
// exist.
func (s *stateSuite) TestAddUserWithPasswordWhichCreatorDoesNotExist(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	nonExistedCreatorUuid, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Try and add admin user with a creator that does not exist.
	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, nonExistedCreatorUuid, "passwordHash", salt)
	c.Assert(err, jc.ErrorIs, usererrors.UserCreatorUUIDNotFound)
}

// TestAddUserWithActivationKey asserts that we can add a user with an
// activation key.
func (s *stateSuite) TestAddUserWithActivationKey(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user with activation key.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}
	adminActivationKey, err := generateActivationKey()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUserWithActivationKey(context.Background(), adminUUID, adminUser, adminUUID, adminActivationKey)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the activation key was set correctly.
	db := s.DB()

	row := db.QueryRow(`
SELECT activation_key
FROM user_activation_key
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var activationKey string
	err = row.Scan(&activationKey)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(activationKey, gc.Equals, string(adminActivationKey))
}

// TestAddUserWithActivationKeyWhichCreatorDoesNotExist asserts that we get an
// error when we try to add a user with an activation key that has a creator
// that does not exist.
func (s *stateSuite) TestAddUserWithActivationKeyWhichCreatorDoesNotExist(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user with activation key.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	nonExistedCreatorUuid, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	// Try and add admin user with an activation key with a creator that does not exist.
	newActivationKey, err := generateActivationKey()
	c.Assert(err, jc.ErrorIsNil)
	err = st.AddUserWithActivationKey(context.Background(), adminUUID, adminUser, nonExistedCreatorUuid, newActivationKey)
	c.Assert(err, jc.ErrorIs, usererrors.UserCreatorUUIDNotFound)
}

// TestSetActivationKey asserts that we can set an activation key for a user.
func (s *stateSuite) TestSetActivationKey(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Add user with password hash.
	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Set activation key.
	adminActivationKey, err := generateActivationKey()
	c.Assert(err, jc.ErrorIsNil)
	err = st.SetActivationKey(context.Background(), adminUUID, adminActivationKey)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the activation key was set correctly, and the password hash was removed.
	db := s.DB()

	row := db.QueryRow(`
SELECT activation_key
FROM user_activation_key
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var activationKey string
	err = row.Scan(&activationKey)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(activationKey, gc.Equals, string(adminActivationKey))

	row = db.QueryRow(`
SELECT password_hash, password_salt
FROM user_password
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var passwordHash, passwordSalt string
	err = row.Scan(&passwordHash, &passwordSalt)
	c.Assert(err, jc.ErrorIs, sql.ErrNoRows)
}

// TestDisableUserAuthentication asserts that we can disable a user.
func (s *stateSuite) TestDisableUserAuthentication(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Add user with password hash.
	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Disable user.
	err = st.DisableUserAuthentication(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the user was disabled correctly.
	db := s.DB()

	row := db.QueryRow(`
SELECT disabled
FROM user_authentication
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var disabled bool
	err = row.Scan(&disabled)
	c.Assert(err, jc.ErrorIsNil)

	c.Check(disabled, gc.Equals, true)
}

// TestEnableUserAuthentication asserts that we can enable a user.
func (s *stateSuite) TestEnableUserAuthentication(c *gc.C) {
	st := NewState(s.TxnRunnerFactory())

	// Add admin user with activation key.
	adminUUID, err := user.NewUUID()
	c.Assert(err, jc.ErrorIsNil)
	adminUser := user.User{
		Name:        "admin",
		DisplayName: "admin",
	}

	salt, err := auth.NewSalt()
	c.Assert(err, jc.ErrorIsNil)

	// Add user with password hash.
	err = st.AddUserWithPasswordHash(context.Background(), adminUUID, adminUser, adminUUID, "passwordHash", salt)
	c.Assert(err, jc.ErrorIsNil)

	// Disable user.
	err = st.DisableUserAuthentication(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Enable user.
	err = st.EnableUserAuthentication(context.Background(), adminUUID)
	c.Assert(err, jc.ErrorIsNil)

	// Check that the user was enabled correctly.
	db := s.DB()

	row := db.QueryRow(`
SELECT disabled
FROM user_authentication
WHERE user_uuid = ?
	`, adminUUID)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var disabled bool
	err = row.Scan(&disabled)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(disabled, gc.Equals, false)
}
