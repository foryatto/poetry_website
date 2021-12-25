<template>
    <div>
        <div class="logo">
            中华诗词
        </div>
        <a-menu theme="light" mode="inline" :default-selected-keys="['11']" @click="showPoets">
                <a-menu-item :key="dynasties[0].id">
                    <a-icon type="read" />
                <span class="nav-text">{{ dynasties[0].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[1].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[1].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[2].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[2].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[3].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[3].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[4].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[4].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[5].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[5].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[6].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[6].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[7].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[7].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[8].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[8].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[9].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[9].name }}</span>
                </a-menu-item>
                <a-menu-item :key="dynasties[10].id">
                    <a-icon type="read" />
                    <span class="nav-text">{{ dynasties[10].name }}</span>
                </a-menu-item>
        </a-menu>
    </div>

</template>

<script>
    export default {
        data(){
            return {
                dynasties: {},
            }
        },
        methods:{
            getDynasties(){
                this.axios.get('/dynasties')
                .then(response => {
                    this.dynasties = response.data.data
                })
                .catch(error => console.log(error))
            },
            showPoets(obj){
                this.$router.push({path:"/poets", query:{dynasty_id: obj.key}});
            }
        },
        mounted(){
            this.getDynasties()
        }
    }
</script>

<style>
</style>