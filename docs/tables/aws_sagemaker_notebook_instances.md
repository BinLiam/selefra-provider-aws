# Table: aws_sagemaker_notebook_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| notebook_instance_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `The tags associated with the notebook instance.` | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| describe_notebook_instance_output | json | X | √ |  | 
| notebook_instance_arn | string | X | √ |  | 


