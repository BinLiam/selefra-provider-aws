# Table: aws_apigateway_api_keys

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enabled | bool | X | √ |  | 
| description | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| name | string | X | √ |  | 
| value | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| customer_id | string | X | √ |  | 
| id | string | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| stage_keys | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_date | timestamp | X | √ |  | 


