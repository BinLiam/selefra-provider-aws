# Table: aws_apigatewayv2_api_integration_responses

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| response_templates | json | X | √ |  | 
| region | string | X | √ |  | 
| api_integration_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigatewayv2_api_integrations_selefra_id | string | X | X | fk to aws_apigatewayv2_api_integrations.selefra_id | 
| integration_response_key | string | X | √ |  | 
| content_handling_strategy | string | X | √ |  | 
| integration_response_id | string | X | √ |  | 
| integration_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| response_parameters | json | X | √ |  | 
| template_selection_expression | string | X | √ |  | 
| account_id | string | X | √ |  | 


