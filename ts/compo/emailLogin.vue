<template>
    <div>
        <input type="text" placeholder="Eメール" v-model="email">
        <input type="text" placeholder="パスワード" v-model="password">
        <button v-on:click="login">ログイン</button>
        <router-link to = "/login/create">アカウントの作成</router-link>
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