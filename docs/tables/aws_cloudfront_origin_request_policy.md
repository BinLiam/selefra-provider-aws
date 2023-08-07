# Table: aws_cloudfront_origin_request_policy

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| id | string | √ | √ |  | 
| comment | string | X | √ |  | 
| etag | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| cookies_config | json | X | √ |  | 
| headers_config | json | X | √ |  | 
| query_strings_config | json | X | √ |  | 


