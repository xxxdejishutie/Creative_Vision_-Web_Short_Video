一个服务器创建一个文件夹，再去写


定义包格式如下：

视频标签包{
    包类型|视频标签
}

URL回复包{
    包类型|url数量|url_1|url_2|...
}

用户收藏包{
    包类型|用户名
}


包类型码定义如下
const FundamentalNumber = 10//基础号
const TYPE_LABEL = FundamentalNumber + 1//标签视频类型
const TYPE_RET_URL = FundamentalNumber + 2//URL回复类型
const TYPE_USER_COL = FundamentalNumber + 3//用户收藏包类型