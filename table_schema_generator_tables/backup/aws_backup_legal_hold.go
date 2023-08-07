package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsBackupLegalHoldGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupLegalHoldGenerator{}

func (x *TableAwsBackupLegalHoldGenerator) GetTableName() string {
	return "aws_backup_legal_hold"
}

func (x *TableAwsBackupLegalHoldGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupLegalHoldGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupLegalHoldGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"arn"},
	}
}

type warpLegalHold struct {
	types.LegalHold
	RecoveryPointSelection *types.RecoveryPointSelection
	RetainRecordUntil      *time.Time
}

func (x *TableAwsBackupLegalHoldGenerator) GetColumns() []*schema.Column {

	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("legal_hold_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("LegalHoldId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("LegalHoldArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cancellation_date").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CancellationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retain_record_until").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("RetainRecordUntil")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recovery_point_selection").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RecoveryPointSelection")).Build(),
	}
}

func (x *TableAwsBackupLegalHoldGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsBackupLegalHoldGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Backup
			var nextToken *string
			var input backup.ListLegalHoldsInput
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListLegalHolds(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, legalHold := range output.LegalHolds {
					wp := &warpLegalHold{
						LegalHold: legalHold,
					}

					func() {
						params := &backup.GetLegalHoldInput{
							LegalHoldId: legalHold.LegalHoldId,
						}
						op, err := svc.GetLegalHold(ctx, params)
						if err != nil {
							return
						}
						wp.RecoveryPointSelection = op.RecoveryPointSelection
						wp.RetainRecordUntil = op.RetainRecordUntil
					}()

					resultChannel <- wp
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

func (x *TableAwsBackupLegalHoldGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}
