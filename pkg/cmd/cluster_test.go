// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/court-listener-sdk-cli/internal/mocktest"
)

func TestClustersRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"clusters", "retrieve",
			"--id", "0",
			"--fields", "fields",
			"--format", "json",
			"--omit", "omit",
		)
	})
}

func TestClustersList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"clusters", "list",
			"--max-items", "10",
			"--id", "0",
			"--citation", "citation",
			"--count", "on",
			"--cursor", "cursor",
			"--date-created", "'2019-12-27T18:11:19.117Z'",
			"--date-created-gte", "'2019-12-27T18:11:19.117Z'",
			"--date-created-lte", "'2019-12-27T18:11:19.117Z'",
			"--date-filed", "'2019-12-27'",
			"--date-filed-gte", "'2019-12-27'",
			"--date-filed-lte", "'2019-12-27'",
			"--date-modified", "'2019-12-27T18:11:19.117Z'",
			"--date-modified-gte", "'2019-12-27T18:11:19.117Z'",
			"--date-modified-lte", "'2019-12-27T18:11:19.117Z'",
			"--docket", "0",
			"--docket-court", "docket__court",
			"--docket-docket-number", "docket__docket_number",
			"--fields", "fields",
			"--format", "json",
			"--id-gt", "0",
			"--id-gte", "0",
			"--id-lt", "0",
			"--id-lte", "0",
			"--id-range", "id__range",
			"--judges", "judges",
			"--omit", "omit",
			"--order-by", "order_by",
			"--page", "1",
		)
	})
}
