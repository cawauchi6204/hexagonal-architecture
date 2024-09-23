// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schemas

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Tag is an object representing the database table.
type Tag struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *tagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TagColumns = struct {
	ID        string
	Name      string
	CreatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
}

var TagTableColumns = struct {
	ID        string
	Name      string
	CreatedAt string
}{
	ID:        "tags.id",
	Name:      "tags.name",
	CreatedAt: "tags.created_at",
}

// Generated where

var TagWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	CreatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "`tags`.`id`"},
	Name:      whereHelperstring{field: "`tags`.`name`"},
	CreatedAt: whereHelpernull_Time{field: "`tags`.`created_at`"},
}

// TagRels is where relationship names are stored.
var TagRels = struct {
	ThreadTags string
	UsersTags  string
}{
	ThreadTags: "ThreadTags",
	UsersTags:  "UsersTags",
}

// tagR is where relationships are stored.
type tagR struct {
	ThreadTags ThreadTagSlice `boil:"ThreadTags" json:"ThreadTags" toml:"ThreadTags" yaml:"ThreadTags"`
	UsersTags  UsersTagSlice  `boil:"UsersTags" json:"UsersTags" toml:"UsersTags" yaml:"UsersTags"`
}

// NewStruct creates a new relationship struct
func (*tagR) NewStruct() *tagR {
	return &tagR{}
}

func (r *tagR) GetThreadTags() ThreadTagSlice {
	if r == nil {
		return nil
	}
	return r.ThreadTags
}

func (r *tagR) GetUsersTags() UsersTagSlice {
	if r == nil {
		return nil
	}
	return r.UsersTags
}

// tagL is where Load methods for each relationship are stored.
type tagL struct{}

var (
	tagAllColumns            = []string{"id", "name", "created_at"}
	tagColumnsWithoutDefault = []string{"id", "name"}
	tagColumnsWithDefault    = []string{"created_at"}
	tagPrimaryKeyColumns     = []string{"id"}
	tagGeneratedColumns      = []string{}
)

type (
	// TagSlice is an alias for a slice of pointers to Tag.
	// This should almost always be used instead of []Tag.
	TagSlice []*Tag
	// TagHook is the signature for custom Tag hook methods
	TagHook func(context.Context, boil.ContextExecutor, *Tag) error

	tagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagType                 = reflect.TypeOf(&Tag{})
	tagMapping              = queries.MakeStructMapping(tagType)
	tagPrimaryKeyMapping, _ = queries.BindMapping(tagType, tagMapping, tagPrimaryKeyColumns)
	tagInsertCacheMut       sync.RWMutex
	tagInsertCache          = make(map[string]insertCache)
	tagUpdateCacheMut       sync.RWMutex
	tagUpdateCache          = make(map[string]updateCache)
	tagUpsertCacheMut       sync.RWMutex
	tagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tagAfterSelectMu sync.Mutex
var tagAfterSelectHooks []TagHook

var tagBeforeInsertMu sync.Mutex
var tagBeforeInsertHooks []TagHook
var tagAfterInsertMu sync.Mutex
var tagAfterInsertHooks []TagHook

var tagBeforeUpdateMu sync.Mutex
var tagBeforeUpdateHooks []TagHook
var tagAfterUpdateMu sync.Mutex
var tagAfterUpdateHooks []TagHook

var tagBeforeDeleteMu sync.Mutex
var tagBeforeDeleteHooks []TagHook
var tagAfterDeleteMu sync.Mutex
var tagAfterDeleteHooks []TagHook

var tagBeforeUpsertMu sync.Mutex
var tagBeforeUpsertHooks []TagHook
var tagAfterUpsertMu sync.Mutex
var tagAfterUpsertHooks []TagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTagHook registers your hook function for all future operations.
func AddTagHook(hookPoint boil.HookPoint, tagHook TagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tagAfterSelectMu.Lock()
		tagAfterSelectHooks = append(tagAfterSelectHooks, tagHook)
		tagAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tagBeforeInsertMu.Lock()
		tagBeforeInsertHooks = append(tagBeforeInsertHooks, tagHook)
		tagBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tagAfterInsertMu.Lock()
		tagAfterInsertHooks = append(tagAfterInsertHooks, tagHook)
		tagAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tagBeforeUpdateMu.Lock()
		tagBeforeUpdateHooks = append(tagBeforeUpdateHooks, tagHook)
		tagBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tagAfterUpdateMu.Lock()
		tagAfterUpdateHooks = append(tagAfterUpdateHooks, tagHook)
		tagAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tagBeforeDeleteMu.Lock()
		tagBeforeDeleteHooks = append(tagBeforeDeleteHooks, tagHook)
		tagBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tagAfterDeleteMu.Lock()
		tagAfterDeleteHooks = append(tagAfterDeleteHooks, tagHook)
		tagAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tagBeforeUpsertMu.Lock()
		tagBeforeUpsertHooks = append(tagBeforeUpsertHooks, tagHook)
		tagBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tagAfterUpsertMu.Lock()
		tagAfterUpsertHooks = append(tagAfterUpsertHooks, tagHook)
		tagAfterUpsertMu.Unlock()
	}
}

