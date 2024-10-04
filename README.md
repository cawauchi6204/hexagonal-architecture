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

## SQL スクリプト

以下は、上記の設計に基づいた SQL スクリプトです。

```sql
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE threads (
    id CHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    user_id CHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE posts (
    id CHAR(36) PRIMARY KEY,
    thread_id CHAR(36),
    user_id CHAR(36),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (thread_id) REFERENCES threads(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE comments (
    id CHAR(36) PRIMARY KEY,
    post_id CHAR(36),
    user_id CHAR(36),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE users_comments (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36),
    comment_id CHAR(36),
    is_like BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (comment_id) REFERENCES comments(id),
    UNIQUE KEY unique_user_comment (user_id, comment_id)
);

CREATE TABLE followers (
    follower_id CHAR(36),
    followed_id CHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, followed_id),
    FOREIGN KEY (follower_id) REFERENCES users(id),
    FOREIGN KEY (followed_id) REFERENCES users(id),
    UNIQUE KEY unique_follow (follower_id, followed_id)
);

CREATE TABLE tags (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_tags (
    user_id CHAR(36),
    tag_id CHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, tag_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    UNIQUE KEY unique_user_tag (user_id, tag_id)
);

CREATE TABLE thread_tags (
    thread_id CHAR(36),
    tag_id CHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (thread_id, tag_id),
    FOREIGN KEY (thread_id) REFERENCES threads(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id),
    UNIQUE KEY unique_thread_tag (thread_id, tag_id)
);
```

## テーブルの関係

- `users` テーブルは、ユーザーの情報を保持します。
- `threads` テーブルは、スレッドの情報を保持します。
- `posts` テーブルは、投稿の情報を保持します。
- `comments` テーブルは、コメントの情報を保持します。
- `users_comments` テーブルは、ユーザーのコメントに対するいいねの情報を保持します。
- `followers` テーブルは、ユーザーのフォロー情報を保持します。


## seeder
-- users テーブルのシードデータ
INSERT INTO users (id, username, email, password_hash, created_at) VALUES
('123e4567-e89b-12d3-a456-426614174000', 'user1', 'user1@example.com', 'hashed_password_1', CURRENT_TIMESTAMP),
('223e4567-e89b-12d3-a456-426614174001', 'user2', 'user2@example.com', 'hashed_password_2', CURRENT_TIMESTAMP),
('323e4567-e89b-12d3-a456-426614174002', 'user3', 'user3@example.com', 'hashed_password_3', CURRENT_TIMESTAMP);

-- threads テーブルのシードデータ
INSERT INTO threads (id, title, user_id, created_at) VALUES
('423e4567-e89b-12d3-a456-426614174003', '最初のスレッド', '123e4567-e89b-12d3-a456-426614174000', CURRENT_TIMESTAMP),
('523e4567-e89b-12d3-a456-426614174004', '2番目のスレッド', '223e4567-e89b-12d3-a456-426614174001', CURRENT_TIMESTAMP);

-- posts テーブルのシードデータ
INSERT INTO posts (id, thread_id, user_id, content, created_at) VALUES
('623e4567-e89b-12d3-a456-426614174005', '423e4567-e89b-12d3-a456-426614174003', '123e4567-e89b-12d3-a456-426614174000', '最初の投稿です', CURRENT_TIMESTAMP),
('723e4567-e89b-12d3-a456-426614174006', '423e4567-e89b-12d3-a456-426614174003', '223e4567-e89b-12d3-a456-426614174001', '2番目の投稿です', CURRENT_TIMESTAMP),
('823e4567-e89b-12d3-a456-426614174007', '523e4567-e89b-12d3-a456-426614174004', '323e4567-e89b-12d3-a456-426614174002', '別のスレッドの投稿です', CURRENT_TIMESTAMP);

-- comments テーブルのシードデータ
INSERT INTO comments (id, post_id, user_id, content, created_at) VALUES
('923e4567-e89b-12d3-a456-426614174008', '623e4567-e89b-12d3-a456-426614174005', '223e4567-e89b-12d3-a456-426614174001', '最初のコメントです', CURRENT_TIMESTAMP),
('a23e4567-e89b-12d3-a456-426614174009', '623e4567-e89b-12d3-a456-426614174005', '323e4567-e89b-12d3-a456-426614174002', '2番目のコメントです', CURRENT_TIMESTAMP);

-- users_comments テーブルのシードデータ
INSERT INTO users_comments (id, user_id, comment_id, is_like, created_at) VALUES
('b23e4567-e89b-12d3-a456-426614174010', '123e4567-e89b-12d3-a456-426614174000', '923e4567-e89b-12d3-a456-426614174008', TRUE, CURRENT_TIMESTAMP),
('c23e4567-e89b-12d3-a456-426614174011', '223e4567-e89b-12d3-a456-426614174001', 'a23e4567-e89b-12d3-a456-426614174009', TRUE, CURRENT_TIMESTAMP);

-- followers テーブルのシードデータ
INSERT INTO followers (follower_id, followed_id, created_at) VALUES
('123e4567-e89b-12d3-a456-426614174000', '223e4567-e89b-12d3-a456-426614174001', CURRENT_TIMESTAMP),
('223e4567-e89b-12d3-a456-426614174001', '323e4567-e89b-12d3-a456-426614174002', CURRENT_TIMESTAMP);

-- tags テーブルのシードデータ
INSERT INTO tags (id, name, created_at) VALUES
('d23e4567-e89b-12d3-a456-426614174012', 'プログラミング', CURRENT_TIMESTAMP),
('e23e4567-e89b-12d3-a456-426614174013', 'デザイン', CURRENT_TIMESTAMP);

-- users_tags テーブルのシードデータ
INSERT INTO users_tags (user_id, tag_id, created_at) VALUES
('123e4567-e89b-12d3-a456-426614174000', 'd23e4567-e89b-12d3-a456-426614174012', CURRENT_TIMESTAMP),
('223e4567-e89b-12d3-a456-426614174001', 'e23e4567-e89b-12d3-a456-426614174013', CURRENT_TIMESTAMP);

-- thread_tags テーブルのシードデータ
INSERT INTO thread_tags (thread_id, tag_id, created_at) VALUES
('423e4567-e89b-12d3-a456-426614174003', 'd23e4567-e89b-12d3-a456-426614174012', CURRENT_TIMESTAMP),
('523e4567-e89b-12d3-a456-426614174004', 'e23e4567-e89b-12d3-a456-426614174013', CURRENT_TIMESTAMP);