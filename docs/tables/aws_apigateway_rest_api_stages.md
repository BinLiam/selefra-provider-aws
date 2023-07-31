# Table: aws_apigateway_rest_api_stages

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deployment_id | string | X | √ |  | 
| description | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| cache_cluster_size | string | X | √ |  | 
| canary_settings | json | X | √ |  | 
| documentation_version | string | X | √ |  | 
| variables | json | X | √ |  | 
| stage_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| access_log_settings | json | X | √ |  | 
| cache_cluster_status | string | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| cache_cluster_enabled | bool | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| method_settings | json | X | √ |  | 
| tracing_enabled | bool | X | √ |  | 


