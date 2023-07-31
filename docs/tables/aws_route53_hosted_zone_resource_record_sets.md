# Table: aws_route53_hosted_zone_resource_record_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| set_identifier | string | X | √ |  | 
| traffic_policy_instance_id | string | X | √ |  | 
| weight | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| hosted_zone_arn | string | X | √ |  | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| type | string | X | √ |  | 
| alias_target | json | X | √ |  | 
| cidr_routing_config | json | X | √ |  | 
| multi_value_answer | bool | X | √ |  | 
| region | string | X | √ |  | 
| health_check_id | string | X | √ |  | 
| ttl | big_int | X | √ |  | 
| name | string | X | √ |  | 
| failover | string | X | √ |  | 
| geo_location | json | X | √ |  | 
| resource_records | json | X | √ |  | 


