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
	"google.golang.org/grpc"
)

func (c *MuxAssetServiceClient) AddOwner(ctx context.Context, in *muxassetpbv1.AddOwnerRequest, opts ...grpc.CallOption) (*muxassetpbv1.AddOwnerResponse, error) {
	return c.AssetServiceClient.AddOwner(ctx, in, opts...)
}

func (c *MuxAssetServiceClient) RemoveOwner(ctx context.Context, in *muxassetpbv1.RemoveOwnerRequest, opts ...grpc.CallOption) (*muxassetpbv1.RemoveOwnerResponse, error) {
	return c.AssetServiceClient.RemoveOwner(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) AddOwner(ctx context.Context, in *cldassetpbv1.AddOwnerRequest, opts ...grpc.CallOption) (*cldassetpbv1.AddOwnerResponse, error) {
	return c.AssetServiceClient.AddOwner(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) RemoveOwner(ctx context.Context, in *cldassetpbv1.RemoveOwnerRequest, opts ...grpc.CallOption) (*cldassetpbv1.RemoveOwnerResponse, error) {
	return c.AssetServiceClient.RemoveOwner(ctx, in, opts...)
}
