# Table: aws_ec2_transit_gateway_peering_attachments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| region | string | X | √ |  | 
| accepter_transit_gateway_attachment_id | string | X | √ |  | 
| options | json | X | √ |  | 
| state | string | X | √ |  | 
| transit_gateway_attachment_id | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| accepter_tgw_info | json | X | √ |  | 
| requester_tgw_info | json | X | √ |  | 
| status | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


