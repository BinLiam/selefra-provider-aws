# Table: aws_frauddetector_model_versions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| training_result | json | X | √ |  | 
| external_events_detail | json | X | √ |  | 
| ingested_events_detail | json | X | √ |  | 
| last_updated_time | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_frauddetector_models_selefra_id | string | X | X | fk to aws_frauddetector_selefra_id | 
| arn | string | √ | √ |  | 
| created_time | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| training_data_schema | json | X | √ |  | 
| training_data_source | string | X | √ |  | 
| training_result_v2 | json | X | √ |  | 
| account_id | string | X | √ |  | 
| model_id | string | X | √ |  | 
| model_type | string | X | √ |  | 
| model_version_number | string | X | √ |  | 


