#!/bin/bash

# Copyright 2021 Flant JSC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

function parse_master_ip_from_log() {
  log_file="$1"
  >&2 echo "  Detect master_ip from bootstrap.log ..."
  if ! master_ip="$(grep -Po '(?<=master_ip_address_for_ssh = ).+$' "$log_file")"; then
    >&2 echo "    ERROR: can't parse master_ip from bootstrap.log"
    return 1
  fi
  echo "${master_ip}"
}

pwdd="$(dirname $0)"
bash "$pwdd/script.sh" $@
exit_code=$?

echo "script exit code: $exit_code"
if [[ "$1" == "run-test" ]]; then
  ssh_string=""
  # save ssh connection script for show in comment
  ssh_string_file="/tmp/ssh-master-connection-string"
  dhctl_log_file_path="/tmp/dhctl-log-file-path"

  >&2 echo "ssh exit_code $exit_code"
  >&2 echo "ssh string content: $(cat $ssh_string_file)"
  >&2 echo "dhctl log file path: $(cat $dhctl_log_file_path)"
  >&2 echo "$(cat $ssh_string_file)"

  if [[ -f "$ssh_string_file" ]]; then
    ssh_string="$(cat $ssh_string_file)"
  elif [ -f "$dhctl_log_file_path" ]; then
    ssh_str_file="$(dhctl_log_file_path)"
    ssh_string="$(parse_master_ip_from_log)"
  fi

  echo '::echo::on'
  echo "::set-output name=ssh_master_connection_string::${ssh_string}"
  echo '::echo::off'
fi

exit "$exit_code"
