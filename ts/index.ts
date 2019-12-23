'use strict';
import Vue from 'vue';
import Router from 'vue-router';
import *as Firebase from 'firebase/app';
const home = require('./compo/home.vue').default;
const login = require('./compo/login.vue').default;
const loginEmail = require('./compo/emailLogin.vue').default;
const emailCreate = require('./compo/emailCreate.vue').default;

const firebaseConfig = {
    apiKey: "AIzaSyAyvxB_NeY-vUxBXoxP8IZQBQG0FllUfOI",
    authDomain: "slope-f3cda.firebaseapp.com",
    projectId: "slope-f3cda",
    appId: "1:997689149698:web:3d95f4025eafd924fb5a1b",
    measurementId: "G-L2WGD42CGZ"
};

Firebase.initializeApp(firebaseConfig);

Vue.use(Router);

const routes = [
    {path: '/', component:home},
    {path: '/login', component:login},
    {path: '/login/email', component:loginEmail},
    {path: '/login/create', component:emailCreate}
];

const router:Router = new Router({
    routes,
});

new Vue({
    el:'#app',
    router: router
});