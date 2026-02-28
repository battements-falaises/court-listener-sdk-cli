// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/court-listener-sdk-cli/internal/mocktest"
)

func TestDocketsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"dockets", "retrieve",
		"--api-key", "string",
		"--id", "0",
		"--fields", "fields",
		"--format", "json",
		"--omit", "omit",
	)
}

func TestDocketsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"dockets", "list",
		"--api-key", "string",
		"--id", "0",
		"--blocked=true",
		"--case-name", "case_name",
		"--cause", "cause",
		"--count", "on",
		"--court", "court",
		"--court-jurisdiction", "court__jurisdiction",
		"--court-jurisdiction", "court__jurisdiction!",
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
		"--date-terminated", "'2019-12-27'",
		"--date-terminated-gte", "'2019-12-27'",
		"--date-terminated-lte", "'2019-12-27'",
		"--docket-number", "docket_number",
		"--fields", "fields",
		"--format", "json",
		"--id-gt", "0",
		"--id-gte", "0",
		"--id-lt", "0",
		"--id-lte", "0",
		"--id-range", "id__range",
		"--nature-of-suit", "nature_of_suit",
		"--omit", "omit",
		"--order-by", "order_by",
		"--page", "1",
		"--source", "0",
	)
}
