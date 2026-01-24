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

	cldassetpbv1 "github.com/mikhail5545/media-service-client/pb/media_service/cloudinary/asset/v1"
	muxassetpbv1 "github.com/mikhail5545/media-service-client/pb/media_service/mux/asset/v1"
)

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *MuxAssetServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.AssetServiceClient = muxassetpbv1.NewAssetServiceClient(c.conn)
	return nil
}

// Connect establishes a gRPC connection to the specified address and initializes the service client.
// It accepts a context for cancellation and a variadic list of connection options.
// See WithInsecure, WithALTS, WithTransportCredentials, WithTLSConfig, WithTLSFromFiles, WithPerRPCCredentials
// and WithExtraDialOpts for available connection options.
func (c *CloudinaryAssetServiceClient) Connect(ctx context.Context, address string, opt ...ConnOption) (err error) {
	c.conn, err = connect(ctx, address, opt...)
	if err != nil {
		return err
	}
	c.AssetServiceClient = cldassetpbv1.NewAssetServiceClient(c.conn)
	return nil
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *MuxAssetServiceClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Close closes the grpc.ClientConn and all underlying connections.
func (c *CloudinaryAssetServiceClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
