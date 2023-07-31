# Table: aws_frauddetector_external_models

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| model_endpoint | string | X | √ |  | 
| model_endpoint_status | string | X | √ |  | 
| model_source | string | X | √ |  | 
| output_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| created_time | string | X | √ |  | 
| input_configuration | json | X | √ |  | 
| invoke_model_endpoint_role_arn | string | X | √ |  | 
| last_updated_time | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


