# データベース設計

## Users テーブル

| カラム名       | データ型     | 制約               | 説明               |
| -------------- | ------------ | ------------------ | ------------------ |
| id             | INT          | PRIMARY KEY        | ユーザーID         |
| username       | VARCHAR(50)  | UNIQUE, NOT NULL   | ユーザー名         |
| email          | VARCHAR(100) | UNIQUE, NOT NULL   | メールアドレス     |
| password_hash  | VARCHAR(255) | NOT NULL           | パスワードハッシュ |
| created_at     | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | 作成日時 |

## Threads テーブル

| カラム名       | データ型     | 制約               | 説明               |
| -------------- | ------------ | ------------------ | ------------------ |
| id             | INT          | PRIMARY KEY        | スレッドID         |
| title          | VARCHAR(255) | NOT NULL           | スレッドタイトル   |
| user_id        | INT          | FOREIGN KEY (Users.id) | スレッド作成者のユーザーID |
| created_at     | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | 作成日時 |

## Posts テーブル

| カラム名       | データ型     | 制約               | 説明               |
| -------------- | ------------ | ------------------ | ------------------ |
| id             | INT          | PRIMARY KEY        | 投稿ID             |
| thread_id      | INT          | FOREIGN KEY (Threads.id) | スレッドID         |
| user_id        | INT          | FOREIGN KEY (Users.id) | 投稿者のユーザーID |
| content        | TEXT         | NOT NULL           | 投稿内容           |
| created_at     | TIMESTAMP    | DEFAULT CURRENT_TIMESTAMP | 作成日時 |

## SQLスクリプト

以下は、上記の設計に基づいたSQLスクリプトです。

```sql
CREATE TABLE Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Threads (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE Posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    thread_id INT,
    user_id INT,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (thread_id) REFERENCES Threads(id),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);