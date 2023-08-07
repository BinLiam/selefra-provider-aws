package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsBackupProtectedResourceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupProtectedResourceGenerator{}

func (x *TableAwsBackupProtectedResourceGenerator) GetTableName() string {
	return "aws_backup_protected_resource"
}

func (x *TableAwsBackupProtectedResourceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupProtectedResourceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupProtectedResourceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("resource_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ResourceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ResourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_backup_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastBackupTime")).Build(),
	}
}

func (x *TableAwsBackupProtectedResourceGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsBackupProtectedResourceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{PrimaryKeys: []string{"resource_arn"}}
}

func (x *TableAwsBackupProtectedResourceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Backup
			var nextToken *string
			var input backup.ListProtectedResourcesInput

			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListProtectedResources(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, resource := range output.Results {
					resultChannel <- resource
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

func (x *TableAwsBackupProtectedResourceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}