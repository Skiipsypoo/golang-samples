// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package howto

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestSearchFeaturedJobs(t *testing.T) {
	tc := testutil.SystemTest(t)
	testutil.Retry(t, 10, 1*time.Second, func(r *testutil.R) {
		buf := &bytes.Buffer{}
		if _, err := searchFeaturedJobs(buf, tc.ProjectID, testCompany.Name, "SWE"); err != nil {
			r.Errorf("searchFeaturedJobs: %v", err)
		}
		want := testJob.Name
		if got := buf.String(); !strings.Contains(got, want) {
			r.Errorf("searchFeaturedJobs got %q, want to contain %q", got, want)
		}
	})
}
