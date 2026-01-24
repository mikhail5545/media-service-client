/*
 * Copyright (c) 2026. Mikhail Kulik.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package client

import (
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type connOptions struct {
	insecure             bool
	transportCredentials credentials.TransportCredentials
	perRPCCredentials    credentials.PerRPCCredentials
	extraDialOpts        []grpc.DialOption
	dialTimeout          time.Duration
	forceALTS            bool
}

type ConnOption func(*connOptions)

// WithInsecure configures the gRPC connection to be insecure (no TLS).
// Use with caution, as this transmits data in plaintext.
func WithInsecure() ConnOption {
	return func(co *connOptions) {
		co.insecure = true
	}
}

// WithTransportCredentials sets custom transport credentials for the gRPC connection.
func WithTransportCredentials(tc credentials.TransportCredentials) ConnOption {
	return func(co *connOptions) {
		co.transportCredentials = tc
	}
}

// WithTLSConfig sets up TLS transport credentials using the provided tls.Config.
func WithTLSConfig(cfg *tls.Config) ConnOption {
	return func(co *connOptions) {
		co.transportCredentials = credentials.NewTLS(cfg)
	}
}

// WithTLSFromFiles sets up TLS transport credentials using the provided certificate files.
func WithTLSFromFiles(caPath, certPath, keyPath string) ConnOption {
	return func(co *connOptions) {
		tc, err := tlsCredentialsFromFiles(caPath, certPath, keyPath)
		if err != nil {
			panic(err)
		}
		co.transportCredentials = tc
	}
}

// WithPerRPCCredentials sets custom per-RPC credentials for the gRPC connection.
func WithPerRPCCredentials(prc credentials.PerRPCCredentials) ConnOption {
	return func(co *connOptions) {
		co.perRPCCredentials = prc
	}
}

// WithALTS forces the use of ALTS transport credentials for the gRPC connection.
func WithALTS() ConnOption {
	return func(co *connOptions) {
		co.forceALTS = true
	}
}

// WithExtraDialOpts appends additional grpc.DialOption to the gRPC connection.
func WithExtraDialOpts(opts ...grpc.DialOption) ConnOption {
	return func(co *connOptions) {
		co.extraDialOpts = append(co.extraDialOpts, opts...)
	}
}
