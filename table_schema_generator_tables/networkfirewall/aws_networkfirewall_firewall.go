package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsNetworkfirewallFirewallGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNetworkfirewallFirewallGenerator{}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetTableName() string {
	return "aws_networkfirewall_firewall"
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Firewall.FirewallArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Firewall.FirewallName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Firewall.VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_protection").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Firewall.DeleteProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Firewall.Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Firewall.FirewallId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Firewall.FirewallPolicyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_change_protection").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Firewall.FirewallPolicyChangeProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_change_protection").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Firewall.SubnetChangeProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_configuration").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Firewall.EncryptionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("firewall_status").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FirewallStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_mappings").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Firewall.SubnetMappings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags_src").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Firewall.Tags")).Build(),
	}
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Networkfirewall
			var nextToken *string
			input := networkfirewall.ListFirewallsInput{}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListFirewalls(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, firewall := range output.Firewalls {
					params := &networkfirewall.DescribeFirewallInput{
						FirewallArn:  firewall.FirewallArn,
						FirewallName: firewall.FirewallName,
					}
					data, err := svc.DescribeFirewall(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- data

				}
				if output.NextToken == nil {
					break
				}
				nextToken = output.NextToken
			}
			return nil
		}}
}

func (x *TableAwsNetworkfirewallFirewallGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("network-firewall")
}
