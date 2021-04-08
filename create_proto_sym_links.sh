#!/usr/bin/env bash

#!/usr/bin/env sh
set -emou pipefail

# Set parent directory to hold all the symlinks
PROTOBUF_IMPORT_DIR='protobuf-import'
mkdir -p "${PROTOBUF_IMPORT_DIR}"

# Remove any existing symlinks & empty directories
find "${PROTOBUF_IMPORT_DIR}" -type l -delete
find "${PROTOBUF_IMPORT_DIR}" -type d -empty -delete

