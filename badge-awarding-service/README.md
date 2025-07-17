### 通知ドメイン

```text
/project-root
├── cmd/                    # main関数やDI、エントリポイント
│   └── main.go
│
├── domain/                 # DDDの "ドメイン層"
│   ├── notification/       # 通知ドメイン（業務ロジックの単位）
│   │   ├── entity.go       # Entity: Notification, Subscriber
│   │   ├── value.go        # Value Object: EmailAddress, Message
│   │   ├── service.go      # ドメインサービス: NotificationSender
│   │   └── repository.go   # インターフェース: NotificationRepository
│   │
│   └── common/             # 共通値オブジェクトなど（任意）
│       └── time_range.go
│
├── usecase/                # アプリケーション層（ユースケース）
│   └── notification/
│       ├── send.go         # 通知を送るユースケース
│       └── subscribe.go    # 通知先の登録ユースケース
│
├── interface/              # インターフェースアダプタ層（外部との接続）
│   └── handler/            # APIハンドラ（REST/GraphQL/Lambdaなど）
│       └── notification_handler.go
│
├── infrastructure/         # 外部インフラ：DB、SNS、SQSなど
│   ├── sns/
│   │   └── sns_publisher.go   # SNS通知の実装
│   ├── dynamodb/
│   │   └── subscriber_repo.go # DynamoDBでSubscriberを保存
│   └── queue/
│       └── sqs_consumer.go    # SQSのConsumer実装
│
├── config/                 # 設定ファイル・環境変数の管理
│
└── internal/               # テスト用モックなど（必要であれば）

```