# Table: aws_elastictranscoder_pipeline_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| aws_elastictranscoder_pipelines_selefra_id | string | X | X | fk to aws_elastictranscoder_pipelines.selefra_id | 
| inputs | json | X | √ |  | 
| output | json | X | √ |  | 
| playlists | json | X | √ |  | 
| status | string | X | √ |  | 
| timing | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| user_metadata | json | X | √ |  | 
| pipeline_id | string | X | √ |  | 
| id | string | X | √ |  | 
| input | json | X | √ |  | 
| output_key_prefix | string | X | √ |  | 
| outputs | json | X | √ |  | 


