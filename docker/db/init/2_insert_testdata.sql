DELETE FROM
  sketch_users;

DELETE FROM
  sketch_items;

DELETE FROM
  sketch_authors;

DELETE FROM
  sketch_tags;

DELETE FROM
  sketch_reviews;

INSERT INTO
  sketch_users (name, password)
VALUES
  (
    'user1',
    '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08'
  ),
(
    'user2',
    '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08'
  ),
(
    'user3',
    '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08'
  );

INSERT INTO
  sketch_authors (user_id, name, birth)
VALUES
  (1, '長女', 2018),
(1, '次女', 2022);

INSERT INTO
  sketch_tags (user_id, name)
VALUES
  (1, '保育園'),
(1, '家'),
(1, '家族'),
(1, 'キャラクター'),
(1, '動物'),
(1, 'クレヨン'),
(1, '工作');

INSERT INTO
  sketch_items (user_id, title, image, status, author_id, tags, date)
VALUES
  (
    1,
    'ライオンとくまの仁義なき戦い',
    'sketch_01.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),
(
    1,
    'だれだ？わからないけど、とりあえずとびま〜す',
    'sketch_02.jpg',
    'PUBLIC',
    1,
    ARRAY [2,4,6],
    '20231123'
  ),(
    1,
    '鉄棒やっとできたー',
    'sketch_03.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220624'
  ),(
    1,
    'よくわからないが、いろんな柄を集めてきたー',
    'sketch_04.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220724'
  ),(
    1,
    'なにかとピング色のキャラクターと虹',
    'sketch_05.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220424'
  ),(
    1,
    'ワニとぞうさんが逆さま',
    'sketch_06.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220114'
  ),(
    1,
    '太陽の下で、これ横向きで合ってる？',
    'sketch_07.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220524'
  ),(
    1,
    'なんとかポケモン、パパはわからないけど',
    'sketch_08.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20230124'
  ),(
    1,
    '星のカービィと虹',
    'sketch_09.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'これな〜に？餃子？',
    'sketch_10.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'むし三昧、カタツムリがかわいい',
    'sketch_11.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '太陽とひよこと蝶々にまた虹',
    'sketch_12.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '仲良しうさぎちゃんと蝶々',
    'sketch_13.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'ポケモン？クォリティ低めw',
    'sketch_14.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'イタリアの国旗らしいよ',
    'sketch_15.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'これまで一番クォリティ低いかも、何歳のときに描いたっけ',
    'sketch_16.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '理解不能のやつだね〜〜〜',
    'sketch_17.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'なにをモチーフしてんだろう、なんか怖い',
    'sketch_18.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '怖いキャラクターにメガネくんがビビってる',
    'sketch_19.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'さまざまな時計',
    'sketch_20.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'ピカチュウ、ダイエット中、ヒゲが。。。',
    'sketch_21.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'むしコラボレーション、たまにある低クォリティー',
    'sketch_22.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'カラスと蝶々とワンちゃんの散歩',
    'sketch_23.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'ピカチュウとなんちゃらポケモン、わからなさすぎ',
    'sketch_24.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '家族でお出かけ、風船を買ってもらった',
    'sketch_25.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'お出かけ日和、小さいときに描いたのはわかる',
    'sketch_26.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'だれかの誕生日ケーキ',
    'sketch_27.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '栗の貼り作品',
    'sketch_28.jpg',
    'PUBLIC',
    2,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'なかなかのアート、将来的に大物画家になれるかも',
    'sketch_29.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '顔、大きめ',
    'sketch_30.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '魔女のハロウィン',
    'sketch_31.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '親がカラスにさらわれる〜ゴッホより～、普通に～、ラッセンが好き～',
    'sketch_32.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '顔真っ青、MADE IN CHINAドラえもん？',
    'sketch_33.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '上手ではないが、なんかキュートな顔',
    'sketch_34.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'かき氷の貼り作品',
    'sketch_35.jpg',
    'PUBLIC',
    2,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '風船と風船',
    'sketch_36.jpg',
    'PUBLIC',
    2,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '大きな魚',
    'sketch_37.jpg',
    'PUBLIC',
    2,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'ぶどうの水彩画',
    'sketch_38.jpg',
    'PUBLIC',
    2,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    'キノコ狩り',
    'sketch_39.jpg',
    'PUBLIC',
    2,
    ARRAY [1,3,5],
    '20220124'
  ),(
    1,
    '魔女とかぼちゃの姉妹',
    'sketch_40.jpg',
    'PUBLIC',
    1,
    ARRAY [1,3,5],
    '20220124'
  );

SELECT
  *
FROM
  sketch_users;

SELECT
  *
FROM
  sketch_items;

SELECT
  *
FROM
  sketch_authors;

SELECT
  *
FROM
  sketch_tags;

SELECT
  *
FROM
  sketch_reviews;