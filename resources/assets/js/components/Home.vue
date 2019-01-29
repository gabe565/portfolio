<template>
    <header class="masthead">
        <transition name="fade" appear>
            <div class="intro-bg" :key="index" :style="{ backgroundImage: `url(${backgrounds[index]})` }" v-if="ready"></div>
        </transition>
        <div class="intro-body">
            <div class="container">
                <div class="row">
                    <div class="col-lg-8 mx-auto">
                        <h1 class="brand-heading">Gabe Cook</h1>
                        <h2 class="intro-text">Software Developer</h2>
                    </div>
                </div>
            </div>
        </div>
    </header>
</template>

<script>
import preloadImage from '../preloadImage'

export default {
    data() {
        return {
            ready: false,
            index: null,
            backgrounds: null,
        }
    },
    created() {
        this.$Progress.start()
        this.updateBackgrounds()
    },
    methods: {
        updateBackgrounds() {
            axios.get('/api/bg')
                .then(response => {
                    this.backgrounds = response.data
                    this.randomBackground()
                })
                .catch(response => {
                    this.$Progress.fail()
                })
        },
        randomBackground() {
            let index = Math.floor(Math.random() * this.backgrounds.length)
            preloadImage(this.backgrounds[index])
                .then(() => {
                    this.index = index
                    this.ready = true
                    this.$Progress.finish()
                    setTimeout(() => this.randomBackground(), 7500)
                })
                .catch(() => {
                    this.$Progress.fail()
                })
        }
    }
}
</script>
