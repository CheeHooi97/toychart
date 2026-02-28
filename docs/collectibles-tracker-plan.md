# Collectibles Tracker Plan

## Goal

Build a product similar to WatchCharts, but for collectibles such as Pop Mart figures and Funko Pops.

Core user value:

- Discover collectible items and series
- See estimated market value over time
- Track owned items and portfolio value
- Compare retail price vs market price
- Follow demand for rare or secret variants

## Product Direction

### MVP

Focus on a narrow, reliable first version:

- Initial catalog coverage:
  - LEGO collectibles
  - Funko
  - Hasbro
  - Mattel
  - Bandai
  - Hot Toys
  - NECA
  - Pop Mart
  - Miniso collaboration collectibles
- Marketplace data source: eBay sold listings
- Catalog / retail reference sources to support later:
  - Miniso
  - Top Pop Culture & Toy Collectibles
  - Ozzie Collectables
- Core pages:
  - Series/collection directory
  - Item detail page
  - Price chart page
  - User collection page
  - Search

### V2

After the basic catalog and pricing model works:

- Wishlist and alerts
- Variant rarity and chase tracking
- Multi-source price aggregation
  - eBay
  - Mercari
  - StockX if relevant for some collectibles
- Portfolio analytics
  - total collection value
  - gain/loss
  - top movers
- Community features
  - public profiles
  - recent sales feed

## Key Domain Decisions

To make this work well for collectibles, the data model needs to separate:

1. Brand
2. Product line / franchise
3. Release collection / wave / blind box set
4. Individual collectible
5. Variant
6. Marketplace sale event
7. Aggregated market price snapshot

This matters because "Labubu", "Skullpanda", and "Funko Pop Marvel" are not the same kind of entity:

- `brand`: Pop Mart, Funko
- `franchise`: Labubu, Skullpanda, Marvel
- `collection`: Exciting Macaron, The Monsters Have a Seat, Wave 23
- `collectible`: one figure within a collection
- `variant`: regular, chase, secret, glow, flocked, metallic

It also matters because your data inputs are not the same kind of entity:

- `brand`: LEGO, Funko, Hasbro, Mattel, Bandai, Hot Toys, NECA, Pop Mart
- `retailer/source`: Miniso, Top Pop Culture & Toy Collectibles, Ozzie Collectables, eBay

The schema should not mix those together. Brands define the catalog. Sources define where data comes from.

## Initial Coverage Strategy

Use three buckets so expansion stays clean:

### 1. Core collectible brands

Seed these as brands:

- LEGO
- Funko
- Hasbro
- Mattel
- Bandai
- Hot Toys
- NECA
- Pop Mart

### 2. Retail / collaboration catalogs

Treat these as external sources first, with optional brand linkage later:

- Miniso
- Top Pop Culture & Toy Collectibles
- Ozzie Collectables

### 3. Resale marketplaces

Use these for market pricing:

- eBay first
- Mercari later
- Whatnot later if you want live collectible market coverage

## What Exists Today

Current backend tables already hint at the right direction:

- `set`
- `toy`
- `toy_price`
- `user_toy`

But they are too flat for WatchCharts-style tracking.

Current limitations:

- `toy.setName` is a string, not a foreign key
- `toy_price` only stores one price value and lacks source, currency, sample size, and pricing date
- `user_toy` only stores quantity and cannot represent condition, purchase price, acquisition date, or variant
- There is no normalized catalog structure for brand, franchise, collection, and variant
- There is no raw sold-listing storage for price verification and reprocessing

## Recommended Architecture

### 1. Catalog Layer

Stores the clean product data:

- brands
- franchises
- collections
- collectibles
- collectible_variants

### 2. Market Data Layer

Stores marketplace ingestion and pricing:

- marketplace_sources
- marketplace_sales_raw
- market_sales
- price_snapshots

### 3. User Portfolio Layer

Stores ownership and watchlists:

- users
- user_collection_items
- user_watchlists
- user_alerts

### 4. Ingestion / Ops Layer

Tracks scraping/import health:

- ingestion_jobs
- external_reference_map

## Proposed Database Structure

Assumption: PostgreSQL remains the main database.

### `brand`

