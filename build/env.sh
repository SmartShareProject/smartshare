#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
sspdir="$workspace/src/github.com/smartshareproject"
if [ ! -L "$sspdir/smartshare" ]; then
    mkdir -p "$sspdir"
    cd "$sspdir"
    ln -s ../../../../../. smartshare
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$sspdir/smartshare"
PWD="$sspdir/smartshare"

# Launch the arguments with the configured environment.
exec "$@"
