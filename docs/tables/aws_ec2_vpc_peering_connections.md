# Table: aws_ec2_vpc_peering_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| accepter_vpc_info | json | X | √ |  | 
| expiration_time | timestamp | X | √ |  | 
| requester_vpc_info | json | X | √ |  | 
| vpc_peering_connection_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status | json | X | √ |  | 
| tags | json | X | √ |  | 


