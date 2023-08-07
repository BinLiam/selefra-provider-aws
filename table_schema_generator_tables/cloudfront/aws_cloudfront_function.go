package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsCloudfrontFunctionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudfrontFunctionGenerator{}

func (x *TableAwsCloudfrontFunctionGenerator) GetTableName() string {
	return "aws_cloudfront_function"
}

func (x *TableAwsCloudfrontFunctionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudfrontFunctionGenerator) GetVersion() uint64 {
	return 0
}

type warpFunction struct {
	ETag           *string
	FunctionConfig *types.FunctionConfig

	FunctionMetadata *types.FunctionMetadata
	Name             *string
	Status           *string

	FunctionARN *string
}

func (x *TableAwsCloudfrontFunctionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FunctionARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("e_tag").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ETag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FunctionConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_metadata").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FunctionMetadata")).Build(),
	}
}

func (x *TableAwsCloudfrontFunctionGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsCloudfrontFunctionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{PrimaryKeys: []string{"arn"}}
}

func (x *TableAwsCloudfrontFunctionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudfront
			input := cloudfront.ListFunctionsInput{}
			var nextMarker *string
			for {
				if nextMarker != nil {
					input.Marker = nextMarker
				}
				output, err := svc.ListFunctions(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, v := range output.FunctionList.Items {
					wp := &warpFunction{
						FunctionConfig:   v.FunctionConfig,
						FunctionMetadata: v.FunctionMetadata,
						Name:             v.Name,
						Status:           v.Status,
					}
					if v.FunctionMetadata != nil {
						wp.FunctionARN = v.FunctionMetadata.FunctionARN
					}
					func() {
						params := &cloudfront.DescribeFunctionInput{
							Name: v.Name,
						}
						data, err := svc.DescribeFunction(ctx, params)
						if err != nil {
							return
						}
						wp.ETag = data.ETag
						if wp.FunctionConfig == nil {
							wp.FunctionConfig = data.FunctionSummary.FunctionConfig
						}
						if wp.FunctionMetadata == nil {
							wp.FunctionMetadata = data.FunctionSummary.FunctionMetadata
						}
						if wp.Name == nil {
							wp.Name = data.FunctionSummary.Name
						}
						if wp.Status == nil {
							wp.Status = data.FunctionSummary.Status
						}
						if wp.FunctionARN == nil && data.FunctionSummary.FunctionMetadata != nil {
							wp.FunctionARN = data.FunctionSummary.FunctionMetadata.FunctionARN
						}
					}()

					resultChannel <- wp
				}

				if output.FunctionList.NextMarker == nil {
					break
				}
				nextMarker = output.FunctionList.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsCloudfrontFunctionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}
