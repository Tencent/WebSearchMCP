# WebSearchMCP-联网搜索MCP(Model Context Protocol，MCP)

联网搜索MCP(Model Context Protocol，MCP)基于腾讯云产品-[联网搜索API（Web Search API，wsa）](https://cloud.tencent.com/product/wsa?Is=sdk-topnav)封装而成，底层搜索引擎来源于搜狗搜索，以互联网全网公开资源为基础，实现了从收录至召回排序全链路的智能搜索增强。合作伙伴通过检索词请求接口，以JSON形式返回搜索结果对应的排序信息、数据字段，可在业务场景中结合搜索结果的内容，扩充信息内容源，丰富展现效果，提升用户的查询满意度。

## 产品特性

### 毫秒级响应

借助分布式计算+边缘节点缓存，最快300ms返回结果，处于行业领先水平。

### 分钟级更新

提供高时效性内容检索，实时获取政策更新、新闻动态、天气变化、金价油价等最新信息。

### 海量资源库

全网内容分层聚类精准索引，直连搜狗百科、企鹅号、腾讯新闻、腾讯视频等内容，覆盖各类优质资源。

### 多模态覆盖

返回结果包含文本、图片等多元模态，支持多模态VR卡，提供丰富多元的信息形式。

## 产品能力

- 自然结果检索
- 多模态VR卡结果
- 指定网址检索
- 指定时间范围检索
- 标准摘要
- 图片列表
- 结果返回条数5～20条
- 结果返回条数扩充至50条（仅尊享版用户独享）
- 动态摘要（仅尊享版用户独享）
- 指定垂域（仅尊享版用户独享）

## 使用方式

以下为您介绍初次使用联网搜索 MCP(Model Context Protocol，MCP) 需要实施的准备工作及入门操作：

1. 注册腾讯云账号，并开通联网搜索API 服务：

| 操作步骤                                                     | 说明                                    |
| ------------------------------------------------------------ | --------------------------------------- |
| [步骤一：登陆注册](https://cloud.tencent.com/document/product/1806/121802#8270649a-f02c-4f88-ab02-656e5e92894a) | 注册腾讯云账号，完成实名认证并登录      |
| [步骤二：开通服务](https://cloud.tencent.com/document/product/1806/121802#63b4e9ef-8c65-4a87-9169-627941e751a1) | 控制台自助开通联网搜索API               |
| [步骤三：获取云API密钥](https://cloud.tencent.com/document/product/1806/121802#21242563-b79b-4db5-962a-12a9a39ebc16) | 获取云 API 密钥的 SecretId 和 SecretKey |

2. 根据您开通的联网搜索 API 服务版本（尊享版/标准版），选择对应的SSE URL：

- 标准版SSE：https://agent.html5.qq.com/prosearch/sse
-  尊享版SSE：https://agent.html5.qq.com/prosearch_vip/sse

3. 配置说明

   将此配置添加到您的 MCP(Model Context Protocol，MCP) 客户端配置文件中：

```json
{
    "mcpServers": {
        "prosearch": {
            "url": "按需替换上述SSE链接",
            "headers": {
                "TENCENTCLOUD_SECRET_KEY": "替换你的腾讯云密钥",
                "TENCENTCLOUD_SECRET_ID": "替换你的腾讯云密钥",
                "source": "tx_cloud_user"
            }
        }
    }
}
```

- TENCENTCLOUD_SECRET_KEY/TENCENTCLOUD_SECRET_ID：腾讯云的密钥，用于身份认证，请妥善保管，切勿泄露。获取方式:
    - 访问 [腾讯云密钥管理](https://console.cloud.tencent.com/cam/capi)。
    - 新建密钥并复制生成的 **SecretId** 和 **SecretKey**。
- source：标记用户来源，固定为：tx_cloud_user