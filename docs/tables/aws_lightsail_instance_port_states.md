# Table: aws_lightsail_instance_port_states

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| from_port | big_int | X | √ |  | 
| protocol | string | X | √ |  | 
| state | string | X | √ |  | 
| to_port | big_int | X | √ |  | 
| region | string | X | √ |  | 
| instance_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_lightsail_instances_selefra_id | string | X | X | fk to aws_lightsail_instances.selefra_id | 
| cidr_list_aliases | string_array | X | √ |  | 
| cidrs | string_array | X | √ |  | 
| ipv6_cidrs | string_array | X | √ |  | 
| account_id | string | X | √ |  | 


