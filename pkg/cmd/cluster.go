// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/battements-falaises/court-listener-sdk-cli/internal/apiquery"
	"github.com/battements-falaises/court-listener-sdk-cli/internal/requestflag"
	"github.com/battements-falaises/court-listener-sdk-go"
	"github.com/battements-falaises/court-listener-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var clustersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Look up a cluster by its ID. The cluster ID matches the ID used in CourtListener\ncase law URLs (e.g. `/opinion/2812209/obergefell-v-hodges/` corresponds to\ncluster ID `2812209`).",
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
	Action:          handleClustersRetrieve,
	HideHelpCommand: true,
}

var clustersList = cli.Command{
	Name:    "list",
	Usage:   "Returns a paginated list of opinion clusters. Each cluster groups together\nopinions from the same panel hearing (e.g. majority, dissent, concurrence). The\ncluster `id` is used in CourtListener case law URLs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "id",
			Usage:     "Filter by cluster ID.",
			QueryPath: "id",
		},
		&requestflag.Flag[string]{
			Name:      "citation",
			Usage:     "Filter by citation.",
			QueryPath: "citation",
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
			Name:      "date-filed",
			Usage:     "Filter by the date the cluster was filed.",
			QueryPath: "date_filed",
		},
		&requestflag.Flag[any]{
			Name:      "date-filed-gte",
			QueryPath: "date_filed__gte",
		},
		&requestflag.Flag[any]{
			Name:      "date-filed-lte",
			QueryPath: "date_filed__lte",
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
		&requestflag.Flag[int64]{
			Name:      "docket",
			Usage:     "Filter by parent docket ID.",
			QueryPath: "docket",
		},
		&requestflag.Flag[string]{
			Name:      "docket-court",
			Usage:     "Filter by the court of the parent docket (e.g. `scotus`).",
			QueryPath: "docket__court",
		},
		&requestflag.Flag[string]{
			Name:      "docket-docket-number",
			Usage:     "Filter by the docket number of the parent docket.",
			QueryPath: "docket__docket_number",
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
			Usage:     "Inclusive range (e.g. `100,500`).",
			QueryPath: "id__range",
		},
		&requestflag.Flag[string]{
			Name:      "judges",
			Usage:     "Filter by judge name string.",
			QueryPath: "judges",
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
	Action:          handleClustersList,
	HideHelpCommand: true,
}

func handleClustersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.ClusterGetParams{}

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
	_, err = client.Clusters.Get(
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
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "clusters retrieve",
		Transform:      transform,
	})
}

func handleClustersList(ctx context.Context, cmd *cli.Command) error {
	client := courtlistenersdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := courtlistenersdk.ClusterListParams{}

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
		_, err = client.Clusters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "clusters list",
			Transform:      transform,
		})
	} else {
		iter := client.Clusters.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "clusters list",
			Transform:      transform,
		})
	}
}
