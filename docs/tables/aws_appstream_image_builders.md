# Table: aws_appstream_image_builders

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| access_endpoints | json | X | √ |  | 
| domain_join_info | json | X | √ |  | 
| image_arn | string | X | √ |  | 
| image_builder_errors | json | X | √ |  | 
| instance_type | string | X | √ |  | 
| network_access_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| display_name | string | X | √ |  | 
| state | string | X | √ |  | 
| state_change_reason | json | X | √ |  | 
| vpc_config | json | X | √ |  | 
| description | string | X | √ |  | 
| enable_default_internet_access | bool | X | √ |  | 
| platform | string | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| appstream_agent_version | string | X | √ |  | 
| arn | string | √ | √ |  | 
| iam_role_arn | string | X | √ |  | 


