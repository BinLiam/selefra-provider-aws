package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsNetworkfirewallRuleGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNetworkfirewallRuleGroupGenerator{}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetTableName() string {
	return "aws_networkfirewall_rule_group"
}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("rule_group_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.Capacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("consumed_capacity").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.ConsumedCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_associations").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.NumberOfAssociations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_group_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.RuleGroupId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_group_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.RuleGroupStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_variables").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RuleGroup.RuleVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules_source").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RuleGroup.RulesSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stateful_rule_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RuleGroup.StatefulRuleOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RuleGroupResponse.Tags")).Build(),
	}
}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}

type warpNetworkFirewallRuleGroup struct {
	Name *string
	Arn  *string
	*networkfirewall.DescribeRuleGroupOutput
}

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Networkfirewall
			var nextToken *string
			input := networkfirewall.ListRuleGroupsInput{}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListRuleGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, ruleGroup := range output.RuleGroups {
					params := &networkfirewall.DescribeRuleGroupInput{
						RuleGroupArn:  ruleGroup.Arn,
						RuleGroupName: ruleGroup.Name,
					}
					data, err := svc.DescribeRuleGroup(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- warpNetworkFirewallRuleGroup{
						Name:                    ruleGroup.Name,
						Arn:                     ruleGroup.Arn,
						DescribeRuleGroupOutput: data,
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

func (x *TableAwsNetworkfirewallRuleGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("network-firewall")
}