Top-level maker or platform.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| slug | varchar(100) | unique |
| name | varchar(255) | `Pop Mart`, `Funko` |
| brand_type | varchar(50) | `manufacturer`, `designer_toy`, `retailer_house_brand`, `licensor` |
| country_code | varchar(8) | optional |
| official_url | text | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(slug)`

### `franchise`

Universe or character family under a brand.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| brand_id | fk -> brand.id | required |
| slug | varchar(150) | unique per brand |
| name | varchar(255) | `Labubu`, `Skullpanda`, `Marvel` |
| description | text | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(brand_id, slug)`
- index `(brand_id, name)`

### `collection`

A release set, wave, blind box series, or product drop.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| brand_id | fk -> brand.id | required |
| franchise_id | fk -> franchise.id | nullable for generic lines |
| slug | varchar(150) | unique per brand |
| name | varchar(255) | `Exciting Macaron` |
| collection_type | varchar(50) | `blind_box`, `wave`, `single_release`, `exclusive_drop`, `brick_set`, `figure_line` |
| release_date | date | optional |
| msrp_amount | numeric(12,2) | base retail price |
| currency_code | char(3) | default `USD` |
| cover_image_url | text | optional |
| status | varchar(30) | `upcoming`, `active`, `retired` |
| metadata_json | jsonb | extra source-specific attributes |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(brand_id, slug)`
- index `(franchise_id, release_date desc)`
- index `(status)`

### `collectible`

One distinct item a user can track.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| collection_id | fk -> collection.id | required |
| franchise_id | fk -> franchise.id | denormalized for filtering |
| sku | varchar(120) | optional internal/external product code |
| number_label | varchar(50) | useful for Funko numbering |
| name | varchar(255) | item name |
| character_name | varchar(255) | optional |
| edition_name | varchar(255) | optional |
| rarity_tier | varchar(30) | `common`, `rare`, `secret`, `chase` |
| is_secret | boolean | default false |
| estimated_odds | numeric(8,4) | optional blind-box odds |
| image_url | text | main image |
| status | varchar(30) | `active`, `retired`, `cancelled` |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- index `(collection_id, name)`
- index `(franchise_id, rarity_tier)`
- index `(number_label)`

### `collectible_variant`

Variant-specific tracking. Important for chase/secret/exclusive versions.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| collectible_id | fk -> collectible.id | required |
| variant_code | varchar(80) | internal stable identifier |
| name | varchar(255) | `Glow Chase`, `Secret`, `Metallic` |
| variant_type | varchar(50) | `regular`, `chase`, `secret`, `exclusive`, `colorway` |
| retailer_exclusive | varchar(255) | optional |
| msrp_amount | numeric(12,2) | optional variant retail price |
| currency_code | char(3) | default `USD` |
| image_url | text | optional |
| attributes_json | jsonb | condition-neutral attributes |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(collectible_id, variant_code)`
- index `(variant_type)`

### `marketplace_source`

Catalog of marketplaces, retailers, official stores, and ingestion channels.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| code | varchar(50) | unique, e.g. `ebay` |
| name | varchar(255) | |
| source_type | varchar(50) | `marketplace`, `retailer`, `official_catalog`, `official_store` |
| brand_id | fk -> brand.id | nullable; useful for official brand stores |
| region_code | varchar(20) | optional |
| base_url | text | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(code, region_code)`

### `marketplace_sale_raw`

Stores raw sold-listing payloads so pricing can be recomputed later.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| source_id | fk -> marketplace_source.id | required |
| external_sale_id | varchar(255) | source listing id |
| search_keyword | varchar(255) | query used to collect it |
| payload_json | jsonb | raw marketplace response |
| sold_at | timestamptz | source sold date |
| fetched_at | timestamptz | ingestion time |
| created_at | timestamptz | |

Indexes:

- unique `(source_id, external_sale_id)`
- index `(sold_at desc)`

### `market_sale`

Normalized sold transaction tied to a collectible or variant.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| source_id | fk -> marketplace_source.id | required |
| collectible_id | fk -> collectible.id | nullable until matched |
| collectible_variant_id | fk -> collectible_variant.id | nullable |
| raw_sale_id | fk -> marketplace_sale_raw.id | required |
| title | varchar(500) | normalized listing title |
| sale_url | text | optional |
| sale_type | varchar(30) | `auction`, `buy_now`, `best_offer` |
| quantity | integer | default 1 |
| price_amount | numeric(12,2) | hammer price only |
| shipping_amount | numeric(12,2) | optional |
| total_amount | numeric(12,2) | optional |
| currency_code | char(3) | required |
| condition_grade | varchar(50) | `new`, `used`, `damaged_box` |
| box_condition | varchar(50) | optional |
| seller_country_code | varchar(8) | optional |
| sold_at | timestamptz | required |
| match_confidence | numeric(5,4) | optional |
| is_outlier | boolean | default false |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- index `(collectible_id, sold_at desc)`
- index `(collectible_variant_id, sold_at desc)`
- index `(source_id, sold_at desc)`
- index `(is_outlier, sold_at desc)`

### `price_snapshot`

Aggregated market pricing for charts.

Recommended grain:

- one row per item per day per source
- allow either `collectible_id` or `collectible_variant_id`

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| source_id | fk -> marketplace_source.id | required |
| collectible_id | fk -> collectible.id | nullable if variant-only |
| collectible_variant_id | fk -> collectible_variant.id | nullable |
| snapshot_date | date | required |
| currency_code | char(3) | required |
| sample_size | integer | sold listings used |
| low_price | numeric(12,2) | optional |
| median_price | numeric(12,2) | key chart metric |
| avg_price | numeric(12,2) | optional |
| high_price | numeric(12,2) | optional |
| retail_price | numeric(12,2) | optional denormalized comparison |
| price_change_7d | numeric(12,2) | optional |
| price_change_30d | numeric(12,2) | optional |
| confidence_score | numeric(5,4) | based on sample quality |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(source_id, collectible_id, collectible_variant_id, snapshot_date, currency_code)`
- index `(collectible_id, snapshot_date desc)`
- index `(collectible_variant_id, snapshot_date desc)`

