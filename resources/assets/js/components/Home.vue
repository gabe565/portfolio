<template>
    <header class="masthead">
        <transition name="fade" appear>
        <div class="intro-bg" :key="index" :style="{ backgroundImage: `url(${backgrounds[index]})` }" v-if="backgrounds"></div>
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
export default {
    data() {
        return {
            index: 0,
            backgrounds: null,
        }
    },
    created() {
        this.updateBackgrounds()
    },
    methods: {
        updateBackgrounds() {
            this.$Progress.start()
            axios.get('/api/bg')
                .then(response => {
                    this.backgrounds = response.data
                    this.randomBackground()
                    setInterval(() => this.randomBackground(), 7500)
                    this.$Progress.finish()
                })
                .catch(response => {
                    this.$Progress.fail()
                })
        },
        randomBackground() {
            this.index = Math.floor(Math.random() * this.backgrounds.length);
        }
    }
}
</script>
