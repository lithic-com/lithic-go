# Changelog

## 0.19.0 (2024-01-09)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/lithic-com/lithic-go/compare/v0.18.0...v0.19.0)

### Features

* **api:** add card renew endpoint ([#198](https://github.com/lithic-com/lithic-go/issues/198)) ([c37e540](https://github.com/lithic-com/lithic-go/commit/c37e540afcb9799aa1de648898c86a77cb75cb05))


### Chores

* add .keep files for examples and custom code directories ([#197](https://github.com/lithic-com/lithic-go/issues/197)) ([182d57a](https://github.com/lithic-com/lithic-go/commit/182d57a9996c02db2ed592615322f3b69e485cb6))
* **internal:** bump license ([#194](https://github.com/lithic-com/lithic-go/issues/194)) ([1267146](https://github.com/lithic-com/lithic-go/commit/1267146acf1efd9971287c036848e2d36d998ac5))
* **internal:** minor updates to pagination ([#195](https://github.com/lithic-com/lithic-go/issues/195)) ([80148a3](https://github.com/lithic-com/lithic-go/commit/80148a3aa0025ebda035e072cd4d93b55aa90242))


### Documentation

* **options:** fix link to readme ([#192](https://github.com/lithic-com/lithic-go/issues/192)) ([37f4590](https://github.com/lithic-com/lithic-go/commit/37f45903423d68297c5f415338527af5b4b6e2ec))


### Refactors

* remove excess whitespace ([#196](https://github.com/lithic-com/lithic-go/issues/196)) ([492ca26](https://github.com/lithic-com/lithic-go/commit/492ca26551b95231bc96900f87d5eb8ab2c92fee))

## 0.18.0 (2023-12-18)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/lithic-com/lithic-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** remove /auth_stream enrollment endpoints ([#191](https://github.com/lithic-com/lithic-go/issues/191)) ([72e530b](https://github.com/lithic-com/lithic-go/commit/72e530b17745df6fcc4b3dc52f12bce0aad0242b))


### Chores

* **ci:** run release workflow once per day ([#189](https://github.com/lithic-com/lithic-go/issues/189)) ([ee96175](https://github.com/lithic-com/lithic-go/commit/ee96175b780f4994810e4228eaa20234894b573b))

## 0.17.0 (2023-12-15)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/lithic-com/lithic-go/compare/v0.16.0...v0.17.0)

### Features

* **api:** rename `token` and `type` to `financial_account_token` and `financial_account_type` ([#188](https://github.com/lithic-com/lithic-go/issues/188)) ([e031830](https://github.com/lithic-com/lithic-go/commit/e031830fa7295d76f449d1becf4bd83dc5f76003))
* **internal:** fallback to json serialization if no serialization methods are defined ([#187](https://github.com/lithic-com/lithic-go/issues/187)) ([7a227de](https://github.com/lithic-com/lithic-go/commit/7a227de3c83785068395225f22ead2fb48e74d77))

## 0.16.0 (2023-12-05)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/lithic-com/lithic-go/compare/v0.15.0...v0.16.0)

### Features

* **api:** remove `CLOSED` account enum and update docstrings ([#184](https://github.com/lithic-com/lithic-go/issues/184)) ([60b3df2](https://github.com/lithic-com/lithic-go/commit/60b3df2099fac7f0932d1e62799c742483efa0c7))

## 0.15.0 (2023-11-28)

Full Changelog: [v0.14.1...v0.15.0](https://github.com/lithic-com/lithic-go/compare/v0.14.1...v0.15.0)

### Features

* **api:** add `get spend_limits` endpoints to `cards` and `accounts` ([#180](https://github.com/lithic-com/lithic-go/issues/180)) ([134e699](https://github.com/lithic-com/lithic-go/commit/134e699e8a32113ba3f894a153d3d37403b22c79))

## 0.14.1 (2023-11-17)

Full Changelog: [v0.14.0...v0.14.1](https://github.com/lithic-com/lithic-go/compare/v0.14.0...v0.14.1)

### Bug Fixes

* stop sending default idempotency headers with GET requests ([#179](https://github.com/lithic-com/lithic-go/issues/179)) ([6dc4ba1](https://github.com/lithic-com/lithic-go/commit/6dc4ba1eea0404c372f1892117a175816c354c5a))


### Chores

* **internal:** update stats file ([#177](https://github.com/lithic-com/lithic-go/issues/177)) ([8f78112](https://github.com/lithic-com/lithic-go/commit/8f7811238c22c2a97cf705dbfdc5e3ef16ea1dd1))

## 0.14.0 (2023-11-16)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/lithic-com/lithic-go/compare/v0.13.0...v0.14.0)

### Features

* **api:** updates ([#176](https://github.com/lithic-com/lithic-go/issues/176)) ([341e188](https://github.com/lithic-com/lithic-go/commit/341e1884699b2073ac93266a1066fee86f5a3df0))


### Refactors

* do not include `JSON` fields when serialising back to json ([#174](https://github.com/lithic-com/lithic-go/issues/174)) ([62660b3](https://github.com/lithic-com/lithic-go/commit/62660b3b27a5a2d09e5067ebef13fe26f8f5e3de))

## 0.13.0 (2023-11-09)

Full Changelog: [v0.12.1...v0.13.0](https://github.com/lithic-com/lithic-go/compare/v0.12.1...v0.13.0)

### Features

* **api:** updates ([#172](https://github.com/lithic-com/lithic-go/issues/172)) ([f5f964b](https://github.com/lithic-com/lithic-go/commit/f5f964be2caa9118feebe77dac4458ba7e1f41a9))

## 0.12.1 (2023-11-08)

Full Changelog: [v0.12.0...v0.12.1](https://github.com/lithic-com/lithic-go/compare/v0.12.0...v0.12.1)

### Bug Fixes

* make options.WithHeader utils case-insensitive ([#170](https://github.com/lithic-com/lithic-go/issues/170)) ([d22eb1f](https://github.com/lithic-com/lithic-go/commit/d22eb1f24c6e012183d97a68ab82f8e68b6336c2))

## 0.12.0 (2023-11-08)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/lithic-com/lithic-go/compare/v0.11.0...v0.12.0)

### Features

* **client:** allow binary returns ([#165](https://github.com/lithic-com/lithic-go/issues/165)) ([b291e7a](https://github.com/lithic-com/lithic-go/commit/b291e7a351c26ef3271be706e7befbb8b708b73b))


### Bug Fixes

* **api:** correct type for other fees details ([#169](https://github.com/lithic-com/lithic-go/issues/169)) ([9b24f3f](https://github.com/lithic-com/lithic-go/commit/9b24f3f67b5a0e254d010cf76998e732652be808))


### Documentation

* improve account holder control person documentation ([#167](https://github.com/lithic-com/lithic-go/issues/167)) ([99615cd](https://github.com/lithic-com/lithic-go/commit/99615cdec543967c9e2c28234d29f8dd1914afa9))
* **readme:** improve example snippets ([#168](https://github.com/lithic-com/lithic-go/issues/168)) ([4d6add2](https://github.com/lithic-com/lithic-go/commit/4d6add26c820292d762090ffe6b615711f2ca9b9))

## 0.11.0 (2023-11-02)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/lithic-com/lithic-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** add verification_attempts response property ([#164](https://github.com/lithic-com/lithic-go/issues/164)) ([d2ae13c](https://github.com/lithic-com/lithic-go/commit/d2ae13c10acee6397896751040da3f8e81fff059))
* **github:** include a devcontainer setup ([#163](https://github.com/lithic-com/lithic-go/issues/163)) ([d4f0c8a](https://github.com/lithic-com/lithic-go/commit/d4f0c8a08fdf99569c327adfda9eb92e2837938e))


### Chores

* **internal:** update gitignore ([#161](https://github.com/lithic-com/lithic-go/issues/161)) ([74c6247](https://github.com/lithic-com/lithic-go/commit/74c6247a00f12b4b6559ffa7ea09c55d65456a23))

## 0.10.0 (2023-10-26)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/lithic-com/lithic-go/compare/v0.9.0...v0.10.0)

### Features

* **api:** add CardProgram and DigitalCardArt resources ([#159](https://github.com/lithic-com/lithic-go/issues/159)) ([b7fd9fd](https://github.com/lithic-com/lithic-go/commit/b7fd9fd605c93c525312835ed6e088c123ab5440))

## 0.9.0 (2023-10-24)

Full Changelog: [v0.8.3...v0.9.0](https://github.com/lithic-com/lithic-go/compare/v0.8.3...v0.9.0)

### Features

* **api:** add AUTH_STREAM_ACCESS to responder endpoints ([#154](https://github.com/lithic-com/lithic-go/issues/154)) ([479ed5c](https://github.com/lithic-com/lithic-go/commit/479ed5cbbb145f12dd99fb2eaa867467598083f9))
* **api:** add verification_failed_reason ([#153](https://github.com/lithic-com/lithic-go/issues/153)) ([21318a2](https://github.com/lithic-com/lithic-go/commit/21318a2f4112fbfc7985ff88ba0ba424c64b2950))
* **api:** updates ([#152](https://github.com/lithic-com/lithic-go/issues/152)) ([f04846c](https://github.com/lithic-com/lithic-go/commit/f04846c1b17dc8a37d9febb64a7b2a7fd68c149e))
* **client:** adjust retry behavior ([#155](https://github.com/lithic-com/lithic-go/issues/155)) ([da9d8df](https://github.com/lithic-com/lithic-go/commit/da9d8df3aa6e44b0c1dd0fb28acb3f13ce7d99b9))


### Chores

* **internal:** rearrange client arguments ([#147](https://github.com/lithic-com/lithic-go/issues/147)) ([1d6367b](https://github.com/lithic-com/lithic-go/commit/1d6367be119cc184e3bcbe17d946a940bd9bf1ae))
* **internal:** reorder code ([#150](https://github.com/lithic-com/lithic-go/issues/150)) ([19d089b](https://github.com/lithic-com/lithic-go/commit/19d089b19af5cbe0407eef8957982470af0f2673))
* update README ([#145](https://github.com/lithic-com/lithic-go/issues/145)) ([0db60d6](https://github.com/lithic-com/lithic-go/commit/0db60d6f126e8178637444ac7a0afe4510624963))


### Documentation

* **api.md:** improve formatting of webhook helpers ([#149](https://github.com/lithic-com/lithic-go/issues/149)) ([1181762](https://github.com/lithic-com/lithic-go/commit/11817620261cbb0a8797665890254f94202cdccb))
* organisation -&gt; organization (UK to US English) ([#151](https://github.com/lithic-com/lithic-go/issues/151)) ([3741c37](https://github.com/lithic-com/lithic-go/commit/3741c37d66794baa44cfcb348df185f6724ff943))

## 0.8.3 (2023-10-11)

Full Changelog: [v0.8.2...v0.8.3](https://github.com/lithic-com/lithic-go/compare/v0.8.2...v0.8.3)

## 0.8.2 (2023-10-05)

Full Changelog: [v0.8.1...v0.8.2](https://github.com/lithic-com/lithic-go/compare/v0.8.1...v0.8.2)

### Features

* **api:** updates ([#141](https://github.com/lithic-com/lithic-go/issues/141)) ([8d31db7](https://github.com/lithic-com/lithic-go/commit/8d31db726756dc1e3450a3ef4ab0f9a85ec74e93))

## 0.8.1 (2023-10-03)

Full Changelog: [v0.8.0...v0.8.1](https://github.com/lithic-com/lithic-go/compare/v0.8.0...v0.8.1)

### Bug Fixes

* prevent index out of range bug during auto-pagination ([#139](https://github.com/lithic-com/lithic-go/issues/139)) ([84ad5e7](https://github.com/lithic-com/lithic-go/commit/84ad5e7025ca45e1b76354930f5f723e06f9a610))


### Chores

* **ci:** remove reviewer ([#137](https://github.com/lithic-com/lithic-go/issues/137)) ([3aa48ba](https://github.com/lithic-com/lithic-go/commit/3aa48bad524652597b418d4dcbbad4373ecf712e))
* **tests:** update test examples ([#140](https://github.com/lithic-com/lithic-go/issues/140)) ([805a28b](https://github.com/lithic-com/lithic-go/commit/805a28b29355776ce1e6a404133e1bda5ea0163d))

## 0.8.0 (2023-09-29)

Full Changelog: [v0.7.4...v0.8.0](https://github.com/lithic-com/lithic-go/compare/v0.7.4...v0.8.0)

### ⚠ BREAKING CHANGES

* **api:** remove `post /webhooks/account_holders` endpoint ([#133](https://github.com/lithic-com/lithic-go/issues/133))

### Refactors

* **api:** remove `post /webhooks/account_holders` endpoint ([#133](https://github.com/lithic-com/lithic-go/issues/133)) ([039145c](https://github.com/lithic-com/lithic-go/commit/039145c76c5e7b834df276e3fadf9f0e9d310c21))

## 0.7.4 (2023-09-25)

Full Changelog: [v0.7.3...v0.7.4](https://github.com/lithic-com/lithic-go/compare/v0.7.3...v0.7.4)

### Features

* **api:** add simulation endpoints, event types, fix transfer request AuthRule ([#129](https://github.com/lithic-com/lithic-go/issues/129)) ([424b35d](https://github.com/lithic-com/lithic-go/commit/424b35d6fe03f5fa9b60cbd811b163256916c815))
* improve retry behavior on context deadline ([#132](https://github.com/lithic-com/lithic-go/issues/132)) ([099165c](https://github.com/lithic-com/lithic-go/commit/099165c14901bb74d07445a1b98f157404041acc))


### Documentation

* **api.md:** rename Top Level to client name ([#130](https://github.com/lithic-com/lithic-go/issues/130)) ([68a64bd](https://github.com/lithic-com/lithic-go/commit/68a64bdf4d35f0ea6f5a63ca01a1419b24862c60))

## 0.7.3 (2023-09-15)

Full Changelog: [v0.7.2...v0.7.3](https://github.com/lithic-com/lithic-go/compare/v0.7.2...v0.7.3)

### Features

* retry on 408 Request Timeout ([#125](https://github.com/lithic-com/lithic-go/issues/125)) ([09f2dc7](https://github.com/lithic-com/lithic-go/commit/09f2dc765bd3e090f5efca803a34f56467977c11))


### Bug Fixes

* **core:** improve retry behavior and related docs ([#126](https://github.com/lithic-com/lithic-go/issues/126)) ([5b18ec6](https://github.com/lithic-com/lithic-go/commit/5b18ec69dbf9ef128961aa03abef986ca42ecaeb))

## 0.7.2 (2023-09-12)

Full Changelog: [v0.7.1...v0.7.2](https://github.com/lithic-com/lithic-go/compare/v0.7.1...v0.7.2)

### Bug Fixes

* **core:** add null check to prevent segfault when canceling context ([#120](https://github.com/lithic-com/lithic-go/issues/120)) ([91276bc](https://github.com/lithic-com/lithic-go/commit/91276bcc574b4e5c2ffe1c3501142974b7b95dde))


### Chores

* **internal:** improve reliability of cancel delay test ([#122](https://github.com/lithic-com/lithic-go/issues/122)) ([69cbd7c](https://github.com/lithic-com/lithic-go/commit/69cbd7cb296670f4faf85da170e709ece9e2d4ad))

## 0.7.1 (2023-09-11)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/lithic-com/lithic-go/compare/v0.7.0...v0.7.1)

### Features

* **api:** add Simulate Return Payment endpoint ([#119](https://github.com/lithic-com/lithic-go/issues/119)) ([5e24144](https://github.com/lithic-com/lithic-go/commit/5e24144ad980f4f15e2577607fdeebe392ed1518))
* **api:** add tokenizations.simulate and correct typo'd enum  ([#117](https://github.com/lithic-com/lithic-go/issues/117)) ([773c215](https://github.com/lithic-com/lithic-go/commit/773c2157cc66177fa08c175eac9bd85c50a767ec))
* **api:** add user defined id ([#114](https://github.com/lithic-com/lithic-go/issues/114)) ([18b2fd6](https://github.com/lithic-com/lithic-go/commit/18b2fd60a1af8ee9d21e311dae2a6a1b731f8363))
* fixes tests where an array has to have unique enum values ([#115](https://github.com/lithic-com/lithic-go/issues/115)) ([2d9beb3](https://github.com/lithic-com/lithic-go/commit/2d9beb3d3ba5625648f559d8fa19ddabbe4c3841))


### Chores

* **ci:** setup workflows to create releases and release PRs ([#109](https://github.com/lithic-com/lithic-go/issues/109)) ([5a5ac6d](https://github.com/lithic-com/lithic-go/commit/5a5ac6d39ce1df0d73b271152fa163802f83f3b3))
* **internal:** cleanup test params ([#118](https://github.com/lithic-com/lithic-go/issues/118)) ([cfb99ae](https://github.com/lithic-com/lithic-go/commit/cfb99ae1aa764465194d53909073bbc894c04e9c))
* **internal:** implement inline json unmarshalling ([#113](https://github.com/lithic-com/lithic-go/issues/113)) ([2fd86e1](https://github.com/lithic-com/lithic-go/commit/2fd86e1f564eb52417d30089c963b79fb959cf75))


### Documentation

* **readme:** add link to api.md ([#116](https://github.com/lithic-com/lithic-go/issues/116)) ([1aebf6b](https://github.com/lithic-com/lithic-go/commit/1aebf6bac01228564cbc8e84b9b216bc92c2d46e))

## [0.7.0](https://github.com/lithic-com/lithic-go/compare/v0.6.8...v0.7.0) (2023-08-15)


### ⚠ BREAKING CHANGES

* **api:** change `key` to `secret` ([#102](https://github.com/lithic-com/lithic-go/issues/102))

### Features

* **api:** change `key` to `secret` ([#102](https://github.com/lithic-com/lithic-go/issues/102)) ([66a2271](https://github.com/lithic-com/lithic-go/commit/66a227152f0bb1db938c256db1764114137ae550))


### Chores

* assign default reviewers to release PRs ([#104](https://github.com/lithic-com/lithic-go/issues/104)) ([6f95aaa](https://github.com/lithic-com/lithic-go/commit/6f95aaaca4aed01adeb217b31ad80ce4454e2028))
* **client:** send Idempotency-Key header ([#105](https://github.com/lithic-com/lithic-go/issues/105)) ([68b7e7f](https://github.com/lithic-com/lithic-go/commit/68b7e7fd330dbfac504398250ee9a4fce7976a17))

## [0.6.8](https://github.com/lithic-com/lithic-go/compare/v0.6.7...v0.6.8) (2023-08-11)


### Features

* allOf models now have toXxx methods to access the separate allOf models ([#100](https://github.com/lithic-com/lithic-go/issues/100)) ([3018520](https://github.com/lithic-com/lithic-go/commit/3018520eb3c3d720cb15cb905de6c3e1e31241e9))
* **api:** add card reissue shipping options ([#99](https://github.com/lithic-com/lithic-go/issues/99)) ([8aaa9d3](https://github.com/lithic-com/lithic-go/commit/8aaa9d34312ff43d6e96630b4584f93c67e2af63))


### Bug Fixes

* **client:** correctly set multipart form data boundary ([#97](https://github.com/lithic-com/lithic-go/issues/97)) ([5785e31](https://github.com/lithic-com/lithic-go/commit/5785e31de9b183ed55763f4b0d790e94d502ce4e))

## [0.6.7](https://github.com/lithic-com/lithic-go/compare/v0.6.6...v0.6.7) (2023-08-08)


### Features

* **api:** add carrier property to card create and reissue params ([#95](https://github.com/lithic-com/lithic-go/issues/95)) ([89f43dc](https://github.com/lithic-com/lithic-go/commit/89f43dcb4bae8fa556fe378742631555b7c8c71a))


### Documentation

* **readme:** remove beta status + document versioning policy ([#93](https://github.com/lithic-com/lithic-go/issues/93)) ([c622205](https://github.com/lithic-com/lithic-go/commit/c622205f8bbdab23116ef76608460bd5bcb5101e))

## [0.6.6](https://github.com/lithic-com/lithic-go/compare/v0.6.5...v0.6.6) (2023-08-01)


### Features

* **api:** updates ([#90](https://github.com/lithic-com/lithic-go/issues/90)) ([7411302](https://github.com/lithic-com/lithic-go/commit/741130245a9146e443a76c9d5e9ce5aa1a885ef1))


### Bug Fixes

* adjust typo of 'descisioning' to 'decisioning' ([#91](https://github.com/lithic-com/lithic-go/issues/91)) ([a8a36a7](https://github.com/lithic-com/lithic-go/commit/a8a36a7df2f8955f819bcabd75a4911975536622))


### Chores

* **internal:** minor reformatting of code ([#88](https://github.com/lithic-com/lithic-go/issues/88)) ([f499c61](https://github.com/lithic-com/lithic-go/commit/f499c61c4a11c274a31bc86aa0abb5446764171c))

## [0.6.5](https://github.com/lithic-com/lithic-go/compare/v0.6.4...v0.6.5) (2023-07-27)


### Features

* add `Bool` param field helper ([#86](https://github.com/lithic-com/lithic-go/issues/86)) ([847e53d](https://github.com/lithic-com/lithic-go/commit/847e53d771cbe02110590b0c4bf609d5212945d0))
* **api:** add payment and external bank accounts resource ([#84](https://github.com/lithic-com/lithic-go/issues/84)) ([98a36d5](https://github.com/lithic-com/lithic-go/commit/98a36d508e45f08767d480a53a6f5440f602a9f7))

## [0.6.4](https://github.com/lithic-com/lithic-go/compare/v0.6.3...v0.6.4) (2023-07-21)


### Features

* **api:** add `with_content` param ([#81](https://github.com/lithic-com/lithic-go/issues/81)) ([004cf68](https://github.com/lithic-com/lithic-go/commit/004cf68be7b1368a8288e72a41556a57fe42614f))

## [0.6.3](https://github.com/lithic-com/lithic-go/compare/v0.6.2...v0.6.3) (2023-07-18)


### Features

* **api:** add event message attempts ([#78](https://github.com/lithic-com/lithic-go/issues/78)) ([2e70c1b](https://github.com/lithic-com/lithic-go/commit/2e70c1b62a32fabe67adbb40c8d8b956e0f522bd))

## [0.6.2](https://github.com/lithic-com/lithic-go/compare/v0.6.1...v0.6.2) (2023-07-17)


### Features

* **api:** add more enum members to event types ([#73](https://github.com/lithic-com/lithic-go/issues/73)) ([7154eaa](https://github.com/lithic-com/lithic-go/commit/7154eaa9483bdf5735193b1d53651e2474bd30db))
* **api:** no longer require `website_url` property on KYB object ([#76](https://github.com/lithic-com/lithic-go/issues/76)) ([929bb45](https://github.com/lithic-com/lithic-go/commit/929bb45ec6dc98c63b8f09de59688f40f19cee7f))


### Chores

* **internal:** add `codegen.log` to `.gitignore` ([#75](https://github.com/lithic-com/lithic-go/issues/75)) ([21eb800](https://github.com/lithic-com/lithic-go/commit/21eb800d9c7e63b14a43c6111269a49c6a17aa2d))

## [0.6.1](https://github.com/lithic-com/lithic-go/compare/v0.6.0...v0.6.1) (2023-07-12)


### Features

* **api:** add digital wallet tokenization result event type ([#70](https://github.com/lithic-com/lithic-go/issues/70)) ([37a2a2d](https://github.com/lithic-com/lithic-go/commit/37a2a2d76b3ce7e3f44c3a102fb6786eead73f37))

## [0.6.0](https://github.com/lithic-com/lithic-go/compare/v0.5.0...v0.6.0) (2023-07-07)


### ⚠ BREAKING CHANGES

* **api:** remove previous_auth_rule_tokens from auth rules ([#64](https://github.com/lithic-com/lithic-go/issues/64))

### Features

* **api:** add `state` query param for cards ([#67](https://github.com/lithic-com/lithic-go/issues/67)) ([6175a18](https://github.com/lithic-com/lithic-go/commit/6175a183cbbfead69e15623069fe9e7b68cc013b))


### Refactors

* **api:** remove previous_auth_rule_tokens from auth rules ([#64](https://github.com/lithic-com/lithic-go/issues/64)) ([4a81426](https://github.com/lithic-com/lithic-go/commit/4a814263872defd55597162dcc7500d437f2b27b))


### Chores

* **internal:** rename some parameter names ([#69](https://github.com/lithic-com/lithic-go/issues/69)) ([14fb9df](https://github.com/lithic-com/lithic-go/commit/14fb9dfc662fff39db9fb78985839baed70c8eb4))

## [0.5.0](https://github.com/lithic-com/lithic-go/compare/v0.4.1...v0.5.0) (2023-06-29)


### ⚠ BREAKING CHANGES

* **types:** singularize array item types ([#63](https://github.com/lithic-com/lithic-go/issues/63))
* rename some response types and remove unnecessary types from paginated endpoints ([#60](https://github.com/lithic-com/lithic-go/issues/60))

### Features

* generate `api.md` file ([#57](https://github.com/lithic-com/lithic-go/issues/57)) ([308263b](https://github.com/lithic-com/lithic-go/commit/308263b05e16a3643f6cfb7d69d97cfd2369337c))


### Chores

* **tests:** minor reformatting of docs and tests ([#52](https://github.com/lithic-com/lithic-go/issues/52)) ([4866ab5](https://github.com/lithic-com/lithic-go/commit/4866ab5ea5f8d22e96e6b205af481b1478830f1d))


### Documentation

* add comments to alias types ([#58](https://github.com/lithic-com/lithic-go/issues/58)) ([da81689](https://github.com/lithic-com/lithic-go/commit/da8168969e06c6d7a1244f5c61826cde5aa84aad))
* add trailing newlines ([#61](https://github.com/lithic-com/lithic-go/issues/61)) ([0008e35](https://github.com/lithic-com/lithic-go/commit/0008e359774415942a67283ebb34e45d2cccef2e))
* **api:** update account limits docstrings ([#55](https://github.com/lithic-com/lithic-go/issues/55)) ([d9efdfc](https://github.com/lithic-com/lithic-go/commit/d9efdfca37f8472cf0dee9edef229dd0bbf54a22))
* **api:** update limits docstrings ([#59](https://github.com/lithic-com/lithic-go/issues/59)) ([4cd42ba](https://github.com/lithic-com/lithic-go/commit/4cd42baab7f9b7240324347b5393baaefcd93b92))


### Styles

* minor reordering of types and properties ([#62](https://github.com/lithic-com/lithic-go/issues/62)) ([89c502e](https://github.com/lithic-com/lithic-go/commit/89c502ec83740ab78f59bc39d2ab2a7dee110ff6))


### Refactors

* rename some response types and remove unnecessary types from paginated endpoints ([#60](https://github.com/lithic-com/lithic-go/issues/60)) ([7ada707](https://github.com/lithic-com/lithic-go/commit/7ada707cb5eb0365f7ae71ca3cefd47da9d85bcf))
* **types:** singularize array item types ([#63](https://github.com/lithic-com/lithic-go/issues/63)) ([ac26936](https://github.com/lithic-com/lithic-go/commit/ac269367dee47d42267f26d10497deb4499f9d1c))

## [0.4.1](https://github.com/lithic-com/lithic-go/compare/v0.4.0...v0.4.1) (2023-06-19)


### Documentation

* **api:** clarify dispute evidence filename docstring ([#50](https://github.com/lithic-com/lithic-go/issues/50)) ([bee16d8](https://github.com/lithic-com/lithic-go/commit/bee16d8d297776b78d4cac21e002702953976af8))

## [0.4.0](https://github.com/lithic-com/lithic-go/compare/v0.3.2...v0.4.0) (2023-06-15)


### ⚠ BREAKING CHANGES

* **api:** add dispute evidence filename ([#46](https://github.com/lithic-com/lithic-go/issues/46))

### Features

* **api:** add dispute evidence filename ([#46](https://github.com/lithic-com/lithic-go/issues/46)) ([78f70af](https://github.com/lithic-com/lithic-go/commit/78f70afe9e7d54d5c1640787109fca90b7afbf5c))
* respect `x-should-retry` header ([#45](https://github.com/lithic-com/lithic-go/issues/45)) ([875dd50](https://github.com/lithic-com/lithic-go/commit/875dd5072eb0574aa499a3fb32d1b136c42b9bfc))


### Refactors

* improve `time.Time` encoding and decoding ([#43](https://github.com/lithic-com/lithic-go/issues/43)) ([cb81762](https://github.com/lithic-com/lithic-go/commit/cb8176278a71c2628704d5aa6acb3a9f168e6984))

## [0.3.2](https://github.com/lithic-com/lithic-go/compare/v0.3.1...v0.3.2) (2023-06-13)


### Features

* implement middleware ([#40](https://github.com/lithic-com/lithic-go/issues/40)) ([2772808](https://github.com/lithic-com/lithic-go/commit/277280844c3ee4cc9627e87f3d5370be6167f94a))


### Documentation

* point to github repo instead of email contact ([#41](https://github.com/lithic-com/lithic-go/issues/41)) ([a2fd417](https://github.com/lithic-com/lithic-go/commit/a2fd41764d365816513e5a647f348d9f3cfb6675))

## [0.3.1](https://github.com/lithic-com/lithic-go/compare/v0.3.0...v0.3.1) (2023-06-12)


### Features

* **api:** update docs ([#26](https://github.com/lithic-com/lithic-go/issues/26)) ([acede1c](https://github.com/lithic-com/lithic-go/commit/acede1ceb6164d14c7889ef58ae248216ba15301))
* make tests give better error message on missing prism server ([#36](https://github.com/lithic-com/lithic-go/issues/36)) ([73d293e](https://github.com/lithic-com/lithic-go/commit/73d293e34cc30c76dc993f9c844a9bfea69b5534))


### Refactors

* improve service types ordering ([#28](https://github.com/lithic-com/lithic-go/issues/28)) ([a0c7143](https://github.com/lithic-com/lithic-go/commit/a0c7143fe09ac3f440a7e2d2ca5859b0c5f9a02b))

## [0.3.0](https://github.com/lithic-com/lithic-go/compare/v0.2.2...v0.3.0) (2023-05-12)


### ⚠ BREAKING CHANGES

* **api:** replace `TransactionToken` param in favour of `TransactionTokens` ([#23](https://github.com/lithic-com/lithic-go/issues/23))

### Refactors

* **api:** replace `TransactionToken` param in favour of `TransactionTokens` ([#23](https://github.com/lithic-com/lithic-go/issues/23)) ([ee26919](https://github.com/lithic-com/lithic-go/commit/ee2691988d3efe2bdf6c9bbe91029d297a8559ee))

## [0.2.2](https://github.com/lithic-com/lithic-go/compare/v0.2.1...v0.2.2) (2023-05-12)


### Refactors

* change `event_types[]` param ([#18](https://github.com/lithic-com/lithic-go/issues/18)) ([d05f43f](https://github.com/lithic-com/lithic-go/commit/d05f43f82e9ea5553bb84c79580f0dd344e9e8ed))

## [0.2.1](https://github.com/lithic-com/lithic-go/compare/v0.2.0...v0.2.1) (2023-05-11)


### Features

* **api:** add support for new `transaction_tokens` query param ([#14](https://github.com/lithic-com/lithic-go/issues/14)) ([d5fa50f](https://github.com/lithic-com/lithic-go/commit/d5fa50f257793794b44e4469b1ad1129eff8ba53))
* **api:** updates ([#12](https://github.com/lithic-com/lithic-go/issues/12)) ([c8fbf7d](https://github.com/lithic-com/lithic-go/commit/c8fbf7de0006f048fcde7853ab5177cc94bb5174))


### Bug Fixes

* **client:** correctly serialise array query params ([#16](https://github.com/lithic-com/lithic-go/issues/16)) ([ad3a4b2](https://github.com/lithic-com/lithic-go/commit/ad3a4b2ed433f0aa68ce2e95193c2003d0dc81ae))

## [0.2.0](https://github.com/lithic-com/lithic-go/compare/v0.1.1...v0.2.0) (2023-05-04)


### ⚠ BREAKING CHANGES

* rename `.JSON.Extras` -> `.JSON.ExtraFields` ([#9](https://github.com/lithic-com/lithic-go/issues/9))

### Code Refactoring

* rename `.JSON.Extras` -&gt; `.JSON.ExtraFields` ([#9](https://github.com/lithic-com/lithic-go/issues/9)) ([4366932](https://github.com/lithic-com/lithic-go/commit/4366932344970e5bf8899cfd7ffba14c42d2382c))

## [0.1.1](https://github.com/lithic-com/lithic-go/compare/v0.1.0...v0.1.1) (2023-05-04)


### Bug Fixes

* **webhooks:** correct parsing of timestamp header ([#6](https://github.com/lithic-com/lithic-go/issues/6)) ([a52ac0a](https://github.com/lithic-com/lithic-go/commit/a52ac0ad7f99f954f88b5927f18ab2de9bc77ff9)), closes [#2](https://github.com/lithic-com/lithic-go/issues/2)

## [0.1.0](https://github.com/lithic-com/lithic-go/compare/v0.0.1...v0.1.0) (2023-05-04)


### ⚠ BREAKING CHANGES

* rename `field.Field` -> `param.Field`
* make JSON structs private, rename Metadata->Field, improve docs
* remove _ in DisputeResolutionReasonNoDisputeRights_3Ds
* **api:** rename _3dsVersion to ThreeDSVersion
* **api:** add tokenization decisioning endpoints and remove unused funding sources API

### Features

* add new services and misc api updates, fix identifier namespacing ([1fb606d](https://github.com/lithic-com/lithic-go/commit/1fb606de6797976367ee177b9aea89c9642f5cd0))
* **api:** add download_url property to dispute evidence ([90fe998](https://github.com/lithic-com/lithic-go/commit/90fe99806180ed20886a3679562076e12cde9da7))
* **api:** add tokenization decisioning endpoints and remove unused funding sources API ([4b89f0a](https://github.com/lithic-com/lithic-go/commit/4b89f0abb56862921f377fa80c5dad98012df25f))
* **api:** more detailed `post /disputes/{dispute_token}/evidences` response ([24e6b12](https://github.com/lithic-com/lithic-go/commit/24e6b12346b4337e6091cd80cdd7b089f1f7e45d))
* better errors ([b4ca8ea](https://github.com/lithic-com/lithic-go/commit/b4ca8ea415652dad20cbc347b2af640c437425c7))
* **docs:** include version references in the README ([7c0f82a](https://github.com/lithic-com/lithic-go/commit/7c0f82ac58ef76d934da75f28dfbf2351915e4c9))
* implement bikesheds ([373d44b](https://github.com/lithic-com/lithic-go/commit/373d44b8decab27a8747b062c1d4a8ee927d6a85))
* implement improved auto-pagination ([ed424a2](https://github.com/lithic-com/lithic-go/commit/ed424a24b71dc0986e486efb153ae801ffcfdb18))
* implement unions ([b6729aa](https://github.com/lithic-com/lithic-go/commit/b6729aaced38a0ee2130559afdcc26647002bc2a))
* implement unions ([b6729aa](https://github.com/lithic-com/lithic-go/commit/b6729aaced38a0ee2130559afdcc26647002bc2a))
* lift fields helpers to main class ([30fa1b7](https://github.com/lithic-com/lithic-go/commit/30fa1b74e6ad8ed8a044e3bdd15417a8b8ee3bdb))
* send package version in X-Stainless-Package-Version ([bbccf72](https://github.com/lithic-com/lithic-go/commit/bbccf72e0785419c000cf1cd9e3e12cbcef4721e))


### Bug Fixes

* **event &gt; payload** type is now any object instead of unknown ([d4c760d](https://github.com/lithic-com/lithic-go/commit/d4c760d271df2ff6b7a7d92612f96a5509d767db))
* add missing properties to AuthRule ([6bc291b](https://github.com/lithic-com/lithic-go/commit/6bc291b19a3216eb4507c1cd2e97aa92c1f03c36))
* **api:** rename _3dsVersion to ThreeDSVersion ([b7c9008](https://github.com/lithic-com/lithic-go/commit/b7c90081cc51cc41dcff660fe03acbfda8d325ba))
* change unknown type generation to `interface{}` ([01976c3](https://github.com/lithic-com/lithic-go/commit/01976c34f76cf4725fcf72073def8ad2eddb094a))
* error that can occur during pagination when there are zero items in the response ([a49d3e1](https://github.com/lithic-com/lithic-go/commit/a49d3e19edff9b6ccc8e4df02c8a077ea8619661))
* error that can occur during pagination when there are zero items in the response ([a49d3e1](https://github.com/lithic-com/lithic-go/commit/a49d3e19edff9b6ccc8e4df02c8a077ea8619661))
* pagination return non-nil on error ([b7d9576](https://github.com/lithic-com/lithic-go/commit/b7d957691767802bbc145c3e79f28c164678236e))
* pagination return non-nil on error ([b7d9576](https://github.com/lithic-com/lithic-go/commit/b7d957691767802bbc145c3e79f28c164678236e))
* remove _ in DisputeResolutionReasonNoDisputeRights_3Ds ([0356df0](https://github.com/lithic-com/lithic-go/commit/0356df02df957ff6afdf3e070f80df0df2da38d2))
* segfault when getting next page if request has no body ([c02a2fd](https://github.com/lithic-com/lithic-go/commit/c02a2fdba60552233a87ba485bd9885d8837d581))
* segfault when getting next page if request has no body ([c02a2fd](https://github.com/lithic-com/lithic-go/commit/c02a2fdba60552233a87ba485bd9885d8837d581))
* segfault when getting next page if request has no body ([3e71bef](https://github.com/lithic-com/lithic-go/commit/3e71bef4b8cce4b4e82d1ac135a595ea1bf68a46))
* update outdate docs in README ([f37232f](https://github.com/lithic-com/lithic-go/commit/f37232f1e65677a778e19d9d6bb2663382209b6c))


### Code Refactoring

* make JSON structs private, rename Metadata-&gt;Field, improve docs ([f6e7936](https://github.com/lithic-com/lithic-go/commit/f6e793680152e9deb9abdd52eaade3e7c7020664))
* rename `field.Field` -&gt; `param.Field` ([97d7533](https://github.com/lithic-com/lithic-go/commit/97d75334d985df2b775414580568f6d9a854e2a5))
