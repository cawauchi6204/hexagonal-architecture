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

```sql
-- users テーブルのシードデータ
INSERT IGNORE INTO users (id, username, email, password_hash, created_at)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'ネコ好き太郎', 'cat_lover@example.com', 'hash1', '2023-01-01 00:00:00'),
    ('22222222-2222-2222-2222-222222222222', '寿司職人見習い', 'sushi_apprentice@example.com', 'hash2', '2023-01-02 00:00:00'),
    ('33333333-3333-3333-3333-333333333333', 'カラオケの帝王', 'karaoke_king@example.com', 'hash3', '2023-01-03 00:00:00'),
    ('44444444-4444-4444-4444-444444444444', 'コーヒー中毒者', 'coffee_addict@example.com', 'hash4', '2023-01-04 00:00:00'),
    ('55555555-5555-5555-5555-555555555555', 'アニメオタク代表', 'anime_otaku@example.com', 'hash5', '2023-01-05 00:00:00'),
    ('66666666-6666-6666-6666-666666666666', '筋トレマニア', 'muscle_maniac@example.com', 'hash6', '2023-01-06 00:00:00'),
    ('77777777-7777-7777-7777-777777777777', 'ラーメンハンター', 'ramen_hunter@example.com', 'hash7', '2023-01-07 00:00:00'),
    ('88888888-8888-8888-8888-888888888888', 'ゲーム廃人予備軍', 'game_addict@example.com', 'hash8', '2023-01-08 00:00:00'),
    ('99999999-9999-9999-9999-999999999999', '猫舌アイス部', 'ice_cream_lover@example.com', 'hash9', '2023-01-09 00:00:00'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '早起き失敗太郎', 'sleepyhead@example.com', 'hash10', '2023-01-10 00:00:00');

-- threads テーブルのシードデータ
INSERT IGNORE INTO threads (id, title, user_id, created_at)
VALUES
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '猫の毛玉で地球を救う方法', '11111111-1111-1111-1111-111111111111', '2023-02-01 00:00:00'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', '寿司ネタで占う今日の運勢', '22222222-2222-2222-2222-222222222222', '2023-02-02 00:00:00'),
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'カラオケで世界平和を実現する歌', '33333333-3333-3333-3333-333333333333', '2023-02-03 00:00:00'),
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'コーヒーで走る車の開発プロジェクト', '44444444-4444-4444-4444-444444444444', '2023-02-04 00:00:00'),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', 'アニメキャラと結婚する方法', '55555555-5555-5555-5555-555555555555', '2023-02-05 00:00:00'),
    ('11111111-2222-3333-4444-555555555555', '筋トレで宇宙旅行を可能にする理論', '66666666-6666-6666-6666-666666666666', '2023-02-06 00:00:00'),
    ('22222222-3333-4444-5555-666666666666', 'ラーメンスープで入浴するメリット', '77777777-7777-7777-7777-777777777777', '2023-02-07 00:00:00'),
    ('33333333-4444-5555-6666-777777777777', 'ゲームの中で永遠に生きる方法', '88888888-8888-8888-8888-888888888888', '2023-02-08 00:00:00'),
    ('44444444-5555-6666-7777-888888888888', 'アイスクリームで家を建てる挑戦', '99999999-9999-9999-9999-999999999999', '2023-02-09 00:00:00'),
    ('55555555-6666-7777-8888-999999999999', '布団から出ずに出勤する裏技', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '2023-02-10 00:00:00');

-- posts テーブルのシードデータ
INSERT IGNORE INTO posts (id, thread_id, user_id, content, created_at)
VALUES
    ('66666666-7777-8888-9999-aaaaaaaaaaaa', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '11111111-1111-1111-1111-111111111111', '猫の毛玉で作った服は暖かいし、地球に優しい！', '2023-03-01 00:00:00'),
    ('77777777-8888-9999-aaaa-bbbbbbbbbbbb', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '22222222-2222-2222-2222-222222222222', 'マグロが出たら大吉、タマゴが出たら中吉...', '2023-03-02 00:00:00'),
    ('88888888-9999-aaaa-bbbb-cccccccccccc', 'dddddddd-dddd-dddd-dddd-dddddddddddd', '33333333-3333-3333-3333-333333333333', '「We Are The World」の100倍平和な歌を作ろう！', '2023-03-03 00:00:00'),
    ('99999999-aaaa-bbbb-cccc-dddddddddddd', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '44444444-4444-4444-4444-444444444444', 'カフェインパワーで東京からNYまで一気に走破！', '2023-03-04 00:00:00'),
    ('aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee', 'ffffffff-ffff-ffff-ffff-ffffffffffff', '55555555-5555-5555-5555-555555555555', '2次元と3次元の壁を越える恋愛テクニック', '2023-03-05 00:00:00'),
    ('bbbbbbbb-cccc-dddd-eeee-ffffffffffff', '11111111-2222-3333-4444-555555555555', '66666666-6666-6666-6666-666666666666', 'プロテインロケットで火星まで筋トレの旅', '2023-03-06 00:00:00'),
    ('cccccccc-dddd-eeee-ffff-111111111111', '22222222-3333-4444-5555-666666666666', '77777777-7777-7777-7777-777777777777', '豚骨スープで美肌になる秘訣とは？', '2023-03-07 00:00:00'),
    ('dddddddd-eeee-ffff-1111-222222222222', '33333333-4444-5555-6666-777777777777', '88888888-8888-8888-8888-888888888888', 'ゲームオーバーを回避し続ける究極のテクニック', '2023-03-08 00:00:00'),
    ('eeeeeeee-ffff-1111-2222-333333333333', '44444444-5555-6666-7777-888888888888', '99999999-9999-9999-9999-999999999999', '溶けないアイスクリームブロックで夢のお城を！', '2023-03-09 00:00:00'),
    ('ffffffff-1111-2222-3333-444444444444', '55555555-6666-7777-8888-999999999999', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'テレポーテーション出勤の実現に向けて', '2023-03-10 00:00:00');

-- comments テーブルのシードデータ
INSERT IGNORE INTO comments (id, post_id, user_id, content, created_at)
VALUES
    ('11111111-2222-3333-4444-555555555555', '66666666-7777-8888-9999-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', 'Comment content 1', '2023-04-01 00:00:00'),
    ('22222222-3333-4444-5555-666666666666', '77777777-8888-9999-aaaa-bbbbbbbbbbbb', '22222222-2222-2222-2222-222222222222', 'Comment content 2', '2023-04-02 00:00:00'),
    ('33333333-4444-5555-6666-777777777777', '88888888-9999-aaaa-bbbb-cccccccccccc', '33333333-3333-3333-3333-333333333333', 'Comment content 3', '2023-04-03 00:00:00'),
    ('44444444-5555-6666-7777-888888888888', '99999999-aaaa-bbbb-cccc-dddddddddddd', '44444444-4444-4444-4444-444444444444', 'Comment content 4', '2023-04-04 00:00:00'),
    ('55555555-6666-7777-8888-999999999999', 'aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee', '55555555-5555-5555-5555-555555555555', 'Comment content 5', '2023-04-05 00:00:00'),
    ('66666666-7777-8888-9999-aaaaaaaaaaaa', 'bbbbbbbb-cccc-dddd-eeee-ffffffffffff', '66666666-6666-6666-6666-666666666666', 'Comment content 6', '2023-04-06 00:00:00'),
    ('77777777-8888-9999-aaaa-bbbbbbbbbbbb', 'cccccccc-dddd-eeee-ffff-111111111111', '77777777-7777-7777-7777-777777777777', 'Comment content 7', '2023-04-07 00:00:00'),
    ('88888888-9999-aaaa-bbbb-cccccccccccc', 'dddddddd-eeee-ffff-1111-222222222222', '88888888-8888-8888-8888-888888888888', 'Comment content 8', '2023-04-08 00:00:00'),
    ('99999999-aaaa-bbbb-cccc-dddddddddddd', 'eeeeeeee-ffff-1111-2222-333333333333', '99999999-9999-9999-9999-999999999999', 'Comment content 9', '2023-04-09 00:00:00'),
    ('aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee', 'ffffffff-1111-2222-3333-444444444444', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Comment content 10', '2023-04-10 00:00:00');

-- users_comments テーブルのシードデータ
INSERT IGNORE INTO users_comments (id, user_id, comment_id, is_like, created_at)
VALUES
    ('bbbbbbbb-cccc-dddd-eeee-ffffffffffff', '11111111-1111-1111-1111-111111111111', '11111111-2222-3333-4444-555555555555', TRUE, '2023-05-01 00:00:00'),
    ('cccccccc-dddd-eeee-ffff-111111111111', '22222222-2222-2222-2222-222222222222', '22222222-3333-4444-5555-666666666666', FALSE, '2023-05-02 00:00:00'),
    ('dddddddd-eeee-ffff-1111-222222222222', '33333333-3333-3333-3333-333333333333', '33333333-4444-5555-6666-777777777777', TRUE, '2023-05-03 00:00:00'),
    ('eeeeeeee-ffff-1111-2222-333333333333', '44444444-4444-4444-4444-444444444444', '44444444-5555-6666-7777-888888888888', FALSE, '2023-05-04 00:00:00'),
    ('ffffffff-1111-2222-3333-444444444444', '55555555-5555-5555-5555-555555555555', '55555555-6666-7777-8888-999999999999', TRUE, '2023-05-05 00:00:00'),
    ('11111111-2222-3333-4444-555555555555', '66666666-6666-6666-6666-666666666666', '66666666-7777-8888-9999-aaaaaaaaaaaa', FALSE, '2023-05-06 00:00:00'),
    ('22222222-3333-4444-5555-666666666666', '77777777-7777-7777-7777-777777777777', '77777777-8888-9999-aaaa-bbbbbbbbbbbb', TRUE, '2023-05-07 00:00:00'),
    ('33333333-4444-5555-6666-777777777777', '88888888-8888-8888-8888-888888888888', '88888888-9999-aaaa-bbbb-cccccccccccc', FALSE, '2023-05-08 00:00:00'),
    ('44444444-5555-6666-7777-888888888888', '99999999-9999-9999-9999-999999999999', '99999999-aaaa-bbbb-cccc-dddddddddddd', TRUE, '2023-05-09 00:00:00'),
    ('55555555-6666-7777-8888-999999999999', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee', FALSE, '2023-05-10 00:00:00');

-- followers テーブルのシードデータ
INSERT IGNORE INTO followers (follower_id, followed_id, created_at)
VALUES
    ('11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', '2023-06-01 00:00:00'),
    ('22222222-2222-2222-2222-222222222222', '33333333-3333-3333-3333-333333333333', '2023-06-02 00:00:00'),
    ('33333333-3333-3333-3333-333333333333', '44444444-4444-4444-4444-444444444444', '2023-06-03 00:00:00'),
    ('44444444-4444-4444-4444-444444444444', '55555555-5555-5555-5555-555555555555', '2023-06-04 00:00:00'),
    ('55555555-5555-5555-5555-555555555555', '66666666-6666-6666-6666-666666666666', '2023-06-05 00:00:00'),
    ('66666666-6666-6666-6666-666666666666', '77777777-7777-7777-7777-777777777777', '2023-06-06 00:00:00'),
    ('77777777-7777-7777-7777-777777777777', '88888888-8888-8888-8888-888888888888', '2023-06-07 00:00:00'),
    ('88888888-8888-8888-8888-888888888888', '99999999-9999-9999-9999-999999999999', '2023-06-08 00:00:00'),
    ('99999999-9999-9999-9999-999999999999', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '2023-06-09 00:00:00'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '11111111-1111-1111-1111-111111111111', '2023-06-10 00:00:00');

-- tags テーブルのシードデータ
INSERT IGNORE INTO tags (id, name, created_at)
VALUES
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '猫毛活用術', '2023-07-01 00:00:00'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', '寿司占い', '2023-07-02 00:00:00'),
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'カラオケ外交', '2023-07-03 00:00:00'),
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'コーヒーパワー', '2023-07-04 00:00:00'),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', '2次元恋愛', '2023-07-05 00:00:00'),
    ('11111111-2222-3333-4444-555555555555', '筋肉宇宙飛行', '2023-07-06 00:00:00'),
    ('22222222-3333-4444-5555-666666666666', 'ラーメン美容', '2023-07-07 00:00:00'),
    ('33333333-4444-5555-6666-777777777777', 'ゲーム人生', '2023-07-08 00:00:00'),
    ('44444444-5555-6666-7777-888888888888', 'アイス建築', '2023-07-09 00:00:00'),
    ('55555555-6666-7777-8888-999999999999', '布団出勤', '2023-07-10 00:00:00');

-- users_tags テーブルのシードデータ
INSERT IGNORE INTO users_tags (user_id, tag_id, created_at)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2023-08-01 00:00:00'),
    ('22222222-2222-2222-2222-222222222222', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2023-08-02 00:00:00'),
    ('33333333-3333-3333-3333-333333333333', 'dddddddd-dddd-dddd-dddd-dddddddddddd', '2023-08-03 00:00:00'),
    ('44444444-4444-4444-4444-444444444444', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '2023-08-04 00:00:00'),
    ('55555555-5555-5555-5555-555555555555', 'ffffffff-ffff-ffff-ffff-ffffffffffff', '2023-08-05 00:00:00'),
    ('66666666-6666-6666-6666-666666666666', '11111111-2222-3333-4444-555555555555', '2023-08-06 00:00:00'),
    ('77777777-7777-7777-7777-777777777777', '22222222-3333-4444-5555-666666666666', '2023-08-07 00:00:00'),
    ('88888888-8888-8888-8888-888888888888', '33333333-4444-5555-6666-777777777777', '2023-08-08 00:00:00'),
    ('99999999-9999-9999-9999-999999999999', '44444444-5555-6666-7777-888888888888', '2023-08-09 00:00:00'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '55555555-6666-7777-8888-999999999999', '2023-08-10 00:00:00');

-- thread_tags テーブルのシードデータ
INSERT IGNORE INTO thread_tags (thread_id, tag_id, created_at)
VALUES
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2023-09-01 00:00:00'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'cccccccc-cccc-cccc-cccc-cccccccccccc', '2023-09-02 00:00:00'),
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'dddddddd-dddd-dddd-dddd-dddddddddddd', '2023-09-03 00:00:00'),
    ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '2023-09-04 00:00:00'),
    ('ffffffff-ffff-ffff-ffff-ffffffffffff', 'ffffffff-ffff-ffff-ffff-ffffffffffff', '2023-09-05 00:00:00'),
    ('11111111-2222-3333-4444-555555555555', '11111111-2222-3333-4444-555555555555', '2023-09-06 00:00:00'),
    ('22222222-3333-4444-5555-666666666666', '22222222-3333-4444-5555-666666666666', '2023-09-07 00:00:00'),
    ('33333333-4444-5555-6666-777777777777', '33333333-4444-5555-6666-777777777777', '2023-09-08 00:00:00'),
    ('44444444-5555-6666-7777-888888888888', '44444444-5555-6666-7777-888888888888', '2023-09-09 00:00:00'),
    ('55555555-6666-7777-8888-999999999999', '55555555-6666-7777-8888-999999999999', '2023-09-10 00:00:00');
```

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