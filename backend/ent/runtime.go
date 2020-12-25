// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/beam19857/playlist-video/ent/playlist"
	"github.com/beam19857/playlist-video/ent/resolution"
	"github.com/beam19857/playlist-video/ent/schema"
	"github.com/beam19857/playlist-video/ent/user"
	"github.com/beam19857/playlist-video/ent/video"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	playlistFields := schema.Playlist{}.Fields()
	_ = playlistFields
	// playlistDescTitle is the schema descriptor for title field.
	playlistDescTitle := playlistFields[0].Descriptor()
	// playlist.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	playlist.TitleValidator = playlistDescTitle.Validators[0].(func(string) error)
	resolutionFields := schema.Resolution{}.Fields()
	_ = resolutionFields
	// resolutionDescValue is the schema descriptor for value field.
	resolutionDescValue := resolutionFields[0].Descriptor()
	// resolution.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	resolution.ValueValidator = resolutionDescValue.Validators[0].(func(int) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescName is the schema descriptor for name field.
	videoDescName := videoFields[0].Descriptor()
	// video.NameValidator is a validator for the "name" field. It is called by the builders before save.
	video.NameValidator = videoDescName.Validators[0].(func(string) error)
	// videoDescURL is the schema descriptor for url field.
	videoDescURL := videoFields[1].Descriptor()
	// video.URLValidator is a validator for the "url" field. It is called by the builders before save.
	video.URLValidator = videoDescURL.Validators[0].(func(string) error)
}
