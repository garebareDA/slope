<template>
<div class="loginCenter">
    <div>
        <button class="button googleLogin" v-on:click="googleLogin">
            <div>Google Login</div>
        </button>
    </div>

    <div>
        <button class="button googleLogin" v-on:click="githubLogin">
            <div>GitHub Login</div>
        </button>
    </div>

    <div>
        <button class="button emailCreate" v-on:click="emailCreate">
            <div>Email Login</div>
        </button>
    </div>

    <div>
        <button class="button guestLogin" v-on:click="guestLogin">
            <div>Guest Login</div>
        </button>
    </div>
</div>
</template>

<script lang="ts">
import Vue from 'vue'
import firebase from 'firebase';
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

        emailCreate():void{
            this.$router.push("login/email/create");
        },

        emailLogin():void{
            this.$router.push("login/email/login");
        },

        guestLogin():void{
            firebase.auth().signInAnonymously().catch((err:ErrorHandler) => {
                if(!err){
                    this.$router.push("/");
                }
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