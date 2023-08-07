package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsCloudfrontResponseHeadersPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudfrontResponseHeadersPolicyGenerator{}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetTableName() string {
	return "aws_cloudfront_response_headers_policy"
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"id"},
	}
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ResponseHeadersPolicy.Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("ResponseHeadersPolicy.LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ETag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("response_headers_policy_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ResponseHeadersPolicy.ResponseHeadersPolicyConfig")).Build(),
	}
}

type warpAwsCloudfrontResponseHeadersPolicy struct {
	types.ResponseHeadersPolicySummary
	Arn  string
	ETag *string
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudfront
			var nextMarker *string
			var input *cloudfront.ListResponseHeadersPoliciesInput
			for {
				if nextMarker != nil {
					input.Marker = nextMarker
				}

				output, err := svc.ListResponseHeadersPolicies(ctx, input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, item := range output.ResponseHeadersPolicyList.Items {
					wp := warpAwsCloudfrontResponseHeadersPolicy{
						ResponseHeadersPolicySummary: item,
						Arn:                          c.ARN("cloudfront", "response-headers-policy", *item.ResponseHeadersPolicy.Id),
					}
					func() {
						params := &cloudfront.GetResponseHeadersPolicyInput{Id: item.ResponseHeadersPolicy.Id}

						data, err := svc.GetResponseHeadersPolicy(ctx, params)
						if err != nil {
							return
						}
						wp.ETag = data.ETag
					}()

					resultChannel <- wp

				}
				if output.ResponseHeadersPolicyList.NextMarker == nil {
					break
				}
				nextMarker = output.ResponseHeadersPolicyList.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsCloudfrontResponseHeadersPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}
