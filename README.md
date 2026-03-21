# ID Generator Service

分布式唯一 ID 生成服务。基于 Twitter Snowflake 算法，通过 gRPC 对外提供高性能、全局唯一的 ID 发号能力。

## 技术栈
- **Go 1.25+** / **Viper** 配置管理 / **Zap** 结构化日志
- **gRPC** + **Protobuf** 高性能通信
- **Snowflake** 雪花算法（毫秒级时间戳 + 节点 ID + 序列号）
- **Docker** 容器化部署

## 目录结构
```text
├── cmd/idgen/main.go                  # 主入口：配置加载、gRPC 注册、优雅停机
├── internal/
│   ├── idgen/snowflake.go             # 雪花算法核心实现
│   └── platform/config/config.go      # Viper 配置结构体 + godotenv 加载
├── configs/config.yaml                # 非敏感配置（提交到 Git）
├── .env                               # 环境变量（不提交，见 .env.example）
└── docker-compose.yaml                # 容器编排
```

## 快速开始

### 1. 配置环境变量
```bash
cp .env.example .env
```

### 2. 本地运行
```bash
go mod tidy
go run cmd/idgen/main.go
```
服务默认监听 `:50059`。

### 3. Docker 部署
```bash
docker-compose up -d --build

# 查看日志
docker logs -f id-generator
```

## 配置说明

### config.yaml
| 字段 | 说明 | 默认值 |
|------|------|--------|
| `app_env` | 运行环境 | `development` |
| `server.port` | gRPC 监听端口 | `50059` |
| `server.mode` | 运行模式 | `debug` |
| `snowflake.node_id` | 雪花节点 ID（集群部署时每个实例必须不同） | `1` |

### .env
| 变量 | 说明 |
|------|------|
| `APP_ENV` | 运行环境，影响日志颜色 |

## gRPC 接口

```protobuf
service IDGenerator {
  rpc NextID(NextIDRequest) returns (NextIDResponse);
}
```

调用 `NextID` 即可获取一个全局唯一的 `int64` ID。
