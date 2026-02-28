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

var opinionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Look up an opinion by its ID. Note that opinion IDs do **not** reliably match\ncluster IDs. If you have a CourtListener case URL, use the cluster API to look\nit up.",
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
	Action:          handleOpinionsRetrieve,
	HideHelpCommand: true,
}

var opinionsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of opinions. Each opinion contains the text of a\njudicial decision and metadata about the authoring judge.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Usage:     "Filter by opinion ID.",
			QueryPath: "id",
		},
		&requestflag.Flag[int64]{
			Name:      "cited-opinion",
			Usage:     "Filter opinions that cite this opinion ID.",
			QueryPath: "cited_opinion",
		},
		&requestflag.Flag[int64]{
			Name:      "cluster",
			Usage:     "Filter by parent cluster ID.",
			QueryPath: "cluster",
		},
		&requestflag.Flag[string]{
			Name:      "cluster-docket-court",
			Usage:     "Filter by court via the cluster's docket (e.g. `scotus`).",
			QueryPath: "cluster__docket__court",
		},
		&requestflag.Flag[string]{
			Name:      "cluster-docket-docket-number",
			Usage:     "Filter by docket number via the cluster's docket.",
			QueryPath: "cluster__docket__docket_number",
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
			Name:      "date-created",
			QueryPath: "date_created",
		},
		&requestflag.Flag[any]{
			Name:      "date-created-gte",
			QueryPath: "date_created__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-created-lte",
			QueryPath: "date_created__lte",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified",
			QueryPath: "date_modified",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified-gte",
			QueryPath: "date_modified__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-modified-lte",
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
		&requestflag.Flag[int64]{
			Name:      "id-gt",
			QueryPath: "id__gt",
		},
		&requestflag.Flag[int64]{
			Name:      "id-gte",
			QueryPath: "id__gte",
		},
		&requestflag.Flag[int64]{
			Name:      "id-lt",
			QueryPath: "id__lt",
		},
		&requestflag.Flag[int64]{
			Name:      "id-lte",
			QueryPath: "id__lte",
		},
		&requestflag.Flag[string]{
			Name:      "id-range",
			QueryPath: "id__range",
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
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter by opinion type. Values are prefixed with numbers for sort order.\nCommon types include combined opinion, lead opinion, concurrence, dissent, etc.\n",
			QueryPath: "type",
		},
	},
	Action:          handleOpinionsList,
	HideHelpCommand: true,
}

func handleOpinionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.OpinionGetParams{}

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
	_, err = client.Opinions.Get(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "opinions retrieve", obj, format, transform)
}

func handleOpinionsList(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.OpinionListParams{}

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
		_, err = client.Opinions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "opinions list", obj, format, transform)
	} else {
		iter := client.Opinions.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "opinions list", iter, format, transform)
	}
}
