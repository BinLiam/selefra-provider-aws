# Table: aws_apigatewayv2_api_deployments

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deployment_status_message | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| auto_deployed | bool | X | √ |  | 
| deployment_id | string | X | √ |  | 
| deployment_status | string | X | √ |  | 
| description | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| created_date | timestamp | X | √ |  | 


