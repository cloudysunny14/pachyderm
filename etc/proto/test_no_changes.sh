#!/bin/bash
# This script detects any changes to generated protobuf code

set -ex
which sha256sum

# cd to top-level pachyderm directory
scriptdir="$(dirname "${0}")"
cd "${scriptdir}/../.."

# hash our generated protobuf code, mostly to see if make proto changed anything
orig_hash="$(
find src -regex ".*\.pb\.go" \
  | sort -u \
  | xargs cat \
  | sha256sum \
  | awk '{print $1}'
)"

make proto

# hash newly-generated code
new_hash="$(
find src -regex ".*\.pb\.go" \
  | sort -u \
  | xargs cat \
  | sha256sum \
  | awk '{print $1}'
)"

# Exit with error if code changed
test "${orig_hash}" == "${new_hash}"
