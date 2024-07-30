# Changelog

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
