import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'
import getters from './getter'
import { user } from "@/store/module/user"
import { router } from "@/store/module/router"
import tagsView from './module/tagsView'
Vue.use(Vuex)



const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
    modules: ['user']
})
export const store = new Vuex.Store({
    modules: {
        user,
        router,
        tagsView
    },
    getters,
    plugins: [vuexLocal.plugin]
})
