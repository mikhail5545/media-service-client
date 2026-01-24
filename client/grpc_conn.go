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
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/credentials/insecure"
)

func connect(ctx context.Context, addr string, opts ...ConnOption) (*grpc.ClientConn, error) {
	co := &connOptions{}
	for _, opt := range opts {
		opt(co)
	}

	dialOpts, err := buildDialOptions(co)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(addr, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func buildDialOptions(co *connOptions) ([]grpc.DialOption, error) {
	var dialOpts []grpc.DialOption

	if co.forceALTS {
		altsCred := alts.NewClientCreds(&alts.ClientOptions{})
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(altsCred))
	} else if co.transportCredentials != nil {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(co.transportCredentials))
	} else if co.insecure {
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// Default: require TLS via system root pool
		tlsCfg := &tls.Config{} // default to system roots
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
	}

	if co.perRPCCredentials != nil {
		if co.perRPCCredentials.RequireTransportSecurity() && co.insecure {
			return nil, fmt.Errorf("per-RPC credentials require transport security, but connection is set to insecure")
		}
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(co.perRPCCredentials))
	}

	dialOpts = append(dialOpts, co.extraDialOpts...)
	return dialOpts, nil
}
