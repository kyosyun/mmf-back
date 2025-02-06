-- Active: 1719815274611@@34.85.62.61@5432@mmf
CREATE TABLE murder_mystery_master (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,                         -- タイトル
    jan_code TEXT,                               -- JANコード
    asin_code TEXT,                              -- ASINコード
    release_date DATE NOT NULL,                  -- 発売日
    official_link TEXT,                          -- 公式リンク
    amazon_link TEXT,                            -- Amazonリンク

    platform TEXT NOT NULL,                      -- プラットフォーム（ボードゲーム, PCゲーム, アプリ, 書籍など）
    language TEXT NOT NULL,                      -- 言語（日本語, 英語など）
    surugaya_link TEXT,                          -- 駿河屋リンク
    steam_link TEXT,                             -- Steamリンク
    epicgames_link TEXT,                         -- EpicGamesリンク
    dlsite_link TEXT,                            -- DLsiteリンク
    rakuten_books_link TEXT,                     -- 楽天ブックスリンク
    bookwalker_link TEXT,                        -- BOOK☆WALKERリンク
    fanza_games_link TEXT,                       -- FANZA GAMESリンク
    
    price NUMERIC(10,2),                         -- 価格（最新価格）
    discount NUMERIC(10,2),                      -- 割引後価格（NULLなら割引なし）
    original_price NUMERIC(10,2),                -- 定価

    image_link TEXT,                             -- 画像リンク
    description TEXT,                            -- 概要

    players_min INTEGER,                         -- 最小プレイ人数
    players_max INTEGER,                         -- 最大プレイ人数
    play_time INTEGER,                           -- プレイ時間（分）
    requires_gm BOOLEAN DEFAULT FALSE,           -- GMの有無（TRUE: GM必須 / FALSE: GM不要）
    online_supported BOOLEAN DEFAULT FALSE,      -- オンライン対応（TRUE: オンライン可 / FALSE: オフラインのみ）

    genre TEXT,                                  -- ジャンル（ホラー, サスペンス, SF など）
    replayability TEXT CHECK (replayability IN ('リプレイ可', '1回限り')), -- リプレイ性
    age_recommendation TEXT,                      -- 推奨年齢（12歳以上, 18+ など）

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- 更新日時
);
