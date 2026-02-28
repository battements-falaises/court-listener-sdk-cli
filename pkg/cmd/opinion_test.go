// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/court-listener-sdk-cli/internal/mocktest"
)

func TestOpinionsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"opinions", "retrieve",
		"--api-key", "string",
		"--id", "0",
		"--fields", "fields",
		"--format", "json",
		"--omit", "omit",
	)
}

func TestOpinionsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"opinions", "list",
		"--api-key", "string",
		"--id", "0",
		"--cited-opinion", "0",
		"--cluster", "0",
		"--cluster-docket-court", "cluster__docket__court",
		"--cluster-docket-docket-number", "cluster__docket__docket_number",
		"--count", "on",
		"--cursor", "cursor",
		"--date-created", "'2019-12-27T18:11:19.117Z'",
		"--date-created-gte", "'2019-12-27T18:11:19.117Z'",
		"--date-created-lte", "'2019-12-27T18:11:19.117Z'",
		"--date-modified", "'2019-12-27T18:11:19.117Z'",
		"--date-modified-gte", "'2019-12-27T18:11:19.117Z'",
		"--date-modified-lte", "'2019-12-27T18:11:19.117Z'",
		"--fields", "fields",
		"--format", "json",
		"--id-gt", "0",
		"--id-gte", "0",
		"--id-lt", "0",
		"--id-lte", "0",
		"--id-range", "id__range",
		"--omit", "omit",
		"--order-by", "order_by",
		"--page", "1",
		"--type", "type",
	)
}
