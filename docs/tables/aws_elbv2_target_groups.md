# Table: aws_elbv2_target_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| health_check_port | string | X | √ |  | 
| health_check_protocol | string | X | √ |  | 
| healthy_threshold_count | big_int | X | √ |  | 
| matcher | json | X | √ |  | 
| health_check_path | string | X | √ |  | 
| target_group_name | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| health_check_enabled | bool | X | √ |  | 
| health_check_timeout_seconds | big_int | X | √ |  | 
| protocol | string | X | √ |  | 
| target_group_arn | string | X | √ |  | 
| target_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| health_check_interval_seconds | big_int | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| load_balancer_arns | string_array | X | √ |  | 
| port | big_int | X | √ |  | 
| protocol_version | string | X | √ |  | 
| unhealthy_threshold_count | big_int | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 


