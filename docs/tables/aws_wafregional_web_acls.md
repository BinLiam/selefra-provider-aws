# Table: aws_wafregional_web_acls

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rules | json | X | √ |  | 
| name | string | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| resources_for_web_acl | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_action | json | X | √ |  | 
| web_acl_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `Web ACL tags.` | 


