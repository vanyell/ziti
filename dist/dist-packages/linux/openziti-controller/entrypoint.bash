#!/usr/bin/env bash
#
# This thin wrapper for the OpenZiti Controller is used by the systemd service and Docker container to call the bootstrap() function before invoking ziti. The bootstrap() function will exit 0 immediately if ZITI_BOOTSTRAP is not set to true. Otherwise,  it will 
#
# usage:
#   entrypoint.bash run config.yml

set -o errexit
set -o nounset
set -o pipefail

# discard debug unless DEBUG
: "${DEBUG:=0}"
if (( DEBUG )); then
  exec 3>&1
  set -o xtrace
else
  exec 3>/dev/null
fi

if ! (( $# )); then
  # if no args, run the controller with the default config file
  set -- run config.yml
elif [[ "${1}" == run && -z "${2:-}" ]]; then
  # if first arg is "run" and second arg is empty, run the controller with the default config file
  set -- run config.yml
fi

# shellcheck disable=SC1090 # default path is set by the systemd service
source "${ZITI_CTRL_BOOTSTRAP_BASH:-/opt/openziti/etc/controller/bootstrap.bash}"

# if first arg is "run", bootstrap the controller with the config file
if [[ "${1}" =~ run|boot && ! -s "${2}" ]]; then
  echo "ERROR: ${2} does not exist" >&2
  hintBootstrap "${PWD}"
  exit 1
elif [[ "${1}" =~ run|boot && ! -w "$(dbFile "${2}")" ]]; then
  echo "ERROR: database file '$(dbFile "${2}")' is not writable" >&2
  hintBootstrap "${PWD}"
  exit 1
elif [[ "${1}" =~ run|boot && "${ZITI_BOOTSTRAP:-}" == true ]]; then
  bootstrap "${2}"
fi

if [[ "${1}" == run ]]; then
  # shellcheck disable=SC2068
  exec ziti controller ${@}
fi
