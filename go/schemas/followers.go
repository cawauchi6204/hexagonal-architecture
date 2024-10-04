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

// Follower is an object representing the database table.
type Follower struct {
	FollowerID string    `boil:"follower_id" json:"follower_id" toml:"follower_id" yaml:"follower_id"`
	FollowedID string    `boil:"followed_id" json:"followed_id" toml:"followed_id" yaml:"followed_id"`
	CreatedAt  null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *followerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L followerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FollowerColumns = struct {
	FollowerID string
	FollowedID string
	CreatedAt  string
}{
	FollowerID: "follower_id",
	FollowedID: "followed_id",
	CreatedAt:  "created_at",
}

var FollowerTableColumns = struct {
	FollowerID string
	FollowedID string
	CreatedAt  string
}{
	FollowerID: "followers.follower_id",
	FollowedID: "followers.followed_id",
	CreatedAt:  "followers.created_at",
}

// Generated where

var FollowerWhere = struct {
	FollowerID whereHelperstring
	FollowedID whereHelperstring
	CreatedAt  whereHelpernull_Time
}{
	FollowerID: whereHelperstring{field: "`followers`.`follower_id`"},
	FollowedID: whereHelperstring{field: "`followers`.`followed_id`"},
	CreatedAt:  whereHelpernull_Time{field: "`followers`.`created_at`"},
}

// FollowerRels is where relationship names are stored.
var FollowerRels = struct {
	Follower string
	Followed string
}{
	Follower: "Follower",
	Followed: "Followed",
}

// followerR is where relationships are stored.
type followerR struct {
	Follower *User `boil:"Follower" json:"Follower" toml:"Follower" yaml:"Follower"`
	Followed *User `boil:"Followed" json:"Followed" toml:"Followed" yaml:"Followed"`
}

// NewStruct creates a new relationship struct
func (*followerR) NewStruct() *followerR {
	return &followerR{}
}

func (r *followerR) GetFollower() *User {
	if r == nil {
		return nil
	}
	return r.Follower
}

func (r *followerR) GetFollowed() *User {
	if r == nil {
		return nil
	}
	return r.Followed
}

// followerL is where Load methods for each relationship are stored.
type followerL struct{}

var (
	followerAllColumns            = []string{"follower_id", "followed_id", "created_at"}
	followerColumnsWithoutDefault = []string{"follower_id", "followed_id"}
	followerColumnsWithDefault    = []string{"created_at"}
	followerPrimaryKeyColumns     = []string{"follower_id", "followed_id"}
	followerGeneratedColumns      = []string{}
)

type (
	// FollowerSlice is an alias for a slice of pointers to Follower.
	// This should almost always be used instead of []Follower.
	FollowerSlice []*Follower
	// FollowerHook is the signature for custom Follower hook methods
	FollowerHook func(context.Context, boil.ContextExecutor, *Follower) error

	followerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	followerType                 = reflect.TypeOf(&Follower{})
	followerMapping              = queries.MakeStructMapping(followerType)
	followerPrimaryKeyMapping, _ = queries.BindMapping(followerType, followerMapping, followerPrimaryKeyColumns)
	followerInsertCacheMut       sync.RWMutex
	followerInsertCache          = make(map[string]insertCache)
	followerUpdateCacheMut       sync.RWMutex
	followerUpdateCache          = make(map[string]updateCache)
	followerUpsertCacheMut       sync.RWMutex
	followerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var followerAfterSelectMu sync.Mutex
var followerAfterSelectHooks []FollowerHook

var followerBeforeInsertMu sync.Mutex
var followerBeforeInsertHooks []FollowerHook
var followerAfterInsertMu sync.Mutex
var followerAfterInsertHooks []FollowerHook

var followerBeforeUpdateMu sync.Mutex
var followerBeforeUpdateHooks []FollowerHook
var followerAfterUpdateMu sync.Mutex
var followerAfterUpdateHooks []FollowerHook

var followerBeforeDeleteMu sync.Mutex
var followerBeforeDeleteHooks []FollowerHook
var followerAfterDeleteMu sync.Mutex
var followerAfterDeleteHooks []FollowerHook

var followerBeforeUpsertMu sync.Mutex
var followerBeforeUpsertHooks []FollowerHook
var followerAfterUpsertMu sync.Mutex
var followerAfterUpsertHooks []FollowerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Follower) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Follower) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Follower) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Follower) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Follower) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Follower) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Follower) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Follower) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Follower) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFollowerHook registers your hook function for all future operations.
func AddFollowerHook(hookPoint boil.HookPoint, followerHook FollowerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		followerAfterSelectMu.Lock()
		followerAfterSelectHooks = append(followerAfterSelectHooks, followerHook)
		followerAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		followerBeforeInsertMu.Lock()
		followerBeforeInsertHooks = append(followerBeforeInsertHooks, followerHook)
		followerBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		followerAfterInsertMu.Lock()
		followerAfterInsertHooks = append(followerAfterInsertHooks, followerHook)
		followerAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		followerBeforeUpdateMu.Lock()
		followerBeforeUpdateHooks = append(followerBeforeUpdateHooks, followerHook)
		followerBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		followerAfterUpdateMu.Lock()
		followerAfterUpdateHooks = append(followerAfterUpdateHooks, followerHook)
		followerAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		followerBeforeDeleteMu.Lock()
		followerBeforeDeleteHooks = append(followerBeforeDeleteHooks, followerHook)
		followerBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		followerAfterDeleteMu.Lock()
		followerAfterDeleteHooks = append(followerAfterDeleteHooks, followerHook)
		followerAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		followerBeforeUpsertMu.Lock()
		followerBeforeUpsertHooks = append(followerBeforeUpsertHooks, followerHook)
		followerBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		followerAfterUpsertMu.Lock()
		followerAfterUpsertHooks = append(followerAfterUpsertHooks, followerHook)
		followerAfterUpsertMu.Unlock()
	}
}

