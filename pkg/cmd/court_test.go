// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/battements-falaises/court-listener-sdk-cli/internal/mocktest"
)

func TestCourtsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"courts", "retrieve",
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
			t,
			"--api-key", "string",
			"courts", "list",
			"--max-items", "10",
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
