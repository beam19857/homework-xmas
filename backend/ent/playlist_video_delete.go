// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/beam19857/playlist-video/ent/playlist_video"
	"github.com/beam19857/playlist-video/ent/predicate"
)

// PlaylistVideoDelete is the builder for deleting a PlaylistVideo entity.
type PlaylistVideoDelete struct {
	config
	hooks      []Hook
	mutation   *PlaylistVideoMutation
	predicates []predicate.Playlist_Video
}

// Where adds a new predicate to the delete builder.
func (pvd *PlaylistVideoDelete) Where(ps ...predicate.Playlist_Video) *PlaylistVideoDelete {
	pvd.predicates = append(pvd.predicates, ps...)
	return pvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pvd *PlaylistVideoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pvd.hooks) == 0 {
		affected, err = pvd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvd.mutation = mutation
			affected, err = pvd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pvd.hooks) - 1; i >= 0; i-- {
			mut = pvd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvd *PlaylistVideoDelete) ExecX(ctx context.Context) int {
	n, err := pvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pvd *PlaylistVideoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: playlist_video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist_video.FieldID,
			},
		},
	}
	if ps := pvd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pvd.driver, _spec)
}

// PlaylistVideoDeleteOne is the builder for deleting a single Playlist_Video entity.
type PlaylistVideoDeleteOne struct {
	pvd *PlaylistVideoDelete
}

// Exec executes the deletion query.
func (pvdo *PlaylistVideoDeleteOne) Exec(ctx context.Context) error {
	n, err := pvdo.pvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{playlist_video.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pvdo *PlaylistVideoDeleteOne) ExecX(ctx context.Context) {
	pvdo.pvd.ExecX(ctx)
}
