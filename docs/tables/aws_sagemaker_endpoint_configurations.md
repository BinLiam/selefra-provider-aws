# Table: aws_sagemaker_endpoint_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| data_capture_config | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| endpoint_config_arn | string | X | √ |  | 
| endpoint_config_name | string | X | √ |  | 
| async_inference_config | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| explainer_config | json | X | √ |  | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| production_variants | json | X | √ |  | 
| shadow_production_variants | json | X | √ |  | 
| tags | json | X | √ | `The tags associated with the model.` | 
| selefra_id | string | √ | √ | primary keys value md5 | 


