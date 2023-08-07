package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsCloudformationStackSetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudformationStackSetGenerator{}

func (x *TableAwsCloudformationStackSetGenerator) GetTableName() string {
	return "aws_cloudformation_stack_set"
}

func (x *TableAwsCloudformationStackSetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudformationStackSetGenerator) GetVersion() uint64 {
	return 0
}

type warpStackSet struct {
	types.StackSetSummary
	StackSetARN                   *string
	PermissionModel               types.PermissionModels
	AdministrationRoleARN         *string
	ExecutionRoleName             *string
	TemplateBody                  *string
	AutoDeployment                *types.AutoDeployment
	Capabilities                  []types.Capability
	OrganizationalUnitIds         []string
	Parameters                    []types.Parameter
	StackSetDriftDetectionDetails *types.StackSetDriftDetectionDetails
	ManagedExecution              *types.ManagedExecution
	Tags                          map[string]string
}

func (x *TableAwsCloudformationStackSetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("stack_set_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("StackSetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack_set_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("StackSetName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("StackSetARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("drift_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DriftStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_drift_check_timestamp").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastDriftCheckTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission_model").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("PermissionModel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administration_role_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AdministrationRoleARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_role_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ExecutionRoleName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_body").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("TemplateBody")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_deployment").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("AutoDeployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capabilities").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Capabilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organizational_unit_ids").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("OrganizationalUnitIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Parameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack_set_drift_detection_details").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("StackSetDriftDetectionDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_execution").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ManagedExecution")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsCloudformationStackSetGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsCloudformationStackSetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{PrimaryKeys: []string{"arn"}}
}

func (x *TableAwsCloudformationStackSetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudformation
			input := &cloudformation.ListStackSetsInput{}
			var nextToken *string
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListStackSets(ctx, input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, v := range output.Summaries {
					wp := warpStackSet{
						StackSetSummary: v,
						Tags:            make(map[string]string),
					}
					func() {
						params := &cloudformation.DescribeStackSetInput{
							StackSetName: v.StackSetName,
						}
						op, err := svc.DescribeStackSet(ctx, params)
						if err != nil {
							return
						}
						wp.StackSetARN = op.StackSet.StackSetARN
						wp.PermissionModel = op.StackSet.PermissionModel
						wp.AdministrationRoleARN = op.StackSet.AdministrationRoleARN
						wp.ExecutionRoleName = op.StackSet.ExecutionRoleName
						wp.TemplateBody = op.StackSet.TemplateBody
						wp.AutoDeployment = op.StackSet.AutoDeployment
						wp.Capabilities = op.StackSet.Capabilities
						wp.OrganizationalUnitIds = op.StackSet.OrganizationalUnitIds
						wp.Parameters = op.StackSet.Parameters
						wp.StackSetDriftDetectionDetails = op.StackSet.StackSetDriftDetectionDetails
						wp.ManagedExecution = op.StackSet.ManagedExecution
						for _, tag := range op.StackSet.Tags {
							wp.Tags[*tag.Key] = *tag.Value
						}
					}()
					resultChannel <- wp
				}
				if output.NextToken == nil {
					break
				}
				nextToken = output.NextToken
			}
			return nil
		}}
}

func (x *TableAwsCloudformationStackSetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cloudformation")
}
