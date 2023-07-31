# Table: aws_ssm_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_association_execution_date | timestamp | X | √ |  | 
| last_ping_date_time | timestamp | X | √ |  | 
| ping_status | string | X | √ |  | 
| agent_version | string | X | √ |  | 
| computer_name | string | X | √ |  | 
| is_latest_version | bool | X | √ |  | 
| platform_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| activation_id | string | X | √ |  | 
| last_successful_association_execution_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| platform_version | string | X | √ |  | 
| registration_date | timestamp | X | √ |  | 
| resource_type | string | X | √ |  | 
| source_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| ip_address | string | X | √ |  | 
| iam_role | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| source_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| association_overview | json | X | √ |  | 
| association_status | string | X | √ |  | 
| platform_type | string | X | √ |  | 


