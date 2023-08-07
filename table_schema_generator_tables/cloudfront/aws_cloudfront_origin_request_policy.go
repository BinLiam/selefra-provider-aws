package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsCloudfrontOriginRequestPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudfrontOriginRequestPolicyGenerator{}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetTableName() string {
	return "aws_cloudfront_origin_request_policy"
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"id"},
	}
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.OriginRequestPolicyConfig.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.OriginRequestPolicyConfig.Comment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cookies_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.OriginRequestPolicyConfig.CookiesConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("headers_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.OriginRequestPolicyConfig.HeadersConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_strings_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("OriginRequestPolicy.OriginRequestPolicyConfig.QueryStringsConfig")).Build(),
	}
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}

type warpCloudFrontOriginRequestPolicy struct {
	types.OriginRequestPolicySummary
	Etag *string
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudfront
			var nextMarker *string
			input := cloudfront.ListOriginRequestPoliciesInput{}
			for {
				if nextMarker != nil {
					input.Marker = nextMarker
				}
				output, err := svc.ListOriginRequestPolicies(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, originRequestPolicy := range output.OriginRequestPolicyList.Items {

					wp := warpCloudFrontOriginRequestPolicy{
						OriginRequestPolicySummary: originRequestPolicy,
					}
					func() {
						params := &cloudfront.GetOriginRequestPolicyInput{
							Id: originRequestPolicy.OriginRequestPolicy.Id,
						}
						op, err := svc.GetOriginRequestPolicy(ctx, params)
						if err != nil {
							return
						}
						wp.Etag = op.ETag
					}()
					resultChannel <- wp
				}
				if output.OriginRequestPolicyList.NextMarker == nil {
					break
				}
				nextMarker = output.OriginRequestPolicyList.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsCloudfrontOriginRequestPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}
