<template>
<div class="loginCenter" >
    <h1 class="title">
        Slope
    </h1>

    <h3 class="login">
        Slopeにログイン
    </h3>

    <div>
        <button class="button googleLogin" v-on:click="googleLogin">
            Googleでログイン
        </button>
    </div>

    <div>
        <button class="button githubLogin" v-on:click="githubLogin">
            GitHubでログイン
        </button>
    </div>

    <div>
        <button class="button emailLogin" v-on:click="emailLogin">
            Eメールでログイン
        </button>
    </div>

    <div>
        <button class="button guestLogin" v-on:click="guestLogin">
            ゲストでログイン
        </button>
    </div>

</div>
</template>

<script lang="ts">
import Vue from 'vue';
import * as firebase from 'firebase/app';
import 'firebase/auth';
import { ErrorHandler } from 'vue-router/types/router';

export default Vue.extend({
    methods:{
        googleLogin():void{
            const provider:firebase.auth.GoogleAuthProvider = new firebase.auth.GoogleAuthProvider();
            provider.addScope('https://www.googleapis.com/auth/contacts.readonly');
            providerLogin(provider);
        },

        githubLogin():void{
            const provider:firebase.auth.GithubAuthProvider = new firebase.auth.GithubAuthProvider();
            provider.addScope('repo');
            providerLogin(provider);
        },

        emailLogin():void{
            this.$router.push("login/email");
        },

        guestLogin():void{
            firebase.auth().signInAnonymously().then(() => {
                this.$router.push("/");
            }).catch((err:ErrorHandler) => {
                    alert(err);
            });
        }
    }
})

const _this:any = this;
function providerLogin(provider:firebase.auth.AuthProvider):void {
    firebase.auth().signInWithPopup(provider).then((result:any) => {
        console.log(result);
        _this.$router.push("/");
    }).catch((err:ErrorHandler) => {
        alert(err);
    });
}
</script>