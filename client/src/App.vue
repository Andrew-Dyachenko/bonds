<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import InputYear from './components/inputYear/inputYear.vue'
import HeaderComponent from './components/header/header.vue'
import Login from './components/login/login.vue'
</script>

<template>
  <HeaderComponent/>
  <img alt="Vue logo" src="./assets/duck.png" />
  <InputYear/>
<!--  <Login/>-->
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

