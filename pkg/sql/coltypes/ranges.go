// Copyright 2017 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package coltypes

import (
	"bytes"
	"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

// TRange represents a XXXRANGE column type.
type TRange struct {
	ParamType T
}

// TypeName implements the ColTypeFormatter interface.
func (node *TRange) TypeName() string {
		return node.ParamType.TypeName() + "RANGE"
}

// Format implements the ColTypeFormatter interface.
func (node *TRange) Format(buf *bytes.Buffer, f lex.EncodeFlags) {
	buf.WriteString(node.TypeName())
}

func canBeInRangeColType(t T) bool {
	// TODO: Any column type with b-tree operator can be customized with range col type,
	// TODO: But for now we just impl the builtin ones.
	switch t.(type) {
	case *TInt, *TDecimal, *TTimestamp, *TTimestampTZ:
		return true
	default:
		return false
	}
}
