# Table: aws_ec2_transit_gateway_multicast_domains

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| options | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| transit_gateway_multicast_domain_arn | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| creation_time | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| transit_gateway_multicast_domain_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


