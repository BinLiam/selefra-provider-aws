package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsGlacierVaultNotificationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlacierVaultNotificationsGenerator{}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetTableName() string {
	return "aws_glacier_vault_notifications"
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"vault_arn",
		},
	}
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Glacier
			p := task.ParentRawResult.(types.DescribeVaultOutput)

			response, err := svc.GetVaultNotifications(ctx, &glacier.GetVaultNotificationsInput{
				VaultName: p.VaultName,
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.VaultNotificationConfig
			return nil
		},
	}
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glacier")
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vault_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_glacier_vaults_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_glacier_vaults.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Events")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sns_topic").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SNSTopic")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsGlacierVaultNotificationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
