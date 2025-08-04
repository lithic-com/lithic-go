# Changelog

## 0.87.0 (2025-08-04)

Full Changelog: [v0.86.0...v0.87.0](https://github.com/lithic-com/lithic-go/compare/v0.86.0...v0.87.0)

### Features

* **api:** adds new Account Activity API ([9e3f757](https://github.com/lithic-com/lithic-go/commit/9e3f757cd46626e20dd0f2b855d2dbe903537aa2))

## 0.86.0 (2025-07-28)

Full Changelog: [v0.85.0...v0.86.0](https://github.com/lithic-com/lithic-go/compare/v0.85.0...v0.86.0)

### Features

* **api:** updates Transaction retrieve response to match API ([890f46b](https://github.com/lithic-com/lithic-go/commit/890f46bf7df7b24940186d3dc59b85e45ae5fc0c))

## 0.85.0 (2025-07-23)

Full Changelog: [v0.84.1...v0.85.0](https://github.com/lithic-com/lithic-go/compare/v0.84.1...v0.85.0)

### Features

* **api:** adds new Auth Rules Scope and Settlement Details type ([ec60225](https://github.com/lithic-com/lithic-go/commit/ec60225d2337da3e87807e71e1d2e151c647f24b))

## 0.84.1 (2025-07-21)

Full Changelog: [v0.84.0...v0.84.1](https://github.com/lithic-com/lithic-go/compare/v0.84.0...v0.84.1)

### Bug Fixes

* **client:** process custom base url ahead of time ([9e4e46f](https://github.com/lithic-com/lithic-go/commit/9e4e46f8e187b0b95bf4ca3c27d590db2e98d97e))


### Chores

* **docs:** update Account Holder deprecation formatting ([f041186](https://github.com/lithic-com/lithic-go/commit/f041186ca233cbf7cb0ce3f6d070ebd0b35b0eb3))

## 0.84.0 (2025-07-18)

Full Changelog: [v0.83.0...v0.84.0](https://github.com/lithic-com/lithic-go/compare/v0.83.0...v0.84.0)

### Features

* **api:** adds Network Programs and Account/Card Sub-statuses ([e64b3c3](https://github.com/lithic-com/lithic-go/commit/e64b3c3fc65f96363a9a9b7d2d95fb8e8f11fc13))


### Bug Fixes

* Remove invalid type references in VelocityLimitParamsPeriodWindow union ([e61a9c2](https://github.com/lithic-com/lithic-go/commit/e61a9c2ef85d84f053c98053ed2812fa2036e8c9))


### Chores

* **internal:** fix lint script for tests ([44efcfe](https://github.com/lithic-com/lithic-go/commit/44efcfe371369e5f11af243f59b4905ee0b1cb37))
* lint tests ([d08e1ec](https://github.com/lithic-com/lithic-go/commit/d08e1ec06ced51838de8ae4c0518e25b88f108d2))
* lint tests in subpackages ([9c70704](https://github.com/lithic-com/lithic-go/commit/9c707043dba73e7912b82287d82f4eda843df1ac))

## 0.83.0 (2025-07-03)

Full Changelog: [v0.82.0...v0.83.0](https://github.com/lithic-com/lithic-go/compare/v0.82.0...v0.83.0)

### Features

* **api:** api update ([d626260](https://github.com/lithic-com/lithic-go/commit/d6262600781baf03b58e55dec6d4d095dea57c80))
* **api:** api update ([b956295](https://github.com/lithic-com/lithic-go/commit/b9562955d2661a1568b0a0292cbc66728ae78a69))

## 0.82.0 (2025-07-02)

Full Changelog: [v0.81.0...v0.82.0](https://github.com/lithic-com/lithic-go/compare/v0.81.0...v0.82.0)

### Features

* **api:** add models for merchant_lock_parameters and conditional_3ds_action_parameters ([4af0ff2](https://github.com/lithic-com/lithic-go/commit/4af0ff261e0e216e921289789a8241c2ac7c152c))
* **api:** api update ([d7fd3b6](https://github.com/lithic-com/lithic-go/commit/d7fd3b6e1ae20c912ebe2d7da17027caed351b8e))


### Chores

* **ci:** only run for pushes and fork pull requests ([8c5c5c5](https://github.com/lithic-com/lithic-go/commit/8c5c5c511ce61f4bc95c574fc65a3f50321dd8ce))

## 0.81.0 (2025-06-27)

Full Changelog: [v0.80.1...v0.81.0](https://github.com/lithic-com/lithic-go/compare/v0.80.1...v0.81.0)

### Features

* **api:** introduce dedicated model for SpendLimitDuration ([7b3838e](https://github.com/lithic-com/lithic-go/commit/7b3838e4ca589b9c3857f038e9a293d7b0536c00))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([ef36797](https://github.com/lithic-com/lithic-go/commit/ef3679747269f83a32013ad86004d90265650730))


### Chores

* **internal:** manual updates ([f10c8dc](https://github.com/lithic-com/lithic-go/commit/f10c8dc233cb4f407753e5aa22c3b20416fdc5d7))

## 0.80.1 (2025-06-27)

Full Changelog: [v0.80.0...v0.80.1](https://github.com/lithic-com/lithic-go/compare/v0.80.0...v0.80.1)

### Bug Fixes

* **pagination:** check if page data is empty in GetNextPage ([8a2b637](https://github.com/lithic-com/lithic-go/commit/8a2b637516190ad244ec0fa1c33cee3955c5b52f))

## 0.80.0 (2025-06-26)

Full Changelog: [v0.79.0...v0.80.0](https://github.com/lithic-com/lithic-go/compare/v0.79.0...v0.80.0)

### Features

* **client:** adds support for on-demand Auth Rule Performance Reports ([2dc23fa](https://github.com/lithic-com/lithic-go/commit/2dc23fae2e661ca0362c1e1d4341b9bd7919c1b6))

## 0.79.0 (2025-06-20)

Full Changelog: [v0.78.0...v0.79.0](https://github.com/lithic-com/lithic-go/compare/v0.78.0...v0.79.0)

### Features

* **api:** add CLOSED account state option and UNVERIFIED verification method ([1106e29](https://github.com/lithic-com/lithic-go/commit/1106e29df2648172730bc56341706dfdbd31821e))
* **client:** add debug log helper ([8da37b4](https://github.com/lithic-com/lithic-go/commit/8da37b4b0a2de149703d961dc7002616fb4b630d))
* **client:** adds endpoint to register an account number on a Financial Account ([8917aeb](https://github.com/lithic-com/lithic-go/commit/8917aeb222e9c69f5f3d6cf729bd146167152855))
* **client:** adds support for 3DS to Auth Rules ([37afe17](https://github.com/lithic-com/lithic-go/commit/37afe17532a7b6710855eda29875a8fbd2d9d63d))


### Chores

* **api:** mark some methods as deprecated ([9d06e4a](https://github.com/lithic-com/lithic-go/commit/9d06e4a5a43feccd874d2951172672bfc242b61e))
* **ci:** enable for pull requests ([13e25ed](https://github.com/lithic-com/lithic-go/commit/13e25edf26288c74571afa6aede47c761f3c1ed8))
* **docs:** grammar improvements ([e215a47](https://github.com/lithic-com/lithic-go/commit/e215a475ffe5414ee77cf1d7b647d8aebbd67d6f))
* improve devcontainer setup ([1edb00d](https://github.com/lithic-com/lithic-go/commit/1edb00df3ef8f2f2a8d4c6dfd793da5b84f8c568))
* make go mod tidy continue on error ([20d4676](https://github.com/lithic-com/lithic-go/commit/20d4676b3c0d62574fce37d1485038c43fcaef5e))

## 0.78.0 (2025-05-15)

Full Changelog: [v0.77.0...v0.78.0](https://github.com/lithic-com/lithic-go/compare/v0.77.0...v0.78.0)

### Features

* **api:** new Funding Events and Card Web Provision API's ([d8c8675](https://github.com/lithic-com/lithic-go/commit/d8c867508174aa82dc445951950427d65b79a356))
* **client:** add support for endpoint-specific base URLs in python ([958dddd](https://github.com/lithic-com/lithic-go/commit/958dddd491fda1ead9aa67844d4053e274aae2c4))

## 0.77.0 (2025-05-12)

Full Changelog: [v0.76.5...v0.77.0](https://github.com/lithic-com/lithic-go/compare/v0.76.5...v0.77.0)

### Features

* **api:** manual updates ([447255f](https://github.com/lithic-com/lithic-go/commit/447255f2ad25c9c5fe52b4259243e0bc0df01535))


### Documentation

* remove or fix invalid readme examples ([df5fc3e](https://github.com/lithic-com/lithic-go/commit/df5fc3e6f3be03bc1ff8d52321ced910c4765930))

## 0.76.5 (2025-05-06)

Full Changelog: [v0.76.4...v0.76.5](https://github.com/lithic-com/lithic-go/compare/v0.76.4...v0.76.5)

### Bug Fixes

* **client:** clean up reader resources ([82aaec8](https://github.com/lithic-com/lithic-go/commit/82aaec8bd1852f7c18fa2c019bead7d8e994dea2))

## 0.76.4 (2025-05-06)

Full Changelog: [v0.76.3...v0.76.4](https://github.com/lithic-com/lithic-go/compare/v0.76.3...v0.76.4)

### Bug Fixes

* **client:** correctly update body in WithJSONSet ([0751972](https://github.com/lithic-com/lithic-go/commit/07519720cff659cd7c9a561b38c5a7fc5c89f013))

## 0.76.3 (2025-05-05)

Full Changelog: [v0.76.2...v0.76.3](https://github.com/lithic-com/lithic-go/compare/v0.76.2...v0.76.3)

### Bug Fixes

* **internals:** fix servers entry for /v2/auth_rules.get ([c4becf6](https://github.com/lithic-com/lithic-go/commit/c4becf61caff760a0072fc867281eca7d9fb08e2))

## 0.76.2 (2025-04-30)

Full Changelog: [v0.76.1...v0.76.2](https://github.com/lithic-com/lithic-go/compare/v0.76.1...v0.76.2)

### Bug Fixes

* **pagination:** handle errors when applying options ([97ab73f](https://github.com/lithic-com/lithic-go/commit/97ab73f6fecc937227f8b7dd2928f2bd2a3391e1))

## 0.76.1 (2025-04-29)

Full Changelog: [v0.76.0...v0.76.1](https://github.com/lithic-com/lithic-go/compare/v0.76.0...v0.76.1)

### Bug Fixes

* handle empty bodies in WithJSONSet ([797902e](https://github.com/lithic-com/lithic-go/commit/797902e763686bd05f78b5dce45db7612b2fd33e))

## 0.76.0 (2025-04-29)

Full Changelog: [v0.75.0...v0.76.0](https://github.com/lithic-com/lithic-go/compare/v0.75.0...v0.76.0)

### Features

* **api:** adds new merchant lock Auth Rule ([fb8e53b](https://github.com/lithic-com/lithic-go/commit/fb8e53b7766aa11b965ef388d59980b74cb6f741))


### Chores

* **ci:** add timeout thresholds for CI jobs ([bf2acdb](https://github.com/lithic-com/lithic-go/commit/bf2acdb6c465016d3c92ef61b3ea847cdbfc83a5))
* **ci:** only use depot for staging repos ([9e57f70](https://github.com/lithic-com/lithic-go/commit/9e57f70ef505d9480441c559d1eadba8a5fe0bab))
* **ci:** run on more branches and use depot runners ([66dc16f](https://github.com/lithic-com/lithic-go/commit/66dc16f3d8d3db69d147faffd64d244114a58282))

## 0.75.0 (2025-04-21)

Full Changelog: [v0.74.0...v0.75.0](https://github.com/lithic-com/lithic-go/compare/v0.74.0...v0.75.0)

### Features

* **api:** updates to Card definition for PCI clarity ([dd32713](https://github.com/lithic-com/lithic-go/commit/dd327135f437a317cd19697768668bc1360aa04e))
* **client:** add support for reading base URL from environment variable ([5642ecb](https://github.com/lithic-com/lithic-go/commit/5642ecbf539f20a354349635de38ff365011c67f))


### Bug Fixes

* **docs:** fix Card in HTML example ([0586d2f](https://github.com/lithic-com/lithic-go/commit/0586d2f885c4d1f0fa56947ca577f102b37d760b))
* **internal:** refresh schemas ([bf2256d](https://github.com/lithic-com/lithic-go/commit/bf2256d9d38e3560a2ec0ed4c43d360e1af75bdc))
* **internals:** fix Card schema definition ([1f3534c](https://github.com/lithic-com/lithic-go/commit/1f3534c43ed127e8457d9918fba7bfbcf8ff0c54))


### Chores

* **docs:** document pre-request options ([d1d78cb](https://github.com/lithic-com/lithic-go/commit/d1d78cbd4a74e96ad2566601acbd067bfc9fd36e))
* **internal:** codegen related update ([f7d5244](https://github.com/lithic-com/lithic-go/commit/f7d5244ee953a29d2b748d9e0dd3a1a93f70d3e3))
* **internal:** reduce CI branch coverage ([628a86c](https://github.com/lithic-com/lithic-go/commit/628a86c6da9466df2d1810b9d23c8a87b243a55b))


### Documentation

* update documentation links to be more uniform ([cbb3f52](https://github.com/lithic-com/lithic-go/commit/cbb3f528661b87c444281849a586a5522045020f))

## 0.74.0 (2025-04-09)

Full Changelog: [v0.73.0...v0.74.0](https://github.com/lithic-com/lithic-go/compare/v0.73.0...v0.74.0)

### Features

* **api:** manual updates ([2ecf324](https://github.com/lithic-com/lithic-go/commit/2ecf324375bdcdb22936c7f7c7f39691e979fc76))


### Chores

* configure new SDK language ([edf97d0](https://github.com/lithic-com/lithic-go/commit/edf97d0f30919ac51b052a92036ab21a8249fd81))
* **internal:** expand CI branch coverage ([#495](https://github.com/lithic-com/lithic-go/issues/495)) ([c50cc15](https://github.com/lithic-com/lithic-go/commit/c50cc15de9de14095862e31320d552f83c7b70e9))

## 0.73.0 (2025-04-08)

Full Changelog: [v0.72.0...v0.73.0](https://github.com/lithic-com/lithic-go/compare/v0.72.0...v0.73.0)

### Features

* **api:** introduce TransactionSeries and update ShippingMethod fields ([#493](https://github.com/lithic-com/lithic-go/issues/493)) ([ecf8985](https://github.com/lithic-com/lithic-go/commit/ecf8985c28c54c79b58cfc9517f8618438ad5dd4))
* **client:** support custom http clients ([#491](https://github.com/lithic-com/lithic-go/issues/491)) ([ef921bb](https://github.com/lithic-com/lithic-go/commit/ef921bb32a5c383ed8d3f208eaa775e992662dfe))


### Chores

* **tests:** improve enum examples ([#494](https://github.com/lithic-com/lithic-go/issues/494)) ([f3b6d31](https://github.com/lithic-com/lithic-go/commit/f3b6d316eb0d1b0bfc665daa7878791505127381))

## 0.72.0 (2025-04-07)

Full Changelog: [v0.71.1...v0.72.0](https://github.com/lithic-com/lithic-go/compare/v0.71.1...v0.72.0)

### Features

* **api:** new resend endpoint for Event Subscriptions ([#488](https://github.com/lithic-com/lithic-go/issues/488)) ([770514b](https://github.com/lithic-com/lithic-go/commit/770514bf6e0792123ef985333a9586212579919b))

## 0.71.1 (2025-04-02)

Full Changelog: [v0.71.0...v0.71.1](https://github.com/lithic-com/lithic-go/compare/v0.71.0...v0.71.1)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#486](https://github.com/lithic-com/lithic-go/issues/486)) ([1c43550](https://github.com/lithic-com/lithic-go/commit/1c43550452db2bdc5f92d53587e26b4eb2bd5bb5))


### Chores

* fix typos ([#482](https://github.com/lithic-com/lithic-go/issues/482)) ([885d231](https://github.com/lithic-com/lithic-go/commit/885d23189694ceefbbe064d40c0253274ff562ce))
* internal codegen changes ([7f19e3a](https://github.com/lithic-com/lithic-go/commit/7f19e3ac6968eabb3d07bc03069f512319d29e6c))
* internal codegen changes ([9b25269](https://github.com/lithic-com/lithic-go/commit/9b25269e63768dc824d5014999624561e41fba0c))
* **internal:** remove workflow condition ([#485](https://github.com/lithic-com/lithic-go/issues/485)) ([0a6bdd8](https://github.com/lithic-com/lithic-go/commit/0a6bdd83ed71986e52e943b9c63e0e4890885cc6))

## 0.71.0 (2025-03-25)

Full Changelog: [v0.70.0...v0.71.0](https://github.com/lithic-com/lithic-go/compare/v0.70.0...v0.71.0)

### Features

* **client:** support v2 ([#477](https://github.com/lithic-com/lithic-go/issues/477)) ([bc010e2](https://github.com/lithic-com/lithic-go/commit/bc010e2aab8e6ede9e5794b9ec5649f3c8e9d0fc))


### Bug Fixes

* **test:** return early after test failure ([#481](https://github.com/lithic-com/lithic-go/issues/481)) ([db2e73f](https://github.com/lithic-com/lithic-go/commit/db2e73ffabe7b4c7fcf2921c03e4eedb22c17219))


### Chores

* add request options to client tests ([#480](https://github.com/lithic-com/lithic-go/issues/480)) ([03d8154](https://github.com/lithic-com/lithic-go/commit/03d8154625dbd600f48da1aeda373546828c2ed5))
* **api:** new attribute targets for Auth Rules and new Financial Account State schema ([#479](https://github.com/lithic-com/lithic-go/issues/479)) ([2e8ec92](https://github.com/lithic-com/lithic-go/commit/2e8ec924be41c370f5d7b009dd2de4eebdab6084))
* **internal:** update .stats.yml ([#475](https://github.com/lithic-com/lithic-go/issues/475)) ([e0c2065](https://github.com/lithic-com/lithic-go/commit/e0c20655be86b2c3f32fe6e00352dbd569477691))

## 0.70.0 (2025-03-18)

Full Changelog: [v0.69.0...v0.70.0](https://github.com/lithic-com/lithic-go/compare/v0.69.0...v0.70.0)

### Features

* **api:** updates to 2 `FinancialAccounts` endpoints and new `ExpireAuthorization` endpoint ([#474](https://github.com/lithic-com/lithic-go/issues/474)) ([7c08b54](https://github.com/lithic-com/lithic-go/commit/7c08b54e3327e4f614d436387d089446fc13f776))
* **client:** improve default client options support ([#471](https://github.com/lithic-com/lithic-go/issues/471)) ([dddd5f9](https://github.com/lithic-com/lithic-go/commit/dddd5f97b45737bd4985f0a9990a99d7506df656))


### Chores

* **internal:** remove extra empty newlines ([#473](https://github.com/lithic-com/lithic-go/issues/473)) ([596060e](https://github.com/lithic-com/lithic-go/commit/596060e680f963368c91f5bdc1d0e5923787768b))

## 0.69.0 (2025-03-11)

Full Changelog: [v0.68.0...v0.69.0](https://github.com/lithic-com/lithic-go/compare/v0.68.0...v0.69.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#466](https://github.com/lithic-com/lithic-go/issues/466)) ([72d2f6d](https://github.com/lithic-com/lithic-go/commit/72d2f6d155ec664667c887d0d46694c035207878))
* **client:** accept RFC6838 JSON content types ([#467](https://github.com/lithic-com/lithic-go/issues/467)) ([26250d6](https://github.com/lithic-com/lithic-go/commit/26250d6143f90e520c736b3fca8831fbd7f021e8))
* **client:** allow custom baseurls without trailing slash ([#464](https://github.com/lithic-com/lithic-go/issues/464)) ([5bd2500](https://github.com/lithic-com/lithic-go/commit/5bd2500c5b39d22d751a5b7141d90c9e518f9169))
* **client:** update currency data type ([#470](https://github.com/lithic-com/lithic-go/issues/470)) ([011d635](https://github.com/lithic-com/lithic-go/commit/011d635da782ed92b76fd2788783d8855d0c0adf))


### Chores

* **api:** release of Network Totals reporting and new filters for Velocity Limit Rules ([#469](https://github.com/lithic-com/lithic-go/issues/469)) ([7047b1d](https://github.com/lithic-com/lithic-go/commit/7047b1dbc776d11541a00232e3ede7d5af069aa8))
* **client:** deprecate some fields ([011d635](https://github.com/lithic-com/lithic-go/commit/011d635da782ed92b76fd2788783d8855d0c0adf))


### Documentation

* update some descriptions ([011d635](https://github.com/lithic-com/lithic-go/commit/011d635da782ed92b76fd2788783d8855d0c0adf))


### Refactors

* tidy up dependencies ([#468](https://github.com/lithic-com/lithic-go/issues/468)) ([b063b41](https://github.com/lithic-com/lithic-go/commit/b063b41c082b43defe76584b0cfe9b903dc9c781))

## 0.68.0 (2025-03-04)

Full Changelog: [v0.67.3...v0.68.0](https://github.com/lithic-com/lithic-go/compare/v0.67.3...v0.68.0)

### Features

* **api:** new Settlement API endpoints and changes to update Account Holder endpoint ([#463](https://github.com/lithic-com/lithic-go/issues/463)) ([ba9317e](https://github.com/lithic-com/lithic-go/commit/ba9317e7236d6fe06587a868774d5803f810dff2))


### Chores

* **api:** adds new `Internal` Category for FinancialTransactions ([#461](https://github.com/lithic-com/lithic-go/issues/461)) ([0d75f20](https://github.com/lithic-com/lithic-go/commit/0d75f2031129724e6ed6ea8c55001afba748e521))
* **internal:** fix devcontainers setup ([#459](https://github.com/lithic-com/lithic-go/issues/459)) ([93f5ea2](https://github.com/lithic-com/lithic-go/commit/93f5ea2b819b2fbd1350b6c1928d7d33b8d5fb39))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#462](https://github.com/lithic-com/lithic-go/issues/462)) ([579a8fc](https://github.com/lithic-com/lithic-go/commit/579a8fc26087b8a2b279ffaa1f1db45eaee9c1e9))

## 0.67.3 (2025-02-20)

Full Changelog: [v0.67.2...v0.67.3](https://github.com/lithic-com/lithic-go/compare/v0.67.2...v0.67.3)

### Bug Fixes

* **client:** mark some request bodies as optional ([#457](https://github.com/lithic-com/lithic-go/issues/457)) ([7d80ac3](https://github.com/lithic-com/lithic-go/commit/7d80ac374c8637e09072a26d0c8f35821c8a3b1e))

## 0.67.2 (2025-02-14)

Full Changelog: [v0.67.1...v0.67.2](https://github.com/lithic-com/lithic-go/compare/v0.67.1...v0.67.2)

### Bug Fixes

* **client:** don't truncate manually specified filenames ([#456](https://github.com/lithic-com/lithic-go/issues/456)) ([2d769b5](https://github.com/lithic-com/lithic-go/commit/2d769b5c4dc14ff44b719fd8767729a51e186558))


### Chores

* **internal:** update test values ([#454](https://github.com/lithic-com/lithic-go/issues/454)) ([f7bfb18](https://github.com/lithic-com/lithic-go/commit/f7bfb18917b43f831ef4c29f670df8d837a1d06b))

## 0.67.1 (2025-02-11)

Full Changelog: [v0.67.0...v0.67.1](https://github.com/lithic-com/lithic-go/compare/v0.67.0...v0.67.1)

### Bug Fixes

* do not call path.Base on ContentType ([#451](https://github.com/lithic-com/lithic-go/issues/451)) ([3310645](https://github.com/lithic-com/lithic-go/commit/331064504c7fc903948563463853db141b0b9e0d))


### Chores

* **api:** new 3DS Event and new `challenge_metadata` property on Authentications ([#453](https://github.com/lithic-com/lithic-go/issues/453)) ([7abae71](https://github.com/lithic-com/lithic-go/commit/7abae71aacf6f9d9e56ccacbf04588a4abeef617))

## 0.67.0 (2025-02-07)

Full Changelog: [v0.66.1...v0.67.0](https://github.com/lithic-com/lithic-go/compare/v0.66.1...v0.67.0)

### Features

* **client:** send `X-Stainless-Timeout` header ([#445](https://github.com/lithic-com/lithic-go/issues/445)) ([41ef2f1](https://github.com/lithic-com/lithic-go/commit/41ef2f1f3d8158b4244a7a53a9b6ee97a315b35a))
* **pagination:** avoid fetching when has_more: false ([#449](https://github.com/lithic-com/lithic-go/issues/449)) ([5de31e5](https://github.com/lithic-com/lithic-go/commit/5de31e5f6bdf08b7068084d96c7e58e75ac221f4))


### Bug Fixes

* fix early cancel when RequestTimeout is provided for streaming requests ([#450](https://github.com/lithic-com/lithic-go/issues/450)) ([f165e52](https://github.com/lithic-com/lithic-go/commit/f165e52fba83b7efa72dcc7df21a89d06ce14cd0))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#448](https://github.com/lithic-com/lithic-go/issues/448)) ([5af84eb](https://github.com/lithic-com/lithic-go/commit/5af84eb0f6db1f21c19f139c7f80b62376e80c05))
* **api:** new PaymentEventType for ACH Returns and small updates to 3DS AuthenticationResult ([#447](https://github.com/lithic-com/lithic-go/issues/447)) ([9ba6cbb](https://github.com/lithic-com/lithic-go/commit/9ba6cbb1c4f17670e3006d2b87100ba0b7ec03f5))

## 0.66.1 (2025-01-31)

Full Changelog: [v0.66.0...v0.66.1](https://github.com/lithic-com/lithic-go/compare/v0.66.0...v0.66.1)

### Bug Fixes

* fix unicode encoding for json ([#443](https://github.com/lithic-com/lithic-go/issues/443)) ([59773ed](https://github.com/lithic-com/lithic-go/commit/59773ed1883db7f61a2551e41e4f3e83ee0a25e7))


### Chores

* refactor client tests ([#441](https://github.com/lithic-com/lithic-go/issues/441)) ([14606ef](https://github.com/lithic-com/lithic-go/commit/14606eff396afde7ab77f9830005d309615bcae7))


### Documentation

* document raw responses ([#444](https://github.com/lithic-com/lithic-go/issues/444)) ([7dd3568](https://github.com/lithic-com/lithic-go/commit/7dd356843bd6e00a2f313e76acfe969b22778ab1))

## 0.66.0 (2025-01-28)

Full Changelog: [v0.65.0...v0.66.0](https://github.com/lithic-com/lithic-go/compare/v0.65.0...v0.66.0)

### Features

* **api:** adds additional request types for updating an Auth Rule ([#438](https://github.com/lithic-com/lithic-go/issues/438)) ([eda879f](https://github.com/lithic-com/lithic-go/commit/eda879f4e8be7bc18dfcbf2ae19715b334e79c97))


### Bug Fixes

* fix interface implementation stub names for unions ([#440](https://github.com/lithic-com/lithic-go/issues/440)) ([b0b3e45](https://github.com/lithic-com/lithic-go/commit/b0b3e456f92720fda3681a1e576c4f03adaa1e30))

## 0.65.0 (2025-01-24)

Full Changelog: [v0.64.0...v0.65.0](https://github.com/lithic-com/lithic-go/compare/v0.64.0...v0.65.0)

### Features

* **api:** adds additional fields to TransactionEvents ([#437](https://github.com/lithic-com/lithic-go/issues/437)) ([6dffc6d](https://github.com/lithic-com/lithic-go/commit/6dffc6d33ab928b9fa66d777f676b623c8b154c4))


### Chores

* **api:** additional field added to 3DS Responses and Tokenization ([#435](https://github.com/lithic-com/lithic-go/issues/435)) ([237c922](https://github.com/lithic-com/lithic-go/commit/237c92222cb24dd37747309d13600e76e688c3c6))

## 0.64.0 (2025-01-21)

Full Changelog: [v0.63.0...v0.64.0](https://github.com/lithic-com/lithic-go/compare/v0.63.0...v0.64.0)

### ⚠ BREAKING CHANGES

* **types:** improve auth rules types ([#428](https://github.com/lithic-com/lithic-go/issues/428))
* **api:** removes AccountHolder `resubmit` endpoint and `KYC_ADVANCED` workflow ([#422](https://github.com/lithic-com/lithic-go/issues/422))

### Features

* **api:** adds EventRuleResult to Transaction Events ([#420](https://github.com/lithic-com/lithic-go/issues/420)) ([c7a271e](https://github.com/lithic-com/lithic-go/commit/c7a271edc42cc0103184ffab170bf5d98c53a628))
* **api:** removes AccountHolder `resubmit` endpoint and `KYC_ADVANCED` workflow ([#422](https://github.com/lithic-com/lithic-go/issues/422)) ([ee46446](https://github.com/lithic-com/lithic-go/commit/ee4644639e08df0837aacddcb64c8e567ab42097))
* **api:** updates to Auth Rules numeric types, new Card Types and Authorization Rule Backtests ([bf025e3](https://github.com/lithic-com/lithic-go/commit/bf025e3a7e0cf9ab9a295715626d1de713d35e8b))
* support deprecated markers ([#434](https://github.com/lithic-com/lithic-go/issues/434)) ([7dcdf7d](https://github.com/lithic-com/lithic-go/commit/7dcdf7d2b03b07fb964c14625d1627245afae32c))
* **types:** improve auth rules types ([#428](https://github.com/lithic-com/lithic-go/issues/428)) ([c90f44e](https://github.com/lithic-com/lithic-go/commit/c90f44e591e25deffa2bf0aae34245cbfb6b9899))


### Bug Fixes

* fix apijson.Port for embedded structs ([#431](https://github.com/lithic-com/lithic-go/issues/431)) ([9162400](https://github.com/lithic-com/lithic-go/commit/9162400a9b1cf570af9a0c79686892bfea1f160b))
* fix apijson.Port for embedded structs ([#432](https://github.com/lithic-com/lithic-go/issues/432)) ([87bfd6f](https://github.com/lithic-com/lithic-go/commit/87bfd6ffa8b1a94a1a0d6ebbf3df678d04b92657))
* **internal:** update next ([88a36f9](https://github.com/lithic-com/lithic-go/commit/88a36f95633138f23af7d28014390c03b699b5fb))
* reuse model in pagination items type ([#433](https://github.com/lithic-com/lithic-go/issues/433)) ([3c5bd15](https://github.com/lithic-com/lithic-go/commit/3c5bd159ff07a68855636ec8c429e50210cf2351))


### Chores

* **api:** add backtest methods to AuthRules ([#419](https://github.com/lithic-com/lithic-go/issues/419)) ([4d63e87](https://github.com/lithic-com/lithic-go/commit/4d63e87a0c7cd01c5d4e65541c0741f6253ab859))
* **api:** adds `dpan` property to Tokenization ([#429](https://github.com/lithic-com/lithic-go/issues/429)) ([5b07971](https://github.com/lithic-com/lithic-go/commit/5b07971752446560331b39267009084eda079bb8))
* **api:** new ConvertPhysical endpoint to convert a virtual card to a physical card ([#421](https://github.com/lithic-com/lithic-go/issues/421)) ([eea41f0](https://github.com/lithic-com/lithic-go/commit/eea41f03f7cd3e6b0776008e9162bf2c27fdbd28))
* **api:** updates to documentation and additional filter for status on Transactions ([#427](https://github.com/lithic-com/lithic-go/issues/427)) ([ef8d694](https://github.com/lithic-com/lithic-go/commit/ef8d694095607c0252827bc27b1aac6fd8533b3b))
* bump license year ([#424](https://github.com/lithic-com/lithic-go/issues/424)) ([d1f5120](https://github.com/lithic-com/lithic-go/commit/d1f5120dce6743f01cc0ff87a8a32800e56e3f29))
* **docs:** updates documentation for DPANs ([#430](https://github.com/lithic-com/lithic-go/issues/430)) ([35260e6](https://github.com/lithic-com/lithic-go/commit/35260e65910787a960ab6f832a3e580859a321d9))
* **internal:** update examples ([#425](https://github.com/lithic-com/lithic-go/issues/425)) ([69db4a6](https://github.com/lithic-com/lithic-go/commit/69db4a6c9e439449546e273f6cea28e9253f376f))


### Documentation

* **readme:** fix misplaced period ([#426](https://github.com/lithic-com/lithic-go/issues/426)) ([7908a0f](https://github.com/lithic-com/lithic-go/commit/7908a0f34301bcd27fd8f0a94538594e5045d723))
* **readme:** fix typo ([#423](https://github.com/lithic-com/lithic-go/issues/423)) ([d46e23f](https://github.com/lithic-com/lithic-go/commit/d46e23fef17323800012f1cc099a3ae0de3c8ddf))

## 0.63.0 (2024-11-19)

Full Changelog: [v0.62.3...v0.63.0](https://github.com/lithic-com/lithic-go/compare/v0.62.3...v0.63.0)

### Features

* **api:** adds PrimeRates API ([#413](https://github.com/lithic-com/lithic-go/issues/413)) ([f8ec404](https://github.com/lithic-com/lithic-go/commit/f8ec404c7d2377436b000612d15fa145d0cb0aa9))


### Chores

* **tests:** limit array example length ([#415](https://github.com/lithic-com/lithic-go/issues/415)) ([022b222](https://github.com/lithic-com/lithic-go/commit/022b2224e9a6737a046d0ad47213776f6ce11d9e))

## 0.62.3 (2024-11-11)

Full Changelog: [v0.62.2...v0.62.3](https://github.com/lithic-com/lithic-go/compare/v0.62.2...v0.62.3)

### Bug Fixes

* **client:** no panic on missing BaseURL ([#411](https://github.com/lithic-com/lithic-go/issues/411)) ([fd4d510](https://github.com/lithic-com/lithic-go/commit/fd4d510d400b78683968c1d3707de2c8dc85e533))

## 0.62.2 (2024-11-11)

Full Changelog: [v0.62.1...v0.62.2](https://github.com/lithic-com/lithic-go/compare/v0.62.1...v0.62.2)

### Bug Fixes

* correct required fields for flattened unions ([#410](https://github.com/lithic-com/lithic-go/issues/410)) ([b42cfdb](https://github.com/lithic-com/lithic-go/commit/b42cfdb6bf83bfdb5e630b8e97f56c034cdb8f5a))


### Chores

* **api:** add business_account_token param for listing Balances ([#408](https://github.com/lithic-com/lithic-go/issues/408)) ([9c5c9db](https://github.com/lithic-com/lithic-go/commit/9c5c9db42f29817f12fc1e3337865cbb9409328f))

## 0.62.1 (2024-11-08)

Full Changelog: [v0.62.0...v0.62.1](https://github.com/lithic-com/lithic-go/compare/v0.62.0...v0.62.1)

### Bug Fixes

* **api:** escape key values when encoding maps ([#407](https://github.com/lithic-com/lithic-go/issues/407)) ([fc95aad](https://github.com/lithic-com/lithic-go/commit/fc95aadd2a99b0e5e8fa8afcf14c3d95d68e1b17))


### Chores

* **api:** adds replacement_account_token to Card create parameters ([#406](https://github.com/lithic-com/lithic-go/issues/406)) ([dd703be](https://github.com/lithic-com/lithic-go/commit/dd703be858ad4700dcb752849be4b98d47eac8e8))


### Refactors

* sort fields for squashed union structs ([#404](https://github.com/lithic-com/lithic-go/issues/404)) ([635b5dd](https://github.com/lithic-com/lithic-go/commit/635b5dda2307d1da4b093f3199f4ea60dd93dd41))

## 0.62.0 (2024-11-05)

Full Changelog: [v0.61.0...v0.62.0](https://github.com/lithic-com/lithic-go/compare/v0.61.0...v0.62.0)

### ⚠ BREAKING CHANGES

* **api:** removes AuthRules V1 ([#403](https://github.com/lithic-com/lithic-go/issues/403))

### Features

* **api:** removes AuthRules V1 ([#403](https://github.com/lithic-com/lithic-go/issues/403)) ([347ca15](https://github.com/lithic-com/lithic-go/commit/347ca157e5acb67d507f6044b8980b75b52c13fa))


### Chores

* **api:** adds `charge_off` functionality to FinancialAccounts ([#402](https://github.com/lithic-com/lithic-go/issues/402)) ([1935628](https://github.com/lithic-com/lithic-go/commit/19356282d48b9743c6360d49cfdd56ebd0eb238c))
* **docs:** updates to documentation for V2 AuthRules ([#400](https://github.com/lithic-com/lithic-go/issues/400)) ([fbc8377](https://github.com/lithic-com/lithic-go/commit/fbc837786077ee053167b0fcaf904d86ed9fbe71))

## 0.61.0 (2024-10-28)

Full Changelog: [v0.60.0...v0.61.0](https://github.com/lithic-com/lithic-go/compare/v0.60.0...v0.61.0)

### Features

* **api:** updates ([#399](https://github.com/lithic-com/lithic-go/issues/399)) ([4203853](https://github.com/lithic-com/lithic-go/commit/4203853bde9c3f10c767a216347da68bf4352949))


### Chores

* **api:** add `PIN_BLOCKED` to `detailed_results` property on Event ([#396](https://github.com/lithic-com/lithic-go/issues/396)) ([64d9116](https://github.com/lithic-com/lithic-go/commit/64d911658dc866ba9604823766a29b9d5eff43a5))
* **api:** adds new result types to Transactions and Events ([#398](https://github.com/lithic-com/lithic-go/issues/398)) ([47a8015](https://github.com/lithic-com/lithic-go/commit/47a801526d37a882f5e1148a6c135d80af9b91f5))

## 0.60.0 (2024-10-23)

Full Changelog: [v0.59.0...v0.60.0](https://github.com/lithic-com/lithic-go/compare/v0.59.0...v0.60.0)

### Features

* **api:** add `interest_details` properties to LoanTapes ([#394](https://github.com/lithic-com/lithic-go/issues/394)) ([98f10bc](https://github.com/lithic-com/lithic-go/commit/98f10bc304d0bad42719fbe3e497066454f5ec64))

## 0.59.0 (2024-10-22)

Full Changelog: [v0.58.0...v0.59.0](https://github.com/lithic-com/lithic-go/compare/v0.58.0...v0.59.0)

### Features

* **api:** removes `transfer_transaction.created` webhook and updates to VerificationApplication ([#393](https://github.com/lithic-com/lithic-go/issues/393)) ([5658bfd](https://github.com/lithic-com/lithic-go/commit/5658bfd250b45dbea868dbbfbac8bbab80317e18))
* move pagination package from internal to packages ([#391](https://github.com/lithic-com/lithic-go/issues/391)) ([4d34239](https://github.com/lithic-com/lithic-go/commit/4d34239d575fcff51017b642cef309b77ddca7aa))

## 0.58.0 (2024-10-11)

Full Changelog: [v0.57.0...v0.58.0](https://github.com/lithic-com/lithic-go/compare/v0.57.0...v0.58.0)

### Features

* **api:** updates to documentation and addition of new 3DS simulation methods ([#390](https://github.com/lithic-com/lithic-go/issues/390)) ([922e670](https://github.com/lithic-com/lithic-go/commit/922e6708668ef16a32360010f2f05da0937cce0b))


### Chores

* fix GetNextPage docstring ([#388](https://github.com/lithic-com/lithic-go/issues/388)) ([eec6cc7](https://github.com/lithic-com/lithic-go/commit/eec6cc72d16d63aa439efa508c67cab2fc8ac9cb))

## 0.57.0 (2024-10-09)

Full Changelog: [v0.56.0...v0.57.0](https://github.com/lithic-com/lithic-go/compare/v0.56.0...v0.57.0)

### Features

* **api:** small updates to Documents, AccountHolders and LoanTapes ([#387](https://github.com/lithic-com/lithic-go/issues/387)) ([20cb506](https://github.com/lithic-com/lithic-go/commit/20cb506fb9094a2efd7275ca728264328e24dca0))


### Chores

* **api:** small updates to verification addresses and Statement and LoanTape fields ([#385](https://github.com/lithic-com/lithic-go/issues/385)) ([bd05727](https://github.com/lithic-com/lithic-go/commit/bd0572705814f44a6436247339dee6c594de02fb))

## 0.56.0 (2024-10-01)

Full Changelog: [v0.55.0...v0.56.0](https://github.com/lithic-com/lithic-go/compare/v0.55.0...v0.56.0)

### Features

* **api:** rename `loan_tape_response.statement_balance` to `previous_statement_balance` ([#384](https://github.com/lithic-com/lithic-go/issues/384)) ([43ab10a](https://github.com/lithic-com/lithic-go/commit/43ab10a2122d6b5955e80398900bea0c5bada633))


### Documentation

* improve and reference contributing documentation ([#382](https://github.com/lithic-com/lithic-go/issues/382)) ([c1f5162](https://github.com/lithic-com/lithic-go/commit/c1f5162de15cb3c7ed0cd7baa4af326d12bcee68))

## 0.55.0 (2024-10-01)

Full Changelog: [v0.54.0...v0.55.0](https://github.com/lithic-com/lithic-go/compare/v0.54.0...v0.55.0)

### Features

* **api:** add Management Operations and Loan Tapes API ([#380](https://github.com/lithic-com/lithic-go/issues/380)) ([06b39cf](https://github.com/lithic-com/lithic-go/commit/06b39cfe31e940439abc952c745c2c039acbeeec))

## 0.54.0 (2024-09-25)

Full Changelog: [v0.53.0...v0.54.0](https://github.com/lithic-com/lithic-go/compare/v0.53.0...v0.54.0)

### Features

* **api:** adds endpoint for migrating auth rules from v1 to V2. marks v1 auth rules as deprecated ([#379](https://github.com/lithic-com/lithic-go/issues/379)) ([4af648c](https://github.com/lithic-com/lithic-go/commit/4af648c6ef59af0bcf503eccd068ccff483666b6))
* **client:** send retry count header ([#377](https://github.com/lithic-com/lithic-go/issues/377)) ([741d4db](https://github.com/lithic-com/lithic-go/commit/741d4dbe323114d3b2cff9e7cafcbb2511fe272a))

## 0.53.0 (2024-09-24)

Full Changelog: [v0.52.0...v0.53.0](https://github.com/lithic-com/lithic-go/compare/v0.52.0...v0.53.0)

### Features

* **api:** add `ACCOUNT_DELINQUENT` to `detailed_results` enum ([#375](https://github.com/lithic-com/lithic-go/issues/375)) ([d0bb73f](https://github.com/lithic-com/lithic-go/commit/d0bb73f7f6cb6582d06c01c45132cc3d9cb397f9))

## 0.52.0 (2024-09-23)

Full Changelog: [v0.51.1...v0.52.0](https://github.com/lithic-com/lithic-go/compare/v0.51.1...v0.52.0)

### Features

* **api:** add `canceled` to `transaction_status` enum values ([#373](https://github.com/lithic-com/lithic-go/issues/373)) ([2628d65](https://github.com/lithic-com/lithic-go/commit/2628d65737c9587dfd592358d1c33c31f568387e))

## 0.51.1 (2024-09-20)

Full Changelog: [v0.51.0...v0.51.1](https://github.com/lithic-com/lithic-go/compare/v0.51.0...v0.51.1)

### Bug Fixes

* remove use of 'CreditProductService' type that doesn't exist ([6e4386d](https://github.com/lithic-com/lithic-go/commit/6e4386d8fa681501973209cc70e7b6519c667393))

## 0.51.0 (2024-09-20)

Full Changelog: [v0.50.1...v0.51.0](https://github.com/lithic-com/lithic-go/compare/v0.50.1...v0.51.0)

### ⚠ BREAKING CHANGES

* **api:** update model `FinancialAccount` ([#369](https://github.com/lithic-com/lithic-go/issues/369))

### Features

* **api:** update model `FinancialAccount` ([#369](https://github.com/lithic-com/lithic-go/issues/369)) ([d7d3d47](https://github.com/lithic-com/lithic-go/commit/d7d3d473ecf8ce1ebfea523d8a22b288c0e6494c))

## 0.50.1 (2024-09-18)

Full Changelog: [v0.50.0...v0.50.1](https://github.com/lithic-com/lithic-go/compare/v0.50.0...v0.50.1)

### Chores

* **internal:** specify API version for each endpoints instead of hardcoded in base URLs ([#366](https://github.com/lithic-com/lithic-go/issues/366)) ([4331c84](https://github.com/lithic-com/lithic-go/commit/4331c84d281155fd91b9d79b16bf8648ddb3e454))
* **tests:** fix GetEmbedURL to specify /v1/ prefix ([d6defc2](https://github.com/lithic-com/lithic-go/commit/d6defc2de23efe60003b7d67f202e6f170cfacf4))

## 0.50.0 (2024-09-17)

Full Changelog: [v0.49.1...v0.50.0](https://github.com/lithic-com/lithic-go/compare/v0.49.1...v0.50.0)

### ⚠ BREAKING CHANGES

* **api:** updates book transfer status, updates to transactions, add currency model ([#365](https://github.com/lithic-com/lithic-go/issues/365))

### Features

* **api:** updates book transfer status, updates to transactions, add currency model ([#365](https://github.com/lithic-com/lithic-go/issues/365)) ([611b4b2](https://github.com/lithic-com/lithic-go/commit/611b4b2617aa310f5146a06e7d7495cf52fbf31e))


### Documentation

* update CONTRIBUTING.md ([#363](https://github.com/lithic-com/lithic-go/issues/363)) ([0f35602](https://github.com/lithic-com/lithic-go/commit/0f3560234b1cbd9876960a682ecebeae89b3db00))

## 0.49.1 (2024-09-10)

Full Changelog: [v0.49.0...v0.49.1](https://github.com/lithic-com/lithic-go/compare/v0.49.0...v0.49.1)

### Bug Fixes

* **requestconfig:** copy over more fields when cloning ([#361](https://github.com/lithic-com/lithic-go/issues/361)) ([a205219](https://github.com/lithic-com/lithic-go/commit/a205219514f80ad9aff5f56d019a72ecff208b22))

## 0.49.0 (2024-09-06)

Full Changelog: [v0.48.0...v0.49.0](https://github.com/lithic-com/lithic-go/compare/v0.48.0...v0.49.0)

### Features

* **api:** add tier and state to financial_accounts ([#360](https://github.com/lithic-com/lithic-go/issues/360)) ([19ef005](https://github.com/lithic-com/lithic-go/commit/19ef0055948ebd66948615f1400598f48a711c04))


### Chores

* **docs:** update description for postal codes ([#358](https://github.com/lithic-com/lithic-go/issues/358)) ([b90dc4d](https://github.com/lithic-com/lithic-go/commit/b90dc4df890e71a3f87155bb6c358af25cc35fb1))

## 0.48.0 (2024-09-03)

Full Changelog: [v0.47.0...v0.48.0](https://github.com/lithic-com/lithic-go/compare/v0.47.0...v0.48.0)

### ⚠ BREAKING CHANGES

* **api:** add shared model Document ([#356](https://github.com/lithic-com/lithic-go/issues/356))

### Features

* **api:** add 'pin status' and 'pending_commands' to Card model ([#355](https://github.com/lithic-com/lithic-go/issues/355)) ([fca6d00](https://github.com/lithic-com/lithic-go/commit/fca6d00715bc0fd002d4cc4c9e1ea47518022dee))
* **api:** add shared model Document ([#356](https://github.com/lithic-com/lithic-go/issues/356)) ([a117730](https://github.com/lithic-com/lithic-go/commit/a117730913073cd0f76de5eee4806b6fb7e88fb0))
* **api:** declare AccountHolderBusinessResponse and remove entity_token from BusinessEntity ([#357](https://github.com/lithic-com/lithic-go/issues/357)) ([9dea351](https://github.com/lithic-com/lithic-go/commit/9dea3517db99c8c44f3e62a55c207dc1f8b21cc0))


### Chores

* **docs:** minor edits ([#353](https://github.com/lithic-com/lithic-go/issues/353)) ([481b68b](https://github.com/lithic-com/lithic-go/commit/481b68b2228b5456ede57a9bb4eb405359a128c7))

## 0.47.0 (2024-08-23)

Full Changelog: [v0.46.0...v0.47.0](https://github.com/lithic-com/lithic-go/compare/v0.46.0...v0.47.0)

### Features

* **api:** add endpoints and webhooks for 3DS challenge decisions and challenges ([#351](https://github.com/lithic-com/lithic-go/issues/351)) ([e1a87ef](https://github.com/lithic-com/lithic-go/commit/e1a87efebd4628add40792b07818c06526f4a652))

## 0.46.0 (2024-08-23)

Full Changelog: [v0.45.0...v0.46.0](https://github.com/lithic-com/lithic-go/compare/v0.45.0...v0.46.0)

### Features

* **api:** new book_transfer_transaction events and Settlement Report field deprecations ([#349](https://github.com/lithic-com/lithic-go/issues/349)) ([caf2767](https://github.com/lithic-com/lithic-go/commit/caf276753f491d11f98216c58ec4eb369ed5fdbd))

## 0.45.0 (2024-08-20)

Full Changelog: [v0.44.0...v0.45.0](https://github.com/lithic-com/lithic-go/compare/v0.44.0...v0.45.0)

### Features

* **api:** add property `next_payment_end_date` and `next_payment_due_date` to Statement model ([#348](https://github.com/lithic-com/lithic-go/issues/348)) ([02cef9a](https://github.com/lithic-com/lithic-go/commit/02cef9a8c8040130445cffb134fd2256ff3050aa))


### Chores

* **docs:** update state description on Card ([#346](https://github.com/lithic-com/lithic-go/issues/346)) ([77b0dfd](https://github.com/lithic-com/lithic-go/commit/77b0dfd8b1b27bb9bab2d31d9bdf96adc1dfd7a8))

## 0.44.0 (2024-08-16)

Full Changelog: [v0.43.0...v0.44.0](https://github.com/lithic-com/lithic-go/compare/v0.43.0...v0.44.0)

### Features

* **api:** add StatementListParams property include_initial_statements ([#344](https://github.com/lithic-com/lithic-go/issues/344)) ([422a2ce](https://github.com/lithic-com/lithic-go/commit/422a2ce1be046d7f148690da77a7aaa98b8b5bab))

## 0.43.0 (2024-08-14)

Full Changelog: [v0.42.0...v0.43.0](https://github.com/lithic-com/lithic-go/compare/v0.42.0...v0.43.0)

### Features

* **api:** add SettlementReport property `is_complete` ([#343](https://github.com/lithic-com/lithic-go/issues/343)) ([528dfbe](https://github.com/lithic-com/lithic-go/commit/528dfbed1f552567eff626800e922e7c9f6670d5))


### Chores

* bump Go to v1.21 ([#340](https://github.com/lithic-com/lithic-go/issues/340)) ([8ba9105](https://github.com/lithic-com/lithic-go/commit/8ba91052ca32f7994ec66fba999aa531a15b91a5))
* **examples:** minor formatting changes ([#342](https://github.com/lithic-com/lithic-go/issues/342)) ([6326899](https://github.com/lithic-com/lithic-go/commit/6326899066cbc2b562158c0625ee1ea274ddcf82))

## 0.42.0 (2024-08-12)

Full Changelog: [v0.41.1...v0.42.0](https://github.com/lithic-com/lithic-go/compare/v0.41.1...v0.42.0)

### Features

* **api:** add property `Account.CardholderCurrency` ([#338](https://github.com/lithic-com/lithic-go/issues/338)) ([7b97846](https://github.com/lithic-com/lithic-go/commit/7b97846987624352af944d09cf9c48aef4849bcd))
* **api:** add property `Card.CardholderCurrency` ([7b97846](https://github.com/lithic-com/lithic-go/commit/7b97846987624352af944d09cf9c48aef4849bcd))
* **api:** add property `CardProgram.CardholderCurrency` ([7b97846](https://github.com/lithic-com/lithic-go/commit/7b97846987624352af944d09cf9c48aef4849bcd))
* **api:** add property `CardProgram.SettlementCurrencies` ([7b97846](https://github.com/lithic-com/lithic-go/commit/7b97846987624352af944d09cf9c48aef4849bcd))

## 0.41.1 (2024-08-09)

Full Changelog: [v0.41.0...v0.41.1](https://github.com/lithic-com/lithic-go/compare/v0.41.0...v0.41.1)

### Bug Fixes

* deserialization of struct unions that implement json.Unmarshaler ([#337](https://github.com/lithic-com/lithic-go/issues/337)) ([416dea1](https://github.com/lithic-com/lithic-go/commit/416dea1c14352eb0284267d30768dfd861317a88))


### Chores

* **ci:** bump prism mock server version ([#335](https://github.com/lithic-com/lithic-go/issues/335)) ([c9ac55e](https://github.com/lithic-com/lithic-go/commit/c9ac55ec54ad23888bb5ffaf0f40c4f7ee4ea29d))

## 0.41.0 (2024-08-09)

Full Changelog: [v0.40.1...v0.41.0](https://github.com/lithic-com/lithic-go/compare/v0.40.1...v0.41.0)

### ⚠ BREAKING CHANGES

* **api:** rename model property 'StatementAccountStanding.State' to 'PeriodState' ([#334](https://github.com/lithic-com/lithic-go/issues/334))

### Features

* **api:** add event type 'card.reissued' ([#331](https://github.com/lithic-com/lithic-go/issues/331)) ([63bef3b](https://github.com/lithic-com/lithic-go/commit/63bef3b5b210c40e571d55a39eb56982064799af))
* **api:** add event type 'statements.created' ([#330](https://github.com/lithic-com/lithic-go/issues/330)) ([df398bc](https://github.com/lithic-com/lithic-go/commit/df398bc8be0f4e3e41d2dd8b13ef9882135b6c82))
* **api:** add methods to simulate enrollment review and enrollment document review ([#332](https://github.com/lithic-com/lithic-go/issues/332)) ([9265daf](https://github.com/lithic-com/lithic-go/commit/9265daf66c81ee07788a139504abae7705b18341))
* **api:** rename model property 'StatementAccountStanding.State' to 'PeriodState' ([#334](https://github.com/lithic-com/lithic-go/issues/334)) ([d4754a6](https://github.com/lithic-com/lithic-go/commit/d4754a6110b71c68f5de87a95b08335c43a2c55c))

## 0.40.1 (2024-07-31)

Full Changelog: [v0.40.0...v0.40.1](https://github.com/lithic-com/lithic-go/compare/v0.40.0...v0.40.1)

### Bug Fixes

* handle nil pagination responses when HTTP status is 200 ([#328](https://github.com/lithic-com/lithic-go/issues/328)) ([d53c62f](https://github.com/lithic-com/lithic-go/commit/d53c62f2d38181f6b40fe5b6accf04e406f78f07))

## 0.40.0 (2024-07-23)

Full Changelog: [v0.39.0...v0.40.0](https://github.com/lithic-com/lithic-go/compare/v0.39.0...v0.40.0)

### ⚠ BREAKING CHANGES

* **api:** deprecate 'auth rule token' in 'card' and 'account holder' models ([#325](https://github.com/lithic-com/lithic-go/issues/325))

### Features

* **api:** deprecate 'auth rule token' in 'card' and 'account holder' models ([#325](https://github.com/lithic-com/lithic-go/issues/325)) ([364e58b](https://github.com/lithic-com/lithic-go/commit/364e58bfefd38d388b04e010c6cade77ec98f1ad))


### Chores

* **ci:** remove unused release doctor ([#322](https://github.com/lithic-com/lithic-go/issues/322)) ([3a3fe06](https://github.com/lithic-com/lithic-go/commit/3a3fe06eb1f02d7d9c27a823c29489a759df3ced))
* **tests:** update prism version ([#324](https://github.com/lithic-com/lithic-go/issues/324)) ([2b65acd](https://github.com/lithic-com/lithic-go/commit/2b65acdb5ec1c6fa295326cf97f362b2f6502f59))

## 0.39.0 (2024-07-19)

Full Changelog: [v0.38.0...v0.39.0](https://github.com/lithic-com/lithic-go/compare/v0.38.0...v0.39.0)

### Features

* **api:** add method to retrieve a transaction's enhanced commercial data ([#321](https://github.com/lithic-com/lithic-go/issues/321)) ([0f46dd0](https://github.com/lithic-com/lithic-go/commit/0f46dd092e69eb46a7cdd8ff8a557f68ed2ae348))


### Chores

* **ci:** limit release doctor target branches ([#319](https://github.com/lithic-com/lithic-go/issues/319)) ([3254f11](https://github.com/lithic-com/lithic-go/commit/3254f1126beea5cf28b0ce503bc55108d1b201a8))

## 0.38.0 (2024-07-17)

Full Changelog: [v0.37.0...v0.38.0](https://github.com/lithic-com/lithic-go/compare/v0.37.0...v0.38.0)

### Features

* **api:** updates ([#317](https://github.com/lithic-com/lithic-go/issues/317)) ([ed54c59](https://github.com/lithic-com/lithic-go/commit/ed54c597883f6ca3a71b52e8b4267ff566cdfbac))

## 0.37.0 (2024-07-11)

Full Changelog: [v0.36.3...v0.37.0](https://github.com/lithic-com/lithic-go/compare/v0.36.3...v0.37.0)

### Features

* **api:** param 'FinancialAccountToken' for 'ExternalBankAccountService.New()' is now required ([#315](https://github.com/lithic-com/lithic-go/issues/315)) ([76303f0](https://github.com/lithic-com/lithic-go/commit/76303f0e8ff486d69356fe6e25feb3c1eb37e034))

## 0.36.3 (2024-07-11)

Full Changelog: [v0.36.2...v0.36.3](https://github.com/lithic-com/lithic-go/compare/v0.36.2...v0.36.3)

### Chores

* **ci:** also run workflows for PRs targeting `next` ([#312](https://github.com/lithic-com/lithic-go/issues/312)) ([7f78938](https://github.com/lithic-com/lithic-go/commit/7f789384e0d13aa8f752ac23cbc1073b85e9302d))


### Documentation

* **examples:** update example values ([#314](https://github.com/lithic-com/lithic-go/issues/314)) ([d998e57](https://github.com/lithic-com/lithic-go/commit/d998e579c22aa15fe92128070d70f7829782bab1))

## 0.36.2 (2024-07-07)

Full Changelog: [v0.36.1...v0.36.2](https://github.com/lithic-com/lithic-go/compare/v0.36.1...v0.36.2)

### Bug Fixes

* **internal:** fix MarshalJSON logic for interface elemnets ([#311](https://github.com/lithic-com/lithic-go/issues/311)) ([6663ba1](https://github.com/lithic-com/lithic-go/commit/6663ba166c8c44d5f49a8f29fc8580b7c3b6cb81))
* use slice instead of appending to r.Options ([#309](https://github.com/lithic-com/lithic-go/issues/309)) ([82edf46](https://github.com/lithic-com/lithic-go/commit/82edf46b49df41c5e1ce7bc9377c0fde3420b204))


### Chores

* gitignore test server logs ([#307](https://github.com/lithic-com/lithic-go/issues/307)) ([e0076c5](https://github.com/lithic-com/lithic-go/commit/e0076c5f9efcbf284f65bb5ea2cb6320b5f98779))
* **internal:** improve deserialization of embedded structs ([#310](https://github.com/lithic-com/lithic-go/issues/310)) ([5abc98f](https://github.com/lithic-com/lithic-go/commit/5abc98fcb88479a111a7e598edea265e0d262d9b))

## 0.36.1 (2024-06-21)

Full Changelog: [v0.36.0...v0.36.1](https://github.com/lithic-com/lithic-go/compare/v0.36.0...v0.36.1)

### Bug Fixes

* fix ExtraFields serialization / deserialization ([#305](https://github.com/lithic-com/lithic-go/issues/305)) ([7ccfaeb](https://github.com/lithic-com/lithic-go/commit/7ccfaeb5f53a5aa6a5b6da0707cd702eca57b5a2))

## 0.36.0 (2024-06-21)

Full Changelog: [v0.35.0...v0.36.0](https://github.com/lithic-com/lithic-go/compare/v0.35.0...v0.36.0)

### ⚠ BREAKING CHANGES

* **api:** remove unused event type 'statement.created'
* **api:** remove unused business account type
* **api:** remove unused embed request params type
* **api:** updates ([#303](https://github.com/lithic-com/lithic-go/issues/303))

### Features

* **api:** add 'reverse' method for book transfers ([ab7fa77](https://github.com/lithic-com/lithic-go/commit/ab7fa7774fd5b8691dd11bc28ad8a1d259a2802c))
* **api:** add field 'trace numbers' to payment method attribute model ([ab7fa77](https://github.com/lithic-com/lithic-go/commit/ab7fa7774fd5b8691dd11bc28ad8a1d259a2802c))
* **api:** remove unused business account type ([ab7fa77](https://github.com/lithic-com/lithic-go/commit/ab7fa7774fd5b8691dd11bc28ad8a1d259a2802c))
* **api:** remove unused embed request params type ([ab7fa77](https://github.com/lithic-com/lithic-go/commit/ab7fa7774fd5b8691dd11bc28ad8a1d259a2802c))
* **api:** remove unused event type 'statement.created' ([ab7fa77](https://github.com/lithic-com/lithic-go/commit/ab7fa7774fd5b8691dd11bc28ad8a1d259a2802c))
* **api:** updates ([#303](https://github.com/lithic-com/lithic-go/issues/303)) ([ab7fa77](https://github.com/lithic-com/lithic-go/commit/ab7fa7774fd5b8691dd11bc28ad8a1d259a2802c))

## 0.35.0 (2024-06-12)

Full Changelog: [v0.34.1...v0.35.0](https://github.com/lithic-com/lithic-go/compare/v0.34.1...v0.35.0)

### Features

* **api:** updates ([#301](https://github.com/lithic-com/lithic-go/issues/301)) ([f6b09d4](https://github.com/lithic-com/lithic-go/commit/f6b09d426a5beaba3c385cd0edef3018542fbc34))

## 0.34.1 (2024-06-06)

Full Changelog: [v0.34.0...v0.34.1](https://github.com/lithic-com/lithic-go/compare/v0.34.0...v0.34.1)

### Bug Fixes

* fix port function for interface{} types ([#299](https://github.com/lithic-com/lithic-go/issues/299)) ([6510ee8](https://github.com/lithic-com/lithic-go/commit/6510ee8a4a95381baf5d2fdbb41f77ec4ab41604))

## 0.34.0 (2024-06-05)

Full Changelog: [v0.33.1...v0.34.0](https://github.com/lithic-com/lithic-go/compare/v0.33.1...v0.34.0)

### ⚠ BREAKING CHANGES

* **api:** remove some endpoints and other API updates ([#297](https://github.com/lithic-com/lithic-go/issues/297))

### Features

* **api:** remove some endpoints and other API updates ([#297](https://github.com/lithic-com/lithic-go/issues/297)) ([0fdb299](https://github.com/lithic-com/lithic-go/commit/0fdb299059fba0542cc65ec4c156791d46b55852))

## 0.33.1 (2024-06-03)

Full Changelog: [v0.33.0...v0.33.1](https://github.com/lithic-com/lithic-go/compare/v0.33.0...v0.33.1)

### Bug Fixes

* **internal:** fix the way that unions are deserialized in nested arrays ([#295](https://github.com/lithic-com/lithic-go/issues/295)) ([42ba4ad](https://github.com/lithic-com/lithic-go/commit/42ba4ad8b02dfa4482ec0dbecae5a3a74a9c5c87))

## 0.33.0 (2024-05-30)

Full Changelog: [v0.32.0...v0.33.0](https://github.com/lithic-com/lithic-go/compare/v0.32.0...v0.33.0)

### Features

* **api:** update detailed_results enum values ([#293](https://github.com/lithic-com/lithic-go/issues/293)) ([e6196d0](https://github.com/lithic-com/lithic-go/commit/e6196d0e2e73bb6a05c40f41e1b8c84b63dd9f73))

## 0.32.0 (2024-05-29)

Full Changelog: [v0.31.0...v0.32.0](https://github.com/lithic-com/lithic-go/compare/v0.31.0...v0.32.0)

### Features

* **api:** add simulate_receipt and simulate_action endpoints ([#291](https://github.com/lithic-com/lithic-go/issues/291)) ([7280fe3](https://github.com/lithic-com/lithic-go/commit/7280fe31cd84fa05b5aeef4cadee31796e9eb70d))

## 0.31.0 (2024-05-29)

Full Changelog: [v0.30.0...v0.31.0](https://github.com/lithic-com/lithic-go/compare/v0.30.0...v0.31.0)

### Features

* **api:** updates ([#290](https://github.com/lithic-com/lithic-go/issues/290)) ([4ade761](https://github.com/lithic-com/lithic-go/commit/4ade761a4a1def031fe2ff457987e04425c4a82b))
* better validation of path params ([#287](https://github.com/lithic-com/lithic-go/issues/287)) ([bfbf2f5](https://github.com/lithic-com/lithic-go/commit/bfbf2f53a1247b4a0c4f4c290e5ffc9ce4e52eb7))


### Chores

* **internal:** fix format script ([#289](https://github.com/lithic-com/lithic-go/issues/289)) ([4a68229](https://github.com/lithic-com/lithic-go/commit/4a68229f455a96c7b555555af82a52ce04c1665c))
* **internal:** support parsing other json content types ([#284](https://github.com/lithic-com/lithic-go/issues/284)) ([7dd0761](https://github.com/lithic-com/lithic-go/commit/7dd0761e35c245ef806ad817a0688d5158d17435))
* **tests:** update some example values ([#286](https://github.com/lithic-com/lithic-go/issues/286)) ([d97dc10](https://github.com/lithic-com/lithic-go/commit/d97dc105214fd08b388ec984f19c981decb5734e))

## 0.30.0 (2024-05-15)

Full Changelog: [v0.29.0...v0.30.0](https://github.com/lithic-com/lithic-go/compare/v0.29.0...v0.30.0)

### Features

* propagate resource description field from stainless config to SDK docs ([#280](https://github.com/lithic-com/lithic-go/issues/280)) ([4f9d6c1](https://github.com/lithic-com/lithic-go/commit/4f9d6c1e0e23c946f8ca254cd65a8f9ccd33d177))


### Bug Fixes

* fix reading the error body more than once ([#283](https://github.com/lithic-com/lithic-go/issues/283)) ([301e35b](https://github.com/lithic-com/lithic-go/commit/301e35b5a1577f394b45f574df61ed6e6efb325d))


### Chores

* **docs:** add SECURITY.md ([#281](https://github.com/lithic-com/lithic-go/issues/281)) ([4fb5392](https://github.com/lithic-com/lithic-go/commit/4fb5392f86a7d0ea73a238919e0024dcc9afaed7))
* **internal:** add slightly better logging to scripts ([#282](https://github.com/lithic-com/lithic-go/issues/282)) ([a646ba8](https://github.com/lithic-com/lithic-go/commit/a646ba8cf8d57b417fa88ffa14886d0c8322111a))
* **internal:** fix bootstrap script ([#277](https://github.com/lithic-com/lithic-go/issues/277)) ([80006cf](https://github.com/lithic-com/lithic-go/commit/80006cf47d3a104a4fee0dfdbbb58ae0801374d5))

## 0.29.0 (2024-05-01)

Full Changelog: [v0.28.0...v0.29.0](https://github.com/lithic-com/lithic-go/compare/v0.28.0...v0.29.0)

### Features

* **api:** changes to balance-related return types and other API changes ([#272](https://github.com/lithic-com/lithic-go/issues/272)) ([8fd8a7c](https://github.com/lithic-com/lithic-go/commit/8fd8a7c2eed94a8a354206ca2dd6a1b0b45a81fd))
* **api:** updates ([#264](https://github.com/lithic-com/lithic-go/issues/264)) ([6442564](https://github.com/lithic-com/lithic-go/commit/644256495e9b85d33a38ef81a2d0c8b523a566e9))
* **api:** updates ([#269](https://github.com/lithic-com/lithic-go/issues/269)) ([0562c47](https://github.com/lithic-com/lithic-go/commit/0562c47d4c685ecf7d4c9aa3f3be5377026e144f))
* **api:** updates ([#276](https://github.com/lithic-com/lithic-go/issues/276)) ([fa52e34](https://github.com/lithic-com/lithic-go/commit/fa52e34c82997485dd5b33e8db023f471d89c86b))
* **option:** add option to provide a raw request body ([#267](https://github.com/lithic-com/lithic-go/issues/267)) ([071d6e6](https://github.com/lithic-com/lithic-go/commit/071d6e653ce9709c8ee04e7fc5afd0f65e0e932a))
* update model params behavior ([#262](https://github.com/lithic-com/lithic-go/issues/262)) ([05a3532](https://github.com/lithic-com/lithic-go/commit/05a3532bc919a455a960f9f7d233766b1062f283))


### Bug Fixes

* make shared package public ([#273](https://github.com/lithic-com/lithic-go/issues/273)) ([00a7799](https://github.com/lithic-com/lithic-go/commit/00a7799ffbff16f8f906651cd8ea774c716e0680))
* **test:** fix test github actions job ([#275](https://github.com/lithic-com/lithic-go/issues/275)) ([0e4ccdd](https://github.com/lithic-com/lithic-go/commit/0e4ccddebba350ed6431ff68427c7d286d2c3f49))


### Chores

* change test names ([#265](https://github.com/lithic-com/lithic-go/issues/265)) ([22a01b7](https://github.com/lithic-com/lithic-go/commit/22a01b74b9b89e07eaed009f6231565a68a5d0b9))
* **internal:** add scripts/test, scripts/mock and add ci job ([#274](https://github.com/lithic-com/lithic-go/issues/274)) ([13d5670](https://github.com/lithic-com/lithic-go/commit/13d5670f6a0e36aa6a457599302d6f04ab542768))
* **internal:** fix Port function for number and boolean enums ([#271](https://github.com/lithic-com/lithic-go/issues/271)) ([c37e9d6](https://github.com/lithic-com/lithic-go/commit/c37e9d63e2114a011ccfd34e9f5733ca03014e9c))
* **internal:** formatting ([#266](https://github.com/lithic-com/lithic-go/issues/266)) ([c453766](https://github.com/lithic-com/lithic-go/commit/c4537667071f9af0e98f12609d745081eebd0734))
* **internal:** use actions/checkout@v4 for codeflow ([#270](https://github.com/lithic-com/lithic-go/issues/270)) ([429a92f](https://github.com/lithic-com/lithic-go/commit/429a92fd2ef06c5e71268ec7cf964c7afac1d2a4))


### Build System

* configure UTF-8 locale in devcontainer ([#268](https://github.com/lithic-com/lithic-go/issues/268)) ([a34339a](https://github.com/lithic-com/lithic-go/commit/a34339a656e9f88244e06955df7bc3edd2649526))

## 0.28.0 (2024-04-05)

Full Changelog: [v0.27.0...v0.28.0](https://github.com/lithic-com/lithic-go/compare/v0.27.0...v0.28.0)

### Features

* **api:** add detailed result CARD_NOT_ACTIVATED ([#259](https://github.com/lithic-com/lithic-go/issues/259)) ([ffcb5e7](https://github.com/lithic-com/lithic-go/commit/ffcb5e794936921c69c09b4fbffe08524d49fee8))
* **api:** add event type digital_wallet.tokenization_two_factor_authentication_code_sent ([#257](https://github.com/lithic-com/lithic-go/issues/257)) ([787ff07](https://github.com/lithic-com/lithic-go/commit/787ff07d9b3d3b9c873cde5660e98ba32916c5e8))
* **api:** add params spend_limit and spend_velocity ([#258](https://github.com/lithic-com/lithic-go/issues/258)) ([df6010a](https://github.com/lithic-com/lithic-go/commit/df6010a9a40d5cf5f6647ca59967d51d3dd5dee5))
* **api:** add settlement_report.updated enum ([#248](https://github.com/lithic-com/lithic-go/issues/248)) ([4ffb933](https://github.com/lithic-com/lithic-go/commit/4ffb93337f866631f7f15fce98d48e2cb0831414))
* **api:** update financial transaction status enum ([#254](https://github.com/lithic-com/lithic-go/issues/254)) ([1f21e5b](https://github.com/lithic-com/lithic-go/commit/1f21e5b12b057eec4be1d9200e69520707ef4389))
* **api:** update link to encrypted PIN block docs ([#261](https://github.com/lithic-com/lithic-go/issues/261)) ([e2c18b3](https://github.com/lithic-com/lithic-go/commit/e2c18b3789fc877348fbe7611a108bf66ae788e7))
* **client:** implement raw requests methods on client ([#252](https://github.com/lithic-com/lithic-go/issues/252)) ([445089d](https://github.com/lithic-com/lithic-go/commit/445089d567cff012c9f1de5acf1a993954ca0242))


### Chores

* **internal:** implement Port function in apijson ([#260](https://github.com/lithic-com/lithic-go/issues/260)) ([f6b0ce2](https://github.com/lithic-com/lithic-go/commit/f6b0ce282a82e431c2749ded09457f5468426b0b))
* **internal:** move pagination types to pagination package ([#253](https://github.com/lithic-com/lithic-go/issues/253)) ([c9dcf15](https://github.com/lithic-com/lithic-go/commit/c9dcf1598bd15c33816157c8ca8ebc4edf435fab))
* **internal:** use a time zone less likely to conflict with the local one ([#256](https://github.com/lithic-com/lithic-go/issues/256)) ([d62266f](https://github.com/lithic-com/lithic-go/commit/d62266febe9f804d765e77647fc66dc0a506b1b5))


### Documentation

* fix typo in docstring for Null() ([#251](https://github.com/lithic-com/lithic-go/issues/251)) ([9787ff1](https://github.com/lithic-com/lithic-go/commit/9787ff10c0c2600dcc93d9e4c2def571324b8f75))
* **readme:** document file uploads ([#250](https://github.com/lithic-com/lithic-go/issues/250)) ([0d7e97c](https://github.com/lithic-com/lithic-go/commit/0d7e97c9ab557f68e226857010e36ed22e4aab15))

## 0.27.0 (2024-03-21)

Full Changelog: [v0.26.1...v0.27.0](https://github.com/lithic-com/lithic-go/compare/v0.26.1...v0.27.0)

### Features

* add IsKnown method to enums ([#244](https://github.com/lithic-com/lithic-go/issues/244)) ([40a5087](https://github.com/lithic-com/lithic-go/commit/40a5087aed7e1837963ec6b4ed9cf60cf753a730))
* **api:** adds closed state ([#247](https://github.com/lithic-com/lithic-go/issues/247)) ([4d2c127](https://github.com/lithic-com/lithic-go/commit/4d2c1272d29d85396a6761e7fae3948e8a3cce73))
* **api:** updates ([#246](https://github.com/lithic-com/lithic-go/issues/246)) ([57dab3a](https://github.com/lithic-com/lithic-go/commit/57dab3ad77824e4b13fff8efc7e2320f773993bc))
* set user-agent header by default when making requests ([#238](https://github.com/lithic-com/lithic-go/issues/238)) ([58da383](https://github.com/lithic-com/lithic-go/commit/58da3834006479f1b9fd1c4f8e26062957de61fc))


### Chores

* add back examples ([dc9c260](https://github.com/lithic-com/lithic-go/commit/dc9c260a686eb9ca92ed4400eb6b676a63140ba4))
* add back removed code ([c46da60](https://github.com/lithic-com/lithic-go/commit/c46da609cd7a9f2b514e9c4d2451824237eef986))
* **internal:** update generated pragma comment ([#243](https://github.com/lithic-com/lithic-go/issues/243)) ([4671964](https://github.com/lithic-com/lithic-go/commit/4671964bb0355d3b4c05f1a9d865f4967a402079))
* temporarily remove examples for migration ([0fe6406](https://github.com/lithic-com/lithic-go/commit/0fe64063882e32404db75192517c2f1528c96b24))
* temporarily remove various code as part of refactor ([#241](https://github.com/lithic-com/lithic-go/issues/241)) ([9766a3a](https://github.com/lithic-com/lithic-go/commit/9766a3a6839cdf63f55c5b9818c5e95796aa2474))


### Documentation

* fix typo in CONTRIBUTING.md ([#242](https://github.com/lithic-com/lithic-go/issues/242)) ([7d03703](https://github.com/lithic-com/lithic-go/commit/7d03703f3725f3516aeddef2d7bcf7efda851896))
* **readme:** consistent use of sentence case in headings ([#245](https://github.com/lithic-com/lithic-go/issues/245)) ([526f33f](https://github.com/lithic-com/lithic-go/commit/526f33fb3606261cc2accda2a6f0a45c9825e873))

## 0.26.1 (2024-03-12)

Full Changelog: [v0.26.0...v0.26.1](https://github.com/lithic-com/lithic-go/compare/v0.26.0...v0.26.1)

### Bug Fixes

* **client:** don't include ? in path unless necessary ([#236](https://github.com/lithic-com/lithic-go/issues/236)) ([c3ab277](https://github.com/lithic-com/lithic-go/commit/c3ab277a354e094bf3ad8f2d4cfa94ebbbf9704f))

## 0.26.0 (2024-03-12)

Full Changelog: [v0.25.0...v0.26.0](https://github.com/lithic-com/lithic-go/compare/v0.25.0...v0.26.0)

### Features

* implement public RawJSON of response structs ([#231](https://github.com/lithic-com/lithic-go/issues/231)) ([df76ab0](https://github.com/lithic-com/lithic-go/commit/df76ab0abebdda1ee03e85288d051f4eef23b670))


### Bug Fixes

* fix String() behavior of param.Field ([#235](https://github.com/lithic-com/lithic-go/issues/235)) ([df2d179](https://github.com/lithic-com/lithic-go/commit/df2d179c5a4b6899816cc45cd57ca826a8edfd0c))


### Chores

* **internal:** improve union deserialization logic ([#233](https://github.com/lithic-com/lithic-go/issues/233)) ([b969527](https://github.com/lithic-com/lithic-go/commit/b969527514f57b8ec7abff1da1d0d0ed50f5ab88))


### Documentation

* **contributing:** add a CONTRIBUTING.md ([#234](https://github.com/lithic-com/lithic-go/issues/234)) ([a3347c8](https://github.com/lithic-com/lithic-go/commit/a3347c8e42a0997070f56c87f9fabcc4b4cc8a7b))

## 0.25.0 (2024-02-29)

Full Changelog: [v0.24.0...v0.25.0](https://github.com/lithic-com/lithic-go/compare/v0.24.0...v0.25.0)

### Features

* **api:** create financial account and retry microdeposits endpoints ([#226](https://github.com/lithic-com/lithic-go/issues/226)) ([c52f602](https://github.com/lithic-com/lithic-go/commit/c52f602fbc00e252a693cf01306eebb707ef0278))
* **api:** tokenizations ([#228](https://github.com/lithic-com/lithic-go/issues/228)) ([957d974](https://github.com/lithic-com/lithic-go/commit/957d974fa6a613ed73bca32913a8dcd56de2f69c))
* **api:** update financial_account_type and documentation ([#225](https://github.com/lithic-com/lithic-go/issues/225)) ([e81a342](https://github.com/lithic-com/lithic-go/commit/e81a342fa9f1450b53ef6e4749405ef72e3072ed))
* **api:** updates ([#229](https://github.com/lithic-com/lithic-go/issues/229)) ([825ff7e](https://github.com/lithic-com/lithic-go/commit/825ff7e599dfa1b03b39e87ab74c7f64f6dcf9a1))


### Chores

* **internal:** bump timeout threshold in tests ([#221](https://github.com/lithic-com/lithic-go/issues/221)) ([87470fd](https://github.com/lithic-com/lithic-go/commit/87470fdbb0c4d8a25710a9d39a28cdd1e8c4bce5))
* **internal:** refactor release environment script ([#223](https://github.com/lithic-com/lithic-go/issues/223)) ([5dfab6d](https://github.com/lithic-com/lithic-go/commit/5dfab6d5879ec2c9e74738963918dece70eb82d2))
* **internal:** update deps ([#227](https://github.com/lithic-com/lithic-go/issues/227)) ([c9fc50f](https://github.com/lithic-com/lithic-go/commit/c9fc50f92bd758a2e5f699ad78715cbac130ce2e))


### Documentation

* **readme:** improve wording ([#230](https://github.com/lithic-com/lithic-go/issues/230)) ([dc0dceb](https://github.com/lithic-com/lithic-go/commit/dc0dcebcfed1775186543dac372f63c4ca3b43a8))

## 0.24.0 (2024-02-08)

Full Changelog: [v0.23.0...v0.24.0](https://github.com/lithic-com/lithic-go/compare/v0.23.0...v0.24.0)

### Features

* **api:** updates ([#218](https://github.com/lithic-com/lithic-go/issues/218)) ([88f852f](https://github.com/lithic-com/lithic-go/commit/88f852f31286177ef8ed19053cd8601b40778f76))


### Chores

* **tests:** add integration test for pagination ([#220](https://github.com/lithic-com/lithic-go/issues/220)) ([95776ad](https://github.com/lithic-com/lithic-go/commit/95776ad7a03481054d52a70ebb00c4d32a361148))

## 0.23.0 (2024-02-06)

Full Changelog: [v0.22.0...v0.23.0](https://github.com/lithic-com/lithic-go/compare/v0.22.0...v0.23.0)

### Features

* **api:** add `account_token` and `card_program_token` to `Card` ([#214](https://github.com/lithic-com/lithic-go/issues/214)) ([4b08af3](https://github.com/lithic-com/lithic-go/commit/4b08af38f6055c6f16f363ea11ac6789c13a64e4))


### Chores

* **interal:** make link to api.md relative ([#216](https://github.com/lithic-com/lithic-go/issues/216)) ([a8187ea](https://github.com/lithic-com/lithic-go/commit/a8187ea13cdda8aab9c26fdb1da55f70b41d8505))
* **internal:** adjust formatting ([#217](https://github.com/lithic-com/lithic-go/issues/217)) ([db8cb51](https://github.com/lithic-com/lithic-go/commit/db8cb512e118aa8d46de5b14c4a517ce2d71f350))

## 0.22.0 (2024-01-31)

Full Changelog: [v0.21.0...v0.22.0](https://github.com/lithic-com/lithic-go/compare/v0.21.0...v0.22.0)

### Features

* remove idempotency headers ([#213](https://github.com/lithic-com/lithic-go/issues/213)) ([6847554](https://github.com/lithic-com/lithic-go/commit/68475543a479ce56c8fadfcec99326949f16c04a))


### Chores

* **internal:** support pre-release versioning ([#211](https://github.com/lithic-com/lithic-go/issues/211)) ([4feaa28](https://github.com/lithic-com/lithic-go/commit/4feaa28aa0df2ab2bcb55e322eff4c570bcdef42))

## 0.21.0 (2024-01-29)

Full Changelog: [v0.20.0...v0.21.0](https://github.com/lithic-com/lithic-go/compare/v0.20.0...v0.21.0)

### Features

* **api:** add search_by_pan endpoint ([#208](https://github.com/lithic-com/lithic-go/issues/208)) ([468d95b](https://github.com/lithic-com/lithic-go/commit/468d95b6f654a47d2bf1dac018ac67d52632f798))


### Bug Fixes

* parse date-time strings more leniently ([#210](https://github.com/lithic-com/lithic-go/issues/210)) ([96b2359](https://github.com/lithic-com/lithic-go/commit/96b2359c878ac95fc7c5143febed0f1824c6ead0))

## 0.20.0 (2024-01-23)

Full Changelog: [v0.19.2...v0.20.0](https://github.com/lithic-com/lithic-go/compare/v0.19.2...v0.20.0)

### ⚠ BREAKING CHANGES

* **api:** change account holder creation response, new settlement detail type ([#207](https://github.com/lithic-com/lithic-go/issues/207))

### Features

* **api:** change account holder creation response, new settlement detail type ([#207](https://github.com/lithic-com/lithic-go/issues/207)) ([fe8b9f5](https://github.com/lithic-com/lithic-go/commit/fe8b9f541612fbc7a4ee0d299a47721516475e1f))


### Chores

* **ci:** rely on Stainless GitHub App for releases ([#205](https://github.com/lithic-com/lithic-go/issues/205)) ([a249b28](https://github.com/lithic-com/lithic-go/commit/a249b28e94adbc63ce68aaea25780b7753762e56))

## 0.19.2 (2024-01-17)

Full Changelog: [v0.19.1...v0.19.2](https://github.com/lithic-com/lithic-go/compare/v0.19.1...v0.19.2)

### Bug Fixes

* **test:** avoid test failures when SKIP_MOCK_TESTS is not set ([#204](https://github.com/lithic-com/lithic-go/issues/204)) ([e100ac3](https://github.com/lithic-com/lithic-go/commit/e100ac3662f42c6b9ea3e1058278b7dbf58ff3cd))


### Chores

* **internal:** speculative retry-after-ms support ([#203](https://github.com/lithic-com/lithic-go/issues/203)) ([aa805ee](https://github.com/lithic-com/lithic-go/commit/aa805eecb9d1c023d3128a00ddef9ed65cdb6ffc))

## 0.19.1 (2024-01-17)

Full Changelog: [v0.19.0...v0.19.1](https://github.com/lithic-com/lithic-go/compare/v0.19.0...v0.19.1)

### Features

* **api:** updates ([#201](https://github.com/lithic-com/lithic-go/issues/201)) ([0be0a85](https://github.com/lithic-com/lithic-go/commit/0be0a85d06c1d4ec377a96d47726b990d5ca42a7))


### Documentation

* **readme:** improve api reference ([#199](https://github.com/lithic-com/lithic-go/issues/199)) ([1cecf89](https://github.com/lithic-com/lithic-go/commit/1cecf89c28f7dd6e343ea5cc66273860862ee776))

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
