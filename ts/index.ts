'use strict';
import Vue from 'vue';
import Router from 'vue-router';
import VModal from 'vue-js-modal';
import *as Firebase from 'firebase/app';

const home = require('./compo/home.vue').default;
const login = require('./compo/loginMenue/login.vue').default;
const loginEmail = require('./compo/loginMenue/emailLogin.vue').default;
const emailCreate = require('./compo/loginMenue/emailCreate.vue').default;
const setting = require('./compo/setting/userSetting.vue').default;

const firebaseConfig = {
    apiKey: "AIzaSyAyvxB_NeY-vUxBXoxP8IZQBQG0FllUfOI",
    authDomain: "slope-f3cda.firebaseapp.com",
    projectId: "slope-f3cda",
    appId: "1:997689149698:web:3d95f4025eafd924fb5a1b",
    measurementId: "G-L2WGD42CGZ"
};

Firebase.initializeApp(firebaseConfig);

Vue.use(Router);
Vue.use(VModal);

const routes = [
    {path: '/', component:home},
    {path: '/login', component:login},
    {path: '/login/email', component:loginEmail},
    {path: '/login/create', component:emailCreate},
    {path:'/setting', component:setting}
];

const router:Router = new Router({
    routes,
});

new Vue({
    el:'#app',
    router: router
});