// One returns a single tag record from the query.
func (q tagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Tag, error) {
	o := &Tag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schemas: failed to execute a one query for tags")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Tag records from the query.
func (q tagQuery) All(ctx context.Context, exec boil.ContextExecutor) (TagSlice, error) {
	var o []*Tag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schemas: failed to assign all query results to Tag slice")
	}

	if len(tagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Tag records in the query.
func (q tagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to count tags rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schemas: failed to check if tags exists")
	}

	return count > 0, nil
}

// ThreadTags retrieves all the thread_tag's ThreadTags with an executor.
func (o *Tag) ThreadTags(mods ...qm.QueryMod) threadTagQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`thread_tags`.`tag_id`=?", o.ID),
	)

	return ThreadTags(queryMods...)
}

// UsersTags retrieves all the users_tag's UsersTags with an executor.
func (o *Tag) UsersTags(mods ...qm.QueryMod) usersTagQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`users_tags`.`tag_id`=?", o.ID),
	)

	return UsersTags(queryMods...)
}

// LoadThreadTags allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tagL) LoadThreadTags(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTag interface{}, mods queries.Applicator) error {
	var slice []*Tag
	var object *Tag

	if singular {
		var ok bool
		object, ok = maybeTag.(*Tag)
		if !ok {
			object = new(Tag)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTag))
			}
		}
	} else {
		s, ok := maybeTag.(*[]*Tag)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTag))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &tagR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`thread_tags`),
		qm.WhereIn(`thread_tags.tag_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load thread_tags")
	}

	var resultSlice []*ThreadTag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice thread_tags")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on thread_tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for thread_tags")
	}

	if len(threadTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ThreadTags = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &threadTagR{}
			}
			foreign.R.Tag = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TagID {
				local.R.ThreadTags = append(local.R.ThreadTags, foreign)
				if foreign.R == nil {
					foreign.R = &threadTagR{}
				}
				foreign.R.Tag = local
				break
			}
		}
	}

	return nil
}

// LoadUsersTags allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tagL) LoadUsersTags(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTag interface{}, mods queries.Applicator) error {
	var slice []*Tag
	var object *Tag

	if singular {
		var ok bool
		object, ok = maybeTag.(*Tag)
		if !ok {
			object = new(Tag)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTag))
			}
		}
	} else {
		s, ok := maybeTag.(*[]*Tag)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTag))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &tagR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tagR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users_tags`),
		qm.WhereIn(`users_tags.tag_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users_tags")
	}

	var resultSlice []*UsersTag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users_tags")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users_tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users_tags")
	}

	if len(usersTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UsersTags = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &usersTagR{}
			}
			foreign.R.Tag = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TagID {
				local.R.UsersTags = append(local.R.UsersTags, foreign)
				if foreign.R == nil {
					foreign.R = &usersTagR{}
				}
				foreign.R.Tag = local
				break
			}
		}
	}

	return nil
}

// AddThreadTags adds the given related objects to the existing relationships
// of the tag, optionally inserting them as new records.
// Appends related to o.R.ThreadTags.
// Sets related.R.Tag appropriately.
func (o *Tag) AddThreadTags(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ThreadTag) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TagID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `thread_tags` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"tag_id"}),
				strmangle.WhereClause("`", "`", 0, threadTagPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ThreadID, rel.TagID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TagID = o.ID
		}
	}

	if o.R == nil {
		o.R = &tagR{
			ThreadTags: related,
		}
	} else {
		o.R.ThreadTags = append(o.R.ThreadTags, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &threadTagR{
				Tag: o,
			}
		} else {
			rel.R.Tag = o
		}
	}
	return nil
}

// AddUsersTags adds the given related objects to the existing relationships
// of the tag, optionally inserting them as new records.
// Appends related to o.R.UsersTags.
// Sets related.R.Tag appropriately.
func (o *Tag) AddUsersTags(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UsersTag) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TagID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `users_tags` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"tag_id"}),
				strmangle.WhereClause("`", "`", 0, usersTagPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.UserID, rel.TagID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TagID = o.ID
		}
	}

	if o.R == nil {
		o.R = &tagR{
			UsersTags: related,
		}
	} else {
		o.R.UsersTags = append(o.R.UsersTags, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &usersTagR{
				Tag: o,
			}
		} else {
			rel.R.Tag = o
		}
	}
	return nil
}

// Tags retrieves all the records using an executor.
func Tags(mods ...qm.QueryMod) tagQuery {
	mods = append(mods, qm.From("`tags`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`tags`.*"})
	}

	return tagQuery{q}
}

// FindTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTag(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Tag, error) {
	tagObj := &Tag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tags` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tagObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schemas: unable to select from tags")
	}

	if err = tagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tagObj, err
	}

	return tagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schemas: no tags provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagInsertCacheMut.RLock()
	cache, cached := tagInsertCache[key]
	tagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagType, tagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tags` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tags` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tags` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tagPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "schemas: unable to insert into tags")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to populate default values for tags")
	}

CacheNoHooks:
	if !cached {
		tagInsertCacheMut.Lock()
		tagInsertCache[key] = cache
		tagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Tag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tagUpdateCacheMut.RLock()
	cache, cached := tagUpdateCache[key]
	tagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("schemas: unable to update tags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tags` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, append(wl, tagPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update tags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by update for tags")
	}

	if !cached {
		tagUpdateCacheMut.Lock()
		tagUpdateCache[key] = cache
		tagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update all for tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to retrieve rows affected for tags")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("schemas: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tags` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update all in tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to retrieve rows affected all in update all tag")
	}
	return rowsAff, nil
}

var mySQLTagUniqueColumns = []string{
	"id",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schemas: no tags provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTagUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tagUpsertCacheMut.RLock()
	cache, cached := tagUpsertCache[key]
	tagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("schemas: unable to upsert tags, could not build update column list")
		}

		ret := strmangle.SetComplement(tagAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`tags`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tags` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagType, tagMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "schemas: unable to upsert for tags")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tagType, tagMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to retrieve unique values for tags")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to populate default values for tags")
	}

CacheNoHooks:
	if !cached {
		tagUpsertCacheMut.Lock()
		tagUpsertCache[key] = cache
		tagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Tag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("schemas: no Tag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagPrimaryKeyMapping)
	sql := "DELETE FROM `tags` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete from tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by delete for tags")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schemas: no tagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete all from tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by deleteall for tags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete all from tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by deleteall for tags")
	}

	if len(tagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tags`.* FROM `tags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to reload all in TagSlice")
	}

	*o = slice

	return nil
}

// TagExists checks if the Tag row exists.
func TagExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tags` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schemas: unable to check if tags exists")
	}

	return exists, nil
}

// Exists checks if the Tag row exists.
func (o *Tag) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TagExists(ctx, exec, o.ID)
}
