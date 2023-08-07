package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks/types"

	"github.com/aws/aws-sdk-go-v2/service/eks"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsEksAddonVersionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEksAddonVersionGenerator{}

func (x *TableAwsEksAddonVersionGenerator) GetTableName() string {
	return "aws_eks_addon_version"
}

func (x *TableAwsEksAddonVersionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEksAddonVersionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEksAddonVersionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEksAddonVersionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_eks_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("addon_name").Extractor(column_value_extractor.StructSelector("AddonName")).ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addon_version").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AddonVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addon_configuration").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ConfigurationSchema")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architecture").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Architecture")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compatibilities").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Compatibilities")).Build(),
	}
}

func (x *TableAwsEksAddonVersionGenerator) GetSubTables() []*schema.Table {
	return nil
}

type warpEksAddonConfiguration struct {
	AddonName *string
	types.AddonVersionInfo
	Type                *string
	ConfigurationSchema *string
}

func (x *TableAwsEksAddonVersionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			// DescribeAddonVersions
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Eks
			var nextToken *string
			var input eks.DescribeAddonVersionsInput
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.DescribeAddonVersions(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, addon := range output.Addons {
					params := &eks.DescribeAddonConfigurationInput{
						AddonName:    addon.AddonName,
						AddonVersion: addon.Type,
					}
					op, _ := svc.DescribeAddonConfiguration(ctx, params)
					var configurationSchema *string
					if op != nil {
						configurationSchema = op.ConfigurationSchema
					}

					for _, version := range addon.AddonVersions {
						resultChannel <- warpEksAddonConfiguration{
							AddonName:           addon.AddonName,
							AddonVersionInfo:    version,
							Type:                addon.Type,
							ConfigurationSchema: configurationSchema,
						}
					}
				}
				if output.NextToken == nil {
					break
				}
				nextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEksAddonVersionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("eks")
}
