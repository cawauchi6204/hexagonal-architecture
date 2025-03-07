// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// ThreadTag is an object representing the database table.
type ThreadTag struct {
	ThreadID  string    `boil:"thread_id" json:"thread_id" toml:"thread_id" yaml:"thread_id"`
	TagID     string    `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *threadTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L threadTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ThreadTagColumns = struct {
	ThreadID  string
	TagID     string
	CreatedAt string
}{
	ThreadID:  "thread_id",
	TagID:     "tag_id",
	CreatedAt: "created_at",
}

var ThreadTagTableColumns = struct {
	ThreadID  string
	TagID     string
	CreatedAt string
}{
	ThreadID:  "thread_tags.thread_id",
	TagID:     "thread_tags.tag_id",
	CreatedAt: "thread_tags.created_at",
}

// Generated where

var ThreadTagWhere = struct {
	ThreadID  whereHelperstring
	TagID     whereHelperstring
	CreatedAt whereHelpernull_Time
}{
	ThreadID:  whereHelperstring{field: "`thread_tags`.`thread_id`"},
	TagID:     whereHelperstring{field: "`thread_tags`.`tag_id`"},
	CreatedAt: whereHelpernull_Time{field: "`thread_tags`.`created_at`"},
}

// ThreadTagRels is where relationship names are stored.
var ThreadTagRels = struct {
	Thread string
	Tag    string
}{
	Thread: "Thread",
	Tag:    "Tag",
}

// threadTagR is where relationships are stored.
type threadTagR struct {
	Thread *Thread `boil:"Thread" json:"Thread" toml:"Thread" yaml:"Thread"`
	Tag    *Tag    `boil:"Tag" json:"Tag" toml:"Tag" yaml:"Tag"`
}

// NewStruct creates a new relationship struct
func (*threadTagR) NewStruct() *threadTagR {
	return &threadTagR{}
}

func (r *threadTagR) GetThread() *Thread {
	if r == nil {
		return nil
	}
	return r.Thread
}

func (r *threadTagR) GetTag() *Tag {
	if r == nil {
		return nil
	}
	return r.Tag
}

// threadTagL is where Load methods for each relationship are stored.
type threadTagL struct{}

var (
	threadTagAllColumns            = []string{"thread_id", "tag_id", "created_at"}
	threadTagColumnsWithoutDefault = []string{"thread_id", "tag_id"}
	threadTagColumnsWithDefault    = []string{"created_at"}
	threadTagPrimaryKeyColumns     = []string{"thread_id", "tag_id"}
	threadTagGeneratedColumns      = []string{}
)

type (
	// ThreadTagSlice is an alias for a slice of pointers to ThreadTag.
	// This should almost always be used instead of []ThreadTag.
	ThreadTagSlice []*ThreadTag
	// ThreadTagHook is the signature for custom ThreadTag hook methods
	ThreadTagHook func(context.Context, boil.ContextExecutor, *ThreadTag) error

	threadTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	threadTagType                 = reflect.TypeOf(&ThreadTag{})
	threadTagMapping              = queries.MakeStructMapping(threadTagType)
	threadTagPrimaryKeyMapping, _ = queries.BindMapping(threadTagType, threadTagMapping, threadTagPrimaryKeyColumns)
	threadTagInsertCacheMut       sync.RWMutex
	threadTagInsertCache          = make(map[string]insertCache)
	threadTagUpdateCacheMut       sync.RWMutex
	threadTagUpdateCache          = make(map[string]updateCache)
	threadTagUpsertCacheMut       sync.RWMutex
	threadTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var threadTagAfterSelectMu sync.Mutex
var threadTagAfterSelectHooks []ThreadTagHook

var threadTagBeforeInsertMu sync.Mutex
var threadTagBeforeInsertHooks []ThreadTagHook
var threadTagAfterInsertMu sync.Mutex
var threadTagAfterInsertHooks []ThreadTagHook

var threadTagBeforeUpdateMu sync.Mutex
var threadTagBeforeUpdateHooks []ThreadTagHook
var threadTagAfterUpdateMu sync.Mutex
var threadTagAfterUpdateHooks []ThreadTagHook

var threadTagBeforeDeleteMu sync.Mutex
var threadTagBeforeDeleteHooks []ThreadTagHook
var threadTagAfterDeleteMu sync.Mutex
var threadTagAfterDeleteHooks []ThreadTagHook

var threadTagBeforeUpsertMu sync.Mutex
var threadTagBeforeUpsertHooks []ThreadTagHook
var threadTagAfterUpsertMu sync.Mutex
var threadTagAfterUpsertHooks []ThreadTagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ThreadTag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ThreadTag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ThreadTag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ThreadTag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ThreadTag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ThreadTag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ThreadTag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ThreadTag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ThreadTag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range threadTagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddThreadTagHook registers your hook function for all future operations.
func AddThreadTagHook(hookPoint boil.HookPoint, threadTagHook ThreadTagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		threadTagAfterSelectMu.Lock()
		threadTagAfterSelectHooks = append(threadTagAfterSelectHooks, threadTagHook)
		threadTagAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		threadTagBeforeInsertMu.Lock()
		threadTagBeforeInsertHooks = append(threadTagBeforeInsertHooks, threadTagHook)
		threadTagBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		threadTagAfterInsertMu.Lock()
		threadTagAfterInsertHooks = append(threadTagAfterInsertHooks, threadTagHook)
		threadTagAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		threadTagBeforeUpdateMu.Lock()
		threadTagBeforeUpdateHooks = append(threadTagBeforeUpdateHooks, threadTagHook)
		threadTagBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		threadTagAfterUpdateMu.Lock()
		threadTagAfterUpdateHooks = append(threadTagAfterUpdateHooks, threadTagHook)
		threadTagAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		threadTagBeforeDeleteMu.Lock()
		threadTagBeforeDeleteHooks = append(threadTagBeforeDeleteHooks, threadTagHook)
		threadTagBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		threadTagAfterDeleteMu.Lock()
		threadTagAfterDeleteHooks = append(threadTagAfterDeleteHooks, threadTagHook)
		threadTagAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		threadTagBeforeUpsertMu.Lock()
		threadTagBeforeUpsertHooks = append(threadTagBeforeUpsertHooks, threadTagHook)
		threadTagBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		threadTagAfterUpsertMu.Lock()
		threadTagAfterUpsertHooks = append(threadTagAfterUpsertHooks, threadTagHook)
		threadTagAfterUpsertMu.Unlock()
	}
}

// One returns a single threadTag record from the query.
func (q threadTagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ThreadTag, error) {
	o := &ThreadTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schemas: failed to execute a one query for thread_tags")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ThreadTag records from the query.
func (q threadTagQuery) All(ctx context.Context, exec boil.ContextExecutor) (ThreadTagSlice, error) {
	var o []*ThreadTag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schemas: failed to assign all query results to ThreadTag slice")
	}

	if len(threadTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ThreadTag records in the query.
func (q threadTagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to count thread_tags rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q threadTagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schemas: failed to check if thread_tags exists")
	}

	return count > 0, nil
}

// Thread pointed to by the foreign key.
func (o *ThreadTag) Thread(mods ...qm.QueryMod) threadQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ThreadID),
	}

	queryMods = append(queryMods, mods...)

	return Threads(queryMods...)
}

// Tag pointed to by the foreign key.
func (o *ThreadTag) Tag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.TagID),
	}

	queryMods = append(queryMods, mods...)

	return Tags(queryMods...)
}

// LoadThread allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (threadTagL) LoadThread(ctx context.Context, e boil.ContextExecutor, singular bool, maybeThreadTag interface{}, mods queries.Applicator) error {
	var slice []*ThreadTag
	var object *ThreadTag

	if singular {
		var ok bool
		object, ok = maybeThreadTag.(*ThreadTag)
		if !ok {
			object = new(ThreadTag)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeThreadTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeThreadTag))
			}
		}
	} else {
		s, ok := maybeThreadTag.(*[]*ThreadTag)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeThreadTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeThreadTag))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &threadTagR{}
		}
		args[object.ThreadID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &threadTagR{}
			}

			args[obj.ThreadID] = struct{}{}

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
		qm.From(`threads`),
		qm.WhereIn(`threads.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Thread")
	}

	var resultSlice []*Thread
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Thread")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for threads")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for threads")
	}

	if len(threadAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Thread = foreign
		if foreign.R == nil {
			foreign.R = &threadR{}
		}
		foreign.R.ThreadTags = append(foreign.R.ThreadTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ThreadID == foreign.ID {
				local.R.Thread = foreign
				if foreign.R == nil {
					foreign.R = &threadR{}
				}
				foreign.R.ThreadTags = append(foreign.R.ThreadTags, local)
				break
			}
		}
	}

	return nil
}

// LoadTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (threadTagL) LoadTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeThreadTag interface{}, mods queries.Applicator) error {
	var slice []*ThreadTag
	var object *ThreadTag

	if singular {
		var ok bool
		object, ok = maybeThreadTag.(*ThreadTag)
		if !ok {
			object = new(ThreadTag)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeThreadTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeThreadTag))
			}
		}
	} else {
		s, ok := maybeThreadTag.(*[]*ThreadTag)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeThreadTag)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeThreadTag))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &threadTagR{}
		}
		args[object.TagID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &threadTagR{}
			}

			args[obj.TagID] = struct{}{}

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
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(tagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.ThreadTags = append(foreign.R.ThreadTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TagID == foreign.ID {
				local.R.Tag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.ThreadTags = append(foreign.R.ThreadTags, local)
				break
			}
		}
	}

	return nil
}

// SetThread of the threadTag to the related item.
// Sets o.R.Thread to related.
// Adds o to related.R.ThreadTags.
func (o *ThreadTag) SetThread(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Thread) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `thread_tags` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"thread_id"}),
		strmangle.WhereClause("`", "`", 0, threadTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ThreadID, o.TagID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ThreadID = related.ID
	if o.R == nil {
		o.R = &threadTagR{
			Thread: related,
		}
	} else {
		o.R.Thread = related
	}

	if related.R == nil {
		related.R = &threadR{
			ThreadTags: ThreadTagSlice{o},
		}
	} else {
		related.R.ThreadTags = append(related.R.ThreadTags, o)
	}

	return nil
}

