# Table: aws_apigateway_usage_plans

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| product_code | string | X | √ |  | 
| quota | json | X | √ |  | 
| throttle | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| api_stages | json | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | X | √ |  | 


