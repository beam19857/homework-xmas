// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/beam19857/playlist-video/ent/playlist_video"
	"github.com/beam19857/playlist-video/ent/user"
	"github.com/beam19857/playlist-video/ent/video"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetName sets the name field.
func (vc *VideoCreate) SetName(s string) *VideoCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetURL sets the url field.
func (vc *VideoCreate) SetURL(s string) *VideoCreate {
	vc.mutation.SetURL(s)
	return vc
}

// SetOwnerID sets the owner edge to User by id.
func (vc *VideoCreate) SetOwnerID(id int) *VideoCreate {
	vc.mutation.SetOwnerID(id)
	return vc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (vc *VideoCreate) SetNillableOwnerID(id *int) *VideoCreate {
	if id != nil {
		vc = vc.SetOwnerID(*id)
	}
	return vc
}

// SetOwner sets the owner edge to User.
func (vc *VideoCreate) SetOwner(u *User) *VideoCreate {
	return vc.SetOwnerID(u.ID)
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (vc *VideoCreate) AddPlaylistVideoIDs(ids ...int) *VideoCreate {
	vc.mutation.AddPlaylistVideoIDs(ids...)
	return vc
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (vc *VideoCreate) AddPlaylistVideos(p ...*Playlist_Video) *VideoCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vc.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	if err := vc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Video
	)
	if len(vc.hooks) == 0 {
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vc *VideoCreate) preSave() error {
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := vc.mutation.Name(); ok {
		if err := video.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := vc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if v, ok := vc.mutation.URL(); ok {
		if err := video.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	v, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	v.ID = int(id)
	return v, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		v     = &Video{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: video.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldName,
		})
		v.Name = value
	}
	if value, ok := vc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldURL,
		})
		v.URL = value
	}
	if nodes := vc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.OwnerTable,
			Columns: []string{video.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaylistVideosTable,
			Columns: []string{video.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return v, _spec
}

// VideoCreateBulk is the builder for creating a bulk of Video entities.
type VideoCreateBulk struct {
	config
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*VideoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
