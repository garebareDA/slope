<template>
    <div>
        <h1 class="title">slope</h1>
        <button class="postButton" v-on:click="openItem">+</button>
        <div class="modal" :class="{'is-open': isModalActive}">
            <div>
                <button class="closeButton" v-on:click="openItem">
                ×
                </button>
            </div>
            <textarea v-model="postText" class="textarea"></textarea>
            <div>
                <button v-on:click="post" class="post">投稿</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import axios, { AxiosError } from 'axios';
import * as firebase from 'firebase/app';
import 'firebase/auth';

export default Vue.extend({
    methods:{
        openItem() {
          this.toggleModal();
        },

        toggleModal() {
          this.isModalActive = ! this.isModalActive;
        },

        post(){
            const text:string = this.$data.postText;
            let _this:any = this;

            firebase.auth().currentUser!.getIdToken(true).then((idToken:any) => {

                axios.post('/postText',{
                    token:idToken,
                    text:text
                }).then(() =>{
                   _this.$data.text = "";
                }).catch((err:AxiosError) => {
                    alert(err);
                });

            }).catch((err:firebase.auth.Error) => {
                alert(err);
            });
        }
    },

    data(){
        return{
            isModalActive: false,
            postText: ""
        }
    }
})
</script>