# Table: aws_route53_traffic_policy_versions

## Primary Keys 

```
traffic_policy_arn, id, version
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| document | json | X | √ |  | 
| type | string | X | √ |  | 
| comment | string | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_route53_traffic_policies_selefra_id | string | X | X | fk to aws_route53_traffic_policies.selefra_id | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| version | int | X | √ |  | 
| traffic_policy_arn | string | X | √ |  | 


