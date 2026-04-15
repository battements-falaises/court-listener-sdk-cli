// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/battements-falaises/court-listener-sdk-cli/internal/apiquery"
	"github.com/battements-falaises/court-listener-sdk-cli/internal/requestflag"
	"github.com/battements-falaises/court-listener-sdk-go"
	"github.com/battements-falaises/court-listener-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var docketsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single docket",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "fields",
			Usage:     "Comma-separated list of fields to include. Supports nested fields via\ndouble-underscore notation (e.g. `educations__id`).\n",
			QueryPath: "fields",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response serialization format. JSON is default when no `Accept` header\nis provided.\n",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:      "omit",
			Usage:     "Comma-separated list of fields to exclude. Supports nested fields via\ndouble-underscore notation.\n",
			QueryPath: "omit",
		},
	},
	Action:          handleDocketsRetrieve,
	HideHelpCommand: true,
}

var docketsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of dockets. Dockets sit at the top of the case law\nhierarchy, linking to clusters of opinions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Usage:     "Filter by docket ID (exact).",
			QueryPath: "id",
		},
		&requestflag.Flag[bool]{
			Name:      "blocked",
			Usage:     "Filter for blocked/unblocked dockets.",
			QueryPath: "blocked",
		},
		&requestflag.Flag[string]{
			Name:      "case-name",
			Usage:     "Filter by case name.",
			QueryPath: "case_name",
		},
		&requestflag.Flag[string]{
			Name:      "cause",
			Usage:     "Filter by cause.",
			QueryPath: "cause",
		},
		&requestflag.Flag[string]{
			Name:      "count",
			Usage:     "Set to `on` to return only the total count of matching items without\nresult data. When enabled, pagination parameters are ignored.\n",
			QueryPath: "count",
		},
		&requestflag.Flag[string]{
			Name:      "court",
			Usage:     "Filter by court identifier (e.g. `scotus`). Supports related court filters via `court__` prefix.",
			QueryPath: "court",
		},
		&requestflag.Flag[string]{
			Name:      "court-jurisdiction",
			Usage:     "Filter by the court's jurisdiction type (e.g. `F`, `FD`, `S`).",
			QueryPath: "court__jurisdiction",
		},
		&requestflag.Flag[string]{
			Name:      "court-jurisdiction",
			Usage:     "Exclude dockets from this jurisdiction type.",
			QueryPath: "court__jurisdiction!",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor token for deep pagination. Returned in the `next` / `previous`\nfields of paginated responses. Available when ordering by `id`,\n`date_modified`, or `date_created`.\n",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "date-created",
			Usage:     "Filter by exact creation date.",
			QueryPath: "date_created",
		},
		&requestflag.Flag[any]{
			Name:      "date-created-gte",
			Usage:     "Created on or after this date.",
			QueryPath: "date_created__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-created-lte",
			Usage:     "Created on or before this date.",
			QueryPath: "date_created__lte",
		},
		&requestflag.Flag[any]{
			Name:      "date-filed",
			Usage:     "Filter by filing date.",
			QueryPath: "date_filed",
		},
		&requestflag.Flag[any]{
			Name:      "date-filed-gte",
			Usage:     "Filed on or after this date.",
			QueryPath: "date_filed__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-filed-lte",
			Usage:     "Filed on or before this date.",
			QueryPath: "date_filed__lte",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified",
			Usage:     "Filter by exact modification date.",
			QueryPath: "date_modified",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified-gte",
			Usage:     "Modified on or after this date.",
			QueryPath: "date_modified__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified-lte",
			Usage:     "Modified on or before this date.",
			QueryPath: "date_modified__lte",
		},
		&requestflag.Flag[any]{
			Name:      "date-terminated",
			Usage:     "Filter by termination date.",
			QueryPath: "date_terminated",
		},
		&requestflag.Flag[any]{
			Name:      "date-terminated-gte",
			QueryPath: "date_terminated__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-terminated-lte",
			QueryPath: "date_terminated__lte",
		},
		&requestflag.Flag[string]{
			Name:      "docket-number",
			Usage:     "Filter by exact docket number (e.g. `23A994`).",
			QueryPath: "docket_number",
		},
		&requestflag.Flag[string]{
			Name:      "fields",
			Usage:     "Comma-separated list of fields to include. Supports nested fields via\ndouble-underscore notation (e.g. `educations__id`).\n",
			QueryPath: "fields",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response serialization format. JSON is default when no `Accept` header\nis provided.\n",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:      "id-gt",
			Usage:     "Docket IDs greater than this value.",
			QueryPath: "id__gt",
		},
		&requestflag.Flag[int64]{
			Name:      "id-gte",
			Usage:     "Docket IDs greater than or equal to this value.",
			QueryPath: "id__gte",
		},
		&requestflag.Flag[int64]{
			Name:      "id-lt",
			Usage:     "Docket IDs less than this value.",
			QueryPath: "id__lt",
		},
		&requestflag.Flag[int64]{
			Name:      "id-lte",
			Usage:     "Docket IDs less than or equal to this value.",
			QueryPath: "id__lte",
		},
		&requestflag.Flag[string]{
			Name:      "id-range",
			Usage:     "Docket IDs within an inclusive range (e.g. `500,1000`).",
			QueryPath: "id__range",
		},
		&requestflag.Flag[string]{
			Name:      "nature-of-suit",
			Usage:     "Filter by nature of suit.",
			QueryPath: "nature_of_suit",
		},
		&requestflag.Flag[string]{
			Name:      "omit",
			Usage:     "Comma-separated list of fields to exclude. Supports nested fields via\ndouble-underscore notation.\n",
			QueryPath: "omit",
		},
		&requestflag.Flag[string]{
			Name:      "order-by",
			Usage:     "Comma-separated list of fields to order by. Prefix with `-` for\ndescending order. Use a secondary field as a tie-breaker for\ndeterministic ordering (e.g. `date_filed,id`).\n",
			QueryPath: "order_by",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number for standard pagination (limited to 100 pages).",
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "source",
			Usage:     "Filter by docket source.",
			QueryPath: "source",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleDocketsList,
	HideHelpCommand: true,
}

func handleDocketsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.DocketGetParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Dockets.Get(
		ctx,
		cmd.Value("id").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "dockets retrieve", obj, format, explicitFormat, transform)
}

func handleDocketsList(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.DocketListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Dockets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "dockets list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Dockets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "dockets list", iter, format, explicitFormat, transform, maxItems)
	}
}
