# Table: aws_directconnect_gateway_attachments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| virtual_interface_id | string | X | √ |  | 
| virtual_interface_owner_account | string | X | √ |  | 
| region | string | X | √ |  | 
| gateway_arn | string | X | √ |  | 
| gateway_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_directconnect_gateways_selefra_id | string | X | X | fk to aws_directconnect_gateways.selefra_id | 
| attachment_state | string | X | √ |  | 
| attachment_type | string | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| state_change_error | string | X | √ |  | 
| virtual_interface_region | string | X | √ |  | 


