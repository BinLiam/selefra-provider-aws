# Table: aws_ssm_instance_patches

## Primary Keys 

```
account_id, region, kb_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kb_id | string | X | √ |  | 
| cve_ids | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| aws_ssm_instances_selefra_id | string | X | X | fk to aws_ssm_instances.selefra_id | 
| classification | string | X | √ |  | 
| installed_time | timestamp | X | √ |  | 
| severity | string | X | √ |  | 
| state | string | X | √ |  | 
| title | string | X | √ |  | 
| account_id | string | X | √ |  | 