// One returns a single follower record from the query.
func (q followerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Follower, error) {
	o := &Follower{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schemas: failed to execute a one query for followers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Follower records from the query.
func (q followerQuery) All(ctx context.Context, exec boil.ContextExecutor) (FollowerSlice, error) {
	var o []*Follower

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schemas: failed to assign all query results to Follower slice")
	}

	if len(followerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Follower records in the query.
func (q followerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to count followers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q followerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schemas: failed to check if followers exists")
	}

	return count > 0, nil
}

// Follower pointed to by the foreign key.
func (o *Follower) Follower(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.FollowerID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Followed pointed to by the foreign key.
func (o *Follower) Followed(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.FollowedID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadFollower allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (followerL) LoadFollower(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFollower interface{}, mods queries.Applicator) error {
	var slice []*Follower
	var object *Follower

	if singular {
		var ok bool
		object, ok = maybeFollower.(*Follower)
		if !ok {
			object = new(Follower)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeFollower)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeFollower))
			}
		}
	} else {
		s, ok := maybeFollower.(*[]*Follower)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeFollower)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeFollower))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &followerR{}
		}
		args[object.FollowerID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &followerR{}
			}

			args[obj.FollowerID] = struct{}{}

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
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.Follower = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.FollowerFollowers = append(foreign.R.FollowerFollowers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FollowerID == foreign.ID {
				local.R.Follower = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.FollowerFollowers = append(foreign.R.FollowerFollowers, local)
				break
			}
		}
	}

	return nil
}

// LoadFollowed allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (followerL) LoadFollowed(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFollower interface{}, mods queries.Applicator) error {
	var slice []*Follower
	var object *Follower

	if singular {
		var ok bool
		object, ok = maybeFollower.(*Follower)
		if !ok {
			object = new(Follower)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeFollower)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeFollower))
			}
		}
	} else {
		s, ok := maybeFollower.(*[]*Follower)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeFollower)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeFollower))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &followerR{}
		}
		args[object.FollowedID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &followerR{}
			}

			args[obj.FollowedID] = struct{}{}

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
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.Followed = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.FollowedFollowers = append(foreign.R.FollowedFollowers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FollowedID == foreign.ID {
				local.R.Followed = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.FollowedFollowers = append(foreign.R.FollowedFollowers, local)
				break
			}
		}
	}

	return nil
}

// SetFollower of the follower to the related item.
// Sets o.R.Follower to related.
// Adds o to related.R.FollowerFollowers.
func (o *Follower) SetFollower(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `followers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"follower_id"}),
		strmangle.WhereClause("`", "`", 0, followerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.FollowerID, o.FollowedID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FollowerID = related.ID
	if o.R == nil {
		o.R = &followerR{
			Follower: related,
		}
	} else {
		o.R.Follower = related
	}

	if related.R == nil {
		related.R = &userR{
			FollowerFollowers: FollowerSlice{o},
		}
	} else {
		related.R.FollowerFollowers = append(related.R.FollowerFollowers, o)
	}

	return nil
}

// SetFollowed of the follower to the related item.
// Sets o.R.Followed to related.
// Adds o to related.R.FollowedFollowers.
func (o *Follower) SetFollowed(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `followers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"followed_id"}),
		strmangle.WhereClause("`", "`", 0, followerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.FollowerID, o.FollowedID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FollowedID = related.ID
	if o.R == nil {
		o.R = &followerR{
			Followed: related,
		}
	} else {
		o.R.Followed = related
	}

	if related.R == nil {
		related.R = &userR{
			FollowedFollowers: FollowerSlice{o},
		}
	} else {
		related.R.FollowedFollowers = append(related.R.FollowedFollowers, o)
	}

	return nil
}

