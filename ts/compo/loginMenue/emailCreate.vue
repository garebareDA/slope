<template>
    <div class="loginCenter">
        <h1 class="title">Slope</h1>
        <h3 class="login">アカウントの作成</h3>

        <div>
            <input type="text" placeholder="Eメール" v-model="email" class="loginText">
        </div>

        <div>
            <input type="password" placeholder="パスワード" v-model="password" class="loginText">
        </div>

        <div>
            <button v-on:click="create" class="button emailLogin">作成</button>
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
        create():void{
            const data = this.$data;
            firebase.auth().createUserWithEmailAndPassword(data.email, data.password).then(() =>{
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