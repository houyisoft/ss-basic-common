#
BloomFilter的应用  
黑名单  
比如邮件黑名单过滤器，判断邮件地址是否在黑名单中
排序(仅限于BitSet)
仔细想想，其实BitSet在set(int value)的时候，“顺便”把value也给排序了。
网络爬虫
判断某个URL是否已经被爬取过
K-V系统快速判断某个key是否存在
典型的例子有Hbase，Hbase的每个Region中都包含一个BloomFilter，用于在查询时快速判断某个key在该region中是否存在，如果不存在，直接返回，节省掉后续的查询。
