SQS介绍
-----

SQS是一种完全托管的消息队列服务，可让您分离和扩展微服务、分布式系统和无服务器应用程序。SQS 消除了与管理和运营消息型中间件相关的复杂性和开销，并使开发人员能够专注于重要工作。借助 SQS，您可以在软件组件之间发送、存储和接收任何规模的消息，而不会丢失消息，并且无需其他服务即可保持可用。

SQS 提供两种消息队列类型。标准队列提供最高吞吐量、最大努力排序和至少一次传送。SQS FIFO 队列旨在确保按照消息的发送顺序对消息进行严格一次处理。

* 无限的队列和消息：在任何区域中创建无限数量的 Amazon SQS 队列，并且消息数量不受限制
* 负载大小：消息负载可包含最多 256 KB 的任意格式的文本。每 64 KB“区块”的已发布数据以 1 次请求计费。例如，1 次负载为 256 KB 的 API 调用将以 4 次请求计费。要发送大小超过 256KB 的消息，您可以使用[适用于 Java 的 Amazon SQS 扩展客户端库](https://github.com/awslabs/amazon-sqs-java-extended-client-lib)，它采用 Amazon S3 来存储消息负载。可使用 SQS 发送消息负载的引用。
* 批处理：批量发送、接收或删除消息，每批最多 10 条消息或 256 KB。批消息与单条消息的消耗量相同，这意味着对于使用批消息的客户而言，SQS 更具成本效益。
* [长轮询](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html)在尽可能快地接收新消息的时候减少外部轮询，从而实现最低成本。在您的队列为空时，长轮询请求会为下一条消息等待至多 20 秒。长轮询请求和普通请求的消耗量相同。
* 可将消息在队列中最多保留 14 天。
* 同时发送和读取消息。
* 消息锁定：收到消息后，在处理期间它会变为“锁定”状态。这可防止其他计算机同时处理该消息。如果消息处理失败，锁定会过期，而消息也重新变为可用。
* 队列共享：匿名或使用特定 AWS 账户安全共享 Amazon SQS 队列。队列共享还可按 IP 地址和一天中的时刻进行限制。
* 服务器端加密 (SSE)：使用由 AWS Key Management Service (AWS KMS) 托管的密钥来保护 Amazon SQS 队列中的消息内容。只要 Amazon SQS 收到消息，SSE 就会对其进行加密。这些消息以加密的形式进行存储，且仅当它们发送到授权的客户时，Amazon SQS 才会对其进行解密。
* 死信队列 (DLQ)：处理具有死信队列的用户未成功处理的消息。当超出消息的最大接收数时，消息将会移动到与原队列相关的 DLQ 中。为 DLQ 设置单独的用户进程，以便帮助分析和理解消息卡住的原因。DLQ 的类型必须与来源队列 (标准队列或 FIFO 队列) 相同。

SQS Ruby SDK API
----------------

```ruby
client_options = {
  access_key_id: 'AKIAJJPZ3JRQMFGXHOPA',
  secret_access_key: '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',
  region: 'us-east-1'
}
client = Aws::SQS::Client.new(options)

queue_options = {
  queue_name: "bim360_analytics_2019-12-17",
  attributes: {
    "DelaySeconds" => 0, # 延迟队列消息传递的时间
    "MaximumMessageSize" => '256kb',
    "MessageRetentionPeriod" => 4.day, # SQS retain message time
    "Policy": '', # queue's policy
    "ReceiveMessageWaitTimeSeconds" => '', # 等待消息到达的时间
    "RedrivePolicy" => {
      "deadLetterTargetArn" => '',
      "maxReceiveCount" => ''
    }, # 查看dead letter queue
    "VisibilityTimeout" => '', # 队列可见性超时
    "KmsMasterKeyId" => '', # 自定义管理密钥
    "KmsDataKeyReusePeriodSeconds" => '',
    "FifoQueue" => 'true/false', # 是否是FIFO QUEUE
    "ContentBasedDeduplication" => 'true/false', # 基于内容重复删除
  },
  tags: {
    'key' => 'value'
  } # 定义队列标签
}

# ContentBasedDeduplication: This option not work

queue_options = {
  queue_name: "bim360_analytics_2019-12-18",
  tags: {
    'Environment' => 'Local'
  }
}
queue_resp = client.create_queue(queue_options)
queue_resp.queue_url

client.get_queue_attributes({queue_url: queue_resp.queue_url, attribute_names: ["All"]})

resp = client.send_message({
  queue_url: "String",
  message_body: "String",
  delay_seconds: 1,
  message_attributes: {
    "String" => {
      string_value: "String",
      binary_value: "data",
      string_list_values: ["String"],
      binary_list_values: ["data"],
      data_type: "String", # required
    },
  },
  message_system_attributes: {
    "AWSTraceHeader" => {
      string_value: "String",
      binary_value: "data",
      string_list_values: ["String"],
      binary_list_values: ["data"],
      data_type: "String", # required
    },
  },
  message_deduplication_id: "String", # 只针对FIFO队列
  message_group_id: "String" # 只针对FIFO队列
})

messages = {
  queue_url: "#{queue_resp.queue_url}",
  entries: [
    {
      id: '1',
      message_body: '123456789',
    },
    {
      id: '2',
      message_body: '987654321',
    }
  ]
}

client.send_message_batch(messages)

resp = client.receive_message({
  queue_url: "String", 
  attribute_names: ["All"],
  message_attribute_names: ["MessageAttributeName"],
  max_number_of_messages: 1, # 返回最大消息条数, default: 1
  visibility_timeout: 1,
  wait_time_seconds: 1,
  receive_request_attempt_id: "String",
})
```