// Followers retrieves all the records using an executor.
func Followers(mods ...qm.QueryMod) followerQuery {
	mods = append(mods, qm.From("`followers`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`followers`.*"})
	}

	return followerQuery{q}
}

// FindFollower retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFollower(ctx context.Context, exec boil.ContextExecutor, followerID string, followedID string, selectCols ...string) (*Follower, error) {
	followerObj := &Follower{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `followers` where `follower_id`=? AND `followed_id`=?", sel,
	)

	q := queries.Raw(query, followerID, followedID)

	err := q.Bind(ctx, exec, followerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schemas: unable to select from followers")
	}

	if err = followerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return followerObj, err
	}

	return followerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Follower) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schemas: no followers provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(followerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	followerInsertCacheMut.RLock()
	cache, cached := followerInsertCache[key]
	followerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			followerAllColumns,
			followerColumnsWithDefault,
			followerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(followerType, followerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(followerType, followerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `followers` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `followers` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `followers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, followerPrimaryKeyColumns))
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
		return errors.Wrap(err, "schemas: unable to insert into followers")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.FollowerID,
		o.FollowedID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to populate default values for followers")
	}

CacheNoHooks:
	if !cached {
		followerInsertCacheMut.Lock()
		followerInsertCache[key] = cache
		followerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Follower.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Follower) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	followerUpdateCacheMut.RLock()
	cache, cached := followerUpdateCache[key]
	followerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			followerAllColumns,
			followerPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("schemas: unable to update followers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `followers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, followerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(followerType, followerMapping, append(wl, followerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "schemas: unable to update followers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by update for followers")
	}

	if !cached {
		followerUpdateCacheMut.Lock()
		followerUpdateCache[key] = cache
		followerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q followerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update all for followers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to retrieve rows affected for followers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FollowerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), followerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `followers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, followerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to update all in follower slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to retrieve rows affected all in update all follower")
	}
	return rowsAff, nil
}

var mySQLFollowerUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Follower) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schemas: no followers provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(followerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFollowerUniqueColumns, o)

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

	followerUpsertCacheMut.RLock()
	cache, cached := followerUpsertCache[key]
	followerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			followerAllColumns,
			followerColumnsWithDefault,
			followerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			followerAllColumns,
			followerPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("schemas: unable to upsert followers, could not build update column list")
		}

		ret := strmangle.SetComplement(followerAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`followers`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `followers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(followerType, followerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(followerType, followerMapping, ret)
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
		return errors.Wrap(err, "schemas: unable to upsert for followers")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(followerType, followerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to retrieve unique values for followers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to populate default values for followers")
	}

CacheNoHooks:
	if !cached {
		followerUpsertCacheMut.Lock()
		followerUpsertCache[key] = cache
		followerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Follower record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Follower) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("schemas: no Follower provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), followerPrimaryKeyMapping)
	sql := "DELETE FROM `followers` WHERE `follower_id`=? AND `followed_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete from followers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by delete for followers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q followerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schemas: no followerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete all from followers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by deleteall for followers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FollowerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(followerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), followerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `followers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, followerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schemas: unable to delete all from follower slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schemas: failed to get rows affected by deleteall for followers")
	}

	if len(followerAfterDeleteHooks) != 0 {
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
func (o *Follower) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFollower(ctx, exec, o.FollowerID, o.FollowedID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FollowerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FollowerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), followerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `followers`.* FROM `followers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, followerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schemas: unable to reload all in FollowerSlice")
	}

	*o = slice

	return nil
}

// FollowerExists checks if the Follower row exists.
func FollowerExists(ctx context.Context, exec boil.ContextExecutor, followerID string, followedID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `followers` where `follower_id`=? AND `followed_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, followerID, followedID)
	}
	row := exec.QueryRowContext(ctx, sql, followerID, followedID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schemas: unable to check if followers exists")
	}

	return exists, nil
}

// Exists checks if the Follower row exists.
func (o *Follower) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FollowerExists(ctx, exec, o.FollowerID, o.FollowedID)
}