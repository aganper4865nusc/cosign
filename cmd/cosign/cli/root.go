// Copyright 2024 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cli implements the cosign command-line interface.
package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags.
	Version = "dev"
	// GitCommit is set at build time via ldflags.
	GitCommit = "unknown"
)

// New returns the root cobra command for cosign.
func New() *cobra.Command {
	var logLevel string

	rootCmd := &cobra.Command{
		Use:   "cosign",
		Short: "A tool for container signing, verification, and storage in an OCI registry.",
		Long: `cosign is a tool for signing and verifying container images and other
artifacts using sigstore infrastructure. It supports keyless signing via
OIDC, as well as traditional key-based signing.

For more information, visit https://sigstore.dev`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// Persistent flags available to all subcommands.
	// Changed default log level from "warn" to "info" for more verbose output during personal use.
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info",
		"Log level (debug, info, warn, error)")

	// Register subcommands.
	rootCmd.AddCommand(versionCmd())

	return rootCmd
}

// Execute runs the root command with the provided context.
func Execute(ctx context.Context) error {
	return New().ExecuteContext(ctx)
}

// versionCmd returns the version subcommand.
func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(os.Stdout, "cosign version %s (commit: %s)\n", Version, GitCommit)
			return nil
		},
	}
}
