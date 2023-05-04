# Changelog

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