// SetTag of the threadTag to the related item.
// Sets o.R.Tag to related.
// Adds o to related.R.ThreadTags.
func (o *ThreadTag) SetTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `thread_tags` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"tag_id"}),
		strmangle.WhereClause("`", "`", 0, threadTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ThreadID, o.TagID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TagID = related.ID
	if o.R == nil {
		o.R = &threadTagR{
			Tag: related,
		}
	} else {
		o.R.Tag = related
	}

	if related.R == nil {
		related.R = &tagR{
			ThreadTags: ThreadTagSlice{o},
		}
	} else {
		related.R.ThreadTags = append(related.R.ThreadTags, o)
	}

	return nil
}

// ThreadTags retrieves all the records using an executor.
func ThreadTags(mods ...qm.QueryMod) threadTagQuery {
	mods = append(mods, qm.From("`thread_tags`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`thread_tags`.*"})
	}

	return threadTagQuery{q}
}

// FindThreadTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindThreadTag(ctx context.Context, exec boil.ContextExecutor, threadID string, tagID string, selectCols ...string) (*ThreadTag, error) {
	threadTagObj := &ThreadTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `thread_tags` where `thread_id`=? AND `tag_id`=?", sel,
	)

	q := queries.Raw(query, threadID, tagID)

	err := q.Bind(ctx, exec, threadTagObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schemas: unable to select from thread_tags")
	}

	if err = threadTagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return threadTagObj, err
	}

	return threadTagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ThreadTag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schemas: no thread_tags provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(threadTagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	threadTagInsertCacheMut.RLock()
	cache, cached := threadTagInsertCache[key]
	threadTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			threadTagAllColumns,
			threadTagColumnsWithDefault,
			threadTagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(threadTagType, threadTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(threadTagType, threadTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `thread_tags` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `thread_tags` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `thread_tags` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, threadTagPrimaryKeyColumns))
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
		return errors.Wrap(err, "schemas: unable to insert into thread_tags")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ThreadID,
		o.TagID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to populate default values for thread_tags")
	}

