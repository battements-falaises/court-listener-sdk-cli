// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/court-listener-sdk-cli/internal/mocktest"
)

func TestCourtsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "courts", "retrieve",
			"--api-key", "string",
			"--id", "id",
			"--fields", "fields",
			"--format", "json",
			"--omit", "omit",
		)
	})
}

func TestCourtsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "courts", "list",
			"--api-key", "string",
			"--id", "id",
			"--count", "on",
			"--cursor", "cursor",
			"--date-modified", "'2019-12-27T18:11:19.117Z'",
			"--date-modified-gte", "'2019-12-27T18:11:19.117Z'",
			"--date-modified-lte", "'2019-12-27T18:11:19.117Z'",
			"--fields", "fields",
			"--format", "json",
			"--full-name", "full_name",
			"--full-name-startswith", "full_name__startswith",
			"--id-in", "id__in",
			"--jurisdiction", "jurisdiction",
			"--omit", "omit",
			"--order-by", "order_by",
			"--page", "1",
		)
	})
}
