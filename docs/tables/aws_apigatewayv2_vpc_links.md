# Table: aws_apigatewayv2_vpc_links

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| vpc_link_version | string | X | √ |  | 
| arn | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| security_group_ids | string_array | X | √ |  | 
| vpc_link_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| vpc_link_status | string | X | √ |  | 
| vpc_link_status_message | string | X | √ |  | 
| account_id | string | X | √ |  | 


