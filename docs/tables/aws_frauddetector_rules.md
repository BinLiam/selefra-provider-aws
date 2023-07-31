# Table: aws_frauddetector_rules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| language | string | X | √ |  | 
| outcomes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| last_updated_time | string | X | √ |  | 
| expression | string | X | √ |  | 
| created_time | string | X | √ |  | 
| detector_id | string | X | √ |  | 
| rule_id | string | X | √ |  | 
| rule_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_frauddetector_detectors_selefra_id | string | X | X | fk to aws_frauddetector_detectors.selefra_id | 
| arn | string | √ | √ |  | 


