# 测试流程

- 创建两个银行

```text
POST /api/bank
// body：
[
    {
        "Name": "农业银行"
    },
    {
        "Name": "中国银行"
    }
]
```

- 创建一个客户，同时一张身份证、两张银行卡

```text
POST /api/customer
// body：
{
    "Name": "张三",
    "RefIDCard": {
        "IDNum": "440783198902039383"
    },
    "BankCards": [
        {
            "BankID": 1,
            "CardNum": "6000493834828912123"
        },
        {
            "BankID": 2,
            "CardNum": "7978008278234234342"
        }
    ]
}
```

- 修改客户的身份证号

```text
PUT /api/customer
// body：
{
    "cond": {
        "ID": 1
    },
    "doc": {
        "RefIDCard": {
            "ID": 1,
            "IDNum": "500382198902039383"
        }
    }
}
```

- 修改客户的银行卡号

```text
// 单个修改：

PUT /api/customer
// body：
{
    "cond": {
        "ID": 1
    },
    "doc": {
        "BankCards": [
            {
                "ID": 2,
                "CardNum": "2220493834828912333"
            }
        ]
    }
}

// 批量修改：

PUT /api/customer
// body：
{
    "cond": {
        "ID": 1
    },
    "doc": {
        "BankCards": [
            {
                "ID": 1,
                "CardNum": "9998008278234234222"
            },
            {
                "ID": 2,
                "CardNum": "1110493834828912333"
            }
        ]
    }
}
```
