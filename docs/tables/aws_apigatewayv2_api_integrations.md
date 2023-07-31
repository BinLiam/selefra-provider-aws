# Table: aws_apigatewayv2_api_integrations

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| connection_id | string | X | √ |  | 
| region | string | X | √ |  | 
| credentials_arn | string | X | √ |  | 
| description | string | X | √ |  | 
| integration_response_selection_expression | string | X | √ |  | 
| response_parameters | json | X | √ |  | 
| content_handling_strategy | string | X | √ |  | 
| integration_method | string | X | √ |  | 
| integration_subtype | string | X | √ |  | 
| integration_uri | string | X | √ |  | 
| request_parameters | json | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| payload_format_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| connection_type | string | X | √ |  | 
| integration_id | string | X | √ |  | 
| integration_type | string | X | √ |  | 
| request_templates | json | X | √ |  | 
| timeout_in_millis | big_int | X | √ |  | 
| tls_config | json | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| passthrough_behavior | string | X | √ |  | 
| template_selection_expression | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


