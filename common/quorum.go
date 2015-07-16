package common

import "github.com/yahoo/coname/proto"

// CheckQuorum evaluates whether the quorum requirement want can be satisfied
// by ratifications of the verifiers in have.
func CheckQuorum(want *proto.QuorumExpr, have map[uint64]struct{}) bool {
	remaining := want.Threshold // unsigned
	if remaining == 0 {
		return true
	}
	for _, verifier := range want.Verifiers {
		if _, yes := have[verifier]; yes {
			if remaining--; remaining == 0 {
				return true
			}
		}
	}
	for _, e := range want.Subexpressions {
		if CheckQuorum(e, have) {
			if remaining--; remaining == 0 {
				return true
			}
		}
	}
	return false
}

// ListQuorum inserts all verifiers mentioned in e to out. If out is nil, a new
// map is allocated.
func ListQuorum(e *proto.QuorumExpr, out map[uint64]struct{}) map[uint64]struct{} {
	if out == nil {
		out = make(map[uint64]struct{}, len(e.Verifiers))
	}
	for _, verifier := range e.Verifiers {
		out[verifier] = struct{}{}
	}
	for _, e := range e.Subexpressions {
		ListQuorum(e, out)
	}
	return out
}
