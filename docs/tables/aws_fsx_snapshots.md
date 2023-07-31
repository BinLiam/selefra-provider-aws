# Table: aws_fsx_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle_transition_reason | json | X | √ |  | 
| resource_arn | string | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| administrative_actions | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 


