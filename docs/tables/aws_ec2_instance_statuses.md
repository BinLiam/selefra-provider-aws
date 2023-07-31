# Table: aws_ec2_instance_statuses

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| availability_zone | string | X | √ |  | 
| instance_status | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| outpost_arn | string | X | √ |  | 
| system_status | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| events | json | X | √ |  | 
| instance_id | string | X | √ |  | 
| instance_state | json | X | √ |  | 


