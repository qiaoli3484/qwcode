# Vue 3 + Vite

This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

## Recommended IDE Setup

- [VS Code](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).


//git中的SSL certificate problem: unable to get local issuer certificate错误的解决办法
//git config --system http.sslverify false
git config --global http.proxy 127.0.0.1:7890
git config --global https.proxy 127.0.0.1:7890
 git config --global --unset http.proxy
 git config --global --unset https.proxy