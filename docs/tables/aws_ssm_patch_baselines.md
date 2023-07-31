# Table: aws_ssm_patch_baselines

## Primary Keys 

```
account_id, region, baseline_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_baseline | bool | X | √ |  | 
| operating_system | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| baseline_description | string | X | √ |  | 
| baseline_id | string | X | √ |  | 
| baseline_name | string | X | √ |  | 


