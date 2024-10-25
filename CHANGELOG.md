# Changelog

## [0.4.6](https://github.com/kurtosis-tech/kardinal/compare/0.4.5...0.4.6) (2024-10-23)


### Bug Fixes

* typo ([#289](https://github.com/kurtosis-tech/kardinal/issues/289)) ([8ba658a](https://github.com/kurtosis-tech/kardinal/commit/8ba658af0b81e4433096dd5c29a7154f9894f9a8))

## [0.4.5](https://github.com/kurtosis-tech/kardinal/compare/0.4.4...0.4.5) (2024-10-18)


### Bug Fixes

* fix server gen go ([#283](https://github.com/kurtosis-tech/kardinal/issues/283)) ([5d52e7c](https://github.com/kurtosis-tech/kardinal/commit/5d52e7cf45303bda934cb9da9fecdf6f09160e64))

## [0.4.4](https://github.com/kurtosis-tech/kardinal/compare/0.4.3...0.4.4) (2024-10-18)


### Features

* [pt1.] version cmd ([#281](https://github.com/kurtosis-tech/kardinal/issues/281)) ([563e0a3](https://github.com/kurtosis-tech/kardinal/commit/563e0a363ab50f8e5a922d4eb19a790e0254d8f9))
* allow overriding env vars in flows ([#270](https://github.com/kurtosis-tech/kardinal/issues/270)) ([657df8a](https://github.com/kurtosis-tech/kardinal/commit/657df8afd0a2e809e465378281a4a192a0b90de1))

## [0.4.3](https://github.com/kurtosis-tech/kardinal/compare/0.4.2...0.4.3) (2024-10-14)


### Bug Fixes

* fix unsupported resources deployment ([#278](https://github.com/kurtosis-tech/kardinal/issues/278)) ([f278c09](https://github.com/kurtosis-tech/kardinal/commit/f278c09089ab0f37cdeb529839e5805f96655ac4))

## [0.4.2](https://github.com/kurtosis-tech/kardinal/compare/0.4.1...0.4.2) (2024-10-14)


### Features

* deploy unsupported Kubernetes resources ([#272](https://github.com/kurtosis-tech/kardinal/issues/272)) ([6c034d6](https://github.com/kurtosis-tech/kardinal/commit/6c034d694338d338b823da744f91c1b65b134f8d))

## [0.4.1](https://github.com/kurtosis-tech/kardinal/compare/0.4.0...0.4.1) (2024-10-03)


### Bug Fixes

* add auth for statefulset during manager deployment ([#268](https://github.com/kurtosis-tech/kardinal/issues/268)) ([9c011c6](https://github.com/kurtosis-tech/kardinal/commit/9c011c66f8315a5ada9183ac0b5873a4156f7246))

## [0.4.0](https://github.com/kurtosis-tech/kardinal/compare/0.3.3...0.4.0) (2024-10-03)


### ⚠ BREAKING CHANGES

* add support to StatefulSet ([#263](https://github.com/kurtosis-tech/kardinal/issues/263))

### Features

* add canonical tag to website pages ([#266](https://github.com/kurtosis-tech/kardinal/issues/266)) ([64bdab1](https://github.com/kurtosis-tech/kardinal/commit/64bdab125ed308ca5282f70f58b0d88949e3c124)), closes [#236](https://github.com/kurtosis-tech/kardinal/issues/236)
* add support to StatefulSet ([#263](https://github.com/kurtosis-tech/kardinal/issues/263)) ([a1632d5](https://github.com/kurtosis-tech/kardinal/commit/a1632d5aecd857f8de5b77db342c39daff91bcb2))


### Bug Fixes

* avoid trying to delete default namespace ([#265](https://github.com/kurtosis-tech/kardinal/issues/265)) ([907ad38](https://github.com/kurtosis-tech/kardinal/commit/907ad38cde24d7a7baac2b3b900187685be92149))

## [0.3.3](https://github.com/kurtosis-tech/kardinal/compare/0.3.2...0.3.3) (2024-10-02)


### Features

* add the `id` flag in the kardinal flow create cmd ([#261](https://github.com/kurtosis-tech/kardinal/issues/261)) ([09309f1](https://github.com/kurtosis-tech/kardinal/commit/09309f173e11c6d4d2b7b8a80f55df759fb98406))

## [0.3.2](https://github.com/kurtosis-tech/kardinal/compare/0.3.1...0.3.2) (2024-09-28)


### Bug Fixes

* add the --service flag in the kardinal flow telepresence intercept command ([#259](https://github.com/kurtosis-tech/kardinal/issues/259)) ([5d22282](https://github.com/kurtosis-tech/kardinal/commit/5d222824139e0b9597f94cb0b4cd0663e948e413))
* fix broken website CSS by refactoring styled-components SSR logic ([#257](https://github.com/kurtosis-tech/kardinal/issues/257)) ([505e885](https://github.com/kurtosis-tech/kardinal/commit/505e88504d020ed1c53ae6f4d018db85b86d1e1c))

## [0.3.1](https://github.com/kurtosis-tech/kardinal/compare/0.3.0...0.3.1) (2024-09-27)


### Bug Fixes

* fix race condition with docs voting feature ([#256](https://github.com/kurtosis-tech/kardinal/issues/256)) ([2d3f769](https://github.com/kurtosis-tech/kardinal/commit/2d3f769e39d5ae008f6b6b6ca25819cccd6847d2))
* remove all Kardinal namespaces when cluster resources is empty during cleanup ([#255](https://github.com/kurtosis-tech/kardinal/issues/255)) ([43c4ac2](https://github.com/kurtosis-tech/kardinal/commit/43c4ac25468f1284fb53b36bc20458af5e7b9409))
* telepresence http server check ([#251](https://github.com/kurtosis-tech/kardinal/issues/251)) ([c26f135](https://github.com/kurtosis-tech/kardinal/commit/c26f135f0656959a0b9d2dd823094fda4d652537))
* try making some changes to change the asset hashes ([#254](https://github.com/kurtosis-tech/kardinal/issues/254)) ([668be13](https://github.com/kurtosis-tech/kardinal/commit/668be13bf4d31dabfe4fe5fc95dc3014794b54c7))
* Update the manager deploy kontrol service location command argument ([#252](https://github.com/kurtosis-tech/kardinal/issues/252)) ([08be35c](https://github.com/kurtosis-tech/kardinal/commit/08be35c92e14178b833af2197fc0032764cb7703))

## [0.3.0](https://github.com/kurtosis-tech/kardinal/compare/0.2.7...0.3.0) (2024-09-26)


### ⚠ BREAKING CHANGES

* new gateway management ([#238](https://github.com/kurtosis-tech/kardinal/issues/238))

### Features

* E2E tests iteration [#2](https://github.com/kurtosis-tech/kardinal/issues/2) ([#237](https://github.com/kurtosis-tech/kardinal/issues/237)) ([c4f6b67](https://github.com/kurtosis-tech/kardinal/commit/c4f6b679b4b1169d22b5aa3fb6b5c716348df433))
* new gateway management ([#238](https://github.com/kurtosis-tech/kardinal/issues/238)) ([b5a9916](https://github.com/kurtosis-tech/kardinal/commit/b5a991623aebf5a49cbd96e740d873e493564f34))


### Bug Fixes

* add missing resources to policy ([#249](https://github.com/kurtosis-tech/kardinal/issues/249)) ([0ade319](https://github.com/kurtosis-tech/kardinal/commit/0ade31954906ff97b200e9e9a0f61382f31d0a3b))
* Clarify the manager CLI command deploy kontrol service location argument ([#247](https://github.com/kurtosis-tech/kardinal/issues/247)) ([54bdc21](https://github.com/kurtosis-tech/kardinal/commit/54bdc2144d7cacab7687d0ede777dd955db10df4))
* Do not alter the istio-system namespace ([#246](https://github.com/kurtosis-tech/kardinal/issues/246)) ([1c98304](https://github.com/kurtosis-tech/kardinal/commit/1c98304e93fcdd5c1db3d87ab0fe060ebbf5bfe7))
* Exit with code 1 upon errors in the CLI ([#243](https://github.com/kurtosis-tech/kardinal/issues/243)) ([63c640e](https://github.com/kurtosis-tech/kardinal/commit/63c640e8529dd6f8231ca01ee570695307782fbb))
* Handle gateway case where the baseline is missing or not labelled ([#245](https://github.com/kurtosis-tech/kardinal/issues/245)) ([b0b6bb2](https://github.com/kurtosis-tech/kardinal/commit/b0b6bb2f60f058528a51e31dc7b7cd9a146b84cc))

## [0.2.7](https://github.com/kurtosis-tech/kardinal/compare/0.2.6...0.2.7) (2024-09-20)


### Features

* Topology node list of versions containing flow id, image tag and baseline flag ([#228](https://github.com/kurtosis-tech/kardinal/issues/228)) ([b42fad2](https://github.com/kurtosis-tech/kardinal/commit/b42fad213b93d26ab93b8f31104f5ab4ac994aa9))


### Bug Fixes

* Display JSON 500 error message ([#235](https://github.com/kurtosis-tech/kardinal/issues/235)) ([1073cb5](https://github.com/kurtosis-tech/kardinal/commit/1073cb55c1fc5eff77d1b41028a5e95ec16b01df))
* using dynamic namespace for the gateway cmd ([#239](https://github.com/kurtosis-tech/kardinal/issues/239)) ([7615ce3](https://github.com/kurtosis-tech/kardinal/commit/7615ce35251484cace736a42cf53b4f18a4890f7))

## [0.2.6](https://github.com/kurtosis-tech/kardinal/compare/0.2.5...0.2.6) (2024-09-17)


### Features

* add flow telepresence-intercept CLI command ([#213](https://github.com/kurtosis-tech/kardinal/issues/213)) ([d5a75f6](https://github.com/kurtosis-tech/kardinal/commit/d5a75f64bec64551fff53ee852de6b8cfe29496b))
* Add health error response ([#224](https://github.com/kurtosis-tech/kardinal/issues/224)) ([d86eaa7](https://github.com/kurtosis-tech/kardinal/commit/d86eaa704e51de33c7cd82dd4ca802c5737c7b0d))
* add the `baseline` column in the `flow ls` CLI command ([#217](https://github.com/kurtosis-tech/kardinal/issues/217)) ([9360ef3](https://github.com/kurtosis-tech/kardinal/commit/9360ef3ad01305c05770e5637837b55dcbd6fc21))
* calculator track events when values change ([#215](https://github.com/kurtosis-tech/kardinal/issues/215)) ([7a16e40](https://github.com/kurtosis-tech/kardinal/commit/7a16e40141d02065033fcc71326e68468ad0e9f6))
* docs anchor links for headings ([#214](https://github.com/kurtosis-tech/kardinal/issues/214)) ([7fbb139](https://github.com/kurtosis-tech/kardinal/commit/7fbb13996c24edf4e75a525f61f189b62573fcd5))
* variant of calculator page with share section ([#207](https://github.com/kurtosis-tech/kardinal/issues/207)) ([f4242f2](https://github.com/kurtosis-tech/kardinal/commit/f4242f2f122c565fec55b67d592c39b826e8387f))
* website get started section ([#220](https://github.com/kurtosis-tech/kardinal/issues/220)) ([805189b](https://github.com/kurtosis-tech/kardinal/commit/805189b6a1ff2d4263c6682cad5fa9517a3c51c5))


### Bug Fixes

* calculator mobile styles ([#216](https://github.com/kurtosis-tech/kardinal/issues/216)) ([9aa6af7](https://github.com/kurtosis-tech/kardinal/commit/9aa6af7bfcfe6c340405924a27583f2268fb4418))
* fix some confusing copy ([#221](https://github.com/kurtosis-tech/kardinal/issues/221)) ([e831db5](https://github.com/kurtosis-tech/kardinal/commit/e831db545217ae7d9a969ccc7652fd98282a3c97))
* replace 'prod' default namespace with 'baseline' in the ci-e2e-tests.yml ([#227](https://github.com/kurtosis-tech/kardinal/issues/227)) ([c81b2d4](https://github.com/kurtosis-tech/kardinal/commit/c81b2d4d67d5abb19c4d7f7c039cb7fe24863846))

## [0.2.5](https://github.com/kurtosis-tech/kardinal/compare/0.2.4...0.2.5) (2024-09-11)


### Features

* calculator fixups and polish ([#200](https://github.com/kurtosis-tech/kardinal/issues/200)) ([a773cd9](https://github.com/kurtosis-tech/kardinal/commit/a773cd98c214d029ec0dbe06893de88451c5c881))
* cost calculator ([#191](https://github.com/kurtosis-tech/kardinal/issues/191)) ([45807a3](https://github.com/kurtosis-tech/kardinal/commit/45807a3431a27d64ebb5b1a4990c57dc66a36aff))
* E2E tests - iteration [#1](https://github.com/kurtosis-tech/kardinal/issues/1) ([#180](https://github.com/kurtosis-tech/kardinal/issues/180)) ([b65a916](https://github.com/kurtosis-tech/kardinal/commit/b65a916a3cac41737f8ceb78ef98f6616564c1fe))
* Support for tenant with no base cluster topology ([#199](https://github.com/kurtosis-tech/kardinal/issues/199)) ([9c91a82](https://github.com/kurtosis-tech/kardinal/commit/9c91a824fd534e4cc9d2f87fc95f9c59f450a23d))


### Bug Fixes

* typo in calculator ([#201](https://github.com/kurtosis-tech/kardinal/issues/201)) ([0f0cf9b](https://github.com/kurtosis-tech/kardinal/commit/0f0cf9b943d7ec29995d381d08019720a1d8726b))

## [0.2.4](https://github.com/kurtosis-tech/kardinal/compare/0.2.3...0.2.4) (2024-09-06)


### Features

* allow concurrent proxies ([#193](https://github.com/kurtosis-tech/kardinal/issues/193)) ([1e9e8ad](https://github.com/kurtosis-tech/kardinal/commit/1e9e8ad61469e7d979aa81484007e512ab5c0c01))

## [0.2.3](https://github.com/kurtosis-tech/kardinal/compare/0.2.2...0.2.3) (2024-09-05)


### Features

* accept json as template args ([#188](https://github.com/kurtosis-tech/kardinal/issues/188)) ([3890abf](https://github.com/kurtosis-tech/kardinal/commit/3890abf9aad9ec7f719b09ca5e2bc4f189e9c5f3))
* adding the `flow delete --all` cmd flag ([#175](https://github.com/kurtosis-tech/kardinal/issues/175)) ([96b1cca](https://github.com/kurtosis-tech/kardinal/commit/96b1cca947ef4e6c1a20d67e8110b202b7b48cad))
* print cluster topology manifest CLI cmd ([#172](https://github.com/kurtosis-tech/kardinal/issues/172)) ([83e6dc2](https://github.com/kurtosis-tech/kardinal/commit/83e6dc2a83feb489ebb5dc266857ba6e0d06c71b))
* replace email form with demo link ([#184](https://github.com/kurtosis-tech/kardinal/issues/184)) ([cd84864](https://github.com/kurtosis-tech/kardinal/commit/cd8486497d94e2f940e5a5dc1e826ea112c31553))
* update video animation on homepage ([#187](https://github.com/kurtosis-tech/kardinal/issues/187)) ([b0eb232](https://github.com/kurtosis-tech/kardinal/commit/b0eb232516ec8e82a6774b47b61dc5dc04d6bd36))


### Bug Fixes

* clean shell name ([#183](https://github.com/kurtosis-tech/kardinal/issues/183)) ([c07ceeb](https://github.com/kurtosis-tech/kardinal/commit/c07ceebfbbd2771bf05596ef742d5cc86d90cb89))
* Karninal manager manifest tmpl ([#178](https://github.com/kurtosis-tech/kardinal/issues/178)) ([3a81875](https://github.com/kurtosis-tech/kardinal/commit/3a81875fa7ad829207bfa9f19d28bec2844c54d5))
* keeping telepresence annotation ([#182](https://github.com/kurtosis-tech/kardinal/issues/182)) ([b997afc](https://github.com/kurtosis-tech/kardinal/commit/b997afc8b8e35e8c203aa4d01fb80585e43f31fe))

## [0.2.2](https://github.com/kurtosis-tech/kardinal/compare/0.2.1...0.2.2) (2024-08-27)


### Features

* adding the cluster resources manifest endpoint in Kontrol API spec and bindings ([#171](https://github.com/kurtosis-tech/kardinal/issues/171)) ([b4e7d71](https://github.com/kurtosis-tech/kardinal/commit/b4e7d7134bdae58f939a4e3813daeec40a27d879))
* report playground user ([#173](https://github.com/kurtosis-tech/kardinal/issues/173)) ([bb4b2fa](https://github.com/kurtosis-tech/kardinal/commit/bb4b2fa71a5fbe7e52d0cf093e9e06f8c0d8348e))
* Tenant show command ([#174](https://github.com/kurtosis-tech/kardinal/issues/174)) ([1faaa44](https://github.com/kurtosis-tech/kardinal/commit/1faaa447ae769858459aaa89ab3ec388e71c7752))


### Bug Fixes

* check if envoy filters are updated before starting gateway ([#168](https://github.com/kurtosis-tech/kardinal/issues/168)) ([7e79072](https://github.com/kurtosis-tech/kardinal/commit/7e790724894ae4f26b4bca92c19b09465eb1156f))

## [0.2.1](https://github.com/kurtosis-tech/kardinal/compare/0.2.0...0.2.1) (2024-08-20)


### Features

* parse namespace k8s objects ([#158](https://github.com/kurtosis-tech/kardinal/issues/158)) ([cd0cef7](https://github.com/kurtosis-tech/kardinal/commit/cd0cef72eee54c2eb08ee36230103c08383a41c2))

## [0.2.0](https://github.com/kurtosis-tech/kardinal/compare/0.1.12...0.2.0) (2024-08-16)


### ⚠ BREAKING CHANGES

* api definitions for templates ([#141](https://github.com/kurtosis-tech/kardinal/issues/141))

### Features

* add `IngressConfig` to api ([#151](https://github.com/kurtosis-tech/kardinal/issues/151)) ([467b3e4](https://github.com/kurtosis-tech/kardinal/commit/467b3e4abf1155217a367ff9515cd1f7e7f56842))
* api definitions for templates ([#141](https://github.com/kurtosis-tech/kardinal/issues/141)) ([c1c7687](https://github.com/kurtosis-tech/kardinal/commit/c1c76872dfbe3af887cd7be47b33e04aba46f38e))
* website changes for launch ([#152](https://github.com/kurtosis-tech/kardinal/issues/152)) ([43fd1af](https://github.com/kurtosis-tech/kardinal/commit/43fd1aff2c0d0207a3c03de5fd3e874f42253d5d))


### Bug Fixes

* update privacy policy link ([#156](https://github.com/kurtosis-tech/kardinal/issues/156)) ([cd88c1b](https://github.com/kurtosis-tech/kardinal/commit/cd88c1bd82822a8c4d81c7e1a45303cf80850ef2))

## [0.1.12](https://github.com/kurtosis-tech/kardinal/compare/0.1.11...0.1.12) (2024-08-09)


### Features

* add multi-service support ([#128](https://github.com/kurtosis-tech/kardinal/issues/128)) ([720d61e](https://github.com/kurtosis-tech/kardinal/commit/720d61e05b14b8db4a51bf9486c32c0cc73d90ab))
* add option to anonymously report install ([#133](https://github.com/kurtosis-tech/kardinal/issues/133)) ([48384f6](https://github.com/kurtosis-tech/kardinal/commit/48384f627db52c5fc18b03b017e46b6e03853613))


### Bug Fixes

* dont hide the nav on docs ([#134](https://github.com/kurtosis-tech/kardinal/issues/134)) ([0695a09](https://github.com/kurtosis-tech/kardinal/commit/0695a09baf4f032c94bee2b8f1a6afa9b3ec960f))

## [0.1.11](https://github.com/kurtosis-tech/kardinal/compare/0.1.10...0.1.11) (2024-08-08)


### Features

* new nav items and mobile nav ([#123](https://github.com/kurtosis-tech/kardinal/issues/123)) ([406d1d2](https://github.com/kurtosis-tech/kardinal/commit/406d1d2f794f70d164b92534bd2b924ecb764d45))


### Bug Fixes

* add install fallback version ([#114](https://github.com/kurtosis-tech/kardinal/issues/114)) ([3ddb011](https://github.com/kurtosis-tech/kardinal/commit/3ddb011f1ff621b5623d92b3730d43fa8326ff2f))

## [0.1.10](https://github.com/kurtosis-tech/kardinal/compare/0.1.9...0.1.10) (2024-08-08)


### Features

* copy button for code blocks in docs ([#103](https://github.com/kurtosis-tech/kardinal/issues/103)) ([be8bdde](https://github.com/kurtosis-tech/kardinal/commit/be8bddebca4cd5f19bf62258091433ae183df01f)), closes [#89](https://github.com/kurtosis-tech/kardinal/issues/89)
* match fonts and gradients to new blog theme ([#107](https://github.com/kurtosis-tech/kardinal/issues/107)) ([0938dc3](https://github.com/kurtosis-tech/kardinal/commit/0938dc308437934044dd8373108c395de1ade2f7))


### Bug Fixes

* allow external links in docs ([#101](https://github.com/kurtosis-tech/kardinal/issues/101)) ([701928f](https://github.com/kurtosis-tech/kardinal/commit/701928f0c0b9a0fd972e4c9b15f44f2f657f21b7)), closes [#87](https://github.com/kurtosis-tech/kardinal/issues/87)
* Improve CLI ([#82](https://github.com/kurtosis-tech/kardinal/issues/82)) ([42c1b69](https://github.com/kurtosis-tech/kardinal/commit/42c1b6960ad1a90c8c340002ab994f91b3085dca))

## [0.1.9](https://github.com/kurtosis-tech/kardinal/compare/0.1.8...0.1.9) (2024-08-06)


### Bug Fixes

* Add missing Kardinal CLI print newlines ([#96](https://github.com/kurtosis-tech/kardinal/issues/96)) ([6e318f3](https://github.com/kurtosis-tech/kardinal/commit/6e318f3d37acce0aa0c87787e4cb22fffdf199f3))
* made the gateway antifragile by asserting prod namespace is alive and healthy ([#100](https://github.com/kurtosis-tech/kardinal/issues/100)) ([642c75e](https://github.com/kurtosis-tech/kardinal/commit/642c75eb26aa229b0bf0fbe414e922d99f1d0897))
* open traffic configuration directly ([#97](https://github.com/kurtosis-tech/kardinal/issues/97)) ([9ffdd11](https://github.com/kurtosis-tech/kardinal/commit/9ffdd11d740cc087236d8af004216fc53450f786))

## [0.1.8](https://github.com/kurtosis-tech/kardinal/compare/0.1.7...0.1.8) (2024-08-06)


### Features

* Add versions to the topology Node object ([#77](https://github.com/kurtosis-tech/kardinal/issues/77)) ([e23afbf](https://github.com/kurtosis-tech/kardinal/commit/e23afbfca95a16508cf7244764ed511f5291d7ea))
* an inbuilt gateway to handle kubectl host proxying ([#91](https://github.com/kurtosis-tech/kardinal/issues/91)) ([d57503a](https://github.com/kurtosis-tech/kardinal/commit/d57503af53632c5d547c445e2ae4858cd2b87ed0))


### Bug Fixes

* docs metadata ([#79](https://github.com/kurtosis-tech/kardinal/issues/79)) ([de822b5](https://github.com/kurtosis-tech/kardinal/commit/de822b56529d9611f8c85be7084c9f10a39e0a33))
* host setup on CLI across kloud and dev envs ([#80](https://github.com/kurtosis-tech/kardinal/issues/80)) ([5fd4937](https://github.com/kurtosis-tech/kardinal/commit/5fd49371e1d26cb1e27790b80f475c5bbff684f2))

## [0.1.7](https://github.com/kurtosis-tech/kardinal/compare/0.1.6...0.1.7) (2024-08-02)


### Bug Fixes

* CLI link printing ([#75](https://github.com/kurtosis-tech/kardinal/issues/75)) ([3b0257d](https://github.com/kurtosis-tech/kardinal/commit/3b0257d2f4340c2df5a705c1828bdf4b16a29f5b))
* flow URL log line ([#73](https://github.com/kurtosis-tech/kardinal/issues/73)) ([26103d0](https://github.com/kurtosis-tech/kardinal/commit/26103d01b30cb78f05144619ae04601791a83569))

## [0.1.6](https://github.com/kurtosis-tech/kardinal/compare/0.1.5...0.1.6) (2024-08-02)


### Features

* allows multiple flows ([#68](https://github.com/kurtosis-tech/kardinal/issues/68)) ([ef41abb](https://github.com/kurtosis-tech/kardinal/commit/ef41abb2a820c02ad7be5072d026c0c41721c827))
* Annotate voting app k8s manifest ([#61](https://github.com/kurtosis-tech/kardinal/issues/61)) ([dc703cb](https://github.com/kurtosis-tech/kardinal/commit/dc703cb5d03e5c0817154c64d16597aa71cc2aa3))
* fetch and apply authorization policies and envoy filters ([#62](https://github.com/kurtosis-tech/kardinal/issues/62)) ([99829e1](https://github.com/kurtosis-tech/kardinal/commit/99829e11a2e97495e66485ed00181649a89736d3))
* improvements to cli ([#71](https://github.com/kurtosis-tech/kardinal/issues/71)) ([eb2e0c4](https://github.com/kurtosis-tech/kardinal/commit/eb2e0c4d81bf26c80b359b48d35bcafd7e0bb4da))
* using IfNotPresent pull policy for local Minikube Kardinal Manager ([#65](https://github.com/kurtosis-tech/kardinal/issues/65)) ([6d2a2ea](https://github.com/kurtosis-tech/kardinal/commit/6d2a2ea61030038947cd727666bbd61db073cde4))


### Bug Fixes

* check if envoy filters and auth policies are not nil ([#66](https://github.com/kurtosis-tech/kardinal/issues/66)) ([ef2f3e0](https://github.com/kurtosis-tech/kardinal/commit/ef2f3e041cd3e7f1ed411b14c2828caa5a804456))

## [0.1.5](https://github.com/kurtosis-tech/kardinal/compare/0.1.4...0.1.5) (2024-07-30)


### Bug Fixes

* use the right namespace for router/router-redis ([#59](https://github.com/kurtosis-tech/kardinal/issues/59)) ([959b970](https://github.com/kurtosis-tech/kardinal/commit/959b970046b790ce4bb1f9084a882c7c595596d9))

## [0.1.4](https://github.com/kurtosis-tech/kardinal/compare/0.1.3...0.1.4) (2024-07-30)


### Features

* added a dashboard cmd that opens the kardinal dashboard in the browser ([#36](https://github.com/kurtosis-tech/kardinal/issues/36)) ([741e2e5](https://github.com/kurtosis-tech/kardinal/commit/741e2e54f1a5676f58d0e1abb054fa7235fd3b64))
* added required types for tracing &lt;&gt; flows ([#51](https://github.com/kurtosis-tech/kardinal/issues/51)) ([c525e3f](https://github.com/kurtosis-tech/kardinal/commit/c525e3f01b1946db33e3c3fc07357d2170bf4901))
* deploy the trace router + redis when manager is deployed ([#53](https://github.com/kurtosis-tech/kardinal/issues/53)) ([d458d22](https://github.com/kurtosis-tech/kardinal/commit/d458d221681cd5ffbf9155f9b52e5d76dfba696d))
* More robust k8s manifest parsing ([#39](https://github.com/kurtosis-tech/kardinal/issues/39)) ([3a18e63](https://github.com/kurtosis-tech/kardinal/commit/3a18e631350c05fb2820d69bf525ea11cb8e3df1))
* move website to public repo ([#47](https://github.com/kurtosis-tech/kardinal/issues/47)) ([368f733](https://github.com/kurtosis-tech/kardinal/commit/368f73396604bc4a6aa0c68224418d53eb047f20))


### Bug Fixes

* better error message if dashboard doesnt open ([#38](https://github.com/kurtosis-tech/kardinal/issues/38)) ([ce8804b](https://github.com/kurtosis-tech/kardinal/commit/ce8804b25294e830a900c0d0f905dd262d08ad22))
* remove extra space in CTA heading ([#49](https://github.com/kurtosis-tech/kardinal/issues/49)) ([8b2fe9f](https://github.com/kurtosis-tech/kardinal/commit/8b2fe9fa7fca3f60a069b73e1b8d38b0c340df55))
* use right type for authorization policies ([#52](https://github.com/kurtosis-tech/kardinal/issues/52)) ([216e606](https://github.com/kurtosis-tech/kardinal/commit/216e60670e4d5eaa0a878067b70491177bd3bade))

## [0.1.3](https://github.com/kurtosis-tech/kardinal/compare/0.1.2...0.1.3) (2024-07-12)


### Features

* remove old urls ([#21](https://github.com/kurtosis-tech/kardinal/issues/21)) ([2ce038f](https://github.com/kurtosis-tech/kardinal/commit/2ce038f8e7fb16979998d88c84d871b76596197b))
* store Kontrol location ([#20](https://github.com/kurtosis-tech/kardinal/issues/20)) ([ad42d22](https://github.com/kurtosis-tech/kardinal/commit/ad42d22ef6b8c1332b9b0447f77b0af248b0948e))
* Take k8s manifest as input instead of docker compose ([#30](https://github.com/kurtosis-tech/kardinal/issues/30)) ([4bcda1f](https://github.com/kurtosis-tech/kardinal/commit/4bcda1fd70349664cfcd552bb6b4c924cd0ec601))
* use install link alias ([#28](https://github.com/kurtosis-tech/kardinal/issues/28)) ([40e9d75](https://github.com/kurtosis-tech/kardinal/commit/40e9d75a93f0309a6c89611818856d62f9f7ec71))


### Bug Fixes

* image for quick demo ([#32](https://github.com/kurtosis-tech/kardinal/issues/32)) ([76d3ff3](https://github.com/kurtosis-tech/kardinal/commit/76d3ff3a3178346ce5e25fd55d449032adfcd4bb))

## [0.1.2](https://github.com/kurtosis-tech/kardinal/compare/0.1.1...0.1.2) (2024-07-09)


### Bug Fixes

* upload clis ([#24](https://github.com/kurtosis-tech/kardinal/issues/24)) ([dfff91b](https://github.com/kurtosis-tech/kardinal/commit/dfff91b67770dcb99b456158188a6b62a7cc39d2))

## [0.1.1](https://github.com/kurtosis-tech/kardinal/compare/v0.1.0...0.1.1) (2024-07-08)


### Features

* add js and nix package to cli-kontrol-api ([#11](https://github.com/kurtosis-tech/kardinal/issues/11)) ([d007773](https://github.com/kurtosis-tech/kardinal/commit/d0077735972bccdc40e48cd1b03a9d7122429d73))
* add more info to docs ([#17](https://github.com/kurtosis-tech/kardinal/issues/17)) ([54400fb](https://github.com/kurtosis-tech/kardinal/commit/54400fb9d65ad9edb01b2e76edb6ef3680b79a29))
* added release please ([#6](https://github.com/kurtosis-tech/kardinal/issues/6)) ([7c38358](https://github.com/kurtosis-tech/kardinal/commit/7c3835830b2616bffb5bbe5a64166154a5953314))
* adding multi tenant support ([#8](https://github.com/kurtosis-tech/kardinal/issues/8)) ([75b736f](https://github.com/kurtosis-tech/kardinal/commit/75b736fe24adce93b5fff7ee7eef9f7f5ead461c))
* adding the `kardinal manager deploy` and `kardinal manager remove` CLI commands ([#10](https://github.com/kurtosis-tech/kardinal/issues/10)) ([af8843c](https://github.com/kurtosis-tech/kardinal/commit/af8843cc3f08f8151741465b1e678b6ec76e80f6))
* adding the health check endpoint for the API ([#14](https://github.com/kurtosis-tech/kardinal/issues/14)) ([15a1b0d](https://github.com/kurtosis-tech/kardinal/commit/15a1b0d2bfb89c01e960657f7209cd1abf59b6c1))
* DRY ref to main branch in CI ([#12](https://github.com/kurtosis-tech/kardinal/issues/12)) ([f32ef8c](https://github.com/kurtosis-tech/kardinal/commit/f32ef8c3d04d24f2aeacec41d89bf627d9221c89))
* publish cli binaries and add install script ([#22](https://github.com/kurtosis-tech/kardinal/issues/22)) ([e1e27e4](https://github.com/kurtosis-tech/kardinal/commit/e1e27e4d0511ac613f66fb71b84cd135cf55e4c5))
* quickstart docs added ([#16](https://github.com/kurtosis-tech/kardinal/issues/16)) ([9f1a759](https://github.com/kurtosis-tech/kardinal/commit/9f1a759a7f1d50778f0a5ee6f17c843566f18356))
* removing kardinal-cli binary ([#13](https://github.com/kurtosis-tech/kardinal/issues/13)) ([31325f1](https://github.com/kurtosis-tech/kardinal/commit/31325f144d60802968572b71e8e997eecd391c52))
* update topology API  ([#9](https://github.com/kurtosis-tech/kardinal/issues/9)) ([2e96ff8](https://github.com/kurtosis-tech/kardinal/commit/2e96ff810e4a05ef1e7eb0ed52f6cff07c6fde4a))


### Bug Fixes

* fix some typos in README.md ([#23](https://github.com/kurtosis-tech/kardinal/issues/23)) ([2be7f55](https://github.com/kurtosis-tech/kardinal/commit/2be7f5544559148448daef072b9b39cf3c3ba403))
* update images pulls ([#18](https://github.com/kurtosis-tech/kardinal/issues/18)) ([212cfe8](https://github.com/kurtosis-tech/kardinal/commit/212cfe810f9c4bc21a25f9ac19751768cf6fcda4))
