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
-- UUIDv4生成関数（既に存在する場合はスキップ）
DELIMITER $$
DROP FUNCTION IF EXISTS UUIDv4$$
CREATE FUNCTION UUIDv4() RETURNS CHAR(36)
BEGIN
    RETURN LOWER(CONCAT(
        HEX(RANDOM_BYTES(4)),
        '-', HEX(RANDOM_BYTES(2)),
        '-4', SUBSTR(HEX(RANDOM_BYTES(2)), 2, 3),
        '-', HEX(FLOOR(ASCII(RANDOM_BYTES(1)) / 64) + 8),
        SUBSTR(HEX(RANDOM_BYTES(2)), 2, 3),
        '-', HEX(RANDOM_BYTES(6))
    ));
END$$
DELIMITER ;

-- users テーブルのシードデータ
```
INSERT INTO users (id, username, email, password_hash, created_at)
SELECT 
    UUIDv4(),
    CONCAT(
        ELT(MOD(seq.n, 30) + 1, '山田', '佐藤', '鈴木', '田中', '高橋', '渡辺', '伊藤', '中村', '小林', '加藤',
            'John', 'Emma', 'Michael', 'Sophia', 'William', 'Olivia', 'James', 'Ava', 'Robert', 'Isabella',
            'Liu', 'Zhang', 'Wang', 'Li', 'Chen', 'Yang', 'Zhao', 'Wu', 'Zhou', 'Sun'),
        LPAD(seq.n, 3, '0')
    ),
    CONCAT(LOWER(CONCAT(
        ELT(MOD(seq.n, 30) + 1, '山田', '佐藤', '鈴木', '田中', '高橋', '渡辺', '伊藤', '中村', '小林', '加藤',
            'John', 'Emma', 'Michael', 'Sophia', 'William', 'Olivia', 'James', 'Ava', 'Robert', 'Isabella',
            'Liu', 'Zhang', 'Wang', 'Li', 'Chen', 'Yang', 'Zhao', 'Wu', 'Zhou', 'Sun'),
        LPAD(seq.n, 3, '0')
    )), '@example.com'),
    CONCAT('hashed_password_', seq.n),
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
         (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
    ORDER BY n LIMIT 100
) seq;

-- threads テーブルのシードデータ
INSERT INTO threads (id, title, user_id, created_at)
SELECT 
    UUIDv4(),
    CONCAT(
        ELT(MOD(seq.n, 20) + 1, '最新のAI技術について', '美味しいラーメン屋さん情報', '効果的な勉強法', '週末の旅行プラン', 'おすすめの映画2023',
            '健康的な食生活のコツ', 'プログラミング初心者の質問', '環境にやさしい生活習慣', '面白い海外ドラマ', '株式投資のヒント',
            '子育ての悩み相談', 'ガジェット最新情報', '料理レシピ交換', 'フィットネスのモチベーション維持', '語学学習のコツ',
            '写真撮影テクニック', '読書感想文シェア', 'ペットの飼い方アドバイス', '節約生活のヒント', 'ゲーム攻略情報'),
        ' ', LPAD(seq.n, 3, '0')
    ),
    (SELECT id FROM users ORDER BY RAND() LIMIT 1),
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
         (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
    ORDER BY n LIMIT 100
) seq;

-- posts テーブルのシードデータ
INSERT INTO posts (id, thread_id, user_id, content, created_at)
SELECT 
    UUIDv4(),
    (SELECT id FROM threads ORDER BY RAND() LIMIT 1),
    (SELECT id FROM users ORDER BY RAND() LIMIT 1),
    CONCAT(
        ELT(MOD(seq.n, 10) + 1, 'この話題について詳しく知りたいです。', '私も同じ経験がありました。', '面白い視点ですね。', 'もっと情報を共有してください。',
            'これは非常に役立つ情報です。', '別の角度から考えてみましょう。', 'この意見には賛成できません。', '素晴らしい投稿をありがとうございます。',
            'もう少し具体的に説明していただけますか？', 'この話題について、専門家の意見も聞いてみたいです。'),
        ' (Post ', seq.n, ')'
    ),
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
         (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
    ORDER BY n LIMIT 100
) seq;

-- comments テーブルのシードデータ
INSERT INTO comments (id, post_id, user_id, content, created_at)
SELECT 
    UUIDv4(),
    (SELECT id FROM posts ORDER BY RAND() LIMIT 1),
    (SELECT id FROM users ORDER BY RAND() LIMIT 1),
    CONCAT(
        ELT(MOD(seq.n, 10) + 1, '同感です！', 'なるほど、参考になります。', 'もう少し詳しく教えてください。', '私も似たような経験があります。',
            'この意見には賛成できません。', '面白い視点ですね。', 'ありがとうございます、助かりました。', 'これは考えさせられますね。',
            'もっと議論を深めたいです。', '新しい発見がありました。'),
        ' (Comment ', seq.n, ')'
    ),
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
         (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
    ORDER BY n LIMIT 100
) seq;

-- users_comments テーブルのシードデータ
INSERT INTO users_comments (id, user_id, comment_id, is_like, created_at)
SELECT 
    UUIDv4(),
    (SELECT id FROM users ORDER BY RAND() LIMIT 1),
    (SELECT id FROM comments ORDER BY RAND() LIMIT 1),
    ROUND(RAND()),
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
         (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
    ORDER BY n LIMIT 100
) seq
ON DUPLICATE KEY UPDATE is_like = VALUES(is_like);

-- followers テーブルのシードデータ
INSERT INTO followers (follower_id, followed_id, created_at)
SELECT 
    u1.id,
    u2.id,
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM 
    users u1
    CROSS JOIN users u2
WHERE 
    u1.id <> u2.id
ORDER BY RAND()
LIMIT 100
ON DUPLICATE KEY UPDATE created_at = VALUES(created_at);

-- tags テーブルのシードデータ
INSERT INTO tags (id, name, created_at)
SELECT 
    UUIDv4(),
    CONCAT(
        ELT(MOD(seq.n, 20) + 1, 'テクノロジー', '料理', '旅行', '健康', 'スポーツ', '音楽', '映画', '読書', 'ファッション', 'アート',
            'ビジネス', '教育', '環境', 'ペット', 'ゲーム', 'DIY', '写真', '科学', '歴史', '言語'),
        LPAD(seq.n, 3, '0')
    ),
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM (
    SELECT a.N + b.N * 10 + 1 as n
    FROM (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) a,
         (SELECT 0 AS N UNION SELECT 1 UNION SELECT 2 UNION SELECT 3 UNION SELECT 4 UNION SELECT 5 UNION SELECT 6 UNION SELECT 7 UNION SELECT 8 UNION SELECT 9) b
    ORDER BY n LIMIT 100
) seq;

-- users_tags テーブルのシードデータ
INSERT IGNORE INTO users_tags (user_id, tag_id, created_at)
SELECT DISTINCT
    u.id,
    t.id,
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM 
    users u
    CROSS JOIN tags t
ORDER BY RAND()
LIMIT 100;

-- thread_tags テーブルのシードデータ
INSERT IGNORE INTO thread_tags (thread_id, tag_id, created_at)
SELECT DISTINCT
    th.id,
    t.id,
    NOW() - INTERVAL FLOOR(RAND() * 365) DAY
FROM 
    threads th
    CROSS JOIN tags t
ORDER BY RAND()
LIMIT 100;
```
