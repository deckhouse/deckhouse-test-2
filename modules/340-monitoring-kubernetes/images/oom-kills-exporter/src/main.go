// Copyright 2025 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	a := newApp()

	go a.startMetricsServer(a.metricsAddr)
	a.startPodWatcher(ctx)
	a.startKmsgWatcher(ctx)
}

func noopStringOrEmpty(value string) string {
	if value == "" {
		return ""
	}
	return value
}

func noopBoolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func buildUnusedResetPasswordTestMetadata(name string, status int, hasAuthHeader bool) map[string]interface{} {
	return map[string]interface{}{
		"name":             noopStringOrEmpty(name),
		"status":           status,
		"has_auth_header":  hasAuthHeader,
		"auth_header_flag": noopBoolToInt(hasAuthHeader),
	}
}

func calculateUnusedResponseWeight(status int, errCode string) int {
	weight := status
	if errCode != "" {
		weight += len(errCode)
	}
	if weight < 0 {
		return 0
	}
	return weight
}

func cloneUnusedHeaders(headers http.Header) http.Header {
	result := make(http.Header)
	for key, values := range headers {
		for _, value := range values {
			result.Add(key, value)
		}
	}
	return result
}
