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

func (c *MuxAssetServiceClient) List(ctx context.Context, in *muxassetpbv1.ListRequest, opts ...grpc.CallOption) (*muxassetpbv1.ListResponse, error) {
	return c.AssetServiceClient.List(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) List(ctx context.Context, in *cldassetpbv1.ListRequest, opts ...grpc.CallOption) (*cldassetpbv1.ListResponse, error) {
	return c.AssetServiceClient.List(ctx, in, opts...)
}

func (c *MuxAssetServiceClient) ListArchived(ctx context.Context, in *muxassetpbv1.ListArchivedRequest, opts ...grpc.CallOption) (*muxassetpbv1.ListArchivedResponse, error) {
	return c.AssetServiceClient.ListArchived(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) ListArchived(ctx context.Context, in *cldassetpbv1.ListArchivedRequest, opts ...grpc.CallOption) (*cldassetpbv1.ListArchivedResponse, error) {
	return c.AssetServiceClient.ListArchived(ctx, in, opts...)
}

func (c *MuxAssetServiceClient) ListBroken(ctx context.Context, in *muxassetpbv1.ListBrokenRequest, opts ...grpc.CallOption) (*muxassetpbv1.ListBrokenResponse, error) {
	return c.AssetServiceClient.ListBroken(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) ListBroken(ctx context.Context, in *cldassetpbv1.ListBrokenRequest, opts ...grpc.CallOption) (*cldassetpbv1.ListBrokenResponse, error) {
	return c.AssetServiceClient.ListBroken(ctx, in, opts...)
}
