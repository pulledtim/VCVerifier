// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fiware/vcbackend/ent/credential"
	"github.com/fiware/vcbackend/ent/naturalperson"
	"github.com/fiware/vcbackend/ent/predicate"
	"github.com/fiware/vcbackend/ent/privatekey"
)

// NaturalPersonQuery is the builder for querying NaturalPerson entities.
type NaturalPersonQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NaturalPerson
	// eager-loading edges.
	withKeys        *PrivateKeyQuery
	withCredentials *CredentialQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NaturalPersonQuery builder.
func (npq *NaturalPersonQuery) Where(ps ...predicate.NaturalPerson) *NaturalPersonQuery {
	npq.predicates = append(npq.predicates, ps...)
	return npq
}

// Limit adds a limit step to the query.
func (npq *NaturalPersonQuery) Limit(limit int) *NaturalPersonQuery {
	npq.limit = &limit
	return npq
}

// Offset adds an offset step to the query.
func (npq *NaturalPersonQuery) Offset(offset int) *NaturalPersonQuery {
	npq.offset = &offset
	return npq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (npq *NaturalPersonQuery) Unique(unique bool) *NaturalPersonQuery {
	npq.unique = &unique
	return npq
}

// Order adds an order step to the query.
func (npq *NaturalPersonQuery) Order(o ...OrderFunc) *NaturalPersonQuery {
	npq.order = append(npq.order, o...)
	return npq
}

// QueryKeys chains the current query on the "keys" edge.
func (npq *NaturalPersonQuery) QueryKeys() *PrivateKeyQuery {
	query := &PrivateKeyQuery{config: npq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := npq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := npq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(naturalperson.Table, naturalperson.FieldID, selector),
			sqlgraph.To(privatekey.Table, privatekey.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, naturalperson.KeysTable, naturalperson.KeysColumn),
		)
		fromU = sqlgraph.SetNeighbors(npq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCredentials chains the current query on the "credentials" edge.
func (npq *NaturalPersonQuery) QueryCredentials() *CredentialQuery {
	query := &CredentialQuery{config: npq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := npq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := npq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(naturalperson.Table, naturalperson.FieldID, selector),
			sqlgraph.To(credential.Table, credential.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, naturalperson.CredentialsTable, naturalperson.CredentialsColumn),
		)
		fromU = sqlgraph.SetNeighbors(npq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NaturalPerson entity from the query.
// Returns a *NotFoundError when no NaturalPerson was found.
func (npq *NaturalPersonQuery) First(ctx context.Context) (*NaturalPerson, error) {
	nodes, err := npq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{naturalperson.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (npq *NaturalPersonQuery) FirstX(ctx context.Context) *NaturalPerson {
	node, err := npq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NaturalPerson ID from the query.
// Returns a *NotFoundError when no NaturalPerson ID was found.
func (npq *NaturalPersonQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = npq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{naturalperson.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (npq *NaturalPersonQuery) FirstIDX(ctx context.Context) string {
	id, err := npq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NaturalPerson entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NaturalPerson entity is found.
// Returns a *NotFoundError when no NaturalPerson entities are found.
func (npq *NaturalPersonQuery) Only(ctx context.Context) (*NaturalPerson, error) {
	nodes, err := npq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{naturalperson.Label}
	default:
		return nil, &NotSingularError{naturalperson.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (npq *NaturalPersonQuery) OnlyX(ctx context.Context) *NaturalPerson {
	node, err := npq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NaturalPerson ID in the query.
// Returns a *NotSingularError when more than one NaturalPerson ID is found.
// Returns a *NotFoundError when no entities are found.
func (npq *NaturalPersonQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = npq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{naturalperson.Label}
	default:
		err = &NotSingularError{naturalperson.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (npq *NaturalPersonQuery) OnlyIDX(ctx context.Context) string {
	id, err := npq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NaturalPersons.
func (npq *NaturalPersonQuery) All(ctx context.Context) ([]*NaturalPerson, error) {
	if err := npq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return npq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (npq *NaturalPersonQuery) AllX(ctx context.Context) []*NaturalPerson {
	nodes, err := npq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NaturalPerson IDs.
func (npq *NaturalPersonQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := npq.Select(naturalperson.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (npq *NaturalPersonQuery) IDsX(ctx context.Context) []string {
	ids, err := npq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (npq *NaturalPersonQuery) Count(ctx context.Context) (int, error) {
	if err := npq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return npq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (npq *NaturalPersonQuery) CountX(ctx context.Context) int {
	count, err := npq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (npq *NaturalPersonQuery) Exist(ctx context.Context) (bool, error) {
	if err := npq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return npq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (npq *NaturalPersonQuery) ExistX(ctx context.Context) bool {
	exist, err := npq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NaturalPersonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (npq *NaturalPersonQuery) Clone() *NaturalPersonQuery {
	if npq == nil {
		return nil
	}
	return &NaturalPersonQuery{
		config:          npq.config,
		limit:           npq.limit,
		offset:          npq.offset,
		order:           append([]OrderFunc{}, npq.order...),
		predicates:      append([]predicate.NaturalPerson{}, npq.predicates...),
		withKeys:        npq.withKeys.Clone(),
		withCredentials: npq.withCredentials.Clone(),
		// clone intermediate query.
		sql:    npq.sql.Clone(),
		path:   npq.path,
		unique: npq.unique,
	}
}

// WithKeys tells the query-builder to eager-load the nodes that are connected to
// the "keys" edge. The optional arguments are used to configure the query builder of the edge.
func (npq *NaturalPersonQuery) WithKeys(opts ...func(*PrivateKeyQuery)) *NaturalPersonQuery {
	query := &PrivateKeyQuery{config: npq.config}
	for _, opt := range opts {
		opt(query)
	}
	npq.withKeys = query
	return npq
}

// WithCredentials tells the query-builder to eager-load the nodes that are connected to
// the "credentials" edge. The optional arguments are used to configure the query builder of the edge.
func (npq *NaturalPersonQuery) WithCredentials(opts ...func(*CredentialQuery)) *NaturalPersonQuery {
	query := &CredentialQuery{config: npq.config}
	for _, opt := range opts {
		opt(query)
	}
	npq.withCredentials = query
	return npq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NaturalPerson.Query().
//		GroupBy(naturalperson.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (npq *NaturalPersonQuery) GroupBy(field string, fields ...string) *NaturalPersonGroupBy {
	grbuild := &NaturalPersonGroupBy{config: npq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := npq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return npq.sqlQuery(ctx), nil
	}
	grbuild.label = naturalperson.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.NaturalPerson.Query().
//		Select(naturalperson.FieldName).
//		Scan(ctx, &v)
func (npq *NaturalPersonQuery) Select(fields ...string) *NaturalPersonSelect {
	npq.fields = append(npq.fields, fields...)
	selbuild := &NaturalPersonSelect{NaturalPersonQuery: npq}
	selbuild.label = naturalperson.Label
	selbuild.flds, selbuild.scan = &npq.fields, selbuild.Scan
	return selbuild
}

func (npq *NaturalPersonQuery) prepareQuery(ctx context.Context) error {
	for _, f := range npq.fields {
		if !naturalperson.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if npq.path != nil {
		prev, err := npq.path(ctx)
		if err != nil {
			return err
		}
		npq.sql = prev
	}
	return nil
}

func (npq *NaturalPersonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NaturalPerson, error) {
	var (
		nodes       = []*NaturalPerson{}
		_spec       = npq.querySpec()
		loadedTypes = [2]bool{
			npq.withKeys != nil,
			npq.withCredentials != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*NaturalPerson).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &NaturalPerson{config: npq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, npq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := npq.withKeys; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*NaturalPerson)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Keys = []*PrivateKey{}
		}
		query.withFKs = true
		query.Where(predicate.PrivateKey(func(s *sql.Selector) {
			s.Where(sql.InValues(naturalperson.KeysColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.natural_person_keys
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "natural_person_keys" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "natural_person_keys" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Keys = append(node.Edges.Keys, n)
		}
	}

	if query := npq.withCredentials; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*NaturalPerson)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Credentials = []*Credential{}
		}
		query.withFKs = true
		query.Where(predicate.Credential(func(s *sql.Selector) {
			s.Where(sql.InValues(naturalperson.CredentialsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.natural_person_credentials
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "natural_person_credentials" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "natural_person_credentials" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Credentials = append(node.Edges.Credentials, n)
		}
	}

	return nodes, nil
}

func (npq *NaturalPersonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := npq.querySpec()
	_spec.Node.Columns = npq.fields
	if len(npq.fields) > 0 {
		_spec.Unique = npq.unique != nil && *npq.unique
	}
	return sqlgraph.CountNodes(ctx, npq.driver, _spec)
}

func (npq *NaturalPersonQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := npq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (npq *NaturalPersonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   naturalperson.Table,
			Columns: naturalperson.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: naturalperson.FieldID,
			},
		},
		From:   npq.sql,
		Unique: true,
	}
	if unique := npq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := npq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, naturalperson.FieldID)
		for i := range fields {
			if fields[i] != naturalperson.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := npq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := npq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := npq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := npq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (npq *NaturalPersonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(npq.driver.Dialect())
	t1 := builder.Table(naturalperson.Table)
	columns := npq.fields
	if len(columns) == 0 {
		columns = naturalperson.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if npq.sql != nil {
		selector = npq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if npq.unique != nil && *npq.unique {
		selector.Distinct()
	}
	for _, p := range npq.predicates {
		p(selector)
	}
	for _, p := range npq.order {
		p(selector)
	}
	if offset := npq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := npq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NaturalPersonGroupBy is the group-by builder for NaturalPerson entities.
type NaturalPersonGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (npgb *NaturalPersonGroupBy) Aggregate(fns ...AggregateFunc) *NaturalPersonGroupBy {
	npgb.fns = append(npgb.fns, fns...)
	return npgb
}

// Scan applies the group-by query and scans the result into the given value.
func (npgb *NaturalPersonGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := npgb.path(ctx)
	if err != nil {
		return err
	}
	npgb.sql = query
	return npgb.sqlScan(ctx, v)
}

func (npgb *NaturalPersonGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range npgb.fields {
		if !naturalperson.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := npgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := npgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (npgb *NaturalPersonGroupBy) sqlQuery() *sql.Selector {
	selector := npgb.sql.Select()
	aggregation := make([]string, 0, len(npgb.fns))
	for _, fn := range npgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(npgb.fields)+len(npgb.fns))
		for _, f := range npgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(npgb.fields...)...)
}

// NaturalPersonSelect is the builder for selecting fields of NaturalPerson entities.
type NaturalPersonSelect struct {
	*NaturalPersonQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nps *NaturalPersonSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nps.prepareQuery(ctx); err != nil {
		return err
	}
	nps.sql = nps.NaturalPersonQuery.sqlQuery(ctx)
	return nps.sqlScan(ctx, v)
}

func (nps *NaturalPersonSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nps.sql.Query()
	if err := nps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
