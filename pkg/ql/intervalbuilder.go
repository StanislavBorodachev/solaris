// Copyright 2024 The Solaris Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ql

import (
	"fmt"
	"github.com/solarisdb/solaris/golibs/errors"
	"github.com/solarisdb/solaris/pkg/intervals"
	"sort"
)

// ParamIntervalBuilder allows to build value intervals from the AST expression
// for a given parameter and comparison operations specified.
type ParamIntervalBuilder[T, K any] struct {
	basis   intervals.Basis[T]
	dialect Dialect[K]

	param string
	ops   map[string]bool
}

var (
	OpsAll  = []string{"<", ">", "<=", ">=", "=", "!="}
	OpsGtLt = []string{"<", ">"}
)

// NewParamIntervalBuilder returns new ParamIntervalBuilder.
func NewParamIntervalBuilder[T, K any](basis intervals.Basis[T], dialect Dialect[K], param string, ops []string) ParamIntervalBuilder[T, K] {
	opsMap := make(map[string]bool)
	for _, op := range ops {
		opsMap[op] = true
	}
	return ParamIntervalBuilder[T, K]{basis: basis, dialect: dialect, param: param, ops: opsMap}
}

// Build returns a list of intervals built from the AST expression.
// Returned intervals are sorted by the L border.
func (ib *ParamIntervalBuilder[T, K]) Build(expr *Expression) ([]intervals.Interval[T], error) {
	var res []intervals.Interval[T]
	for _, or := range expr.Or {
		tt, err := ib.buildOR(or)
		if err != nil {
			return nil, err
		}
		if len(tt) > 0 {
			res = append(res, tt...)
		}
	}
	return ib.union(res), nil
}

func (ib *ParamIntervalBuilder[T, K]) buildOR(or *OrCondition) ([]intervals.Interval[T], error) {
	var groups [][]intervals.Interval[T]
	for _, and := range or.And {
		group, err := ib.buildXCond(and)
		if err != nil {
			return nil, err
		}
		if len(group) > 0 {
			groups = append(groups, group)
		}
	}
	return ib.intersect(groups), nil
}

func (ib *ParamIntervalBuilder[T, K]) buildXCond(and *XCondition) ([]intervals.Interval[T], error) {
	var res []intervals.Interval[T]
	if and.Expr != nil {
		var err error
		res, err = ib.Build(and.Expr)
		if err != nil {
			return nil, err
		}
	} else {
		tt, err := ib.buildCond(and.Cond)
		if err != nil {
			return nil, err
		}
		if len(tt) > 0 {
			res = append(res, tt...)
		}
	}
	if !and.Not {
		return res, nil
	}
	var negated []intervals.Interval[T]
	for _, t := range res {
		negated = append(negated, ib.basis.Negate(t)...)
	}
	return negated, nil
}

