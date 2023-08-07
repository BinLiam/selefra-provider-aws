package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/amplify"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAmplifyApp struct{}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAmplifyApp{}

func (x *TableAwsAmplifyApp) GetTableName() string {
	return "aws_amplify_app"
}

func (x *TableAwsAmplifyApp) GetTableDescription() string {
	return ""
}

func (x *TableAwsAmplifyApp) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAmplifyApp) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"app_id",
		},
	}
}

func (x *TableAwsAmplifyApp) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Amplify
			var input amplify.ListAppsInput
			var nextToken *string
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListApps(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, app := range output.Apps {
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

func (x *TableAwsAmplifyApp) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("amplify")

}

func (x *TableAwsAmplifyApp) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AppId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AppArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("basic_auth_credentials").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("BasicAuthCredentials")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_headers").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("CustomHeaders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_domain").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DefaultDomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_auto_branch_creation").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("EnableAutoBranchCreation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_basic_auth").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("EnableBasicAuth")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_branch_auto_build").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("EnableBranchAutoBuild")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_branch_auto_deletion").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("EnableBranchAutoDeletion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_service_role_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("IamServiceRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Platform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Repository")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repository_clone_method").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("RepositoryCloneMethod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_branch_creation_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("AutoBranchCreationConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_branch_creation_patterns").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("AutoBranchCreationPatterns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_spec").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("BuildSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_rules").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("CustomRules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_variables").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("EnvironmentVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("production_branch").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ProductionBranch")).Build(),
	}
}

func (x *TableAwsAmplifyApp) GetSubTables() []*schema.Table {
	return nil
}
