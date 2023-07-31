# Table: aws_ses_configuration_sets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| sending_options | json | X | √ |  | 
| suppression_options | json | X | √ |  | 
| tags | json | X | √ |  | 
| vdm_options | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| configuration_set_name | string | X | √ |  | 
| delivery_options | json | X | √ |  | 
| reputation_options | json | X | √ |  | 
| tracking_options | json | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 


