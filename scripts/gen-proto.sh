#!/usr/bin/env bash
#
# Copyright (c) 2026. Mikhail Kulik.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published
# by the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.
#

set -euo pipefail

if ! command -v buf &>/dev/null; then
  echo "install buf: https://docs.buf.build/installation"
  exit 2
fi

OUT_DIR=$(mktemp -d)
trap 'rm -rf "${OUT_DIR}"' EXIT

buf generate buf.build/vitainmove/protos --template buf.gen.yaml --output "${OUT_DIR}"

rm -rf pb_tmp || true
mv "${OUT_DIR}" pb_tmp
echo "Generated pb in pb_tmp/"