CacheNoHooks:
	if !cached {
		threadTagInsertCacheMut.Lock()
		threadTagInsertCache[key] = cache
		threadTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ThreadTag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ThreadTag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	threadTagUpdateCacheMut.RLock()
	cache, cached := threadTagUpdateCache[key]
	threadTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			threadTagAllColumns,
			threadTagPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("schemas: unable to update thread_tags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `thread_tags` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, threadTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(threadTagType, threadTagMapping, append(wl, threadTagPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schemas: unable to update thread_tags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by update for thread_tags")
	}

	if !cached {
		threadTagUpdateCacheMut.Lock()
		threadTagUpdateCache[key] = cache
		threadTagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q threadTagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update all for thread_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to retrieve rows affected for thread_tags")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ThreadTagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), threadTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `thread_tags` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, threadTagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update all in threadTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to retrieve rows affected all in update all threadTag")
	}
	return rowsAff, nil
}

var mySQLThreadTagUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ThreadTag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schemas: no thread_tags provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(threadTagColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLThreadTagUniqueColumns, o)

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

	threadTagUpsertCacheMut.RLock()
	cache, cached := threadTagUpsertCache[key]
	threadTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			threadTagAllColumns,
			threadTagColumnsWithDefault,
			threadTagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			threadTagAllColumns,
			threadTagPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("schemas: unable to upsert thread_tags, could not build update column list")
		}

		ret := strmangle.SetComplement(threadTagAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`thread_tags`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `thread_tags` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(threadTagType, threadTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(threadTagType, threadTagMapping, ret)
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
		return errors.Wrap(err, "schemas: unable to upsert for thread_tags")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(threadTagType, threadTagMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to retrieve unique values for thread_tags")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to populate default values for thread_tags")
	}

CacheNoHooks:
	if !cached {
		threadTagUpsertCacheMut.Lock()
		threadTagUpsertCache[key] = cache
		threadTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ThreadTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ThreadTag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("schemas: no ThreadTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), threadTagPrimaryKeyMapping)
	sql := "DELETE FROM `thread_tags` WHERE `thread_id`=? AND `tag_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete from thread_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by delete for thread_tags")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q threadTagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schemas: no threadTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete all from thread_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by deleteall for thread_tags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ThreadTagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(threadTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), threadTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `thread_tags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, threadTagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete all from threadTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by deleteall for thread_tags")
	}

	if len(threadTagAfterDeleteHooks) != 0 {
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
func (o *ThreadTag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindThreadTag(ctx, exec, o.ThreadID, o.TagID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ThreadTagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ThreadTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), threadTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `thread_tags`.* FROM `thread_tags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, threadTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to reload all in ThreadTagSlice")
	}

	*o = slice

	return nil
}

// ThreadTagExists checks if the ThreadTag row exists.
func ThreadTagExists(ctx context.Context, exec boil.ContextExecutor, threadID string, tagID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `thread_tags` where `thread_id`=? AND `tag_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, threadID, tagID)
	}
	row := exec.QueryRowContext(ctx, sql, threadID, tagID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schemas: unable to check if thread_tags exists")
	}

	return exists, nil
}

// Exists checks if the ThreadTag row exists.
func (o *ThreadTag) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ThreadTagExists(ctx, exec, o.ThreadID, o.TagID)
}
