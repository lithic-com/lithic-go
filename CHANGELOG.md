# Changelog

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
