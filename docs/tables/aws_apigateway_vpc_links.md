# Table: aws_apigateway_vpc_links

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status_message | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| target_arns | string_array | X | √ |  | 


