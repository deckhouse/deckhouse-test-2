#!/bin/bash
{{- /*
# Copyright 2024 Flant JSC
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
*/}}
{{- if and (ne .nodeGroup.nodeType "Static") (.provider )}}
  {{- if $bootstrap_script_network := $.Files.Get (printf "deckhouse/candi/cloud-providers/%s/bashible/bootstrap-networks.sh.tpl" .provider) | default ($.Files.Get (printf "candi/cloud-providers/%s/bashible/bootstrap-networks.sh.tpl" .provider) ) }}
    {{- tpl ($bootstrap_script_network) $ | nindent 0 }}
  {{- end }}
{{- end }}
