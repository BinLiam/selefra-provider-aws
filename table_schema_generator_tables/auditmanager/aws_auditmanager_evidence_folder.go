package auditmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAuditmanagerEvidenceFolderGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAuditmanagerEvidenceFolderGenerator{}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetTableName() string {
	return "aws_auditmanager_evidence_folder"
}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"id"},
	}
}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(warpAssessment)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Auditmanager
			params := auditmanager.GetEvidenceFoldersByAssessmentInput{
				AssessmentId: r.Id,
			}
			var nextToken *string
			for {
				if nextToken != nil {
					params.NextToken = nextToken
				}
				output, err := svc.GetEvidenceFoldersByAssessment(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, evidenceFolder := range output.EvidenceFolders {
					resultChannel <- evidenceFolder
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

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_auditmanager_assessment_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
			task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

			extractor := func() (any, error) {
				return client.(*aws_client.Client).ARN("auditmanager", "evidence-folder", *result.(types.AssessmentEvidenceFolder).Id), nil
			}
			extractResultValue, err := extractor()
			if err != nil {
				return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
			} else {
				return extractResultValue, nil
			}
		})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assessment_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AssessmentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_set_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ControlSetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assessment_report_selection_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("AssessmentReportSelectionCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("author").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Author")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ControlId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ControlName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_source").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DataSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("Date")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_aws_service_source_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceAwsServiceSourceCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_by_type_compliance_check_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceByTypeComplianceCheckCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_by_type_compliance_check_issues_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceByTypeComplianceCheckIssuesCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_by_type_configuration_data_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceByTypeConfigurationDataCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_by_type_manual_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceByTypeManualCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_by_type_user_activity_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceByTypeUserActivityCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_resources_included_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("EvidenceResourcesIncludedCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_evidence").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("TotalEvidence")).Build(),
	}
}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("auditmanager")
}

func (x *TableAwsAuditmanagerEvidenceFolderGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAuditmanagerEvidenceGenerator{}),
	}
}
