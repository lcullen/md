
常规操作
1. mapping setting
    * 对数组格式的fields 进行mapping index 设置
    * analysis: 预处理扫描加工
        * 精确值  keyword(term): 查询
        * 全文本分词处理: 
    * tokenizer: 分词器
        * token filter: token 完之后进行后续处理
        
2. query
    term:
    text:
        {
            "query": {
                "match": {
                    "title": {
                        "query": "life"
                    }
                }
            }
        }


运维相关

### es 的近实时搜索的底层实现
