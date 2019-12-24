<template>
    <div class="loginCenter">

        <h1 class="title">Slope</h1>
        <h3 class="login">Eメールでログイン</h3>

        <div>
            <input type="text" placeholder="Eメール" v-model="email" class="loginText">
        </div>

        <div>
            <input type="password" placeholder="パスワード" v-model="password" class="loginText">
        </div>

        <div>
            <button v-on:click="login" class="button emailLogin">ログイン</button>
        </div>

        <div>
            <router-link to = "/login/create" >アカウントの作成</router-link>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import * as firebase from 'firebase/app';
import 'firebase/auth';
import { ErrorHandler } from 'vue-router/types/router';

export default Vue.extend({
    methods:{
        login():void{
            const data = this.$data;
            firebase.auth().signInWithEmailAndPassword(data.email, data.password).then(() =>{
                this.$router.push("/");
            }).catch((err:ErrorHandler) =>{
                alert(err);
            });
        }
    },

    data(){
        return{
            email:"",
            password:"",
        }
    }
})
</script>