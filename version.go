package pool

// Copyright (c) 2013-2017 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Heavily inspired by https://github.com/btcsuite/btcd/blob/master/version.go
// Copyright (C) 2015-2019 The Lightning Network Developers

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"strings"

	"google.golang.org/grpc/metadata"
)

// SetAgentName overwrites the default agent name which can be used to identify
// the software Pool is bundled in (for example LiT). This function panics if
// the agent name contains characters outside of the allowed semantic alphabet.
func SetAgentName(newAgentName string) {
	for _, r := range newAgentName {
		if !strings.ContainsRune(semanticAlphabet, r) {
			panic(fmt.Errorf("rune: %v is not in the semantic "+
				"alphabet", r))
		}
	}

	agentName = newAgentName
}

// Version returns the application version as a properly formed string per the
// semantic versioning 2.0.0 spec (http://semver.org/) and the commit it was
// built on.
func Version() string {
	// Append commit hash of current build to version.
	return fmt.Sprintf("%s commit=%s", semanticVersion(), Commit)
}

// ContextWithInitiator creates a new context with the given initiator string
// added (provided it is not empty).
func ContextWithInitiator(ctx context.Context,
	initiator string) context.Context {

	trimmed := strings.TrimSpace(initiator)
	if trimmed == "" {
		return ctx
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}

	md[additionalInfoInitiator] = []string{trimmed}
	return metadata.NewIncomingContext(ctx, md)
}
