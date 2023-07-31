# Table: aws_route53_hosted_zone_traffic_policy_instances

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| hosted_zone_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| traffic_policy_version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| message | string | X | √ |  | 
| traffic_policy_id | string | X | √ |  | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| id | string | X | √ |  | 
| arn | string | X | √ | `Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.` | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| traffic_policy_type | string | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| name | string | X | √ |  | 
| ttl | big_int | X | √ |  | 


