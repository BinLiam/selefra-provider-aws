package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsNetworkfirewallFirewallPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNetworkfirewallFirewallPolicyGenerator{}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetTableName() string {
	return "aws_networkfirewall_firewall_policy"
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.FirewallPolicyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.FirewallPolicyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("firewall_policy_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.FirewallPolicyId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("consumed_stateful_rule_capacity").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.ConsumedStatefulRuleCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("consumed_stateless_rule_capacity").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.ConsumedStatelessRuleCapacity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("firewall_policy_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.FirewallPolicyStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_associations").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.NumberOfAssociations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_configuration").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.EncryptionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("firewall_policy").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FirewallPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FirewallPolicyResponse.Tags")).Build(),
	}
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Networkfirewall
			var nextToken *string
			input := networkfirewall.ListFirewallPoliciesInput{}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListFirewallPolicies(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, firewallPolicy := range output.FirewallPolicies {
					params := &networkfirewall.DescribeFirewallPolicyInput{
						FirewallPolicyArn:  firewallPolicy.Arn,
						FirewallPolicyName: firewallPolicy.Name,
					}
					data, err := svc.DescribeFirewallPolicy(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- data
				}
			}
		},
	}
}

func (x *TableAwsNetworkfirewallFirewallPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("network-firewall")

}
