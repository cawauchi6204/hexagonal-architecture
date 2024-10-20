1. make compose-up
2. make migrate-up

cloudflare pagesのビルド構成の設定
<img width="403" alt="image" src="https://github.com/user-attachments/assets/1e1ad3e4-0f21-4bff-8759-448889e75dca">


上記で環境立ち上げられます.
localhost:8888にアクセスするとブラウザでデータベースの中身を見ることができます

# データベース設計

## users テーブル

| カラム名      | データ型     | 制約                      | 説明               |
| ------------- | ------------ | ------------------------- | ------------------ |
| id            | CHAR(36)     | PRIMARY KEY               | ユーザー ID        |
| name          | VARCHAR(50)  | UNIQUE, NOT NULL          | ユーザー名         |
| email         | VARCHAR(100) | UNIQUE, NOT NULL          | メールアドレス     |
| password_hash | VARCHAR(255) | NOT NULL                  | パスワードハッシュ |
| created_at    | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | 作成日時           |

## threads テーブル

| カラム名   | データ型     | 制約                      | 説明                        |
| ---------- | ------------ | ------------------------- | --------------------------- |
| id         | CHAR(36)     | PRIMARY KEY               | スレッド ID                 |
| title      | VARCHAR(255) | NOT NULL                  | スレッドタイトル            |
| user_id    | CHAR(36)     | FOREIGN KEY (users.id)    | スレッド作成者のユーザー ID |
| created_at | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | 作成日時                    |

## posts テーブル

| カラム名   | データ型  | 制約                      | 説明                |
| ---------- | --------- | ------------------------- | ------------------- |
| id         | CHAR(36)  | PRIMARY KEY               | 投稿 ID             |
| thread_id  | CHAR(36)  | FOREIGN KEY (threads.id)  | スレッド ID         |
| user_id    | CHAR(36)  | FOREIGN KEY (users.id)    | 投稿者のユーザー ID |
| content    | TEXT      | NOT NULL                  | 投稿内容            |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | 作成日時            |

## comments テーブル

| カラム名   | データ型  | 制約                      | 説明                        |
| ---------- | --------- | ------------------------- | --------------------------- |
| id         | CHAR(36)  | PRIMARY KEY               | コメント ID                 |
| post_id    | CHAR(36)  | FOREIGN KEY (posts.id)    | 紐づく投稿の ID             |
| user_id    | CHAR(36)  | FOREIGN KEY (users.id)    | コメント投稿者のユーザー ID |
| content    | TEXT      | NOT NULL                  | コメント内容                |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | 作成日時                    |

## users_comments テーブル

| カラム名   | データ型  | 制約                      | 説明         |
| ---------- | --------- | ------------------------- | ------------ |
| user_id    | CHAR(36)  | FOREIGN KEY (users.id)    | ユーザー ID  |
| comment_id | CHAR(36)  | FOREIGN KEY (comments.id) | コメント ID  |
| is_like    | BOOLEAN   | NOT NULL                  | いいねフラグ |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | 作成日時     |

## followers テーブル

| カラム名    | データ型  | 制約                      | 説明                          |
| ----------- | --------- | ------------------------- | ----------------------------- |
| follower_id | CHAR(36)  | FOREIGN KEY (users.id)    | フォロワーのユーザー ID       |
| followed_id | CHAR(36)  | FOREIGN KEY (users.id)    | フォローされているユーザー ID |
| created_at  | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | 作成日時                      |

## users_tags テーブル

| カラム名   | データ型  | 制約                      | 説明        |
| ---------- | --------- | ------------------------- | ----------- |
| user_id    | CHAR(36)  | FOREIGN KEY (users.id)    | ユーザー ID |
| tag_id     | CHAR(36)  | FOREIGN KEY (tags.id)     | タグ ID     |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | 作成日時    |

## tags テーブル

| カラム名   | データ型    | 制約                      | 説明     |
| ---------- | ----------- | ------------------------- | -------- |
| id         | CHAR(36)    | PRIMARY KEY               | タグ ID  |
| name       | VARCHAR(50) | UNIQUE, NOT NULL          | タグ名   |
| created_at | TIMESTAMP   | DEFAULT CURRENT_TIMESTAMP | 作成日時 |

## thread_tags テーブル

| カラム名   | データ型  | 制約                      | 説明        |
| ---------- | --------- | ------------------------- | ----------- |
| thread_id  | CHAR(36)  | FOREIGN KEY (threads.id)  | スレッド ID |
| tag_id     | CHAR(36)  | FOREIGN KEY (tags.id)     | タグ ID     |
| created_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | 作成日時    |

truncateするやつ
```sql
-- 外部キー制約を一時的に無効にする
SET FOREIGN_KEY_CHECKS = 0;

-- テーブルをtruncateする（依存関係の逆順）
TRUNCATE TABLE thread_tags;
TRUNCATE TABLE users_tags;
TRUNCATE TABLE users_comments;
TRUNCATE TABLE followers;
TRUNCATE TABLE comments;
TRUNCATE TABLE posts;
TRUNCATE TABLE threads;
TRUNCATE TABLE tags;
TRUNCATE TABLE users;

-- 外部キー制約を再度有効にする
SET FOREIGN_KEY_CHECKS = 1;
```
