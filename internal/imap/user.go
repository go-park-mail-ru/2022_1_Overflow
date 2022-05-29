package imap_server

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/models"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/auth_proto"
	"OverflowBackend/proto/folder_manager_proto"
	"OverflowBackend/proto/mailbox_proto"
	"OverflowBackend/proto/profile_proto"
	"OverflowBackend/proto/utils_proto"
	"context"
	"encoding/json"
	"errors"

	imap_backend "github.com/emersion/go-imap/backend"
)

type User struct {
	auth auth_proto.AuthClient
	profile profile_proto.ProfileClient
	mailbox mailbox_proto.MailboxClient
	folderManager folder_manager_proto.FolderManagerClient
	config *config.Config

	userSession *utils_proto.Session
}

func (s *User) Init(
	config *config.Config,
	auth auth_proto.AuthClient,
	profile profile_proto.ProfileClient,
	mailbox mailbox_proto.MailboxClient,
	folderManager folder_manager_proto.FolderManagerClient,
	userSession *utils_proto.Session,
	) {
	s.config = config
	s.auth = auth
	s.profile = profile
	s.mailbox = mailbox
	s.folderManager = folderManager
	s.userSession = userSession
}

// Username returns this user's username.
func (s *User) Username() string {
	return s.userSession.Username
}

// ListMailboxes returns a list of mailboxes belonging to this user. If
// subscribed is set to true, only returns subscribed mailboxes.
func (s *User) ListMailboxes(subscribed bool) ([]imap_backend.Mailbox, error) {
	if subscribed {
		return nil, nil
	}
	resp, err := s.folderManager.ListFolders(context.Background(), &folder_manager_proto.ListFoldersRequest{
		Data: s.userSession,
		Limit: 100,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}
	var response pkg.JsonResponse
	err = json.Unmarshal(resp.Response.Response, &response)
	if err != nil {
		return nil, err
	}
	if response != pkg.NO_ERR {
		return nil, errors.New(response.Message)
	}
	var folders models.FolderList
	err = json.Unmarshal(resp.Folders, &folders)
	if err != nil {
		return nil, err
	}
	// доработать 
	return nil, nil
}

// GetMailbox returns a mailbox. If it doesn't exist, it returns
// ErrNoSuchMailbox.
func (s *User) GetMailbox(name string) (imap_backend.Mailbox, error) {
	// добавить метод проверки в use case существует ли данная папка пользователя
	return nil, nil
}

// CreateMailbox creates a new mailbox.
//
// If the mailbox already exists, an error must be returned. If the mailbox
// name is suffixed with the server's hierarchy separator character, this is a
// declaration that the client intends to create mailbox names under this name
// in the hierarchy.
//
// If the server's hierarchy separator character appears elsewhere in the
// name, the server SHOULD create any superior hierarchical names that are
// needed for the CREATE command to be successfully completed.  In other
// words, an attempt to create "foo/bar/zap" on a server in which "/" is the
// hierarchy separator character SHOULD create foo/ and foo/bar/ if they do
// not already exist.
//
// If a new mailbox is created with the same name as a mailbox which was
// deleted, its unique identifiers MUST be greater than any unique identifiers
// used in the previous incarnation of the mailbox UNLESS the new incarnation
// has a different unique identifier validity value.
func (s *User) CreateMailbox(name string) error

// DeleteMailbox permanently remove the mailbox with the given name. It is an
// error to // attempt to delete INBOX or a mailbox name that does not exist.
//
// The DELETE command MUST NOT remove inferior hierarchical names. For
// example, if a mailbox "foo" has an inferior "foo.bar" (assuming "." is the
// hierarchy delimiter character), removing "foo" MUST NOT remove "foo.bar".
//
// The value of the highest-used unique identifier of the deleted mailbox MUST
// be preserved so that a new mailbox created with the same name will not
// reuse the identifiers of the former incarnation, UNLESS the new incarnation
// has a different unique identifier validity value.
func (s *User) DeleteMailbox(name string) error

// RenameMailbox changes the name of a mailbox. It is an error to attempt to
// rename from a mailbox name that does not exist or to a mailbox name that
// already exists.
//
// If the name has inferior hierarchical names, then the inferior hierarchical
// names MUST also be renamed.  For example, a rename of "foo" to "zap" will
// rename "foo/bar" (assuming "/" is the hierarchy delimiter character) to
// "zap/bar".
//
// If the server's hierarchy separator character appears in the name, the
// server SHOULD create any superior hierarchical names that are needed for
// the RENAME command to complete successfully.  In other words, an attempt to
// rename "foo/bar/zap" to baz/rag/zowie on a server in which "/" is the
// hierarchy separator character SHOULD create baz/ and baz/rag/ if they do
// not already exist.
//
// The value of the highest-used unique identifier of the old mailbox name
// MUST be preserved so that a new mailbox created with the same name will not
// reuse the identifiers of the former incarnation, UNLESS the new incarnation
// has a different unique identifier validity value.
//
// Renaming INBOX is permitted, and has special behavior.  It moves all
// messages in INBOX to a new mailbox with the given name, leaving INBOX
// empty.  If the server implementation supports inferior hierarchical names
// of INBOX, these are unaffected by a rename of INBOX.
func (s *User) RenameMailbox(existingName, newName string) error

// Logout is called when this User will no longer be used, likely because the
// client closed the connection.
func (s *User) Logout() error