### `external_reference_map`

Maps internal records to external ids from Pop Mart, Funko, eBay, etc.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| entity_type | varchar(50) | `collection`, `collectible`, `variant` |
| entity_id | varchar(36) | internal id |
| source_id | fk -> marketplace_source.id | or external source bucket |
| external_id | varchar(255) | required |
| external_url | text | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(entity_type, entity_id, source_id)`
- unique `(source_id, external_id)`

### `user_collection_item`

Represents what a user owns.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| user_id | fk -> user.id | required |
| collectible_id | fk -> collectible.id | required |
| collectible_variant_id | fk -> collectible_variant.id | nullable |
| quantity | integer | default 1 |
| acquisition_price | numeric(12,2) | optional |
| acquisition_currency | char(3) | optional |
| acquired_at | date | optional |
| condition_grade | varchar(50) | `sealed`, `opened`, `mint` |
| box_condition | varchar(50) | optional |
| notes | text | optional |
| visibility | varchar(20) | `private`, `public` |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- index `(user_id, collectible_id)`
- index `(user_id, collectible_variant_id)`

### `user_watchlist`

Items a user wants to follow or buy.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| user_id | fk -> user.id | required |
| collectible_id | fk -> collectible.id | required |
| collectible_variant_id | fk -> collectible_variant.id | nullable |
| target_buy_price | numeric(12,2) | optional |
| currency_code | char(3) | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- unique `(user_id, collectible_id, collectible_variant_id)`

### `user_alert`

Alert rules for price movement.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| user_id | fk -> user.id | required |
| collectible_id | fk -> collectible.id | required |
| collectible_variant_id | fk -> collectible_variant.id | nullable |
| alert_type | varchar(50) | `price_below`, `price_above`, `new_sale` |
| threshold_amount | numeric(12,2) | optional |
| currency_code | char(3) | optional |
| is_active | boolean | default true |
| last_triggered_at | timestamptz | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- index `(user_id, is_active)`
- index `(collectible_id, is_active)`

### `ingestion_job`

Operational tracking for imports and scraping jobs.

| column | type | notes |
| --- | --- | --- |
| id | uuid / varchar(36) | primary key |
| source_id | fk -> marketplace_source.id | required |
| job_type | varchar(50) | `catalog_import`, `sales_scrape`, `price_rollup` |
| status | varchar(30) | `queued`, `running`, `success`, `failed` |
| started_at | timestamptz | optional |
| finished_at | timestamptz | optional |
| items_processed | integer | default 0 |
| error_message | text | optional |
| metadata_json | jsonb | optional |
| created_at | timestamptz | |
| updated_at | timestamptz | |

Indexes:

- index `(source_id, created_at desc)`
- index `(job_type, status, created_at desc)`

## Recommended Relationships

High-level relational flow:

- `brand -> franchise -> collection -> collectible -> collectible_variant`
- `collectible / collectible_variant -> market_sale -> price_snapshot`
- `user -> user_collection_item`
- `user -> user_watchlist`
- `user -> user_alert`

## Mapping From Current Tables

If you want to evolve the current schema instead of replacing it immediately:

- `set` -> migrate toward `collection`
- `toy` -> migrate toward `collectible`
- `toy_price` -> replace with `price_snapshot`
- `user_toy` -> replace with `user_collection_item`

Recommended near-term compatibility approach:

1. Keep current tables working for existing endpoints.
2. Add new normalized tables in parallel.
3. Build new read paths on the normalized schema.
4. Backfill old data into new tables.
5. Deprecate old tables after API migration.

## Suggested Seed Data

The first seed pass should include:

### Brands

- LEGO
- Funko
- Hasbro
- Mattel
- Bandai
- Hot Toys
- NECA
- Pop Mart

### Example franchises

- LEGO: Star Wars, Marvel, Harry Potter, Icons
- Funko: Pop!, Soda, Bitty Pop!
- Hasbro: Star Wars Black Series, Marvel Legends, Transformers, G.I. Joe
- Mattel: Barbie, Hot Wheels, Masters of the Universe
- Bandai: Tamashii Nations, S.H. Figuarts, Gundam
- Hot Toys: Marvel, Star Wars, DC
- NECA: TMNT, Alien, Predator, horror licenses
- Pop Mart: Labubu, Skullpanda, Dimoo, Molly

### Sources

- eBay as first resale source
- Miniso as optional retail / collaboration catalog source
- Top Pop Culture & Toy Collectibles as optional retailer source
- Ozzie Collectables as optional retailer source

## Suggested API Surface

For the app to feel like WatchCharts, the backend should expose:

- `GET /brands`
- `GET /collections`
- `GET /collections/:slug`
- `GET /collectibles/:id`
- `GET /collectibles/:id/price-history`
- `GET /collectibles/:id/sales`
- `POST /collectibles/search`
- `POST /user/collection`
- `GET /user/collection`
- `POST /user/watchlist`
- `POST /user/alerts`

## Delivery Plan

### Phase 1: Normalize the catalog

- Add `brand`, `franchise`, `collection`, `collectible`
- Seed Pop Mart and Funko
- Add slugs and images
- Keep only the item fields needed for search and detail pages

### Phase 2: Build market ingestion

- Save eBay sold listing raw payloads
- Normalize sold listings into `market_sale`
- Build matching rules from listing title to collectible / variant
- Add outlier filtering

### Phase 3: Price chart generation

- Create daily `price_snapshot` rollups
- Store median price and sample size
- Expose chart endpoint

### Phase 4: User portfolio

- Add `user_collection_item`
- Show collection market value
- Compare acquisition price vs estimated value

### Phase 5: Alerts and watchlists

- Add `user_watchlist`
- Add `user_alert`
- Notify users on thresholds

## Recommendation For This Repo

Given the current codebase, the most pragmatic next implementation step is:

1. Introduce new models for `brand`, `franchise`, `collection`, `collectible`, `market_sale`, and `price_snapshot`
2. Leave current `user` and auth tables as-is
3. Add a small importer for brand catalogs, starting with LEGO, Funko, and Pop Mart
4. Store retailer/source references for Miniso, Top Pop, and Ozzie separately from the brand catalog
5. Change eBay sold scraping to write raw and normalized sale rows
6. Add brand-specific matching rules where naming formats differ, especially LEGO set numbers and Funko numbering
7. Generate daily snapshots from normalized sales instead of writing one-off `toy_price` rows

## MVP Query Patterns To Design For

These are the queries that will dominate product performance:

- search collectibles by keyword, brand, franchise, collection
- get collectible detail by slug/id
- fetch 30d / 90d / 1y price history
- fetch recent sold listings for an item
- compute user portfolio market value

Important indexes:

- text/search indexes on collectible and collection names
- descending indexes on `market_sale.sold_at`
- unique daily index for `price_snapshot`
- compound user ownership indexes on `user_collection_item`

## Final Recommendation

Do not model this as only `set -> toy -> toy_price`.

That shape is enough for a demo, but not for a collectible market tracker. The durable version should treat catalog data, raw marketplace data, normalized sales, and user portfolio data as separate concerns.

If you want, the next step can be turning this plan into actual GORM models and migrations inside the existing Go backend.