func (ib *ParamIntervalBuilder[T, K]) buildCond(cond *Condition) ([]intervals.Interval[T], error) {
	// param1
	p1 := cond.FirstParam
	dp1, ok := ib.dialect[p1.ID()]
	if !ok {
		return nil, fmt.Errorf("the parameter %s must be known: %w", p1.Name(false), errors.ErrInvalid)
	}
	if dp1.Flags&PfLValue == 0 {
		return nil, fmt.Errorf("the parameter %s must be on the left side of the condition: %w", p1.Name(false), errors.ErrInvalid)
	}
	if dp1.Flags&PfNop != 0 {
		return nil, fmt.Errorf("the parameter %s must allow operation (%s): %w", p1.Name(false), cond.Op, errors.ErrInvalid)
	}
	if p1.Name(false) != ib.param { // skip not the param we look for
		return nil, nil
	}

	// param2
	p2 := cond.SecondParam
	if p2 == nil {
		return nil, fmt.Errorf("the second parameter must be specified for the parameter %s and the operation %q: %w", p1.Name(false), cond.Op, errors.ErrInvalid)
	}
	dp2, ok := ib.dialect[p2.ID()]
	if !ok {
		return nil, fmt.Errorf("the second parameter %s must be known: %w", p2.Name(false), errors.ErrInvalid)
	}
	if dp2.Flags&PfRValue == 0 {
		return nil, fmt.Errorf("the second parameter %s must be on the right side of the condition: %w", p2.Name(false), errors.ErrInvalid)
	}
	if dp2.Flags&PfNop != 0 {
		return nil, fmt.Errorf("the second parameter %s must allow operation (%s): %w", p2.Name(false), cond.Op, errors.ErrInvalid)
	}
	if p2.Const == nil { // skip not a constant param
		return nil, nil
	}

	// operation
	if !ib.ops[cond.Op] { // skip not the ops we look for
		return nil, nil
	}
	switch cond.Op {
	case "<", ">":
		if dp1.Flags&PfComparable == 0 && dp1.Flags&PfGreaterLess == 0 {
			return nil, fmt.Errorf("the first parameter %s must be comparable for the operation %s: %w", p1.Name(false), cond.Op, errors.ErrInvalid)
		}
		if dp2.Flags&PfComparable == 0 && dp2.Flags&PfGreaterLess == 0 {
			return nil, fmt.Errorf("the second parameter %s must be comparable for the operation %s: %w", p2.Name(false), cond.Op, errors.ErrInvalid)
		}
	case "<=", ">=", "=", "!=":
		if dp1.Flags&PfComparable == 0 {
			return nil, fmt.Errorf("the first parameter %s must be comparable for the operation %s: %w", p1.Name(false), cond.Op, errors.ErrInvalid)
		}
		if dp2.Flags&PfComparable == 0 {
			return nil, fmt.Errorf("the second parameter %s must be comparable for the operation %s: %w", p2.Name(false), cond.Op, errors.ErrInvalid)
		}
	}

	// value
	vf, err := castValueF(dp2.ValueF, dp2.Type, dp1.Type)
	if err != nil {
		return nil, err
	}
	kVal, err := vf(cond.SecondParam, *new(K))
	if err != nil {
		return nil, err
	}
	tVal, ok := kVal.(T)
	if !ok {
		return nil, fmt.Errorf("cannot cast the second parameter value(type=%T) to interval point(type=%T): %w", kVal, tVal, errors.ErrInvalid)
	}

	// intervals
	return ib.getIntervals(cond.Op, tVal), nil
}

func (ib *ParamIntervalBuilder[T, K]) union(intervalsL []intervals.Interval[T]) []intervals.Interval[T] {
	if len(intervalsL) == 0 {
		return intervalsL
	}
	sort.Slice(intervalsL, func(i, j int) bool {
		return ib.basis.StartsBefore(intervalsL[i], intervalsL[j])
	})
	var res []intervals.Interval[T]
	prev := intervalsL[0]
	for i := 1; i < len(intervalsL); i++ {
		curr := intervalsL[i]
		if union, ok := ib.basis.Union(prev, curr); ok {
			prev = union
			continue
		}
		res = append(res, prev)
		prev = curr
	}
	res = append(res, prev)
	return res
}

func (ib *ParamIntervalBuilder[T, K]) intersect(groups [][]intervals.Interval[T]) []intervals.Interval[T] {
	if len(groups) == 0 {
		return nil
	}
	prev := groups[0]
	for i := 1; i < len(groups); i++ {
		curr := groups[i]
		var group []intervals.Interval[T]
		for j := 0; j < len(prev); j++ {
			for k := 0; k < len(curr); k++ {
				if t, ok := ib.basis.Intersect(prev[j], curr[k]); ok {
					group = append(group, t)
				}
			}
		}
		prev = group
	}
	return prev
}

func (ib *ParamIntervalBuilder[T, K]) getIntervals(op string, val T) []intervals.Interval[T] {
	switch op {
	case "<":
		return []intervals.Interval[T]{ib.basis.OpenR(ib.basis.Min, val)}
	case ">":
		return []intervals.Interval[T]{ib.basis.OpenL(val, ib.basis.Max)}
	case "<=":
		return []intervals.Interval[T]{ib.basis.Closed(ib.basis.Min, val)}
	case ">=":
		return []intervals.Interval[T]{ib.basis.Closed(val, ib.basis.Max)}
	case "=":
		return []intervals.Interval[T]{ib.basis.Closed(val, val)}
	case "!=":
		return ib.basis.Negate(ib.basis.Closed(val, val))
	}
	return nil
}
