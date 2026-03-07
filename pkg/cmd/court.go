// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/court-listener-sdk-cli/internal/apiquery"
	"github.com/stainless-sdks/court-listener-sdk-cli/internal/requestflag"
	"github.com/stainless-sdks/court-listener-sdk-go"
	"github.com/stainless-sdks/court-listener-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var courtsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single court",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
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
	Action:          handleCourtsRetrieve,
	HideHelpCommand: true,
}

var courtsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of courts. Results can generally be cached as court\ndata changes infrequently.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Filter by court identifier (e.g. `scotus`, `ca9`, `dcd`).",
			QueryPath: "id",
		},
		&requestflag.Flag[string]{
			Name:      "count",
			Usage:     "Set to `on` to return only the total count of matching items without\nresult data. When enabled, pagination parameters are ignored.\n",
			QueryPath: "count",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor token for deep pagination. Returned in the `next` / `previous`\nfields of paginated responses. Available when ordering by `id`,\n`date_modified`, or `date_created`.\n",
			QueryPath: "cursor",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified",
			Usage:     "Filter by exact date modified (ISO-8601).",
			QueryPath: "date_modified",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified-gte",
			Usage:     "Filter courts modified on or after this date.",
			QueryPath: "date_modified__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified-lte",
			Usage:     "Filter courts modified on or before this date.",
			QueryPath: "date_modified__lte",
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
			Name:      "full-name",
			Usage:     "Filter by the full name of the court.",
			QueryPath: "full_name",
		},
		&requestflag.Flag[string]{
			Name:      "full-name-startswith",
			Usage:     "Filter courts whose full name starts with the given value.",
			QueryPath: "full_name__startswith",
		},
		&requestflag.Flag[string]{
			Name:      "id-in",
			Usage:     "Filter by multiple court identifiers (comma-separated).",
			QueryPath: "id__in",
		},
		&requestflag.Flag[string]{
			Name:      "jurisdiction",
			Usage:     "Filter by jurisdiction type. Common values:\n`F` (Federal Appellate), `FD` (Federal District),\n`FB` (Federal Bankruptcy), `FBP` (Federal Bankruptcy Panel),\n`FS` (Federal Special), `S` (State Supreme),\n`SA` (State Appellate), `ST` (State Trial),\n`SS` (State Special), `SAG` (State Attorney General),\n`T` (Tribal), `I` (International), `C` (Committee),\n`TES` (Testing).\n",
			QueryPath: "jurisdiction",
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
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleCourtsList,
	HideHelpCommand: true,
}

func handleCourtsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.CourtGetParams{}

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
	_, err = client.Courts.Get(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "courts retrieve", obj, format, transform)
}

func handleCourtsList(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.CourtListParams{}

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
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Courts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "courts list", obj, format, transform)
	} else {
		iter := client.Courts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "courts list", iter, format, transform, maxItems)
	}
}
