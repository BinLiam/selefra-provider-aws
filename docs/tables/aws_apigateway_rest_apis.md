# Table: aws_apigateway_rest_apis

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| endpoint_configuration | json | X | √ |  | 
| minimum_compression_size | big_int | X | √ |  | 
| warnings | string_array | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| api_key_source | string | X | √ |  | 
| binary_media_types | string_array | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| description | string | X | √ |  | 
| disable_execute_api_endpoint | bool | X | √ |  | 
| name | string | X | √ |  | 
| policy | string | X | √ |  | 


