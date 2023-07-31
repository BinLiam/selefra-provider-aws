# Table: aws_ec2_key_pairs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| key_fingerprint | string | X | √ |  | 
| key_name | string | X | √ |  | 
| key_pair_id | string | X | √ |  | 
| public_key | string | X | √ |  | 
| tags | json | X | √ |  | 
| key_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


