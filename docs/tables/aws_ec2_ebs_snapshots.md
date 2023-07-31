# Table: aws_ec2_ebs_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| snapshot_id | string | X | √ |  | 
| storage_tier | string | X | √ |  | 
| tags | json | X | √ |  | 
| state | string | X | √ |  | 
| volume_size | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| attribute | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| restore_expiry_time | timestamp | X | √ |  | 
| state_message | string | X | √ |  | 
| arn | string | √ | √ |  | 
| data_encryption_key_id | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| owner_alias | string | X | √ |  | 
| progress | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| volume_id | string | X | √ |  | 


