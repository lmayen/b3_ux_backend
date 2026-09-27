package auth

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/entities/session"
	"b3_ux_backend/internal/entities/user"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/olekukonko/errors"
)

func NewSession(ctx context.Context, userId uuid.UUID) (string, error) {
	// --- Generate credential --- //
	token, err := GenerateSessionToken()
	if err != nil {
		return "", errors.Wrapf(err, "NewSession ERR: could not generate session token")
	}

	tokenHash := HashSecret(token)

	// --- Ensure artist exists and is active --- //
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "NewSession ERR: could not create transaction")
	}
	defer tx.Rollback()

	exists, err := tx.User.
		Query().
		Where(user.IDEQ(userId)).
		Exist(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "NewSession ERR: could not find active artist")
	}
	if !exists {
		return "", errors.Wrapf(err, "NewSession ERR: artist not found")
	}

	// --- Record session --- //
	_, err = tx.Session.Delete().Where(session.HasArtistWith(user.IDEQ(userId))).Exec(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "NewSession ERR: could not delete old sessions")
	}

	_, err = tx.Session.
		Create().
		SetToken(tokenHash).
		SetArtistID(userId).
		Save(ctx)
	if err != nil {
		return "", errors.Wrapf(err, "NewSession ERR: could not create session")
	}

	err = tx.Commit()
	if err != nil {
		return "", errors.Wrapf(err, "NewSession ERR: could not commit transaction")
	}

	return token, nil
}

func ResolveSession(ctx context.Context, token string) (*entities.Session, *entities.User, error) {
	if token == "" {
		return nil, nil, errInvalidSession
	}

	tokenHash := HashSecret(token)

	entitySession, err := db.Client.Session.Query().
		Where(session.TokenEQ(tokenHash)).
		WithArtist().
		Only(ctx)
	if err != nil {
		if entities.IsNotFound(err) {
			return nil, nil, errInvalidSession
		}

		return nil, nil, errors.Wrapf(err, "ResolveSession ERR: could not query session")
	}

	// An expired session has no further value.
	if !entitySession.ExpiresAt.After(time.Now()) {
		err = db.Client.Session.DeleteOneID(entitySession.ID).Exec(ctx)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "ResolveSession ERR: could not delete expired session")
		}

		return nil, nil, errInvalidSession
	}

	entityArtist := entitySession.Edges.Artist
	if entityArtist == nil {
		return nil, nil, errors.New("ResolveSession ERR: session has no artist")
	}

	return entitySession, entityArtist, nil
}

func DeleteOneSession(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}

	tokenHash := HashSecret(token)

	_, err := db.Client.Session.
		Delete().
		Where(session.TokenEQ(tokenHash)).
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "DeleteSession ERR: could not delete session")
	}

	return nil
}

func DeleteExpiredSessions(ctx context.Context) error {
	_, err := db.Client.Session.
		Delete().
		Where(session.ExpiresAtLTE(time.Now())).
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "DeleteExpiredSessions ERR: could not delete expired sessions")
	}

	return nil
}
