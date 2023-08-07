package appconfig

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/aws/aws-sdk-go-v2/service/appconfig"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAppconfigApplicationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAppconfigApplicationGenerator{}

func (x *TableAwsAppconfigApplicationGenerator) GetTableName() string {
	return "aws_appconfig_application"
}

func (x *TableAwsAppconfigApplicationGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAppconfigApplicationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAppconfigApplicationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"id"},
	}
}

func (x *TableAwsAppconfigApplicationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Appconfig
			var input appconfig.ListApplicationsInput
			var nextToken *string
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListApplications(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, app := range output.Items {
					resultChannel <- app
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

func (x *TableAwsAppconfigApplicationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("appconfig")
}

func (x *TableAwsAppconfigApplicationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
	}
}

func (x *TableAwsAppconfigApplicationGenerator) GetSubTables() []*schema.Table {
	return nil
}
