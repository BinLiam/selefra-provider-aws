# Table: aws_shield_attacks

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attack_id | string | X | √ |  | 
| attack_properties | json | X | √ |  | 
| sub_resources | json | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | √ | √ | `The unique identifier (ID) of the attack` | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| attack_counters | json | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| mitigations | json | X | √ |  | 
| resource_arn | string | X | √ |  | 


