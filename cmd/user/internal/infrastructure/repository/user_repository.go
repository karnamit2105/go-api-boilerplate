/*
Package repository holds event sourced repositories
*/
package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/vardius/go-api-boilerplate/cmd/user/internal/domain/user"
	"github.com/vardius/go-api-boilerplate/pkg/application"
	"github.com/vardius/go-api-boilerplate/pkg/errors"
	"github.com/vardius/go-api-boilerplate/pkg/eventbus"
	"github.com/vardius/go-api-boilerplate/pkg/eventstore"
)

type userRepository struct {
	eventStore eventstore.EventStore
	eventBus   eventbus.EventBus
}

// NewUserRepository creates new user event sourced repository
func NewUserRepository(store eventstore.EventStore, bus eventbus.EventBus) user.Repository {
	return &userRepository{store, bus}
}

// Save current user changes to event store and publish each event with an event bus
func (r *userRepository) Save(ctx context.Context, u user.User) error {
	err := r.eventStore.Store(u.Changes())
	if err != nil {
		return errors.Wrap(err)
	}

	for _, event := range u.Changes() {
		if err := r.eventBus.Publish(ctx, event); err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

// Get user with current state applied
func (r *userRepository) Get(id uuid.UUID) (user.User, error) {
	events := r.eventStore.GetStream(id, user.StreamName)

	if len(events) == 0 {
		return user.User{}, application.ErrNotFound
	}

	return user.FromHistory(events), nil
}
